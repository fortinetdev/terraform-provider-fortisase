# GUI: Endpoint management -> Configuration -> Profiles
resource "fortisase_endpoint_profile" "endpoint_profile" {
  primary_key = "exampleEndpointProfile"
  enabled     = true
}
