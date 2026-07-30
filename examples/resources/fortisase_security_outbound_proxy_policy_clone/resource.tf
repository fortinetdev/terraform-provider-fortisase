# Note: Destroy this resource won't delete the cloned policy.
resource "fortisase_security_outbound_proxy_policy_clone" "clone_example" {
  based_on    = "existing_proxy_policy_name" # The policy to clone from
  primary_key = "new_proxy_policy_name"      # The name of the new policy
}
