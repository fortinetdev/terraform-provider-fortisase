resource "fortisase_endpoint_policies" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}

resource "fortisase_endpoint_fsso_profiles" "example" {
  primary_key = fortisase_endpoint_policies.endpoint_profile.primary_key
  port        = 443
}
