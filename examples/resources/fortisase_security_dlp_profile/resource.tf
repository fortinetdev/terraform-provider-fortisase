resource "fortisase_security_profile_group" "example" {
  primary_key = "example_profile_name" # The name of the new profile group

  # Data Loss Prevention (DLP)
  dlp_filter_profile = {
    status = "enable"
  }
}

resource "fortisase_security_dlp_profile" "dlp_profile" {
  primary_key = fortisase_security_profile_group.example.primary_key # The name of the existing profile group
  dlp_rules   = []
}
