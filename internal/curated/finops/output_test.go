package finops

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
)

func TestRenderJSONBody(t *testing.T) {
	for _, style := range []clipkg.JSONStyle{clipkg.JSONStyleCompact, clipkg.JSONStylePretty} {
		t.Run(string(style), func(t *testing.T) {
			var out bytes.Buffer
			deps := &clipkg.Deps{
				Stdout:  &out,
				Config:  cliconfig.CommonConfig{Output: "json"},
				Printer: clipkg.Printer{Style: style},
			}
			body := []byte(`{"id":9007199254740993,"date":"","html":"<>&"}`)
			if err := renderJSONBody(deps, body); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(out.String(), "9007199254740993") || !strings.Contains(out.String(), "<>&") {
				t.Fatalf("response changed values: %s", &out)
			}
			if strings.Contains(out.String(), "\n  ") != (style == clipkg.JSONStylePretty) {
				t.Fatalf("wrong style: %q", out.String())
			}
		})
	}
}

func TestRenderJSONBodyRejectsMalformedWithoutOutput(t *testing.T) {
	for _, body := range []string{"not JSON", `{"ok":true} {}`, `{"ok":true} trailing`, " "} {
		var out bytes.Buffer
		deps := &clipkg.Deps{Stdout: &out, Config: cliconfig.CommonConfig{Output: "json"}}
		if err := renderJSONBody(deps, []byte(body)); err == nil {
			t.Fatalf("expected error for %q", body)
		}
		if out.Len() != 0 {
			t.Fatal("malformed response wrote output")
		}
	}
}

func TestOptimaEmptyStatusAndTableContracts(t *testing.T) {
	var out bytes.Buffer
	deps := &clipkg.Deps{Stdout: &out, Config: cliconfig.CommonConfig{Output: "table"}}
	if err := renderJSONBody(deps, nil); err != nil || out.Len() != 0 {
		t.Fatalf("empty body: %v, %q", err, out.String())
	}
	if err := renderJSONBody(deps, []byte(`[1,2]`)); err == nil {
		t.Fatal("unsupported table must not silently emit JSON")
	}
	if err := emitOptima(deps, &http.Response{StatusCode: 500}, []byte(`{}`)); err == nil {
		t.Fatal("expected status error")
	}
	if out.Len() != 0 {
		t.Fatal("failed response wrote output")
	}
}

type outputFailure struct{}

func (outputFailure) Write([]byte) (int, error) { return 0, io.ErrClosedPipe }

func TestFinopsDryRunUsesFormatterNotOutputFormat(t *testing.T) {
	var out bytes.Buffer
	deps := &clipkg.Deps{
		Stdout:  &out,
		Config:  cliconfig.CommonConfig{Output: "table"},
		Printer: clipkg.Printer{Style: clipkg.JSONStyleCompact},
	}
	done, err := resolveWriteOp(true, false, true, deps, map[string]any{"body": "<>&"})
	if err != nil || !done {
		t.Fatalf("done=%v err=%v", done, err)
	}
	if out.String() != "{\"destructive\":true,\"dryRun\":true,\"plan\":{\"body\":\"<>&\"}}\n" {
		t.Fatal(out.String())
	}
	deps.Stdout = outputFailure{}
	if _, err := resolveWriteOp(true, false, true, deps, nil); !errors.Is(err, io.ErrClosedPipe) {
		t.Fatalf("writer error lost: %v", err)
	}
	if done, err := resolveWriteOp(false, false, true, deps, nil); done || err == nil {
		t.Fatalf("confirmation changed: done=%v err=%v", done, err)
	}
}
