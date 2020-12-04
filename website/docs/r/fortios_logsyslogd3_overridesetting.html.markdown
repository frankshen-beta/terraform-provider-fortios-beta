---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logsyslogd3_overridesetting"
description: |-
  Override settings for remote syslog server.
---

# fortios_logsyslogd3_overridesetting
Override settings for remote syslog server.

## Argument Reference

The following arguments are supported:

* `override` - Enable/disable override syslog settings. (`ver <= 6.2.0`)
* `status` - Enable/disable remote syslog logging.
* `server` - Address of remote syslog server.
* `mode` - Remote syslog logging over UDP/Reliable TCP.
* `port` - Server listen port.
* `facility` - Remote syslog facility.
* `source_ip` - Source IP address of syslog.
* `format` - Log format.
* `priority` - Set log transmission priority.
* `max_log_rate` - Syslog maximum log rate in MBps (0 = unlimited).
* `enc_algorithm` - Enable/disable reliable syslogging with TLS encryption.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `certificate` - Certificate used to communicate with Syslog server.
* `custom_field_name` - Custom field name for CEF format logging. The structure of `custom_field_name` block is documented below.
* `interface_select_method` - Specify how to select outgoing interface to reach server. (`ver 6.2.6,6.4.2,6.6.0`)
* `interface` - Specify outgoing interface to reach server. (`ver 6.2.6,6.4.2,6.6.0`)
* `syslog_type` - Hidden setting index of Syslog. (`ver <= 6.2.0`)

The `custom_field_name` block supports:

* `id` - Entry ID.
* `name` - Field name.
* `custom` - Field custom name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogSyslogd3 OverrideSetting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logsyslogd3_overridesetting.labelname LogSyslogd3OverrideSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
