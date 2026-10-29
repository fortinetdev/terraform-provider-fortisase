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
  # Reference remote certificate
  idp_certificate = {
    primary_key = fortisase_system_certificate.idp.primary_key
    datasource  = "system/certificate/remote-certificates"
  }

  # [Service Provider Certificate]
  # Reference local certificate
  sp_cert = {
    primary_key = fortisase_system_certificate.sp.primary_key
    datasource  = "system/certificate/local-certificates"
  }

  # [Digest Method]
  digest_method = "sha256" # "sha256", "sha1"

  # [SCIM]
  scim_enabled = false
}

resource "fortisase_system_certificate" "idp" {
  certificate_type = "remote-certificate"
  primary_key      = "tf-swg-idp-cert"
  file_content     = filebase64("./path/to/remote-cert-file.crt")
}

resource "fortisase_system_certificate" "sp" {
  certificate_type = "local-certificate"
  primary_key      = "tf-swg-sp-cert"
  format           = "pkcs12"
  file_content     = filebase64("./path/to/local-cert-file.p12")
  password         = "example_password"
}
