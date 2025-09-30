# To configure this resource, please enable proxy configuration.
resource "fortisase_auth_swg_saml_server" "swg-sso" {
  primary_key     = "$sase-global"
  enabled         = true
  digest_method   = "sha256"
  idp_entity_id   = "https://sts.windows.net/example/"
  idp_sign_on_url = "https://login.microsoftonline.com/example/saml2"
  idp_log_out_url = "https://login.microsoftonline.com/example/saml2"
  idp_certificate = {
    primary_key = "certificate"
    datasource  = "system/certificate/remote-certificates"
  }
  username    = "username"
  group_name  = "group"
  group_match = ""
  sp_cert = {
    primary_key = "FortiSASE Default Certificate"
    datasource  = "system/certificate/local-certificates"
  }
  scim_enabled = false
}
