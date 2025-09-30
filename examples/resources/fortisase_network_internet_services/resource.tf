resource "fortisase_network_internet_services" "services" {
  primary_key     = "example_name"
  direction       = "dst"
  ip_range_number = 34948
  ip_number       = 19597162
  icon_id         = 1
}
