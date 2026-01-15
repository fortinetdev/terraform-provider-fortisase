# Same behavior as "fortisase_endpoint_profile"
resource "fortisase_endpoint_policies" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}
