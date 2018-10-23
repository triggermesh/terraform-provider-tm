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
    url = "https://gist.githubusercontent.com/tzununbekov/cb89f19c4339b39fe0d8b4730523cdca/raw/20bd381cb4a8a9fc53e5261c40d574ff8a78fa6e/kaniko.yaml"
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
