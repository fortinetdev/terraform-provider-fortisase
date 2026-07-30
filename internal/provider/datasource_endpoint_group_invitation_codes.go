// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// datasourceEndpointGroupInvitationCodes keeps the deprecated Terraform type available while
// reusing the canonical data source implementation.
var _ datasource.DataSource = &datasourceEndpointGroupInvitationCodes{}

func newDatasourceEndpointGroupInvitationCodes() datasource.DataSource {
	return &datasourceEndpointGroupInvitationCodes{
		datasourceEndpointGroupInvitationCode: &datasourceEndpointGroupInvitationCode{},
	}
}

type datasourceEndpointGroupInvitationCodes struct {
	*datasourceEndpointGroupInvitationCode
}

func (r *datasourceEndpointGroupInvitationCodes) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_group_invitation_codes"
}

func (r *datasourceEndpointGroupInvitationCodes) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	r.datasourceEndpointGroupInvitationCode.Schema(ctx, req, resp)
	resp.Schema.DeprecationMessage = "fortisase_endpoint_group_invitation_codes is deprecated. Please use fortisase_endpoint_group_invitation_code instead."
	resp.Schema.MarkdownDescription += "\n" + resp.Schema.DeprecationMessage
}

func (r *datasourceEndpointGroupInvitationCodes) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	r.datasourceEndpointGroupInvitationCode.Configure(ctx, req, resp)
	r.datasourceEndpointGroupInvitationCode.resourceName = "fortisase_endpoint_group_invitation_codes"
}
