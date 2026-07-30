resource "fortisase_security_cert_remote_cert" "remote_cert_sslvpn" {
  cert_name    = "tf_remote_cert_sslvpn"
  file_content = base64encode(file("../path/to/certificate.crt"))
}

resource "fortisase_auth_sslvpn_saml_server" "sslvpn-sso" {
  digest_method   = "sha256"
  idp_entity_id   = "https://sts.windows.net/2a925438-d60d-42e8-a50a-2b0c527812e4/"
  idp_sign_on_url = "https://login.microsoftonline.com/2a925438-d60d-42e8-a50a-2b0c527812e4/saml2"
  idp_log_out_url = "https://login.microsoftonline.com/2a925438-d60d-42e8-a50a-2b0c527812e4/saml2"
  idp_certificate = {
    primary_key = fortisase_security_cert_remote_cert.remote_cert_sslvpn.cert_name
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
