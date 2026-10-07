package cli

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	flexera "github.com/flexera-public/unified-go-client"
)

func TestPrinterJSONStyles(t *testing.T) {
	for _, tc := range []struct {
		name        string
		style       JSONStyle
		tty, pretty bool
	}{
		{"zero-buffer", "", false, false},
		{"zero-tty", "", true, true},
		{"auto-buffer", JSONStyleAuto, false, false},
		{"auto-tty", JSONStyleAuto, true, true},
		{"pretty-buffer", JSONStylePretty, false, true},
		{"pretty-tty", JSONStylePretty, true, true},
		{"compact-buffer", JSONStyleCompact, false, false},
		{"compact-tty", JSONStyleCompact, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			p := Printer{Style: tc.style, IsTTY: func(w io.Writer) bool {
				if w != &out {
					t.Fatal("TTY check received wrong writer")
				}
				return tc.tty
			}}
			if err := p.Render(&out, "json", map[string]any{"html": "<>&", "id": int64(9007199254740993)}); err != nil {
				t.Fatal(err)
			}
			want := "{\"html\":\"<>&\",\"id\":9007199254740993}\n"
			if tc.pretty {
				want = "{\n  \"html\": \"<>&\",\n  \"id\": 9007199254740993\n}\n"
			}
			if out.String() != want {
				t.Fatalf("got %q, want %q", out.String(), want)
			}
		})
	}
}

func TestPrinterActualDestination(t *testing.T) {
	var stdout, stderr bytes.Buffer
	p := Printer{IsTTY: func(w io.Writer) bool { return w == &stderr }}
	if err := p.Render(&stdout, "json", map[string]any{"ok": true}); err != nil {
		t.Fatal(err)
	}
	if err := p.RenderError(&stderr, errors.New("bad \"value\" <>&")); err != nil {
		t.Fatal(err)
	}
	if stdout.String() != "{\"ok\":true}\n" {
		t.Fatal(stdout.String())
	}
	if stderr.String() != "{\n  \"error\": \"bad \\\"value\\\" <>&\"\n}\n" {
		t.Fatal(stderr.String())
	}
}

func TestPrinterDefaultDetection(t *testing.T) {
	var out bytes.Buffer
	if writerIsTTY(&out) {
		t.Fatal("buffer is not a terminal")
	}
	f, err := os.CreateTemp(t.TempDir(), "output")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if writerIsTTY(f) {
		t.Fatal("regular file is not a terminal")
	}
	if err := (Printer{}).Render(&out, "", nil); err != nil {
		t.Fatal(err)
	}
	if out.String() != "null\n" {
		t.Fatal(out.String())
	}
}

type failingWriter struct{ err error }

func (w failingWriter) Write([]byte) (int, error) { return 0, w.err }

func TestPrinterFailures(t *testing.T) {
	sentinel := errors.New("write failed")
	p := Printer{}
	if err := p.Render(failingWriter{sentinel}, "json", true); !errors.Is(err, sentinel) {
		t.Fatal(err)
	}
	if err := p.RenderError(failingWriter{sentinel}, sentinel); !errors.Is(err, sentinel) {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		p      Printer
		format string
		value  any
	}{
		{Printer{Style: "invalid"}, "json", true},
		{p, "yaml", true},
		{p, "json", make(chan int)},
	} {
		var out bytes.Buffer
		if err := tc.p.Render(&out, tc.format, tc.value); err == nil {
			t.Fatal("expected error")
		}
		if out.Len() != 0 {
			t.Fatal("failure wrote partial output")
		}
	}
}

func TestPrinterTableDelegatesToSDK(t *testing.T) {
	value := &flexera.IamUserList{}
	var got, want bytes.Buffer
	p := Printer{Style: JSONStylePretty, IsTTY: func(io.Writer) bool { t.Fatal("table queried TTY"); return false }}
	if err := p.Render(&got, " TABLE ", value); err != nil {
		t.Fatal(err)
	}
	if err := flexera.WriteTable(&want, value); err != nil {
		t.Fatal(err)
	}
	if got.String() != want.String() {
		t.Fatalf("got %q, want %q", got.String(), want.String())
	}
	if err := p.Render(io.Discard, "table", []string{"unsupported"}); err == nil {
		t.Fatal("expected SDK table rejection")
	}
}

func TestJSONStyleConfigPrecedence(t *testing.T) {
	for _, tc := range []struct {
		name, file, env, flag string
		want                  JSONStyle
	}{
		{"default", "", "", "", JSONStyleAuto},
		{"file", "pretty", "", "", JSONStylePretty},
		{"env", "pretty", "compact", "", JSONStyleCompact},
		{"flag", "pretty", "compact", "auto", JSONStyleAuto},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("FLEXERA_CLI_JSON_STYLE", tc.env)
			path := filepath.Join(t.TempDir(), "config.yaml")
			if err := os.WriteFile(path, []byte("json-style: "+tc.file+"\n"), 0600); err != nil {
				t.Fatal(err)
			}
			fs := newBoundFlags(t)
			if tc.flag != "" {
				if err := fs.Set(FlagJSONStyle, tc.flag); err != nil {
					t.Fatal(err)
				}
			}
			v, err := NewViper(fs, path)
			if err != nil {
				t.Fatal(err)
			}
			got, err := ResolveJSONStyle(v)
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestInvalidJSONStyleFailsResolve(t *testing.T) {
	t.Setenv("FLEXERA_CLI_JSON_STYLE", "invalid")
	fs := newBoundFlags(t)
	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(path, []byte("{}\n"), 0600); err != nil {
		t.Fatal(err)
	}
	v, err := NewViper(fs, path)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Resolve(v); err == nil || !strings.Contains(err.Error(), "JSON style") {
		t.Fatalf("got %v", err)
	}
}
