# Method 1: fortisase_system_certificate is recommended for new configurations.
resource "fortisase_system_certificate" "local_certificate_regular" {
  certificate_type = "local-certificate"
  primary_key      = "example_local_cert_regular"
  format           = "regular"
  file_content     = filebase64("./path/to/local-cert-file.crt")
  key_file_content = filebase64("./path/to/local-cert-file-key.pem")
  password         = "example_password"
}

resource "fortisase_system_certificate" "local_certificate_pkcs12" {
  certificate_type = "local-certificate"
  primary_key      = "example_local_cert_pkcs12"
  format           = "pkcs12"
  file_content     = filebase64("./path/to/local-cert-file.p12")
  password         = "example_password"
}

# Method 2
resource "fortisase_security_cert_local_cert" "local_cert" {
  cert_name        = "local_cert_name"
  format           = "regular"
  password         = "your_password"
  file_content     = filebase64("./path/to/local-cert-file2.crt")
  key_file_content = filebase64("./path/to/local-cert-file-key2.pem")
}
