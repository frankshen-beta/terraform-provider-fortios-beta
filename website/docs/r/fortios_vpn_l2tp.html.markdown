---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_l2tp"
description: |-
  Configure L2TP.
---

# fortios_vpn_l2tp
Configure L2TP.

## Argument Reference

The following arguments are supported:

* `eip` - End IP.
* `sip` - Start IP.
* `status` - (Required) Enable/disable FortiGate as a L2TP gateway.
* `usrgrp` - User group.
* `enforce_ipsec` - Enable/disable IPsec enforcement.
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests. (`ver >= 6.4.2`)
* `lcp_max_echo_fails` - Maximum number of missed LCP echo messages before disconnect. (`ver >= 6.4.2`)
* `compress` - Enable/disable data compression.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn L2Tp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpn_l2tp.labelname VpnL2Tp
$ unset "FORTIOS_IMPORT_TABLE"
```
