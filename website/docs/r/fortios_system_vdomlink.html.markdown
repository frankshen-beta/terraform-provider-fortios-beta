---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomlink"
description: |-
  Configure VDOM links.
---

# fortios_system_vdomlink
Configure VDOM links.

## Argument Reference

The following arguments are supported:

* `name` - (Required) VDOM link name (maximum = 8 characters).
* `vcluster` - Virtual cluster.
* `type` - VDOM link type: PPP or Ethernet.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VdomLink can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_vdomlink.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
