resource "fortisase_security_dlp_fingerprint_databases" "dlp_fingerprint_databases" {
  primary_key = "dlp_fingerprint_databases"

  server                 = "example-server.com"
  sensitivity            = "Warning"
  include_subdirectories = "enable"
  server_directory       = "/path/to/directory/"
  file_pattern           = "*.txt"
  schedule = {
    period      = "daily"
    sync_hour   = 2
    sync_minute = 0
  }
  remove_deleted_file_fingerprints = "enable"
  keep_modified                    = "enable"
  scan_on_creation                 = "enable"
  authentication = {
    username = "admin"
    password = "password123"
  }
}
