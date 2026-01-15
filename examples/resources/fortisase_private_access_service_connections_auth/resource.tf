# Update the authentication method of an existing service connection.
# Note: If your secure private access resource was originally created with
#       `fortisase_private_access_service_connections`, you should update that
#       resource directly instead of using this one.

resource "fortisase_private_access_service_connections_auth" "example" {
  service_connection_id = "existing_service_connection_id"

  ## Method1: Pre-shared Key
  auth                 = "psk"
  ipsec_pre_shared_key = "new_shared_key"

  ## Method2: Certificate-Based Authentication
  # auth                  = "pki"
  # ipsec_cert_name       = "existing_cert_name"
  # ipsec_peer_name       = "existing_pki_name"
}
