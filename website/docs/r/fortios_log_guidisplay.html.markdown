---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_guidisplay"
description: |-
  Configure how log messages are displayed on the GUI.
---

# fortios_log_guidisplay
Configure how log messages are displayed on the GUI.

## Example Usage

```hcl
resource "fortios_log_guidisplay" "trname" {
  fortiview_unscanned_apps = "disable"
  resolve_apps             = "enable"
  resolve_hosts            = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `resolve_hosts` - Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup
* `resolve_apps` - Resolve unknown applications on the GUI using Fortinet's remote application database.
* `fortiview_unscanned_apps` - Enable/disable showing unscanned traffic in FortiView application charts.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log GuiDisplay can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_log_guidisplay.labelname LogGuiDisplay
$ unset "FORTIOS_IMPORT_TABLE"
```
