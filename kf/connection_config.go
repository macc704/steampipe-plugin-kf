package kf

import (
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/schema"
)

type KFConfigOrg struct {
	Url         *string `cty:"url"`
	Username    *string `cty:"username"`
	Password    *string `cty:"password"`
	CommunityId *string `cty:"communityId"`
}

type KFConfig struct {
	Url         string
	Username    string
	Password    string
	CommunityId string
	Token       string
}

var ConfigSchema = map[string]*schema.Attribute{
	"url": {
		Type: schema.TypeString,
	},
	"username": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"communityId": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &KFConfigOrg{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) KFConfig {
	if connection == nil || connection.Config == nil {
		return KFConfig{}
	}
	config, _ := connection.Config.(KFConfigOrg)
	var kfConfig KFConfig
	kfConfig.Url = *config.Url
	kfConfig.Username = *config.Username
	kfConfig.Password = *config.Password
	kfConfig.CommunityId = *config.CommunityId
	return kfConfig
}
