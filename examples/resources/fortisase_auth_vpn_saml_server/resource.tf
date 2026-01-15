# Agent user single sign-on (SSO)
# Populate the following fields with configuration from your Identity Provider.
resource "fortisase_auth_vpn_saml_server" "vpn_sso" {
  # [Identity Provider Configuration]
  idp_entity_id   = "https://sts.windows.net/example/"
  idp_sign_on_url = "https://login.microsoftonline.com/example/saml"
  idp_log_out_url = "https://login.microsoftonline.com/example/saml"

  # [SAML Claims Mapping]
  username   = "example_username"
  group_name = "example_group_name"

  # [SAML Group Matching Group ID]
  ## Disable SAML Claims Mapping
  group_id = ""
  ## Enable SAML Claims Mapping
  # group_id = "123"

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

  # [Microsoft Entra ID Options]
  ## Disable
  entra_id_enabled = false
  ## Enable
  # entra_id_enabled = true
  # domain_name      = "example.domain.name"
  # application_id   = "1234"

  # [SCIM]
  scim_enabled = false
}
