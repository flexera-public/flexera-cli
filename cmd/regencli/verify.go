// Symbol verification: regencli/gencli derive Go identifiers used by
// generated CLI commands (client methods, request-body / params / enum
// types) purely from string manipulation of OpenAPI operation IDs. That
// derivation can silently drift from the actual unified-go-client package
// (renamed operationId, regenerated types, etc.) without ever failing —
// the mismatch would only surface as a `go build` error downstream, far
// from the regen step that caused it.
//
// verifyFile below closes that gap: it parses each generated cmd_gen.go
// file and type-checks every "flexera.X" / "client.X" identifier it
// references against the real, compiled unified-go-client package (loaded
// once via go/packages + go/types), catching name and arity mismatches at
// regen time instead of at build time or, worse, silently.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

const (
	flexeraClientPkg  = "github.com/flexera-public/unified-go-client"
	flexeraClientType = "ClientWithResponses"
)

// symbolVerifier holds the exported symbol table of the real
// unified-go-client package, resolved once via go/types, so every
// generated file can be checked against it cheaply.
type symbolVerifier struct {
	// pkgSymbols is every exported package-level identifier (types, funcs,
	// consts, vars) in unified-go-client, e.g. "BudgetBudgetCreateJSONRequestBody",
	// "CollectPages", "ResponseError".
	pkgSymbols map[string]bool
	// clientMethods is every method (including promoted/embedded ones) on
	// *flexera.ClientWithResponses, e.g. "BudgetBudgetCreateWithResponse".
	clientMethods map[string]*types.Func
}

// newSymbolVerifier type-checks the unified-go-client module dependency
// and builds the symbol table used to verify generated files against it.
func newSymbolVerifier() (*symbolVerifier, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo |
			packages.NeedSyntax | packages.NeedDeps,
	}
	pkgs, err := packages.Load(cfg, flexeraClientPkg)
	if err != nil {
		return nil, fmt.Errorf("loading %s: %w", flexeraClientPkg, err)
	}
	if len(pkgs) != 1 || pkgs[0].Types == nil {
		return nil, fmt.Errorf("could not load types for %s", flexeraClientPkg)
	}
	pkg := pkgs[0]
	if len(pkg.Errors) > 0 {
		msgs := make([]string, 0, len(pkg.Errors))
		for _, e := range pkg.Errors {
			msgs = append(msgs, e.Error())
		}
		return nil, fmt.Errorf("errors loading %s: %s", flexeraClientPkg, strings.Join(msgs, "; "))
	}

	sv := &symbolVerifier{
		pkgSymbols:    map[string]bool{},
		clientMethods: map[string]*types.Func{},
	}

	scope := pkg.Types.Scope()
	for _, name := range scope.Names() {
		if obj := scope.Lookup(name); obj != nil && obj.Exported() {
			sv.pkgSymbols[name] = true
		}
	}

	obj := scope.Lookup(flexeraClientType)
	if obj == nil {
		return nil, fmt.Errorf("%s: type %s not found", flexeraClientPkg, flexeraClientType)
	}
	tn, ok := obj.(*types.TypeName)
	if !ok {
		return nil, fmt.Errorf("%s: %s is not a type", flexeraClientPkg, flexeraClientType)
	}
	named, ok := tn.Type().(*types.Named)
	if !ok {
		return nil, fmt.Errorf("%s: %s is not a named type", flexeraClientPkg, flexeraClientType)
	}
	// *ClientWithResponses is what deps.APIClient() returns; its method set
	// includes methods promoted from any embedded interfaces/structs.
	mset := types.NewMethodSet(types.NewPointer(named))
	for i := 0; i < mset.Len(); i++ {
		if fn, ok := mset.At(i).Obj().(*types.Func); ok {
			sv.clientMethods[fn.Name()] = fn
		}
	}
	return sv, nil
}

// verifyFile parses a single generated cmd_gen.go file and reports every
// flexera.* / client.* reference that does not resolve against the real
// unified-go-client symbol table, plus call-site arity mismatches.
func (sv *symbolVerifier) verifyFile(path string) []string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return []string{fmt.Sprintf("%s: parse error: %v", path, err)}
	}

	var issues []string
	pos := func(p token.Pos) string {
		return fmt.Sprintf("%s:%d", path, fset.Position(p).Line)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.CallExpr:
			sel, ok := node.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			ident, ok := sel.X.(*ast.Ident)
			if !ok || ident.Name != "client" {
				return true
			}
			name := sel.Sel.Name
			fn, found := sv.clientMethods[name]
			if !found {
				issues = append(issues, fmt.Sprintf(
					"%s: client.%s: no such method on *flexera.%s (unified-go-client symbol drift)",
					pos(node.Pos()), name, flexeraClientType))
				return true
			}
			if msg := checkArity(fn, node); msg != "" {
				issues = append(issues, fmt.Sprintf("%s: client.%s: %s", pos(node.Pos()), name, msg))
			}
		case *ast.SelectorExpr:
			ident, ok := node.X.(*ast.Ident)
			if !ok || ident.Name != "flexera" {
				return true
			}
			if !sv.pkgSymbols[node.Sel.Name] {
				issues = append(issues, fmt.Sprintf(
					"%s: flexera.%s: no such exported symbol in unified-go-client (spec/client drift)",
					pos(node.Pos()), node.Sel.Name))
			}
		}
		return true
	})
	return issues
}

// checkArity compares a call site's argument count against the real
// method's signature, accounting for the trailing variadic
// "...RequestEditorFn" every WithResponse method takes.
func checkArity(fn *types.Func, call *ast.CallExpr) string {
	sig, ok := fn.Type().(*types.Signature)
	if !ok {
		return ""
	}
	nParams := sig.Params().Len()
	nArgs := len(call.Args)
	if sig.Variadic() {
		if nArgs < nParams-1 {
			return fmt.Sprintf("expects at least %d arg(s), call site has %d", nParams-1, nArgs)
		}
		return ""
	}
	if nArgs != nParams {
		return fmt.Sprintf("expects %d arg(s), call site has %d", nParams, nArgs)
	}
	return ""
}
