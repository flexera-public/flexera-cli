package finops

import (
	"bytes"
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

func TestOptimaEmptyBodyAndTableContracts(t *testing.T) {
	var out bytes.Buffer
	deps := &clipkg.Deps{Stdout: &out, Config: cliconfig.CommonConfig{Output: "table"}}
	if err := renderJSONBody(deps, nil); err != nil || out.Len() != 0 {
		t.Fatalf("empty body: %v, %q", err, out.String())
	}
	if err := renderJSONBody(deps, []byte(`[1,2]`)); err == nil {
		t.Fatal("unsupported table must not silently emit JSON")
	}
	if out.Len() != 0 {
		t.Fatal("failed response wrote output")
	}
}
