package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	flexera "github.com/flexera-public/unified-go-client"
	"github.com/itchyny/gojq"
	"golang.org/x/term"
)

// JSONStyle controls whitespace independently of the output format.
type JSONStyle string

const (
	JSONStyleAuto    JSONStyle = "auto"
	JSONStylePretty  JSONStyle = "pretty"
	JSONStyleCompact JSONStyle = "compact"
)

// ParseJSONStyle validates a style; the empty value means auto.
func ParseJSONStyle(value string) (JSONStyle, error) {
	if value == "" {
		return JSONStyleAuto, nil
	}
	switch JSONStyle(value) {
	case JSONStyleAuto, JSONStylePretty, JSONStyleCompact:
		return JSONStyle(value), nil
	default:
		return "", fmt.Errorf("unsupported JSON style %q (expected auto|pretty|compact)", value)
	}
}

// Printer owns JSON response policy while retaining SDK table rendering.
// Its zero value uses auto style and detects terminals on *os.File writers.
type Printer struct {
	Context   context.Context
	Style     JSONStyle
	IsTTY     func(io.Writer) bool
	JQ        *gojq.Code
	Fields    []FieldPath
	RawOutput bool
	Envelope  string
}

// WriteJSON is the shared formatter for responses, errors and safety plans.
// It never applies response shaping. IsTTY may be nil for production detection.
func WriteJSON(w io.Writer, value any, style JSONStyle, isTTY func(io.Writer) bool) error {
	resolved, err := ParseJSONStyle(string(style))
	if err != nil {
		return err
	}
	pretty := resolved == JSONStylePretty
	if resolved == JSONStyleAuto {
		if isTTY == nil {
			isTTY = writerIsTTY
		}
		pretty = isTTY(w)
	}
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	if pretty {
		encoder.SetIndent("", "  ")
	}
	return encoder.Encode(value)
}

func writerIsTTY(w io.Writer) bool {
	file, ok := w.(*os.File)
	return ok && term.IsTerminal(int(file.Fd()))
}

// Render writes value to w in the given format ("json" or "table").
func (p Printer) Render(w io.Writer, format string, value any) error {
	if p.RawOutput && p.JQ == nil {
		return Exit(2, fmt.Errorf("--raw-output requires --out-jq"))
	}
	if p.JQ != nil || len(p.Fields) != 0 {
		return p.renderShaped(w, format, value)
	}
	switch strings.ToLower(strings.TrimSpace(format)) {
	case "", "json":
		return WriteJSON(w, value, p.Style, p.IsTTY)
	case "table":
		if table, ok := value.(interface{ TableRows() ([]string, [][]string) }); ok {
			headers, rows := table.TableRows()
			tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
			if _, err := fmt.Fprintln(tw, strings.Join(headers, "\t")); err != nil {
				return err
			}
			for _, row := range rows {
				if _, err := fmt.Fprintln(tw, strings.Join(row, "\t")); err != nil {
					return err
				}
			}
			return tw.Flush()
		}
		return flexera.WriteTable(w, value)
	default:
		return fmt.Errorf("unsupported output format %q", format)
	}
}

// RenderError writes err to w as a JSON error object, matching the legacy
// CLI's stderr error contract.
func (p Printer) RenderError(w io.Writer, err error) error {
	return WriteJSON(w, errorProperties(err), p.Style, p.IsTTY)
}
