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
var _ datasource.DataSource = &datasourceSecurityFileFilterProfile{}

func newDatasourceSecurityFileFilterProfile() datasource.DataSource {
	return &datasourceSecurityFileFilterProfile{}
}

type datasourceSecurityFileFilterProfile struct {
	fortiClient *FortiClient
}

// datasourceSecurityFileFilterProfileModel describes the datasource data model.
type datasourceSecurityFileFilterProfileModel struct {
	PrimaryKey                  types.String                                      `tfsdk:"primary_key"`
	Block                       []datasourceSecurityFileFilterProfileBlockModel   `tfsdk:"block"`
	Monitor                     []datasourceSecurityFileFilterProfileMonitorModel `tfsdk:"monitor"`
	BlockPasswordProtectedFiles types.Bool                                        `tfsdk:"block_password_protected_files"`
	Direction                   types.String                                      `tfsdk:"direction"`
}

func (r *datasourceSecurityFileFilterProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_file_filter_profile"
}

func (r *datasourceSecurityFileFilterProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"block_password_protected_files": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"direction": schema.StringAttribute{
				Description: "The direction of the target resource.",
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				Computed: true,
				Optional: true,
			},
			"block": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("security/antivirus-filetypes"),
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"monitor": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("security/antivirus-filetypes"),
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityFileFilterProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityFileFilterProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityFileFilterProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityFileFilterProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityFileFilterProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityFileFilterProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityFileFilterProfileModel) refreshSecurityFileFilterProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["block"]; ok {
		m.Block = m.flattenSecurityFileFilterProfileBlockList(ctx, v, &diags)
	}

	if v, ok := o["monitor"]; ok {
		m.Monitor = m.flattenSecurityFileFilterProfileMonitorList(ctx, v, &diags)
	}

	if v, ok := o["blockPasswordProtectedFiles"]; ok {
		m.BlockPasswordProtectedFiles = parseBoolValue(v)
	}

	return diags
}

func (data *datasourceSecurityFileFilterProfileModel) getURLObjectSecurityFileFilterProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityFileFilterProfileBlockModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityFileFilterProfileMonitorModel struct {
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityFileFilterProfileBlockModel) flattenSecurityFileFilterProfileBlock(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityFileFilterProfileBlockModel {
	if input == nil {
		return &datasourceSecurityFileFilterProfileBlockModel{}
	}
	if m == nil {
		m = &datasourceSecurityFileFilterProfileBlockModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *datasourceSecurityFileFilterProfileModel) flattenSecurityFileFilterProfileBlockList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityFileFilterProfileBlockModel {
	if o == nil {
		return []datasourceSecurityFileFilterProfileBlockModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument block is not type of []interface{}.", "")
		return []datasourceSecurityFileFilterProfileBlockModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityFileFilterProfileBlockModel{}
	}

	values := make([]datasourceSecurityFileFilterProfileBlockModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityFileFilterProfileBlockModel
		values[i] = *m.flattenSecurityFileFilterProfileBlock(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityFileFilterProfileMonitorModel) flattenSecurityFileFilterProfileMonitor(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityFileFilterProfileMonitorModel {
	if input == nil {
		return &datasourceSecurityFileFilterProfileMonitorModel{}
	}
	if m == nil {
		m = &datasourceSecurityFileFilterProfileMonitorModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *datasourceSecurityFileFilterProfileModel) flattenSecurityFileFilterProfileMonitorList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityFileFilterProfileMonitorModel {
	if o == nil {
		return []datasourceSecurityFileFilterProfileMonitorModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument monitor is not type of []interface{}.", "")
		return []datasourceSecurityFileFilterProfileMonitorModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityFileFilterProfileMonitorModel{}
	}

	values := make([]datasourceSecurityFileFilterProfileMonitorModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityFileFilterProfileMonitorModel
		values[i] = *m.flattenSecurityFileFilterProfileMonitor(ctx, ele, diags)
	}

	return values
}
