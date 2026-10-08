package catalog

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type SearchResult struct {
	Command     string  `json:"command"`
	Summary     string  `json:"summary"`
	Action      string  `json:"action,omitempty"`
	Destructive bool    `json:"destructive"`
	Score       float64 `json:"score"`
	Usage       string  `json:"usage"`
	Schema      string  `json:"schema,omitempty"`
}
type SearchResults []SearchResult

func (results SearchResults) TableRows() ([]string, [][]string) {
	rows := make([][]string, 0, len(results))
	for _, result := range results {
		rows = append(rows, []string{result.Command, fmt.Sprintf("%.2f", result.Score), result.Summary})
	}
	return []string{"COMMAND", "SCORE", "SUMMARY"}, rows
}

type SearchOptions struct {
	Limit       int
	Tag, Action string
	// Service matches the top-level service command, one of its aliases, or
	// the x-flexera-service id (e.g. "bill-analysis", "ba", "bill_analysis").
	Service  string
	ReadOnly bool
}
type searchDocument struct {
	result   SearchResult
	tag      string
	services []string
	readOnly bool
	terms    map[string]float64
	length   float64
}
type SearchIndex struct {
	docs        []searchDocument
	frequencies map[string]int
	average     float64
}

func tokens(text string) []string {
	runes := []rune(text)
	var b strings.Builder
	for i, r := range runes {
		if unicode.IsUpper(r) && i > 0 && (unicode.IsLower(runes[i-1]) || unicode.IsDigit(runes[i-1]) || i+1 < len(runes) && unicode.IsLower(runes[i+1])) {
			b.WriteByte(' ')
		}
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteByte(' ')
		}
	}
	words := strings.Fields(b.String())
	out := []string{}
	for i, word := range words {
		if i+1 < len(words) && word == "cloud" && words[i+1] == "account" {
			out = append(out, "connector")
		}
		if canonical, ok := synonyms[word]; ok {
			word = canonical
		}
		for _, suffix := range []string{"ing", "ed", "es", "s"} {
			if strings.HasSuffix(word, suffix) && len(word) > len(suffix)+3 {
				word = strings.TrimSuffix(word, suffix)
				break
			}
		}
		if canonical, ok := synonyms[word]; ok {
			word = canonical
		}
		switch word {
		case "a", "an", "the", "to", "for", "of", "has", "do", "what":
			continue
		}
		out = append(out, word)
	}
	return out
}

// topLevel returns the ancestor of cmd that is a direct child of root.
func topLevel(root, cmd *cobra.Command) *cobra.Command {
	for c := cmd; c != nil; c = c.Parent() {
		if c.Parent() == root {
			return c
		}
	}
	return nil
}

func matchesAny(want string, values []string) bool {
	for _, v := range values {
		if strings.EqualFold(want, v) {
			return true
		}
	}
	return false
}

// NewSearchIndex walks only the supplied tree. It never runs commands or
// initializes configuration, authentication or an API client.
func NewSearchIndex(root *cobra.Command) (*SearchIndex, error) {
	loaded, err := Load()
	if err != nil {
		return nil, err
	}
	index := &SearchIndex{frequencies: map[string]int{}}
	var walk func(*cobra.Command) error
	walk = func(cmd *cobra.Command) error {
		if cmd.Hidden || cmd.Name() == "help" || cmd.Name() == "completion" {
			return nil
		}
		if !cmd.HasSubCommands() && (cmd.RunE != nil || cmd.Run != nil) && !(cmd.Parent() != nil && cmd.Parent().Name() == "cli" && cmd.Name() == "search") {
			path := strings.TrimPrefix(cmd.CommandPath(), root.Name()+" ")
			doc := searchDocument{result: SearchResult{Command: cmd.CommandPath(), Summary: cmd.Short, Usage: cmd.CommandPath()}, terms: map[string]float64{}, readOnly: cmd.Annotations["flexera.readOnly"] == "true"}
			add := func(text string, weight float64) {
				for _, term := range tokens(text) {
					doc.terms[term] += weight
					doc.length += weight
				}
			}
			add(path, 3)
			add(cmd.Short, 2)
			add(cmd.Long, 1)
			if top := topLevel(root, cmd); top != nil {
				doc.services = append([]string{top.Name()}, top.Aliases...)
				if top != cmd {
					add(top.Short, 1) // service title, e.g. "Identity and Access Management API"
				}
			}
			if id := cmd.Annotations["flexera.operationId"]; id != "" {
				e, found := loaded.Lookup(id)
				if !found {
					return fmt.Errorf("search: command %s has no catalog operation %s", path, id)
				}
				doc.tag = e.Tag
				if e.Service != "" {
					doc.services = append(doc.services, e.Service)
				}
				doc.result.Action = e.Action
				doc.result.Destructive = e.Destructive
				doc.result.Schema = root.Name() + " cli schema " + path
				doc.readOnly = e.Method == "GET" || e.Method == "HEAD" || e.Method == "OPTIONS"
				add(e.Tag+" "+e.Resource, 2)
				add(id, 2)
				add(e.Description, 1)
				for _, p := range e.Params {
					add(p.Flag, 1)
					if p.Required {
						doc.result.Usage += " --" + p.Flag + " " + strings.ToUpper(strings.ReplaceAll(p.Flag, "-", "_"))
					}
				}
				if len(e.RequestSchema) > 0 {
					if cmd.Flags().Lookup("body") != nil {
						doc.result.Usage += " --body BODY_JSON_OR_@FILE_OR_@-"
					} else if cmd.Flags().Lookup("file") != nil {
						doc.result.Usage += " --file REQUEST_JSON_FILE"
					}
				}
			} else {
				doc.result.Destructive = cmd.Annotations["flexera.destructive"] == "true"
				cmd.LocalNonPersistentFlags().VisitAll(func(f *pflag.Flag) {
					if len(f.Annotations[cobra.BashCompOneRequiredFlag]) > 0 {
						doc.result.Usage += " --" + f.Name + " " + strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
					}
				})
			}
			cmd.LocalNonPersistentFlags().VisitAll(func(f *pflag.Flag) { add(f.Name+" "+f.Usage, 1) })
			for term := range doc.terms {
				index.frequencies[term]++
			}
			index.average += doc.length
			index.docs = append(index.docs, doc)
		}
		for _, child := range cmd.Commands() {
			if err := walk(child); err != nil {
				return err
			}
		}
		return nil
	}
	if err := walk(root); err != nil {
		return nil, err
	}
	if len(index.docs) > 0 {
		index.average /= float64(len(index.docs))
	}
	return index, nil
}

func (index *SearchIndex) Search(query string, options SearchOptions) SearchResults {
	results := SearchResults{}
	terms := tokens(query)
	unique := map[string]bool{}
	for _, doc := range index.docs {
		if options.Service != "" && !matchesAny(options.Service, doc.services) ||
			options.Tag != "" && !strings.EqualFold(options.Tag, doc.tag) || options.Action != "" && !strings.EqualFold(options.Action, doc.result.Action) || options.ReadOnly && !doc.readOnly {
			continue
		}
		score := 0.0
		clear(unique)
		for _, term := range terms {
			if unique[term] {
				continue
			}
			unique[term] = true
			tf := doc.terms[term]
			if tf == 0 {
				continue
			}
			df := float64(index.frequencies[term])
			idf := math.Log(1 + (float64(len(index.docs))-df+0.5)/(df+0.5))
			score += idf * (tf * 2.2) / (tf + 1.2*(0.25+0.75*doc.length/index.average))
		}
		if score > 0 {
			result := doc.result
			result.Score = score
			results = append(results, result)
		}
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			return results[i].Command < results[j].Command
		}
		return results[i].Score > results[j].Score
	})
	limit := options.Limit
	if limit <= 0 {
		limit = 10
	}
	if len(results) > limit {
		results = results[:limit]
	}
	return results
}
