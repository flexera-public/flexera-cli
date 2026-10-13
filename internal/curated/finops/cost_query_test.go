package finops

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
)

func TestCostQueryCmdFlags(t *testing.T) {
	cmd := newCostQueryCmd()
	expectedFlags := []string{
		"start",
		"start-at",
		"end",
		"end-at",
		"granularity",
		"endpoint",
		"dimensions",
		"metrics",
		"billing-center-ids",
		"no-chunk",
		"summarized",
		"export",
		"no-wait",
		"export-poll-interval",
		"line-items",
		"limit",
		"period-type",
		"adj-dimension-granularity",
		"no-slice",
		"progress-interval",
		"filter",
		"body",
		"optima-base-url",
	}

	for _, name := range expectedFlags {
		flag := cmd.Flags().Lookup(name)
		if flag == nil {
			t.Errorf("missing expected flag --%s", name)
		}
	}

	// Verify defaults
	if d := cmd.Flags().Lookup("granularity").DefValue; d != "month" {
		t.Errorf("expected granularity default 'month', got %q", d)
	}
	if d := cmd.Flags().Lookup("endpoint").DefValue; d != "auto" {
		t.Errorf("expected endpoint default 'auto', got %q", d)
	}
	if d := cmd.Flags().Lookup("export-poll-interval").DefValue; d != (5 * time.Second).String() {
		t.Errorf("expected export-poll-interval default '5s', got %q", d)
	}
	if d := cmd.Flags().Lookup("progress-interval").DefValue; d != (30 * time.Second).String() {
		t.Errorf("expected progress-interval default '30s', got %q", d)
	}
}

func TestCostQueryValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr string
	}{
		{
			name:    "missing start and end",
			args:    []string{},
			wantErr: "--start and --end are required",
		},
		{
			name:    "missing end",
			args:    []string{"--start", "2024-01"},
			wantErr: "--start and --end are required",
		},
		{
			name:    "invalid endpoint",
			args:    []string{"--start", "2024-01", "--end", "2024-02", "--endpoint", "invalid"},
			wantErr: "invalid --endpoint",
		},
		{
			name:    "invalid granularity",
			args:    []string{"--start", "2024-01", "--end", "2024-02", "--granularity", "year"},
			wantErr: "invalid --granularity",
		},
		{
			name:    "line-items with aggregated endpoint",
			args:    []string{"--start", "2024-01", "--end", "2024-02", "--line-items", "--endpoint", "aggregated"},
			wantErr: "--line-items cannot be used with --endpoint aggregated",
		},
		{
			name:    "invalid period-type",
			args:    []string{"--start", "2024-01", "--end", "2024-02", "--period-type", "weekly"},
			wantErr: "invalid --period-type",
		},
		{
			name:    "invalid adj-dimension-granularity",
			args:    []string{"--start", "2024-01", "--end", "2024-02", "--adj-dimension-granularity", "invalid"},
			wantErr: "invalid --adj-dimension-granularity",
		},
		{
			name:    "negative limit",
			args:    []string{"--start", "2024-01", "--end", "2024-02", "--limit", "-1"},
			wantErr: "--limit must be non-negative",
		},
		{
			name:    "invalid filter JSON",
			args:    []string{"--start", "2024-01", "--end", "2024-02", "--filter", "not-json"},
			wantErr: "invalid --filter JSON",
		},
		{
			name:    "invalid body JSON",
			args:    []string{"--body", "not-json"},
			wantErr: "--body is not valid JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := newCostQueryCmd()
			var stdout, stderr bytes.Buffer
			deps := &clipkg.Deps{
				Stdout:  &stdout,
				Stderr:  &stderr,
				Config:  cliconfig.CommonConfig{OrgID: 100},
				Printer: clipkg.Printer{Style: clipkg.JSONStyleCompact},
			}
			cmd.SetContext(clipkg.WithDeps(context.Background(), deps))
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("expected error containing %q, got %q", tt.wantErr, err.Error())
			}
		})
	}
}
