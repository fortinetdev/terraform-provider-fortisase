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
var _ datasource.DataSource = &datasourceUsageSecurityIpThreatFeeds{}

func newDatasourceUsageSecurityIpThreatFeeds() datasource.DataSource {
	return &datasourceUsageSecurityIpThreatFeeds{}
}

type datasourceUsageSecurityIpThreatFeeds struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceUsageSecurityIpThreatFeedsModel describes the datasource data model.
type datasourceUsageSecurityIpThreatFeedsModel struct {
	Type       types.String  `tfsdk:"type"`
	Ftntcount  types.Float64 `tfsdk:"ftntcount"`
	PrimaryKey types.String  `tfsdk:"primary_key"`
}

func (r *datasourceUsageSecurityIpThreatFeeds) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_usage_security_ip_threat_feeds"
}

func (r *datasourceUsageSecurityIpThreatFeeds) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ftntcount": schema.Float64Attribute{
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
			},
		},
	}
}

func (r *datasourceUsageSecurityIpThreatFeeds) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_usage_security_ip_threat_feeds"
}

func (r *datasourceUsageSecurityIpThreatFeeds) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceUsageSecurityIpThreatFeedsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectUsageSecurityIpThreatFeeds(ctx, "read", diags))

	read_output, err := c.ReadUsageSecurityIpThreatFeeds(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshUsageSecurityIpThreatFeeds(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceUsageSecurityIpThreatFeedsModel) refreshUsageSecurityIpThreatFeeds(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["count"]; ok {
		m.Ftntcount = parseFloat64Value(v)
	}

	return diags
}

func (data *datasourceUsageSecurityIpThreatFeedsModel) getURLObjectUsageSecurityIpThreatFeeds(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
