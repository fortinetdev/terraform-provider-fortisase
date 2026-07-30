# Method 1: fortisase_system_certificate is recommended for new configurations.
resource "fortisase_system_certificate" "remote_certificate" {
  certificate_type = "remote-certificate"
  primary_key      = "example_remote_cert"
  file_content     = base64encode(file("../path/to/my-server-cert.pem"))
}

# Method 2
resource "fortisase_security_cert_remote_cert" "remote_cert" {
  cert_name    = "remote_cert_name"
  file_content = base64encode(file("./path/to/cert.pem"))
}
