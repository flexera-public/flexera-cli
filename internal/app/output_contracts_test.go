package app_test

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/flexera-public/flexera-cli/internal/app"
	"github.com/flexera-public/flexera-cli/internal/cli"
	"github.com/spf13/cobra"
)

func TestRealTreeUnsupportedShapingBeforeSideEffects(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("FLEXERA_CLI_CONFIG", "")
	if err := os.MkdirAll(filepath.Join(os.Getenv("HOME"), ".flexera"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(os.Getenv("HOME"), ".flexera", "config.yaml"), []byte("broken: ["), 0600); err != nil {
		t.Fatal(err)
	}
	for _, path := range [][]string{{"curated", "list"}, {"policy", "applied-policy", "log"}, {"budget", "delete"}, {"bill-months", "download"}, {"completion", "fish"}} {
		args := append(append([]string{}, path...), "--out-jq", ".")
		code, out, stderr := schemaRun(t, args...)
		if code != 2 || out != "" || !json.Valid([]byte(stderr)) || !strings.Contains(stderr, "does not support") {
			t.Errorf("output guard %v: %d %s %s", args, code, out, stderr)
		}
	}
}

func TestEveryGeneratedLeafHasOutputContract(t *testing.T) {
	root, _ := app.NewRootCmd(io.Discard, io.Discard, func(string) string { return "" }, nil, "test")
	var walk func(*cobra.Command)
	walk = func(c *cobra.Command) {
		if id := c.Annotations["flexera.operationId"]; id != "" && !strings.HasPrefix(id, "Auth_") && !(c.Parent() != nil && c.Parent().Parent() != nil && c.Parent().Parent().Name() == "policy") {
			switch c.Annotations[cli.OutputAnnotation] {
			case "structured", "text", "binary", "mixed":
			default:
				t.Errorf("missing generated output contract: %s", c.CommandPath())
			}
		}
		for _, child := range c.Commands() {
			walk(child)
		}
	}
	walk(root)
}
