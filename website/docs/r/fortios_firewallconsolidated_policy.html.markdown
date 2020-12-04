---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallconsolidated_policy"
description: |-
  Configure consolidated IPv4/IPv6 policies.
---

# fortios_firewallconsolidated_policy
Configure consolidated IPv4/IPv6 policies. Applies to FortiOS Version `<= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `policyid` - (Required) Policy ID.
* `status` - Enable or disable this policy.
* `name` - Policy name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `srcintf` - Incoming (ingress) interface. The structure of `srcintf` block is documented below.
* `dstintf` - Outgoing (egress) interface. The structure of `dstintf` block is documented below.
* `srcaddr4` - Source IPv4 address name and address group names. The structure of `srcaddr4` block is documented below.
* `dstaddr4` - Destination IPv4 address name and address group names. The structure of `dstaddr4` block is documented below.
* `srcaddr6` - Source IPv6 address name and address group names. The structure of `srcaddr6` block is documented below.
* `dstaddr6` - Destination IPv6 address name and address group names. The structure of `dstaddr6` block is documented below.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. (`ver <= 6.4.0`)
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. (`ver <= 6.4.0`)
* `service_negate` - When enabled service specifies what the service must NOT be. (`ver <= 6.4.0`)
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.  (`ver <= 6.4.0`)
* `internet_service_name` - Internet Service name. The structure of `internet_service_name` block is documented below. (`ver 6.4.0`)
* `internet_service_id` - Internet Service ID. The structure of `internet_service_id` block is documented below.
* `internet_service_group` - Internet Service group name. The structure of `internet_service_group` block is documented below. (`ver <= 6.4.0`)
* `internet_service_custom` - Custom Internet Service name. The structure of `internet_service_custom` block is documented below. (`ver <= 6.4.0`)
* `internet_service_custom_group` - Custom Internet Service group name. The structure of `internet_service_custom_group` block is documented below. (`ver <= 6.4.0`)
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.  (`ver <= 6.4.0`)
* `internet_service_src_name` - Internet Service source name. The structure of `internet_service_src_name` block is documented below. (`ver 6.4.0`)
* `internet_service_src_id` - Internet Service source ID. The structure of `internet_service_src_id` block is documented below.
* `internet_service_src_group` - Internet Service source group name. The structure of `internet_service_src_group` block is documented below. (`ver <= 6.4.0`)
* `internet_service_src_custom` - Custom Internet Service source name. The structure of `internet_service_src_custom` block is documented below. (`ver <= 6.4.0`)
* `internet_service_src_custom_group` - Custom Internet Service source group name. The structure of `internet_service_src_custom_group` block is documented below. (`ver <= 6.4.0`)
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. (`ver <= 6.4.0`)
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. (`ver <= 6.4.0`)
* `action` - Policy action (allow/deny/ipsec).
* `schedule` - Schedule name.
* `service` - Service and service group names. The structure of `service` block is documented below.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode. (`ver <= 6.4.0`)
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy. (`ver <= 6.4.0`)
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. (`ver <= 6.4.0`)
* `webproxy_profile` - Webproxy profile name. (`ver <= 6.4.0`)
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `profile_group` - Name of profile group.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `emailfilter_profile` - Name of an existing email filter profile. (`ver <= 6.4.0`)
* `spamfilter_profile` - Name of an existing Spam filter profile. (`ver <= 6.2.0`)
* `dlp_sensor` - Name of an existing DLP sensor.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `voip_profile` - Name of an existing VoIP profile.
* `icap_profile` - Name of an existing ICAP profile.
* `cifs_profile` - Name of an existing CIFS profile. (`ver <= 6.4.0`)
* `waf_profile` - Name of an existing Web application firewall profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions.
* `logtraffic_start` - Record logs when a session starts.
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. (`ver <= 6.4.0`)
* `groups` - Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.
* `users` - Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. 
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `webproxy_forward_server` - Webproxy forward server name. (`ver <= 6.4.0`)
* `wanopt` - Enable/disable WAN optimization. (`ver <= 6.4.0`)
* `wanopt_detection` - WAN optimization auto-detection mode. (`ver <= 6.4.0`)
* `wanopt_passive_opt` - WAN optimization passive mode options. This option decides what IP address will be used to connect to server. (`ver <= 6.4.0`)
* `wanopt_profile` - WAN optimization profile. (`ver <= 6.4.0`)
* `wanopt_peer` - WAN optimization peer. (`ver <= 6.4.0`)
* `webcache` - Enable/disable web cache. (`ver <= 6.4.0`)
* `webcache_https` - Enable/disable web cache for HTTPS. (`ver <= 6.4.0`)
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per-IP traffic shaper.
* `nat` - Enable/disable source NAT.
* `fixedport` - Enable to prevent source NAT from changing a session's source port.
* `ippool` - Enable to use IP Pools for source NAT.
* `poolname4` - IPv4 pool names. The structure of `poolname4` block is documented below.
* `poolname6` - IPv6 pool names. The structure of `poolname6` block is documented below.
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `comments` - Comment.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
* `captive_portal_exempt` - Enable exemption of some users from the captive portal. (`ver <= 6.4.0`)
* `fsso_groups` - Names of FSSO groups. The structure of `fsso_groups` block is documented below. (`ver <= 6.4.0`)
* `application` - Application ID list. The structure of `application` block is documented below. (`ver <= 6.2.0`)
* `app_category` - Application category ID list. The structure of `app_category` block is documented below. (`ver <= 6.2.0`)
* `url_category` - URL category ID list. The structure of `url_category` block is documented below. (`ver <= 6.2.0`)
* `app_group` - Application group names. The structure of `app_group` block is documented below. (`ver <= 6.2.0`)

The `srcintf` block supports:

* `name` - Interface name.

The `dstintf` block supports:

* `name` - Interface name.

The `srcaddr4` block supports:

* `name` - Address name.

The `dstaddr4` block supports:

* `name` - Address name.

The `srcaddr6` block supports:

* `name` - Address name.

The `dstaddr6` block supports:

* `name` - Address name.

The `internet_service_name` block supports (`ver 6.4.0`):

* `name` - Internet Service name.

The `internet_service_id` block supports:

* `id` - Internet Service ID.

The `internet_service_group` block supports (`ver <= 6.4.0`):

* `name` - Internet Service group name.

The `internet_service_custom` block supports (`ver <= 6.4.0`):

* `name` - Custom Internet Service name.

The `internet_service_custom_group` block supports (`ver <= 6.4.0`):

* `name` - Custom Internet Service group name.

The `internet_service_src_name` block supports (`ver 6.4.0`):

* `name` - Internet Service name.

The `internet_service_src_id` block supports:

* `id` - Internet Service ID.

The `internet_service_src_group` block supports (`ver <= 6.4.0`):

* `name` - Internet Service group name.

The `internet_service_src_custom` block supports (`ver <= 6.4.0`):

* `name` - Custom Internet Service name.

The `internet_service_src_custom_group` block supports (`ver <= 6.4.0`):

* `name` - Custom Internet Service group name.

The `service` block supports:

* `name` - Service name.

The `groups` block supports:

* `name` - Group name.

The `users` block supports:

* `name` - User name.

The `poolname4` block supports:

* `name` - IPv4 pool name.

The `poolname6` block supports:

* `name` - IPv6 pool name.

The `fsso_groups` block supports (`ver <= 6.4.0`):

* `name` - Names of FSSO groups.

The `application` block supports (`ver <= 6.2.0`):

* `id` - Application IDs.

The `app_category` block supports (`ver <= 6.2.0`):

* `id` - Category IDs.

The `url_category` block supports (`ver <= 6.2.0`):

* `id` - URL category ID.

The `app_group` block supports (`ver <= 6.2.0`):

* `name` - Application group names.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

FirewallConsolidated Policy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallconsolidated_policy.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
