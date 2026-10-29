# GUI: Security > Policies > Secure private access > From hubs
resource "fortisase_security_internal_reverse_policy" "example" {
  primary_key = "IT Admins: Remote Access by TF"
  # [Destination Scope]
  scope = "vpn-user"

  # [Source]
  sources = [
    {
      primary_key = fortisase_network_host.admin_subnet.primary_key
      datasource  = "network/hosts"
    }
  ]

  # [Service]
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

  # [Profile Group]
  profile_group = {
    group = {
      primary_key = fortisase_security_profile_group.internal_reverse.primary_key
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
  comments = "Allow IT Admins remote access to machines of mobile workers"
}

resource "fortisase_network_host" "admin_subnet" {
  primary_key = "IT Admin Subnet"
  type        = "ipmask"
  location    = "private-access"
  subnet      = "192.168.4.0/24" # Replace with the administrator subnet behind the SPA hub.
}

resource "fortisase_security_profile_group" "internal_reverse" {
  primary_key = "tf-internal-reverse-profile"
}
