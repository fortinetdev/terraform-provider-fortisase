# GUI: Security > Policies > Secure private access > Same PoP client-to-client
resource "fortisase_security_endpoint_to_endpoint_policies" "example" {
  primary_key = "example"
  enabled     = true
  services = [
    {
      primary_key = "ALL"
      datasource  = "security/services"
    }
  ]
  action      = "accept"
  log_traffic = "all"
  profile_group = {
    group = {
      primary_key = "existing_profile_group"
      datasource  = "security/profile-groups"
    }
    force_cert_inspection = true
  }
  sources = [
    {
      primary_key = "existing_tag"
      datasource  = "endpoint/ztna-tags"
    }
  ]
  schedule = {
    primary_key = "always"
    datasource  = "security/recurring-schedules"
  }
  comments = "Your comment"
}
