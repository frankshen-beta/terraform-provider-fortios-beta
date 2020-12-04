---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_snmpuser"
description: |-
  Configure FortiSwitch SNMP v3 users globally.
---

# fortios_switchcontroller_snmpuser
Configure FortiSwitch SNMP v3 users globally.

## Argument Reference

The following arguments are supported:

* `name` - (Required) SNMP user name.
* `queries` - Enable/disable SNMP queries for this user.
* `query_port` - SNMPv3 query port (default = 161).
* `security_level` - Security level for message authentication and encryption.
* `auth_proto` - Authentication protocol.
* `auth_pwd` - Password for authentication protocol.
* `priv_proto` - Privacy (encryption) protocol.
* `priv_pwd` - Password for privacy (encryption) protocol.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SnmpUser can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_snmpuser.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
