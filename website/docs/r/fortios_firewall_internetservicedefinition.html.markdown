---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicedefinition"
description: |-
  Internet Service definition.
---

# fortios_firewall_internetservicedefinition
Internet Service definition.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) Internet Service application list ID.
* `entry` - Protocol and port information in an Internet Service entry. The structure of `entry` block is documented below.

The `entry` block supports:

* `seq_num` - Entry sequence number.
* `category_id` - Internet Service category ID.
* `name` - Internet Service name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the definition entry. The structure of `port_range` block is documented below.
* `port` - Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535). 0 means undefined. (`ver <= 6.2.0`)

The `port_range` block supports:

* `id` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (1 to 65535).
* `end_port` - Ending TCP/UDP/SCTP destination port (1 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceDefinition can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetservicedefinition.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
