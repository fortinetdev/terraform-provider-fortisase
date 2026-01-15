---
page_title: "Authenticating Terraform FortiSASE with Username/Password or API Token"
subcategory: ""
description: |-
  Obtain Your API Username and Password
---

## Obtain Your API Username and Password

1. Log into the [FortiCloud Identity & Access Management (IAM) portal](https://support.fortinet.com/iam) as an IAM user with Admin permissions.

2. Set up a resource-based permission profile of type Local allowing IAM users to access FortiSASE as a portal. See [Creating a permission profile](https://docs.fortinet.com/document/forticloud/latest/identity-access-management-iam/836213/creating-a-permission-profile).

3. Create a new API user, select the desired permission profile, and download credentials for this API user. See [Adding an API user](https://docs.fortinet.com/document/forticloud/latest/identity-access-management-iam/282341/adding-an-api-user).

Once you have obtained the API username and password, you can configure Terraform FortiSASE as shown below:

```terraform
terraform {
  required_providers {
    fortisase = {
      version = "~> 1.0"
      source  = "fortinet.com/fortinetdev/fortisase"
    }
  }
}

provider "fortisase" {
  # method1: username and password
  username = "ABCDEFG"
  password = "ABCDEFG"
}
```

## Generate an API token

Alternatively, you can authenticate using a temporary API token that expires in one hour.

Sending a POST request to `https://customerapiauth.fortinet.com/api/v1/oauth/token/`
```bash
curl -X POST "https://customerapiauth.fortinet.com/api/v1/oauth/token/" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "<your username>",
    "password": "<your password>",
    "client_id": "FortiSASE",
    "client_secret": "",
    "grant_type": "password"
  }'
```

A successful response returns a JSON object similar to the following:
```json
{
    "access_token": "<access token value>",
    "expires_in": 3600,
    "token_type": "Bearer",
    "scope": "read write",
    "refresh_token": "<refresh token value>",
    "message": "successfully authenticated",
    "status": "success"
}
```

"The "access_token" field contains the token you will use for authentication.

You can configure Terraform FortiSASE with the token as follows:

```terraform
terraform {
  required_providers {
    fortisase = {
      version = "~> 1.0"
      source  = "fortinet.com/fortinetdev/fortisase"
    }
  }
}

provider "fortisase" {
  # method2: access_token
  access_token = "ABCDEFG"
}
```
