---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomdns"
description: |-
  Configure DNS servers for a non-management VDOM.
---

# fortios_system_vdomdns
Configure DNS servers for a non-management VDOM.

## Argument Reference

The following arguments are supported:

* `vdom_dns` - Enable/disable configuring DNS servers for the current VDOM.
* `primary` - Primary DNS server IP address for the VDOM.
* `secondary` - Secondary DNS server IP address for the VDOM.
* `dns_over_tls` - Enable/disable/enforce DNS over TLS.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `server_hostname` - DNS server host name list. The structure of `server_hostname` block is documented below.
* `ip6_primary` - Primary IPv6 DNS server IP address for the VDOM.
* `ip6_secondary` - Secondary IPv6 DNS server IP address for the VDOM.
* `source_ip` - Source IP for communications with the DNS server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.

The `server_hostname` block supports:

* `hostname` - DNS server host name list separated by space (maximum 4 domains).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VdomDns can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_vdomdns.labelname SystemVdomDns
$ unset "FORTIOS_IMPORT_TABLE"
```
