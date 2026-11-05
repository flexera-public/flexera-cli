package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// ResolveBody produces the JSON request body for a write operation.
//
// If raw is non-empty it takes precedence (the --body escape hatch): it may
// be inline JSON, "@<path>" to read a file, or "@-" to read stdin. Otherwise
// the typed value (assembled from per-field flags) is marshaled to JSON.
//
// stdin is injected so tests can supply a reader; pass os.Stdin in production.
func ResolveBody(raw string, typed any, stdin io.Reader) (json.RawMessage, error) {
	if strings.TrimSpace(raw) != "" {
		b, err := readBodyArg(raw, stdin)
		if err != nil {
			return nil, err
		}
		if !json.Valid(b) {
			return nil, Exit(2, fmt.Errorf("--body is not valid JSON"))
		}
		return json.RawMessage(b), nil
	}
	if typed == nil {
		return nil, nil
	}
	b, err := json.Marshal(typed)
	if err != nil {
		return nil, fmt.Errorf("encoding request body: %w", err)
	}
	return json.RawMessage(b), nil
}

// ResolveRawBody reads an arbitrary request body from --body. Unlike
// ResolveBody, it does not require JSON; this is used for binary uploads.
func ResolveRawBody(raw string, stdin io.Reader) ([]byte, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, nil
	}
	return readBodyArg(raw, stdin)
}

func BuildMultipartBody(fileField, filePath string, fields map[string]string) ([]byte, string, error) {
	if strings.TrimSpace(fileField) == "" || strings.TrimSpace(filePath) == "" {
		return nil, "", Exit(2, fmt.Errorf("a multipart upload file is required"))
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("opening multipart upload file %q: %w", filePath, err)
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	keys := make([]string, 0, len(fields))
	for key := range fields {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	for _, key := range keys {
		if err := writer.WriteField(key, fields[key]); err != nil {
			return nil, "", fmt.Errorf("writing multipart field %q: %w", key, err)
		}
	}
	part, err := writer.CreateFormFile(fileField, filepath.Base(filePath))
	if err != nil {
		return nil, "", fmt.Errorf("creating multipart file field %q: %w", fileField, err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, "", fmt.Errorf("reading multipart upload file %q: %w", filePath, err)
	}
	if err := file.Close(); err != nil {
		return nil, "", fmt.Errorf("closing multipart upload file %q: %w", filePath, err)
	}
	if err := writer.Close(); err != nil {
		return nil, "", fmt.Errorf("finishing multipart request body: %w", err)
	}
	return body.Bytes(), writer.FormDataContentType(), nil
}

// ConfirmWrite enforces the shared write-op confirmation contract used by
// generated mutating commands. On --dry-run it prints a plan summary to w and
// returns done=true (the caller should stop without calling the API).
// Destructive ops require --yes. done=true means the dry-run early-exit fired.
func ConfirmWrite(dryRun, yes, destructive bool, w io.Writer, plan map[string]any, printers ...Printer) (bool, error) {
	if dryRun {
		printer := Printer{}
		if len(printers) != 0 {
			printer = printers[0]
		}
		return true, WriteJSON(w, map[string]any{"dryRun": true, "destructive": destructive, "plan": plan}, printer.Style, printer.IsTTY)
	}
	if destructive && !yes {
		return false, fmt.Errorf("destructive operation requires --yes (or use --dry-run to preview)")
	}
	return false, nil
}

func readBodyArg(arg string, stdin io.Reader) ([]byte, error) {
	switch {
	case arg == "@-":
		if stdin == nil {
			stdin = os.Stdin
		}
		b, err := io.ReadAll(stdin)
		if err != nil {
			return nil, fmt.Errorf("reading body from stdin: %w", err)
		}
		return b, nil
	case strings.HasPrefix(arg, "@"):
		path := arg[1:]
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("reading body file %q: %w", path, err)
		}
		return b, nil
	default:
		return []byte(arg), nil
	}
}
