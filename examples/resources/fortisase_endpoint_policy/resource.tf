# Same behavior as "fortisase_endpoint_profile"
resource "fortisase_endpoint_policy" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}
