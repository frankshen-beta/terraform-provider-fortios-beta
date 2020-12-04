---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_flowtracking"
description: |-
  Configure FortiSwitch flow tracking and export via ipfix/netflow.
---

# fortios_switchcontroller_flowtracking
Configure FortiSwitch flow tracking and export via ipfix/netflow.

## Argument Reference

The following arguments are supported:

* `sample_mode` - Configure sample mode for the flow tracking.
* `sample_rate` - Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
* `format` - Configure flow tracking protocol.
* `collector_ip` - Configure collector ip address.
* `collector_port` - Configure collector port number(0-65535, default=0).
* `transport` - Configure L4 transport protocol for exporting packets.
* `level` - Configure flow tracking level.
* `max_export_pkt_size` - Configure flow max export packet size (512-9216, default=512 bytes).
* `timeout_general` - Configure flow session general timeout (60-604800, default=3600 seconds).
* `timeout_icmp` - Configure flow session ICMP timeout (60-604800, default=300 seconds).
* `timeout_max` - Configure flow session max timeout (60-604800, default=604800 seconds).
* `timeout_tcp` - Configure flow session TCP timeout (60-604800, default=3600 seconds).
* `timeout_tcp_fin` - Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
* `timeout_tcp_rst` - Configure flow session TCP RST timeout (60-604800, default=120 seconds).
* `timeout_udp` - Configure flow session UDP timeout (60-604800, default=300 seconds).
* `aggregates` - Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.

The `aggregates` block supports:

* `id` - Aggregate id.
* `ip` - IP address to group all matching traffic sessions to a flow.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController FlowTracking can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_flowtracking.labelname SwitchControllerFlowTracking
$ unset "FORTIOS_IMPORT_TABLE"
```
