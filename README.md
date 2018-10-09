
Terraform provider for knative resources based on [triggermesh CLI](https://github.com/triggermesh/tm)


### Terraform provider usage

Plugin requires either [tm](https://github.com/triggermesh/tm/blob/master/README.md)/kubectl tools configured or be running in Kubernetes cluster environment. Golang is needed to build plugin binary and Terraform to run example manifest.

Sample `main.tf` will install buildtemplate, build and deploy nodejs application source-to-URL.   


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
