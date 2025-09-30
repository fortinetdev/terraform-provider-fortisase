# Note: Destroy this resource won't delete the cloned endpoint profile.
resource "fortisase_endpoint_policies_clone" "clone" {
  based_on    = "Default"        # The profile to clone from
  primary_key = "newProfileName" # The name of the new profile
  enabled     = true
}
