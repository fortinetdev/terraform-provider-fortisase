resource "fortisase_security_profile_group" "example" {
  direction   = "outbound-profiles"    # "internal-profiles" or "outbound-profiles"
  primary_key = "example_profile_name" # The name of the new profile group
}

resource "fortisase_security_ssl_ssh_profile" "ssl_ssh_profile" {
  direction                               = fortisase_security_profile_group.example.direction   # "internal-profiles" or "outbound-profiles"
  primary_key                             = fortisase_security_profile_group.example.primary_key # The name of the existing profile group
  inspection_mode                         = "certificate-inspection"
  expired_certificate_action              = "block"
  revoked_certificate_action              = "block"
  timed_out_validation_certificate_action = "allow"
  validation_failed_certificate_action    = "block"
  cert_probe_failure                      = "allow"
  ca_certificate = {
    primary_key = "Fortinet_CA_SSL"
    datasource  = "system/certificate/ca-certificates"
  }
  host_exemptions = [
    {
      primary_key = "FortiClient"
      datasource  = "network/hosts"
    },
    {
      primary_key = "Fortinet Services"
      datasource  = "network/host-groups"
    }
  ]
  url_category_exemptions = [
    {
      primary_key = "Finance and Banking"
      datasource  = "security/fortiguard-categories"
    },
    {
      primary_key = "Health and Wellness"
      datasource  = "security/fortiguard-categories"
    }
  ]
}
