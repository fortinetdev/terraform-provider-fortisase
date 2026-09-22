# GUI: Access & authentication -> Fortinet Single Sign-On (FSSO)
resource "fortisase_auth_fsso_agent" "fsso_agent" {
  primary_key      = "fsso_agent_name"
  name             = "fsso_agent_name"
  server           = "1.2.3.4"
  port             = "8001"
  password         = "password"
  ssl_trusted_cert = "existing_remote_ca_certificate_name"
}
