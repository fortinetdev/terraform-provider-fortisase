# GUI: Security -> Schedules -> Schedule groups
resource "fortisase_security_schedule_group" "example" {
  primary_key = "tf-workdays-group"

  # Please only select one of the following
  # Option: Recurring schedule
  members = [
    {
      datasource  = "security/recurring-schedules"
      primary_key = fortisase_security_recurring_schedule.workdays.primary_key
    }
  ]

  # Option: Always
  # members = [
  #   {
  #     datasource  = "security/recurring-schedules"
  #     primary_key = "always"
  #   }
  # ]

  # Option: One-time schedule
  # members = [
  #   {
  #     datasource  = "security/onetime-schedules"
  #     primary_key = "one_time_key_name"
  #   }
  # ]
}

resource "fortisase_security_recurring_schedule" "workdays" {
  primary_key = "tf-workdays"
  days        = ["monday", "tuesday", "wednesday", "thursday", "friday"]
  # Recurring schedule times are in GMT (UTC).
  start_time = "09:00"
  end_time   = "17:00"
}
