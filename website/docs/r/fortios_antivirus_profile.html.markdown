---
subcategory: "FortiGate AntiVirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_profile"
description: |-
  Configure AntiVirus profiles.
---

# fortios_antivirus_profile
Configure AntiVirus profiles.

## Example Usage

```hcl
resource "fortios_antivirus_profile" "trname" {
  analytics_bl_filetype = 0
  analytics_db          = "disable"
  analytics_max_upload  = 10
  analytics_wl_filetype = 0
  av_block_log          = "enable"
  av_virus_log          = "enable"
  extended_log          = "disable"
  ftgd_analytics        = "disable"
  inspection_mode       = "flow-based"
  mobile_malware_db     = "enable"
  name                  = "1"
  scan_mode             = "quick"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Profile name.
* `comment` - Comment.
* `replacemsg_group` - Replacement message group customized for this profile.
* `feature_set` - Flow/proxy feature set. (`ver >= 6.4.0`)
* `inspection_mode` - Inspection mode. (`ver <= 6.2.0`)
* `ftgd_analytics` - Settings to control which files are uploaded to FortiSandbox.
* `analytics_max_upload` - Maximum size of files that can be uploaded to FortiSandbox (1 - 395 MBytes, default = 10).
* `analytics_ignore_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox. (`ver >= 6.6.0`)
* `analytics_accept_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox. (`ver >= 6.6.0`)
* `analytics_wl_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox. (`ver <= 6.4.2`)
* `analytics_bl_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox. (`ver <= 6.4.2`)
* `analytics_db` - Enable/disable using the FortiSandbox signature database to supplement the AV signature databases.
* `mobile_malware_db` - Enable/disable using the mobile malware signature database.
* `http` - Configure HTTP AntiVirus options. The structure of `http` block is documented below.
* `ftp` - Configure FTP AntiVirus options. The structure of `ftp` block is documented below.
* `imap` - Configure IMAP AntiVirus options. The structure of `imap` block is documented below.
* `pop3` - Configure POP3 AntiVirus options. The structure of `pop3` block is documented below.
* `smtp` - Configure SMTP AntiVirus options. The structure of `smtp` block is documented below.
* `mapi` - Configure MAPI AntiVirus options. The structure of `mapi` block is documented below.
* `nntp` - Configure NNTP AntiVirus options. The structure of `nntp` block is documented below.
* `cifs` - Configure CIFS AntiVirus options. The structure of `cifs` block is documented below.
* `ssh` - Configure SFTP and SCP AntiVirus options. The structure of `ssh` block is documented below.
* `smb` - Configure SMB AntiVirus options. The structure of `smb` block is documented below. (`ver <= 6.2.0`)
* `nac_quar` - Configure AntiVirus quarantine settings. The structure of `nac_quar` block is documented below.
* `outbreak_prevention` - Configure Virus Outbreak Prevention settings. The structure of `outbreak_prevention` block is documented below.
* `content_disarm` - AV Content Disarm and Reconstruction settings. The structure of `content_disarm` block is documented below.
* `av_virus_log` - Enable/disable AntiVirus logging.
* `av_block_log` - Enable/disable logging for AntiVirus file blocking.
* `extended_log` - Enable/disable extended logging for antivirus.
* `scan_mode` - Choose between full scan mode and quick scan mode.

The `http` block supports:

* `options` - Enable/disable HTTP AntiVirus scanning, monitoring, and quarantine.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol.

The `ftp` block supports:

* `options` - Enable/disable FTP AntiVirus scanning, monitoring, and quarantine.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.

The `imap` block supports:

* `options` - Enable/disable IMAP AntiVirus scanning, monitoring, and quarantine.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol.

The `pop3` block supports:

* `options` - Enable/disable POP3 AntiVirus scanning, monitoring, and quarantine.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol.

The `smtp` block supports:

* `options` - Enable/disable SMTP AntiVirus scanning, monitoring, and quarantine.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol.

The `mapi` block supports:

* `options` - Enable/disable MAPI AntiVirus scanning, monitoring, and quarantine.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.

The `nntp` block supports:

* `options` - Enable/disable NNTP AntiVirus scanning, monitoring, and quarantine.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.

The `cifs` block supports:

* `options` - Enable/disable CIFS AntiVirus scanning, monitoring, and quarantine.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.

The `ssh` block supports:

* `options` - Enable/disable SFTP and SCP AntiVirus scanning, monitoring, and quarantine.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.

The `smb` block supports (`ver <= 6.2.0`):

* `options` - Enable/disable SMB AntiVirus scanning, monitoring, and quarantine.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.

The `nac_quar` block supports:

* `infected` - Enable/Disable quarantining infected hosts to the banned user list.
* `expiry` - Duration of quarantine.
* `log` - Enable/disable AntiVirus quarantine logging.

The `outbreak_prevention` block supports:

* `ftgd_service` - Enable/disable FortiGuard Virus outbreak prevention service.
* `external_blocklist` - Enable/disable external malware blocklist.

The `content_disarm` block supports:

* `original_file_destination` - Destination to send original file if active content is removed.
* `error_action` - Action to be taken if CDR engine encounters an unrecoverable error. (`ver >= 6.4.2`)
* `office_macro` - Enable/disable stripping of macros in Microsoft Office documents.
* `office_hylink` - Enable/disable stripping of hyperlinks in Microsoft Office documents.
* `office_linked` - Enable/disable stripping of linked objects in Microsoft Office documents.
* `office_embed` - Enable/disable stripping of embedded objects in Microsoft Office documents.
* `office_dde` - Enable/disable stripping of Dynamic Data Exchange events in Microsoft Office documents.
* `office_action` - Enable/disable stripping of PowerPoint action events in Microsoft Office documents.
* `pdf_javacode` - Enable/disable stripping of JavaScript code in PDF documents.
* `pdf_embedfile` - Enable/disable stripping of embedded files in PDF documents.
* `pdf_hyperlink` - Enable/disable stripping of hyperlinks from PDF documents.
* `pdf_act_gotor` - Enable/disable stripping of links to other PDFs in PDF documents.
* `pdf_act_launch` - Enable/disable stripping of links to external applications in PDF documents.
* `pdf_act_sound` - Enable/disable stripping of embedded sound files in PDF documents.
* `pdf_act_movie` - Enable/disable stripping of embedded movies in PDF documents.
* `pdf_act_java` - Enable/disable stripping of actions that execute JavaScript code in PDF documents.
* `pdf_act_form` - Enable/disable stripping of actions that submit data to other targets in PDF documents.
* `cover_page` - Enable/disable inserting a cover page into the disarmed document.
* `detect_only` - Enable/disable only detect disarmable files, do not alter content.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Antivirus Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_antivirus_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
