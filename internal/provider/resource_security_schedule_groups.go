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
var _ resource.Resource = &resourceSecurityScheduleGroups{}

func newResourceSecurityScheduleGroups() resource.Resource {
	return &resourceSecurityScheduleGroups{}
}

type resourceSecurityScheduleGroups struct {
	fortiClient *FortiClient
}

// resourceSecurityScheduleGroupsModel describes the resource data model.
type resourceSecurityScheduleGroupsModel struct {
	ID         types.String                                 `tfsdk:"id"`
	PrimaryKey types.String                                 `tfsdk:"primary_key"`
	Members    []resourceSecurityScheduleGroupsMembersModel `tfsdk:"members"`
}

func (r *resourceSecurityScheduleGroups) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_schedule_groups"
}

func (r *resourceSecurityScheduleGroups) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthBetween(1, 31),
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
								stringvalidator.OneOf("security/onetime-schedules", "security/recurring-schedules"),
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

func (r *resourceSecurityScheduleGroups) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityScheduleGroups) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityScheduleGroupsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityScheduleGroups(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityScheduleGroups(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityScheduleGroups(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityScheduleGroups(ctx, "read", diags))

	read_output, err := c.ReadSecurityScheduleGroups(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityScheduleGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityScheduleGroups) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityScheduleGroupsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityScheduleGroupsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityScheduleGroups(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityScheduleGroups(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityScheduleGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityScheduleGroups(ctx, "read", diags))

	read_output, err := c.ReadSecurityScheduleGroups(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityScheduleGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityScheduleGroups) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityScheduleGroupsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityScheduleGroups(ctx, "delete", diags))

	err := c.DeleteSecurityScheduleGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceSecurityScheduleGroups) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityScheduleGroupsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityScheduleGroups(ctx, "read", diags))

	read_output, err := c.ReadSecurityScheduleGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityScheduleGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityScheduleGroups) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityScheduleGroupsModel) refreshSecurityScheduleGroups(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["members"]; ok {
		m.Members = m.flattenSecurityScheduleGroupsMembersList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityScheduleGroupsModel) getCreateObjectSecurityScheduleGroups(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	result["members"] = data.expandSecurityScheduleGroupsMembersList(ctx, data.Members, diags)

	return &result
}

func (data *resourceSecurityScheduleGroupsModel) getUpdateObjectSecurityScheduleGroups(ctx context.Context, state resourceSecurityScheduleGroupsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if len(data.Members) > 0 || !isSameStruct(data.Members, state.Members) {
		result["members"] = data.expandSecurityScheduleGroupsMembersList(ctx, data.Members, diags)
	}

	return &result
}

func (data *resourceSecurityScheduleGroupsModel) getURLObjectSecurityScheduleGroups(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityScheduleGroupsMembersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityScheduleGroupsMembersModel) flattenSecurityScheduleGroupsMembers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityScheduleGroupsMembersModel {
	if input == nil {
		return &resourceSecurityScheduleGroupsMembersModel{}
	}
	if m == nil {
		m = &resourceSecurityScheduleGroupsMembersModel{}
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

func (s *resourceSecurityScheduleGroupsModel) flattenSecurityScheduleGroupsMembersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityScheduleGroupsMembersModel {
	if o == nil {
		return []resourceSecurityScheduleGroupsMembersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument members is not type of []interface{}.", "")
		return []resourceSecurityScheduleGroupsMembersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityScheduleGroupsMembersModel{}
	}

	values := make([]resourceSecurityScheduleGroupsMembersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityScheduleGroupsMembersModel
		values[i] = *m.flattenSecurityScheduleGroupsMembers(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityScheduleGroupsMembersModel) expandSecurityScheduleGroupsMembers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityScheduleGroupsModel) expandSecurityScheduleGroupsMembersList(ctx context.Context, l []resourceSecurityScheduleGroupsMembersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityScheduleGroupsMembers(ctx, diags)
	}
	return result
}
