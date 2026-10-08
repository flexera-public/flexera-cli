package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// finalHelpPrefix is the registered invocation prefix for a tag command,
// mirroring writeRegister (service first, nested tags renamed).
func finalHelpPrefix(tag genTag) string {
	return "flexera-cli " + strings.Join(tag.finalPath(), " ") + " "
}

// rewriteStagedHelp runs after registration and before audit/publication. It
// changes only string literals assigned to Example in Cobra Command composites;
// source-offset edits preserve comments and unrelated literals verbatim. Files
// with rewritten help are formatted before they can be published.
func rewriteStagedHelp(stage string, tags []genTag) error {
	for _, tag := range tags {
		from, to := "flexera-cli "+tag.Cmd+" ", finalHelpPrefix(tag)
		if from == to {
			continue
		}
		path := filepath.Join(stage, filepath.FromSlash(tag.Pkg), "cmd_gen.go")
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("staged help %s: %w", path, err)
		}
		updated, err := rewriteCobraExamples(data, from, to)
		if err != nil {
			return fmt.Errorf("staged help %s: %w", path, err)
		}
		if string(updated) == string(data) {
			continue
		}
		if err := os.WriteFile(path, updated, 0o644); err != nil {
			return fmt.Errorf("staged help %s: %w", path, err)
		}
	}
	return nil
}

func rewriteCobraExamples(data []byte, from, to string) ([]byte, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "cmd_gen.go", data, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	// Recognize aliases by their import path, not by the spelling "cobra".
	aliases := map[string]bool{}
	for _, imp := range file.Imports {
		path, err := strconv.Unquote(imp.Path.Value)
		if err != nil || path != "github.com/spf13/cobra" {
			continue
		}
		name := "cobra"
		if imp.Name != nil {
			name = imp.Name.Name
		}
		aliases[name] = true
	}
	type edit struct {
		start, end int
		text       string
	}
	var edits []edit
	ast.Inspect(file, func(node ast.Node) bool {
		composite, ok := node.(*ast.CompositeLit)
		if !ok {
			return true
		}
		selector, ok := composite.Type.(*ast.SelectorExpr)
		if !ok || selector.Sel.Name != "Command" {
			return true
		}
		pkg, ok := selector.X.(*ast.Ident)
		if !ok || !aliases[pkg.Name] {
			return true
		}
		for _, element := range composite.Elts {
			field, ok := element.(*ast.KeyValueExpr)
			if !ok {
				continue
			}
			key, ok := field.Key.(*ast.Ident)
			if !ok || key.Name != "Example" {
				continue
			}
			literal, ok := field.Value.(*ast.BasicLit)
			if !ok || literal.Kind != token.STRING {
				continue
			}
			value, unquoteErr := strconv.Unquote(literal.Value)
			if unquoteErr != nil {
				err = unquoteErr
				return false
			}
			updated := strings.ReplaceAll(value, from, to)
			updated = strings.ReplaceAll(updated, strings.Replace(from, "flexera-cli ", "flexera-cli cli schema ", 1), strings.Replace(to, "flexera-cli ", "flexera-cli cli schema ", 1))
			if updated != value {
				edits = append(edits, edit{fset.Position(literal.Pos()).Offset, fset.Position(literal.End()).Offset, strconv.Quote(updated)})
			}
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	if len(edits) == 0 {
		return data, nil
	}
	// AST traversal visits literals in source order; apply backwards so offsets
	// remain valid even when a replacement has a different length.
	updated := append([]byte(nil), data...)
	for i := len(edits) - 1; i >= 0; i-- {
		e := edits[i]
		replacement := append([]byte(e.text), updated[e.end:]...)
		updated = append(updated[:e.start], replacement...)
	}
	return format.Source(updated)
}
