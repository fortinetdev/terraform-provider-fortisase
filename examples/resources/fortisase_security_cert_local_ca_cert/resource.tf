# Method 1: fortisase_system_certificate is recommended for new configurations.
resource "fortisase_system_certificate" "ca_certificate_regular" {
  certificate_type = "ca-certificate"
  primary_key      = "example_ca_cert_regular"
  format           = "regular"
  file_content     = filebase64("./path/to/ca-cert-file.crt")
  key_file_content = filebase64("./path/to/ca-cert-file-key.pem")
  password         = "example_password"
}

resource "fortisase_system_certificate" "ca_certificate_pkcs12" {
  certificate_type = "ca-certificate"
  primary_key      = "example_ca_cert_pkcs12"
  format           = "pkcs12"
  file_content     = filebase64("./path/to/ca-cert-file.p12")
  password         = "example_password"
}

# Method 2
resource "fortisase_security_cert_local_ca_cert" "local_ca_cert" {
  cert_name        = "local_ca_cert"
  format           = "regular"
  password         = "your_password"
  file_content     = filebase64("./path/to/ca-cert-file2.crt")
  key_file_content = filebase64("./path/to/ca-cert-file-key2.pem")
}
