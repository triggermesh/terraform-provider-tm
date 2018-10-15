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
	"time"
)

type Service struct {
	Metadata `json:"Metadata"`
	Spec     `json:"Spec"`
	Status   `json:"Status"`
}

type Metadata struct {
	Name              string    `json:"name"`
	Namespace         string    `json:"namespace"`
	UID               string    `json:"uid"`
	ResourceVersion   string    `json:"resourceVersion"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
}

type Spec struct {
	RunLatest struct {
		Configuration struct {
			RevisionTemplate struct {
				Spec struct {
					Container struct {
						Name  string `json:"name"`
						Image string `json:"image"`
						Env   []struct {
							Name  string `json:"name"`
							Value string `json:"value"`
						} `json:"env"`
						ImagePullPolicy string `json:"imagePullPolicy"`
					} `json:"container"`
				} `json:"spec"`
			} `json:"revisionTemplate"`
		} `json:"configuration"`
	} `json:"runLatest"`
}

type Status struct {
	Domain         string `json:"domain"`
	DomainInternal string `json:"domainInternal"`
	Conditions     `json:"conditions"`
	Traffic        `json:"traffic"`
}

type Conditions []struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type Traffic []struct {
	RevisionName      string `json:"revisionName"`
	ConfigurationName string `json:"configurationName"`
	Percent           int    `json:"percent"`
}

func flatMetadata(m Metadata) []interface{} {
	return []interface{}{map[string]interface{}{
		"name":               m.Name,
		"namespace":          m.Namespace,
		"uid":                m.UID,
		"resource_version":   m.ResourceVersion,
		"creation_timestamp": m.CreationTimestamp,
	}}
}

func flatStatus(s Status) []interface{} {
	return []interface{}{map[string]interface{}{
		"domain":          s.Domain,
		"internal_domain": s.DomainInternal,
		"traffic":         flatTraffic(s.Traffic),
		"conditions":      flatConditions(s.Conditions),
	}}
}

func flatSpec(s Spec) []interface{} {
	return []interface{}{map[string]interface{}{
		"image": s.RunLatest.Configuration.RevisionTemplate.Spec.Container.Image,
	}}
}

func flatTraffic(t Traffic) []interface{} {
	traffic := make([]interface{}, len(t), len(t))
	for _, v := range t {
		tmp := make(map[string]interface{})
		tmp["revision_name"] = v.RevisionName
		tmp["configuration_name"] = v.ConfigurationName
		tmp["percent"] = v.Percent
		traffic = append(traffic, tmp)
	}
	return traffic
}

func flatConditions(c Conditions) []interface{} {
	condition := make([]interface{}, len(c), len(c))
	for _, v := range c {
		tmp := make(map[string]interface{})
		tmp["type"] = v.Type
		tmp["status"] = v.Status
		tmp["reason"] = v.Reason
		tmp["message"] = v.Message
		condition = append(condition, tmp)
	}
	return condition
}
