---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_federatedupgrade"
description: |-
  Coordinate federated upgrades within the Security Fabric.
---

# fortios_system_federatedupgrade
Coordinate federated upgrades within the Security Fabric. Applies to FortiOS Version `>= 6.6.0`.

## Argument Reference

The following arguments are supported:

* `status` - Current status of the upgrade.
* `node_list` - Nodes which will be included in the upgrade. The structure of `node_list` block is documented below.

The `node_list` block supports:

* `serial` - Serial number of the node to include.
* `timing` - Whether the upgrade should be run immediately, or at a scheduled time.
* `time` - Scheduled time for the upgrade. Format hh:mm yyyy/mm/dd UTC.
* `setup_time` - When the upgrade was configured. Format hh:mm yyyy/mm/dd UTC.
* `upgrade_path` - Image IDs to upgrade through.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FederatedUpgrade can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_federatedupgrade.labelname SystemFederatedUpgrade
$ unset "FORTIOS_IMPORT_TABLE"
```
