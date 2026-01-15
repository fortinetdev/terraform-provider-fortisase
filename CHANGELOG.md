## 1.2.0 (Unreleased)

## 1.1.0 (January 15, 2026)

DEPRECATIONS:
- Deprecate the `direction` attribute in `fortisase_security_profile_group`, `fortisase_security_profile_group_clone`, and all related security profile resources. While `direction` remains supported, it is recommended to omit this attribute.
- Remove datasource/fortisase_security_profile_groups. Please use datasource/fortisase_security_profile_group instead;
- resource/fortisase_auth_vpn_saml_server: The attribute `enabled` has been removed. The Terraform FortiSASE provider handles the enabled state internally;
- resource/fortisase_auth_swg_saml_server: The attribute `enabled` has been removed. The Terraform FortiSASE provider handles the enabled state internally;

FEATURES:
- **New Resource:** `fortisase_network_dns_rules`
- **New Resource:** `fortisase_endpoint_on_net_rules`
- **New Resource:** `fortisase_endpoint_profile`
- **New Resource:** `fortisase_endpoint_profile_clone`
- **New Resource:** `fortisase_security_cert_local_ca_certs`
- **New Resource:** `fortisase_security_cert_remote_certs`
- **New Resource:** `fortisase_infra_ipam_setting`
- **New Resource:** `fortisase_infra_secure_web_gateway_supplementary_data`

- **New Data Source:** `fortisase_network_dns_rules`
- **New Data Source:** `fortisase_endpoint_on_net_rules`
- **New Data Source:** `fortisase_endpoint_profile`
- **New Data Source:** `fortisase_security_cert_local_ca_certs`
- **New Data Source:** `fortisase_security_cert_remote_certs`
- **New Data Source:** `fortisase_infra_secure_web_gateway_supplementary_data`

IMPROVEMENTS:
- Support the schema of FortiSASE API 25.3.c;
- Include the possible values for each attribute in the descriptions;
- Improve examples and documentation for the resources and datasources;
- Add documentation for how to get username and password;
- Improve the error message displayed when an error is returned from the FortiSASE API;
- Handle requests sequentially when the FortiSASE API does not support parallel calls;
- Retry requests multiple times when the FortiSASE API returns errors due to internal instability;


BUG FIXES:
- Fix an issue where the attribute could not be set to empty using `var = []`;
- Fix an issue where user got "Provider produced inconsistent result after apply" for security profile related resources, the attributes `fortiguard_filters`, `fortiguard_local_category_filters`, `application_category_controls`, `fqdn_threat_feed_filters`, `domain_threat_feed_filters` in the related resources have been fixed;
- resource/fortisase_private_access_service_connections: Fix attribute `backup_links`;
- resource/fortisase_private_access_service_connections: Fix attribute `config.region_cost`, it can return correct result after apply;
- resource/fortisase_endpoint_sandbox_profiles: Change attribute `notification_type` from a string to a numeric type;
- resource/fortisase_endpoint_connection_profiles: Fix attribute `secure_internet_access.failover_sequence` produced inconsistent result after apply;
- resource/fortisase_endpoint_connection_profiles: Fix attribute `on_fabric_rule_set`, user can unset it by commenting out or removing the attribute;
- resource/fortisase_endpoint_connection_profiles: Ensure that destroying the resource clears the `on_fabric_rule_set` and `posture_check` attributes;
- resource/fortisase_endpoint_connection_profiles: Fix an issue where the resource may return before all configurations are fully applied on the server;
- resource/fortisase_security_profile_group: User can specify the enable status of the profiles. Check the example in the documentation for more details;
- resource/fortisase_security_web_filter_profile: Fix attribute `status` produced unexpected error when creating or updating the resource;
- resource/fortisase_security_profile_group: Fix error when `inspection_mode` is "certificate-inspection" or "no-inspection";
- resource/fortisase_security_profile_group: Fix the `host_exemptions` attribute to support "network/hosts";
- resource/fortisase_security_file_filter_profile: Fix the `monitor` attribute by adding the missing `primary_key`;
- resource/fortisase_security_outbound_policies: Fix issue where the `sources` attribute was not sent to the API;
- resource/fortisase_security_outbound_policies: A more detailed example has been added to the documentation to show how to use the resource;
- resource/fortisase_auth_vpn_saml_server: Fix issue where the resource cannot be deleted;
- resource/fortisase_private_access_service_connections_region_cost: Fix an issue where the resource may return before all configurations are fully applied on the server.
- resource/fortisase_endpoint_policies: Fix an issue where changing primary_key could cause an error.

## 1.0.0 (September 30, 2025)

FEATURES:
- Initial release. 74 resources, 110 datasources.
