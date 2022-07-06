package kf

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableKFRecords(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "kf_records",
		Description: "KF records",
		List: &plugin.ListConfig{
			Hydrate: listRecords,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "kf-id"},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "kf-type"},
			{Name: "author_id", Type: proto.ColumnType_STRING, Description: "kf-author-id"},
			{Name: "target_id", Type: proto.ColumnType_STRING, Description: "kf-target-id"},
			{Name: "data", Type: proto.ColumnType_JSON, Description: "kf-data"},
			{Name: "historical_object_id", Type: proto.ColumnType_STRING, Description: "kf"},
			{Name: "historical_object_type", Type: proto.ColumnType_STRING, Description: "kf"},
			{Name: "historical_variable_name", Type: proto.ColumnType_STRING, Description: "kf"},
			{Name: "historical_operation_type", Type: proto.ColumnType_STRING, Description: "kf"},
			{Name: "timestamp", Type: proto.ColumnType_TIMESTAMP, Description: "kf-timestamp"},
		},
	}
}

type Record struct {
	ID                      string `json:"_id"`
	Type                    string
	AuthorID                string `json:"authorId"`
	TargetID                string `json:"targetId"`
	Data                    interface{}
	HistoricalObjectID      string
	HistoricalObjectType    string
	HistoricalVariableName  string
	HistoricalOperationType string
	Timestamp               time.Time
}

func listRecords(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	kfConfig := GetConfig(d.Connection)
	records := getRecords(&kfConfig)
	for _, t := range records {
		d.StreamListItem(ctx, t)
	}
	return nil, nil
}

func getRecords(config *KFConfig) []Record {
	login(config)

	var body = connect(config, "api/records/search/"+config.CommunityId, "POST", `{}`)
	var objects []Record
	if err := json.Unmarshal(body, &objects); err != nil {
		fmt.Println(err)
	}
	return objects
}
