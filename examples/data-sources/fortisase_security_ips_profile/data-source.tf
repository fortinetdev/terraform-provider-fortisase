data "fortisase_security_ips_profile" "example" {
  direction   = "outbound-profiles" # "internal-profiles" or "outbound-profiles"
  primary_key = "default"           # The name of your existing profile group
}
