---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_quarantine"
description: |-
  Configure FortiSwitch quarantine support.
---

# fortios_switchcontroller_quarantine
Configure FortiSwitch quarantine support.

## Argument Reference

The following arguments are supported:

* `quarantine` - Enable/disable quarantine.
* `targets` - Quarantine MACs. The structure of `targets` block is documented below.

The `targets` block supports:

* `mac` - Quarantine MAC.
* `entry_id` - FSW entry id for the quarantine MAC. (`ver <= 6.2.0`)
* `description` - Description for the quarantine MAC.
* `tag` - Tags for the quarantine MAC. The structure of `tag` block is documented below.

The `tag` block supports:

* `tags` - Tag string(eg. string1 string2 string3).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController Quarantine can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_quarantine.labelname SwitchControllerQuarantine
$ unset "FORTIOS_IMPORT_TABLE"
```
