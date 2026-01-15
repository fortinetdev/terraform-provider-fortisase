resource "fortisase_security_profile_group" "example" {
  primary_key = "example_profile_name" # The name of the new profile group

  # AntiVirus
  antivirus_profile = {
    status = "enable"
  }

  # Web Filter With Inline-CASB
  web_filter_profile = {
    status = "enable"
  }

  # Intrusion Prevention
  intrusion_prevention_profile = {
    status = "enable"
  }

  # File Filter
  file_filter_profile = {
    status = "disable"
  }

  # Data Loss Prevention (DLP)
  dlp_filter_profile = {
    status = "disable"
  }

  # DNS Filter
  dns_filter_profile = {
    status = "enable"
  }

  # Application Control With Inline-CASB
  application_control_profile = {
    status = "disable"
  }

  # Video Filter
  video_filter_profile = {
    status = "disable"
  }
}
