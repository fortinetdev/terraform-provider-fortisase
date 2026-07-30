# GUI: Endpoint management -> Configuration -> Profiles
resource "fortisase_endpoint_profile" "endpoint_profile" {
  primary_key = "example_endpoint_profile"
  enabled     = true
}

resource "fortisase_endpoint_ztna_profile" "example" {
  primary_key             = fortisase_endpoint_profile.endpoint_profile.primary_key
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
