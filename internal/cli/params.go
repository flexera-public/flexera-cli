package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ParamInput struct {
	Value   any
	Present bool
}

// ValidateParams distinguishes absence from legitimate zero/false inputs.
// Skipping schema constraints never skips required CLI/config inputs.
func ValidateParams(index *catalog.Catalog, entry catalog.Entry, inputs map[string]ParamInput, skip bool) (map[string]any, error) {
	effective := map[string]any{}
	for _, param := range entry.Params {
		input := inputs[param.Flag]
		if !input.Present {
			if param.Required {
				return nil, Exit(2, fmt.Errorf("--%s is required (or its documented configuration source)", param.Flag))
			}
			continue
		}
		if param.Required && param.Type == "string" {
			if value, ok := input.Value.(string); ok && strings.TrimSpace(value) == "" {
				return nil, Exit(2, fmt.Errorf("--%s must not be empty", param.Flag))
			}
		}
		effective[param.Flag] = input.Value
		if skip || len(param.Schema) == 0 {
			continue
		}
		valueBytes, err := json.Marshal(input.Value)
		if err != nil {
			return nil, Exit(2, fmt.Errorf("--%s cannot be encoded", param.Flag))
		}
		value, err := ParseRequestJSON(valueBytes)
		if err != nil {
			return nil, err
		}
		schema, err := index.RequestSchema(param.Schema)
		if err != nil {
			return nil, Exit(2, err)
		}
		if err := validateSchemaValue(schema, value); err != nil {
			details := validationDetails(err)
			for i := range details {
				details[i].Path = "/params/" + pointerSegment(param.Flag) + details[i].Path
			}
			return nil, Exit(2, &ValidationError{Details: details, Schema: "flexera-cli cli schema " + strings.Join(entry.Command, " ")})
		}
	}
	return effective, nil
}

func parameterFlagValue(flags *pflag.FlagSet, flag *pflag.Flag) (any, error) {
	switch flag.Value.Type() {
	case "string":
		return flags.GetString(flag.Name)
	case "int":
		return flags.GetInt(flag.Name)
	case "int64":
		return flags.GetInt64(flag.Name)
	case "bool":
		return flags.GetBool(flag.Name)
	case "stringSlice":
		return flags.GetStringSlice(flag.Name)
	default:
		return nil, fmt.Errorf("unsupported parameter flag type %s", flag.Value.Type())
	}
}

// ValidateCommandParams is called by generated templates before body reads or
// client creation. Only the generated/config adapters are used, not auth fields.
func ValidateCommandParams(cmd *cobra.Command, operationID string) (map[string]any, error) {
	index, err := catalog.Load()
	if err != nil {
		return nil, Exit(2, err)
	}
	entry, found := index.Lookup(operationID)
	if !found {
		return nil, Exit(2, fmt.Errorf("unknown parameter operation %s", operationID))
	}
	deps := DepsFrom(cmd.Context())
	if deps == nil {
		return nil, fmt.Errorf("parameter validation requires CLI dependencies")
	}
	inputs := map[string]ParamInput{}
	for _, param := range entry.Params {
		if param.Source == "config" {
			if param.Flag != FlagOrgID {
				return nil, Exit(2, fmt.Errorf("unsupported config parameter %s", param.Flag))
			}
			value := any(deps.Config.OrgID)
			if param.Type == "string" {
				value = fmt.Sprint(deps.Config.OrgID)
			}
			inputs[param.Flag] = ParamInput{Value: value, Present: deps.OrgIDPresent || deps.Config.OrgID != 0 || cmd.Root().PersistentFlags().Changed(FlagOrgID)}
			continue
		}
		flag := cmd.Flags().Lookup(param.Flag)
		if flag == nil {
			return nil, Exit(2, fmt.Errorf("catalog parameter --%s has no registered flag", param.Flag))
		}
		value, err := parameterFlagValue(cmd.Flags(), flag)
		if err != nil {
			return nil, Exit(2, err)
		}
		inputs[param.Flag] = ParamInput{Value: value, Present: flag.Changed}
	}
	skip, _ := cmd.Root().PersistentFlags().GetBool(FlagNoValidate)
	return ValidateParams(index, entry, inputs, skip)
}
