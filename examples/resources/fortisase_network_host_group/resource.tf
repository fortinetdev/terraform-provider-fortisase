resource "fortisase_network_host_group" "network_host_group" {
  primary_key = "network_host_group"
  members = [
    {
      datasource  = "network/hosts"
      primary_key = "all"
    }
  ]
}
