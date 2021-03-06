---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_authpath"
description: |-
  Configure authentication based routing.
---

# fortios_router_authpath
Configure authentication based routing.

## Example Usage

```hcl
resource "fortios_router_authpath" "trname" {
  device  = "port3"
  gateway = "1.1.1.1"
  name    = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the entry.
* `device` - (Required) Outgoing interface.
* `gateway` - Gateway IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router AuthPath can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_authpath.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
