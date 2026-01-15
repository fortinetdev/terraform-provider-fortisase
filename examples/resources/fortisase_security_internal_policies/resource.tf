# GUI: Security > Policies > Secure private access > To hubs
resource "fortisase_security_internal_policies" "example" {
  primary_key = "Secure Private Access: Finance App"
  enabled     = true
  scope       = "vpn-user"
  users = [
    {
      primary_key = "EntraID_Finance"
      datasource  = "auth/user-groups"
    }
  ]
  destinations = [
    {
      primary_key = "Finance Application"
      datasource  = "network/hosts"
    }
  ]
  services = [
    {
      primary_key = "HTTPS"
      datasource  = "security/services"
    },
    {
      primary_key = "HTTP"
      datasource  = "security/services"
    }
  ]
  action      = "accept"
  log_traffic = "all"

  profile_group = {
    group = {
      primary_key = "internal"
      datasource  = "security/profile-groups"
    }
    force_cert_inspection = false
  }

  sources = [
    {
      primary_key = "Compliant"
      datasource  = "endpoint/ztna-tags"
    }
  ]

  schedule = {
    primary_key = "always"
    datasource  = "security/recurring-schedules"
  }

  comments = "Access for private Finance App"

}
