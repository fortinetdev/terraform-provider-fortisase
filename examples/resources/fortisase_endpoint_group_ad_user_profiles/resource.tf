resource "fortisase_endpoint_policies" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}

resource "fortisase_endpoint_group_ad_user_profiles" "endpoint_group_ad_user_profile" {
  primary_key = fortisase_endpoint_policies.endpoint_profile.primary_key
}
