---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_decryptedtrafficmirror"
description: |-
  Configure decrypted traffic mirror.
---

# fortios_firewall_decryptedtrafficmirror
Configure decrypted traffic mirror. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `dstmac` - Set destination MAC address for mirrored traffic.
* `traffic_type` - Types of decrypted traffic to be mirrored.
* `traffic_source` - Source of decrypted traffic to be mirrored.
* `interface` - Decrypted traffic mirror interface The structure of `interface` block is documented below.

The `interface` block supports:

* `name` - Decrypted traffic mirror interface.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall DecryptedTrafficMirror can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_decryptedtrafficmirror.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
