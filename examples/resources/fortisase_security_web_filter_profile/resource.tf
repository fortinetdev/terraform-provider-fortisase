resource "fortisase_security_profile_group" "example" {
  direction   = "outbound-profiles"    # "internal-profiles" or "outbound-profiles"
  primary_key = "example_profile_name" # The name of the new profile group
}

resource "fortisase_security_web_filter_profile" "web_filter_profile" {
  direction               = fortisase_security_profile_group.example.direction   # "internal-profiles" or "outbound-profiles"
  primary_key             = fortisase_security_profile_group.example.primary_key # The name of the existing profile group
  use_fortiguard_filters  = "disable"
  block_invalid_url       = "disable"
  enforce_safe_search     = "disable"
  traffic_on_rating_error = "enable"
  content_filters = [
    {
      action       = "exempt"
      lang         = "western"
      pattern      = "your_string"
      pattern_type = "wildcard"
      score        = 10
      status       = "enable"
    },
  ]
  http_headers = []
  url_filters = [
    {
      action = "block"
      status = "enable"
      type   = "simple"
      url    = "your_string"
    },
  ]
}
