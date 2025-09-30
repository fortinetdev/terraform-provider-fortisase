resource "fortisase_auth_radius_servers" "radius_server" {
  primary_key                    = "radius_server"
  primary_secret                 = "radius"
  primary_server                 = "2.3.4.5"
  auth_type                      = "auto"
  included_in_default_user_group = false
}
