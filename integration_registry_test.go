//go:build integration

// This file implements the live, read-only integration suite invoked by
// `make test-integration` (see Makefile). It requires FLEXERA_NAM_REFRESH_TOKEN
// to be exported and exercises the CLI's process entrypoint (run, in main.go)
// against the real Flexera One NAM zone API — no mocking. Only list/get (and
// other side-effect-free) operations are exercised; nothing is created,
// updated, or deleted here since the target orgs are live, shared, non-test
// environments.
//
// Add new read-only entries by extending the per-org subtests in
// TestReadOnlyCommandsLive below.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
)

const testZone = "nam"

// defaultTestOrgIDs are used when FLEXERA_NAM_ORG_IDS is not set. 6 and 1105
// are Flexera-internal test orgs reachable with the maintainers' NAM refresh
// token; other users of this suite should export FLEXERA_NAM_ORG_IDS with
// org(s) their own credentials can access.
var defaultTestOrgIDs = []int{6, 1105}

// testOrgIDs returns FLEXERA_NAM_ORG_IDS (comma-separated) if set, otherwise
// defaultTestOrgIDs.
func testOrgIDs(t *testing.T) []int {
	t.Helper()
	raw := strings.TrimSpace(os.Getenv("FLEXERA_NAM_ORG_IDS"))
	if raw == "" {
		return defaultTestOrgIDs
	}
	var ids []int
	for _, part := range strings.Split(raw, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		id, err := strconv.Atoi(part)
		if err != nil {
			t.Fatalf("invalid FLEXERA_NAM_ORG_IDS entry %q: %v", part, err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		t.Fatal("FLEXERA_NAM_ORG_IDS is set but contains no org IDs")
	}
	return ids
}

// liveRefreshToken returns FLEXERA_NAM_REFRESH_TOKEN or skips the test.
func liveRefreshToken(t *testing.T) string {
	t.Helper()
	tok := os.Getenv("FLEXERA_NAM_REFRESH_TOKEN")
	if strings.TrimSpace(tok) == "" {
		t.Skip("FLEXERA_NAM_REFRESH_TOKEN not set; skipping live integration test")
	}
	return tok
}

// runLive invokes the CLI's process entrypoint against the real API and
// returns stdout, stderr, and the process exit code.
func runLive(t *testing.T, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()
	var outBuf, errBuf bytes.Buffer
	code := run(context.Background(), args, &outBuf, &errBuf, os.Getenv, &http.Client{})
	return outBuf.String(), errBuf.String(), code
}

// firstItem returns the first element of a CLI JSON response: the response
// may be a bare array, an envelope object with a "values" or "rules" array,
// or a single object (treated as its own "first item"). Returns nil if there
// is no item (empty array).
func firstItem(v any) any {
	switch val := v.(type) {
	case []any:
		if len(val) == 0 {
			return nil
		}
		return val[0]
	case map[string]any:
		for _, key := range []string{"values", "rules", "items"} {
			if arr, ok := val[key].([]any); ok {
				if len(arr) == 0 {
					return nil
				}
				return arr[0]
			}
		}
		return val
	default:
		return nil
	}
}

// firstID decodes a CLI JSON response and returns the named field of its
// first item as a string (handling both numeric and string IDs). ok is false
// if the response has no items or the field is absent.
func firstID(t *testing.T, raw, idField string) (id string, ok bool) {
	t.Helper()
	var v any
	if err := json.Unmarshal([]byte(raw), &v); err != nil {
		t.Fatalf("decode JSON response: %v\nraw: %s", err, raw)
	}
	item := firstItem(v)
	if item == nil {
		return "", false
	}
	m, isMap := item.(map[string]any)
	if !isMap {
		return "", false
	}
	val, present := m[idField]
	if !present {
		return "", false
	}
	switch tv := val.(type) {
	case string:
		return tv, true
	case float64:
		return strconv.FormatFloat(tv, 'f', -1, 64), true
	default:
		return fmt.Sprint(tv), true
	}
}

// checkResource runs a list command and expects success. If the response
// contains at least one item, it extracts idField from the first item, builds
// a get command via getArgsFn, and expects that to succeed too. Pass a nil
// getArgsFn to only exercise a list-only endpoint. forbiddenOK treats a 403
// response as an environment/permission limitation (skip) rather than a
// failure, for resources not enabled on every org (e.g. MSP customer index).
func checkResource(t *testing.T, base, listArgs []string, idField string, getArgsFn func(id string) []string, forbiddenOK bool) {
	t.Helper()
	out, errOut, code := runLive(t, append(append([]string{}, listArgs...), base...)...)
	if code != 0 {
		if forbiddenOK && strings.Contains(out+errOut, "status 403") {
			t.Skipf("forbidden in this org (not enabled for every org): %s", strings.TrimSpace(out+errOut))
		}
		t.Fatalf("%v failed (exit %d): %s", listArgs, code, strings.TrimSpace(out+errOut))
	}
	if getArgsFn == nil {
		return
	}
	id, ok := firstID(t, out, idField)
	if !ok {
		t.Skipf("no items returned by %v; skipping get", listArgs)
		return
	}
	if _, errOut, code := runLive(t, append(getArgsFn(id), base...)...); code != 0 {
		t.Fatalf("get for id %q failed (exit %d): %s", id, code, strings.TrimSpace(errOut))
	}
}

// TestReadOnlyCommandsLive exercises read-only (list/get) commands across a
// broad swath of the generated + curated command tree against the real NAM
// API, using orgs 6 and 1105. See Makefile's test-integration target.
func TestReadOnlyCommandsLive(t *testing.T) {
	token := liveRefreshToken(t)

	t.Run("user-orgs_list", func(t *testing.T) {
		t.Parallel()
		out, errOut, code := runLive(t, "user-orgs", "list", "--zone", testZone, "--refresh-token", token)
		if code != 0 {
			t.Fatalf("user-orgs list failed (exit %d): %s", code, strings.TrimSpace(errOut))
		}
		if _, ok := firstID(t, out, "id"); !ok {
			t.Log("user-orgs list returned no orgs for this token (unexpected, but not fatal)")
		}
	})

	for _, orgID := range testOrgIDs(t) {
		t.Run(fmt.Sprintf("org_%d", orgID), func(t *testing.T) {
			base := []string{"--zone", testZone, "--org-id", strconv.Itoa(orgID), "--refresh-token", token}

			t.Run("organization_get", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"organization", "get"}, "id", nil, false)
			})

			t.Run("budget", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"budget", "list"}, "id", func(id string) []string {
					return []string{"budget", "get", "--id", id}
				}, false)
			})

			t.Run("cloud_vendor_account_list", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"cloud-vendor-account", "list"}, "", nil, false)
			})

			t.Run("contracts", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"contracts", "list"}, "id", func(id string) []string {
					return []string{"contracts", "get", "--id", id}
				}, false)
			})

			t.Run("tag_dimension", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"tag-dimension", "list"}, "id", func(id string) []string {
					return []string{"tag-dimension", "get", "--id", id}
				}, false)
			})

			t.Run("rule_based_dimension", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"rule-based-dimension", "list"}, "id", func(id string) []string {
					return []string{"rule-based-dimension", "get", "--id", id}
				}, false)
			})

			t.Run("role", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"role", "list"}, "id", func(id string) []string {
					return []string{"role", "get", "--id", id}
				}, false)
			})

			t.Run("group", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"group", "list"}, "id", func(id string) []string {
					return []string{"group", "get", "--id", id}
				}, false)
			})

			t.Run("ip_access_control_list", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"ip-access-control", "list"}, "", nil, false)
			})

			t.Run("org_login_policy_get", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"org-login-policy", "list"}, "", nil, false)
			})

			t.Run("msp_customer", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"msp-customer", "list"}, "id", func(id string) []string {
					return []string{"msp-customer", "get", "--customer-id", id}
				}, true)
			})

			t.Run("service_account", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"service-account", "list"}, "id", func(id string) []string {
					return []string{"service-account", "get", "--id", id}
				}, false)
			})

			t.Run("user", func(t *testing.T) {
				t.Parallel()
				checkResource(t, base, []string{"user", "users"}, "id", func(id string) []string {
					return []string{"user", "get", "--id", id}
				}, false)
			})

			t.Run("group_membership_list", func(t *testing.T) {
				t.Parallel()
				out, errOut, code := runLive(t, append([]string{"group", "list"}, base...)...)
				if code != 0 {
					t.Fatalf("group list failed (exit %d): %s", code, strings.TrimSpace(errOut))
				}
				groupID, ok := firstID(t, out, "id")
				if !ok {
					t.Skip("no groups in this org; skipping group-membership list")
				}
				if _, errOut, code := runLive(t, append([]string{"group-membership", "list", "--group-id", groupID}, base...)...); code != 0 {
					t.Fatalf("group-membership list for group %q failed (exit %d): %s", groupID, code, strings.TrimSpace(errOut))
				}
			})

			t.Run("service_account_client", func(t *testing.T) {
				t.Parallel()
				out, errOut, code := runLive(t, append([]string{"service-account", "list"}, base...)...)
				if code != 0 {
					t.Fatalf("service-account list failed (exit %d): %s", code, strings.TrimSpace(errOut))
				}
				saID, ok := firstID(t, out, "id")
				if !ok {
					t.Skip("no service accounts in this org; skipping service-account-client checks")
				}
				clientsOut, errOut, code := runLive(t, append([]string{"service-account-client", "list", "--service-account-id", saID}, base...)...)
				if code != 0 {
					t.Fatalf("service-account-client list for account %q failed (exit %d): %s", saID, code, strings.TrimSpace(errOut))
				}
				clientID, ok := firstID(t, clientsOut, "clientId")
				if !ok {
					t.Skip("no service account clients; skipping service-account-client get")
				}
				getArgs := []string{"service-account-client", "get", "--service-account-id", saID, "--client-id", clientID}
				if _, errOut, code := runLive(t, append(getArgs, base...)...); code != 0 {
					t.Fatalf("service-account-client get for client %q failed (exit %d): %s", clientID, code, strings.TrimSpace(errOut))
				}
			})
		})
	}
}
