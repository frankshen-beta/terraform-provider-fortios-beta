---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_dnsbl"
description: |-
  Configure AntiSpam DNSBL/ORBL.
---

# fortios_spamfilter_dnsbl
Configure AntiSpam DNSBL/ORBL. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_spamfilter_dnsbl" "trname" {
  comment = "test"
  fosid   = 1
  name    = "s1"

  entries {
    action = "reject"
    server = "1.1.1.1"
    status = "enable"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table.
* `comment` - Optional comments.
* `entries` - Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.

The `entries` block supports:

* `status` - Enable/disable status.
* `id` - DNSBL/ORBL entry ID.
* `server` - DNSBL or ORBL server name.
* `action` - Reject connection or mark as spam email.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Spamfilter Dnsbl can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_spamfilter_dnsbl.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
