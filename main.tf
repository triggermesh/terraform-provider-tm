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

resource "tm_buildtemplate" "nodejs" {
    name = "runtime-nodejs"
    url = "https://raw.githubusercontent.com/triggermesh/nodejs-runtime/master/knative-build-template.yaml"
}