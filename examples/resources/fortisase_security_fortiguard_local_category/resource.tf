resource "fortisase_security_fortiguard_local_category" "fortiguard_local_category" {
  primary_key   = "example_name"
  threat_weight = "low"
  urls          = ["test", "test2"]
}
