---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_setting"
description: |-
  VDOM wireless controller configuration.
---

# fortios_wirelesscontroller_setting
VDOM wireless controller configuration.

## Argument Reference

The following arguments are supported:

* `account_id` - FortiCloud customer account ID.
* `country` - Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.
* `duplicate_ssid` - Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM.
* `fapc_compatibility` - Enable/disable FAP-C series compatibility.
* `wfa_compatibility` - Enable/disable WFA compatibility.
* `phishing_ssid_detect` - Enable/disable phishing SSID detection.
* `fake_ssid_action` - Actions taken for detected fake SSID.
* `offending_ssid` - Configure offending SSID. The structure of `offending_ssid` block is documented below.
* `device_weight` - Upper limit of confidence of device for identification (0 - 255, default = 1, 0 = disable). (`ver >= 6.4.2`)
* `device_holdoff` - Lower limit of creation time of device for identification in minutes (0 - 60, default = 5). (`ver >= 6.4.2`)
* `device_idle` - Upper limit of idle time of device for identification in minutes (0 - 14400, default = 1440). (`ver >= 6.4.2`)
* `darrp_optimize` - Time for running Dynamic Automatic Radio Resource Provisioning (DARRP) optimizations (0 - 86400 sec, default = 86400, 0 = disable).
* `darrp_optimize_schedules` - Firewall schedules for DARRP running time. DARRP will run periodically based on darrp-optimize within the schedules. Separate multiple schedule names with a space. The structure of `darrp_optimize_schedules` block is documented below.

The `offending_ssid` block supports:

* `id` - ID.
* `ssid_pattern` - Define offending SSID pattern (case insensitive), eg: word, word*, *word, wo*rd.
* `action` - Actions taken for detected offending SSID.

The `darrp_optimize_schedules` block supports:

* `name` - Schedule name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_setting.labelname WirelessControllerSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
