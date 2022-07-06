package kf

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableKFAuthors(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kf_authors",
		Description: "KF authors",
		List: &plugin.ListConfig{
			Hydrate: listAuthors,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "kf-id"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "kf-type"},
			//{Name: "authors", Type: proto.ColumnType_STRING, Description: "kf-authors"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "kf-name"},
			{Name: "first_name", Type: proto.ColumnType_STRING, Description: "kf-fname"},
			{Name: "last_name", Type: proto.ColumnType_STRING, Description: "kf-lname"},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "kf-created"},
			{Name: "modified", Type: proto.ColumnType_TIMESTAMP, Description: "kf-modified"},
		},
	}
}

type Author struct {
	ID      string `json:"_id"`
	Type    string
	Title   string
	Authors []string
	Data    struct {
		Body string
	}
	//Text4Search string
	FirstName  string
	LastName   string
	Name       string
	Status     string
	Permission string
	Created    time.Time
	Modified   time.Time
}

func listAuthors(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kfConfig := GetConfig(d.Connection)
	authors := getAuthors(&kfConfig)
	for _, t := range authors {
		t.Name = t.FirstName + " " + t.LastName
		d.StreamListItem(ctx, t)
	}
	return nil, nil
}

func getAuthors(config *KFConfig) []Author {
	login(config)

	var body = connect(config, "api/communities/"+config.CommunityId+"/authors", "GET", ``)
	var objects []Author
	if err := json.Unmarshal(body, &objects); err != nil {
		fmt.Println(err)
	}
	return objects
}
