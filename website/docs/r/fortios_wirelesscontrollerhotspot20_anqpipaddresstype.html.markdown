---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpipaddresstype"
description: |-
  Configure IP address type availability.
---

# fortios_wirelesscontrollerhotspot20_anqpipaddresstype
Configure IP address type availability.

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_anqpipaddresstype" "trname" {
  ipv4_address_type = "public"
  ipv6_address_type = "not-available"
  name              = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) IP type name.
* `ipv6_address_type` - IPv6 address type.
* `ipv4_address_type` - IPv4 address type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 AnqpIpAddressType can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_anqpipaddresstype.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
