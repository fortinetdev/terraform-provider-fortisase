# GUI: Endpoint management -> Configuration -> Profiles
resource "fortisase_endpoint_profile" "endpoint_profile" {
  primary_key = "example_endpoint_profile"
  enabled     = true
}

resource "fortisase_endpoint_group_ad_user_profile" "endpoint_group_ad_user_profile" {
  primary_key = fortisase_fortisase_endpoint_profileendpoint_policy.endpoint_profile.primary_key
}
