---
subcategory: "FortiGate SSHFilter"
layout: "fortios"
page_title: "FortiOS: fortios_sshfilter_profile"
description: |-
  SSH filter profile.
---

# fortios_sshfilter_profile
SSH filter profile.

## Example Usage

```hcl
resource "fortios_sshfilter_profile" "trname" {
  block               = "x11"
  default_command_log = "enable"
  log                 = "x11"
  name                = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) SSH filter profile name.
* `block` - SSH blocking options.
* `log` - SSH logging options.
* `default_command_log` - Enable/disable logging unmatched shell commands.
* `shell_commands` - SSH command filter. The structure of `shell_commands` block is documented below.
* `file_filter` - File filter. The structure of `file_filter` block is documented below. (`ver <= 6.4.0`)

The `shell_commands` block supports:

* `id` - Id.
* `type` - Matching type.
* `pattern` - SSH shell command pattern.
* `action` - Action to take for URL filter matches.
* `log` - Enable/disable logging.
* `alert` - Enable/disable alert.
* `severity` - Log severity.

The `file_filter` block supports (`ver <= 6.4.0`):

* `status` - Enable/disable file filter.
* `log` - Enable/disable file filter logging.
* `scan_archive_contents` - Enable/disable file filter archive contents scan.
* `entries` - File filter entries. The structure of `entries` block is documented below.

The `entries` block supports:

* `filter` - Add a file filter.
* `comment` - Comment.
* `action` - Action taken for matched file.
* `direction` - Match files transmitted in the session's originating or reply direction.
* `password_protected` - Match password-protected files.
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block supports:

* `name` - File type name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SshFilter Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_sshfilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
