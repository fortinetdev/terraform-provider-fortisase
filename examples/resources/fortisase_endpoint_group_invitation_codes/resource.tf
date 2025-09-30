resource "fortisase_endpoint_group_invitation_codes" "example" {
  primary_key = "example_name"
  expire_date = "2026-03-08T12:45:30Z"
  group_assignment = {
    enabled = true
    group = {
      id   = "1"
      path = "All Groups"
    }
  }
}
