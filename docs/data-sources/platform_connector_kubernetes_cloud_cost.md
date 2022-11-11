---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_kubernetes_cloud_cost Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Datasource for looking up a Kubernetes Cloud Cost connector.
---

# harness_platform_connector_kubernetes_cloud_cost (Data Source)

Datasource for looking up a Kubernetes Cloud Cost connector.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `org_id` (String) Unique identifier of the Organization.
- `project_id` (String) Unique identifier of the Project.

### Read-Only

- `connector_ref` (String) Reference of the Connector.
- `description` (String) Description of the resource.
- `features_enabled` (Set of String) Indicates which feature to enable among Billing, Optimization, and Visibility.
- `id` (String) The ID of this resource.
- `tags` (Set of String) Tags to associate with the resource. Tags should be in the form `name:value`.

