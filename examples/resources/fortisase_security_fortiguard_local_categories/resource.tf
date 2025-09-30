resource "fortisase_security_fortiguard_local_categories" "fortiguard_local_category" {
  primary_key   = "example_name"
  threat_weight = "low"
  urls          = ["test", "test2"]
}
