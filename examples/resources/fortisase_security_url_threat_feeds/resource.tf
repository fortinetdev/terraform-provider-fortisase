resource "fortisase_security_url_threat_feeds" "url_threat_feeds" {
  primary_key          = "url_threat_feeds"
  refresh_rate         = 10
  status               = "enable"
  uri                  = "https://www.virustotal.com/api/v3/domains/google.com/threat-feed"
  basic_authentication = "enable"
  username             = "your_username"
  password             = "your_password"
}
