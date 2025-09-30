resource "fortisase_endpoint_policies" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}

resource "fortisase_endpoint_sandbox_profiles" "endpoint_sandbox_profile" {
  primary_key = fortisase_endpoint_policies.endpoint_profile.primary_key
}
