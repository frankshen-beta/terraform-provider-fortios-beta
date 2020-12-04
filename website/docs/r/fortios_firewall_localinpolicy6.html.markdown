---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_localinpolicy6"
description: |-
  Configure user defined IPv6 local-in policies.
---

# fortios_firewall_localinpolicy6
Configure user defined IPv6 local-in policies.

## Example Usage

```hcl
resource "fortios_firewall_localinpolicy6" "trname" {
  action   = "accept"
  intf     = "port4"
  policyid = 1
  schedule = "always"
  status   = "enable"

  dstaddr {
    name = "all"
  }

  service {
    name = "ALL"
  }

  srcaddr {
    name = "all"
  }
}
```

## Argument Reference

The following arguments are supported:

* `policyid` - (Required) User defined local in policy ID.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset). (`ver >= 6.4.0`)
* `intf` - (Required) Incoming interface name from available options.
* `srcaddr` - (Required) Source address object from available options. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) Destination address object from available options. The structure of `dstaddr` block is documented below.
* `action` - Action performed on traffic matching the policy (default = deny).
* `service` - (Required) Service object from available options. Separate names with a space. The structure of `service` block is documented below.
* `schedule` - (Required) Schedule object from available options.
* `status` - Enable/disable this local-in policy.
* `comments` - Comment.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Service name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall LocalInPolicy6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_localinpolicy6.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
