resource "fortisase_network_host" "network_host" {
  primary_key = "network_host_example"
  type        = "ipmask"
  location    = "internal"
  subnet      = "192.168.4.0/24"
}
