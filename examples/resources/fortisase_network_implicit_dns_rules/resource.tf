resource "fortisase_network_implicit_dns_rules" "network_implicit_dns_rule" {
  primary_key = "implicit_all"
  dns_server  = "google"
}
