package config

import "testing"

func TestResolveCommonUsesEnvFallback(t *testing.T) {
	cfg, err := ResolveCommon(CommonOptions{}, func(key string) string {
		switch key {
		case EnvZone:
			return "eu"
		case EnvOrgID:
			return "42"
		case EnvAccessToken:
			return "env-token"
		default:
			return ""
		}
	})
	if err != nil {
		t.Fatalf("ResolveCommon returned error: %v", err)
	}
	if cfg.Zone != "eu" {
		t.Fatalf("expected zone eu, got %q", cfg.Zone)
	}
	if cfg.OrgID != 42 {
		t.Fatalf("expected org ID 42, got %d", cfg.OrgID)
	}
	if cfg.AccessToken != "env-token" {
		t.Fatalf("expected env access token, got %q", cfg.AccessToken)
	}
}

func TestResolveCommonPrefersFlagsOverEnv(t *testing.T) {
	cfg, err := ResolveCommon(CommonOptions{
		Zone:        "apac",
		OrgID:       99,
		AccessToken: "flag-token",
	}, func(key string) string {
		switch key {
		case EnvZone:
			return "eu"
		case EnvOrgID:
			return "42"
		case EnvAccessToken:
			return "env-token"
		default:
			return ""
		}
	})
	if err != nil {
		t.Fatalf("ResolveCommon returned error: %v", err)
	}
	if cfg.Zone != "au" {
		t.Fatalf("expected zone au, got %q", cfg.Zone)
	}
	if cfg.OrgID != 99 {
		t.Fatalf("expected org ID 99, got %d", cfg.OrgID)
	}
	if cfg.AccessToken != "flag-token" {
		t.Fatalf("expected flag access token, got %q", cfg.AccessToken)
	}
}

func TestResolveCommonAllowsTableOutput(t *testing.T) {
	cfg, err := ResolveCommon(CommonOptions{Output: "table"}, func(string) string { return "" })
	if err != nil {
		t.Fatalf("ResolveCommon returned error: %v", err)
	}
	if cfg.Output != "table" {
		t.Fatalf("expected output table, got %q", cfg.Output)
	}
}

func TestParseZoneTest(t *testing.T) {
	for _, raw := range []string{"test", "Test", " staging ", "flexeratest"} {
		cfg, err := ResolveCommon(CommonOptions{Zone: raw, AccessToken: "x"}, func(string) string { return "" })
		if err != nil {
			t.Fatalf("zone %q: %v", raw, err)
		}
		if string(cfg.Zone) != "test" {
			t.Fatalf("zone %q resolved to %q, want \"test\"", raw, cfg.Zone)
		}
	}
}

func TestParseZoneRejectsUnknown(t *testing.T) {
	_, err := ResolveCommon(CommonOptions{Zone: "mars", AccessToken: "x"}, func(string) string { return "" })
	if err == nil {
		t.Fatal("expected error for unknown zone")
	}
}
