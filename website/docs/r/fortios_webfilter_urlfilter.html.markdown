---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_urlfilter"
description: |-
  Configure URL filter lists.
---

# fortios_webfilter_urlfilter
Configure URL filter lists.

## Example Usage

```hcl
resource "fortios_webfilter_urlfilter" "trname" {
  fosid                 = 1
  ip_addr_block         = "enable"
  name                  = "ei"
  one_arm_ips_urlfilter = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of URL filter list.
* `comment` - Optional comments.
* `one_arm_ips_urlfilter` - Enable/disable DNS resolver for one-arm IPS URL filter operation.
* `ip_addr_block` - Enable/disable blocking URLs when the hostname appears as an IP address.
* `entries` - URL filter entries. The structure of `entries` block is documented below.

The `entries` block supports:

* `id` - Id.
* `url` - URL to be filtered.
* `type` - Filter type (simple, regex, or wildcard).
* `action` - Action to take for URL filter matches.
* `antiphish_action` - Action to take for AntiPhishing matches. (`ver >= 6.4.0`)
* `status` - Enable/disable this URL filter.
* `exempt` - If action is set to exempt, select the security profile operations that exempt URLs skip. Separate multiple options with a space.
* `web_proxy_profile` - Web proxy profile.
* `referrer_host` - Referrer host name.
* `dns_address_family` - Resolve IPv4 address, IPv6 address, or both from DNS server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter Urlfilter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webfilter_urlfilter.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
