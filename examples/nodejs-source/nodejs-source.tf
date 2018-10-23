# Simple nodejs service built from source
provider "tm" {
    namespace = "default"
    registry = "knative.registry.svc.cluster.local"
}

resource "tm_service" "nodejs" {
    name = "tf-nodejs"
    depends_on = ["tm_buildtemplate.nodejs"]
    build_template = "${tm_buildtemplate.nodejs.name}"
    build_argument = ["DIRECTORY=example-module"]
    source = "https://github.com/triggermesh/nodejs-runtime.git"
}

resource "tm_buildtemplate" "nodejs" {
    name = "runtime-nodejs"
    url = "https://raw.githubusercontent.com/triggermesh/nodejs-runtime/master/knative-build-template.yaml"
}

data "tm_service" "nodejs" {
    metadata {
        name = "${tm_service.nodejs.name}"
    }
}

output "image" {
    value = "${data.tm_service.nodejs.spec.0.image}"
}
output "domain" {
  value = "${data.tm_service.nodejs.status.0.domain}"
}
