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
var _ datasource.DataSource = &datasourceInfraIpamSetting{}

func newDatasourceInfraIpamSetting() datasource.DataSource {
	return &datasourceInfraIpamSetting{}
}

type datasourceInfraIpamSetting struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceInfraIpamSettingModel describes the datasource data model.
type datasourceInfraIpamSettingModel struct {
	PrimaryKey types.String                           `tfsdk:"primary_key"`
	Pools      []datasourceInfraIpamSettingPoolsModel `tfsdk:"pools"`
}

func (r *datasourceInfraIpamSetting) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_ipam_setting"
}

func (r *datasourceInfraIpamSetting) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("$sase-global"),
				},
				Required: true,
			},
			"pools": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"subnet": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"excluded_subnets": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"subnet": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
								},
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

func (r *datasourceInfraIpamSetting) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_infra_ipam_setting"
}

func (r *datasourceInfraIpamSetting) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceInfraIpamSettingModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadInfraIpamSetting(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraIpamSetting(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceInfraIpamSettingModel) refreshInfraIpamSetting(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["pools"]; ok {
		m.Pools = m.flattenInfraIpamSettingPoolsList(ctx, v, &diags)
	}

	return diags
}

type datasourceInfraIpamSettingPoolsModel struct {
	Name            types.String                                          `tfsdk:"name"`
	Subnet          types.String                                          `tfsdk:"subnet"`
	ExcludedSubnets []datasourceInfraIpamSettingPoolsExcludedSubnetsModel `tfsdk:"excluded_subnets"`
}

type datasourceInfraIpamSettingPoolsExcludedSubnetsModel struct {
	Subnet types.String `tfsdk:"subnet"`
}

func (m *datasourceInfraIpamSettingPoolsModel) flattenInfraIpamSettingPools(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceInfraIpamSettingPoolsModel {
	if input == nil {
		return &datasourceInfraIpamSettingPoolsModel{}
	}
	if m == nil {
		m = &datasourceInfraIpamSettingPoolsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["subnet"]; ok {
		m.Subnet = parseStringValue(v)
	}

	if v, ok := o["excludedSubnets"]; ok {
		m.ExcludedSubnets = m.flattenInfraIpamSettingPoolsExcludedSubnetsList(ctx, v, diags)
	}

	return m
}

func (s *datasourceInfraIpamSettingModel) flattenInfraIpamSettingPoolsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceInfraIpamSettingPoolsModel {
	if o == nil {
		return []datasourceInfraIpamSettingPoolsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument pools is not type of []interface{}.", "")
		return []datasourceInfraIpamSettingPoolsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceInfraIpamSettingPoolsModel{}
	}

	values := make([]datasourceInfraIpamSettingPoolsModel, len(l))
	for i, ele := range l {
		var m datasourceInfraIpamSettingPoolsModel
		values[i] = *m.flattenInfraIpamSettingPools(ctx, ele, diags)
	}

	return values
}

func (m *datasourceInfraIpamSettingPoolsExcludedSubnetsModel) flattenInfraIpamSettingPoolsExcludedSubnets(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceInfraIpamSettingPoolsExcludedSubnetsModel {
	if input == nil {
		return &datasourceInfraIpamSettingPoolsExcludedSubnetsModel{}
	}
	if m == nil {
		m = &datasourceInfraIpamSettingPoolsExcludedSubnetsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["subnet"]; ok {
		m.Subnet = parseStringValue(v)
	}

	return m
}

func (s *datasourceInfraIpamSettingPoolsModel) flattenInfraIpamSettingPoolsExcludedSubnetsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceInfraIpamSettingPoolsExcludedSubnetsModel {
	if o == nil {
		return []datasourceInfraIpamSettingPoolsExcludedSubnetsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument excluded_subnets is not type of []interface{}.", "")
		return []datasourceInfraIpamSettingPoolsExcludedSubnetsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceInfraIpamSettingPoolsExcludedSubnetsModel{}
	}

	values := make([]datasourceInfraIpamSettingPoolsExcludedSubnetsModel, len(l))
	for i, ele := range l {
		var m datasourceInfraIpamSettingPoolsExcludedSubnetsModel
		values[i] = *m.flattenInfraIpamSettingPoolsExcludedSubnets(ctx, ele, diags)
	}

	return values
}
