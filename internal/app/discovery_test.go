package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/app"
	"github.com/flexera-public/flexera-cli/internal/cli"
)

type offlineDoer struct{ t *testing.T }

func (d offlineDoer) Do(*http.Request) (*http.Response, error) {
	d.t.Fatal("offline discovery attempted HTTP")
	return nil, nil
}

func schemaRun(t *testing.T, args ...string) (int, string, string) {
	t.Helper()
	var out, errOut bytes.Buffer
	root, deps := app.NewRootCmd(&out, &errOut, os.Getenv, offlineDoer{t}, "test")
	root.SetArgs(args)
	code := cli.Execute(context.Background(), root, deps, args)
	return code, out.String(), errOut.String()
}

func TestOfflineSchema(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("FLEXERA_CLI_CONFIG", "")
	t.Setenv("FLEXERA_CLI_ZONE", "invalid")
	t.Setenv("FLEXERA_CLI_OUTPUT", "json")
	t.Setenv("FLEXERA_CLI_JSON_STYLE", "compact")
	if err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".flexera"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(os.Getenv("HOME"), ".flexera", "config.yaml"), []byte("invalid: ["), 0600); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{{"cli", "schema", "budget", "create", "--example"}, {"cli", "schema", "budget", "create", "--part", "request", "--depth", "0"}, {"cli", "schema", "budget", "create", "--part", "params"}, {"cli", "schema", "budget", "get", "--part", "response"}} {
		code, out, stderr := schemaRun(t, args...)
		if code != 0 || !json.Valid([]byte(out)) {
			t.Fatalf("%v code=%d stdout=%s stderr=%s", args, code, out, stderr)
		}
	}
	code, _, stderr := schemaRun(t, "budget", "list")
	if code != 2 || !strings.Contains(stderr, "reading config") {
		t.Fatalf("API command ignored malformed implicit config: %d %s", code, stderr)
	}
}

func TestSchemaUsageErrors(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("FLEXERA_CLI_CONFIG", "")
	t.Setenv("FLEXERA_CLI_OUTPUT", "json")
	t.Setenv("FLEXERA_CLI_JSON_STYLE", "compact")
	for _, args := range [][]string{{"cli", "schema"}, {"cli", "schema", "absent", "create"}, {"cli", "schema", "budget"}, {"cli", "schema", "curated", "list"}, {"cli", "schema", "budget", "delete", "--example"}, {"cli", "schema", "budget", "create", "--depth", "-1"}, {"cli", "schema", "budget", "create", "--part", "invalid"}, {"cli", "schema", "budget", "create", "--example", "--part", "request"}, {"cli", "schema", "budget", "create", "--example", "-o", "table"}, {"cli", "schema", "budget", "create", "--example", "--out-jq", "."}, {"cli", "schema", "budget", "create", "--example", "--out-fields", "name"}, {"cli", "schema", "budget", "create", "--example", "--out-jq", ".", "--raw-output"}} {
		code, out, stderr := schemaRun(t, args...)
		if code != 2 || out != "" || stderr == "" {
			t.Errorf("%v code=%d out=%s err=%s", args, code, out, stderr)
		}
	}
}

func TestSchemaConfigAndAliases(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("FLEXERA_CLI_CONFIG", "")
	t.Setenv("FLEXERA_CLI_OUTPUT", "json")
	t.Setenv("FLEXERA_CLI_JSON_STYLE", "compact")
	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(path, []byte("invalid: ["), 0600); err != nil {
		t.Fatal(err)
	}
	code, out, _ := schemaRun(t, "cli", "schema", "budget", "create", "--config", path)
	if code != 2 || out != "" {
		t.Fatal("malformed explicit config accepted")
	}
	if err := os.WriteFile(path, []byte("zone: invalid\njson-style: pretty\noutput: json\n"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("FLEXERA_CLI_JSON_STYLE", "")
	code, out, stderr := schemaRun(t, "cli", "schema", "budget", "create", "--part", "request", "--depth", "0", "--config", path)
	if code != 0 || !strings.Contains(out, "\n  ") {
		t.Fatalf("explicit style ignored %d %s %s", code, out, stderr)
	}
	var stdout, stderrBuf bytes.Buffer
	root, deps := app.NewRootCmd(&stdout, &stderrBuf, os.Getenv, offlineDoer{t}, "test")
	for _, c := range root.Commands() {
		if c.Name() == "budget" {
			c.Aliases = append(c.Aliases, "budgets")
		}
	}
	root.SetArgs([]string{"cli", "schema", "budgets", "create", "--part", "request", "--depth", "0", "--json-style", "compact"})
	if code := cli.Execute(context.Background(), root, deps); code != 0 || !json.Valid(stdout.Bytes()) {
		t.Fatalf("alias resolution failed: %d %s", code, stderrBuf.String())
	}
}

func TestOfflineStyleEnvironmentAndFlagPrecedence(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("FLEXERA_CLI_CONFIG", "")
	t.Setenv("FLEXERA_CLI_JSON_STYLE", "pretty")
	t.Setenv("FLEXERA_CLI_OUTPUT", "json")
	code, out, stderr := schemaRun(t, "cli", "schema", "budget", "create", "--part", "request", "--depth", "0")
	if code != 0 || !strings.Contains(out, "\n  ") {
		t.Fatalf("env style ignored: %d %s %s", code, out, stderr)
	}
	code, out, stderr = schemaRun(t, "cli", "schema", "budget", "create", "--part", "request", "--depth", "0", "--json-style", "compact")
	if code != 0 || strings.Contains(out, "\n  ") {
		t.Fatalf("flag did not override env: %d %s %s", code, out, stderr)
	}
}

func TestRenamedCommandsKeepLegacyPaths(t *testing.T) {
	var out, errOut bytes.Buffer
	root, _ := app.NewRootCmd(&out, &errOut, os.Getenv, offlineDoer{t}, "test")
	for canonical, legacy := range map[string]string{
		"saas data upload":                                             "saas saa-s-data update",
		"credential aws-sts-credential get":                            "credential awssts-credential get",
		"credential oauth2-credential project replace":                 "credential o-auth2-credential project replace",
		"unified-onboarding saas-connector list":                       "unified-onboarding saa-s-connector list",
		"unified-onboarding saas-onboarding validate":                  "unified-onboarding saa-s-onboarding validate",
		"recommendations update-status":                                "recommendations updatestatus",
		"iam organization-invitation resend":                           "iam organization-invitation update",
		"bill-analysis bill-months reprocess":                          "bill-analysis bill-months update",
		"finops-onboarding processing-history reprocess":               "finops-onboarding processing-history update",
		"finops-billing shared-cost-rules reorder":                     "finops-billing shared-cost-rules operations-reorder",
		"it-visibility query execute":                                  "it-visibility query get-all",
		"iam contracts update-last-seen":                               "iam contracts update",
		"iam contracts set-activation":                                 "iam contracts activation",
		"iam user-profile get":                                         "iam user-profile list",
		"iam user-setting-blob download-url":                           "iam user-setting-blob list",
		"iam user-setting-blob upload-url":                             "iam user-setting-blob replace",
		"finops-onboarding bill-connect validate-all":                  "finops-onboarding bill-connect validations",
		"finops-onboarding bill-connect gcp validate":                  "finops-onboarding bill-connect gcp validation",
		"finops-customizations rule-based-dimension from-csv generate": "finops-customizations rule-based-dimension from_csv generate",
	} {
		want, _, err := root.Find(strings.Fields(canonical))
		if err != nil || want.CommandPath() != "flexera-cli "+canonical {
			t.Errorf("canonical %q: got %v, %v", canonical, want.CommandPath(), err)
			continue
		}
		if got, _, err := root.Find(strings.Fields(legacy)); err != nil || got != want {
			t.Errorf("legacy %q resolved to %v, %v; want %q", legacy, got.CommandPath(), err, canonical)
		}
	}
}
