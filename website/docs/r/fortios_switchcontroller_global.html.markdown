---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_global"
description: |-
  Configure FortiSwitch global settings.
---

# fortios_switchcontroller_global
Configure FortiSwitch global settings.

## Example Usage

```hcl
resource "fortios_switchcontroller_global" "trname" {
  allow_multiple_interfaces = "disable"
  https_image_push          = "disable"
  log_mac_limit_violations  = "disable"
  mac_aging_interval        = 332
  mac_retention_period      = 24
  mac_violation_timer       = 0
}
```

## Argument Reference

The following arguments are supported:

* `mac_aging_interval` - Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
* `allow_multiple_interfaces` - Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. (`ver <= 6.4.0`)
* `https_image_push` - Enable/disable image push to FortiSwitch using HTTPS.
* `vlan_all_mode` - VLAN configuration mode, user-defined-vlans or all-possible-vlans.
* `vlan_optimization` - FortiLink VLAN optimization.
* `disable_discovery` - Prevent this FortiSwitch from discovering. The structure of `disable_discovery` block is documented below.
* `mac_retention_period` - Time in hours after which an inactive MAC is removed from client DB.
* `default_virtual_switch_vlan` - Default VLAN for ports when added to the virtual-switch.
* `log_mac_limit_violations` - Enable/disable logs for Learning Limit Violations.
* `mac_violation_timer` - Set timeout for Learning Limit Violations (0 = disabled).
* `sn_dns_resolution` - Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number.
* `mac_event_logging` - Enable/disable MAC address event logging.
* `bounce_quarantined_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device.
* `quarantine_mode` - Quarantine mode. (`ver >= 6.4.0`)
* `update_user_device` - Control which sources update the device user list. (`ver >= 6.4.2`)
* `custom_command` - List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `custom_command` block is documented below.

The `disable_discovery` block supports:

* `name` - Managed device ID.

The `custom_command` block supports:

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Name of custom command to push to all FortiSwitches in VDOM.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController Global can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_global.labelname SwitchControllerGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
