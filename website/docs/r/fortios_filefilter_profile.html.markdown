---
subcategory: "FortiGate File-Filter"
layout: "fortios"
page_title: "FortiOS: fortios_filefilter_profile"
description: |-
  Configure file-filter profiles.
---

# fortios_filefilter_profile
Configure file-filter profiles. Applies to FortiOS Version `>= 6.4.2`.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Profile name.
* `comment` - Comment.
* `feature_set` - Flow/proxy feature set.
* `replacemsg_group` - Replacement message group
* `log` - Enable/disable file-filter logging.
* `extended_log` - Enable/disable file-filter extended logging.
* `scan_archive_contents` - Enable/disable archive contents scan. (Not for CIFS)
* `rules` - File filter rules. The structure of `rules` block is documented below.

The `rules` block supports:

* `name` - File-filter rule name.
* `comment` - Comment.
* `protocol` - Protocols to apply rule to.
* `action` - Action taken for matched file.
* `direction` - Traffic direction. (HTTP, FTP, SSH, CIFS only)
* `password_protected` - Match password-protected files.
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block supports:

* `name` - File type name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FileFilter Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_filefilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
