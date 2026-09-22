# GUI: Endpoint management -> Configuration -> Profiles
resource "fortisase_endpoint_profile" "endpoint_profile" {
  primary_key = "exampleEndpointProfile"
  enabled     = true
}

resource "fortisase_endpoint_group_ad_user_profile" "endpoint_group_ad_user_profile" {
  primary_key = fortisase_endpoint_profile.endpoint_profile.primary_key
}
