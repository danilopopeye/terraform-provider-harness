---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_azure_cloud_provider Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating an Azure Cloud Provider in Harness.
---

# harness_platform_connector_azure_cloud_provider (Resource)

Resource for creating an Azure Cloud Provider in Harness.

## Example Usage

```terraform
resource "harness_platform_connector_azure_cloud_provider" "manual_config_secret" {
  identifier  = "identifier"
  name        = "name"
  description = "example"
  tags        = ["foo:bar"]

  credentials {
    type = "ManualConfig"
    azure_manual_details {
      application_id = "application_id"
      tenant_id      = "tenant_id"
      auth {
        type = "Secret"
        azure_client_secret_key {
          secret_ref = "account.${harness_platform_secret_text.test.id}"
        }
      }
    }
  }

  azure_environment_type = "AZURE"
  delegate_selectors     = ["harness-delegate"]
}

resource "harness_platform_connector_azure_cloud_provider" "manual_config_certificate" {
  identifier  = "identifier"
  name        = "name"
  description = "example"
  tags        = ["foo:bar"]

  credentials {
    type = "ManualConfig"
    azure_manual_details {
      application_id = "application_id"
      tenant_id      = "tenant_id"
      auth {
        type = "Certificate"
        azure_client_key_cert {
          certificate_ref = "account.${harness_platform_secret_text.test.id}"
        }
      }
    }
  }

  azure_environment_type = "AZURE"
  delegate_selectors     = ["harness-delegate"]
}

resource "harness_platform_connector_azure_cloud_provider" "inherit_from_delegate_user_assigned_managed_identity" {
  identifier  = "identifier"
  name        = "name"
  description = "example"
  tags        = ["foo:bar"]

  credentials {
    type = "InheritFromDelegate"
    azure_inherit_from_delegate_details {
      auth {
        azure_msi_auth_ua {
          client_id = "client_id"
        }
        type = "UserAssignedManagedIdentity"
      }
    }
  }

  azure_environment_type = "AZURE"
  delegate_selectors     = ["harness-delegate"]
}

resource "harness_platform_connector_azure_cloud_provider" "inherit_from_delegate_system_assigned_managed_identity" {
  identifier  = "identifier"
  name        = "name"
  description = "example"
  tags        = ["foo:bar"]

  credentials {
    type = "InheritFromDelegate"
    azure_inherit_from_delegate_details {
      auth {
        type = "SystemAssignedManagedIdentity"
      }
    }
  }

  azure_environment_type = "AZURE"
  delegate_selectors     = ["harness-delegate"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `credentials` (Block List, Min: 1, Max: 1) Contains Azure connector credentials. (see [below for nested schema](#nestedblock--credentials))
- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.

### Optional

- `azure_environment_type` (String) Specifies the Azure Environment type, which is AZURE by default. Can either be AZURE or AZURE_US_GOVERNMENT
- `delegate_selectors` (Set of String) Connect using only the delegates which have these tags.
- `description` (String) Description of the resource.
- `execute_on_delegate` (Boolean) Execute on delegate or not.
- `org_id` (String) Unique identifier of the Organization.
- `project_id` (String) Unique identifier of the Project.
- `tags` (Set of String) Tags to associate with the resource. Tags should be in the form `name:value`.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--credentials"></a>
### Nested Schema for `credentials`

Required:

- `type` (String) Type can either be InheritFromDelegate or ManualConfig.

Optional:

- `azure_inherit_from_delegate_details` (Block List, Max: 1) Authenticate to Azure Cloud Provider using details inheriting from delegate. (see [below for nested schema](#nestedblock--credentials--azure_inherit_from_delegate_details))
- `azure_manual_details` (Block List, Max: 1) Authenticate to Azure Cloud Provider using manual details. (see [below for nested schema](#nestedblock--credentials--azure_manual_details))

<a id="nestedblock--credentials--azure_inherit_from_delegate_details"></a>
### Nested Schema for `credentials.azure_inherit_from_delegate_details`

Optional:

- `auth` (Block List, Max: 1) Auth to authenticate to Azure Cloud Provider using details inheriting from delegate. (see [below for nested schema](#nestedblock--credentials--azure_inherit_from_delegate_details--auth))

<a id="nestedblock--credentials--azure_inherit_from_delegate_details--auth"></a>
### Nested Schema for `credentials.azure_inherit_from_delegate_details.auth`

Required:

- `type` (String) Type can either be SystemAssignedManagedIdentity or UserAssignedManagedIdentity.

Optional:

- `azure_msi_auth_ua` (Block List, Max: 1) Azure UserAssigned MSI auth details. (see [below for nested schema](#nestedblock--credentials--azure_inherit_from_delegate_details--auth--azure_msi_auth_ua))

<a id="nestedblock--credentials--azure_inherit_from_delegate_details--auth--azure_msi_auth_ua"></a>
### Nested Schema for `credentials.azure_inherit_from_delegate_details.auth.azure_msi_auth_ua`

Optional:

- `client_id` (String) Client Id of the ManagedIdentity resource.




<a id="nestedblock--credentials--azure_manual_details"></a>
### Nested Schema for `credentials.azure_manual_details`

Optional:

- `application_id` (String) Application ID of the Azure App.
- `auth` (Block List, Max: 1) Contains Azure auth details. (see [below for nested schema](#nestedblock--credentials--azure_manual_details--auth))
- `tenant_id` (String) The Azure Active Directory (AAD) directory ID where you created your application.

<a id="nestedblock--credentials--azure_manual_details--auth"></a>
### Nested Schema for `credentials.azure_manual_details.auth`

Optional:

- `azure_client_key_cert` (Block List, Max: 1) Azure client key certificate details. (see [below for nested schema](#nestedblock--credentials--azure_manual_details--auth--azure_client_key_cert))
- `azure_client_secret_key` (Block List, Max: 1) Azure Client Secret Key details. (see [below for nested schema](#nestedblock--credentials--azure_manual_details--auth--azure_client_secret_key))
- `type` (String) Type can either be Certificate or Secret.

<a id="nestedblock--credentials--azure_manual_details--auth--azure_client_key_cert"></a>
### Nested Schema for `credentials.azure_manual_details.auth.azure_client_key_cert`

Optional:

- `certificate_ref` (String) Reference of the secret for the certificate.


<a id="nestedblock--credentials--azure_manual_details--auth--azure_client_secret_key"></a>
### Nested Schema for `credentials.azure_manual_details.auth.azure_client_secret_key`

Optional:

- `secret_ref` (String) Reference of the secret for the secret key.

## Import

Import is supported using the following syntax:

```shell
# Import using azure cloud provider connector id
terraform import harness_platform_connector_azure_cloud_provider.example <connector_id>
```