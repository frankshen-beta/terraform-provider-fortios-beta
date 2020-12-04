---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_group"
description: |-
  Configure firewall application groups.
---

# fortios_application_group
Configure firewall application groups.

## Example Usage

```hcl
resource "fortios_application_group" "trname" {
  comment = "group1 test"
  name    = "1"
  type    = "category"
  category {
    id = 2
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Application group name.
* `comment` - Comment
* `type` - Application group type.
* `application` - Application ID list. The structure of `application` block is documented below.
* `category` - Application category ID list. The structure of `category` block is documented below.
* `risk` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below. (`ver >= 6.6.0`)
* `protocols` - Application protocol filter. (`ver >= 6.6.0`)
* `vendor` - Application vendor filter. (`ver >= 6.6.0`)
* `technology` - Application technology filter. (`ver >= 6.6.0`)
* `behavior` - Application behavior filter. (`ver >= 6.6.0`)
* `popularity` - Application popularity filter (1 - 5, from least to most popular). (`ver >= 6.6.0`)

The `application` block supports:

* `id` - Application IDs.

The `category` block supports:

* `id` - Category IDs.

The `risk` block supports (`ver >= 6.6.0`):

* `level` - Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Application Group can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_application_group.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
