# GUI: Security -> Proxy policies
resource "fortisase_security_outbound_proxy_policy" "example" {
  primary_key = "outbound_proxy_name"
  enabled     = true
  action      = "accept"
  log_traffic = "all"
  comments    = "tf test outbound proxy policy"

  users = []

  destinations = [
    {
      primary_key = "all"
      datasource  = "network/hosts"
    }
  ]

  sources = []

  profile_group = {
    group = {
      primary_key = "outbound"
      datasource  = "security/profile-groups"
    }
    force_cert_inspection = false
  }

  schedule = {
    primary_key = "always"
    datasource  = "security/recurring-schedules"
  }
}
