# Please make sure to use **single quotes** for the primary_key, otherwise $sase can be interpreted as a variable, not a string.
terraform import fortisase_auth_vpn_saml_server.{{your_resource_name}} '$sase-global'
