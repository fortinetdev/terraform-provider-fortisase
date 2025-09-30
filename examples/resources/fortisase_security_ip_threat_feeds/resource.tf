resource "fortisase_security_ip_threat_feeds" "ip_threat_feeds" {
  primary_key          = "example_name"
  refresh_rate         = 10
  status               = "enable"
  uri                  = "https://www.virustotal.com/api/v3/domains/google.com/threat-feed"
  username             = "fortinet"
  password             = "fortinet"
  basic_authentication = "enable"
}
