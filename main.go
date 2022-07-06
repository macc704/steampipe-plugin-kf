package main

import (
	"github.com/turbot/steampipe-plugin-kf/kf"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: kf.Plugin})
}
