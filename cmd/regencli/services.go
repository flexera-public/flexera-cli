package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
)

// serviceCLI is the CLI-local presentation of one x-flexera-services entry.
// Command names stay downstream: unified-openapi only supplies the service id
// and documentation metadata.
type serviceCLI struct {
	Cmd     string
	Aliases []string
	// Hoist names the tag whose operations become direct children of the
	// service command (e.g. `budget list` instead of `budget budget list`).
	Hoist string
	// Rename overrides the derived command name of a tag, e.g. when it would
	// collide with a hoisted leaf.
	Rename map[string]string
}

// serviceCommands maps every generated x-flexera-service id to its top-level
// command. A service missing from this table fails regeneration so new
// upstream services get an explicit name instead of a derived one.
var serviceCommands = map[string]serviceCLI{
	"bill_analysis": {Cmd: "bill-analysis", Aliases: []string{"ba"}},
	"bill_upload":   {Cmd: "bill-upload", Hoist: "BillUpload"},
	"billing_center_service": {Cmd: "billing-center", Aliases: []string{"bc"}, Hoist: "BillingCenters",
		// The hoisted BillingCenters tag already has a per-center
		// `allocation-table` leaf.
		Rename: map[string]string{"AllocationTable": "org-allocation-table"}},
	"budget":                 {Cmd: "budget", Hoist: "Budget"},
	"cred":                   {Cmd: "credential", Aliases: []string{"cred"}, Hoist: "Credential"},
	"divnt":                  {Cmd: "data-inventory", Aliases: []string{"divnt"}},
	"finops_billing":         {Cmd: "finops-billing"},
	"finops_customizations":  {Cmd: "finops-customizations"},
	"finops_onboarding":      {Cmd: "finops-onboarding"},
	"graphql":                {Cmd: "graphql", Hoist: "graphql"},
	"grs":                    {Cmd: "grs"},
	"iam":                    {Cmd: "iam"},
	"optima_recommendations": {Cmd: "recommendations", Aliases: []string{"optima-recommendations"}, Hoist: "Recommendations"},
	"policy":                 {Cmd: "policy"},
	"risk":                   {Cmd: "risk"},
	"saas":                   {Cmd: "saas"},
	"uobs":                   {Cmd: "unified-onboarding", Aliases: []string{"uobs"}},
	"vis":                    {Cmd: "it-visibility", Aliases: []string{"vis"}},
}

// nestedTags places tags beneath a sibling tag of the same service instead of
// directly under the service command.
var nestedTags = map[string]map[string]struct{ Parent, Use string }{
	"finops_onboarding": {
		"Bill Connect - AWS":                   {"Bill Connect", "aws"},
		"Bill Connect - Azure CSP":             {"Bill Connect", "azure-csp"},
		"Bill Connect - Azure EA Management":   {"Bill Connect", "azure-ea-management"},
		"Bill Connect - Azure MCA":             {"Bill Connect", "azure-mca"},
		"Bill Connect - Common Bill Ingestion": {"Bill Connect", "common-bill-ingestion"},
		"Bill Connect - Databricks":            {"Bill Connect", "databricks"},
		"Bill Connect - GCP":                   {"Bill Connect", "gcp"},
		"Bill Connect - Snowflake":             {"Bill Connect", "snowflake"},
	},
}

// serviceGroup carries the registry metadata rendered on a service command.
type serviceGroup struct {
	ID      string
	Cmd     string
	Aliases []string
	Short   string
	Long    string
}

type specService struct {
	Title   string `json:"title"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// planTags derives one generated package per (service, tag) pair. Tags shared
// across services (e.g. IAM and GRS "Project") generate separately so leaf
// naming is scoped to a single service. Pairs whose operations are all owned
// by exact curated wrappers are skipped.
func planTags(specData []byte) ([]genTag, error) {
	var doc struct {
		Services map[string]specService                `json:"x-flexera-services"`
		Paths    map[string]map[string]json.RawMessage `json:"paths"`
	}
	if err := json.Unmarshal(specData, &doc); err != nil {
		return nil, err
	}
	if len(doc.Services) == 0 {
		return nil, fmt.Errorf("spec has no x-flexera-services registry; regenerate unified-openapi")
	}
	type pair struct{ service, tag string }
	generated := map[pair]int{}
	owned := map[pair]int{}
	for p, item := range doc.Paths {
		for method, raw := range item {
			if !isHTTPMethod(method) {
				continue
			}
			var op struct {
				ID      string   `json:"operationId"`
				Action  string   `json:"x-flexera-action"`
				Service string   `json:"x-flexera-service"`
				Tags    []string `json:"tags"`
			}
			if err := json.Unmarshal(raw, &op); err != nil {
				return nil, fmt.Errorf("%s %s: %w", method, p, err)
			}
			if op.Action == "" {
				continue
			}
			if op.Service == "" {
				return nil, fmt.Errorf("%s %s (%s) has no x-flexera-service", strings.ToUpper(method), p, op.ID)
			}
			if _, ok := doc.Services[op.Service]; !ok {
				return nil, fmt.Errorf("%s references unknown service %q", op.ID, op.Service)
			}
			for _, tag := range op.Tags {
				if _, curated := curatedOperations[op.ID]; curated {
					owned[pair{op.Service, tag}]++
				} else {
					generated[pair{op.Service, tag}]++
				}
			}
		}
	}

	groups := map[string]*serviceGroup{}
	byService := map[string][]string{}
	for key := range generated {
		byService[key.service] = append(byService[key.service], key.tag)
	}
	var skipped []string
	for key := range owned {
		if generated[key] == 0 {
			skipped = append(skipped, fmt.Sprintf("%s tag %q", key.service, key.tag))
		}
	}
	sort.Strings(skipped)
	for _, skip := range skipped {
		fmt.Fprintf(os.Stderr, "regencli: skip %s (exact curated wrappers own every operation)\n", skip)
	}
	serviceIDs := make([]string, 0, len(byService))
	for id := range byService {
		serviceIDs = append(serviceIDs, id)
	}
	sort.Strings(serviceIDs)

	var tags []genTag
	usedCmds := map[string]string{}
	usedAliases := map[string]string{}
	for _, id := range serviceIDs {
		cli, ok := serviceCommands[id]
		if !ok {
			return nil, fmt.Errorf("service %q has no CLI name; add it to serviceCommands in cmd/regencli/services.go", id)
		}
		if other, dup := usedCmds[cli.Cmd]; dup {
			return nil, fmt.Errorf("services %q and %q share command %q", other, id, cli.Cmd)
		}
		usedCmds[cli.Cmd] = id
		for _, alias := range cli.Aliases {
			if other, dup := usedAliases[alias]; dup {
				return nil, fmt.Errorf("services %q and %q share alias %q", other, id, alias)
			}
			usedAliases[alias] = id
		}
		info := doc.Services[id]
		group := &serviceGroup{ID: id, Cmd: cli.Cmd, Aliases: cli.Aliases, Short: info.Title + " API"}
		group.Long = fmt.Sprintf("Commands for the %s API", info.Title)
		if info.Name != "" {
			detail := info.Name
			if info.Version != "" {
				detail += ", " + info.Version
			}
			group.Long += " (" + detail + ")"
		}
		group.Long += ".\n\nService id: " + id
		groups[id] = group

		serviceTags := byService[id]
		sort.Strings(serviceTags)
		servicePkg := strings.ReplaceAll(cli.Cmd, "-", "")
		usedPkg := map[string]bool{}
		cmdOwner := map[string]string{}
		planned := map[string]*genTag{}
		for _, tag := range serviceTags {
			cmd := kebab(cleanIdent(tag))
			if renamed := cli.Rename[tag]; renamed != "" {
				cmd = renamed
			}
			if other, dup := cmdOwner[cmd]; dup {
				return nil, fmt.Errorf("service %q tags %q and %q both map to command %q", id, other, tag, cmd)
			}
			cmdOwner[cmd] = tag
			t := genTag{
				Tag:     tag,
				Service: id,
				Pkg:     path.Join(servicePkg, uniquePkg(cleanIdent(tag), usedPkg)),
				Cmd:     cmd,
				Group:   group,
			}
			planned[tag] = &t
		}
		if cli.Hoist != "" && planned[cli.Hoist] == nil {
			return nil, fmt.Errorf("service %q hoists unknown tag %q", id, cli.Hoist)
		}
		for tag := range cli.Rename {
			if planned[tag] == nil {
				return nil, fmt.Errorf("service %q renames unknown tag %q", id, tag)
			}
		}
		for _, tag := range serviceTags {
			t := planned[tag]
			switch nest, nested := nestedTags[id][tag]; {
			case tag == cli.Hoist:
				t.Path = []string{cli.Cmd}
			case nested:
				parent := planned[nest.Parent]
				if parent == nil || nest.Parent == cli.Hoist {
					return nil, fmt.Errorf("service %q nests %q under unavailable tag %q", id, tag, nest.Parent)
				}
				t.Path = []string{cli.Cmd, parent.Cmd, nest.Use}
			default:
				t.Path = []string{cli.Cmd, t.Cmd}
			}
			tags = append(tags, *t)
		}
	}
	return tags, nil
}

func isHTTPMethod(method string) bool {
	switch method {
	case "get", "put", "post", "delete", "patch", "head", "options", "trace":
		return true
	}
	return false
}
