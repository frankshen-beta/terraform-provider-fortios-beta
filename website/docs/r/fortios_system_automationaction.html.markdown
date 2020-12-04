---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationaction"
description: |-
  Action for automation stitches.
---

# fortios_system_automationaction
Action for automation stitches.

## Example Usage

```hcl
resource "fortios_system_automationaction" "trname" {
  action_type      = "email"
  aws_domain       = "amazonaws.com"
  delay            = 0
  email_subject    = "SUBJECT1"
  method           = "post"
  minimum_interval = 1
  name             = "1"
  protocol         = "http"
  required         = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `action_type` - Action type.
* `tls_certificate` - Custom TLS certificate for API request.
* `email_to` - Email addresses. The structure of `email_to` block is documented below.
* `email_from` - Email sender name.
* `email_subject` - Email subject.
* `email_body` - Email body.
* `minimum_interval` - Limit execution to no more than once in this interval (in seconds).
* `delay` - Delay before execution (in seconds).
* `required` - Required in action chain.
* `aws_api_id` - AWS API Gateway ID.
* `aws_region` - AWS region.
* `aws_domain` - AWS domain.
* `aws_api_stage` - AWS API Gateway deployment stage name.
* `aws_api_path` - AWS API Gateway path.
* `aws_api_key` - AWS API Gateway API key.
* `azure_app` - Azure function application name.
* `azure_function` - Azure function name.
* `azure_domain` - Azure function domain.
* `azure_function_authorization` - Azure function authorization level.
* `azure_api_key` - Azure function API key.
* `gcp_function_region` - Google Cloud function region.
* `gcp_project` - Google Cloud Platform project name.
* `gcp_function_domain` - Google Cloud function domain.
* `gcp_function` - Google Cloud function name.
* `alicloud_account_id` - AliCloud account ID.
* `alicloud_region` - AliCloud region.
* `alicloud_function_domain` - AliCloud function domain.
* `alicloud_version` - AliCloud version.
* `alicloud_service` - AliCloud service name.
* `alicloud_function` - AliCloud function name.
* `alicloud_function_authorization` - AliCloud function authorization type.
* `alicloud_access_key_id` - AliCloud AccessKey ID.
* `alicloud_access_key_secret` - AliCloud AccessKey secret.
* `message` - Message content. (`ver >= 6.4.0`)
* `replacement_message` - Enable/disable replacement message. (`ver >= 6.6.0`)
* `protocol` - Request protocol.
* `method` - Request method (POST, PUT, GET, PATCH or DELETE).
* `uri` - Request API URI.
* `http_body` - Request body (if necessary). Should be serialized json string.
* `port` - Protocol port.
* `headers` - Request headers. The structure of `headers` block is documented below.
* `script` - CLI script.
* `accprofile` - Access profile for CLI script action to access FortiGate features. (`ver >= 6.6.0`)
* `security_tag` - NSX security tag.
* `sdn_connector` - NSX SDN connector names. The structure of `sdn_connector` block is documented below.

The `email_to` block supports:

* `name` - Email address.

The `headers` block supports:

* `header` - Request header.

The `sdn_connector` block supports:

* `name` - SDN connector name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationAction can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_automationaction.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
