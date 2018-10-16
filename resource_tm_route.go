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
	"encoding/json"
	"io/ioutil"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/triggermesh/tm/cmd/delete"
	"github.com/triggermesh/tm/cmd/describe"
	"github.com/triggermesh/tm/cmd/set"
	"github.com/triggermesh/tm/pkg/client"
)

func resourceTmRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceTmRouteCreate,
		Read:   resourceTmRouteRead,
		Update: resourceTmRouteCreate,
		Delete: resourceTmRouteDelete,
		Exists: resourceTmRouteExists,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			"revisions": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"configs": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceTmRouteCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name, ok := d.Get("metadata.0.name").(string)
	if !ok {
		name = d.Get("name").(string)
	}

	var revisions, configs []string
	for _, v := range d.Get("revisions").([]interface{}) {
		revisions = append(revisions, v.(string))
	}
	for _, v := range d.Get("configs").([]interface{}) {
		configs = append(configs, v.(string))
	}

	r := set.Route{
		Revisions: revisions,
		Configs:   configs,
	}

	if err := r.SetPercentage([]string{name}, &config); err != nil {
		return err
	}
	d.SetId(config.Namespace + "/" + name)
	return nil
}

func resourceTmRouteRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name, ok := d.Get("metadata.0.name").(string)
	if !ok {
		name = d.Get("name").(string)
	}

	var r Route
	output, err := describe.Route(name, &config)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(output, &r); err != nil {
		return err
	}
	if err = ioutil.WriteFile("/tmp/out.log", output, 0644); err != nil {
		return err
	}
	d.Set("metadata", flatMetadata(r.Metadata))
	d.Set("status", flatStatus(r.Status))
	d.Set("spec", flatRouteSpec(r.RouteSpec))
	return nil
}

func resourceTmRouteDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name, ok := d.Get("metadata.0.name").(string)
	if !ok {
		name = d.Get("name").(string)
	}
	if err := delete.Route([]string{name}, &config); err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceTmRouteExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	config := meta.(client.ClientSet)
	name, ok := d.Get("metadata.0.name").(string)
	if !ok {
		name = d.Get("name").(string)
	}
	if _, err := describe.Route(name, &config); err != nil {
		return false, nil
	}
	return true, nil
}
