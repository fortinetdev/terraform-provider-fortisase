resource "fortisase_security_profile_group" "example" {
  primary_key = "example_profile_name" # The name of the new profile group

  # Video Filter
  video_filter_profile = {
    status = "enable"
  }
}

resource "fortisase_security_video_filter_profile" "video_filter_profile" {
  primary_key    = fortisase_security_profile_group.example.primary_key # The name of the existing profile group
  default_action = "monitor"
  channels       = []
}
