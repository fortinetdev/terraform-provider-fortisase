resource "fortisase_endpoints_disable_management" "endpoints_disable_management" {
  endpoints = [{
    device_id = "1"
    hostname  = "test"
  }]
}
