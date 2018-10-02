package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	tm "github.com/triggermesh/tm/cmd"
)

func resourceTmService() *schema.Resource {
	return &schema.Resource{
		Create: resourceTmServiceDeploy,
		Read:   resourceTmServiceGet,
		Update: resourceTmServiceDeploy,
		Delete: resourceTmServiceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
			},
			"image": &schema.Schema{
				Type: schema.TypeString,
			},
			"url": &schema.Schema{
				Type: schema.TypeString,
			},
			"source": &schema.Schema{
				Type: schema.TypeString,
			},
			"revision": &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func resourceTmServiceDeploy(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	if err := tm.DeployService([]string{name}); err != nil {
		return err
	}
	d.SetId(name)
	return nil
}

func resourceTmServiceGet(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTmServiceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
