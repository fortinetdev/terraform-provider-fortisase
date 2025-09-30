resource "fortisase_private_access_network_configuration" "example" {
  bgp_design            = "loopback"
  bgp_router_ids_subnet = "172.1.0.0/24"
  as_number             = "65400"
  sdwan_rule_enable     = "true"
  sdwan_health_check_vm = "10.255.255.100"
  recursive_next_hop    = "true"
}

resource "fortisase_private_access_service_connections" "example" {
  type                 = fortisase_private_access_network_configuration.example.bgp_design
  alias                = "AWS-Ireland-Primary"
  ipsec_remote_gw      = "1.1.1.1"
  ipsec_ike_version    = "2"
  auth                 = "psk"
  ipsec_pre_shared_key = "example_shared_key"
  route_map_tag        = "100"
  bgp_peer_ip          = "10.255.255.100"
  overlay_network_id   = "100"
}
