---
layout: "fortios"
page_title: "FortiOS: fortios_system_admin_profiles"
sidebar_current: "docs-fortios-resource-system-admin-profiles"
subcategory: "FortiGate OldVersion"
description: |-
  Provides a resource to configure access profiles of FortiOS.
---

# fortios_system_admin_profiles
Provides a resource to configure access profiles of FortiOS.

!> **Warning:** The resource will be deprecated and replaced by new resource `fortios_system_accprofile`, we recommend that you use the new resource.

## Example Usage
```hcl
resource "fortios_system_admin_profiles" "test1" {
  name                  = "223d3"
  scope                 = "vdom"
  comments              = "test"
  secfabgrp             = "read-write"
  ftviewgrp             = "read"
  authgrp               = "none"
  sysgrp                = "read"
  netgrp                = "none"
  loggrp                = "none"
  fwgrp                 = "none"
  vpngrp                = "none"
  utmgrp                = "none"
  wanoptgrp             = "none"
  wifi                  = "none"
  admintimeout_override = "disable"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Profile name.
* `scope` - Scope of admin access.
* `secfabgrp` - Security Fabric.
* `ftviewgrp` - FortiView.
* `authgrp` - Administrator access to Users and Devices.
* `sysgrp` - System Configuration.
* `netgrp` - Network Configuration.
* `loggrp` - Administrator access to Logging and Reporting including viewing log messages.
* `fwgrp` - Administrator access to the Firewall configuration.
* `vpngrp` - Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
* `utmgrp` - Administrator access to Security Profiles.
* `wanoptgrp` - Administrator access to WAN Opt & Cache.
* `wifi` - Administrator access to the WiFi controller and Switch controller.
* `admintimeout_override` - Enable/disable overriding the global administrator idle timeout.
* `comments` - Comment.

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the access profile item.
* `name` - Profile name.
* `scope` - Scope of admin access.
* `secfabgrp` - Security Fabric.
* `ftviewgrp` - FortiView.
* `authgrp` - Administrator access to Users and Devices.
* `sysgrp` - System Configuration.
* `netgrp` - Network Configuration.
* `loggrp` - Administrator access to Logging and Reporting including viewing log messages.
* `fwgrp` - Administrator access to the Firewall configuration.
* `vpngrp` - Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
* `utmgrp` - Administrator access to Security Profiles.
* `wanoptgrp` - Administrator access to WAN Opt & Cache.
* `wifi` - Administrator access to the WiFi controller and Switch controller.
* `admintimeout_override` - Enable/disable overriding the global administrator idle timeout.
* `comments` - Comment.

