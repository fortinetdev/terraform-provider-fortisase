# Replace these values with the actual endpoint device IDs.
# Destroying this resource does not undo the unquarantine operation.
# To undo this operation, apply a fortisase_endpoints_quarantine resource.
resource "fortisase_endpoints_unquarantine" "endpoints_unquarantine" {
  device_ids = [12345, 67890]
}
