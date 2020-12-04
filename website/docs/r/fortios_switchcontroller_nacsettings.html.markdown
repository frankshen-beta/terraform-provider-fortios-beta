---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_nacsettings"
description: |-
  Configure integrated NAC settings for FortiSwitch.
---

# fortios_switchcontroller_nacsettings
Configure integrated NAC settings for FortiSwitch. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - (Required) NAC settings name.
* `mode` - Set NAC mode to be used on the FortiSwitch ports.
* `inactive_timer` - Time interval after which inactive NAC devices will be expired (in minutes, 0 means no expiry).
* `onboarding_vlan` - Default NAC Onboarding VLAN when NAC devices are discovered.
* `auto_auth` - Enable/disable NAC device auto authorization when discovered and nac-policy matched.
* `bounce_nac_port` - Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device.
* `link_down_flush` - Clear NAC devices on switch ports on link down event.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController NacSettings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_nacsettings.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
