package kf

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableKFViews(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kf_views",
		Description: "KF views",
		List: &plugin.ListConfig{
			Hydrate: listViews,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "kf-id"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "kf-type"},
			{Name: "authors", Type: proto.ColumnType_STRING, Description: "kf-authors"},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "kf-title"},
			{Name: "data", Type: proto.ColumnType_JSON, Description: "kf-data"},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "kf-created"},
			{Name: "modified", Type: proto.ColumnType_TIMESTAMP, Description: "kf-modified"},
		},
	}
}

type View struct {
	ID      string `json:"_id"`
	Type    string
	Title   string
	Authors []string
	Author  string
	Data    struct {
		Body string
	}
	//Text4Search string
	Status     string
	Permission string
	Created    time.Time
	Modified   time.Time
}

func listViews(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kfConfig := GetConfig(d.Connection)
	views := getViews(&kfConfig)
	for _, t := range views {
		t.Author = t.Authors[0]
		d.StreamListItem(ctx, t)
	}
	return nil, nil
}

func getViews(config *KFConfig) []View {
	login(config)

	var body = connect(config, "api/communities/"+config.CommunityId+"/views", "GET", ``)
	var objects []View
	if err := json.Unmarshal(body, &objects); err != nil {
		fmt.Println(err)
	}
	return objects
}
