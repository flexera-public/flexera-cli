package app

import (
	"io"
	"testing"
)

func TestNewRootCmdWithUserAgentVersion(t *testing.T) {
	displayVersion := "1.2.3 (commit abc123, built today)"
	root, deps := NewRootCmdWithUserAgentVersion(io.Discard, io.Discard, nil, nil, displayVersion, "1.2.3")

	if root.Version != displayVersion {
		t.Fatalf("root version = %q, want %q", root.Version, displayVersion)
	}
	if got := deps.Factory().UserAgentVersion; got != "1.2.3" {
		t.Fatalf("user-agent version = %q, want %q", got, "1.2.3")
	}
}
