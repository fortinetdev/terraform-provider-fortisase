resource "fortisase_security_profile_group" "example" {
  direction   = "outbound-profiles"    # "internal-profiles" or "outbound-profiles"
  primary_key = "example_profile_name" # The name of the new profile group
}

# To configure this resource, please disable proxy configuration.
resource "fortisase_security_application_control_profile" "application_control_profile" {
  direction                           = fortisase_security_profile_group.example.direction   # "internal-profiles" or "outbound-profiles"
  primary_key                         = fortisase_security_profile_group.example.primary_key # The name of the existing profile group
  application_controls                = []
  unknown_application_action          = "allow"
  network_protocol_enforcement        = "disable"
  network_protocols                   = []
  block_non_default_port_applications = "disable"
}
