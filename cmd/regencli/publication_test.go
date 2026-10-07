package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func publicationFixture(t *testing.T, existing [2]bool) (string, []publicationArtifact) {
	t.Helper()
	root := t.TempDir()
	artifacts := []publicationArtifact{
		{Staged: filepath.Join(root, "commands-stage"), Destination: filepath.Join(root, "commands")},
		{Staged: filepath.Join(root, "catalog-stage.json"), Destination: filepath.Join(root, "catalog.json")},
	}
	for i, a := range artifacts {
		staged, destination := a.Staged, a.Destination
		if i == 0 {
			staged = filepath.Join(staged, "cmd_gen.go")
			destination = filepath.Join(destination, "cmd_gen.go")
		}
		writeTestFile(t, staged, "new")
		if existing[i] {
			writeTestFile(t, destination, "old")
		}
	}
	return root, artifacts
}

func assertPublicationContent(t *testing.T, path, want string) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.IsDir() {
		path = filepath.Join(path, "cmd_gen.go")
	}
	got, err := os.ReadFile(path)
	if err != nil || string(got) != want {
		t.Fatalf("%s: got %q, err %v; want %q", path, got, err, want)
	}
}

func assertNoPublicationBackups(t *testing.T, root string) {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join(root, ".publication-backup-*"))
	if err != nil || len(paths) != 0 {
		t.Fatalf("unexpected backups %v: %v", paths, err)
	}
}

func TestPublishArtifactsRenameFailures(t *testing.T) {
	for _, existing := range [][2]bool{{true, true}, {false, false}, {true, false}, {false, true}} {
		count := 2
		for _, present := range existing {
			if present {
				count++
			}
		}
		for failAt := 1; failAt <= count; failAt++ {
			t.Run(fmt.Sprintf("existing-%v/fail-%d", existing, failAt), func(t *testing.T) {
				root, artifacts := publicationFixture(t, existing)
				injected := errors.New("rename failed")
				calls := 0
				err := publishArtifacts(artifacts, func(from, to string) error {
					calls++
					if calls == failAt {
						return injected
					}
					return os.Rename(from, to)
				})
				if !errors.Is(err, injected) {
					t.Fatalf("expected injected error, got %v", err)
				}
				for i, a := range artifacts {
					if existing[i] {
						assertPublicationContent(t, a.Destination, "old")
					} else if _, err := os.Lstat(a.Destination); !errors.Is(err, os.ErrNotExist) {
						t.Fatalf("destination should be absent: %s: %v", a.Destination, err)
					}
					assertPublicationContent(t, a.Staged, "new")
				}
				assertNoPublicationBackups(t, root)
			})
		}
	}
}

func TestPublishArtifactsSuccess(t *testing.T) {
	for _, existing := range [][2]bool{{true, true}, {false, false}, {true, false}, {false, true}} {
		t.Run(fmt.Sprint(existing), func(t *testing.T) {
			root, artifacts := publicationFixture(t, existing)
			if err := publishArtifacts(artifacts, os.Rename); err != nil {
				t.Fatal(err)
			}
			for _, a := range artifacts {
				assertPublicationContent(t, a.Destination, "new")
				if _, err := os.Lstat(a.Staged); !errors.Is(err, os.ErrNotExist) {
					t.Fatalf("staged artifact remains: %s: %v", a.Staged, err)
				}
			}
			assertNoPublicationBackups(t, root)
		})
	}
}

func TestPublishArtifactsRollbackFailure(t *testing.T) {
	for _, failAt := range []int{5, 6, 7} {
		t.Run(fmt.Sprintf("rollback-rename-%d", failAt), func(t *testing.T) {
			root, artifacts := publicationFixture(t, [2]bool{true, true})
			publishErr, rollbackErr := errors.New("publish failed"), errors.New("rollback failed")
			calls := 0
			err := publishArtifacts(artifacts, func(from, to string) error {
				calls++
				if calls == 4 {
					return publishErr
				}
				if calls == failAt {
					return rollbackErr
				}
				return os.Rename(from, to)
			})
			if !errors.Is(err, publishErr) || !errors.Is(err, rollbackErr) {
				t.Fatalf("lost failure: %v", err)
			}
			backups, globErr := filepath.Glob(filepath.Join(root, ".publication-backup-*"))
			if globErr != nil || len(backups) != 1 {
				t.Fatalf("expected retained backup, got %v: %v", backups, globErr)
			}
			failed := 0
			if failAt == 5 {
				failed = 1
			}
			backup := filepath.Join(backups[0], filepath.Base(artifacts[failed].Destination))
			assertPublicationContent(t, backup, "old")
			if !strings.Contains(err.Error(), backup) || !strings.Contains(err.Error(), artifacts[failed].Destination) {
				t.Fatalf("missing recovery paths: %v", err)
			}
			assertPublicationContent(t, artifacts[1-failed].Destination, "old")
		})
	}
}
