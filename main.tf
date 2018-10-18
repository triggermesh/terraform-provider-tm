provider "tm" {
    namespace = "default"
    registry = "knative-local-registry:5000"
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

resource "tm_buildtemplate" "nodejs" {
    name = "runtime-nodejs"
    url = "https://raw.githubusercontent.com/triggermesh/nodejs-runtime/master/knative-build-template.yaml"
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
