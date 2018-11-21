# Pipeline to build and push kaniko image from URL to local registry
provider "tm" {
    namespace = "default"
    registry = "knative.registry.svc.cluster.local"
}

resource "tm_build" "main" {
    name = "tf-test-build"
    url = "https://github.com/GoogleContainerTools/skaffold"
    depends_on = ["tm_buildtemplate.kaniko"]
    build_template = "kaniko"
    build_argument = ["DIRECTORY=examples/kaniko"]
}

resource "tm_buildtemplate" "kaniko" {
    name = "kaniko"
    url = "https://raw.githubusercontent.com/triggermesh/build-templates/master/kaniko/kaniko.yaml"
}
