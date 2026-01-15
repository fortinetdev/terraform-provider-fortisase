resource "fortisase_security_profile_group" "example" {
  primary_key = "example_profile_name" # The name of the new profile group

  # AntiVirus
  antivirus_profile = {
    status = "enable"
  }
}

resource "fortisase_security_antivirus_profile" "antivirus_profile" {
  primary_key = fortisase_security_profile_group.example.primary_key # The name of the existing profile group
  http        = "enable"
  smtp        = "enable"
  pop3        = "enable"
  imap        = "enable"
  ftp         = "enable"
  cifs        = "enable"
}
