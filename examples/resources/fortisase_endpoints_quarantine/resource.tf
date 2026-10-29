# Replace these values with the actual endpoint device IDs.
# Destroying this resource does not undo the quarantine operation.
# To undo this operation, apply a fortisase_endpoints_unquarantine resource.
resource "fortisase_endpoints_quarantine" "endpoints_quarantine" {
  device_ids = [12345, 67890]
}
