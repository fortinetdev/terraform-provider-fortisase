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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceNetworkHostGroups{}

func newResourceNetworkHostGroups() resource.Resource {
	return &resourceNetworkHostGroups{}
}

type resourceNetworkHostGroups struct {
	fortiClient *FortiClient
}

// resourceNetworkHostGroupsModel describes the resource data model.
type resourceNetworkHostGroupsModel struct {
	ID         types.String                            `tfsdk:"id"`
	PrimaryKey types.String                            `tfsdk:"primary_key"`
	Members    []resourceNetworkHostGroupsMembersModel `tfsdk:"members"`
}

func (r *resourceNetworkHostGroups) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_host_groups"
}

func (r *resourceNetworkHostGroups) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthBetween(1, 79),
				},
				Required: true,
			},
			"members": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("network/hosts", "network/host-groups"),
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

func (r *resourceNetworkHostGroups) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceNetworkHostGroups) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceNetworkHostGroupsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectNetworkHostGroups(ctx, diags))
	input_model.URLParams = *(data.getURLObjectNetworkHostGroups(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateNetworkHostGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectNetworkHostGroups(ctx, "read", diags))

	read_output, err := c.ReadNetworkHostGroups(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkHostGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkHostGroups) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceNetworkHostGroupsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceNetworkHostGroupsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectNetworkHostGroups(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectNetworkHostGroups(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateNetworkHostGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectNetworkHostGroups(ctx, "read", diags))

	read_output, err := c.ReadNetworkHostGroups(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkHostGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkHostGroups) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceNetworkHostGroupsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkHostGroups(ctx, "delete", diags))

	err := c.DeleteNetworkHostGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceNetworkHostGroups) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceNetworkHostGroupsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkHostGroups(ctx, "read", diags))

	read_output, err := c.ReadNetworkHostGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshNetworkHostGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkHostGroups) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceNetworkHostGroupsModel) refreshNetworkHostGroups(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["members"]; ok {
		m.Members = m.flattenNetworkHostGroupsMembersList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceNetworkHostGroupsModel) getCreateObjectNetworkHostGroups(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	result["members"] = data.expandNetworkHostGroupsMembersList(ctx, data.Members, diags)

	return &result
}

func (data *resourceNetworkHostGroupsModel) getUpdateObjectNetworkHostGroups(ctx context.Context, state resourceNetworkHostGroupsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if len(data.Members) > 0 || !isSameStruct(data.Members, state.Members) {
		result["members"] = data.expandNetworkHostGroupsMembersList(ctx, data.Members, diags)
	}

	return &result
}

func (data *resourceNetworkHostGroupsModel) getURLObjectNetworkHostGroups(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceNetworkHostGroupsMembersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceNetworkHostGroupsMembersModel) flattenNetworkHostGroupsMembers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceNetworkHostGroupsMembersModel {
	if input == nil {
		return &resourceNetworkHostGroupsMembersModel{}
	}
	if m == nil {
		m = &resourceNetworkHostGroupsMembersModel{}
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

func (s *resourceNetworkHostGroupsModel) flattenNetworkHostGroupsMembersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceNetworkHostGroupsMembersModel {
	if o == nil {
		return []resourceNetworkHostGroupsMembersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument members is not type of []interface{}.", "")
		return []resourceNetworkHostGroupsMembersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceNetworkHostGroupsMembersModel{}
	}

	values := make([]resourceNetworkHostGroupsMembersModel, len(l))
	for i, ele := range l {
		var m resourceNetworkHostGroupsMembersModel
		values[i] = *m.flattenNetworkHostGroupsMembers(ctx, ele, diags)
	}

	return values
}

func (data *resourceNetworkHostGroupsMembersModel) expandNetworkHostGroupsMembers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceNetworkHostGroupsModel) expandNetworkHostGroupsMembersList(ctx context.Context, l []resourceNetworkHostGroupsMembersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandNetworkHostGroupsMembers(ctx, diags)
	}
	return result
}
