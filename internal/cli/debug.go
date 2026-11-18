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
	// Keep http.Client recognizable to the SDK's secure FNMS download helper.
	if client, ok := base.(*http.Client); ok {
		clone := *client
		transport := client.Transport
		if transport == nil {
			transport = http.DefaultTransport
		}
		clone.Transport = &safeDebugTransport{logger: &safeDebugDoer{base: roundTripDoer{transport}, stderr: stderr}}
		return &clone
	}
	return &safeDebugDoer{base: base, stderr: stderr}
}

type roundTripDoer struct{ transport http.RoundTripper }

func (d roundTripDoer) Do(req *http.Request) (*http.Response, error) {
	return d.transport.RoundTrip(req)
}

type safeDebugTransport struct{ logger *safeDebugDoer }

func (t *safeDebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.logger.Do(req)
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
