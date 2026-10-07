package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func shapingTempDir(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("TMPDIR", dir)
	t.Cleanup(func() {
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Error(err)
		} else if len(entries) != 0 {
			t.Errorf("output spool files were not cleaned up: %v", entries)
		}
	})
}

func shapingPrinter(t *testing.T, expression string) Printer {
	t.Helper()
	shapingTempDir(t)
	code, err := CompileJQ(expression)
	if err != nil {
		t.Fatal(err)
	}
	return Printer{JQ: code, Style: JSONStyleCompact}
}

func TestShapingCancellation(t *testing.T) {
	p := shapingPrinter(t, "repeat(.)")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	p.Context = ctx
	var out bytes.Buffer
	if err := p.Render(&out, "json", 1); err == nil || out.Len() != 0 {
		t.Fatalf("cancelled evaluation must fail atomically: %v, %q", err, out.String())
	}
}

func TestShapingResults(t *testing.T) {
	for _, tc := range []struct {
		name, query, want string
		raw               bool
	}{
		{"zero", "empty", "", false},
		{"one", `{id: .id}`, "{\"id\":7}\n", false},
		{"multiple", `.id, "<>&", null`, "7\n\"<>&\"\nnull\n", false},
		{"raw-zero", "empty", "", true},
		{"raw-one-string", `"hello\nworld"`, "hello\nworld\n", true},
		{"raw-one-number", ".id", "7\n", true},
		{"raw-multiple", `"hello", .id, {ok: true}, null`, "hello\n7\n{\"ok\":true}\nnull\n", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := shapingPrinter(t, tc.query)
			p.RawOutput = tc.raw
			var out bytes.Buffer
			if err := p.Render(&out, "json", map[string]any{"id": 7}); err != nil {
				t.Fatal(err)
			}
			if out.String() != tc.want {
				t.Fatalf("got %q, want %q", out.String(), tc.want)
			}
		})
	}
}

func TestShapingIncompatibilities(t *testing.T) {
	for _, tc := range []struct {
		name, query, format, message string
		style                        JSONStyle
		raw                          bool
	}{
		{"pretty-multiple", "1, 2", "json", "multiple jq results", JSONStylePretty, false},
		{"table-multiple", "1, 2", "table", "multiple jq results", JSONStyleCompact, false},
		{"raw-pretty-multiple", `"a", "b"`, "json", "multiple jq results", JSONStylePretty, true},
		{"raw-table", `"a"`, "table", "--raw-output", JSONStyleCompact, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := shapingPrinter(t, tc.query)
			p.Style, p.RawOutput = tc.style, tc.raw
			var out bytes.Buffer
			err := p.Render(&out, tc.format, nil)
			var exit *ExitError
			if !errors.As(err, &exit) || exit.Code != 2 || !strings.Contains(err.Error(), tc.message) {
				t.Fatalf("got %v, want exit code 2 containing %q", err, tc.message)
			}
			if out.Len() != 0 {
				t.Fatalf("failure leaked stdout: %q", out.String())
			}
		})
	}
}

func TestShapingSinglePretty(t *testing.T) {
	p := shapingPrinter(t, ".")
	p.Style = JSONStylePretty
	var out bytes.Buffer
	if err := p.Render(&out, "json", map[string]any{"id": 7}); err != nil {
		t.Fatal(err)
	}
	if want := "{\n  \"id\": 7\n}\n"; out.String() != want {
		t.Fatalf("got %q, want %q", out.String(), want)
	}
}

func TestShapingRuntimeErrorIsAtomic(t *testing.T) {
	for _, raw := range []bool{false, true} {
		t.Run(map[bool]string{false: "json", true: "raw"}[raw], func(t *testing.T) {
			p := shapingPrinter(t, `"first", error("boom")`)
			p.RawOutput = raw
			var out bytes.Buffer
			err := p.Render(&out, "json", nil)
			if err == nil || !strings.Contains(err.Error(), "jq evaluation failed") || !strings.Contains(err.Error(), "boom") {
				t.Fatalf("expected runtime error, got %v", err)
			}
			if out.Len() != 0 {
				t.Fatalf("runtime error leaked stdout: %q", out.String())
			}
		})
	}
}

func TestShapingFieldsBeforeJQRetainsEnvelope(t *testing.T) {
	p := shapingPrinter(t, `. + {projected: (.values[0] | keys == ["id", "owner"])}`)
	fields, err := ParseFields("id,owner.email")
	if err != nil {
		t.Fatal(err)
	}
	p.Fields, p.Envelope = fields, "values"
	input := map[string]any{
		"values":   []any{map[string]any{"id": 7, "owner": map[string]any{"email": "a", "secret": true}, "extra": 9}},
		"nextPage": "next", "total": 4,
	}
	var out bytes.Buffer
	if err := p.Render(&out, "json", input); err != nil {
		t.Fatal(err)
	}
	want := "{\"nextPage\":\"next\",\"projected\":true,\"total\":4,\"values\":[{\"id\":7,\"owner\":{\"email\":\"a\"}}]}\n"
	if out.String() != want {
		t.Fatalf("got %q, want %q", out.String(), want)
	}
}

func TestShapingLargeInteger(t *testing.T) {
	for _, query := range []string{".id", ".id + 1", ".id, (.id + 1)"} {
		t.Run(query, func(t *testing.T) {
			p := shapingPrinter(t, query)
			p.Fields = []FieldPath{{"id"}}
			var out bytes.Buffer
			if err := p.Render(&out, "json", map[string]any{"id": json.Number("9007199254740993")}); err != nil {
				t.Fatal(err)
			}
			want := map[string]string{
				".id": "9007199254740993\n", ".id + 1": "9007199254740994\n",
				".id, (.id + 1)": "9007199254740993\n9007199254740994\n",
			}[query]
			if out.String() != want {
				t.Fatalf("got %q, want %q", out.String(), want)
			}
		})
	}
}

func TestShapingWriterFailureCleansSpool(t *testing.T) {
	p := shapingPrinter(t, "1, 2")
	sentinel := errors.New("destination failed")
	if err := p.Render(failingWriter{sentinel}, "json", nil); !errors.Is(err, sentinel) {
		t.Fatalf("got %v, want destination error", err)
	}
}

func TestShapingRestrictedJQ(t *testing.T) {
	t.Setenv("FLEXERA_SHAPING_SECRET", "environment-secret")
	for _, expression := range []string{"env", "$ENV"} {
		p := shapingPrinter(t, expression)
		var out bytes.Buffer
		if err := p.Render(&out, "json", nil); err != nil || out.String() != "{}\n" {
			t.Fatalf("%s exposed environment: %q, %v", expression, out.String(), err)
		}
	}
	for _, expression := range []string{"input", "inputs", "$__loc__", `import "secret" as secret; secret::value`} {
		_, err := CompileJQ(expression)
		var exit *ExitError
		if !errors.As(err, &exit) || exit.Code != 2 {
			t.Fatalf("%s must reject external input/location/modules, got %v", expression, err)
		}
	}
}

func TestShapingFieldProjectionCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	p := Printer{Context: ctx, Fields: []FieldPath{{"id"}}, Style: JSONStyleCompact}
	var out bytes.Buffer
	// Cancel after initial context checks, while preparing projection input.
	err := p.Render(&out, "json", shapingCancelMarshal{cancel})
	if !errors.Is(err, context.Canceled) || out.Len() != 0 {
		t.Fatalf("projection ignored cancellation: %v, %q", err, out.String())
	}
}

type shapingCancelMarshal struct{ cancel context.CancelFunc }

func (v shapingCancelMarshal) MarshalJSON() ([]byte, error) {
	v.cancel()
	return []byte(`{"id":7}`), nil
}

// Cancellation is triggered synchronously only after a result has reached the
// spool, avoiding timing assumptions about goroutine scheduling or jq speed.
type shapingSpoolCancelContext struct {
	context.Context
	cancel context.CancelFunc
	t      *testing.T
	seen   bool
}

func (c *shapingSpoolCancelContext) Err() error {
	entries, err := os.ReadDir(os.TempDir())
	if err != nil {
		c.t.Fatal(err)
	}
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			c.t.Fatal(err)
		}
		if info.Size() > 0 {
			if info.Mode().Perm()&0077 != 0 {
				c.t.Fatalf("spool permissions expose data: %v", info.Mode())
			}
			c.seen = true
			c.cancel()
		}
	}
	return c.Context.Err()
}

func TestShapingCancellationAfterEvaluationBegins(t *testing.T) {
	p := shapingPrinter(t, "repeat(.)")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	observed := &shapingSpoolCancelContext{Context: ctx, cancel: cancel, t: t}
	p.Context = observed
	var out bytes.Buffer
	if err := p.Render(&out, "json", 7); !errors.Is(err, context.Canceled) || !observed.seen || out.Len() != 0 {
		t.Fatalf("stream cancellation must discard spool: %v, started=%v, output=%q", err, observed.seen, out.String())
	}
}

func TestShapingCancellationDuringSpoolCopy(t *testing.T) {
	p := shapingPrinter(t, "range(0; 20000)")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	p.Context = ctx
	w := &shapingCancelWriter{t: t, cancel: cancel}
	err := p.Render(w, "json", nil)
	if !errors.Is(err, context.Canceled) || w.writes != 1 || w.Len() == 0 || int64(w.Len()) >= w.spoolSize {
		t.Fatalf("copy did not stop after cancellation: %v, writes=%d, bytes=%d, spool=%d", err, w.writes, w.Len(), w.spoolSize)
	}
}

type shapingCancelWriter struct {
	bytes.Buffer
	t         *testing.T
	cancel    context.CancelFunc
	writes    int
	spoolSize int64
}

func (w *shapingCancelWriter) Write(b []byte) (int, error) {
	w.writes++
	entries, err := os.ReadDir(os.TempDir())
	if err != nil || len(entries) != 1 {
		w.t.Fatalf("expected one spool during commit: %v, %v", entries, err)
	}
	info, err := os.Stat(filepath.Join(os.TempDir(), entries[0].Name()))
	if err != nil {
		w.t.Fatal(err)
	}
	if info.Mode().Perm()&0077 != 0 {
		w.t.Fatalf("spool permissions expose data: %v", info.Mode())
	}
	w.spoolSize = info.Size()
	n, err := w.Buffer.Write(b)
	w.cancel()
	return n, err
}

func TestShapingContextReaderCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	r := shapingContextReader{ctx, strings.NewReader("first then second")}
	buf := make([]byte, 5)
	if n, err := r.Read(buf); n != 5 || err != nil {
		t.Fatalf("initial read: %d, %v", n, err)
	}
	cancel()
	if n, err := r.Read(buf); n != 0 || !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled read: %d, %v", n, err)
	}
	if _, ok := any(r).(io.WriterTo); ok {
		t.Fatal("context reader must not permit bypassing Read checks")
	}
}
