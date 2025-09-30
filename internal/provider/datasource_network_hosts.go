// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceNetworkHosts{}

func newDatasourceNetworkHosts() datasource.DataSource {
	return &datasourceNetworkHosts{}
}

type datasourceNetworkHosts struct {
	fortiClient *FortiClient
}

// datasourceNetworkHostsModel describes the datasource data model.
type datasourceNetworkHostsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Type       types.String `tfsdk:"type"`
	Location   types.String `tfsdk:"location"`
	Subnet     types.String `tfsdk:"subnet"`
	StartIp    types.String `tfsdk:"start_ip"`
	EndIp      types.String `tfsdk:"end_ip"`
	Fqdn       types.String `tfsdk:"fqdn"`
	CountryId  types.String `tfsdk:"country_id"`
}

func (r *datasourceNetworkHosts) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hosts"
}

func (r *datasourceNetworkHosts) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 79),
				},
				Required: true,
			},
			"type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("ipmask", "iprange", "fqdn", "geography"),
				},
				Computed: true,
				Optional: true,
			},
			"location": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("internal", "external", "private-access", "unspecified"),
				},
				Computed: true,
				Optional: true,
			},
			"subnet": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"start_ip": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"end_ip": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"fqdn": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 255),
				},
				Computed: true,
				Optional: true,
			},
			"country_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(2, 2),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceNetworkHosts) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceNetworkHosts) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceNetworkHostsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkHosts(ctx, "read", diags))

	read_output, err := c.ReadNetworkHosts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkHosts(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceNetworkHostsModel) refreshNetworkHosts(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["location"]; ok {
		m.Location = parseStringValue(v)
	}

	if v, ok := o["subnet"]; ok {
		m.Subnet = parseStringValue(v)
	}

	if v, ok := o["startIp"]; ok {
		m.StartIp = parseStringValue(v)
	}

	if v, ok := o["endIp"]; ok {
		m.EndIp = parseStringValue(v)
	}

	if v, ok := o["fqdn"]; ok {
		m.Fqdn = parseStringValue(v)
	}

	if v, ok := o["countryId"]; ok {
		m.CountryId = parseStringValue(v)
	}

	return diags
}

func (data *datasourceNetworkHostsModel) getURLObjectNetworkHosts(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
