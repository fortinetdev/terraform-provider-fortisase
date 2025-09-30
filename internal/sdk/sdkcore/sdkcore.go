// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),
// Yue Wang (@yuew-ftnt)

// Description: Description: SDK for FortiSASE Provider

package forticlient

func (c *FortiSDKClient) CreateEndpointsAccessProxyAuthorize(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/monitor-api/v1/endpoints/access-proxy/authorize"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) CreateEndpointsAccessProxyDisconnect(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/monitor-api/v1/endpoints/access-proxy/disconnect"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointsClientUserDetails(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/monitor-api/v1/endpoints/client-user-details/{clientUserId}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointsDetails(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/monitor-api/v1/endpoints/details/{deviceId}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) CreateEndpointsDisableManagement(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/monitor-api/v1/endpoints/disable-management"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointsDonut(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/monitor-api/v1/endpoints/donut/{donutType}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) CreateEndpointsEnableManagement(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/monitor-api/v1/endpoints/enable-management"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointsEndpointsWithSoftware(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/monitor-api/v1/endpoints/endpoints-with-software/{softwareId}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointsGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/monitor-api/v1/endpoints/groups/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointsSoftwareOnClientUser(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/monitor-api/v1/endpoints/software-on-client-user/{clientUserId}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointsSoftwareOnEndpoint(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/monitor-api/v1/endpoints/software-on-endpoint/{deviceId}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityBotnetDomainsStat(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/monitor-api/v1/security/botnet-domains/stat"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) CreateUserSwgSessionsDeauth(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/monitor-api/v1/user/swg/sessions/deauth"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) CreateUserVpnSessionsDeauth(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/monitor-api/v1/user/vpn/sessions/deauth"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateAuthFssoAgents(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/auth/fsso-agents/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadAuthFssoAgents(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/auth/fsso-agents/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteAuthFssoAgents(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/auth/fsso-agents/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateAuthFssoAgents(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/auth/fsso-agents"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateAuthLdapServers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/auth/ldap-servers/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadAuthLdapServers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/auth/ldap-servers/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteAuthLdapServers(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/auth/ldap-servers/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateAuthLdapServers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/auth/ldap-servers"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateAuthRadiusServers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/auth/radius-servers/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadAuthRadiusServers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/auth/radius-servers/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteAuthRadiusServers(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/auth/radius-servers/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateAuthRadiusServers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/auth/radius-servers"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateAuthSwgSamlServer(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/auth/swg-saml-server"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadAuthSwgSamlServer(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/auth/swg-saml-server"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateAuthUserGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/auth/user-groups/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadAuthUserGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/auth/user-groups/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteAuthUserGroups(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/auth/user-groups/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateAuthUserGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/auth/user-groups"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateAuthUsers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/auth/users/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadAuthUsers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/auth/users/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteAuthUsers(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/auth/users/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateAuthUsers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/auth/users"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateAuthVpnSamlServer(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/auth/vpn-saml-server"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadAuthVpnSamlServer(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/auth/vpn-saml-server"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateDemCustomSaasApps(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/dem/custom-saas-apps/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadDemCustomSaasApps(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/dem/custom-saas-apps/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteDemCustomSaasApps(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/dem/custom-saas-apps/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateDemCustomSaasApps(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/dem/custom-saas-apps"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateDemSpaApplications(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/dem/spa-applications/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadDemSpaApplications(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/dem/spa-applications/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteDemSpaApplications(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/dem/spa-applications/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateDemSpaApplications(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/dem/spa-applications"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateEndpointConnectionProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/endpoint/connection-profiles/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointConnectionProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/connection-profiles/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateEndpointFssoProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/endpoint/fsso-profiles/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointFssoProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/fsso-profiles/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateEndpointGroupAdUserProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/endpoint/group-ad-user-profiles/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointGroupAdUserProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/group-ad-user-profiles/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateEndpointGroupInvitationCodes(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/endpoint/group-invitation-codes/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointGroupInvitationCodes(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/group-invitation-codes/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteEndpointGroupInvitationCodes(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/endpoint/group-invitation-codes/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateEndpointGroupInvitationCodes(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/endpoint/group-invitation-codes"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateEndpointPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/endpoint/policies/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/policies/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteEndpointPolicies(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/endpoint/policies/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateEndpointPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/endpoint/policies"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) CreateEndpointPoliciesClone(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/endpoint/policies/{based_on}/clone"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateEndpointProtectionProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/endpoint/protection-profiles/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointProtectionProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/protection-profiles/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateEndpointSandboxProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/endpoint/sandbox-profiles/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointSandboxProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/sandbox-profiles/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateEndpointSettingProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/endpoint/setting-profiles/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointSettingProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/setting-profiles/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateEndpointZtnaProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/endpoint/ztna-profiles/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointZtnaProfiles(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/ztna-profiles/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateEndpointZtnaRules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/endpoint/ztna-rules/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointZtnaRules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/ztna-rules/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteEndpointZtnaRules(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/endpoint/ztna-rules/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateEndpointZtnaRules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/endpoint/ztna-rules"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadEndpointZtnaTags(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/endpoint/ztna-tags/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteEndpointZtnaTags(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/endpoint/ztna-tags/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateEndpointZtnaTags(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/endpoint/ztna-tags"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadInfraExtenders(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/infra/extenders/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadInfraFortigates(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/infra/fortigates/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateInfraSsids(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/infra/ssids/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadInfraSsids(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/infra/ssids/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteInfraSsids(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/infra/ssids/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateInfraSsids(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/infra/ssids"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadNetworkBasicInternetServices(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/network/basic-internet-services/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateNetworkHostGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/network/host-groups/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadNetworkHostGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/network/host-groups/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteNetworkHostGroups(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/network/host-groups/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateNetworkHostGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/network/host-groups"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateNetworkHosts(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/network/hosts/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadNetworkHosts(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/network/hosts/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteNetworkHosts(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/network/hosts/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateNetworkHosts(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/network/hosts"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateNetworkImplicitDnsRules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/network/implicit-dns-rules/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadNetworkImplicitDnsRules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/network/implicit-dns-rules/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadNetworkWildcardFqdnCustoms(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/network/wildcard-fqdn-customs/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityAntivirusFiletypes(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/antivirus-filetypes/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityAntivirusProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/antivirus-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityAntivirusProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/antivirus-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityAppCustomSignatures(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/app-custom-signatures/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityAppCustomSignatures(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/app-custom-signatures/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityAppCustomSignatures(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/app-custom-signatures/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityAppCustomSignatures(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/app-custom-signatures"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityApplicationCategories(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/application-categories/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityApplicationControlProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/application-control-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityApplicationControlProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/application-control-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityApplications(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/applications/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityDlpDataTypes(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/dlp-data-types/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityDlpDictionaries(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/dlp-dictionaries/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityDlpDictionaries(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/dlp-dictionaries/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityDlpDictionaries(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/dlp-dictionaries/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityDlpDictionaries(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/dlp-dictionaries"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityDlpExactDataMatches(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/dlp-exact-data-matches/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityDlpExactDataMatches(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/dlp-exact-data-matches/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityDlpExactDataMatches(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/dlp-exact-data-matches/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityDlpExactDataMatches(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/dlp-exact-data-matches"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityDlpFilePatterns(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/dlp-file-patterns/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityDlpFilePatterns(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/dlp-file-patterns/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityDlpFilePatterns(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/dlp-file-patterns/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityDlpFilePatterns(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/dlp-file-patterns"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityDlpFingerprintDatabases(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/dlp-fingerprint-databases/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityDlpFingerprintDatabases(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/dlp-fingerprint-databases/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityDlpFingerprintDatabases(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/dlp-fingerprint-databases/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityDlpFingerprintDatabases(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/dlp-fingerprint-databases"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityDlpProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/dlp-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityDlpProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/dlp-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityDlpSensors(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/dlp-sensors/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityDlpSensors(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/dlp-sensors/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityDlpSensors(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/dlp-sensors/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityDlpSensors(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/dlp-sensors"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityDnsFilterProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/dns-filter-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityDnsFilterProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/dns-filter-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityDomainThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/domain-threat-feeds/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityDomainThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/domain-threat-feeds/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityDomainThreatFeeds(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/domain-threat-feeds/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityDomainThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/domain-threat-feeds"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityEndpointToEndpointPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/endpoint-to-endpoint-policies/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityEndpointToEndpointPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/endpoint-to-endpoint-policies/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityEndpointToEndpointPolicies(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/endpoint-to-endpoint-policies/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityEndpointToEndpointPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/endpoint-to-endpoint-policies"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityEndpointToEndpointPoliciesClone(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/endpoint-to-endpoint-policies/{based_on}/clone"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityFileFilterProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/file-filter-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityFileFilterProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/file-filter-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityFortiguardCategories(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/fortiguard-categories/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityFortiguardLocalCategories(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/fortiguard-local-categories/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityFortiguardLocalCategories(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/fortiguard-local-categories/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityFortiguardLocalCategories(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/fortiguard-local-categories/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityFortiguardLocalCategories(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/fortiguard-local-categories"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityGeoipCountries(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/geoip-countries/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityInternalPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/internal-policies/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityInternalPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/internal-policies/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityInternalPolicies(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/internal-policies/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityInternalPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/internal-policies"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityInternalPoliciesClone(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/internal-policies/{based_on}/clone"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityInternalReversePolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/internal-reverse-policies/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityInternalReversePolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/internal-reverse-policies/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityInternalReversePolicies(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/internal-reverse-policies/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityInternalReversePolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/internal-reverse-policies"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityInternalReversePoliciesClone(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/internal-reverse-policies/{based_on}/clone"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityIpThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/ip-threat-feeds/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityIpThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/ip-threat-feeds/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityIpThreatFeeds(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/ip-threat-feeds/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityIpThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/ip-threat-feeds"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityIpsCustomSignatures(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/ips-custom-signatures/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityIpsCustomSignatures(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/ips-custom-signatures/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityIpsCustomSignatures(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/ips-custom-signatures/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityIpsCustomSignatures(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/ips-custom-signatures"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityIpsProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/ips-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityIpsProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/ips-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityOnetimeSchedules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/onetime-schedules/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityOnetimeSchedules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/onetime-schedules/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityOnetimeSchedules(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/onetime-schedules/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityOnetimeSchedules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/onetime-schedules"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityOutboundPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/outbound-policies/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityOutboundPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/outbound-policies/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityOutboundPolicies(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/outbound-policies/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityOutboundPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/outbound-policies"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityOutboundPoliciesClone(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/outbound-policies/{based_on}/clone"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityProfileGroup(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/profile-group/{direction}/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityProfileGroup(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/profile-group/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityProfileGroup(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/profile-group/{direction}/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityProfileGroup(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/profile-group/{direction}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityProfileGroupClone(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/profile-group/{direction}/{based_on}/clone"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityProfileGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/profile-groups/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityRecurringSchedules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/recurring-schedules/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityRecurringSchedules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/recurring-schedules/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityRecurringSchedules(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/recurring-schedules/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityRecurringSchedules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/recurring-schedules"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityScheduleGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/schedule-groups/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityScheduleGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/schedule-groups/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityScheduleGroups(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/schedule-groups/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityScheduleGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/schedule-groups"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityServiceCategories(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/service-categories/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityServiceGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/service-groups/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityServiceGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/service-groups/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityServiceGroups(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/service-groups/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityServiceGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/service-groups"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityServices(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/services/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityServices(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/services/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityServices(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/services/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityServices(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/services"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecuritySslSshProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/ssl-ssh-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecuritySslSshProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/ssl-ssh-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityUrlThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/url-threat-feeds/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityUrlThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/url-threat-feeds/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityUrlThreatFeeds(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v2/security/url-threat-feeds/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityUrlThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v2/security/url-threat-feeds"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityVideoFilterFortiguardCategories(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/video-filter-fortiguard-categories/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityVideoFilterProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/video-filter-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityVideoFilterProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/video-filter-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityVideoFilterYoutubeKey(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/video-filter-youtube-key"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityVideoFilterYoutubeKey(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/video-filter-youtube-key"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityWebFilterProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v2/security/web-filter-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityWebFilterProfile(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/security/web-filter-profile/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageAuthFssoAgents(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/auth/fsso-agents/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageAuthLdapServers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/auth/ldap-servers/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageAuthRadiusServers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/auth/radius-servers/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageAuthUserGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/auth/user-groups/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageEndpointZtnaTags(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/endpoint/ztna-tags/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageInfraSsids(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/infra/ssids/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageNetworkHostGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/network/host-groups/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageNetworkHosts(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/network/hosts/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityAppCustomSignatures(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/app-custom-signatures/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityDlpDictionaries(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/dlp-dictionaries/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityDlpExactDataMatches(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/dlp-exact-data-matches/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityDlpFilePatterns(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/dlp-file-patterns/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityDlpFingerprintDatabases(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/dlp-fingerprint-databases/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityDlpSensors(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/dlp-sensors/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityDomainThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/domain-threat-feeds/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityEndpointToEndpointPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/endpoint-to-endpoint-policies/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityFortiguardLocalCategories(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/fortiguard-local-categories/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityInternalPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/internal-policies/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityInternalReversePolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/internal-reverse-policies/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityIpThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/ip-threat-feeds/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityIpsCustomSignatures(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/ips-custom-signatures/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityOnetimeSchedules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/onetime-schedules/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityOutboundPolicies(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/outbound-policies/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityProfileGroup(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/profile-group/{direction}/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityRecurringSchedules(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/recurring-schedules/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityScheduleGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/schedule-groups/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityServiceGroups(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/service-groups/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityServices(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/services/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) ReadUsageSecurityUrlThreatFeeds(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v2/usage/security/url-threat-feeds/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) UpdatePrivateAccessNetworkConfiguration(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v1/private-access/network-configuration"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadPrivateAccessNetworkConfiguration(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v1/private-access/network-configuration"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeletePrivateAccessNetworkConfiguration(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v1/private-access/network-configuration"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreatePrivateAccessNetworkConfiguration(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v1/private-access/network-configuration"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdatePrivateAccessServiceConnections(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v1/private-access/service-connections/{service-connection-id}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadPrivateAccessServiceConnections(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v1/private-access/service-connections/{service-connection-id}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeletePrivateAccessServiceConnections(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v1/private-access/service-connections/{service-connection-id}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreatePrivateAccessServiceConnections(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v1/private-access/service-connections"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) CreatePrivateAccessServiceConnectionsAuth(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v1/private-access/service-connections/{service-connection-id}/auth"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) CreatePrivateAccessServiceConnectionsRegionCost(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v1/private-access/service-connections/region_cost"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityCertLocalCerts(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v1/security/cert/local-certs/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityCertLocalCerts(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v1/security/cert/local-certs/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityCertLocalCerts(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v1/security/cert/local-certs"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityCertRemoteCaCerts(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v1/security/cert/remote-ca-certs/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityCertRemoteCaCerts(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v1/security/cert/remote-ca-certs/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityCertRemoteCaCerts(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v1/security/cert/remote-ca-certs"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) UpdateSecurityPkiUsers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "PUT"
	input_model.URL = "/resource-api/v1/security/pki-users/{primaryKey}"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
func (c *FortiSDKClient) ReadSecurityPkiUsers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "GET"
	input_model.URL = "/resource-api/v1/security/pki-users/{primaryKey}"
	input_model.update()

	output, err = read(c, input_model)
	return
}
func (c *FortiSDKClient) DeleteSecurityPkiUsers(input_model *InputModel) (err error) {
	input_model.HTTPMethod = "DELETE"
	input_model.URL = "/resource-api/v1/security/pki-users/{primaryKey}"
	input_model.update()

	err = delete(c, input_model)
	return
}
func (c *FortiSDKClient) CreateSecurityPkiUsers(input_model *InputModel) (output map[string]interface{}, err error) {
	input_model.HTTPMethod = "POST"
	input_model.URL = "/resource-api/v1/security/pki-users"
	input_model.update()

	output, err = createUpdate(c, input_model)
	return
}
