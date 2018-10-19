variable "registry" {
    type    = "string"
    default = "knative-local-registry:5000"
}

variable "namespace" {
    type = "string"
    default = "default"
}



provider "tm" {
    namespace = "${var.namespace}"
    registry = "${var.registry}"
}

resource "tm_service" "main" {
    name = "tf-nodejs"
    depends_on = ["tm_buildtemplate.nodejs"]
    build_template = "${tm_buildtemplate.nodejs.name}"
    build_argument = ["DIRECTORY=example-module", "SKIP_TLS_VERIFY=true"]
    source = "https://github.com/triggermesh/nodejs-runtime.git"
}

resource "tm_route" "main" {
    name = "${tm_service.main.name}"
    revisions = ["${tm_service.main.name}-00001=100"]
}
resource "tm_build" "main" {
    name = "tf-test-build"
    depends_on = ["tm_buildtemplate.kaniko"]
    url = "https://github.com/GoogleContainerTools/skaffold"
    build_template = "kaniko"
    build_argument = ["IMAGE=${var.registry}/tf-test-build:latest", "DIRECTORY=examples/kaniko", "DOCKERFILE=examples/kaniko/Dockerfile", "SKIP_TLS_VERIFY=true"]
}

resource "tm_buildtemplate" "nodejs" {
    name = "runtime-nodejs"
    url = "https://raw.githubusercontent.com/triggermesh/nodejs-runtime/master/knative-build-template.yaml"
}

resource "tm_buildtemplate" "kaniko" {
    name = "kaniko"
    url = "https://raw.githubusercontent.com/knative/build-templates/master/kaniko/kaniko.yaml"
}


data "tm_service" "main" {
    metadata {
        name = "${tm_service.main.name}"
    }
}

output "image" {
    value = "${data.tm_service.main.spec.0.image}"
}
output "domain" {
  value = "${data.tm_service.main.status.0.domain}"
}
