resource "fortisase_security_service_group" "service_group" {
  primary_key = "service_group_name"
  proxy       = false
  members = [
    {
      primary_key = "ALL"
      datasource  = "security/services"
    }
  ]
}
