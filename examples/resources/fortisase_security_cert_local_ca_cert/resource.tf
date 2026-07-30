# Method 1: fortisase_system_certificate is recommended for new configurations.
resource "fortisase_system_certificate" "ca_certificate" {
  certificate_type = "ca-certificate"
  primary_key      = "example_ca_cert"
  format           = "regular" # "regular" or "pkcs12"
  file_content     = base64encode(file("../path/to/my-ca-cert.pem"))
  key_file_content = base64encode(file("../path/to/my-ca-key.pem")) # Only required if format is "regular"
}

# Method 2
resource "fortisase_security_cert_local_ca_cert" "local_ca_cert" {
  cert_name        = "local_ca_cert"
  format           = "regular"
  password         = "your_password"
  file_content     = base64encode(file("./path/to/ca_cert.crt"))
  key_file_content = base64encode(file("./path/to/private.key"))
}
