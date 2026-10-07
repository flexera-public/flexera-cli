package cli

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type safeDebugTestDoer func(*http.Request) (*http.Response, error)

func (f safeDebugTestDoer) Do(req *http.Request) (*http.Response, error) { return f(req) }

func TestSafeDebugCredentialsNeverTraced(t *testing.T) {
	for _, tc := range []struct{ name, contentType, body string }{
		{"oauth", "application/x-www-form-urlencoded", "client_secret=oauth-secret&refresh_token=refresh-secret"},
		{"json", "application/json", `{"password":"json-secret","access_token":"body-token"}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "https://url-user:url-password@example.test/oauth/token?api_key=query-secret#fragment-secret", strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", tc.contentType)
			req.Header.Set("Authorization", "Bearer header-secret")
			req.Header.Set("X-Custom-Credential", "custom-secret")
			req.Header.Set("Cookie", "session=cookie-secret")
			requestBody, requestURL := req.Body, req.URL
			headers := req.Header.Clone()
			resp := &http.Response{
				StatusCode: 201, Status: "201 response-status-secret",
				Header: http.Header{"Set-Cookie": {"response-cookie-secret"}},
				Body:   io.NopCloser(strings.NewReader(`{"access_token":"response-body-secret"}`)),
			}
			responseBody := resp.Body
			responseHeaders := resp.Header.Clone()
			var trace bytes.Buffer
			base := safeDebugTestDoer(func(got *http.Request) (*http.Response, error) {
				if got != req || got.Body != requestBody || got.URL != requestURL || !reflect.DeepEqual(got.Header, headers) {
					t.Fatal("debug wrapper modified request")
				}
				body, err := io.ReadAll(got.Body)
				if err != nil || string(body) != tc.body {
					t.Fatalf("request body was consumed: %q, %v", body, err)
				}
				return resp, nil
			})
			got, err := NewSafeDebugDoer(base, &trace).Do(req)
			if err != nil || got != resp || got.Body != responseBody || !reflect.DeepEqual(got.Header, responseHeaders) {
				t.Fatalf("debug wrapper modified response: %v", err)
			}
			body, err := io.ReadAll(got.Body)
			if err != nil || string(body) != `{"access_token":"response-body-secret"}` {
				t.Fatalf("response body was consumed: %q, %v", body, err)
			}
			want := "HTTP request method=\"POST\" target=\"example.test/oauth/token\"\nHTTP response status=201\n"
			if trace.String() != want {
				t.Fatalf("trace must contain only allowed metadata: %q", trace.String())
			}
		})
	}
}

func TestSafeDebugPreservesTransportErrorAndResponse(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "https://example.test/path?token=query-secret", nil)
	sentinel := errors.New("transport error with query-secret and credential-secret")
	for _, resp := range []*http.Response{nil, {StatusCode: 503}} {
		var trace bytes.Buffer
		got, err := NewSafeDebugDoer(safeDebugTestDoer(func(*http.Request) (*http.Response, error) {
			return resp, sentinel
		}), &trace).Do(req)
		if got != resp || err != sentinel {
			t.Fatal("transport result was changed")
		}
		want := "HTTP request method=\"GET\" target=\"example.test/path\"\n"
		if resp != nil {
			want += "HTTP response status=503\n"
		}
		if trace.String() != want {
			t.Fatalf("transport error leaked: %q", trace.String())
		}
	}
}

func TestSafeDebugDoesNotCloseBodies(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "https://example.test/path", nil)
	body := &safeDebugUntouchedBody{t: t}
	req.Body = body
	resp := &http.Response{StatusCode: 200, Body: body}
	got, err := NewSafeDebugDoer(safeDebugTestDoer(func(*http.Request) (*http.Response, error) {
		return resp, nil
	}), nil).Do(req)
	if err != nil || got != resp {
		t.Fatalf("got %v, %v", got, err)
	}
}

type safeDebugUntouchedBody struct{ t *testing.T }

func (b *safeDebugUntouchedBody) Read([]byte) (int, error) {
	b.t.Fatal("debug wrapper read a body")
	return 0, io.EOF
}

func (b *safeDebugUntouchedBody) Close() error {
	b.t.Fatal("debug wrapper closed a body")
	return nil
}
