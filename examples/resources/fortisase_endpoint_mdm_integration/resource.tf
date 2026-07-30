# This resource is available only with EMS 7.4.
resource "fortisase_endpoint_mdm_integration" "example" {
  primary_key     = "$sase-global"
  enabled         = false
  vendor          = "intune"
  auth_type       = "client_secret"
  deployment_type = "cloud"
  tenant_id       = "00000000-0000-0000-0000-000000000001"
  client_id       = "00000000-0000-0000-0000-000000000002"
  client_secret   = "tf-test-client-secret"
  url             = "https://graph.microsoft.com"
  region          = "northAmerica"
  site_name       = "tf-test-site"
  smart_group     = "tf-test-smart-group"
  username        = "tfadmin"
  password        = "fortinet"
  api_key         = "tf-test-api-key"
}
