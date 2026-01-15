resource "fortisase_security_profile_group" "example" {
  primary_key = "example_profile_name" # The name of the new profile group

  # Web Filter With Inline-CASB
  web_filter_profile = {
    status = "enable"
  }
}

resource "fortisase_security_web_filter_profile" "web_filter_profile" {
  primary_key = fortisase_security_profile_group.example.primary_key # The name of the existing profile group

  # FortiGuard Category Based Filter
  use_fortiguard_filters = "enable"

  # FortiGuard categories
  fortiguard_filters = [{
    action = "block" # "allow" or "monitor" or "block" or "warning"
    category = {
      primary_key = "Drug Abuse"
      datasource  = "security/fortiguard-categories"
    }
  }]

  # FortiGuard custom categories
  # fortiguard_local_category_filters = [
  #   {
  #     action = "disable" # "allow" or "monitor" or "block" or "warning" or "disable"
  #     category = {
  #       datasource  = "security/fortiguard-local-categories"
  #       primary_key = "custom1"
  #     }
  #   }
  # ]

  # URL Filter
  url_filters = [
    {
      action = "block"
      status = "enable"
      type   = "simple"
      url    = "your_string"
    }
  ]
  # Content Filter
  content_filters = [
    {
      action       = "exempt"
      lang         = "western"
      pattern      = "your_string"
      pattern_type = "wildcard"
      score        = 10
      status       = "disable"
    },
  ]

  # Options
  block_invalid_url       = "disable"
  traffic_on_rating_error = "disable"
  enforce_safe_search     = "disable"
  log_searched_keywords   = "disable"

  # Inline-CASB Headers
  http_headers = [{
    name    = "example_header_name"
    action  = "add-to-request"
    content = "example_header_content"
    destinations = [{
      primary_key = "existing_host_example"
      datasource  = "network/hosts"
    }]
  }]
}
