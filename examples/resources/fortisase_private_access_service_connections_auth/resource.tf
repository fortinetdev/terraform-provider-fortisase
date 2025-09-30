resource "fortisase_private_access_network_configuration" "example" {
  bgp_design            = "loopback"
  bgp_router_ids_subnet = "172.1.0.0/24"
  as_number             = "65400"
  sdwan_rule_enable     = "true"
  sdwan_health_check_vm = "10.255.255.100"
  recursive_next_hop    = "true"
}

resource "fortisase_private_access_service_connections" "example" {
  type                  = fortisase_private_access_network_configuration.example.bgp_design
  service_connection_id = "1"
  alias                 = "AWS-Ireland-Primary"
  ipsec_remote_gw       = "1.1.1.1"
  ipsec_ike_version     = "2"
  auth                  = "psk"
  ipsec_pre_shared_key  = "example_shared_key"
  route_map_tag         = "100"
  bgp_peer_ip           = "10.255.255.100"
  overlay_network_id    = "100"
}

resource "fortisase_private_access_service_connections_auth" "example" {
  service_connection_id = fortisase_private_access_service_connections.example.id
  auth                  = "pki"
  ipsec_pre_shared_key  = fortisase_private_access_service_connections.example.ipsec_pre_shared_key
  ipsec_cert_name       = "certificate_local"
  ipsec_peer_name       = "HUB_CERT_1"
}
