---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logdisk_filter"
description: |-
  Configure filters for local disk logging. Use these filters to determine the log messages to record according to severity and type.
---

# fortios_logdisk_filter
Configure filters for local disk logging. Use these filters to determine the log messages to record according to severity and type.

## Example Usage

```hcl
resource "fortios_logdisk_filter" "trname" {
  anomaly           = "enable"
  dlp_archive       = "enable"
  dns               = "enable"
  filter_type       = "include"
  forward_traffic   = "enable"
  gtp               = "enable"
  local_traffic     = "enable"
  multicast_traffic = "enable"
  severity          = "information"
  sniffer_traffic   = "enable"
  ssh               = "enable"
  voip              = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `severity` - Log to disk every message above and including this severity level.
* `forward_traffic` - Enable/disable forward traffic logging.
* `local_traffic` - Enable/disable local in or out traffic logging.
* `multicast_traffic` - Enable/disable multicast traffic logging.
* `sniffer_traffic` - Enable/disable sniffer traffic logging.
* `anomaly` - Enable/disable anomaly logging.
* `netscan_discovery` - Enable/disable netscan discovery event logging. (`ver <= 6.2.0`)
* `netscan_vulnerability` - Enable/disable netscan vulnerability event logging. (`ver <= 6.2.0`)
* `voip` - Enable/disable VoIP logging.
* `dlp_archive` - Enable/disable DLP archive logging.
* `gtp` - Enable/disable GTP messages logging.
* `free_style` - Free Style Filters The structure of `free_style` block is documented below. (`ver >= 6.6.0`)
* `dns` - Enable/disable detailed DNS event logging. (`ver <= 6.2.0`)
* `ssh` - Enable/disable SSH logging. (`ver <= 6.2.0`)
* `event` - Enable/disable event logging. (`ver <= 6.4.0`)
* `system` - Enable/disable system activity logging. (`ver <= 6.4.0`)
* `radius` - Enable/disable RADIUS messages logging. (`ver <= 6.4.0`)
* `ipsec` - Enable/disable IPsec negotiation messages logging. (`ver <= 6.4.0`)
* `dhcp` - Enable/disable DHCP service messages logging. (`ver <= 6.4.0`)
* `ppp` - Enable/disable L2TP/PPTP/PPPoE logging. (`ver <= 6.4.0`)
* `admin` - Enable/disable admin login/logout logging. (`ver <= 6.4.0`)
* `ha` - Enable/disable HA logging. (`ver <= 6.4.0`)
* `auth` - Enable/disable firewall authentication logging. (`ver <= 6.4.0`)
* `pattern` - Enable/disable pattern update logging. (`ver <= 6.4.0`)
* `sslvpn_log_auth` - Enable/disable SSL user authentication logging. (`ver <= 6.4.0`)
* `sslvpn_log_adm` - Enable/disable SSL administrator login logging. (`ver <= 6.4.0`)
* `sslvpn_log_session` - Enable/disable SSL session logging. (`ver <= 6.4.0`)
* `vip_ssl` - Enable/disable VIP SSL logging. (`ver <= 6.4.0`)
* `ldb_monitor` - Enable/disable VIP real server health monitoring logging. (`ver <= 6.4.0`)
* `wan_opt` - Enable/disable WAN optimization event logging. (`ver <= 6.4.0`)
* `wireless_activity` - Enable/disable wireless activity event logging. (`ver <= 6.4.0`)
* `cpu_memory_usage` - Enable/disable CPU & memory usage logging every 5 minutes. (`ver <= 6.4.0`)
* `filter` - Disk log filter. (`ver <= 6.4.2`)
* `filter_type` - Include/exclude logs that match the filter. (`ver <= 6.4.2`)

The `free_style` block supports (`ver >= 6.6.0`):

* `id` - Entry ID.
* `category` - Log category.
* `filter` - Free style filter string.
* `filter_type` - Include/exclude logs that match the filter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogDisk Filter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logdisk_filter.labelname LogDiskFilter
$ unset "FORTIOS_IMPORT_TABLE"
```
