---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_peer"
description: |-
  Configure peer users.
---

# fortios_user_peer
Configure peer users.

## Example Usage

```hcl
resource "fortios_user_peer" "trname1" {
  ca                  = "EC-ACC"
  cn_type             = "string"
  ldap_mode           = "password"
  mandatory_ca_verify = "enable"
  name                = "u1"
  two_factor          = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Peer name.
* `mandatory_ca_verify` - Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid.
* `ca` - Name of the CA certificate as returned by the execute vpn certificate ca list command.
* `subject` - Peer certificate name constraints.
* `cn` - Peer certificate common name.
* `cn_type` - Peer certificate common name type.
* `ldap_server` - Name of an LDAP server defined under the user ldap command. Performs client access rights check.
* `ldap_username` - Username for LDAP server bind.
* `ldap_password` - Password for LDAP server bind.
* `ldap_mode` - Mode for LDAP peer authentication.
* `ocsp_override_server` - Online Certificate Status Protocol (OCSP) server for certificate retrieval.
* `two_factor` - Enable/disable two-factor authentication, applying certificate and password-based authentication.
* `passwd` - Peer's password used for two-factor authentication.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Peer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_peer.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
