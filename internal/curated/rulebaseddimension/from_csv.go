// Package rulebaseddimension provides CSV-backed rule-based-dimension
// workflows built on the unified client local transformation API.
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

// NewFromCSVCommand builds the CSV workflow subcommand for the generated
// rule-based-dimension command.
func NewFromCSVCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "from_csv",
		Short: "Generate and apply rule-based dimensions from CSV",
	}
	c.AddCommand(newGenerateCmd(), newApplyCmd("create"), newApplyCmd("update"))
	return c
}

type csvFlags struct {
	file            string
	separatorHeader string
	idTemplate      string
	nameTemplate    string
	effectiveAt     string
	columnConfig    string
	caseInsensitive bool
}

func (f *csvFlags) addTo(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.file, "file", "f", "", "CSV file path, or - for stdin (required)")
	cmd.Flags().StringVar(&f.file, "csv", "", "alias for --file")
	cmd.Flags().StringVar(&f.separatorHeader, "separator-header", "", "CSV header separating condition columns from output columns")
	cmd.Flags().StringVar(&f.idTemplate, "id-template", "", "Go template for generated IDs")
	cmd.Flags().StringVar(&f.nameTemplate, "name-template", "", "Go template for generated names")
	cmd.Flags().StringVar(&f.effectiveAt, "effective-at", "", "effective month for generated rules (for example 2024-01)")
	cmd.Flags().StringVar(&f.columnConfig, "column-config", "", "JSON map of output-column overrides ({header:{id,name,skip}})")
	cmd.Flags().BoolVar(&f.caseInsensitive, "case-insensitive", true, "make generated conditions case-insensitive")
}

func (f csvFlags) options() (flexera.RuleBasedDimensionCSVOptions, error) {
	opts := flexera.DefaultRuleBasedDimensionCSVOptions()
	if strings.TrimSpace(f.separatorHeader) != "" {
		opts.SeparatorHeader = f.separatorHeader
	}
	if strings.TrimSpace(f.idTemplate) != "" {
		opts.IDTemplate = f.idTemplate
	}
	if strings.TrimSpace(f.nameTemplate) != "" {
		opts.NameTemplate = f.nameTemplate
	}
	if strings.TrimSpace(f.effectiveAt) != "" {
		opts.EffectiveAt = f.effectiveAt
	}
	opts.CaseInsensitive = f.caseInsensitive
	if strings.TrimSpace(f.columnConfig) != "" {
		if err := json.Unmarshal([]byte(f.columnConfig), &opts.ColumnConfig); err != nil {
			return flexera.RuleBasedDimensionCSVOptions{}, fmt.Errorf("decoding --column-config: %w", err)
		}
	}
	return opts, nil
}

func (f csvFlags) generate() ([]flexera.RuleBasedDimensionSpec, error) {
	if strings.TrimSpace(f.file) == "" {
		return nil, errors.New("--file is required")
	}
	opts, err := f.options()
	if err != nil {
		return nil, err
	}
	return flexera.GenerateRuleBasedDimensionsFromCSVFile(f.file, opts)
}

func newGenerateCmd() *cobra.Command {
	var flags csvFlags
	c := &cobra.Command{
		Use:   "generate",
		Short: "Generate rule-based-dimension specs from CSV",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			specs, err := generateFromFlags(cmd, flags)
			if err != nil {
				return err
			}
			deps := clipkg.DepsFrom(cmd.Context())
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, specs)
		},
	}
	flags.addTo(c)
	return c
}

func newApplyCmd(name string) *cobra.Command {
	var flags csvFlags
	var dryRun, continueOnError bool
	c := &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Create or update rule-based dimensions from CSV (%s)", name),
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			specs, err := generateFromFlags(cmd, flags)
			if err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			output, err := flexera.NewRuleBasedDimensionBulkTool(client).Invoke(
				cmd.Context(),
				flexera.RuleBasedDimensionBulkInput{
					OrgID:           deps.Config.OrgID,
					Dimensions:      specs,
					DryRun:          dryRun,
					ContinueOnError: continueOnError,
				},
			)
			if err != nil {
				return fmt.Errorf("rule-based-dimensions %s: %w", name, err)
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, output)
		},
	}
	flags.addTo(c)
	c.Flags().BoolVar(&dryRun, "dry-run", false, "validate and report dimensions without making API calls")
	c.Flags().BoolVar(&continueOnError, "continue-on-error", false, "continue processing after a dimension fails")
	return c
}

func generateFromFlags(cmd *cobra.Command, flags csvFlags) ([]flexera.RuleBasedDimensionSpec, error) {
	if flags.file == "-" {
		opts, err := flags.options()
		if err != nil {
			return nil, err
		}
		return flexera.GenerateRuleBasedDimensionsFromCSV(cmd.InOrStdin(), opts)
	}
	return flags.generate()
}
