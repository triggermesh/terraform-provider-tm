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
	"github.com/triggermesh/tm/cmd/deploy"
	"github.com/triggermesh/tm/cmd/describe"
	"github.com/triggermesh/tm/pkg/client"
)

func resourceTmService() *schema.Resource {
	return &schema.Resource{
		Create: resourceTmServiceCreate,
		Read:   resourceTmServiceRead,
		Update: resourceTmServiceCreate,
		Delete: resourceTmServiceDelete,
		Exists: resourceTmServiceExists,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Required: true,
				Type:     schema.TypeString,
			},
			"image": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"url", "source", "path"},
			},
			"source": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"url", "image", "path"},
			},
			"url": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"image", "source", "path"},
			},
			"path": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"image", "source", "url"},
			},
			"revision": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"pull_policy": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"image_tag": &schema.Schema{
				Optional: true,
				Default:  "latest",
				Type:     schema.TypeString,
			},
			"build_template": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"build_argument": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"env": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"labels": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceTmServiceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name, ok := d.Get("metadata.0.name").(string)
	if !ok {
		name = d.Get("name").(string)
	}
	var buildArgs, env, labels []string
	for _, v := range d.Get("build_argument").([]interface{}) {
		buildArgs = append(buildArgs, v.(string))
	}
	for _, v := range d.Get("env").([]interface{}) {
		env = append(env, v.(string))
	}
	for _, v := range d.Get("labels").([]interface{}) {
		labels = append(labels, v.(string))
	}
	options := deploy.Options{
		PullPolicy:     d.Get("pull_policy").(string),
		ResultImageTag: d.Get("image_tag").(string),
		Buildtemplate:  d.Get("build_template").(string),
		BuildArgs:      buildArgs,
		Env:            env,
		Labels:         labels,
	}
	source := deploy.Source{
		URL:      d.Get("source").(string),
		Revision: d.Get("revision").(string),
	}
	registry := deploy.Registry{
		URL: d.Get("image").(string),
	}
	img := deploy.Image{
		Repository: source,
		Image:      registry,
		URL:        d.Get("url").(string),
		Path:       d.Get("path").(string),
	}
	s := deploy.Service{
		From:    img,
		Options: options,
	}
	if err := s.DeployService([]string{name}, &config); err != nil {
		return err
	}
	d.SetId(config.Namespace + "/" + name)
	return nil
}

func resourceTmServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name, ok := d.Get("metadata.0.name").(string)
	if !ok {
		name = d.Get("name").(string)
	}
	var s Service
	output, err := describe.Service(name, &config)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(output, &s); err != nil {
		return err
	}
	if err = ioutil.WriteFile("/tmp/out.log", output, 0644); err != nil {
		return err
	}
	d.Set("metadata", flatMetadata(s.Metadata))
	d.Set("spec", flatServiceSpec(s.ServiceSpec))
	d.Set("status", flatStatus(s.Status))
	return nil
}

func resourceTmServiceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(client.ClientSet)
	name, ok := d.Get("metadata.0.name").(string)
	if !ok {
		name = d.Get("name").(string)
	}
	if err := delete.Service([]string{name}, &config); err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceTmServiceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	config := meta.(client.ClientSet)
	name, ok := d.Get("metadata.0.name").(string)
	if !ok {
		name = d.Get("name").(string)
	}
	if _, err := describe.Service(name, &config); err != nil {
		return false, nil
	}
	return true, nil
}
