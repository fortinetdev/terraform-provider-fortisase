# GUI: Access & authentication -> Fortinet Single Sign-On (FSSO)
resource "fortisase_auth_fsso_agent" "fsso_agent" {
  primary_key = "fsso_agent_name"
  name        = "fsso_agent_name"
  server      = "1.2.3.4"
  port        = "8001"
  password    = "password"
  ssl_trusted_cert = {
    primary_key = fortisase_system_certificate.remote_ca_certificate.primary_key
    datasource  = "system/certificate/remote-ca-certificates"
  }
}

resource "fortisase_system_certificate" "remote_ca_certificate" {
  certificate_type = "remote-ca-certificate"
  primary_key      = "example_remote_ca_cert"
  file_content     = filebase64("./path/to/remote-ca-file.crt")
}
