package cli

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	flexera "github.com/flexera-public/unified-go-client"
)

// NewSafeDebugDoer logs only the request method, URL host and escaped path,
// and numeric response status. It never inspects headers or bodies, logs query
// strings or transport errors, or modifies the request or response.
func NewSafeDebugDoer(base flexera.HttpRequestDoer, stderr io.Writer) flexera.HttpRequestDoer {
	if base == nil {
		base = http.DefaultClient
	}
	if stderr == nil {
		stderr = io.Discard
	}
	return &safeDebugDoer{base: base, stderr: stderr}
}

type safeDebugDoer struct {
	base   flexera.HttpRequestDoer
	stderr io.Writer
	mu     sync.Mutex
}

func (d *safeDebugDoer) Do(req *http.Request) (*http.Response, error) {
	method, target := "", ""
	if req != nil {
		method = req.Method
		if req.URL != nil {
			target = req.URL.Host + req.URL.EscapedPath()
		}
	}
	// Quoting prevents control characters from forging additional trace lines.
	d.mu.Lock()
	_, _ = fmt.Fprintf(d.stderr, "HTTP request method=%q target=%q\n", method, target)
	d.mu.Unlock()
	resp, err := d.base.Do(req)
	if resp != nil {
		d.mu.Lock()
		_, _ = fmt.Fprintf(d.stderr, "HTTP response status=%d\n", resp.StatusCode)
		d.mu.Unlock()
	}
	return resp, err
}
