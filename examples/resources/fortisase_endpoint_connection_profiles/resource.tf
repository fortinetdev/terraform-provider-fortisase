resource "fortisase_endpoint_policies" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}

resource "fortisase_endpoint_connection_profiles" "connection_profile" {
  primary_key = fortisase_endpoint_policies.endpoint_profile.primary_key
  lockdown = {
    max_attempts = 5
  }
}
