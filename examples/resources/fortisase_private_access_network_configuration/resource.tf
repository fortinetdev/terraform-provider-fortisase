resource "fortisase_private_access_network_configuration" "example" {
  bgp_design            = "loopback"
  bgp_router_ids_subnet = "172.1.0.0/24"
  as_number             = "65400"
  sdwan_rule_enable     = "true"
  sdwan_health_check_vm = "10.255.255.100"
  recursive_next_hop    = "true"
}
