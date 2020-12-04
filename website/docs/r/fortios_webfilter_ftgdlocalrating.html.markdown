---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ftgdlocalrating"
description: |-
  Configure local FortiGuard Web Filter local ratings.
---

# fortios_webfilter_ftgdlocalrating
Configure local FortiGuard Web Filter local ratings.

## Example Usage

```hcl
resource "fortios_webfilter_ftgdlocalrating" "trname" {
  rating = "1"
  status = "enable"
  url    = "sgala.com"
}
```

## Argument Reference

The following arguments are supported:

* `url` - (Required) URL to rate locally.
* `status` - Enable/disable local rating.
* `comment` - Comment. (`ver >= 6.6.0`)
* `rating` - (Required) Local rating.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{url}}.

## Import

Webfilter FtgdLocalRating can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webfilter_ftgdlocalrating.labelname {{url}}
$ unset "FORTIOS_IMPORT_TABLE"
```
