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
