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
)

func dataTmService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKubernetesServiceRead,

		Schema: map[string]*schema.Schema{
			"metadata": {
				Type:        schema.TypeList,
				Description: "",
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creation_timestamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"spec": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"image": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"conditions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reason": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"internal_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"traffic": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"revision_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"configuration_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"percent": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"latest_ready_revision_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"latest_created_revision_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceKubernetesServiceRead(d *schema.ResourceData, meta interface{}) error {
	d.SetId(meta.(client.ClientSet).Namespace + "/" + d.Get("metadata.0.name").(string))

	return resourceTmServiceRead(d, meta)
}
