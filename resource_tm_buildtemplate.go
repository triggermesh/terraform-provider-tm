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

func resourceTmBuildtemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceTmBuildtemplateCreate,
		Read:   resourceTmBuildtemplateRead,
		Update: resourceTmBuildtemplateCreate,
		Delete: resourceTmBuildtemplateDelete,
		Exists: resourceTmBuildtemplateExists,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			"url": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"file"},
			},
			"file": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"url"},
			},
		},
	}
}

func resourceTmBuildtemplateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name := d.Get("name").(string)

	bt := deploy.Buildtemplate{
		URL:  d.Get("url").(string),
		Path: d.Get("file").(string),
	}
	if err := bt.DeployBuildTemplate([]string{name}, &config); err != nil {
		return err
	}
	d.SetId(config.Namespace + "/" + name)
	return nil
}

func resourceTmBuildtemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name := d.Get("name").(string)
	output, err := describe.BuildTemplate(name, &config)
	if err != nil {
		return err
	}
	d.Set("data", string(output))
	return nil
}

func resourceTmBuildtemplateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name := d.Get("name").(string)
	if err := delete.BuildTemplate([]string{name}, &config); err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceTmBuildtemplateExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	config := meta.(client.ClientSet)
	name := d.Get("name").(string)
	if _, err := describe.BuildTemplate(name, &config); err != nil {
		return false, nil
	}
	return true, nil
}
