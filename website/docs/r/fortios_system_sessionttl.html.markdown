---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sessionttl"
description: |-
  Configure global session TTL timers for this FortiGate.
---

# fortios_system_sessionttl
Configure global session TTL timers for this FortiGate.

## Example Usage

```hcl
resource "fortios_system_sessionttl" "trname" {
  default = "3600"
}
```

## Argument Reference

The following arguments are supported:

* `default` - Default timeout.
* `port` - Session TTL port. The structure of `port` block is documented below.

The `port` block supports:

* `id` - Table entry ID.
* `protocol` - Protocol (0 - 255).
* `start_port` - Start port number.
* `end_port` - End port number.
* `timeout` - Session timeout (TTL).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SessionTtl can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_sessionttl.labelname SystemSessionTtl
$ unset "FORTIOS_IMPORT_TABLE"
```
