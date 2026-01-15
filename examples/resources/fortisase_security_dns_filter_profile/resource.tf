resource "fortisase_security_profile_group" "example" {
  primary_key = "example_profile_name" # The name of the new profile group

  # DNS Filter
  dns_filter_profile = {
    status = "enable"
  }
}

resource "fortisase_security_dns_filter_profile" "dns_filter_profile" {
  primary_key                        = fortisase_security_profile_group.example.primary_key # The name of the existing profile group
  use_fortiguard_filters             = "enable"
  enable_all_logs                    = "disable"
  enable_botnet_blocking             = "enable"
  enable_safe_search                 = "disable"
  allow_dns_requests_on_rating_error = "enable"
  dns_translation_entries            = []
  domain_filters                     = []
  domain_threat_feed_filters         = []
  use_for_edge_devices               = false
}
