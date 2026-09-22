# GUI: Endpoint management -> Configuration -> Profiles
resource "fortisase_endpoint_profile" "endpoint_profile" {
  primary_key = "exampleEndpointProfile"
  enabled     = true
}

resource "fortisase_endpoint_fsso_profile" "example" {
  primary_key = fortisase_endpoint_profile.endpoint_profile.primary_key
  port        = 443
}
