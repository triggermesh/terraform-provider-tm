
[![CircleCI](https://circleci.com/gh/triggermesh/terraform-provider-tm/tree/master.svg?style=svg)](https://circleci.com/gh/triggermesh/terraform-provider-tm/tree/master)

Terraform provider for knative resources based on [triggermesh CLI](https://github.com/triggermesh/tm)

### Terraform provider usage

Plugin requires either [tm](https://github.com/triggermesh/tm/blob/master/README.md) or [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) tools configured or be running in a Kubernetes cluster environment. Golang is needed to build the plugin binary and [Terraform](https://www.terraform.io/intro/getting-started/install.html) to run example manifest.

Sample `main.tf` will install buildtemplate, build and deploy nodejs application source-to-URL.   

### Getting Started

```
git clone https://github.com/triggermesh/terraform-provider-tm.git
cd terraform-provider-tm
go build
terraform init
terraform apply
```

After applying manifest you may check service status using tm CLI:

```
tm -n default describe service tf-nodejs
```

### Support

We would love your feedback on this Terraform plugin so don't hesitate to let us know what is wrong and how we could improve it, just file an [issue](https://github.com/triggermesh/terraform-provider-tm/issues/new)

### Code of Conduct

This plugin is by no means part of [CNCF](https://www.cncf.io/) but we abide by its [code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)
