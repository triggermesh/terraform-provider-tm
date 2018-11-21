# Qrcode faas service
# https://github.com/faas-and-furious/qrcode
provider "tm" {
    namespace = "default"
    registry = "knative.registry.svc.cluster.local"
}

resource "tm_service" "qr" {
    name = "tf-qrcode"
    depends_on = ["tm_buildtemplate.kaniko"]
    build_template = "${tm_buildtemplate.kaniko.name}"
    source = "https://github.com/faas-and-furious/qrcode.git"
}

resource "tm_buildtemplate" "kaniko" {
    name = "kaniko"
    url = "https://raw.githubusercontent.com/triggermesh/build-templates/master/kaniko/kaniko.yaml"
}

data "tm_service" "qr" {
    metadata {
        name = "${tm_service.qr.name}"
    }
}

output "domain" {
  value = "${data.tm_service.qr.status.0.domain}"
}
