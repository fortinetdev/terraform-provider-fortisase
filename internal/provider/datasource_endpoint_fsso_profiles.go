// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointFssoProfiles{}

func newDatasourceEndpointFssoProfiles() datasource.DataSource {
	return &datasourceEndpointFssoProfiles{}
}

type datasourceEndpointFssoProfiles struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointFssoProfilesModel describes the datasource data model.
type datasourceEndpointFssoProfilesModel struct {
	Enabled       types.Bool    `tfsdk:"enabled"`
	PreferEntraId types.String  `tfsdk:"prefer_entra_id"`
	Host          types.String  `tfsdk:"host"`
	Port          types.Float64 `tfsdk:"port"`
	PreSharedKey  types.String  `tfsdk:"pre_shared_key"`
	PrimaryKey    types.String  `tfsdk:"primary_key"`
}

func (r *datasourceEndpointFssoProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_fsso_profiles"
}

func (r *datasourceEndpointFssoProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"prefer_entra_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"host": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"port": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(65535),
				},
				Computed: true,
				Optional: true,
			},
			"pre_shared_key": schema.StringAttribute{
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

func (r *datasourceEndpointFssoProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_fsso_profiles"
}

func (r *datasourceEndpointFssoProfiles) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointFssoProfilesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointFssoProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointFssoProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointFssoProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointFssoProfilesModel) refreshEndpointFssoProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["preferEntraId"]; ok {
		m.PreferEntraId = parseStringValue(v)
	}

	if v, ok := o["host"]; ok {
		m.Host = parseStringValue(v)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	if v, ok := o["preSharedKey"]; ok {
		m.PreSharedKey = parseStringValue(v)
	}

	return diags
}

func (data *datasourceEndpointFssoProfilesModel) getURLObjectEndpointFssoProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
