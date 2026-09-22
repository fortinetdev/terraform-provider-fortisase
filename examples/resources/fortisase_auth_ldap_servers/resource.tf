resource "fortisase_auth_ldap_server" "ldap_server" {
  primary_key = "ldap_server"
  server      = "10.0.10.20"
  port        = 389

  # [Common name identifier]
  cnid = "cn"

  # [Distinguished name]
  dn = "cn=admin,dc=example,dc=com"

  # [Secure connection]
  ### Option A: Disable secure connection
  secure_connection = false

  ### Option B: Enable secure connection without client certificate authentication
  # secure_connection             = true
  # server_identity_check_enabled = true
  # password_renewal_enabled      = false
  # certificate = {
  #   primary_key = "remote_ca_name"
  #   datasource  = "system/certificate/remote-ca-certificates"
  # }
  # client_cert_auth_enabled = false

  ### Option C: Enable secure connection with client certificate authentication
  # secure_connection             = true
  # server_identity_check_enabled = true
  # password_renewal_enabled      = false
  # certificate = {
  #   primary_key = "certificategui"
  #   datasource  = "system/certificate/remote-ca-certificates"
  # }
  # client_cert_auth_enabled = true
  # client_cert = {
  #   primary_key = "local_cert_name"
  #   datasource  = "system/certificate/local-certificates"
  # }

  # [Advanced group matching]
  ### Option A: Disable advanced group matching
  advanced_group_matching_enabled = false

  ### Option B: Enable advanced group matching and check the user attribute
  # advanced_group_matching_enabled = true
  # group_member_check              = "user-attr"
  # member_attribute                = "memberOf"
  # group_filter                    = "(objectClass=group)"
  # group_search_base               = "ou=groups,dc=example,dc=com"

  ### Option C: Enable advanced group matching and check the group object
  # advanced_group_matching_enabled = true
  # group_member_check              = "group-object"
  # member_attribute                = "member"
  # group_object_filter             = "(objectClass=group)"

  ### Option D: Enable advanced group matching and check the POSIX group object
  # advanced_group_matching_enabled = true
  # group_member_check              = "posix-group-object"
  # member_attribute                = "memberUid"
  # group_search_base               = "ou=groups,dc=example,dc=com"
  # group_object_filter             = "(objectClass=posixGroup)"


  # [Bind type]
  ### Option A: Simple
  bind_type = "simple"

  ### Option B: Anonymous
  # bind_type = "anonymous"

  ### Option C: Regular
  # bind_type = "regular"
  # username  = "cn=ldap-reader,dc=example,dc=com"
  # password  = "your_password"
}
