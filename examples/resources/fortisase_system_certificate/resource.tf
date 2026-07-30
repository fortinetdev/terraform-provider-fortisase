# [Local Certificate]
resource "fortisase_system_certificate" "local_certificate" {
  certificate_type = "local-certificate"
  primary_key      = "example_local_cert"
  format           = "regular" # "regular" or "pkcs12"
  file_content     = base64encode(file("../path/to/my-server-cert.pem"))
  key_file_content = base64encode(file("../path/to/my-server-key.pem")) # Only required if format is "regular"
}

# [CA Certificate]
resource "fortisase_system_certificate" "ca_certificate" {
  certificate_type = "ca-certificate"
  primary_key      = "example_ca_cert"
  format           = "regular" # "regular" or "pkcs12"
  file_content     = base64encode(file("../path/to/my-ca-cert.pem"))
  key_file_content = base64encode(file("../path/to/my-ca-key.pem")) # Only required if format is "regular"
}

# [Remote Certificate]
resource "fortisase_system_certificate" "remote_certificate" {
  certificate_type = "remote-certificate"
  primary_key      = "example_remote_cert"
  file_content     = base64encode(file("../path/to/my-server-cert.pem"))
}

# [Remote CA Certificate]
resource "fortisase_system_certificate" "remote_ca_certificate" {
  certificate_type = "remote-ca-certificate"
  primary_key      = "example_remote_ca_cert"
  file_content     = base64encode(file("../path/to/my-ca-cert.pem"))
}

# [HSM Local Certificate]
# Requires HSM to be enabled; otherwise the request is rejected with 403.
resource "fortisase_system_certificate" "hsm_local_certificate" {
  certificate_type = "hsm-local-certificate"
  primary_key      = "example_hsm_local_cert"
  file_content     = base64encode(file("../path/to/my-ca-cert.pem"))
}
