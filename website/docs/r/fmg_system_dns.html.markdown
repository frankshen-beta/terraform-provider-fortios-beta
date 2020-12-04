---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_system_dns"
sidebar_current: "docs-fortios-fortimanager-resource-system-dns"
subcategory: "FortiManager"
description: |-
  Provides a resource to configure system dns setting for FortiManager.
---

# fortios_fmg_system_dns
This resource supports modifying system dns setting for FortiManager.

## Example Usage
```hcl
resource "fortios_fmg_system_dns" "test1" {
  primary   = "208.91.112.52"
  secondary = "208.91.112.54"
}
```

## Argument Reference
The following arguments are supported:

* `primary` - Primary DNS IP.
* `secondary` - Secondary DNS IP.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `primary` - Primary DNS IP.
* `secondary` - Secondary DNS IP.
