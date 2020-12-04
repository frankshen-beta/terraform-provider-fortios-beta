---
subcategory: "FortiGate IPS"
layout: "fortios"
page_title: "FortiOS: fortios_ips_global"
description: |-
  Configure IPS global parameter.
---

# fortios_ips_global
Configure IPS global parameter.

## Example Usage

```hcl
resource "fortios_ips_global" "trname" {
  anomaly_mode           = "continuous"
  database               = "regular"
  deep_app_insp_db_limit = 0
  deep_app_insp_timeout  = 0
  engine_count           = 0
  exclude_signatures     = "industrial"
  fail_open              = "disable"
  intelligent_mode       = "enable"
  session_limit_mode     = "heuristic"
  socket_size            = 0
  sync_session_ttl       = "enable"
  traffic_submit         = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `fail_open` - Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes.
* `database` - Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks.
* `traffic_submit` - Enable/disable submitting attack data found by this FortiGate to FortiGuard.
* `anomaly_mode` - Global blocking mode for rate-based anomalies.
* `session_limit_mode` - Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics).
* `intelligent_mode` - Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. (`ver <= 6.4.2`)
* `socket_size` - IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
* `engine_count` - Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
* `sync_session_ttl` - Enable/disable use of kernel session TTL for IPS sessions.
* `skype_client_public_ipaddr` - Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
* `deep_app_insp_timeout` - Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
* `deep_app_insp_db_limit` - Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
* `exclude_signatures` - Excluded signatures.
* `packet_log_queue_depth` - Packet/pcap log queue depth per IPS engine.
* `ngfw_max_scan_range` - NGFW policy-mode app detection threshold. (`ver >= 6.6.0`)
* `tls_active_probe` - TLS active probe configuration. The structure of `tls_active_probe` block is documented below. (`ver 6.2.6,6.6.0`)

The `tls_active_probe` block supports (`ver 6.2.6,6.6.0`):

* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.
* `vdom` - Virtual domain name for TLS active probe.
* `source_ip` - Source IP address used for TLS active probe.
* `source_ip6` - Source IPv6 address used for TLS active probe.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ips Global can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_ips_global.labelname IpsGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
