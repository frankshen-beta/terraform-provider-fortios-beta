---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerinitialconfig_template"
description: |-
  Configure template for auto-generated VLANs.
---

# fortios_switchcontrollerinitialconfig_template
Configure template for auto-generated VLANs. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Initial config template name
* `vlanid` - Unique VLAN ID.
* `ip` - Interface IPv4 address and subnet mask.
* `allowaccess` - Permitted types of management access to this interface.
* `auto_ip` - Automatically allocate interface address and subnet block.
* `dhcp_server` - Enable/disable a DHCP server on this interface.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerInitialConfig Template can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollerinitialconfig_template.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
