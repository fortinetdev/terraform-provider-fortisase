// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityBotnetDomainsStat2Edl{}

func newDatasourceSecurityBotnetDomainsStat() datasource.DataSource {
	return &datasourceSecurityBotnetDomainsStat2Edl{}
}

type datasourceSecurityBotnetDomainsStat2Edl struct {
	fortiClient *FortiClient
}

// datasourceSecurityBotnetDomainsStat2EdlModel describes the datasource data model.
type datasourceSecurityBotnetDomainsStat2EdlModel struct {
	TotalEntries types.Float64 `tfsdk:"total_entries"`
}

func (r *datasourceSecurityBotnetDomainsStat2Edl) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_botnet_domains_stat"
}

func (r *datasourceSecurityBotnetDomainsStat2Edl) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"total_entries": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityBotnetDomainsStat2Edl) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Always perform a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*FortiClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *FortiClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.fortiClient = client
}

func (r *datasourceSecurityBotnetDomainsStat2Edl) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityBotnetDomainsStat2EdlModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := "SecurityBotnetDomainsStat"

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadSecurityBotnetDomainsStat(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityBotnetDomainsStat2Edl(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityBotnetDomainsStat2EdlModel) refreshSecurityBotnetDomainsStat2Edl(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["totalEntries"]; ok {
		m.TotalEntries = parseFloat64Value(v)
	}

	return diags
}
