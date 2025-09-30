resource "fortisase_security_domain_threat_feeds" "domain_threat_feed" {
  primary_key          = "example_name"
  refresh_rate         = 10
  status               = "enable"
  uri                  = "https://www.virustotal.com/api/v3/domains/google.com/threat-feed"
  username             = "example_username"
  password             = "example_password"
  basic_authentication = "enable"
}
