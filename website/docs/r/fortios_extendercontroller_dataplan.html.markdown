---
subcategory: "FortiGate Extender-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extendercontroller_dataplan"
description: |-
  FortiExtender dataplan configuration.
---

# fortios_extendercontroller_dataplan
FortiExtender dataplan configuration. Applies to FortiOS Version `>= 6.4.2`.

## Argument Reference

The following arguments are supported:

* `name` - (Required) FortiExtender dataplan name
* `modem_id` - Dataplan's modem specifics, if any.
* `type` - Type preferences configuration.
* `slot` - SIM slot configuration.
* `iccid` - ICCID configuration.
* `carrier` - Carrier configuration.
* `APN` - APN configuration.
* `auth_type` - Authentication type.
* `username` - Username.
* `password` - Password.
* `PDN` - PDN type.
* `signal_threshold` - Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
* `signal_period` - Signal period (600 to 18000 seconds).
* `capacity` - Capacity in MB (0 - 102400000).
* `monthly_fee` - Monthly fee of dataplan (0 - 100000, in local currency).
* `billing_date` - Billing day of the month (1 - 31).
* `overage` - Enable/disable dataplan overage detection.
* `preferred_subnet` - Preferred subnet mask (8 - 32).
* `private_network` - Enable/disable dataplan private network support.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtenderController Dataplan can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_extendercontroller_dataplan.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
