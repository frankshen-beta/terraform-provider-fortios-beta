---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_userbookmark"
description: |-
  Configure SSL VPN user bookmark.
---

# fortios_vpnsslweb_userbookmark
Configure SSL VPN user bookmark.

## Example Usage

```hcl
resource "fortios_vpnsslweb_userbookmark" "trname" {
  custom_lang = "big5"
  name        = "userbookmarks1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) User and group name.
* `custom_lang` - Personal language.
* `bookmarks` - Bookmark table. The structure of `bookmarks` block is documented below.

The `bookmarks` block supports:

* `name` - Bookmark name.
* `apptype` - Application type.
* `url` - URL parameter.
* `host` - Host name/IP parameter.
* `folder` - Network shared file folder parameter.
* `domain` - Login domain. (`ver >= 6.4.2`)
* `additional_params` - Additional parameters.
* `listening_port` - Listening port (0 - 65535).
* `remote_port` - Remote port (0 - 65535).
* `show_status_window` - Enable/disable showing of status window.
* `description` - Description.
* `server_layout` - Server side keyboard layout.
* `security` - Security mode for RDP connection.
* `preconnection_id` - The numeric ID of the RDP source (0-2147483648).
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `port` - Remote port.
* `logon_user` - Logon user.
* `logon_password` - Logon password.
* `sso` - Single Sign-On.
* `form_data` - Form data. The structure of `form_data` block is documented below.
* `sso_credential` - Single sign-on credentials.
* `sso_username` - SSO user name.
* `sso_password` - SSO password.
* `sso_credential_sent_once` - Single sign-on credentials are only sent once to remote server.

The `form_data` block supports:

* `name` - Name.
* `value` - Value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnSslWeb UserBookmark can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpnsslweb_userbookmark.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
