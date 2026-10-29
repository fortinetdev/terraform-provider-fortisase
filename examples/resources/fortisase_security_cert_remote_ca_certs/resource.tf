# Method 1: fortisase_system_certificate is recommended for new configurations.
resource "fortisase_system_certificate" "remote_ca_certificate" {
  certificate_type = "remote-ca-certificate"
  primary_key      = "example_remote_ca_cert"
  file_content     = filebase64("./path/to/remote-ca-file.crt")
}

# Method 2
resource "fortisase_security_cert_remote_ca_cert" "remote_ca_cert" {
  cert_name    = "remote_ca_cert_name"
  file_content = filebase64("./path/to/remote-ca-file2.crt")
}
