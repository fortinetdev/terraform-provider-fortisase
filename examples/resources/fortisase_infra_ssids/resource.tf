resource "fortisase_infra_ssids" "infra_ssids" {
  primary_key    = "terraform"
  broadcast_ssid = "enable"
  security_mode  = "wpa2-only-personal"
  pre_shared_key = "1234567890"
  wifi_ssid      = "wifi_ssid"
  client_limit   = 100
}
