// Package billupload provides the higher-level bill-upload workflows that are
// implemented by unified-go-client rather than the OpenAPI operation set.
package billupload

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

// NewPushCmd builds the bill-upload push workflow command.
func NewPushCmd() *cobra.Command {
	var (
		billConnectID string
		billingPeriod string
		files         []string
	)
	c := &cobra.Command{
		Use:   "push",
		Short: "Create, upload, and commit a bill upload",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			input := flexera.BillUploadPushInput{
				OrgID:         deps.Config.OrgID,
				BillConnectID: billConnectID,
				BillingPeriod: billingPeriod,
				Files:         files,
			}
			output, err := flexera.NewBillUploadPushTool(client.ClientInterface).Invoke(cmd.Context(), input)
			if err != nil {
				return err
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, output)
		},
	}
	c.Flags().StringVar(&billConnectID, "bill-connect-id", "", "bill connect ID (required)")
	c.Flags().StringVar(&billingPeriod, "billing-period", "", "billing period in yyyy-mm format (required)")
	c.Flags().StringSliceVar(&files, "file", nil, "local bill file to upload (repeatable, required)")
	return c
}

// NewVerifyCmd builds the bill-upload verify workflow command. Verification
// is a purely local, offline CSV check: it never makes network calls.
func NewVerifyCmd() *cobra.Command {
	var files []string
	c := &cobra.Command{
		Use:   "verify",
		Short: "Verify local CBI bill-upload CSV file(s) before uploading",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if len(files) == 0 {
				return errors.New("at least one --file is required")
			}
			deps := clipkg.DepsFrom(cmd.Context())

			results := make(map[string]flexera.BillUploadVerifyResult, len(files))
			invalid := false
			for _, path := range files {
				result, err := flexera.VerifyCBIBillUploadCSVFile(path)
				if err != nil {
					return fmt.Errorf("verify %q: %w", path, err)
				}
				if !result.Valid {
					invalid = true
				}
				results[path] = result
			}

			if err := deps.Printer.Render(deps.Stdout, deps.Config.Output, results); err != nil {
				return err
			}
			if invalid {
				return fmt.Errorf("one or more bill-upload CSV files failed verification")
			}
			return nil
		},
	}
	c.Flags().StringSliceVar(&files, "file", nil, "local CSV file to verify (repeatable, required)")
	return c
}

// Attach adds the curated push and verify workflows below the generated
// bill-upload command, preserving the generated direct endpoint commands.
func Attach(root *cobra.Command) error {
	for _, command := range root.Commands() {
		if command.Name() == "bill-upload" {
			command.AddCommand(NewPushCmd())
			command.AddCommand(NewVerifyCmd())
			return nil
		}
	}
	return fmt.Errorf("bill-upload generated command is not registered")
}
