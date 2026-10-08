package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

// fixtureSpec builds a minimal spec with the x-flexera-services registry and
// one operation per (service, tag) pair in ops ("service/Tag/operationId").
func fixtureSpec(t *testing.T, ops ...string) []byte {
	t.Helper()
	services := map[string]any{}
	paths := map[string]any{}
	for i, op := range ops {
		parts := strings.SplitN(op, "/", 3)
		services[parts[0]] = map[string]string{"title": strings.ToUpper(parts[0][:1]) + parts[0][1:], "name": parts[0] + "-api", "version": "v1"}
		paths["/p"+strings.Repeat("x", i)] = map[string]any{"get": map[string]any{
			"operationId":       parts[2],
			"tags":              []string{parts[1]},
			"x-flexera-action":  "list",
			"x-flexera-service": parts[0],
		}}
	}
	data, err := json.Marshal(map[string]any{"x-flexera-services": services, "paths": paths})
	if err != nil {
		t.Fatal(err)
	}
	return data
}

// fixtureTags plans a finops_onboarding service with nested Bill Connect
// children plus the hoisted budget service.
func fixtureTags(t *testing.T) []genTag {
	t.Helper()
	tags, err := planTags(fixtureSpec(t,
		"finops_onboarding/Bill Connect/BC_index",
		"finops_onboarding/Bill Connect - AWS/BCA_index",
		"budget/Budget/Budget_index",
		"budget/Cloud Vendor Account/CVA_index",
	))
	if err != nil {
		t.Fatal(err)
	}
	return tags
}

func tagByName(t *testing.T, tags []genTag, name string) genTag {
	t.Helper()
	for _, tag := range tags {
		if tag.Tag == name {
			return tag
		}
	}
	t.Fatalf("tag %q not planned", name)
	return genTag{}
}

func TestPlanTagsGroupsByService(t *testing.T) {
	tags := fixtureTags(t)
	got := map[string][]string{}
	pkgs := map[string]string{}
	for _, tag := range tags {
		got[tag.Tag] = tag.finalPath()
		pkgs[tag.Tag] = tag.Pkg
	}
	want := map[string][]string{
		"Bill Connect":         {"finops-onboarding", "bill-connect"},
		"Bill Connect - AWS":   {"finops-onboarding", "bill-connect", "aws"},
		"Budget":               {"budget"},
		"Cloud Vendor Account": {"budget", "cloud-vendor-account"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("paths = %v, want %v", got, want)
	}
	if pkgs["Bill Connect - AWS"] != "finopsonboarding/billconnectaws" || pkgs["Budget"] != "budget/budget" {
		t.Fatalf("packages = %v", pkgs)
	}
	group := tagByName(t, tags, "Budget").Group
	if group == nil || group.Cmd != "budget" || group.Short != "Budget API" ||
		group.Long != "Commands for the Budget API (budget-api, v1).\n\nService id: budget" {
		t.Fatalf("service group = %+v", group)
	}
}

func TestPlanTagsSkipsCuratedOwnedTags(t *testing.T) {
	tags, err := planTags(fixtureSpec(t,
		"iam/User Memberships/Iam_User_Memberships_index",
		"iam/User Memberships/Iam_User_Memberships_indexProjects",
		"policy/Applied Policy/Policy_Applied_Policy_index",
		"policy/Incident/Policy_Incident_index",
	))
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, tag := range tags {
		names = append(names, tag.Service+"/"+tag.Tag)
	}
	if want := []string{"iam/User Memberships", "policy/Incident"}; !reflect.DeepEqual(names, want) {
		t.Fatalf("planned %v, want %v", names, want)
	}
}

func TestPlanTagsRejectsUnmappedMetadata(t *testing.T) {
	for name, data := range map[string][]byte{
		"no registry":     []byte(`{"paths":{}}`),
		"unnamed service": fixtureSpec(t, "brand_new/Thing/Thing_index"),
		"missing service": []byte(`{"x-flexera-services":{"iam":{}},"paths":{"/a":{"get":{"operationId":"A","tags":["T"],"x-flexera-action":"list"}}}}`),
		"unknown service": []byte(`{"x-flexera-services":{"iam":{}},"paths":{"/a":{"get":{"operationId":"A","tags":["T"],"x-flexera-action":"list","x-flexera-service":"nope"}}}}`),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := planTags(data); err == nil {
				t.Fatal("expected error")
			}
		})
	}
}
