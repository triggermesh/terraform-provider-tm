# Simple nodejs service built from source
provider "tm" {
    namespace = "default"
    registry = "knative.registry.svc.cluster.local"
}

resource "tm_service" "nodeinfo" {
    name = "tf-nodeinfo"
    depends_on = ["tm_buildtemplate.kaniko"]
    build_template = "${tm_buildtemplate.kaniko.name}"
    build_argument = ["DIRECTORY=sample-functions/NodeInfo"]
    source = "https://github.com/openfaas/faas.git"
}

resource "tm_buildtemplate" "kaniko" {
    name = "kaniko"
    url = "https://raw.githubusercontent.com/triggermesh/build-templates/master/kaniko/kaniko.yaml"
}

data "tm_service" "nodeinfo" {
    depends_on = ["tm_service.nodeinfo"]
    metadata {
        name = "${tm_service.nodeinfo.name}"
    }
}

output "domain" {
  value = "${data.tm_service.nodeinfo.status.0.domain}"
}
