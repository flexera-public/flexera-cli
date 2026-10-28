package cli

import (
	"context"
	"fmt"

	flexera "github.com/flexera-public/unified-go-client"
)

// TypedOutput reports whether output renders typed SDK responses: table
// output without --jq or --fields. Other formats keep raw JSON values so no
// response fields are lost.
func (d *Deps) TypedOutput() bool {
	return d.Config.Output == "table" && d.Printer.JQ == nil && len(d.Printer.Fields) == 0
}

// CollectPages follows values/nextPage pagination. For typed output, fetch
// must return the generated response type T, and the merged pages are
// restored to T so table output can render them. typed only supplies T (a nil
// value such as new(flexera.XResponse).JSON200 is enough); pass any(nil) when
// the response type varies by status code.
func CollectPages[T any](ctx context.Context, deps *Deps, noPaginate bool, skipToken *string, typed T, fetch flexera.PageFetcher) (any, error) {
	if !deps.TypedOutput() {
		return flexera.CollectPages(ctx, noPaginate, skipToken, fetch)
	}
	return flexera.CollectPagesAs(ctx, noPaginate, skipToken, func(ctx context.Context, token *string) (T, error) {
		page, err := fetch(ctx, token)
		if err != nil {
			return typed, err
		}
		value, ok := page.(T)
		if !ok {
			return typed, fmt.Errorf("pagination: page is %T, want %T", page, typed)
		}
		return value, nil
	})
}
