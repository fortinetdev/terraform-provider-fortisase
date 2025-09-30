resource "fortisase_auth_users" "user" {
  primary_key = "user_001@example.com"
  auth_type   = "password"
  status      = "enable"
  email       = "user_001@example.com"
  password    = "example_password" # This value is fixed once created and cannot be changed.
}
