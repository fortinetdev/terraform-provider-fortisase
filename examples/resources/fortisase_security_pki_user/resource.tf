resource "fortisase_security_pki_user" "pki_user" {
  primary_key = "example_name"
  ca = {
    name = "Fortinet_CA"
  }
}
