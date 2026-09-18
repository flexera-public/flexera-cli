// Package rulebaseddimension provides hand-written workflows built on top of
// the generated rule-based-dimension commands.
package rulebaseddimension

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

type bulkInput struct {
	OrgID           int                 `json:"org_id"`
	Dimensions      []bulkDimensionSpec `json:"dimensions"`
	DryRun          bool                `json:"dry_run"`
	ContinueOnError bool                `json:"continue_on_error"`
}

type bulkDimensionSpec struct {
	ID          string                                                      `json:"id"`
	Name        string                                                      `json:"name"`
	EffectiveAt string                                                      `json:"effective_at"`
	Rules       []flexera.FinopsCustomizationsRuleBasedDimensionRulePayload `json:"rules"`
}

// NewBulkCmd builds the rule-based-dimension bulk workflow command.
func NewBulkCmd() *cobra.Command {
	var input, file string
	var dryRun, continueOnError bool

	c := &cobra.Command{
		Use:   "bulk",
		Short: "Create or update rule-based dimensions from JSON input",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			raw, err := resolveInput(input, file, cmd.InOrStdin())
			if err != nil {
				return err
			}
			var decoded bulkInput
			if err := json.Unmarshal(raw, &decoded); err != nil {
				return fmt.Errorf("decoding bulk input: %w", err)
			}

			orgID := decoded.OrgID
			if orgID == 0 {
				if err := deps.Config.RequireOrgID(); err != nil {
					return err
				}
				orgID = deps.Config.OrgID
			}
			dimensions := make([]flexera.RuleBasedDimensionSpec, 0, len(decoded.Dimensions))
			for _, dim := range decoded.Dimensions {
				dimensions = append(dimensions, flexera.RuleBasedDimensionSpec{
					ID:          dim.ID,
					Name:        dim.Name,
					EffectiveAt: dim.EffectiveAt,
					Rules:       dim.Rules,
				})
			}

			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			tool := flexera.NewRuleBasedDimensionBulkTool(client)
			output, err := tool.Invoke(cmd.Context(), flexera.RuleBasedDimensionBulkInput{
				OrgID:           orgID,
				Dimensions:      dimensions,
				DryRun:          decoded.DryRun || dryRun,
				ContinueOnError: decoded.ContinueOnError || continueOnError,
			})
			if err != nil {
				return fmt.Errorf("%s: %w", tool.Name(), err)
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, output)
		},
	}
	c.Flags().StringVar(&input, "input", "", "inline JSON input (mutually exclusive with --file)")
	c.Flags().StringVar(&file, "file", "", "path to JSON input, or - to read JSON from stdin")
	c.Flags().BoolVar(&dryRun, "dry-run", false, "validate and report dimensions without making API calls")
	c.Flags().BoolVar(&continueOnError, "continue-on-error", false, "continue processing after a dimension fails")
	return c
}

func resolveInput(inline, file string, stdin interface{ Read([]byte) (int, error) }) ([]byte, error) {
	inline = strings.TrimSpace(inline)
	file = strings.TrimSpace(file)
	if inline == "" && file == "" {
		return nil, errors.New("bulk input is required: pass --input or --file (use --file - for stdin)")
	}
	if inline != "" && file != "" {
		return nil, errors.New("--input and --file are mutually exclusive")
	}
	if inline != "" {
		if !json.Valid([]byte(inline)) {
			return nil, errors.New("--input is not valid JSON")
		}
		return []byte(inline), nil
	}
	arg := file
	if arg != "-" {
		arg = "@" + arg
	} else {
		arg = "@-"
	}
	return clipkg.ResolveBody(arg, nil, stdin)
}

// Attach adds the curated bulk workflow below the generated
// rule-based-dimension command.
func Attach(root *cobra.Command) error {
	for _, command := range root.Commands() {
		if command.Name() == "rule-based-dimension" {
			command.Aliases = append(command.Aliases, "rule-base-dimension")
			command.AddCommand(NewBulkCmd())
			command.AddCommand(NewFromCSVCommand())
			return nil
		}
	}
	return errors.New("rule-based-dimension generated command is not registered")
}
