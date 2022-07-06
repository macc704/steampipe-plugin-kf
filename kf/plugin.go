package kf

import (
	"context"
	//"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-kf",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		// DefaultTransform: transform.FromJSONTag().NullIfZero(),
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"kf_notes":   tableKFNotes(ctx),
			"kf_authors": tableKFAuthors(ctx),
			"kf_views":   tableKFViews(ctx),
			"kf_links":   tableKFLinks(ctx),
		},
	}
	return p
}
