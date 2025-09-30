resource "fortisase_endpoint_policies" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}

resource "fortisase_endpoint_protection_profiles" "endpoint_protection_profile" {
  primary_key        = fortisase_endpoint_policies.endpoint_profile.primary_key
  vulnerability_scan = "enable"
}
