---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_country"
description: |-
  Define country table.
---

# fortios_firewall_country
Define country table. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) Country ID.
* `name` - Country name.
* `region` - Region ID list. The structure of `region` block is documented below.

The `region` block supports:

* `id` - Region ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall Country can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_country.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
