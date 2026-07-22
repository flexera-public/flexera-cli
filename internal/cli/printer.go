package cli

import (
	"io"

	flexera "github.com/flexera-public/unified-go-client"
)

// Printer renders API response values in the configured output format
// (json|table). It is a thin seam over the unified-go-client package so that
// generated leaf commands never format output themselves.
type Printer struct{}

// Render writes value to w in the given format ("json" or "table").
func (Printer) Render(w io.Writer, format string, value any) error {
	return flexera.Write(w, format, value)
}

// RenderError writes err to w as a JSON error object, matching the legacy
// CLI's stderr error contract.
func (Printer) RenderError(w io.Writer, err error) error {
	return flexera.WriteErrorJSON(w, err)
}
