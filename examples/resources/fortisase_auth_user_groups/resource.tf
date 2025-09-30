resource "fortisase_auth_users" "user" {
  primary_key = "user_001@example.com"
  auth_type   = "password"
  status      = "enable"
  email       = "user_001@example.com"
  password    = "password"
}

resource "fortisase_auth_ldap_servers" "ldap_server" {
  primary_key                     = "ldap_server"
  server                          = "1.2.3.4"
  port                            = 1234
  cnid                            = "test"
  dn                              = "cn=admin,dc=example,dc=com"
  client_cert_auth_enabled        = false
  bind_type                       = "simple"
  secure_connection               = false
  server_identity_check_enabled   = true
  advanced_group_matching_enabled = true
  group_member_check              = "user-attr"
  group_filter                    = "cn=group,dc=example,dc=com"
  group_search_base               = "dc=example,dc=com"
}

resource "fortisase_auth_user_groups" "user_group" {
  primary_key = "user_group"
  group_type  = "firewall"
  local_users = [
    {
      primary_key = fortisase_auth_users.user.primary_key
      datasource  = "auth/users"
    }
  ]
  remote_user_groups = [
    {
      server = {
        primary_key = fortisase_auth_ldap_servers.ldap_server.primary_key
        datasource  = "auth/ldap-servers"
      }
      matches = ["group1"]
    }
  ]
}
