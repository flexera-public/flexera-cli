// Package coverage implements the unified-go-client "hand-written symbol"
// scanner used by cmd/checkcoverage.
//
// unified-go-client mixes two kinds of exported Go API:
//
//  1. Generated code (oapi-codegen), which has 1:1 CLI coverage via
//     cmd/gencli reading the OpenAPI spec directly -- new endpoints there
//     automatically get new generated commands.
//  2. Hand-written extensions (bill_upload.go, graphql.go, the anomaly
//     package, etc.) that have no OpenAPI operation backing them and so are
//     invisible to cmd/gencli. These only get CLI coverage when a human
//     notices them and writes a curated command (internal/curated/...).
//
// This package finds every exported top-level declaration in a hand-written
// file so that gap (2) can be tracked against a checked-in allowlist
// instead of relying on someone remembering to look. For the root package
// it prefers unified-go-client's own generator-maintained
// client_extensions_manifest.json (see manifest.go); everywhere else
// (sub-packages, and the root package on older checkouts without a
// manifest) it falls back to identifying hand-written files by the absence
// of the standard "// Code generated ... DO NOT EDIT." header used
// consistently across every generated file in the module.
package coverage

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// generatedHeader matches the standard generated-code marker. Every
// generated file in unified-go-client carries this on one of its first
// lines (oapi-codegen's own convention, the same one goimports/gofmt-adjacent
// tooling recognizes); see https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source.
var generatedHeader = regexp.MustCompile(`^// Code generated .* DO NOT EDIT\.$`)

// Symbol identifies one exported, hand-written declaration.
type Symbol struct {
	// Package is the import path relative to the module root ("." for the
	// root package, "anomaly", "service/graphql/v1", etc).
	Package string
	// Name is the declaration name. Methods are named "(Recv).Method"
	// (the receiver's pointer-ness is not encoded, since a type cannot have
	// both a value- and pointer-receiver method of the same name).
	Name string
	// Kind is "func", "method", "type", "var", or "const".
	Kind string
	// File is the path (relative to moduleDir) the symbol was declared in,
	// for error messages.
	File string
}

// Key returns the stable identifier used to look symbols up in an
// Allowlist: "<package>#<name>".
func (s Symbol) Key() string {
	return s.Package + "#" + s.Name
}

// isGenerated reports whether file starts with the standard generated-code
// header within its first few lines.
func isGenerated(path string) (bool, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	lines := strings.SplitN(string(b), "\n", 6)
	for _, line := range lines {
		if generatedHeader.MatchString(strings.TrimRight(line, "\r")) {
			return true, nil
		}
	}
	return false, nil
}

// Scan walks moduleDir (the root of an unified-go-client checkout or module
// cache directory) and returns every exported top-level func, method, type,
// var, and const declared in hand-written (non-generated) .go files. Test
// files, example/tool directories under cmd/, and vendor are skipped.
//
// For the root package, Scan prefers unified-go-client's own
// client_extensions_manifest.json (see loadExtensionsManifest) over
// AST-parsing root-package files itself, falling back to the AST scan only
// when no manifest is present. Sub-packages (service/*, rightscale/*, ...)
// have no equivalent manifest upstream, so they are always AST-scanned.
func Scan(moduleDir string) ([]Symbol, error) {
	rootSymbols, haveManifest, err := loadExtensionsManifest(moduleDir)
	if err != nil {
		return nil, err
	}

	symbols := append([]Symbol(nil), rootSymbols...)
	err = filepath.WalkDir(moduleDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if name == "vendor" || name == "testdata" || strings.HasPrefix(name, ".") {
				return filepath.SkipDir
			}
			// cmd/ under the client module holds standalone example/dev
			// binaries (e.g. cmd/split-client), not library API surface.
			if path != moduleDir && name == "cmd" {
				return filepath.SkipDir
			}
			return nil
		}
		if haveManifest && filepath.Dir(path) == moduleDir {
			// Root-package symbols already came from the manifest above.
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		generated, err := isGenerated(path)
		if err != nil {
			return err
		}
		if generated {
			return nil
		}
		fileSymbols, err := scanFile(moduleDir, path)
		if err != nil {
			return err
		}
		symbols = append(symbols, fileSymbols...)
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(symbols, func(i, j int) bool {
		if symbols[i].Package != symbols[j].Package {
			return symbols[i].Package < symbols[j].Package
		}
		return symbols[i].Name < symbols[j].Name
	})
	return symbols, nil
}

func scanFile(moduleDir, path string) ([]Symbol, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}
	pkg := packagePath(moduleDir, path)
	rel, err := filepath.Rel(moduleDir, path)
	if err != nil {
		rel = path
	}

	var symbols []Symbol
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if !d.Name.IsExported() {
				continue
			}
			if d.Recv == nil || len(d.Recv.List) == 0 {
				symbols = append(symbols, Symbol{Package: pkg, Name: d.Name.Name, Kind: "func", File: rel})
				continue
			}
			recv := strings.TrimPrefix(receiverName(d.Recv.List[0].Type), "*")
			if !ast.IsExported(recv) {
				continue
			}
			// Receiver pointer-ness is dropped from Name (Go forbids a type
			// from having both a value- and pointer-receiver method of the
			// same name, so this can't collide), matching the shape of
			// unified-go-client's client_extensions_manifest.json, which
			// records only the bare receiver type name.
			symbols = append(symbols, Symbol{Package: pkg, Name: "(" + recv + ")." + d.Name.Name, Kind: "method", File: rel})
		case *ast.GenDecl:
			switch d.Tok {
			case token.TYPE:
				for _, spec := range d.Specs {
					ts, ok := spec.(*ast.TypeSpec)
					if !ok || !ts.Name.IsExported() {
						continue
					}
					symbols = append(symbols, Symbol{Package: pkg, Name: ts.Name.Name, Kind: "type", File: rel})
				}
			case token.VAR, token.CONST:
				kind := "var"
				if d.Tok == token.CONST {
					kind = "const"
				}
				for _, spec := range d.Specs {
					vs, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}
					for _, name := range vs.Names {
						if name.IsExported() {
							symbols = append(symbols, Symbol{Package: pkg, Name: name.Name, Kind: kind, File: rel})
						}
					}
				}
			}
		}
	}
	return symbols, nil
}

func receiverName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.StarExpr:
		return "*" + receiverName(t.X)
	case *ast.Ident:
		return t.Name
	case *ast.IndexExpr:
		return receiverName(t.X)
	case *ast.IndexListExpr:
		return receiverName(t.X)
	default:
		return ""
	}
}

// packagePath returns path's directory relative to moduleDir, using "." for
// the module root itself, matching the Package field convention used
// throughout this package and the checked-in allowlist.
func packagePath(moduleDir, path string) string {
	dir := filepath.Dir(path)
	rel, err := filepath.Rel(moduleDir, dir)
	if err != nil {
		return dir
	}
	rel = filepath.ToSlash(rel)
	if rel == "." {
		return "."
	}
	return rel
}
