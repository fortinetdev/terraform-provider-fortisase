// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceInfraIpamSetting{}

func newResourceInfraIpamSetting() resource.Resource {
	return &resourceInfraIpamSetting{}
}

type resourceInfraIpamSetting struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceInfraIpamSettingModel describes the resource data model.
type resourceInfraIpamSettingModel struct {
	ID         types.String                         `tfsdk:"id"`
	PrimaryKey types.String                         `tfsdk:"primary_key"`
	Pools      []resourceInfraIpamSettingPoolsModel `tfsdk:"pools"`
}

func (r *resourceInfraIpamSetting) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_ipam_setting"
}

func (r *resourceInfraIpamSetting) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("$sase-global"),
				},
				Default:  stringdefault.StaticString("$sase-global"),
				Computed: true,
				Optional: true,
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

func (r *resourceInfraIpamSetting) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceInfraIpamSetting) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceInfraIpamSettingModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectInfraIpamSetting(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateInfraIpamSetting(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output, err := c.ReadInfraIpamSetting(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraIpamSetting(ctx, read_output)...)
	if diags.HasError() {
		return
	}
	data.ID = types.StringValue(fmt.Sprintf("%v", read_output["primaryKey"]))

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraIpamSetting) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceInfraIpamSettingModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceInfraIpamSettingModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectInfraIpamSetting(ctx, state, diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateInfraIpamSetting(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey

	read_output, err := c.ReadInfraIpamSetting(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraIpamSetting(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceInfraIpamSetting) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceInfraIpamSetting) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceInfraIpamSettingModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadInfraIpamSetting(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
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

func (r *resourceInfraIpamSetting) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceInfraIpamSettingModel) refreshInfraIpamSetting(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["pools"]; ok {
		m.Pools = m.flattenInfraIpamSettingPoolsList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceInfraIpamSettingModel) getCreateObjectInfraIpamSetting(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	result["pools"] = data.expandInfraIpamSettingPoolsList(ctx, data.Pools, diags)

	return &result
}

func (data *resourceInfraIpamSettingModel) getUpdateObjectInfraIpamSetting(ctx context.Context, state resourceInfraIpamSettingModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if data.Pools != nil {
		result["pools"] = data.expandInfraIpamSettingPoolsList(ctx, data.Pools, diags)
	}

	return &result
}

type resourceInfraIpamSettingPoolsModel struct {
	Name            types.String                                        `tfsdk:"name"`
	Subnet          types.String                                        `tfsdk:"subnet"`
	ExcludedSubnets []resourceInfraIpamSettingPoolsExcludedSubnetsModel `tfsdk:"excluded_subnets"`
}

type resourceInfraIpamSettingPoolsExcludedSubnetsModel struct {
	Subnet types.String `tfsdk:"subnet"`
}

func (m *resourceInfraIpamSettingPoolsModel) flattenInfraIpamSettingPools(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceInfraIpamSettingPoolsModel {
	if input == nil {
		return &resourceInfraIpamSettingPoolsModel{}
	}
	if m == nil {
		m = &resourceInfraIpamSettingPoolsModel{}
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

func (s *resourceInfraIpamSettingModel) flattenInfraIpamSettingPoolsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceInfraIpamSettingPoolsModel {
	if o == nil {
		return []resourceInfraIpamSettingPoolsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument pools is not type of []interface{}.", "")
		return []resourceInfraIpamSettingPoolsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceInfraIpamSettingPoolsModel{}
	}

	values := make([]resourceInfraIpamSettingPoolsModel, len(l))
	for i, ele := range l {
		var m resourceInfraIpamSettingPoolsModel
		values[i] = *m.flattenInfraIpamSettingPools(ctx, ele, diags)
	}

	return values
}

func (m *resourceInfraIpamSettingPoolsExcludedSubnetsModel) flattenInfraIpamSettingPoolsExcludedSubnets(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceInfraIpamSettingPoolsExcludedSubnetsModel {
	if input == nil {
		return &resourceInfraIpamSettingPoolsExcludedSubnetsModel{}
	}
	if m == nil {
		m = &resourceInfraIpamSettingPoolsExcludedSubnetsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["subnet"]; ok {
		m.Subnet = parseStringValue(v)
	}

	return m
}

func (s *resourceInfraIpamSettingPoolsModel) flattenInfraIpamSettingPoolsExcludedSubnetsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceInfraIpamSettingPoolsExcludedSubnetsModel {
	if o == nil {
		return []resourceInfraIpamSettingPoolsExcludedSubnetsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument excluded_subnets is not type of []interface{}.", "")
		return []resourceInfraIpamSettingPoolsExcludedSubnetsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceInfraIpamSettingPoolsExcludedSubnetsModel{}
	}

	values := make([]resourceInfraIpamSettingPoolsExcludedSubnetsModel, len(l))
	for i, ele := range l {
		var m resourceInfraIpamSettingPoolsExcludedSubnetsModel
		values[i] = *m.flattenInfraIpamSettingPoolsExcludedSubnets(ctx, ele, diags)
	}

	return values
}

func (data *resourceInfraIpamSettingPoolsModel) expandInfraIpamSettingPools(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Name.IsNull() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Subnet.IsNull() {
		result["subnet"] = data.Subnet.ValueString()
	}

	result["excludedSubnets"] = data.expandInfraIpamSettingPoolsExcludedSubnetsList(ctx, data.ExcludedSubnets, diags)

	return result
}

func (s *resourceInfraIpamSettingModel) expandInfraIpamSettingPoolsList(ctx context.Context, l []resourceInfraIpamSettingPoolsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandInfraIpamSettingPools(ctx, diags)
	}
	return result
}

func (data *resourceInfraIpamSettingPoolsExcludedSubnetsModel) expandInfraIpamSettingPoolsExcludedSubnets(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Subnet.IsNull() {
		result["subnet"] = data.Subnet.ValueString()
	}

	return result
}

func (s *resourceInfraIpamSettingPoolsModel) expandInfraIpamSettingPoolsExcludedSubnetsList(ctx context.Context, l []resourceInfraIpamSettingPoolsExcludedSubnetsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandInfraIpamSettingPoolsExcludedSubnets(ctx, diags)
	}
	return result
}
