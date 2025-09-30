resource "fortisase_endpoint_policies" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}

resource "fortisase_endpoint_setting_profiles" "endpoint_setting_profile" {
  primary_key = fortisase_endpoint_policies.endpoint_profile.primary_key
}
