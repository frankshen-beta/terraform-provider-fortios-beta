---
subcategory: "FortiGate EmailFilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_bwl"
description: |-
  Configure anti-spam black/white list.
---

# fortios_emailfilter_bwl
Configure anti-spam black/white list. Applies to FortiOS Version `<= 6.4.2`.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - Name of table.
* `comment` - Optional comments.
* `entries` - Anti-spam black/white list entries. The structure of `entries` block is documented below.

The `entries` block supports:

* `status` - Enable/disable status.
* `id` - Entry ID.
* `type` - Entry type.
* `action` - Reject, mark as spam or good email.
* `addr_type` - IP address type.
* `ip4_subnet` - IPv4 network address/subnet mask bits.
* `ip6_subnet` - IPv6 network address/subnet mask bits.
* `pattern_type` - Wildcard pattern or regular expression.
* `email_pattern` - Email address pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter Bwl can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_emailfilter_bwl.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
