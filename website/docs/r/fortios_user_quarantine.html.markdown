---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_quarantine"
description: |-
  Configure quarantine support.
---

# fortios_user_quarantine
Configure quarantine support.

## Example Usage

```hcl
resource "fortios_user_quarantine" "trname" {
  quarantine = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `quarantine` - Enable/disable quarantine.
* `traffic_policy` - Traffic policy for quarantined MACs.
* `firewall_groups` - Firewall address group which includes all quarantine MAC address. (`ver >= 6.4.0`)
* `targets` - Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.

The `targets` block supports:

* `entry` - Quarantine entry name.
* `description` - Description for the quarantine entry.
* `macs` - Quarantine MACs. The structure of `macs` block is documented below.

The `macs` block supports:

* `mac` - Quarantine MAC.
* `entry_id` - FSW entry id for the quarantine MAC. (`ver <= 6.2.0`)
* `description` - Description for the quarantine MAC.
* `drop` - Enable/Disable dropping of quarantined device traffic (`ver >= 6.4.0`)
* `parent` - Parent entry name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

User Quarantine can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_quarantine.labelname UserQuarantine
$ unset "FORTIOS_IMPORT_TABLE"
```
