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
	d.SetId(name)
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
