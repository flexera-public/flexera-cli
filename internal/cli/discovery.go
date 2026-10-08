package cli

import (
	"fmt"
	"strings"
	"sync"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/spf13/cobra"
)

// NewDiscoveryCmd uses the existing tree; it never constructs another root.
func NewDiscoveryCmd(root *cobra.Command) *cobra.Command {
	group := &cobra.Command{Use: "cli", Short: "Discover CLI commands and API schemas", Annotations: map[string]string{"flexera.meta": "true"}}
	group.AddCommand(newSchemaCmd(root), newSearchCmd(root))
	return group
}

func newSearchCmd(root *cobra.Command) *cobra.Command {
	var options catalog.SearchOptions
	var once sync.Once
	var index *catalog.SearchIndex
	var buildErr error
	cmd := &cobra.Command{Use: "search <words...>", Short: "Find commands for a task", Long: "Search registered commands offline using deterministic lexical ranking. Usage values are synopses, not copy-ready commands: replace uppercase placeholders with your own values. Read-only excludes unclassified curated side effects; it is not a guarantee against normal authentication or output-file activity.", Annotations: map[string]string{"flexera.offline": "true", "flexera.readOnly": "true"}, Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || strings.TrimSpace(strings.Join(args, " ")) == "" {
			return Exit(2, fmt.Errorf("provide words describing the task to search"))
		}
		return nil
	}, RunE: func(cmd *cobra.Command, args []string) error {
		if options.Limit < 1 {
			return Exit(2, fmt.Errorf("--limit must be positive"))
		}
		once.Do(func() { index, buildErr = catalog.NewSearchIndex(root) })
		if buildErr != nil {
			return buildErr
		}
		results := index.Search(strings.Join(args, " "), options)
		deps := DepsFrom(cmd.Context())
		return deps.Printer.Render(deps.Stdout, deps.Config.Output, results)
	}}
	cmd.Flags().IntVar(&options.Limit, "limit", 10, "maximum results")
	cmd.Flags().StringVar(&options.Service, "service", "", "restrict to an API service (command, alias, or service id, e.g. bill-analysis, ba, bill_analysis)")
	cmd.Flags().StringVar(&options.Tag, "tag", "", "restrict to a spec tag")
	cmd.Flags().StringVar(&options.Action, "action", "", "restrict to a spec action (list|get|create|update|replace|delete|action)")
	cmd.Flags().BoolVar(&options.ReadOnly, "read-only", false, "only HTTP reads and explicitly classified read-only curated commands")
	return cmd
}

func resolveCommandPath(root *cobra.Command, path []string) (*cobra.Command, error) {
	current := root
	for _, token := range path {
		var match *cobra.Command
		for _, child := range current.Commands() {
			matches := child.Name() == token
			for _, alias := range child.Aliases {
				matches = matches || alias == token
			}
			if matches {
				if match != nil {
					return nil, Exit(2, fmt.Errorf("ambiguous command path at %q", token))
				}
				match = child
			}
		}
		if match == nil {
			return nil, Exit(2, fmt.Errorf("unknown command path %q; inspect --help for available commands", strings.Join(path, " ")))
		}
		current = match
	}
	if current.RunE == nil && current.Run == nil {
		return nil, Exit(2, fmt.Errorf("%s is a command group; specify an operation", current.CommandPath()))
	}
	return current, nil
}

func newSchemaCmd(root *cobra.Command) *cobra.Command {
	var example bool
	var part string
	var depth int
	c := &cobra.Command{Use: "schema <command path...>", Short: "Inspect a command's API schemas, parameters and illustrative example", Long: "Inspect an exact spec-backed operation without authentication or network access. Examples are illustrative, not a promise of server acceptance. --example emits only a schema-validated, non-sensitive request body; unavailable or invalid examples are errors.", Annotations: map[string]string{"flexera.offline": "true"}, Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return Exit(2, fmt.Errorf("provide a command path, for example: cli schema budget create"))
		}
		return nil
	}, RunE: func(cmd *cobra.Command, args []string) error {
		deps := DepsFrom(cmd.Context())
		if depth < 0 {
			return Exit(2, fmt.Errorf("--depth must not be negative"))
		}
		if deps.Config.Output == "table" {
			return Exit(2, fmt.Errorf("cli schema supports JSON output only"))
		}
		if example && (cmd.Flags().Changed("part") || cmd.Root().PersistentFlags().Changed(FlagOutJQ) || cmd.Root().PersistentFlags().Changed(FlagOutFields) || cmd.Root().PersistentFlags().Changed(FlagRawOutput)) {
			return Exit(2, fmt.Errorf("--example cannot be combined with --part, --out-jq, --out-fields or --raw-output"))
		}
		if part != "" && part != "request" && part != "response" && part != "params" {
			return Exit(2, fmt.Errorf("--part must be request, response or params"))
		}
		leaf, err := resolveCommandPath(root, args)
		if err != nil {
			return err
		}
		id := leaf.Annotations["flexera.operationId"]
		if id == "" {
			return Exit(2, fmt.Errorf("%s has no exact API schema; use its --help for inputs", leaf.CommandPath()))
		}
		index, err := catalog.Load()
		if err != nil {
			return Exit(2, err)
		}
		entry, found := index.Lookup(id)
		if !found {
			return Exit(2, fmt.Errorf("catalog has no operation %s", id))
		}
		if example {
			body, err := index.ValidExample(entry)
			if err != nil {
				return Exit(2, err)
			}
			return WriteJSON(deps.Stdout, body, deps.Printer.Style, deps.Printer.IsTTY)
		}
		// Validation must use the complete original schema, not the
		// depth-limited presentation with unresolved deeper references.
		if _, err := index.ValidExample(entry); err != nil {
			entry.RequestExample = nil
		}
		entry.RequestSchema, err = index.Expand(entry.RequestSchema, depth)
		if err != nil {
			return Exit(2, err)
		}
		entry.ResponseSchema, err = index.Expand(entry.ResponseSchema, depth)
		if err != nil {
			return Exit(2, err)
		}
		for i := range entry.Params {
			entry.Params[i].Schema, err = index.Expand(entry.Params[i].Schema, depth)
			if err != nil {
				return Exit(2, err)
			}
		}
		var output any = entry
		switch part {
		case "request":
			if len(entry.RequestSchema) == 0 {
				return Exit(2, fmt.Errorf("%s has no request schema", leaf.CommandPath()))
			}
			output = entry.RequestSchema
		case "response":
			if len(entry.ResponseSchema) == 0 {
				return Exit(2, fmt.Errorf("%s has no response schema", leaf.CommandPath()))
			}
			output = entry.ResponseSchema
		case "params":
			output = entry.Params
		}
		return deps.Printer.Render(deps.Stdout, deps.Config.Output, output)
	}}
	c.Flags().BoolVar(&example, "example", false, "emit only a validated illustrative request body")
	c.Flags().StringVar(&part, "part", "", "schema part (request|response|params)")
	c.Flags().IntVar(&depth, "depth", 3, "reference expansion depth; 0 preserves references")
	return c
}
