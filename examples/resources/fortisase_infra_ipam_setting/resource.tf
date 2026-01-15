resource "random_string" "unique_id" {
  length  = 8
  upper   = false
  lower   = true
  numeric = true
  special = false
}

# Note: Destroying this resource couldn't remove the IPAM setting. You need to set pools = [] to remove the IPAM setting.
resource "fortisase_infra_ipam_setting" "example" {
  pools = [
    {
      name             = "${random_string.unique_id.result}1" # Random unique ID to avoid conflicts
      subnet           = "100.65.0.0/16"
      excluded_subnets = []
    },
    {
      name   = "${random_string.unique_id.result}2" # Random unique ID to avoid conflicts
      subnet = "172.16.0.0/12"
      excluded_subnets = [
        {
          subnet = "172.16.0.0/16"
        }
      ],
    }
  ]
}
