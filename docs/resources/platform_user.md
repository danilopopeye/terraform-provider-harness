---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_user Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Harness User. This requires your authentication mechanism to be set to SAML, LDAP, or OAuth, and the feature flag AUTOACCEPTSAMLACCOUNTINVITES to be enabled.
---

# harness_platform_user (Resource)

Resource for creating a Harness User. This requires your authentication mechanism to be set to SAML, LDAP, or OAuth, and the feature flag AUTO_ACCEPT_SAML_ACCOUNT_INVITES to be enabled.

## Example Usage

```terraform
# Create user at project level
resource "harness_platform_user" "example" {
  org_id      = "org_id"
  project_id  = "project_id"
  email       = "john.doe@harness.io"
  user_groups = ["_project_all_users"]
  role_bindings {
    resource_group_identifier = "_all_project_level_resources"
    role_identifier           = "_project_viewer"
    role_name                 = "Project Viewer"
    resource_group_name       = "All Project Level Resources"
    managed_role              = true
  }
}

# Create user at org level
resource "harness_platform_user" "example" {
  org_id      = "org_id"
  email       = "john.doe@harness.io"
  user_groups = ["_project_all_users"]
  role_bindings {
    resource_group_identifier = "_all_project_level_resources"
    role_identifier           = "_project_viewer"
    role_name                 = "Project Viewer"
    resource_group_name       = "All Project Level Resources"
    managed_role              = true
  }
}

# Create user at account level
resource "harness_platform_user" "example" {
  email       = "john.doe@harness.io"
  user_groups = ["_project_all_users"]
  role_bindings {
    resource_group_identifier = "_all_project_level_resources"
    role_identifier           = "_project_viewer"
    role_name                 = "Project Viewer"
    resource_group_name       = "All Project Level Resources"
    managed_role              = true
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String) The email of the user.
- `role_bindings` (Block List, Min: 1) Role Bindings of the user. Cannot be updated. (see [below for nested schema](#nestedblock--role_bindings))
- `user_groups` (Set of String) The user group of the user. Cannot be updated.

### Optional

- `org_id` (String) Organization identifier of the user.
- `project_id` (String) Project identifier of the user.

### Read-Only

- `disabled` (Boolean) Whether or not the user account is disabled.
- `externally_managed` (Boolean) Whether or not the user account is externally managed.
- `id` (String) The ID of this resource.
- `identifier` (String) Unique identifier of the user.
- `locked` (Boolean) Whether or not the user account is locked.
- `name` (String) Name of the user.

<a id="nestedblock--role_bindings"></a>
### Nested Schema for `role_bindings`

Optional:

- `managed_role` (Boolean) Managed Role of the user.
- `resource_group_identifier` (String) Resource Group Identifier of the user.
- `resource_group_name` (String) Resource Group Name of the user.
- `role_identifier` (String) Role Identifier of the user.
- `role_name` (String) Role Name Identifier of the user.

## Import

Import is supported using the following syntax:

```shell
# Import account level
terraform import harness_platform_user.john_doe <email_id>

# Import org level 
terraform import harness_platform_user.john_doe <email_id>/<org_id>

# Import project level
terraform import harness_platform_user.john_doe <email_id>/<org_id>/<project_id>
```