package main

type Route struct {
	Metadata  `json:"metadata"`
	RouteSpec `json:"spec"`
	Status    `json:"status"`
}

type RouteSpec struct {
	Traffic `json:"traffic"`
}

func flatRouteSpec(r RouteSpec) []interface{} {
	traffic := make([]interface{}, len(r.Traffic), len(r.Traffic))
	for _, v := range r.Traffic {
		tmp := make(map[string]interface{})
		tmp["revision_name"] = v.RevisionName
		tmp["configuration_name"] = v.ConfigurationName
		tmp["percent"] = v.Percent
		traffic = append(traffic, tmp)
	}
	return traffic
}
