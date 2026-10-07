package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// RenderHuman shows effective inputs, not a server-side diff. It never writes a
// body file, prints secrets or invents a replayable command pointing to a file.
func (p Plan) RenderHuman(w io.Writer) error {
	var encoded bytes.Buffer
	if err := p.RenderJSON(&encoded, Printer{Style: JSONStyleCompact}); err != nil {
		return err
	}
	value, err := ParseRequestJSON(encoded.Bytes())
	if err != nil {
		return err
	}
	output := value.(map[string]any)
	plan := output["plan"].(map[string]any)
	var text bytes.Buffer
	fmt.Fprintln(&text, "flexera-cli will perform the following action:")
	fmt.Fprintf(&text, "\n  # %s (%s %s)\n", strconv.Quote(p.Command), p.Method, p.Path)
	if p.OrgID != 0 {
		fmt.Fprintf(&text, "  org: %d\n", p.OrgID)
	}
	if params, ok := plan["params"].(map[string]any); ok {
		keys := sortedPlanKeys(params)
		for _, key := range keys {
			raw, err := json.Marshal(params[key])
			if err != nil {
				return err
			}
			fmt.Fprintf(&text, "  param %s = %s\n", strconv.Quote(key), raw)
		}
	}
	if body, exists := plan["body"]; exists {
		if fields, ok := body.(map[string]any); ok && !p.RawUpload {
			for _, key := range sortedPlanKeys(fields) {
				raw, err := json.Marshal(fields[key])
				if err != nil {
					return err
				}
				fmt.Fprintf(&text, "  + %s = %s\n", strconv.Quote(key), raw)
			}
		} else {
			raw, err := json.Marshal(body)
			if err != nil {
				return err
			}
			fmt.Fprintf(&text, "  body = %s\n", raw)
		}
	}
	if p.Validation != nil {
		fmt.Fprintf(&text, "\n  Validation: %s\n", p.Validation.Status)
	}
	fmt.Fprintln(&text, "\n  '+' marks supplied request fields, not a server-side diff.")
	if output["redacted"] == true {
		fmt.Fprintln(&text, "  Sensitive values are redacted; this preview is not replayable.")
	}
	if output["redacted"] != true && !p.RawUpload && len(p.Body) > 0 && len(p.Body) <= 1024 {
		command := p.Command
		if params, ok := plan["params"].(map[string]any); ok {
			for _, key := range sortedPlanKeys(params) {
				raw, err := json.Marshal(params[key])
				if err != nil {
					return err
				}
				value := string(raw)
				if text, ok := params[key].(string); ok {
					value = text
				}
				command += " --" + key + " " + shellQuote(value)
			}
		}
		fmt.Fprintln(&text, "\n  Reviewed input invocation (authentication/config may still be required):")
		fmt.Fprintln(&text, "  "+command+" --body "+shellQuote(string(p.Body)))
	} else {
		fmt.Fprintln(&text, "  No reusable command or body file has been exported. Keep your original request in a private file.")
	}
	_, err = w.Write(text.Bytes())
	return err
}

func shellQuote(value string) string { return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'" }

func sortedPlanKeys(values map[string]any) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
