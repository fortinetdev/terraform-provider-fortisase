resource "fortisase_security_dlp_exact_data_matches" "example" {
  primary_key = "example_name"

  external_resource_data = {
    resource      = "https://example-resource.com"
    username      = "admin"
    password      = "password123"
    refresh_rate  = 3600
    update_method = "feed"
  }
  optional_count = 0

  columns = [{
    index    = 1
    optional = false
    type = {
      datasource  = "security/dlp-data-types"
      primary_key = "credit-card"
    }
  }]
}
