resource "fortisase_dem_spa_applications" "spa_application" {
  primary_key          = "example_name"
  server               = "string"
  latency_threshold    = 10000000
  jitter_threshold     = 10000000
  packetloss_threshold = 100
  interval             = 20
  fail_time            = 1
  recovery_time        = 1
}
