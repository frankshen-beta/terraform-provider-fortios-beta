---
subcategory: "FortiGate NSX-T"
layout: "fortios"
page_title: "FortiOS: fortios_nsxt_servicechain"
description: |-
  Configure NSX-T service chain.
---

# fortios_nsxt_servicechain
Configure NSX-T service chain. Applies to FortiOS Version `>= 6.6.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) Chain ID.
* `name` - Chain name.
* `service_index` - Configure service index. The structure of `service_index` block is documented below.

The `service_index` block supports:

* `id` - Service index.
* `reverse_index` - Reverse service index.
* `name` - Index name.
* `vd` - VDOM name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Nsxt ServiceChain can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_nsxt_servicechain.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
