---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_DoSpolicy6"
description: |-
  Configure IPv6 DoS policies.
---

# fortios_firewall_DoSpolicy6
Configure IPv6 DoS policies.

## Argument Reference

The following arguments are supported:

* `policyid` - (Required) Policy ID.
* `status` - Enable/disable this policy.
* `name` - Policy name. (`ver >= 6.4.2`)
* `comments` - Comment.
* `interface` - (Required) Incoming interface name from available interfaces.
* `srcaddr` - (Required) Source address name from available addresses. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) Destination address name from available addresses. The structure of `dstaddr` block is documented below.
* `service` - Service object from available options. The structure of `service` block is documented below.
* `anomaly` - Anomaly name. The structure of `anomaly` block is documented below.

The `srcaddr` block supports:

* `name` - Service name.

The `dstaddr` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Service name.

The `anomaly` block supports:

* `name` - Anomaly name.
* `status` - Enable/disable this anomaly.
* `log` - Enable/disable anomaly logging.
* `action` - Action taken when the threshold is reached.
* `quarantine` - Quarantine method.
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging.
* `threshold` - Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.
* `thresholddefault` - Number of detected instances per minute which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall DosPolicy6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_DoSpolicy6.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
