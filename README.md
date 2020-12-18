# terraform-provider-fortios-beta


## Quick Starts
Covers FortiOS 6.0 6.2.0 6.2.4 6.2.6 6.4.0 6.4.2 6.6.0.

## Build Steps

```
# go version
go version go1.14.7 linux/amd64

# cd /yourworkdir/
# export GOPATH=/yourworkdir/
# mkdir -p $GOPATH/src/github.com/terraform-providers
# cd $GOPATH/src/github.com/terraform-providers
# git clone https://github.com/frankshen-beta/terraform-provider-fortios-beta
# mv terraform-provider-fortios-beta/ terraform-provider-fortios
# cd terraform-provider-fortios
# export GO111MODULE=off
# make build

# cd /yourworkdir/bin
```
After executing 'terraform init' for a tf file with fortios provider included, then move compiled terraform-provider-fortios file to .terraform/plugins/registry.terraform.io/fortinetdev/fortios/x.x.x/linux_amd64/terraform-provider-fortios_vx.x.x to overwrite the provider binary file.

