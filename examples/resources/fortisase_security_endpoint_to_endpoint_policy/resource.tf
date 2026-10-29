# GUI: Security > Policies > Secure private access > Same PoP client-to-client
resource "fortisase_security_endpoint_to_endpoint_policy" "example" {
  primary_key = "example"

  # [Security Posture Tag]
  sources = [
    # {
    #   primary_key = "existing_ztna_tag"
    #   datasource  = "endpoint/ztna-tags"
    # }
  ]

  # [Service]
  services = [
    {
      primary_key = "ALL"
      datasource  = "security/services"
    }
  ]

  # [Schedule]
  schedule = {
    primary_key = "always"
    datasource  = "security/recurring-schedules"
  }

  # [Action]
  action = "accept"

  # [Profile Group] Required if action = "accept"
  profile_group = {
    group = {
      primary_key = fortisase_security_profile_group.endpoint_to_endpoint.primary_key
      datasource  = "security/profile-groups"
    }
    force_cert_inspection = true
  }

  # [Status]
  enabled = true

  # [Log Violation Traffic]
  log_traffic = "all"

  # [Comments]
  comments = "Your comment"
}

resource "fortisase_security_profile_group" "endpoint_to_endpoint" {
  primary_key = "endpoint-to-endpoint-profile"
}
