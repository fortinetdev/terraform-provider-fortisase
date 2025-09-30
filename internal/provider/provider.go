// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure FortisaseProvider satisfies various provider interfaces.
var _ provider.Provider = &FortisaseProvider{}

// FortisaseProvider defines the provider implementation.
type FortisaseProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// FortisaseProviderModel describes the provider data model.
type FortisaseProviderModel struct {
	Username     types.String `tfsdk:"username"`
	Password     types.String `tfsdk:"password"`
	AccessToken  types.String `tfsdk:"access_token"`
	RefreshToken types.String `tfsdk:"refresh_token"`
}

func (p *FortisaseProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "fortisase"
}

func (p *FortisaseProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Description: "The username of API user.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "The password of API user.",
				Optional:    true,
			},
			"access_token": schema.StringAttribute{
				Description: "The access token of API user.",
				Optional:    true,
			},
			"refresh_token": schema.StringAttribute{
				Description: "The refresh token of API user.",
				Optional:    true,
			},
		},
	}
}

func (p *FortisaseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data FortisaseProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	config := Config{
		Username:     data.Username.ValueString(),
		Password:     data.Password.ValueString(),
		AccessToken:  data.AccessToken.ValueString(),
		RefreshToken: data.RefreshToken.ValueString(),
	}

	sdkClient, err := config.CreateClient()
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error to create client: %v", err),
			"",
		)
	}
	resp.DataSourceData = sdkClient
	resp.ResourceData = sdkClient
	resp.EphemeralResourceData = sdkClient
}
func (p *FortisaseProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newResourceEndpointsAccessProxyAuthorize,
		newResourceEndpointsAccessProxyDisconnect,
		newResourceEndpointsDisableManagement,
		newResourceEndpointsEnableManagement,
		newResourceUserSwgSessionsDeauth,
		newResourceUserVpnSessionsDeauth,
		newResourceAuthFssoAgents,
		newResourceAuthLdapServers,
		newResourceAuthRadiusServers,
		newResourceAuthSwgSamlServer,
		newResourceAuthUserGroups,
		newResourceAuthUsers,
		newResourceAuthVpnSamlServer,
		newResourceDemCustomSaasApps,
		newResourceDemSpaApplications,
		newResourceEndpointConnectionProfiles,
		newResourceEndpointFssoProfiles,
		newResourceEndpointGroupAdUserProfiles,
		newResourceEndpointGroupInvitationCodes,
		newResourceEndpointPolicies,
		newResourceEndpointPoliciesClone,
		newResourceEndpointProtectionProfiles,
		newResourceEndpointSandboxProfiles,
		newResourceEndpointSettingProfiles,
		newResourceEndpointZtnaProfiles,
		newResourceEndpointZtnaRules,
		newResourceEndpointZtnaTags,
		newResourceInfraSsids,
		newResourceNetworkHostGroups,
		newResourceNetworkHosts,
		newResourceNetworkImplicitDnsRules,
		newResourceSecurityAntivirusProfile,
		newResourceSecurityAppCustomSignatures,
		newResourceSecurityApplicationControlProfile,
		newResourceSecurityDlpDictionaries,
		newResourceSecurityDlpExactDataMatches,
		newResourceSecurityDlpFilePatterns,
		newResourceSecurityDlpFingerprintDatabases,
		newResourceSecurityDlpProfile,
		newResourceSecurityDlpSensors,
		newResourceSecurityDnsFilterProfile,
		newResourceSecurityDomainThreatFeeds,
		newResourceSecurityEndpointToEndpointPolicies,
		newResourceSecurityEndpointToEndpointPoliciesClone,
		newResourceSecurityFileFilterProfile,
		newResourceSecurityFortiguardLocalCategories,
		newResourceSecurityInternalPolicies,
		newResourceSecurityInternalPoliciesClone,
		newResourceSecurityInternalReversePolicies,
		newResourceSecurityInternalReversePoliciesClone,
		newResourceSecurityIpThreatFeeds,
		newResourceSecurityIpsCustomSignatures,
		newResourceSecurityIpsProfile,
		newResourceSecurityOnetimeSchedules,
		newResourceSecurityOutboundPolicies,
		newResourceSecurityOutboundPoliciesClone,
		newResourceSecurityProfileGroup,
		newResourceSecurityProfileGroupClone,
		newResourceSecurityRecurringSchedules,
		newResourceSecurityScheduleGroups,
		newResourceSecurityServiceGroups,
		newResourceSecurityServices,
		newResourceSecuritySslSshProfile,
		newResourceSecurityUrlThreatFeeds,
		newResourceSecurityVideoFilterProfile,
		newResourceSecurityVideoFilterYoutubeKey,
		newResourceSecurityWebFilterProfile,
		newResourcePrivateAccessNetworkConfiguration,
		newResourcePrivateAccessServiceConnections,
		newResourcePrivateAccessServiceConnectionsAuth,
		newResourcePrivateAccessServiceConnectionsRegionCost,
		newResourceSecurityCertLocalCerts,
		newResourceSecurityCertRemoteCaCerts,
		newResourceSecurityPkiUsers,
	}
}

func (p *FortisaseProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newDatasourceEndpointsClientUserDetails,
		newDatasourceEndpointsDetails,
		newDatasourceEndpointsDonut,
		newDatasourceEndpointsEndpointsWithSoftware,
		newDatasourceEndpointsGroups,
		newDatasourceEndpointsSoftwareOnClientUser,
		newDatasourceEndpointsSoftwareOnEndpoint,
		newDatasourceSecurityBotnetDomainsStat,
		newDatasourceAuthFssoAgents,
		newDatasourceAuthLdapServers,
		newDatasourceAuthRadiusServers,
		newDatasourceAuthSwgSamlServer,
		newDatasourceAuthUserGroups,
		newDatasourceAuthUsers,
		newDatasourceAuthVpnSamlServer,
		newDatasourceDemCustomSaasApps,
		newDatasourceDemSpaApplications,
		newDatasourceEndpointConnectionProfiles,
		newDatasourceEndpointFssoProfiles,
		newDatasourceEndpointGroupAdUserProfiles,
		newDatasourceEndpointGroupInvitationCodes,
		newDatasourceEndpointPolicies,
		newDatasourceEndpointProtectionProfiles,
		newDatasourceEndpointSandboxProfiles,
		newDatasourceEndpointSettingProfiles,
		newDatasourceEndpointZtnaProfiles,
		newDatasourceEndpointZtnaRules,
		newDatasourceEndpointZtnaTags,
		newDatasourceInfraExtenders,
		newDatasourceInfraFortigates,
		newDatasourceInfraSsids,
		newDatasourceNetworkBasicInternetServices,
		newDatasourceNetworkHostGroups,
		newDatasourceNetworkHosts,
		newDatasourceNetworkImplicitDnsRules,
		newDatasourceNetworkWildcardFqdnCustoms,
		newDatasourceSecurityAntivirusFiletypes,
		newDatasourceSecurityAntivirusProfile,
		newDatasourceSecurityAppCustomSignatures,
		newDatasourceSecurityApplicationCategories,
		newDatasourceSecurityApplicationControlProfile,
		newDatasourceSecurityApplications,
		newDatasourceSecurityDlpDataTypes,
		newDatasourceSecurityDlpDictionaries,
		newDatasourceSecurityDlpExactDataMatches,
		newDatasourceSecurityDlpFilePatterns,
		newDatasourceSecurityDlpFingerprintDatabases,
		newDatasourceSecurityDlpProfile,
		newDatasourceSecurityDlpSensors,
		newDatasourceSecurityDnsFilterProfile,
		newDatasourceSecurityDomainThreatFeeds,
		newDatasourceSecurityEndpointToEndpointPolicies,
		newDatasourceSecurityFileFilterProfile,
		newDatasourceSecurityFortiguardCategories,
		newDatasourceSecurityFortiguardLocalCategories,
		newDatasourceSecurityGeoipCountries,
		newDatasourceSecurityInternalPolicies,
		newDatasourceSecurityInternalReversePolicies,
		newDatasourceSecurityIpThreatFeeds,
		newDatasourceSecurityIpsCustomSignatures,
		newDatasourceSecurityIpsProfile,
		newDatasourceSecurityOnetimeSchedules,
		newDatasourceSecurityOutboundPolicies,
		newDatasourceSecurityProfileGroup,
		newDatasourceSecurityProfileGroups,
		newDatasourceSecurityRecurringSchedules,
		newDatasourceSecurityScheduleGroups,
		newDatasourceSecurityServiceCategories,
		newDatasourceSecurityServiceGroups,
		newDatasourceSecurityServices,
		newDatasourceSecuritySslSshProfile,
		newDatasourceSecurityUrlThreatFeeds,
		newDatasourceSecurityVideoFilterFortiguardCategories,
		newDatasourceSecurityVideoFilterProfile,
		newDatasourceSecurityVideoFilterYoutubeKey,
		newDatasourceSecurityWebFilterProfile,
		newDatasourceUsageAuthFssoAgents,
		newDatasourceUsageAuthLdapServers,
		newDatasourceUsageAuthRadiusServers,
		newDatasourceUsageAuthUserGroups,
		newDatasourceUsageEndpointZtnaTags,
		newDatasourceUsageInfraSsids,
		newDatasourceUsageNetworkHostGroups,
		newDatasourceUsageNetworkHosts,
		newDatasourceUsageSecurityAppCustomSignatures,
		newDatasourceUsageSecurityDlpDictionaries,
		newDatasourceUsageSecurityDlpExactDataMatches,
		newDatasourceUsageSecurityDlpFilePatterns,
		newDatasourceUsageSecurityDlpFingerprintDatabases,
		newDatasourceUsageSecurityDlpSensors,
		newDatasourceUsageSecurityDomainThreatFeeds,
		newDatasourceUsageSecurityEndpointToEndpointPolicies,
		newDatasourceUsageSecurityFortiguardLocalCategories,
		newDatasourceUsageSecurityInternalPolicies,
		newDatasourceUsageSecurityInternalReversePolicies,
		newDatasourceUsageSecurityIpThreatFeeds,
		newDatasourceUsageSecurityIpsCustomSignatures,
		newDatasourceUsageSecurityOnetimeSchedules,
		newDatasourceUsageSecurityOutboundPolicies,
		newDatasourceUsageSecurityProfileGroup,
		newDatasourceUsageSecurityRecurringSchedules,
		newDatasourceUsageSecurityScheduleGroups,
		newDatasourceUsageSecurityServiceGroups,
		newDatasourceUsageSecurityServices,
		newDatasourceUsageSecurityUrlThreatFeeds,
		newDatasourcePrivateAccessNetworkConfiguration,
		newDatasourcePrivateAccessServiceConnections,
		newDatasourceSecurityCertLocalCerts,
		newDatasourceSecurityCertRemoteCaCerts,
		newDatasourceSecurityPkiUsers,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &FortisaseProvider{
			version: version,
		}
	}
}
