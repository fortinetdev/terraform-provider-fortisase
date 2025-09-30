resource "fortisase_security_dlp_file_patterns" "dlp_file_pattern" {
  # Please don't set "primary_key", the value of this variable is only known after the resource is created.
  tag = "test"
  entries = [
    {
      pattern     = "string"
      filter_type = "type"
      file_type   = "7z"
    }
  ]
}
