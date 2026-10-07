package workflows

import (
	"bytes"
	"strings"
	"testing"
)

func TestCuratedListPreservesText(t *testing.T) {
	var out bytes.Buffer
	cmd := newListCmd()
	cmd.SetOut(&out)
	cmd.SetArgs(nil)
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(out.String(), "Available tools:\n  anomaly-investigation") {
		t.Fatalf("text contract changed: %q", out.String())
	}
}
