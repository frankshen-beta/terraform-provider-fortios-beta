---
subcategory: "FortiGate EmailFilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_mheader"
description: |-
  Configure AntiSpam MIME header.
---

# fortios_emailfilter_mheader
Configure AntiSpam MIME header.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - Name of table.
* `comment` - Optional comments.
* `entries` - Spam filter mime header content. The structure of `entries` block is documented below.

The `entries` block supports:

* `status` - Enable/disable status.
* `id` - Mime header entry ID.
* `fieldname` - Pattern for header field name.
* `fieldbody` - Pattern for the header field body.
* `pattern_type` - Wildcard pattern or regular expression.
* `action` - Mark spam or good.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter Mheader can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_emailfilter_mheader.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
