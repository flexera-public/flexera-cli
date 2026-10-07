package cli

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/flexera-public/flexera-cli/internal/catalog"
	"github.com/itchyny/gojq"
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

type UsageError struct{ Err error }

func (e *UsageError) Error() string { return e.Err.Error() }
func (e *UsageError) Unwrap() error { return e.Err }

// UnknownCommandError captures resolution tokens, never reconstructed from text.
type UnknownCommandError struct {
	Parent      string
	Tokens      []string
	Err         error
	Suggestions []string
}

func (e *UnknownCommandError) Error() string { return e.Err.Error() }
func (e *UnknownCommandError) Unwrap() error { return e.Err }

type ValidationDetail struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}
type ValidationError struct {
	Details []ValidationDetail
	Schema  string
}

func (e *ValidationError) Error() string { return "request body failed validation" }

func errorProperties(err error) map[string]any {
	value := map[string]any{"error": err.Error()}
	var unknown *UnknownCommandError
	if errors.As(err, &unknown) {
		suggestions := unknown.Suggestions
		if suggestions == nil {
			suggestions = []string{}
		}
		value["suggestions"] = suggestions
		value["parent"] = unknown.Parent
		value["tokens"] = unknown.Tokens
	}
	var validation *ValidationError
	if errors.As(err, &validation) {
		value["details"] = validation.Details
		value["schema"] = validation.Schema
	}
	var parse *gojq.ParseError
	if errors.As(err, &parse) {
		value["position"] = parse.Offset
	}
	return value
}

// InstallExecutionErrors wraps argument validation (before PersistentPreRunE)
// without adding child initialization hooks. It never runs a leaf command.
func InstallExecutionErrors(root *cobra.Command) {
	root.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error { return &UsageError{Err: err} })
	var walk func(*cobra.Command)
	walk = func(cmd *cobra.Command) {
		if cmd.Annotations == nil {
			cmd.Annotations = map[string]string{}
		}
		if cmd.Annotations["flexera.executionErrors"] != "true" {
			original := cmd.Args
			if cmd.HasSubCommands() {
				// Making groups runnable permits Cobra's pre-init argument hook
				// to report unknown children instead of silently showing help.
				if cmd.RunE == nil && cmd.Run == nil {
					cmd.RunE = func(cmd *cobra.Command, args []string) error { return cmd.Help() }
				}
				cmd.Args = func(cmd *cobra.Command, args []string) error {
					if len(args) > 0 {
						text := cobra.NoArgs(cmd, args)
						return &UnknownCommandError{Parent: cmd.CommandPath(), Tokens: append([]string(nil), args...), Err: text}
					}
					if original != nil {
						if err := original(cmd, args); err != nil {
							return &UsageError{Err: err}
						}
					}
					return ValidateOutputOptions(cmd)
				}
			} else if original != nil {
				cmd.Args = func(cmd *cobra.Command, args []string) error {
					if err := original(cmd, args); err != nil {
						return &UsageError{Err: err}
					}
					return ValidateOutputOptions(cmd)
				}
			} else {
				cmd.Args = func(cmd *cobra.Command, args []string) error { return ValidateOutputOptions(cmd) }
			}
			cmd.Annotations["flexera.executionErrors"] = "true"
		}
		for _, child := range cmd.Commands() {
			walk(child)
		}
	}
	walk(root)
}

// Execute runs the command tree and returns a process exit code. Errors are
// rendered to the root command's stderr as a JSON error object (matching the
// legacy CLI contract) and mapped to an exit code (ExitError.Code when set,
// otherwise 1). On success it returns 0.
func Execute(ctx context.Context, root *cobra.Command, deps *Deps, argv ...[]string) int {
	// Cobra creates completion commands late; initialize them before installing
	// output-contract argument guards so their bytes cannot silently ignore flags.
	root.InitDefaultCompletionCmd()
	root.InitDefaultHelpCmd()
	var helpError error
	originalHelp := root.HelpFunc()
	root.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		if err := ValidateOutputOptions(cmd); err != nil {
			helpError = err
			return
		}
		originalHelp(cmd, args)
	})
	var resolutionError error
	var failingParent *cobra.Command
	var versionError error
	if len(argv) > 0 {
		version, shaping := false, false
		for _, token := range argv[0] {
			if token == "--" {
				break
			}
			version = version || token == "--version" || token == "--version=true"
			for _, name := range []string{"--out-jq", "--out-fields", "--raw-output", "-r"} {
				shaping = shaping || token == name || strings.HasPrefix(token, name+"=")
			}
		}
		if version && shaping {
			versionError = Exit(2, fmt.Errorf("version preserves text output and does not support --out-jq, --out-fields or --raw-output"))
		}
		// Capture Cobra's resolution error from the original argv before installing
		// validation adapters. Tokens themselves come from parsed argv below.
		failingParent, _, resolutionError = root.Find(argv[0])
		root.SetArgs(argv[0])
	}
	InstallExecutionErrors(root)
	err := versionError
	if err == nil {
		err = root.ExecuteContext(ctx)
	}
	if helpError != nil {
		err = helpError
	}
	if err != nil {
		stderr := root.ErrOrStderr()
		if deps != nil && deps.Stderr != nil {
			stderr = deps.Stderr
		}
		var unknown *UnknownCommandError
		if errors.As(err, &unknown) {
			if resolutionError != nil && failingParent != nil && unknown.Parent == failingParent.CommandPath() {
				unknown.Err = resolutionError
			}
			unknown.Suggestions = []string{}
			if index, indexErr := catalog.NewSearchIndex(root); indexErr == nil {
				query := strings.Join(unknown.Tokens, " ")
				if unknown.Parent != root.CommandPath() {
					query = strings.TrimPrefix(unknown.Parent, root.Name()+" ") + " " + query
				}
				for _, result := range index.Search(query, catalog.SearchOptions{Limit: 3}) {
					unknown.Suggestions = append(unknown.Suggestions, result.Command)
				}
			}
		}
		style := JSONStyleAuto
		// Already-parsed flags are safe even if config/root initialization failed.
		if flag := root.PersistentFlags().Lookup(FlagJSONStyle); flag != nil && flag.Changed {
			if parsed, styleErr := ParseJSONStyle(flag.Value.String()); styleErr == nil {
				style = parsed
			}
		} else if deps != nil && deps.Stdout != nil {
			style = deps.Printer.Style
		}
		_ = WriteJSON(stderr, errorProperties(err), style, nil)
		var ee *ExitError
		if errors.As(err, &ee) && ee.Code != 0 {
			return ee.Code
		}
		var usage *UsageError
		var validation *ValidationError
		if unknown != nil || errors.As(err, &usage) || errors.As(err, &validation) {
			return 2
		}
		return 1
	}
	return 0
}
