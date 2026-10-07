package cli

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// newBoundFlags mirrors the root command's persistent flag set for testing
// viper binding in isolation.
func newBoundFlags(t *testing.T) *pflag.FlagSet {
	t.Helper()
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	RegisterPersistentFlags(fs)
	return fs
}

func TestResolvePrecedenceFlagBeatsEnv(t *testing.T) {
	t.Setenv("FLEXERA_CLI_ORG_ID", "111")
	t.Setenv("FLEXERA_CLI_ACCESS_TOKEN", "env-token")

	fs := newBoundFlags(t)
	if err := fs.Parse([]string{"--org-id", "222", "--output", "table"}); err != nil {
		t.Fatalf("parse: %v", err)
	}
	v, err := NewViper(fs, "")
	if err != nil {
		t.Fatalf("NewViper: %v", err)
	}
	cfg, err := Resolve(v)
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if cfg.OrgID != 222 {
		t.Errorf("org-id: flag should beat env, got %d want 222", cfg.OrgID)
	}
	if cfg.AccessToken != "env-token" {
		t.Errorf("access-token: env should apply when no flag, got %q", cfg.AccessToken)
	}
	if cfg.Output != "table" {
		t.Errorf("output: got %q want table", cfg.Output)
	}
}

func TestResolveEnvBeatsConfigFileBeatsDefault(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(cfgPath, []byte("org-id: 999\nzone: eu\noutput: json\n"), 0o600); err != nil {
		t.Fatalf("write config: %v", err)
	}
	// env overrides the file value for org-id; zone comes from file; output
	// default (json) when neither flag nor env set.
	t.Setenv("FLEXERA_CLI_ORG_ID", "777")

	fs := newBoundFlags(t)
	if err := fs.Parse(nil); err != nil {
		t.Fatalf("parse: %v", err)
	}
	v, err := NewViper(fs, cfgPath)
	if err != nil {
		t.Fatalf("NewViper: %v", err)
	}
	cfg, err := Resolve(v)
	if err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if cfg.OrgID != 777 {
		t.Errorf("org-id: env should beat config file, got %d want 777", cfg.OrgID)
	}
	if string(cfg.Zone) == "" {
		t.Errorf("zone: expected value from config file, got empty")
	}
	if cfg.Output != "json" {
		t.Errorf("output: expected default json, got %q", cfg.Output)
	}
}

func TestNewViperMissingExplicitConfigErrors(t *testing.T) {
	fs := newBoundFlags(t)
	_ = fs.Parse(nil)
	_, err := NewViper(fs, filepath.Join(t.TempDir(), "does-not-exist.yaml"))
	if err == nil {
		t.Fatal("expected error for missing explicit config file")
	}
}

func TestResolveBodyRawWins(t *testing.T) {
	typed := map[string]any{"name": "from-flags"}
	got, err := ResolveBody(`{"name":"from-body"}`, typed, nil)
	if err != nil {
		t.Fatalf("ResolveBody: %v", err)
	}
	if !strings.Contains(string(got), "from-body") {
		t.Errorf("raw --body should win, got %s", got)
	}
}

func TestResolveBodyTypedFallback(t *testing.T) {
	typed := map[string]any{"name": "from-flags"}
	got, err := ResolveBody("", typed, nil)
	if err != nil {
		t.Fatalf("ResolveBody: %v", err)
	}
	if !strings.Contains(string(got), "from-flags") {
		t.Errorf("typed body should be used when raw empty, got %s", got)
	}
}

func TestResolveBodyStdin(t *testing.T) {
	got, err := ResolveBody("@-", nil, strings.NewReader(`{"k":1}`))
	if err != nil {
		t.Fatalf("ResolveBody: %v", err)
	}
	if strings.TrimSpace(string(got)) != `{"k":1}` {
		t.Errorf("stdin body mismatch, got %s", got)
	}
}

func TestResolveBodyInvalidJSON(t *testing.T) {
	if _, err := ResolveBody("not json", nil, nil); err == nil {
		t.Fatal("expected invalid JSON error for raw --body")
	}
}

func TestRootConfiguresOutputShaping(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		want string
		code int
	}{
		{name: "combined", args: []string{"--out-fields", "id,name", "--out-jq", ".name"}, want: `"sample"` + "\n"},
		{name: "raw", args: []string{"--out-jq", ".name", "--raw-output"}, want: "sample\n"},
		{name: "raw-without-jq", args: []string{"--raw-output"}, code: 2},
		{name: "empty-jq", args: []string{"--out-jq", ""}, code: 2},
		{name: "invalid-jq", args: []string{"--out-jq", "["}, code: 2},
		{name: "empty-fields", args: []string{"--out-fields", ""}, code: 2},
		{name: "invalid-fields", args: []string{"--out-fields", "id,,name"}, code: 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config := filepath.Join(t.TempDir(), "config.yaml")
			if err := os.WriteFile(config, []byte("zone: nam\n"), 0o600); err != nil {
				t.Fatal(err)
			}
			var stdout, stderr bytes.Buffer
			root, deps := NewRootCmd(RootOptions{})
			root.SetOut(&stdout)
			root.SetErr(&stderr)
			root.AddCommand(&cobra.Command{
				Use: "shape",
				RunE: func(cmd *cobra.Command, _ []string) error {
					return DepsFrom(cmd.Context()).Printer.Render(cmd.OutOrStdout(), "json", map[string]any{"id": 4, "name": "sample", "ignored": true})
				},
			})
			args := append([]string{"--config", config}, tc.args...)
			args = append(args, "shape")
			root.SetArgs(args)
			code := Execute(context.Background(), root, deps)
			if code != tc.code {
				t.Fatalf("exit code %d, want %d; stderr=%s", code, tc.code, stderr.String())
			}
			if tc.code == 0 && stdout.String() != tc.want {
				t.Fatalf("stdout %q, want %q", stdout.String(), tc.want)
			}
		})
	}
}

func TestOutputShapingFlagsAreCLIOnly(t *testing.T) {
	for _, key := range []string{FlagOutJQ, FlagOutFields, FlagRawOutput} {
		for _, bound := range persistentKeys {
			if key == bound {
				t.Fatalf("--%s must not be bound as configuration", key)
			}
		}
	}
}
