package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/spf13/cobra"
)

// EnrichMetadataHelp uses the published, sample-sanitized catalog, never the
// raw upstream examples. It does not change defaults or input requirements.
func EnrichMetadataHelp(root *cobra.Command) error {
	index, err := catalog.Load()
	if err != nil {
		return err
	}
	var walk func(*cobra.Command) error
	walk = func(cmd *cobra.Command) error {
		if id := cmd.Annotations["flexera.operationId"]; id != "" {
			entry, found := index.Lookup(id)
			if !found {
				return fmt.Errorf("help: command %s has no catalog operation %s", cmd.CommandPath(), id)
			}
			if cmd.Long == "" {
				cmd.Long = cmd.Short
				if description := operationHelpText(entry.Description); description != "" && description != cmd.Short {
					cmd.Long += "\n\n" + description
				}
				if entry.RequestDescription != "" {
					cmd.Long += "\n\nRequest body: " + operationHelpText(entry.RequestDescription)
				}
			}
			if entry.Deprecated {
				cmd.Long += "\n\nDeprecated in the upstream API."
			}
			enrich := func(flag, description string, raw, example json.RawMessage, required bool) error {
				f := cmd.LocalNonPersistentFlags().Lookup(flag)
				if f == nil {
					return nil // Config/pagination inputs and curated body adapters have their own help.
				}
				schema, err := index.Expand(raw, 1)
				if err != nil {
					return err
				}
				guidance, err := schemaHelp(schema, example, required, flag)
				if err != nil {
					return err
				}
				parts := []string{f.Usage}
				if description != "" && !strings.Contains(f.Usage, description) {
					parts = append(parts, helpDescription(description))
				}
				parts = append(parts, guidance...)
				unique := []string{}
				seen := map[string]bool{}
				for _, part := range parts {
					if !seen[part] {
						unique = append(unique, part)
						seen[part] = true
					}
				}
				f.Usage = strings.Join(unique, "; ")
				return nil
			}
			for _, param := range entry.Params {
				example, err := parameterHelpExample(param.Example, param.Examples)
				if err != nil {
					return fmt.Errorf("help %s --%s example: %w", id, param.Flag, err)
				}
				if err := enrich(param.Flag, param.Description, param.Schema, example, param.Required); err != nil {
					return fmt.Errorf("help %s --%s: %w", id, param.Flag, err)
				}
			}
			for _, field := range entry.BodyFields {
				if err := enrich(field.Flag, "", field.Schema, nil, field.Required); err != nil {
					return fmt.Errorf("help %s --%s: %w", id, field.Flag, err)
				}
			}
		}
		for _, child := range cmd.Commands() {
			if err := walk(child); err != nil {
				return err
			}
		}
		return nil
	}
	return walk(root)
}

func schemaHelp(raw, example json.RawMessage, required bool, name string) ([]string, error) {
	parts := []string{}
	if required {
		parts = append(parts, "required by API")
	}
	if len(raw) == 0 {
		return parts, nil
	}
	var schema map[string]any
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	if err := decoder.Decode(&schema); err != nil {
		return nil, err
	}
	if description, ok := schema["description"].(string); ok && description != "" {
		parts = append(parts, helpDescription(description))
	}
	if format, ok := schema["format"].(string); ok && format != "" {
		parts = append(parts, "format: "+format)
	}
	secret := schema["writeOnly"] == true || schema["format"] == "password" || sensitiveName(name)
	for _, key := range []string{"enum", "minimum", "maximum", "exclusiveMinimum", "exclusiveMaximum", "multipleOf", "minLength", "maxLength", "pattern", "minItems", "maxItems", "uniqueItems", "default"} {
		if value, found := schema[key]; found && !secret {
			encoded, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			label := key
			if key == "default" {
				label = "API default"
			}
			parts = append(parts, label+": "+shortHelpValue(encoded))
		}
	}
	if schema["type"] == "array" {
		parts = append(parts, "CLI: comma-separated values or repeated flag")
		if items, ok := schema["items"].(map[string]any); ok && !secret && items["writeOnly"] != true && items["format"] != "password" {
			for _, key := range []string{"enum", "format"} {
				if value, found := items[key]; found {
					encoded, err := json.Marshal(value)
					if err != nil {
						return nil, err
					}
					parts = append(parts, "items."+key+": "+shortHelpValue(encoded))
				}
			}
		}
	}
	if !secret {
		if len(example) == 0 {
			if value, found := schema["example"]; found {
				var err error
				example, err = json.Marshal(value)
				if err != nil {
					return nil, err
				}
			}
		}
		if len(example) != 0 {
			parts = append(parts, "illustrative example: "+shortHelpValue(example))
		}
	}
	return parts, nil
}

func parameterHelpExample(example, examples json.RawMessage) (json.RawMessage, error) {
	if len(example) != 0 || len(examples) == 0 {
		return example, nil
	}
	var named map[string]struct {
		Value json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(examples, &named); err != nil {
		return nil, err
	}
	names := make([]string, 0, len(named))
	for name := range named {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		if len(named[name].Value) != 0 {
			return named[name].Value, nil
		}
	}
	return nil, nil
}

func shortHelpValue(raw []byte) string {
	text := strings.Join(strings.Fields(string(raw)), " ")
	runes := []rune(text)
	if len(runes) > 180 {
		return string(runes[:180]) + "... (see cli schema)"
	}
	return text
}

func helpDescription(description string) string {
	text := strings.Join(strings.Fields(description), " ")
	runes := []rune(text)
	if len(runes) > 360 {
		return string(runes[:360]) + "... (see cli schema)"
	}
	return text
}

func operationHelpText(text string) string {
	lines := strings.Split(strings.TrimSpace(text), "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " \t\r")
	}
	return strings.Join(lines, "\n")
}
