package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/triggermesh/tm/pkg/client"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
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
				Default:     "default",
				Description: "kubernetes namespace for tm to work in",
			},
			"registry": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "registry host address",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"tm_service":       resourceTmService(),
			"tm_buildtemplate": resourceTmBuildtemplate(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"tm_service": dataTmService(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return client.NewClient(d.Get("config_path").(string), d.Get("namespace").(string), d.Get("registry").(string))
}
