# GUI: Security -> Proxy policies
resource "fortisase_security_internal_proxy_policy" "example" {
  primary_key = "internal_proxy_name"

  # [User]
  users = [
    # {
    #   primary_key = "existing-user@example.com"
    #   datasource  = "auth/users"
    # }
  ]

  # [Destination]
  destinations = [
    {
      primary_key = "all"
      datasource  = "network/hosts"
    }
  ]

  # [Security Posture Tag]
  sources = [
    # {
    #   primary_key = "existing_tag_rule_name"
    #   datasource  = "endpoint/ztna-tag-rules" # EMS 7.4; for EMS 7.2, use "endpoint/ztna-tags" and the tag's primary key.
    # }
  ]

  # [Profile Group]
  profile_group = {
    group = {
      primary_key = fortisase_security_profile_group.internal_proxy.primary_key
      datasource  = "security/profile-groups"
    }
    force_cert_inspection = false
  }

  # [Schedule]
  schedule = {
    primary_key = "always"
    datasource  = "security/recurring-schedules"
  }

  # [Action]
  action = "accept"

  # [Status]
  enabled = true

  # [Log Violation Traffic]
  log_traffic = "all"

  # [Comments]
  comments = "tf test internal proxy policy"
}

resource "fortisase_security_profile_group" "internal_proxy" {
  primary_key = "tf-internal-proxy-profile"
}
