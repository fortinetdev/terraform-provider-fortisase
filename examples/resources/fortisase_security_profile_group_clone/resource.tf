# Note: Destroy this resource won't delete the cloned profile group.
resource "fortisase_security_profile_group_clone" "example" {
  direction   = "outbound-profiles"     # "internal-profiles" or "outbound-profiles"
  based_on    = "default"               # The profile to clone from
  primary_key = "your-new-profile-name" # The name of the new profile group
}
