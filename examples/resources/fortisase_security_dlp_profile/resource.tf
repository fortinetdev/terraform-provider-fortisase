resource "fortisase_security_profile_group" "example" {
  primary_key = "example_profile_name" # The name of the new profile group

  # Data Loss Prevention (DLP)
  dlp_filter_profile = {
    status = "enable"
  }
}

resource "fortisase_security_dlp_profile" "dlp_profile" {
  primary_key = fortisase_security_profile_group.example.primary_key # The name of the existing profile group
  dlp_rules = [
    {
      primary_key     = "tf-monitor-all-files"
      datasource_type = "none"
      severity        = "medium"
      action          = "monitor"
      dlp_rule_type   = "file"
      file_type       = "specify"
      dlp_file_pattern = {
        primary_key = fortisase_security_dlp_file_pattern.all_files.primary_key
        datasource  = "security/dlp-file-patterns"
      }
      protocols = ["http-get"]
    }
  ]
}

resource "fortisase_security_dlp_file_pattern" "all_files" {
  tag = "tf-all-files"
  entries = [{
    filter_type = "pattern"
    pattern     = "*"
  }]
}
