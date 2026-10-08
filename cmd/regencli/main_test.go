package main

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func writeTestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func treeContents(t *testing.T, root string) map[string]string {
	t.Helper()
	files := map[string]string{}
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		files[rel] = string(data)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func assertNoTemporaryTrees(t *testing.T, parent string) {
	t.Helper()
	entries, err := os.ReadDir(parent)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if entry.Name() != "commands" {
			t.Errorf("temporary tree left behind: %s", entry.Name())
		}
	}
}

// flatTags plans root-level tag commands without service groups.
func flatTags(names ...string) []genTag {
	used := map[string]bool{}
	tags := make([]genTag, 0, len(names))
	for _, name := range names {
		tags = append(tags, genTag{Tag: name, Pkg: uniquePkg(cleanIdent(name), used), Cmd: kebab(cleanIdent(name))})
	}
	return tags
}

func testHooks(t *testing.T) generationHooks {
	t.Helper()
	return generationHooks{
		generate: func(tag genTag, out string) error {
			writeTestFile(t, out, "package "+path.Base(tag.Pkg)+"\n")
			return nil
		},
		verify: func(stage string, tags []genTag) error {
			// Exercise the real file parser without loading the SDK or running gencli.
			return verifyFiles(&symbolVerifier{}, stage, tags)
		},
		rename: os.Rename,
	}
}

func TestRegenerateFailuresLeaveOldTreeUntouched(t *testing.T) {
	injected := errors.New("injected failure")
	for _, failure := range []string{"generation", "missing output", "symbol verification", "registration", "backup rename", "publish rename"} {
		t.Run(failure, func(t *testing.T) {
			parent := t.TempDir()
			destination := filepath.Join(parent, "commands")
			writeTestFile(t, filepath.Join(destination, "old", "cmd_gen.go"), "old command\n")
			writeTestFile(t, filepath.Join(destination, "register_gen.go"), "old registry\n")
			before := treeContents(t, destination)
			hooks := testHooks(t)
			generate := hooks.generate
			calls := 0
			hooks.generate = func(tag genTag, out string) error {
				calls++
				if failure == "missing output" {
					return nil
				}
				if err := generate(tag, out); err != nil {
					return err
				}
				if failure == "generation" && calls == 2 {
					return injected // A partial file and earlier tag are already staged.
				}
				if failure == "symbol verification" {
					writeTestFile(t, out, "package "+path.Base(tag.Pkg)+"\nvar _ = flexera.MissingSymbol\n")
				}
				return nil
			}
			if failure == "registration" {
				verify := hooks.verify
				hooks.verify = func(stage string, tags []genTag) error {
					if err := verify(stage, tags); err != nil {
						return err
					}
					return os.Mkdir(filepath.Join(stage, "register_gen.go"), 0o755)
				}
			}
			renameCalls := 0
			hooks.rename = func(from, to string) error {
				renameCalls++
				if failure == "backup rename" && renameCalls == 1 || failure == "publish rename" && renameCalls == 2 {
					return injected
				}
				return os.Rename(from, to)
			}
			err := regenerate(destination, flatTags("Alpha", "Beta", "Gamma"), hooks)
			if err == nil {
				t.Fatal("expected failure")
			}
			if failure == "generation" || failure == "backup rename" || failure == "publish rename" {
				if !errors.Is(err, injected) {
					t.Fatalf("lost injected error: %v", err)
				}
			}
			if failure == "generation" && calls != 2 {
				t.Errorf("generation did not abort immediately: %d calls", calls)
			}
			if after := treeContents(t, destination); !reflect.DeepEqual(before, after) {
				t.Errorf("old outputs changed: before %v, after %v", before, after)
			}
			assertNoTemporaryTrees(t, parent)
		})
	}
}

func TestCatalogPreparationFailureLeavesTreeUntouched(t *testing.T) {
	parent := t.TempDir()
	destination := filepath.Join(parent, "commands")
	writeTestFile(t, filepath.Join(destination, "register_gen.go"), "old registry")
	before := treeContents(t, destination)
	hooks := testHooks(t)
	injected := errors.New("injected coverage failure")
	hooks.prepare = func(stage string, tags []genTag) ([]publicationArtifact, error) { return nil, injected }
	if err := regenerate(destination, flatTags("Budget"), hooks); !errors.Is(err, injected) {
		t.Fatalf("lost preparation error: %v", err)
	}
	if !reflect.DeepEqual(before, treeContents(t, destination)) {
		t.Fatal("preparation failure changed live outputs")
	}
	assertNoTemporaryTrees(t, parent)
}

func TestRegeneratePublishesCompleteTree(t *testing.T) {
	for _, existing := range []bool{false, true} {
		t.Run(map[bool]string{false: "first generation", true: "replacement"}[existing], func(t *testing.T) {
			parent := t.TempDir()
			destination := filepath.Join(parent, "commands")
			if existing {
				writeTestFile(t, filepath.Join(destination, "stale", "cmd_gen.go"), "stale")
				writeTestFile(t, filepath.Join(destination, "register_gen.go"), "old registry")
			}
			hooks := testHooks(t)
			verify := hooks.verify
			hooks.verify = func(stage string, tags []genTag) error {
				if filepath.Dir(stage) != parent || stage == destination {
					t.Errorf("not staging in sibling tree: %s", stage)
				}
				if existing && treeContents(t, destination)["register_gen.go"] != "old registry" {
					t.Error("live registry changed before verification")
				}
				if _, err := os.Stat(filepath.Join(stage, "register_gen.go")); !errors.Is(err, os.ErrNotExist) {
					t.Errorf("registry written before verification: %v", err)
				}
				return verify(stage, tags)
			}
			tags := fixtureTags(t)
			if err := regenerate(destination, tags, hooks); err != nil {
				t.Fatal(err)
			}
			files := treeContents(t, destination)
			if len(files) != len(tags)+1 {
				t.Fatalf("expected %d packages and registry, got %v", len(tags), files)
			}
			registry := files["register_gen.go"]
			for _, expected := range []string{
				`finopsonboardingbillconnect "` + modulePath + `/internal/commands/finopsonboarding/billconnect"`,
				`finopsonboardingbillconnectaws "` + modulePath + `/internal/commands/finopsonboarding/billconnectaws"`,
				`budgetbudget "` + modulePath + `/internal/commands/budget/budget"`,
				`nest(finopsonboardingbillconnectaws.NewCmd(), "aws")`,
				`service(&cobra.Command{Use: "finops-onboarding"}, "finops-onboarding"`,
				`service(budgetbudget.NewCmd(), "budget"`,
			} {
				if !strings.Contains(registry, expected) {
					t.Errorf("registry missing %q", expected)
				}
			}
			if strings.Contains(registry, parent) || strings.Contains(registry, ".commands-stage-") {
				t.Error("registry contains temporary import paths")
			}
			assertNoTemporaryTrees(t, parent)
		})
	}
}

func TestPublishPreservesBackupWhenRollbackFails(t *testing.T) {
	parent := t.TempDir()
	destination := filepath.Join(parent, "commands")
	stage := filepath.Join(parent, "stage")
	writeTestFile(t, filepath.Join(destination, "register_gen.go"), "old registry")
	writeTestFile(t, filepath.Join(stage, "register_gen.go"), "new registry")
	injected := errors.New("rename failed")
	calls := 0
	err := publish(stage, destination, func(from, to string) error {
		calls++
		if calls > 1 {
			return injected
		}
		return os.Rename(from, to)
	})
	if !errors.Is(err, injected) || !strings.Contains(err.Error(), "old commands preserved at") {
		t.Fatalf("expected actionable rollback error, got %v", err)
	}
	backups, err := filepath.Glob(filepath.Join(parent, ".commands-backup-*", "commands"))
	if err != nil || len(backups) != 1 {
		t.Fatalf("backup not retained: %v, %v", backups, err)
	}
	if treeContents(t, backups[0])["register_gen.go"] != "old registry" {
		t.Fatal("retained backup changed")
	}
}
