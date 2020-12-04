---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerautoconfig_policy"
description: |-
  Configure FortiSwitch Auto-Config QoS policy.
---

# fortios_switchcontrollerautoconfig_policy
Configure FortiSwitch Auto-Config QoS policy.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Auto-Config QoS policy name
* `qos_policy` - Auto-Config QoS policy.
* `storm_control_policy` - Auto-Config storm control policy.
* `poe_status` - Enable/disable PoE status.
* `igmp_flood_report` - Enable/disable IGMP flood report.
* `igmp_flood_traffic` - Enable/disable IGMP flood traffic.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerAutoConfig Policy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollerautoconfig_policy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
