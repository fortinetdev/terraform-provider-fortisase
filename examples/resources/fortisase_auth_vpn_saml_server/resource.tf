resource "fortisase_auth_vpn_saml_server" "vpn_sso" {
  primary_key     = "$sase-global"
  enabled         = true
  digest_method   = "sha256"
  idp_entity_id   = "https://sts.windows.net/example1/"
  idp_sign_on_url = "https://login.microsoftonline.com/example/saml"
  idp_log_out_url = "https://login.microsoftonline.com/example/saml"
  idp_certificate = {
    primary_key = "certificate"
    datasource  = "system/certificate/remote-certificates"
  }
  username   = "example_username"
  group_name = "example_group_name"
  sp_cert = {
    primary_key = "FortiSASE Default Certificate"
    datasource  = "system/certificate/local-certificates"
  }
  scim_enabled     = false
  group_id         = ""
  entra_id_enabled = false
}
