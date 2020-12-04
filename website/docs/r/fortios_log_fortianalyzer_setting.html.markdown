---
layout: "fortios"
page_title: "FortiOS: fortios_log_fortianalyzer_setting"
sidebar_current: "docs-fortios-resource-log-fortianalyzer-setting"
subcategory: "FortiGate OldVersion"
description: |-
  Provides a resource to configure configure logging to FortiAnalyzer log management devices.
---

# fortios_log_fortianalyzer_setting
Provides a resource to configure configure logging to FortiAnalyzer log management devices.

!> **Warning:** The resource will be deprecated and replaced by new resource `fortios_logfortianalyzer_setting`, we recommend that you use the new resource.

## Example Usage
```hcl
resource "fortios_log_fortianalyzer_setting" "test1" {
  status         = "enable"
  server         = "10.2.2.99"
  source_ip      = "10.2.2.99"
  upload_option  = "realtime"
  reliable       = "enable"
  hmac_algorithm = "sha256"
  enc_algorithm  = "high-medium"
}
```

## Argument Reference
The following arguments are supported:

* `status` - (Required) Enable/disable logging to FortiAnalyzer.
* `server` - The remote FortiAnalyzer.
* `source_ip` - Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer.
* `reliable` - Enable/disable reliable logging to FortiAnalyzer.
* `hmac_algorithm` - FortiAnalyzer IPsec tunnel HMAC algorithm.
* `enc_algorithm` - Enable/disable sending FortiAnalyzer log data with SSL encryption.

## Attributes Reference
The following attributes are exported:

* `status` - Enable/disable logging to FortiAnalyzer.
* `server` - The remote FortiAnalyzer.
* `source_ip` - Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer.
* `reliable` - Enable/disable reliable logging to FortiAnalyzer.
* `hmac_algorithm` - FortiAnalyzer IPsec tunnel HMAC algorithm.
* `enc_algorithm` - Enable/disable sending FortiAnalyzer log data with SSL encryption.
