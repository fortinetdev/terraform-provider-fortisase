resource "fortisase_security_schedule_group" "example" {
  primary_key = "example_name"
  members = [{
    datasource  = "security/recurring-schedules"
    primary_key = "always"
  }]
}
