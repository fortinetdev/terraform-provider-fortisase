data "fortisase_security_file_filter_profile" "example" {
  direction   = "outbound-profiles" # "internal-profiles" or "outbound-profiles"
  primary_key = "default"           # The name of your existing profile group
}
