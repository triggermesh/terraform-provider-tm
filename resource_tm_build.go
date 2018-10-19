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
	"github.com/triggermesh/tm/cmd/delete"
	"github.com/triggermesh/tm/cmd/deploy"
	"github.com/triggermesh/tm/cmd/describe"
	"github.com/triggermesh/tm/pkg/client"
)

func resourceTmBuild() *schema.Resource {
	return &schema.Resource{
		Create: resourceTmBuildCreate,
		Read:   resourceTmBuildRead,
		Update: resourceTmBuildCreate,
		Delete: resourceTmBuildDelete,
		Exists: resourceTmBuildExists,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			"url": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			"revision": &schema.Schema{
				Optional: true,
				Default:  "master",
				Type:     schema.TypeString,
			},
			"build_template": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			"build_argument": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceTmBuildCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name := d.Get("name").(string)

	var buildArgs []string
	for _, v := range d.Get("build_argument").([]interface{}) {
		buildArgs = append(buildArgs, v.(string))
	}
	b := deploy.Build{
		Source:        d.Get("url").(string),
		Revision:      d.Get("revision").(string),
		Buildtemplate: d.Get("build_template").(string),
		Args:          buildArgs,
	}
	if err := b.DeployBuild([]string{name}, &config); err != nil {
		return err
	}
	d.SetId(config.Namespace + "/" + name)
	return nil
}

func resourceTmBuildRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name := d.Get("name").(string)
	output, err := describe.Build(name, &config)
	if err != nil {
		return err
	}
	d.Set("data", string(output))
	return nil
}

func resourceTmBuildDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name := d.Get("name").(string)
	if err := delete.Build([]string{name}, &config); err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceTmBuildExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	config := meta.(client.ClientSet)
	name := d.Get("name").(string)
	if _, err := describe.Build(name, &config); err != nil {
		return false, nil
	}
	return true, nil
}
