# Method 1: fortisase_system_certificate is recommended for new configurations.
resource "fortisase_system_certificate" "local_certificate" {
  certificate_type = "local-certificate"
  primary_key      = "example_local_cert"
  format           = "regular" # "regular" or "pkcs12"
  file_content     = base64encode(file("../path/to/my-server-cert.pem"))
  key_file_content = base64encode(file("../path/to/my-server-key.pem")) # Only required if format is "regular"
}

# Method 2
resource "fortisase_security_cert_local_cert" "local_cert" {
  cert_name        = "local_cert_name"
  format           = "regular"
  password         = "your_password"
  file_content     = base64encode(file("./path/to/cert.pem"))
  key_file_content = base64encode(file("./path/to/key.pem"))
}
