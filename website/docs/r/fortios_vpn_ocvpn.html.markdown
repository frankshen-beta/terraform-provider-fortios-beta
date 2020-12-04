---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_ocvpn"
description: |-
  Configure Overlay Controller VPN settings.
---

# fortios_vpn_ocvpn
Configure Overlay Controller VPN settings.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable Overlay Controller cloud assisted VPN.
* `role` - Set device role.
* `multipath` - Enable/disable multipath redundancy. (`ver >= 6.4.0`)
* `sdwan` - Enable/disable adding OCVPN tunnels to SDWAN. (`ver >= 6.4.0`)
* `wan_interface` - FortiGate WAN interfaces to use with OCVPN. The structure of `wan_interface` block is documented below. (`ver >= 6.4.0`)
* `ip_allocation_block` - Class B subnet reserved for private IP address assignment. (`ver >= 6.4.0`)
* `poll_interval` - Overlay Controller VPN polling interval.
* `auto_discovery` - Enable/disable auto-discovery shortcuts.
* `eap` - Enable/disable EAP client authentication.
* `eap_users` - EAP authentication user group.
* `nat` - Enable/disable inter-overlay source NAT.
* `overlays` - Network overlays to register with Overlay Controller VPN service. The structure of `overlays` block is documented below.
* `forticlient_access` - Configure FortiClient settings. The structure of `forticlient_access` block is documented below. (`ver >= 6.4.0`)

The `wan_interface` block supports (`ver >= 6.4.0`):

* `name` - Interface name.

The `overlays` block supports:

* `overlay_name` - Overlay name. (`ver >= 6.4.0`)
* `inter_overlay` - Allow or deny traffic from other overlays. (`ver >= 6.4.0`)
* `id` - ID.
* `name` - Overlay name.
* `assign_ip` - Enable/disable client address assignment.
* `ipv4_start_ip` - Start of client IPv4 range.
* `ipv4_end_ip` - End of client IPv4 range.
* `subnets` - Internal subnets to register with OCVPN service. The structure of `subnets` block is documented below.

The `subnets` block supports:

* `id` - ID.
* `type` - Subnet type.
* `subnet` - IPv4 address and subnet mask.
* `interface` - LAN interface.

The `forticlient_access` block supports (`ver >= 6.4.0`):

* `status` - Enable/disable FortiClient to access OCVPN networks.
* `psksecret` - Pre-shared secret for FortiClient PSK authentication (ASCII string or hexadecimal encoded with a leading 0x).
* `auth_groups` - FortiClient user authentication groups. The structure of `auth_groups` block is documented below.

The `auth_groups` block supports:

* `name` - Group name.
* `auth_group` - Authentication user group for FortiClient access.
* `overlays` - OCVPN overlays to allow access to. The structure of `overlays` block is documented below.

The `overlays` block supports:

* `overlay_name` - Overlay name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn Ocvpn can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpn_ocvpn.labelname VpnOcvpn
$ unset "FORTIOS_IMPORT_TABLE"
```
