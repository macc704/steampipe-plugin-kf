package kf

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableKFLinks(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kf_links",
		Description: "KF links",
		List: &plugin.ListConfig{
			Hydrate: listLinks,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "kf-id"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "kf-type"},
			{Name: "from", Type: proto.ColumnType_STRING, Description: "kf-from"},
			{Name: "to", Type: proto.ColumnType_STRING, Description: "kf-to"},
			{Name: "data", Type: proto.ColumnType_JSON, Description: "kf-data"},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "kf-created"},
			{Name: "modified", Type: proto.ColumnType_TIMESTAMP, Description: "kf-modified"},
		},
	}
}

type Link struct {
	ID   string `json:"_id"`
	Type string
	From string
	To   string
	Data interface{}
	//Text4Search string
	Status     string
	Permission string
	Created    time.Time
	Modified   time.Time
}

func listLinks(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kfConfig := GetConfig(d.Connection)
	links := getLinks(&kfConfig)
	for _, t := range links {
		d.StreamListItem(ctx, t)
	}
	return nil, nil
}

func getLinks(config *KFConfig) []Link {
	login(config)

	var body = connect(config, "api/links/"+config.CommunityId+"/search", "POST", `{"query": {"pagesize": 10000}}`)
	var objects []Link
	if err := json.Unmarshal(body, &objects); err != nil {
		fmt.Println(err)
	}
	return objects
}
