resource "fortisase_auth_fsso_agents" "fsso_agents" {
  primary_key      = "fsso_agent"
  name             = "fsso_agent"
  server           = "1.2.3.4"
  status           = "disconnected"
  password         = "password"
  ssl_trusted_cert = "remote_ca_certs"
}
