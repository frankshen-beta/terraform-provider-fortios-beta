---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_setting"
description: |-
  Configure router settings.
---

# fortios_router_setting
Configure router settings.

## Example Usage

```hcl
resource "fortios_router_setting" "trname" {
  hostname = "s1"
}
```

## Argument Reference

The following arguments are supported:

* `show_filter` - Prefix-list as filter for showing routes.
* `hostname` - Hostname for this virtual domain router.
* `ospf_debug_lsa_flags` - ospf_debug_lsa_flags (`ver <= 6.2.0`)
* `ospf_debug_nfsm_flags` - ospf_debug_nfsm_flags (`ver <= 6.2.0`)
* `ospf_debug_packet_flags` - ospf_debug_packet_flags (`ver <= 6.2.0`)
* `ospf_debug_events_flags` - ospf_debug_events_flags (`ver <= 6.2.0`)
* `ospf_debug_route_flags` - ospf_debug_route_flags (`ver <= 6.2.0`)
* `ospf_debug_ifsm_flags` - ospf_debug_ifsm_flags (`ver <= 6.2.0`)
* `ospf_debug_nsm_flags` - ospf_debug_nsm_flags (`ver <= 6.2.0`)
* `rip_debug_flags` - rip_debug_flags (`ver <= 6.2.0`)
* `bgp_debug_flags` - bgp_debug_flags (`ver <= 6.2.0`)
* `igmp_debug_flags` - igmp_debug_flags (`ver <= 6.2.0`)
* `pimdm_debug_flags` - pimdm_debug_flags (`ver <= 6.2.0`)
* `pimsm_debug_simple_flags` - pimsm_debug_simple_flags (`ver <= 6.2.0`)
* `pimsm_debug_timer_flags` - pimsm_debug_timer_flags (`ver <= 6.2.0`)
* `pimsm_debug_joinprune_flags` - pimsm_debug_joinprune_flags (`ver <= 6.2.0`)
* `imi_debug_flags` - imi_debug_flags (`ver <= 6.2.0`)
* `isis_debug_flags` - isis_debug_flags (`ver <= 6.2.0`)
* `ospf6_debug_lsa_flags` - ospf6_debug_lsa_flags (`ver <= 6.2.0`)
* `ospf6_debug_nfsm_flags` - ospf6_debug_nfsm_flags (`ver <= 6.2.0`)
* `ospf6_debug_packet_flags` - ospf6_debug_packet_flags (`ver <= 6.2.0`)
* `ospf6_debug_events_flags` - ospf6_debug_events_flags (`ver <= 6.2.0`)
* `ospf6_debug_route_flags` - ospf6_debug_route_flags (`ver <= 6.2.0`)
* `ospf6_debug_ifsm_flags` - ospf6_debug_ifsm_flags (`ver <= 6.2.0`)
* `ospf6_debug_nsm_flags` - ospf6_debug_nsm_flags (`ver <= 6.2.0`)
* `ripng_debug_flags` - ripng_debug_flags (`ver <= 6.2.0`)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_setting.labelname RouterSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
