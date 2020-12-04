---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_proberesponse"
description: |-
  Configure system probe response.
---

# fortios_system_proberesponse
Configure system probe response.

## Example Usage

```hcl
resource "fortios_system_proberesponse" "trname" {
  http_probe_value = "OK"
  mode             = "none"
  port             = 8008
  security_mode    = "none"
  timeout          = 300
  ttl_mode         = "retain"
}
```

## Argument Reference

The following arguments are supported:

* `port` - Port number to response.
* `http_probe_value` - Value to respond to the monitoring server.
* `ttl_mode` - Mode for TWAMP packet TTL modification.
* `mode` - SLA response mode.
* `security_mode` - Twamp respondor security mode.
* `password` - Twamp respondor password in authentication mode
* `timeout` - An inactivity timer for a twamp test session.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System ProbeResponse can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_proberesponse.labelname SystemProbeResponse
$ unset "FORTIOS_IMPORT_TABLE"
```
