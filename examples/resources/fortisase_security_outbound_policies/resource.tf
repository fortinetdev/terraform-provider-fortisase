resource "fortisase_security_outbound_policies" "outbound_policy" {
  primary_key = "Secure SaaS Access"
  enabled     = true
  scope       = "vpn-user"
  users = [
    {
      primary_key = "EntraID_Finance"
      datasource  = "auth/user-groups"
    },
    {
      primary_key = "EntraID_Sales"
      datasource  = "auth/user-groups"
    }
  ]
  destinations = [
    {
      primary_key = "Salesforce-Web"
      datasource  = "network/internet-services"
    },
    {
      primary_key = "SAP-Web"
      datasource  = "network/internet-services"
    }
  ]
  # services is required by the backend API. It can be set to "ALL" for outbound policies.
  services = [
    {
      primary_key = "ALL",
      datasource  = "security/services"
    }
  ]
  action      = "accept"
  log_traffic = "all"
  profile_group = {
    group = {
      primary_key = "Secure SaaS Access"
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
  comments = "Secure SaaS Access Policy"
}
