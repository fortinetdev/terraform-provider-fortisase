// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceSecurityGeoipCountries keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceSecurityGeoipCountries{}

func newDatasourceSecurityGeoipCountries() datasource.DataSource {
	return &datasourceSecurityGeoipCountries{
		datasourceSecurityGeoipCountry: &datasourceSecurityGeoipCountry{},
	}
}

type datasourceSecurityGeoipCountries struct {
	*datasourceSecurityGeoipCountry
}

func (r *datasourceSecurityGeoipCountries) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_geoip_countries"
}

func (r *datasourceSecurityGeoipCountries) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceSecurityGeoipCountry.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_security_geoip_countries is deprecated. Please use fortisase_security_geoip_country instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceSecurityGeoipCountries) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceSecurityGeoipCountry.Configure(ctx, req, resp)
	r.datasourceSecurityGeoipCountry.resourceName = "fortisase_security_geoip_countries"
}
