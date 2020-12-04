---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_name"
description: |-
  Configure application signatures.
---

# fortios_application_name
Configure application signatures.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Application name.
* `fosid` - Application ID.
* `category` - (Required) Application category ID.
* `sub_category` - Application sub-category ID. (`ver <= 6.4.0`)
* `popularity` - Application popularity.
* `risk` - Application risk.
* `weight` - Application weight.
* `protocol` - Application protocol.
* `technology` - Application technology.
* `behavior` - Application behavior.
* `vendor` - Application vendor.
* `parameters` - Application parameters. The structure of `parameters` block is documented below. (`ver >= 6.4.0`)
* `parameter` - Application parameter name.
* `metadata` - Meta data. The structure of `metadata` block is documented below.

The `parameters` block supports (`ver >= 6.4.0`):

* `name` - Parameter name.

The `metadata` block supports:

* `id` - ID.
* `metaid` - Meta ID.
* `valueid` - Value ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Application Name can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_application_name.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
