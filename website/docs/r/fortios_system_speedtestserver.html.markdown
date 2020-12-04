---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_speedtestserver"
description: |-
  Configure speed test server list.
---

# fortios_system_speedtestserver
Configure speed test server list.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Speed test server name.
* `timestamp` - Speed test server timestamp.
* `host` - Hosts of the server. The structure of `host` block is documented below.

The `host` block supports:

* `id` - Server host ID.
* `ip` - Server host IPv4 address.
* `port` - Server host port number to communicate with client.
* `user` - Speed test host user name.
* `password` - Speed test host password.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SpeedTestServer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_speedtestserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
