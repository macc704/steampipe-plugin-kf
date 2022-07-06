package kf

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableKFNotes(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kf_notes",
		Description: "KF notes",
		List: &plugin.ListConfig{
			Hydrate: listNotes,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "kf-id"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "kf-type"},
			{Name: "authors", Type: proto.ColumnType_STRING, Description: "kf-authors"},
			{Name: "author", Type: proto.ColumnType_STRING, Description: "kf-author"},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "kf-title"},
			{Name: "body", Type: proto.ColumnType_STRING, Description: "kf-body"},
			{Name: "data", Type: proto.ColumnType_JSON, Description: "kf-data"},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "kf-created"},
			{Name: "modified", Type: proto.ColumnType_TIMESTAMP, Description: "kf-modified"},
		},
	}
}

type Note struct {
	ID      string `json:"_id"`
	Type    string
	Title   string
	Authors []string
	Author  string
	Data    struct {
		Body string
	}
	Body string
	//Text4Search string
	Status     string
	Permission string
	Created    time.Time
	Modified   time.Time
}

func listNotes(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kfConfig := GetConfig(d.Connection)
	notes := getNotes(&kfConfig)
	for _, t := range notes {
		t.Author = t.Authors[0]
		t.Body = t.Data.Body
		d.StreamListItem(ctx, t)
	}
	return nil, nil
}

func getNotes(config *KFConfig) []Note {
	login(config)

	var body = connect(config, "api/contributions/"+config.CommunityId+"/search", "POST", `{"query": {"pagesize": 10000}}`)
	var objects []Note
	if err := json.Unmarshal(body, &objects); err != nil {
		fmt.Println(err)
	}
	return objects
}
