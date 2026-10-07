package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type publicationArtifact struct {
	Staged      string
	Destination string
}

// publishArtifacts publishes in order and rolls back in reverse order. Backups
// are siblings of their destinations so their renames stay on one filesystem.
// Failed rollback leaves recovery data intact and reports its paths.
func publishArtifacts(artifacts []publicationArtifact, rename func(string, string) error) error {
	type publicationState struct {
		publicationArtifact
		backupDir, backup   string
		backedUp, published bool
	}
	states := make([]publicationState, 0, len(artifacts))
	rollback := func(cause error) error {
		for i := len(states) - 1; i >= 0; i-- {
			s := &states[i]
			if s.published {
				if err := rename(s.Destination, s.Staged); err != nil {
					cause = errors.Join(cause, fmt.Errorf("rollback %s to %s failed (backup path %q): %w", s.Destination, s.Staged, s.backup, err))
					continue // Never overwrite a publication that could not be moved.
				}
			}
			if s.backedUp {
				if err := rename(s.backup, s.Destination); err != nil {
					cause = errors.Join(cause, fmt.Errorf("rollback %s failed; old artifact preserved at %s (staged at %s): %w", s.Destination, s.backup, s.Staged, err))
					continue
				}
			}
			if s.backupDir != "" {
				if err := os.RemoveAll(s.backupDir); err != nil {
					cause = errors.Join(cause, fmt.Errorf("clean backup directory %s: %w", s.backupDir, err))
				}
			}
		}
		return cause
	}
	for _, artifact := range artifacts {
		states = append(states, publicationState{publicationArtifact: artifact})
		s := &states[len(states)-1]
		if _, err := os.Lstat(s.Destination); err == nil {
			s.backupDir, err = os.MkdirTemp(filepath.Dir(s.Destination), ".publication-backup-")
			if err != nil {
				return rollback(fmt.Errorf("create backup for %s: %w", s.Destination, err))
			}
			s.backup = filepath.Join(s.backupDir, filepath.Base(s.Destination))
			if err := rename(s.Destination, s.backup); err != nil {
				return rollback(fmt.Errorf("backup %s to %s: %w", s.Destination, s.backup, err))
			}
			s.backedUp = true
		} else if !errors.Is(err, os.ErrNotExist) {
			return rollback(fmt.Errorf("inspect %s: %w", s.Destination, err))
		}
		if err := rename(s.Staged, s.Destination); err != nil {
			return rollback(fmt.Errorf("publish %s to %s: %w", s.Staged, s.Destination, err))
		}
		s.published = true
	}
	var cleanupErr error
	for _, s := range states {
		if s.backupDir != "" {
			if err := os.RemoveAll(s.backupDir); err != nil {
				cleanupErr = errors.Join(cleanupErr, fmt.Errorf("clean backup directory %s: %w", s.backupDir, err))
			}
		}
	}
	return cleanupErr
}
