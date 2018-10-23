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
    url = "https://gist.githubusercontent.com/tzununbekov/cb89f19c4339b39fe0d8b4730523cdca/raw/20bd381cb4a8a9fc53e5261c40d574ff8a78fa6e/kaniko.yaml"
}
