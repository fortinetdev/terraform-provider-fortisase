resource "fortisase_endpoint_policies" "endpoint_profile" {
  primary_key = "example"
  enabled     = true
}

resource "fortisase_endpoint_protection_profiles" "endpoint_protection_profile" {
  primary_key = fortisase_endpoint_policies.endpoint_profile.primary_key

  # Malware
  antivirus      = "enable"
  antivirus_scan = "enable"    # Required if antivirus is "enable"
  scheduled_antivirus_scan = { # Required if antivirus_scan is "enable"
    scan_type = "full"         # "full" or "quick"
    repeat    = "daily"        # "daily", "weekly", "monthly"
    time      = "00:00"        # from "00:00" to "23:59"
    # day     = 1              # 1 ~ 7 if repeat is "weekly", 1 ~ 31 if repeat is "monthly"
  }

  # Anti-Ransomware, this feature applies to Windows endpoints only
  antiransomware = "enable"
  protected_folders_path = [ # Required if antiransomware is "enable"
    "%USERPROFILE%\\Documents\\", "%USERPROFILE%\\Pictures\\", "%USERPROFILE%\\Videos\\",
    "%USERPROFILE%\\Music\\", "%USERPROFILE%\\Desktop\\", "%USERPROFILE%\\Favorites\\"
  ]

  # Scan for vulnerabilities
  vulnerability_scan = "enable"
  scheduled_scan = { # Required if vulnerability_scan is "enable"
    repeat = "daily" # "daily", "weekly", "monthly"
    time   = "00:00" # from "00:00" to "23:59"
    # day  = 1       # 1 ~ 7 if repeat is "weekly", 1 ~ 31 if repeat is "monthly"
  }

  # Event-based scanning
  event_based_scanning = "enable"

  # Automatically patch vulnerabilities
  automatically_patch_vulnerabilities = "enable"
  automatic_vulnerability_patch_level = "medium" # Required if automatically_patch_vulnerabilities is "enable"

  # Exclude specified folders/files
  exclusions = {
    files   = []
    folders = []
  }

  # Default removable media access
  default_action = "allow"

  # Notify endpoint of blocks
  notify_endpoint_of_blocks = "enable"

  # Access control rules
  rules = [{
    action       = "allow"
    class        = "Bluetooth"
    description  = "example"
    manufacturer = "test"
    product_id   = "123"
    revision     = "123"
    type         = "simple"
    vendor_id    = "123"
  }]
}
