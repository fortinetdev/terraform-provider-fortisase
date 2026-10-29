# GUI: Access & authentication -> SSO
# Requires VPN migration mode.
resource "fortisase_auth_sslvpn_saml_server" "sslvpn-sso" {
  digest_method   = "sha256"
  idp_entity_id   = "https://sts.windows.net/2a925438-d60d-42e8-a50a-2b0c527812e4/"
  idp_sign_on_url = "https://login.microsoftonline.com/2a925438-d60d-42e8-a50a-2b0c527812e4/saml2"
  idp_log_out_url = "https://login.microsoftonline.com/2a925438-d60d-42e8-a50a-2b0c527812e4/saml2"
  idp_certificate = {
    primary_key = fortisase_system_certificate.remote_cert_sslvpn.primary_key
    datasource  = "system/certificate/remote-certificates"
  }
  username   = "username"
  group_name = "groups"
  sp_cert = {
    primary_key = "FortiSASE Default Certificate"
    datasource  = "system/certificate/local-certificates"
  }
  scim_enabled     = false
  group_id         = ""
  entra_id_enabled = false
  domain_name      = "3"
  application_id   = "2"
}

resource "fortisase_system_certificate" "remote_cert_sslvpn" {
  certificate_type = "remote-certificate"
  primary_key      = "tf_remote_cert_sslvpn"
  file_content     = filebase64("./path/to/remote-cert-file.crt")
}
