---
page_title: "Resource Type Name Migration"
subcategory: ""
description: |-
  Learn how to migrate Terraform state from a deprecated plural FortiSASE resource type to its singular replacement without recreating the remote object.
---

# Migrate Resources to Singular Type Names

The FortiSASE Terraform provider is transitioning resource type names from plural to singular forms. Plural resource types remain functional for backward compatibility, but you should migrate existing configurations to their singular replacements.

This guide demonstrates how to migrate the deprecated `fortisase_auth_users` resource type to `fortisase_auth_user` while preserving the existing Terraform state and remote user.

## Prerequisites

- Terraform CLI 1.8 or later.
- A version of the FortiSASE provider that supports both `fortisase_auth_users` and `fortisase_auth_user`, including state moves between them.
- A backup of the current Terraform state. Store the backup securely because Terraform state may contain sensitive values.

You can create a local state backup before starting:

```bash
terraform state pull > terraform-state-backup.json
```

## Existing configuration

Assume that the user has already been created with the deprecated resource type:

```terraform
resource "fortisase_auth_users" "user" {
  primary_key = "example@example.com"
  auth_type   = "password"
  status      = "enable"
  email       = "example@example.com"
  password    = "password"
}
```

## Update the resource type name

Change only the resource type from `fortisase_auth_users` to `fortisase_auth_user`. Keep the local resource name (`user`) and the resource arguments unchanged:

```terraform
resource "fortisase_auth_user" "user" {
  primary_key = "example@example.com"
  auth_type   = "password"
  status      = "enable"
  email       = "example@example.com"
  password    = "password"
}
```

## Declare the state move

Add the following `moved` block in the same Terraform module as the resource:

```terraform
moved {
  from = fortisase_auth_users.user
  to   = fortisase_auth_user.user
}
```

The `moved` block tells Terraform that both addresses represent the same remote object. The provider transfers the existing state to the new resource type without calling the create or delete API.

## Review and apply the migration

Upgrade the provider dependency if necessary, and then create a plan:

```bash
terraform init -upgrade
terraform plan
```

The plan should report that the resource has moved, similar to the following:

```text
# fortisase_auth_users.user has moved to fortisase_auth_user.user
```

Do not apply the plan if it proposes destroying `fortisase_auth_users.user` or creating a new `fortisase_auth_user.user`. Confirm that the installed Terraform CLI and FortiSASE provider versions satisfy the prerequisites before trying again.

After confirming that the plan contains only the expected state move, apply it:

```bash
terraform apply
```

Optionally, verify the new state address:

```bash
terraform state show fortisase_auth_user.user
```

Keep the `moved` block in the configuration until every workspace and environment that uses the module has completed the migration. Removing it too early can cause Terraform to interpret the old and new addresses as different resources in environments that have not yet migrated.
