resource "fortisase_security_profile_group" "example" {
  direction   = "outbound-profiles"    # "internal-profiles" or "outbound-profiles"
  primary_key = "example_profile_name" # The name of the new profile group
}

resource "fortisase_security_file_filter_profile" "file_filter_profile" {
  direction                      = fortisase_security_profile_group.example.direction   # "internal-profiles" or "outbound-profiles"
  primary_key                    = fortisase_security_profile_group.example.primary_key # The name of the existing profile group
  block                          = []
  monitor                        = []
  block_password_protected_files = false
}
