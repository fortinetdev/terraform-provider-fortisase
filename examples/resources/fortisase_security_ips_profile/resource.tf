resource "fortisase_security_profile_group" "example" {
  direction   = "outbound-profiles"    # "internal-profiles" or "outbound-profiles"
  primary_key = "example_profile_name" # The name of the new profile group
}


resource "fortisase_security_ips_profile" "ips_profile" {
  direction                 = fortisase_security_profile_group.example.direction   # "internal-profiles" or "outbound-profiles"
  primary_key               = fortisase_security_profile_group.example.primary_key # The name of the existing profile group
  profile_type              = "recommended"
  custom_rule_groups        = []
  is_blocking_malicious_url = false
  botnet_scanning           = "block"
  is_extended_log_enabled   = false
  comment                   = "Recommended"
  entries = [
    {
      rule               = [],
      location           = "all",
      severity           = "all",
      protocol           = "all",
      os                 = "all",
      application        = "all",
      cve                = [],
      status             = "default",
      log                = "enable",
      log_packet         = "disable",
      log_attack_context = "disable",
      action             = "default",
      quarantine         = "none",
      exempt_ip          = [],
      vuln_type          = [],
      default_action     = "all",
      default_status     = "all"
    }
  ]
}
