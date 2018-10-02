package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	tm "github.com/triggermesh/tm/cmd"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"config_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TM_CONFIG", "~/.tm/config.json"),
				Description: "Path to the tm config file, defaults to ~/.tm/config.json",
			},
			"namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TM_NAMESPACE", "default"),
				Description: "kubernetes namespace for tm to work in",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"tm_service": resourceTmService(),
			"tm_build":   nil,
		},

		DataSourcesMap: map[string]*schema.Resource{
			"tm_service": dataTmService(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	tm.CfgFile = d.Get("config_path").(string)
	tm.Namespace = d.Get("namespace").(string)
	tm.InitConfig()
	return nil, nil
}
