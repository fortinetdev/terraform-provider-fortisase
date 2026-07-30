## 1.4.0 (Unreleased)

## 1.3.0 (July 30, 2026)

FEATURES:
- **New Resource:** `fortisase_endpoint_mdm_integration`
- **New Resource:** `fortisase_security_per_ip_traffic_shaper`
- **New Resource:** `fortisase_security_traffic_shaper`
- **New Resource:** `fortisase_security_traffic_shaping_policy`
- **New Resource:** `fortisase_security_html_template`
- **New Resource:** `fortisase_security_html_template_image`
- **New Resource:** `fortisase_security_internal_proxy_policy`
- **New Resource:** `fortisase_security_outbound_proxy_policy`
- **New Resource:** `fortisase_security_internal_proxy_policy_clone`
- **New Resource:** `fortisase_security_outbound_proxy_policy_clone`
- **New Resource:** `fortisase_system_certificate`

- **New Data Source:** `fortisase_endpoint_mdm_integration`
- **New Data Source:** `fortisase_security_per_ip_traffic_shaper`
- **New Data Source:** `fortisase_security_traffic_shaper`
- **New Data Source:** `fortisase_security_traffic_shaping_policy`
- **New Data Source:** `fortisase_security_html_template`
- **New Data Source:** `fortisase_security_html_template_image`
- **New Data Source:** `fortisase_security_internal_proxy_policy`
- **New Data Source:** `fortisase_security_outbound_proxy_policy`

IMPROVEMENTS:
- Support the schema of FortiSASE API 26.2.2;

BUG FIXES:
- resource/fortisase_endpoint_ztna_tag_rule: Fix error related with argument sensitivities.
- datasource: Mark all non-configurable arguments as read-only.

Breaking Changes:
- resource/fortisase_security_application_control_profile: This resource has been redesigned due to backend API changes. Please refer to the documentation for the updated configuration and usage.

DEPRECATIONS:
- resource/fortisase_endpoint_connection_profiles: "connect_to_forti_sase" is deprecated, please use "connect_to_fortisase" instead.
- resource/fortisase_endpoint_connection_profiles: "available_vp_ns" is deprecated, please use "available_vpns" instead.
- Resources and data sources using plural naming are deprecated. They will remain supported in future releases for backward compatibility, but users are encouraged to migrate to the corresponding singularly named resources and data sources.
- The certificate-related resources `fortisase_security_cert_*` are deprecated. Please use fortisase_system_certificate instead.

## 1.2.0 (April 29, 2026)

FEATURES:
- **New Resource:** `fortisase_auth_sslvpn_saml_server`
- **New Resource:** `fortisase_endpoint_ztna_tag_rule`

- **New Data Source:** `fortisase_auth_sslvpn_saml_server`
- **New Data Source:** `fortisase_browser_provision`
- **New Data Source:** `fortisase_endpoint_ztna_tag_rule`
- **New Data Source:** `fortisase_infra_data_transfer`

DEPRECATIONS:
- All fortisase_usage_xxx data sources are deprecated and will be removed in version 1.4.0. If you still require these data sources, please contact us by opening a GitHub issue.

BUG FIXES:
- resource/fortisase_auth_vpn_saml_server: Fix an issue where the resource could report an error during deletion;
- resource/fortisase_endpoint_profile: Improve stability to reduce errors during creation;
- resource/fortisase_endpoint_policies: Improve stability to reduce errors during creation;
- resource/fortisase_endpoint_connection_profiles: Improve stability to reduce errors during creation;
- resource/fortisase_security_internal_policies: Report error when the resource is not created;
- resource/fortisase_security_internal_reverse_policies: Report error when the resource is not created;

IMPROVEMENTS:
- Support the schema of FortiSASE API 26.1.1;
- Retry requests if the API returns 502 error;
- Add hidden resource lock to reduce the API 500 error;
- Report a warning instead of an error when user input fails validation;
- Report an error if the specified resource is not compatible with the user's EMS version;

## 1.1.0 (January 15, 2026)

DEPRECATIONS:
- Deprecate the `direction` attribute in `fortisase_security_profile_group`, `fortisase_security_profile_group_clone`, and all related security profile resources. While `direction` remains supported, it is recommended to omit this attribute;
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
- resource/fortisase_private_access_service_connections_region_cost: Fix an issue where the resource may return before all configurations are fully applied on the server;
- resource/fortisase_endpoint_policies: Fix an issue where changing primary_key could cause an error;

## 1.0.0 (September 30, 2025)

FEATURES:
- Initial release. 74 resources, 110 datasources.
