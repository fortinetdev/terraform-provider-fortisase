# GUI: Security > Policies > Secure private access > To hubs
resource "fortisase_security_internal_policy" "example" {
  primary_key = "internal_policy_name"

  # [Source Scope]
  scope = "vpn-user"

  # [Security Posture Tag]
  # Option 1: Reference a tagging rule.
  sources = [
    {
      primary_key = fortisase_endpoint_ztna_tag_rule.compliant.primary_key
      datasource  = "endpoint/ztna-tag-rules" # EMS 7.4: use fortisase_endpoint_ztna_tag_rule; EMS 7.2: use "endpoint/ztna-tags"
    }
  ]
  # Option 2: No security posture tag 
  # sources = []

  # [User]
  users = [
    {
      primary_key = fortisase_auth_user.private_app.primary_key
      datasource  = "auth/users"
    }
  ]

  # [Destination]
  destinations = [
    {
      primary_key = fortisase_network_host.private_app.primary_key
      datasource  = "network/hosts"
    }
  ]

  # [Service]
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

  # [Profile Group]
  profile_group = {
    group = {
      primary_key = fortisase_security_profile_group.private_app.primary_key
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
  comments = "Access for private application"
}

resource "fortisase_endpoint_ztna_tag_rule" "compliant" {
  primary_key = "Compliant"
  status      = "enable"
  description = "Windows endpoints with antivirus software installed and running"
  rules = [{
    id      = 1
    os      = "windows"
    type    = "anti-virus"
    content = "AV Software is installed and running"
    negated = false
  }]
  logic = {
    windows = jsonencode({
      op    = "and"
      rules = [{ id = 1 }]
    })
  }
}

resource "fortisase_auth_user" "private_app" {
  primary_key = "private-app-user@example.com"
  auth_type   = "password"
  status      = "enable"
  email       = "private-app-user@example.com"
  password    = "Example-PrivateApp-Password123!" # Replace with the user's initial password before applying.
}

resource "fortisase_network_host" "private_app" {
  primary_key = "Private Application"
  type        = "ipmask"
  location    = "private-access"
  subnet      = "192.168.4.10/32" # Replace with the private application's IP address.
}

resource "fortisase_security_profile_group" "private_app" {
  primary_key = "tf-private-app-profile"
}
