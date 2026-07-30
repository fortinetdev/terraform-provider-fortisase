# GUI: Security -> Bandwidth control -> Shared profile
resource "fortisase_security_traffic_shaper" "example" {
  primary_key          = "example_trafficshaper"
  bandwidth_unit       = "kbps"
  guaranteed_bandwidth = 0
  maximum_bandwidth    = 1024
  per_policy           = "disable"
  priority             = "low"
}
