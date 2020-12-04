---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_system_ntp"
sidebar_current: "docs-fortios-fortimanager-resource-system-ntp"
subcategory: "FortiManager"
description: |-
  Provides a resource to configure system ntp setting for FortiManager.
---

# fortios_fmg_system_ntp
This resource supports modifying system ntp setting for FortiManager.

## Example Usage
```hcl
resource "fortios_fmg_system_ntp" "test1" {
  server        = "ntp1.fortinet.com"
  status        = "enable"
  sync_interval = 30
}
```

## Argument Reference
The following arguments are supported:

* `server` - IP address/hostname of NTP Server.
* `status` - Enable/disable NTP.
* `sync_interval` - NTP sync interval (minute).

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `server` - IP address/hostname of NTP Server.
* `status` - Enable/disable NTP.
* `sync_interval` - NTP sync interval.
