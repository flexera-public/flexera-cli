package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

type tokenCreateDoer func(*http.Request) (*http.Response, error)

func (f tokenCreateDoer) Do(req *http.Request) (*http.Response, error) { return f(req) }

func executeTokenCreate(t *testing.T, doer tokenCreateDoer, args ...string) (string, string, error) {
	t.Helper()
	// Viper reads the real environment, independently of RootOptions.Getenv.
	t.Setenv("HOME", t.TempDir())
	for _, key := range []string{
		"CONFIG", "ZONE", "API_BASE_URL", "LOGIN_BASE_URL", "OUTPUT", "JSON_STYLE",
		"ACCESS_TOKEN", "CLIENT_ID", "CLIENT_SECRET", "REFRESH_TOKEN", "ORG_ID", "DEBUG",
	} {
		t.Setenv("FLEXERA_CLI_"+key, "")
	}
	root, deps := clipkg.NewRootCmd(clipkg.RootOptions{BaseHTTP: doer, Getenv: func(string) string { return "" }})
	root.AddCommand(NewCmd())
	var stdout, stderr bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stderr)
	root.SetArgs(append([]string{"auth", "token", "create", "--login-base-url", "http://login.invalid"}, args...))
	err := root.Execute()
	if deps.Config.AccessToken != "" || deps.Config.ClientID != "" || deps.Config.ClientSecret != "" || deps.Config.RefreshToken != "" {
		t.Fatal("test unexpectedly configured prerequisite authentication")
	}
	return stdout.String(), stderr.String(), err
}

func tokenCreateResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		Header:     http.Header{"Content-Type": {"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestDecodeTokenBody(t *testing.T) {
	for _, grant := range []string{"authorization_code", "client_credentials", "refresh_token"} {
		t.Run(grant, func(t *testing.T) {
			body, err := decodeTokenBody([]byte(`{"grant_type":"` + grant + `","client_id":"ID","client_secret":"secret","code":"code","redirect_uri":"https://example.invalid","refresh_token":"refresh"}`))
			if err != nil {
				t.Fatal(err)
			}
			if string(body.GrantType) != grant {
				t.Fatalf("grant_type = %q, want %q", body.GrantType, grant)
			}
		})
	}
	for _, tc := range []struct{ name, raw string }{
		{"missing grant", `{}`},
		{"invalid grant", `{"grant_type":"secret"}`},
		{"empty grant", `{"grant_type":""}`},
		{"multiple objects", `{"grant_type":"client_credentials"} {"client_secret":"secret"}`},
		{"trailing value", `{"grant_type":"client_credentials"} "secret"`},
		{"unknown field", `{"grant_type":"client_credentials","secret":"secret"}`},
		{"non-string", `{"grant_type":"client_credentials","client_secret":{"secret":"secret"}}`},
		{"null property", `{"grant_type":"client_credentials","client_secret":null}`},
		{"null object", `null`},
		{"array", `["secret"]`},
		{"string", `"secret"`},
		{"malformed", `{"client_secret":"secret"`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := decodeTokenBody([]byte(tc.raw))
			if err == nil {
				t.Fatal("expected invalid token body to fail")
			}
			if strings.Contains(err.Error(), "secret") {
				t.Fatalf("error leaked input: %v", err)
			}
		})
	}
}

func TestTokenCreateFormWithoutAuthentication(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		want url.Values
	}{
		{
			name: "body flags",
			args: []string{"--grant-type", "client_credentials", "--body-client-id", "ID", "--body-client-secret", "secret"},
			want: url.Values{"grant_type": {"client_credentials"}, "client_id": {"ID"}, "client_secret": {"secret"}},
		},
		{
			name: "body overrides flags",
			args: []string{"--grant-type", "client_credentials", "--body-client-id", "ignored-ID", "--body-client-secret", "ignored-secret", "--body", `{"grant_type":"refresh_token","refresh_token":"body-refresh"}`},
			want: url.Values{"grant_type": {"refresh_token"}, "refresh_token": {"body-refresh"}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			doer := tokenCreateDoer(func(req *http.Request) (*http.Response, error) {
				calls++
				if req.Method != http.MethodPost || req.URL.String() != "http://login.invalid/oidc/token" {
					t.Fatalf("unexpected request: %s %s", req.Method, req.URL)
				}
				if got := req.Header.Get("Content-Type"); got != "application/x-www-form-urlencoded" {
					t.Fatalf("Content-Type = %q", got)
				}
				if got := req.Header.Get("Authorization"); got != "" {
					t.Fatalf("unexpected Authorization header: %q", got)
				}
				if err := req.ParseForm(); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(req.PostForm, tc.want) {
					t.Fatalf("form = %v, want %v", req.PostForm, tc.want)
				}
				return tokenCreateResponse(http.StatusOK, `{"access_token":"issued-token","token_type":"Bearer","expires_in":3600}`), nil
			})
			stdout, stderr, err := executeTokenCreate(t, doer, tc.args...)
			if err != nil {
				t.Fatal(err)
			}
			if calls != 1 {
				t.Fatalf("HTTP calls = %d, want 1", calls)
			}
			var response flexera.AuthTokenResponseBody
			if err := json.Unmarshal([]byte(stdout), &response); err != nil {
				t.Fatalf("invalid response output %q: %v", stdout, err)
			}
			if response.AccessToken != "issued-token" || string(response.TokenType) != "Bearer" || response.ExpiresIn != 3600 {
				t.Fatalf("unexpected response: %+v", response)
			}
			if stderr != "" {
				t.Fatalf("unexpected stderr: %q", stderr)
			}
		})
	}
}

func TestTokenCreateDryRunRedactsWithoutHTTP(t *testing.T) {
	calls := 0
	doer := tokenCreateDoer(func(*http.Request) (*http.Response, error) {
		calls++
		return nil, errors.New("unexpected HTTP")
	})
	stdout, stderr, err := executeTokenCreate(t, doer, "--dry-run", "--body", `{"grant_type":"authorization_code","client_id":"private-id","client_secret":"private-secret","code":"private-code","redirect_uri":"https://private.invalid","refresh_token":"private-refresh"}`)
	if err != nil {
		t.Fatal(err)
	}
	if calls != 0 {
		t.Fatalf("dry-run made %d HTTP calls", calls)
	}
	var preview struct {
		DryRun     bool                    `json:"dryRun"`
		Redacted   bool                    `json:"redacted"`
		Validation struct{ Status string } `json:"validation"`
		Plan       struct {
			Method      string            `json:"method"`
			ContentType string            `json:"contentType"`
			Body        map[string]string `json:"body"`
		} `json:"plan"`
	}
	if err := json.Unmarshal([]byte(stdout), &preview); err != nil {
		t.Fatal(err)
	}
	if !preview.DryRun || !preview.Redacted || preview.Validation.Status != "unsupported" || preview.Plan.Method != "POST /oidc/token" || preview.Plan.ContentType != "application/x-www-form-urlencoded" {
		t.Fatalf("unexpected preview: %+v", preview)
	}
	want := map[string]string{"grant_type": "authorization_code", "client_id": "[REDACTED]", "client_secret": "[REDACTED]", "code": "[REDACTED]", "redirect_uri": "[REDACTED]", "refresh_token": "[REDACTED]"}
	if !reflect.DeepEqual(preview.Plan.Body, want) {
		t.Fatalf("preview body = %v, want %v", preview.Plan.Body, want)
	}
	if strings.Contains(stdout+stderr, "private") {
		t.Fatal("dry-run output leaked request values")
	}
}

func TestTokenCreateErrorsDoNotLeakSecrets(t *testing.T) {
	for _, tc := range []struct {
		name      string
		args      []string
		doer      tokenCreateDoer
		wantCalls int
		wantError string
	}{
		{"debug rejected", []string{"--debug"}, nil, 0, "--debug is not supported"},
		{"invalid body", []string{"--body", `{"grant_type":"client_credentials","unknown":"sensitive-value"}`}, nil, 0, "supported string properties"},
		{"transport", nil, func(*http.Request) (*http.Response, error) { return nil, errors.New("sensitive-value") }, 1, "sending or decoding"},
		{"malformed response", nil, func(*http.Request) (*http.Response, error) {
			return tokenCreateResponse(200, `{"access_token":"sensitive-value"`), nil
		}, 1, "sending or decoding"},
		{"service rejection", nil, func(*http.Request) (*http.Response, error) {
			return tokenCreateResponse(400, `{"message":"sensitive-value"}`), nil
		}, 1, "HTTP 400"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			doer := tokenCreateDoer(func(req *http.Request) (*http.Response, error) {
				calls++
				if tc.doer == nil {
					return nil, errors.New("unexpected HTTP: sensitive-value")
				}
				return tc.doer(req)
			})
			args := append([]string{"--grant-type", "client_credentials", "--body-client-id", "ID", "--body-client-secret", "sensitive-value"}, tc.args...)
			stdout, stderr, err := executeTokenCreate(t, doer, args...)
			if err == nil || !strings.Contains(err.Error(), tc.wantError) {
				t.Fatalf("error = %v, want containing %q", err, tc.wantError)
			}
			if calls != tc.wantCalls {
				t.Fatalf("HTTP calls = %d, want %d", calls, tc.wantCalls)
			}
			if strings.Contains(stdout+stderr+err.Error(), "sensitive-value") {
				t.Fatal("error or output leaked secret")
			}
		})
	}
}
