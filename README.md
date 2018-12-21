[![Go Report Card](https://goreportcard.com/badge/github.com/triggermesh/terraform-provider-tm)](https://goreportcard.com/report/github.com/triggermesh/terraform-provider-tm) [![CircleCI](https://circleci.com/gh/triggermesh/terraform-provider-tm/tree/master.svg?style=shield)](https://circleci.com/gh/triggermesh/terraform-provider-tm/tree/master)

Terraform provider for knative resources based on [triggermesh CLI](https://github.com/triggermesh/tm)

### Terraform provider usage

Plugin requires either [tm](https://github.com/triggermesh/tm/blob/master/README.md) or [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) tools configured or be running in a Kubernetes cluster environment. Golang is needed to build the plugin binary and [Terraform](https://www.terraform.io/intro/getting-started/install.html) to run example manifests.

### Getting Started

```
git clone https://github.com/triggermesh/terraform-provider-tm.git
cd terraform-provider-tm
go build
terraform init
```

Repository contains `examples` directory with manifests for different services, for example you may deploy
[qrcode](https://github.com/faas-and-furious/qrcode) service by running:

```
terraform apply examples/qrcode/
```
Terraform will return domain name which will be available in a minutes after service creation:

```
curl tf-qrcode.default.example.com/function/qrcode --data "Triggermesh" > qrcode.png
```

Note: do not apply more then one example manifest at the same time; destroy old resource before creating new one. 

If you're updating already existing tm provider, please don't forget to run `go build` and `terraform init` commands after pulling latest changes

### Support

We would love your feedback on this Terraform plugin so don't hesitate to let us know what is wrong and how we could improve it, just file an [issue](https://github.com/triggermesh/terraform-provider-tm/issues/new)

### Code of Conduct

This plugin is by no means part of [CNCF](https://www.cncf.io/) but we abide by its [code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)
