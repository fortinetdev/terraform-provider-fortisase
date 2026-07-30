# GUI: Security -> Bandwidth control -> Per IP profile
resource "fortisase_security_per_ip_traffic_shaper" "example_per_ip" {
  primary_key                 = "example_per_ip"
  bandwidth_unit              = "kbps"
  maximum_bandwidth           = 1024
  max_concurrent_sessions     = 330
  max_concurrent_tcp_sessions = 100
  max_concurrent_udp_sessions = 100
}
