// Package fnms adapts report jobs and redirect-backed binary downloads.
package fnms

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
	"github.com/spf13/cobra"
)

func Attach(root *cobra.Command) error {
	if err := clipkg.AttachCommands(root, []string{"fnms"}, fileGroup()); err != nil {
		return err
	}
	if err := clipkg.AttachCommands(root, []string{"fnms", "report", "execute-async"}, reportCmd("create", "Fnms_Report_execute_async", false, false)); err != nil {
		return err
	}
	if err := clipkg.AttachCommands(root, []string{"fnms", "report", "exports"}, reportCmd("create", "Fnms_Report_export_report", true, false)); err != nil {
		return err
	}
	return clipkg.AttachCommands(root, []string{"fnms", "report"}, reportCmd("execute", "", false, true), reportCmd("export", "", true, true), waitCmd(false), waitCmd(true), downloadCmd("download-export", ""))
}

func fileGroup() *cobra.Command {
	c := &cobra.Command{Use: "file", Short: "FNMS files"}
	c.AddCommand(downloadCmd("get", "Fnms_File_show"))
	return c
}

func annotations(id, output string, readOnly bool) map[string]string {
	a := map[string]string{clipkg.OutputAnnotation: output}
	if id != "" {
		a["flexera.operationId"] = id
	}
	if readOnly {
		a["flexera.readOnly"] = "true"
	}
	return a
}

func pollFlags(c *cobra.Command, options *flexera.FnmsPollOptions) {
	c.Flags().DurationVar(&options.Interval, "poll-interval", 2*time.Second, "interval between status requests")
	c.Flags().DurationVar(&options.Timeout, "poll-timeout", 5*time.Minute, "maximum time to wait")
}
func validatePoll(options flexera.FnmsPollOptions) error {
	if options.Interval <= 0 || options.Timeout <= 0 {
		return clipkg.Exit(2, fmt.Errorf("--poll-interval and --poll-timeout must be positive"))
	}
	return nil
}

func reportCmd(name, id string, export, wait bool) *cobra.Command {
	var reportID, search, filters, accept string
	var limit int64
	var force, dryRun, yes bool
	var options flexera.FnmsPollOptions
	short := "Start an asynchronous FNMS report execution"
	if export {
		short = "Start an FNMS report export"
	}
	if wait {
		short += " and wait for completion"
	}
	c := &cobra.Command{Use: name, Short: short, Args: cobra.NoArgs, Annotations: annotations(id, "structured", false)}
	c.RunE = func(cmd *cobra.Command, _ []string) error {
		deps := clipkg.DepsFrom(cmd.Context())
		if err := deps.Config.RequireOrgID(); err != nil {
			return err
		}
		var object map[string]interface{}
		if cmd.Flags().Changed("filter-params") {
			if err := json.Unmarshal([]byte(filters), &object); err != nil || object == nil {
				return clipkg.Exit(2, fmt.Errorf("--filter-params must be a JSON object"))
			}
		}
		if wait {
			if err := validatePoll(options); err != nil {
				return err
			}
		}
		if !export && cmd.Flags().Changed("limit") && (limit < 5000 || limit > 10000) {
			return clipkg.Exit(2, fmt.Errorf("--limit must be between 5000 and 10000"))
		}
		if export && accept != "application/json" && accept != "text/csv" {
			return clipkg.Exit(2, fmt.Errorf("--accept must be application/json or text/csv"))
		}
		params := map[string]any{"id": reportID}
		if object != nil {
			params["filterParams"] = object
		}
		if cmd.Flags().Changed("search-text") {
			params["searchText"] = search
		}
		path := fmt.Sprintf("/fnms/v1/orgs/%d/reports/%s/execute-async", deps.Config.OrgID, reportID)
		if export {
			path = fmt.Sprintf("/fnms/v1/orgs/%d/reports/%s/exports", deps.Config.OrgID, reportID)
			params["Accept"] = accept
		} else {
			if cmd.Flags().Changed("limit") {
				params["limit"] = limit
			}
			if cmd.Flags().Changed("force-new-execution") {
				params["forceNewExecution"] = force
			}
		}
		done, err := clipkg.ConfirmPlan(dryRun, yes, deps.Stdout, clipkg.Plan{Command: cmd.CommandPath(), Method: "POST", Path: path, OrgID: deps.Config.OrgID, Params: params}, deps.Printer)
		if err != nil || done {
			return err
		}
		client, err := deps.Factory().NewFNMSClient(deps.Config)
		if err != nil {
			return err
		}
		var output any
		if export {
			p := &flexera.FnmsReportExportReportParams{}
			if object != nil {
				p.FilterParams = &object
			}
			if cmd.Flags().Changed("search-text") {
				p.SearchText = &search
			}
			format := flexera.FnmsReportExportReportParamsAccept(accept)
			p.Accept = &format
			if wait {
				output, err = client.FnmsExportReport(cmd.Context(), int64(deps.Config.OrgID), reportID, p, &options)
			} else {
				var exportID string
				exportID, err = client.FnmsStartReportExport(cmd.Context(), int64(deps.Config.OrgID), reportID, p)
				output = map[string]string{"exportId": exportID}
			}
		} else {
			p := &flexera.FnmsReportExecuteAsyncParams{}
			if object != nil {
				p.FilterParams = &object
			}
			if cmd.Flags().Changed("search-text") {
				p.SearchText = &search
			}
			if cmd.Flags().Changed("limit") {
				p.Limit = &limit
			}
			if cmd.Flags().Changed("force-new-execution") {
				p.ForceNewExecution = &force
			}
			if wait {
				output, err = client.FnmsExecuteReport(cmd.Context(), int64(deps.Config.OrgID), reportID, p, &options)
			} else {
				var jobID string
				jobID, err = client.FnmsStartReportExecution(cmd.Context(), int64(deps.Config.OrgID), reportID, p)
				output = map[string]string{"jobId": jobID}
			}
		}
		if err != nil {
			return err
		}
		return deps.Printer.Render(deps.Stdout, deps.Config.Output, output)
	}
	c.Flags().StringVar(&reportID, "id", "", "saved report ID (required)")
	_ = c.MarkFlagRequired("id")
	c.Flags().StringVar(&search, "search-text", "", "text to filter report results")
	c.Flags().StringVar(&filters, "filter-params", "", "report filters as a JSON object (deepObject query parameters)")
	c.Flags().BoolVar(&dryRun, "dry-run", false, "preview without making API requests")
	c.Flags().BoolVar(&yes, "yes", false, "confirm operation (not required for non-destructive report jobs)")
	if export {
		c.Flags().StringVar(&accept, "accept", "application/json", "export format: application/json or text/csv")
	} else {
		c.Flags().Int64Var(&limit, "limit", 10000, "rows per result page (5000–10000)")
		c.Flags().BoolVar(&force, "force-new-execution", false, "skip job deduplication")
	}
	if wait {
		pollFlags(c, &options)
	}
	return c
}

func waitCmd(export bool) *cobra.Command {
	var reportID, jobID, exportID string
	var options flexera.FnmsPollOptions
	name, short := "wait-execution", "Wait for an existing report execution (does not retrieve result pages)"
	if export {
		name, short = "wait-export", "Wait for an existing report export and return its download path"
	}
	c := &cobra.Command{Use: name, Short: short, Args: cobra.NoArgs, Annotations: annotations("", "structured", true)}
	c.RunE = func(cmd *cobra.Command, _ []string) error {
		deps := clipkg.DepsFrom(cmd.Context())
		if err := deps.Config.RequireOrgID(); err != nil {
			return err
		}
		if err := validatePoll(options); err != nil {
			return err
		}
		client, err := deps.Factory().NewFNMSClient(deps.Config)
		if err != nil {
			return err
		}
		var output any
		if export {
			output, err = client.FnmsWaitReportExport(cmd.Context(), int64(deps.Config.OrgID), exportID, &options)
		} else {
			output, err = client.FnmsWaitReportExecution(cmd.Context(), int64(deps.Config.OrgID), reportID, jobID, &options)
		}
		if err != nil {
			return err
		}
		return deps.Printer.Render(deps.Stdout, deps.Config.Output, output)
	}
	if export {
		c.Flags().StringVar(&exportID, "export-id", "", "export ID (required)")
		_ = c.MarkFlagRequired("export-id")
	} else {
		c.Flags().StringVar(&reportID, "id", "", "saved report ID (required)")
		_ = c.MarkFlagRequired("id")
		c.Flags().StringVar(&jobID, "job-id", "", "execution job ID (required)")
		_ = c.MarkFlagRequired("job-id")
	}
	pollFlags(c, &options)
	return c
}

func downloadCmd(name, id string) *cobra.Command {
	var target, destination string
	c := &cobra.Command{Use: name, Short: "Stream FNMS download bytes unchanged to stdout or a new file", Args: cobra.NoArgs, Annotations: annotations(id, "binary", true)}
	c.RunE = func(cmd *cobra.Command, _ []string) error {
		deps := clipkg.DepsFrom(cmd.Context())
		if id != "" {
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
		}
		client, err := deps.Factory().NewFNMSClient(deps.Config)
		if err != nil {
			return err
		}
		var dst io.Writer = deps.Stdout
		var file *os.File
		if destination != "-" {
			file, err = os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
			if err != nil {
				return err
			}
			dst = file
		}
		var count int64
		if id != "" {
			count, err = client.FnmsDownloadFile(cmd.Context(), int64(deps.Config.OrgID), target, dst, nil)
		} else {
			count, err = client.FnmsDownloadReportExport(cmd.Context(), target, dst, nil)
		}
		if file != nil {
			closeErr := file.Close()
			if err == nil {
				err = closeErr
			}
			if err != nil {
				_ = os.Remove(destination)
			}
		}
		if err != nil {
			return fmt.Errorf("FNMS download after %d bytes: %w", count, err)
		}
		return nil
	}
	flag := "path"
	if id != "" {
		flag = "filename"
	}
	c.Flags().StringVar(&target, flag, "", "FNMS download target (required)")
	_ = c.MarkFlagRequired(flag)
	c.Flags().StringVar(&destination, "destination", "-", "write raw bytes to a new file, or - for stdout; existing files are never overwritten")
	return c
}
