---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_region"
description: |-
  Define region table.
---

# fortios_firewall_region
Define region table. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) Region ID.
* `name` - Region name.
* `city` - City ID list. The structure of `city` block is documented below.

The `city` block supports:

* `id` - City ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall Region can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_region.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
