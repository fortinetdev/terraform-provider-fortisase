resource "fortisase_security_recurring_schedules" "example" {
  primary_key = "example_name"
  days = [
    "sunday",
    "monday",
    "tuesday",
    "wednesday",
    "thursday"
  ]
  end_time   = "17:00"
  start_time = "09:02"
}
