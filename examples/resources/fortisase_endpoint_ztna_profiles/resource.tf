resource "fortisase_endpoint_policies" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}

resource "fortisase_endpoint_ztna_profiles" "example" {
  primary_key             = fortisase_endpoint_policies.endpoint_profile.primary_key
  allow_automatic_sign_on = "disable"
  connection_rules = [
    {
      id         = 1
      address    = "192.168.1.1"
      uid        = "1"
      gateways   = []
      mask       = "255.255.255.0"
      name       = "test"
      port       = "80"
      encryption = "enable"
    }
  ]
}
