# GUI: Endpoint management -> Configuration -> Profiles
resource "fortisase_endpoint_profile" "endpoint_profile" {
  primary_key = "exampleEndpointProfile"
  enabled     = true
}

resource "fortisase_endpoint_ztna_profile" "example" {
  primary_key             = fortisase_endpoint_profile.endpoint_profile.primary_key
  status                  = "enable"
  allow_automatic_sign_on = "disable"
}
