resource "fortisase_endpoint_policies" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}

resource "fortisase_endpoint_sandbox_profiles" "endpoint_sandbox_profile" {
  primary_key                      = fortisase_endpoint_policies.endpoint_profile.primary_key
  sandbox_mode                     = "FortiSASE"
  notification_type                = 1
  detection_verdict_level          = "Medium"
  timeout_awaiting_sandbox_results = 300
  file_submission_options = {
    all_files_removable_media       = "enable"
    all_files_mapped_network_drives = "enable"
    all_web_downloads               = "enable"
    all_email_downloads             = "enable"

  }
  remediation_actions = "quarantine"
  exceptions = {
    exclude_files_from_trusted_sources = "disable"
    files                              = []
    folders                            = ["/", "/123"]
  }
}
