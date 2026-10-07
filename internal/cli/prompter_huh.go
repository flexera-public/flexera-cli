package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/charmbracelet/huh"
)

type HuhPrompter struct {
	Input      io.Reader
	Output     io.Writer
	Accessible bool
}

func (p *HuhPrompter) run(ctx context.Context, field huh.Field) error {
	return huh.NewForm(huh.NewGroup(field)).WithInput(p.Input).WithOutput(p.Output).WithAccessible(p.Accessible).RunWithContext(ctx)
}
func (p *HuhPrompter) Ask(ctx context.Context, f PromptField) (any, error) {
	if f.Schema == nil {
		return nil, fmt.Errorf("missing prompt schema")
	}
	if len(f.Schema.Enum) > 0 {
		options := []huh.Option[int]{}
		selected := 0
		for i, item := range f.Schema.Enum {
			raw, _ := json.Marshal(item)
			options = append(options, huh.NewOption(string(raw), i))
			if f.Value != nil {
				prefill, _ := json.Marshal(f.Value)
				if string(prefill) == string(raw) {
					selected = i
				}
			}
		}
		if err := p.run(ctx, huh.NewSelect[int]().Title(f.Title).Options(options...).Value(&selected)); err != nil {
			return nil, err
		}
		return f.Schema.Enum[selected], nil
	}
	if f.Schema.Type != nil && f.Schema.Type.Is("boolean") {
		value, _ := f.Value.(bool)
		if err := p.run(ctx, huh.NewConfirm().Title(f.Title).Value(&value)); err != nil {
			return nil, err
		}
		return value, nil
	}
	text := ""
	if f.Value != nil {
		if value, ok := f.Value.(string); ok {
			text = value
		} else {
			raw, _ := json.Marshal(f.Value)
			text = string(raw)
		}
	}
	parse := func(text string) (any, error) {
		if f.Schema.Type != nil && (f.Schema.Type.Is("number") || f.Schema.Type.Is("integer")) {
			value, err := ParseRequestJSON([]byte(text))
			if err != nil {
				return nil, fmt.Errorf("enter a valid number")
			}
			if _, ok := value.(json.Number); !ok {
				return nil, fmt.Errorf("enter a valid number")
			}
			return value, nil
		}
		return text, nil
	}
	input := huh.NewInput().Title(f.Title).Value(&text).Validate(func(text string) error {
		value, err := parse(text)
		if err != nil {
			return err
		}
		if err := validateSchemaValue(f.Schema, value); err != nil {
			return fmt.Errorf("input does not satisfy this field's schema")
		}
		return nil
	})
	if f.Secret {
		input.EchoMode(huh.EchoModePassword)
	}
	if err := p.run(ctx, input); err != nil {
		return nil, err
	}
	return parse(text)
}
func (p *HuhPrompter) SelectFields(ctx context.Context, title string, options, selected []string) ([]string, error) {
	values := append([]string{}, selected...)
	choices := make([]huh.Option[string], 0, len(options))
	for _, value := range options {
		choices = append(choices, huh.NewOption(value, value))
	}
	if err := p.run(ctx, huh.NewMultiSelect[string]().Title(title).Options(choices...).Value(&values)); err != nil {
		return nil, err
	}
	return values, nil
}
func (p *HuhPrompter) Approve(ctx context.Context) (string, error) {
	value := ""
	if err := p.run(ctx, huh.NewInput().Title("Do you want to perform this action? Only 'yes' will be accepted to approve.").Value(&value)); err != nil {
		return "", err
	}
	return value, nil
}
