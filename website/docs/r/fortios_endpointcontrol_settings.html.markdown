---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_settings"
description: |-
  Configure endpoint control settings.
---

# fortios_endpointcontrol_settings
Configure endpoint control settings.

## Example Usage

```hcl
resource "fortios_endpointcontrol_settings" "trname" {
  download_location                     = "fortiguard"
  forticlient_avdb_update_interval      = 8
  forticlient_dereg_unsupported_client  = "enable"
  forticlient_ems_rest_api_call_timeout = 5000
  forticlient_keepalive_interval        = 60
  forticlient_offline_grace             = "disable"
  forticlient_offline_grace_interval    = 120
  forticlient_reg_key_enforce           = "disable"
  forticlient_reg_timeout               = 7
  forticlient_sys_update_interval       = 720
  forticlient_user_avatar               = "enable"
  forticlient_warning_interval          = 1
}
```

## Argument Reference

The following arguments are supported:

* `forticlient_reg_key_enforce` - Enable/disable requiring or enforcing FortiClient registration keys. (`ver <= 6.2.0`)
* `forticlient_reg_key` - FortiClient registration key. (`ver <= 6.2.0`)
* `forticlient_reg_timeout` - FortiClient registration license timeout (days, min = 1, max = 180, 0 means unlimited). (`ver <= 6.2.0`)
* `download_custom_link` - Customized URL for downloading FortiClient. (`ver <= 6.2.0`)
* `download_location` - FortiClient download location (FortiGuard or custom). (`ver <= 6.2.0`)
* `forticlient_offline_grace` - Enable/disable grace period for offline registered clients. (`ver <= 6.2.0`)
* `forticlient_offline_grace_interval` - Grace period for offline registered FortiClient (60 - 600 sec, default = 120). (`ver <= 6.2.0`)
* `forticlient_keepalive_interval` - Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).
* `forticlient_sys_update_interval` - Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).
* `forticlient_avdb_update_interval` - Period of time between FortiClient AntiVirus database updates (0 - 24 hours, default = 8). (`ver <= 6.2.0`)
* `forticlient_warning_interval` - Period of time between FortiClient portal warnings (0 - 24 hours, default = 1). (`ver <= 6.2.0`)
* `forticlient_user_avatar` - Enable/disable uploading FortiClient user avatars.
* `forticlient_disconnect_unsupported_client` - Enable/disable disconnecting of unsupported FortiClient endpoints.
* `forticlient_dereg_unsupported_client` - Enable/disable deregistering unsupported FortiClient endpoints. (`ver <= 6.2.0`)
* `forticlient_ems_rest_api_call_timeout` - FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000). (`ver <= 6.2.0`)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

EndpointControl Settings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_endpointcontrol_settings.labelname EndpointControlSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
