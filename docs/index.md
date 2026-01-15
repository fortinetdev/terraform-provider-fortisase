---
page_title: "fortisase Provider"
description: |-
  
---

# fortisase Provider

The FortiSASE provider is used to interact with the resources supported by FortiSASE.



## Example Usage

```terraform
terraform {
  required_providers {
    fortisase = {
      version = "~> 1.0.0"
      source  = "fortinet.com/fortinetdev/fortisase"
    }
  }
}

provider "fortisase" {
  # method1: username and password
  username = "ABCDEFG"
  password = "ABCDEFG"

  # method2: access_token
  # access_token = "ABCDEFG"
}
```

## Schema

### Optional

- `access_token` (String) The access token of API user.
- `password` (String) The password of API user.
- `refresh_token` (String) The refresh token of API user.
- `username` (String) The username of API user.
