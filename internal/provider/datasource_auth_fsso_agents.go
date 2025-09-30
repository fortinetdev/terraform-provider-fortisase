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
var _ datasource.DataSource = &datasourceAuthFssoAgents{}

func newDatasourceAuthFssoAgents() datasource.DataSource {
	return &datasourceAuthFssoAgents{}
}

type datasourceAuthFssoAgents struct {
	fortiClient *FortiClient
}

// datasourceAuthFssoAgentsModel describes the datasource data model.
type datasourceAuthFssoAgentsModel struct {
	PrimaryKey     types.String `tfsdk:"primary_key"`
	ActiveServer   types.String `tfsdk:"active_server"`
	Status         types.String `tfsdk:"status"`
	Name           types.String `tfsdk:"name"`
	Server         types.String `tfsdk:"server"`
	Server2        types.String `tfsdk:"server2"`
	Server3        types.String `tfsdk:"server3"`
	Server4        types.String `tfsdk:"server4"`
	Server5        types.String `tfsdk:"server5"`
	SslTrustedCert types.String `tfsdk:"ssl_trusted_cert"`
}

func (r *datasourceAuthFssoAgents) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auth_fsso_agents"
}

func (r *datasourceAuthFssoAgents) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 35),
				},
				Required: true,
			},
			"active_server": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("connected", "disconnected"),
				},
				Computed: true,
				Optional: true,
			},
			"name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 35),
				},
				Computed: true,
				Optional: true,
			},
			"server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"server2": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"server3": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"server4": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"server5": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(63),
				},
				Computed: true,
				Optional: true,
			},
			"ssl_trusted_cert": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(79),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceAuthFssoAgents) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceAuthFssoAgents) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceAuthFssoAgentsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectAuthFssoAgents(ctx, "read", diags))

	read_output, err := c.ReadAuthFssoAgents(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshAuthFssoAgents(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceAuthFssoAgentsModel) refreshAuthFssoAgents(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["activeServer"]; ok {
		m.ActiveServer = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["server"]; ok {
		m.Server = parseStringValue(v)
	}

	if v, ok := o["server2"]; ok {
		m.Server2 = parseStringValue(v)
	}

	if v, ok := o["server3"]; ok {
		m.Server3 = parseStringValue(v)
	}

	if v, ok := o["server4"]; ok {
		m.Server4 = parseStringValue(v)
	}

	if v, ok := o["server5"]; ok {
		m.Server5 = parseStringValue(v)
	}

	if v, ok := o["sslTrustedCert"]; ok {
		m.SslTrustedCert = parseStringValue(v)
	}

	return diags
}

func (data *datasourceAuthFssoAgentsModel) getURLObjectAuthFssoAgents(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
