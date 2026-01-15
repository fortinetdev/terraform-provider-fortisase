# Proxy user single sign-on (SSO)
# To configure this resource, please enable proxy configuration.
# Populate the following fields with configuration from your Identity Provider.
resource "fortisase_auth_swg_saml_server" "swg_sso" {
  # [Identity Provider Configuration]
  idp_entity_id   = "https://sts.windows.net/example/"
  idp_sign_on_url = "https://login.microsoftonline.com/example/saml"
  idp_log_out_url = "https://login.microsoftonline.com/example/saml"

  # [SAML Claims Mapping]
  username   = "example_username"
  group_name = "example_group_name"

  # [SAML Group Matching Group ID]
  ## Disable SAML Claims Mapping
  group_match = ""
  ## Enable SAML Claims Mapping
  # group_match = "123"

  # [IdP Certificate]
  # References a remote certificate stored in the system
  idp_certificate = {
    primary_key = "certificate"
    datasource  = "system/certificate/remote-certificates"
  }

  # [Service Provider Certificate]
  # References a local certificate stored in the system
  sp_cert = {
    primary_key = "FortiSASE Default Certificate"
    datasource  = "system/certificate/local-certificates"
  }

  # [Digest Method]
  digest_method = "sha256" # "sha256", "sha1"

  # [SCIM]
  scim_enabled = false
}
