# GUI: Security > Policies > Secure private access > From hubs
resource "fortisase_security_internal_reverse_policies" "example" {
  primary_key = "IT Admins: Remote Access by TF"
  enabled     = true
  scope       = "vpn-user"

  services = [
    {
      primary_key = "SSH"
      datasource  = "security/services"
    },
    {
      primary_key = "RDP"
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
      primary_key = "IT Admin Subnet"
      datasource  = "network/hosts"
    }
  ]

  schedule = {
    primary_key = "always"
    datasource  = "security/recurring-schedules"
  }

  comments = "Allow IT Admins remote access to machines of mobile workers"
}
