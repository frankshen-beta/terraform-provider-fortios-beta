---
subcategory: "FortiGate EmailFilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_bword"
description: |-
  Configure AntiSpam banned word list.
---

# fortios_emailfilter_bword
Configure AntiSpam banned word list.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - Name of table.
* `comment` - Optional comments.
* `entries` - Spam filter banned word. The structure of `entries` block is documented below.

The `entries` block supports:

* `status` - Enable/disable status.
* `id` - Banned word entry ID.
* `pattern` - Pattern for the banned word.
* `pattern_type` - Wildcard pattern or regular expression.
* `action` - Mark spam or good.
* `where` - Component of the email to be scanned.
* `language` - Language for the banned word.
* `score` - Score value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter Bword can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_emailfilter_bword.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
