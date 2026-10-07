package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/itchyny/gojq"
)

// CompileJQ deliberately supplies no environment or input extensions.
func CompileJQ(expression string) (*gojq.Code, error) {
	query, err := gojq.Parse(expression)
	if err != nil {
		return nil, Exit(2, err)
	}
	code, err := gojq.Compile(query)
	if err != nil {
		return nil, Exit(2, err)
	}
	return code, nil
}

// Results are spooled to a private file so evaluation errors cannot leak partial
// stdout and result streams do not accumulate in memory. No output cap is imposed.
func (p Printer) renderShaped(w io.Writer, format string, value any) error {
	ctx := p.Context
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	format = strings.ToLower(strings.TrimSpace(format))
	if format == "" {
		format = "json"
	}
	if format != "json" && format != "table" {
		return Exit(2, fmt.Errorf("unsupported output format %q", format))
	}
	if format == "table" && p.RawOutput {
		return Exit(2, fmt.Errorf("table output does not support --raw-output"))
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	var input any
	if err := decoder.Decode(&input); err != nil {
		return err
	}
	if len(p.Fields) != 0 {
		code, err := CompileFields(p.Fields, p.Envelope == "values")
		if err != nil {
			return err
		}
		result, ok := code.RunWithContext(ctx, input).Next()
		if !ok {
			return fmt.Errorf("field projection produced no result")
		}
		if err, ok := result.(error); ok {
			return err
		}
		input = result
	}
	if p.JQ == nil {
		if err := ctx.Err(); err != nil {
			return err
		}
		p.Fields = nil
		return p.Render(shapingContextWriter{ctx, w}, format, input)
	}
	spool, err := os.CreateTemp("", "flexera-cli-output-*")
	if err != nil {
		return fmt.Errorf("create output spool: %w", err)
	}
	defer os.Remove(spool.Name())
	defer spool.Close()
	iter := p.JQ.RunWithContext(ctx, input)
	var first any
	count := 0
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		item, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := item.(error); ok {
			return fmt.Errorf("jq evaluation failed: %w", err)
		}
		// Only distinguish zero, one and multiple; never let an unbounded
		// result stream overflow a counter.
		if count < 2 {
			count++
		}
		if count == 1 {
			first = item
		}
		if count > 1 && (format == "table" || p.Style == JSONStylePretty) {
			return Exit(2, fmt.Errorf("multiple jq results require compact JSON output, not pretty JSON or table"))
		}
		if text, ok := item.(string); ok && p.RawOutput {
			_, err = fmt.Fprintln(spool, text)
		} else {
			err = WriteJSON(spool, item, JSONStyleCompact, nil)
		}
		if err != nil {
			return fmt.Errorf("write output spool: %w", err)
		}
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if count == 0 {
		return nil
	}
	if count == 1 && !p.RawOutput {
		p.JQ = nil
		p.Fields = nil
		return p.Render(shapingContextWriter{ctx, w}, format, first)
	}
	if _, err := spool.Seek(0, io.SeekStart); err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	// The wrapper also prevents io.Copy from bypassing checks via File.WriteTo.
	_, err = io.Copy(shapingContextWriter{ctx, w}, shapingContextReader{ctx, spool})
	if err == nil {
		err = ctx.Err()
	}
	return err
}

type shapingContextReader struct {
	ctx context.Context
	r   io.Reader
}

func (r shapingContextReader) Read(b []byte) (int, error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	return r.r.Read(b)
}

type shapingContextWriter struct {
	ctx context.Context
	w   io.Writer
}

func (w shapingContextWriter) Write(b []byte) (int, error) {
	if err := w.ctx.Err(); err != nil {
		return 0, err
	}
	n, err := w.w.Write(b)
	if err == nil {
		err = w.ctx.Err()
	}
	return n, err
}
