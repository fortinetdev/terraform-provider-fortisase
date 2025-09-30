resource "fortisase_security_service_groups" "service_group" {
  primary_key = "service_group_name"
  proxy       = false
  members = [
    {
      primary_key = "ALL"
      datasource  = "security/services"
    }
  ]
}
