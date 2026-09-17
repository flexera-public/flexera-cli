package coverage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// extensionsManifestFileName is the file unified-go-client's `make
// generate-client` (cmd/split-client) emits alongside its generated code:
// an inventory of every exported, hand-written declaration in the root
// package, mirroring client_gen_manifest.json's coverage of generated ones.
// See the "client_extensions_manifest.json" section of unified-go-client's
// README.
const extensionsManifestFileName = "client_extensions_manifest.json"

// clientExtensionsManifest mirrors the JSON shape cmd/split-client emits.
// Only the fields this package needs are declared here.
type clientExtensionsManifest struct {
	Files []clientExtensionsManifestFile `json:"files"`
}

type clientExtensionsManifestFile struct {
	Path         string                        `json:"path"`
	Declarations []clientExtensionsDeclaration `json:"declarations"`
}

type clientExtensionsDeclaration struct {
	Name     string `json:"name"`
	Kind     string `json:"kind"` // func, method, type, var, const
	Receiver string `json:"receiver,omitempty"`
}

// loadExtensionsManifest reads moduleDir's client_extensions_manifest.json,
// if present, and converts its entries into root-package Symbols. It
// returns ok=false (with no error) when the file doesn't exist, so Scan can
// fall back to AST-parsing root-package files directly -- e.g. for older
// unified-go-client checkouts that predate the manifest, or the synthetic
// fixtures this package's own tests construct.
//
// Using the manifest instead of AST-parsing root-package files directly
// means Scan trusts unified-go-client's own generator-maintained inventory
// rather than re-deriving it via the generated-header-regex heuristic
// isGenerated relies on elsewhere in this package.
func loadExtensionsManifest(moduleDir string) ([]Symbol, bool, error) {
	path := filepath.Join(moduleDir, extensionsManifestFileName)
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	var m clientExtensionsManifest
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, false, fmt.Errorf("parse %s: %w", path, err)
	}

	var symbols []Symbol
	for _, f := range m.Files {
		for _, d := range f.Declarations {
			name := d.Name
			if d.Receiver != "" {
				name = "(" + d.Receiver + ")." + d.Name
			}
			symbols = append(symbols, Symbol{Package: ".", Name: name, Kind: d.Kind, File: f.Path})
		}
	}
	return symbols, true, nil
}
