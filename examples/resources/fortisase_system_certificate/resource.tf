# [Local Certificate]
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

# [CA Certificate]
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

# [Remote Certificate]
resource "fortisase_system_certificate" "remote_certificate" {
  certificate_type = "remote-certificate"
  primary_key      = "example_remote_cert"
  file_content     = filebase64("./path/to/remote-cert-file.crt")
}

# [Remote CA Certificate]
resource "fortisase_system_certificate" "remote_ca_certificate" {
  certificate_type = "remote-ca-certificate"
  primary_key      = "example_remote_ca_cert"
  file_content     = filebase64("./path/to/remote-ca-file.crt")
}

# [HSM Local Certificate]
# This feature required HSM setup. Otherwise the request is rejected.
resource "fortisase_system_certificate" "hsm_local_certificate" {
  certificate_type = "hsm-local-certificate"
  primary_key      = "example_hsm_local_cert"
  file_content     = filebase64("./path/to/hsm-local-file.crt")
}
