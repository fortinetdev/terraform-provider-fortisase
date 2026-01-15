resource "fortisase_security_profile_group" "example" {
  primary_key = "example_profile_name" # The name of the new profile group
}

# Please use one of the following examples:
# 1. Certificate inspection
resource "fortisase_security_ssl_ssh_profile" "certificate_inspection" {
  primary_key = fortisase_security_profile_group.example.primary_key

  # Inspection method
  inspection_mode = "certificate-inspection"

  # CA certificate
  ca_certificate = {
    primary_key = "Fortinet_CA_SSL"
    datasource  = "system/certificate/ca-certificates"
  }
  # Handling invalid certificates
  expired_certificate_action              = "block"   # "allow", "block"
  revoked_certificate_action              = "block"   # "allow", "block"
  timed_out_validation_certificate_action = "allow"   # "allow", "block"
  validation_failed_certificate_action    = "block"   # "allow", "block"
  cert_probe_failure                      = "allow"   # "allow", "block"
  quic                                    = "inspect" # "inspect", "bypass", "block"

  # Performance optimization
  profile_protocol_options = {
    unknown_content_encoding = "inspect" # "block", "inspect", "bypass"
    oversized_action         = "block"   # "allow", "block"
    uncompressed_limit       = 10
    compressed_limit         = 10
  }
}

# 2. Deep inspection
resource "fortisase_security_ssl_ssh_profile" "deep_inspection" {
  primary_key = fortisase_security_profile_group.example.primary_key

  # Inspection method
  inspection_mode = "deep-inspection"

  # CA certificate
  ca_certificate = {
    primary_key = "Fortinet_CA_SSL"
    datasource  = "system/certificate/ca-certificates"
  }
  # Handling invalid certificates
  expired_certificate_action              = "block"   # "allow", "block"
  revoked_certificate_action              = "block"   # "allow", "block"
  timed_out_validation_certificate_action = "allow"   # "allow", "block"
  validation_failed_certificate_action    = "block"   # "allow", "block"
  cert_probe_failure                      = "allow"   # "allow", "block"
  quic                                    = "inspect" # "inspect", "bypass", "block"

  # Exemptions
  host_exemptions = [
    {
      primary_key = "FortiClient"
      datasource  = "network/hosts"
    },
    {
      primary_key = "Fortinet Services"
      datasource  = "network/host-groups"
    }
  ]
  url_category_exemptions = [
    {
      primary_key = "Finance and Banking"
      datasource  = "security/fortiguard-categories"
    },
    {
      primary_key = "Health and Wellness"
      datasource  = "security/fortiguard-categories"
    }
  ]

  # Performance optimization
  profile_protocol_options = {
    unknown_content_encoding = "inspect" # "block", "inspect", "bypass"
    oversized_action         = "block"   # "allow", "block"
    uncompressed_limit       = 10
    compressed_limit         = 10
  }
}

# 3. No inspection
resource "fortisase_security_ssl_ssh_profile" "no_inspection" {
  primary_key = fortisase_security_profile_group.example.primary_key

  # Inspection method
  inspection_mode = "no-inspection"

  # Performance optimization
  profile_protocol_options = {
    unknown_content_encoding = "inspect" # "block", "inspect", "bypass"
    oversized_action         = "block"   # "allow", "block"
    uncompressed_limit       = 10
    compressed_limit         = 10
  }
}
