---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_wtpgroup"
description: |-
  Configure WTP groups.
---

# fortios_wirelesscontroller_wtpgroup
Configure WTP groups.

## Argument Reference

The following arguments are supported:

* `name` - (Required) WTP group name.
* `platform_type` - FortiAP models to define the WTP group platform type.
* `wtps` - WTP list. The structure of `wtps` block is documented below.

The `wtps` block supports:

* `wtp_id` - WTP ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController WtpGroup can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_wtpgroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
