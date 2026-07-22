package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// ExitError carries an explicit process exit code out of a command's RunE.
// Commands that need a non-1 exit return an *ExitError; everything else maps
// to exit code 1 (handled in main).
type ExitError struct {
	Code int
	Err  error
}

func (e *ExitError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("exit code %d", e.Code)
	}
	return e.Err.Error()
}

func (e *ExitError) Unwrap() error { return e.Err }

// Exit wraps err with an explicit exit code.
func Exit(code int, err error) error { return &ExitError{Code: code, Err: err} }

// Execute runs the command tree and returns a process exit code. Errors are
// rendered to the root command's stderr as a JSON error object (matching the
// legacy CLI contract) and mapped to an exit code (ExitError.Code when set,
// otherwise 1). On success it returns 0.
func Execute(ctx context.Context, root *cobra.Command, deps *Deps) int {
	if err := root.ExecuteContext(ctx); err != nil {
		stderr := root.ErrOrStderr()
		if deps != nil && deps.Stderr != nil {
			stderr = deps.Stderr
		}
		_ = (Printer{}).RenderError(stderr, err)
		var ee *ExitError
		if errors.As(err, &ee) && ee.Code != 0 {
			return ee.Code
		}
		return 1
	}
	return 0
}
