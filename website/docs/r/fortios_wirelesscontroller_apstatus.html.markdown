---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_apstatus"
description: |-
  Configure access point status (rogue | accepted | suppressed).
---

# fortios_wirelesscontroller_apstatus
Configure access point status (rogue | accepted | suppressed).

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) AP ID.
* `bssid` - Access Point's (AP's) BSSID.
* `ssid` - Access Point's (AP's) SSID.
* `status` - Access Point's (AP's) status: rogue, accepted, or supressed.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController ApStatus can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_apstatus.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
