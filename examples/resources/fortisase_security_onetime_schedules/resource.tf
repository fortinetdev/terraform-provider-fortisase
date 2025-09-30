resource "fortisase_security_onetime_schedules" "example" {
  primary_key = "example_name"
  start_utc   = 1422835200 # It is the number of seconds since 1970-01-01 00:00:00 UTC
  end_utc     = 1485993600 # It is the number of seconds since 1970-01-01 00:00:00 UTC
}