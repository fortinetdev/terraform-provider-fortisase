resource "fortisase_security_profile_group" "example" {
  direction   = "outbound-profiles" # "internal-profiles" or "outbound-profiles"
  primary_key = "TF1"               # the name of the new profile group
}
