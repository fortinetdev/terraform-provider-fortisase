# Note: Destroying this resource won't delete the cloned profile group.
resource "fortisase_security_profile_group_clone" "template" {
  based_on    = "existing_profile_name" # The profile to clone from
  primary_key = "new_profile_name"      # The name of the new profile group
}

# Clone Default Internet Access
resource "fortisase_security_profile_group_clone" "clone_default_outbound_profile" {
  based_on    = "outbound"
  primary_key = "cloned_outbound_profile"
}

# Clone Default Private Access
resource "fortisase_security_profile_group_clone" "clone_default_internal_profile" {
  based_on    = "internal"
  primary_key = "cloned_internal_profile"
}
