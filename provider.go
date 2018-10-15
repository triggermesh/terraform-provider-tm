/*
Copyright (c) 2018 TriggerMesh, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
