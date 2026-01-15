resource "fortisase_security_profile_group" "example" {
  primary_key = "example_profile_name" # The name of the new profile group

  # File Filter
  file_filter_profile = {
    status = "enable"
  }
}

resource "fortisase_security_file_filter_profile" "file_filter_profile" {
  primary_key = fortisase_security_profile_group.example.primary_key # The name of the existing profile group

  # The primary_key not in both block and monitor is considered as "allow"
  block = [
    {
      primary_key = "upx"
      datasource  = "security/antivirus-filetypes"
    },
    {
      primary_key = "torrent"
      datasource  = "security/antivirus-filetypes"
    },
  ]
  monitor = [
    {
      primary_key = "exe"
      datasource  = "security/antivirus-filetypes"
    }
  ]
  block_password_protected_files = false
}
