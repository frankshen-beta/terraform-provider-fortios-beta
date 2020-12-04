---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_profileprotocoloptions"
description: |-
  Configure protocol options.
---

# fortios_firewall_profileprotocoloptions
Configure protocol options.

## Example Usage

```hcl
resource "fortios_firewall_profileprotocoloptions" "trname" {
  name                    = "sp"
  oversize_log            = "disable"
  rpc_over_http           = "disable"
  switching_protocols_log = "disable"

  dns {
    ports  = 53
    status = "enable"
  }

  ftp {
    comfort_amount              = 1
    comfort_interval            = 10
    inspect_all                 = "disable"
    options                     = "splice"
    oversize_limit              = 10
    ports                       = 21
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  http {
    block_page_status_code      = 403
    comfort_amount              = 1
    comfort_interval            = 10
    fortinet_bar                = "disable"
    fortinet_bar_port           = 8011
    http_policy                 = "disable"
    inspect_all                 = "disable"
    oversize_limit              = 10
    ports                       = 80
    range_block                 = "disable"
    retry_count                 = 0
    scan_bzip2                  = "enable"
    status                      = "enable"
    streaming_content_bypass    = "enable"
    strip_x_forwarded_for       = "disable"
    switching_protocols         = "bypass"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  imap {
    inspect_all                 = "disable"
    options                     = "fragmail"
    oversize_limit              = 10
    ports                       = 143
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  mail_signature {
    status = "disable"
  }

  mapi {
    options                     = "fragmail"
    oversize_limit              = 10
    ports                       = 135
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  nntp {
    inspect_all                 = "disable"
    options                     = "splice"
    oversize_limit              = 10
    ports                       = 119
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  pop3 {
    inspect_all                 = "disable"
    options                     = "fragmail"
    oversize_limit              = 10
    ports                       = 110
    scan_bzip2                  = "enable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }

  smtp {
    inspect_all                 = "disable"
    options                     = "fragmail splice"
    oversize_limit              = 10
    ports                       = 25
    scan_bzip2                  = "enable"
    server_busy                 = "disable"
    status                      = "enable"
    uncompressed_nest_limit     = 12
    uncompressed_oversize_limit = 10
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `comment` - Optional comments.
* `feature_set` - Flow/proxy feature set. (`ver 6.4.0`)
* `replacemsg_group` - Name of the replacement message group to be used
* `oversize_log` - Enable/disable logging for antivirus oversize file blocking.
* `switching_protocols_log` - Enable/disable logging for HTTP/HTTPS switching protocols.
* `http` - Configure HTTP protocol options. The structure of `http` block is documented below.
* `ftp` - Configure FTP protocol options. The structure of `ftp` block is documented below.
* `imap` - Configure IMAP protocol options. The structure of `imap` block is documented below.
* `mapi` - Configure MAPI protocol options. The structure of `mapi` block is documented below.
* `pop3` - Configure POP3 protocol options. The structure of `pop3` block is documented below.
* `smtp` - Configure SMTP protocol options. The structure of `smtp` block is documented below.
* `nntp` - Configure NNTP protocol options. The structure of `nntp` block is documented below.
* `ssh` - Configure SFTP and SCP protocol options. The structure of `ssh` block is documented below.
* `dns` - Configure DNS protocol options. The structure of `dns` block is documented below.
* `cifs` - Configure CIFS protocol options. The structure of `cifs` block is documented below.
* `mail_signature` - Configure Mail signature. The structure of `mail_signature` block is documented below.
* `rpc_over_http` - Enable/disable inspection of RPC over HTTP.

The `http` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 80).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). (`ver >= 6.4.0`)
* `options` - One or more options that can be applied to the session.
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 10240 bytes, default = 1).
* `range_block` - Enable/disable blocking of partial downloads.
* `http_policy` - Enable/disable HTTP policy check. (`ver <= 6.2.0`)
* `strip_x_forwarded_for` - Enable/disable stripping of HTTP X-Forwarded-For header.
* `post_lang` - ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets).
* `fortinet_bar` - Enable/disable Fortinet bar on HTML content. (`ver <= 6.4.0`)
* `fortinet_bar_port` - Port for use by Fortinet Bar (1 - 65535, default = 8011). (`ver <= 6.4.0`)
* `streaming_content_bypass` - Enable/disable bypassing of streaming content from buffering.
* `switching_protocols` - Bypass from scanning, or block a connection that attempts to switch protocol.
* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. (`ver >= 6.4.0`)
* `tunnel_non_http` - Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port. (`ver >= 6.4.0`)
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `block_page_status_code` - Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).
* `retry_count` - Number of attempts to retry HTTP connection (0 - 100, default = 0).
* `tcp_window_type` - Specify type of TCP window to use for this protocol.
* `tcp_window_minimum` - Minimum dynamic TCP window size (default = 128KB).
* `tcp_window_maximum` - Maximum dynamic TCP window size (default = 8MB).
* `tcp_window_size` - Set TCP static window size (default = 256KB).
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `ftp` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 21).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `options` - One or more options that can be applied to the session.
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 10240 bytes, default = 1).
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.). (`ver >= 6.6.0`)
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `tcp_window_type` - TCP window type to use for this protocol. (`ver >= 6.6.0`)
* `tcp_window_minimum` - Minimum dynamic TCP window size. (`ver >= 6.6.0`)
* `tcp_window_maximum` - Maximum dynamic TCP window size. (`ver >= 6.6.0`)
* `tcp_window_size` - Set TCP static window size. (`ver >= 6.6.0`)
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `imap` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 143).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). (`ver >= 6.4.0`)
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `mapi` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 135).
* `status` - Enable/disable the active status of scanning for this protocol.
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.

The `pop3` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 110).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). (`ver >= 6.4.0`)
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `smtp` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 25).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). (`ver >= 6.4.0`)
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `server_busy` - Enable/disable SMTP server busy when server not available.
* `ssl_offloaded` - SSL decryption and encryption performed by an external device.

The `nntp` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 119).
* `status` - Enable/disable the active status of scanning for this protocol.
* `inspect_all` - Enable/disable the inspection of all ports for the protocol.
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). (`ver >= 6.4.0`)
* `options` - One or more options that can be applied to the session.
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.

The `ssh` block supports:

* `options` - One or more options that can be applied to the session.
* `comfort_interval` - Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
* `comfort_amount` - Amount of data to send in a transmission for client comforting (1 - 65535 bytes, default = 1).
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
* `stream_based_uncompressed_limit` - Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.). (`ver >= 6.6.0`)
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files.
* `tcp_window_type` - TCP window type to use for this protocol. (`ver >= 6.6.0`)
* `tcp_window_minimum` - Minimum dynamic TCP window size. (`ver >= 6.6.0`)
* `tcp_window_maximum` - Maximum dynamic TCP window size. (`ver >= 6.6.0`)
* `tcp_window_size` - Set TCP static window size. (`ver >= 6.6.0`)
* `ssl_offloaded` - SSL decryption and encryption performed by an external device. (`ver >= 6.6.0`)

The `dns` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 53).
* `status` - Enable/disable the active status of scanning for this protocol.

The `cifs` block supports:

* `ports` - Ports to scan for content (1 - 65535, default = 445).
* `status` - Enable/disable the active status of scanning for this protocol.
* `options` - One or more options that can be applied to the session. (`ver >= 6.4.0`)
* `oversize_limit` - Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10). (`ver >= 6.4.0`)
* `uncompressed_oversize_limit` - Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10). (`ver >= 6.4.0`)
* `uncompressed_nest_limit` - Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12). (`ver >= 6.4.0`)
* `scan_bzip2` - Enable/disable scanning of BZip2 compressed files. (`ver >= 6.4.0`)
* `tcp_window_type` - Specify type of TCP window to use for this protocol. (`ver >= 6.4.0`)
* `tcp_window_minimum` - Minimum dynamic TCP window size (default = 128KB). (`ver >= 6.4.0`)
* `tcp_window_maximum` - Maximum dynamic TCP window size (default = 8MB). (`ver >= 6.4.0`)
* `tcp_window_size` - Set TCP static window size (default = 256KB). (`ver >= 6.4.0`)
* `server_credential_type` - CIFS server credential type. (`ver 6.2.6,6.4.2,6.6.0`)
* `domain_controller` - Domain for which to decrypt CIFS traffic. (`ver >= 6.4.2`)
* `server_keytab` - Server keytab. The structure of `server_keytab` block is documented below. (`ver 6.2.6,6.4.2,6.6.0`)

The `server_keytab` block supports (`ver 6.2.6,6.4.2,6.6.0`):

* `principal` - Service principal.  For example, "host/cifsserver.example.com@example.com".
* `keytab` - Base64 encoded keytab file containing credential of the server.

The `mail_signature` block supports:

* `status` - Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate.
* `signature` - Email signature to be added to outgoing email (if the signature contains spaces, enclose with quotation marks).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProfileProtocolOptions can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_profileprotocoloptions.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
