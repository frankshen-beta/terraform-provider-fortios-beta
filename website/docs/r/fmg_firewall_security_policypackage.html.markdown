---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_firewall_security_policypackage"
sidebar_current: "docs-fortios-fortimanager-resource-firewall-security-policypackage"
subcategory: "FortiManager"
description: |-
  Provides a resource to configure firewall security policypackage on FortiManager which could be installed to the FortiGate later
---

# fortios_fmg_firewall_security_policypackage
This resource supports Create/Read/Update/Delete firewall security policypackage on FortiManager which could be installed to the FortiGate later


## Example Usage
```hcl
resource "fortios_fmg_firewall_security_policypackage" "test1" {
  name   = "test-pkg"
  target = "FGVM64-test"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Security policy package name.
* `target` - The installation target.
* `adom` - Source ADOM name. default is 'root'
* `vdom` - Vdom of managed device. default is 'root'
* `inspection_mode` - Inspection Mode. Enum:[flow, proxy]. default is 'flow'

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Security policy package name.
* `target` - The installation target.
* `adom` - Source ADOM name.
* `vdom` - Vdom of managed device.
* `inspection_mode` - Inspection Mode.
