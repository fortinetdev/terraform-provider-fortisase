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
var _ resource.Resource = &resourceSecurityServiceGroups{}

func newResourceSecurityServiceGroups() resource.Resource {
	return &resourceSecurityServiceGroups{}
}

type resourceSecurityServiceGroups struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityServiceGroupsModel describes the resource data model.
type resourceSecurityServiceGroupsModel struct {
	ID         types.String                                `tfsdk:"id"`
	PrimaryKey types.String                                `tfsdk:"primary_key"`
	Proxy      types.Bool                                  `tfsdk:"proxy"`
	Members    []resourceSecurityServiceGroupsMembersModel `tfsdk:"members"`
}

func (r *resourceSecurityServiceGroups) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_service_groups"
}

func (r *resourceSecurityServiceGroups) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthAtMost(79),
				},
				Required: true,
			},
			"proxy": schema.BoolAttribute{
				Computed: true,
				Optional: true,
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
								stringvalidator.OneOf("security/services", "security/service-groups"),
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

func (r *resourceSecurityServiceGroups) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_service_groups"
}

func (r *resourceSecurityServiceGroups) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityServiceGroups")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityServiceGroupsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityServiceGroups(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityServiceGroups(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityServiceGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityServiceGroups(ctx, "read", diags))

	read_output, err := c.ReadSecurityServiceGroups(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityServiceGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityServiceGroups) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityServiceGroups")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityServiceGroupsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityServiceGroupsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityServiceGroups(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityServiceGroups(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityServiceGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityServiceGroups(ctx, "read", diags))

	read_output, err := c.ReadSecurityServiceGroups(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityServiceGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityServiceGroups) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityServiceGroups")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityServiceGroupsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityServiceGroups(ctx, "delete", diags))

	output, err := c.DeleteSecurityServiceGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityServiceGroups) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityServiceGroupsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityServiceGroups(ctx, "read", diags))

	read_output, err := c.ReadSecurityServiceGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityServiceGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityServiceGroups) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityServiceGroupsModel) refreshSecurityServiceGroups(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["proxy"]; ok {
		m.Proxy = parseBoolValue(v)
	}

	if v, ok := o["members"]; ok {
		m.Members = m.flattenSecurityServiceGroupsMembersList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityServiceGroupsModel) getCreateObjectSecurityServiceGroups(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Proxy.IsNull() {
		result["proxy"] = data.Proxy.ValueBool()
	}

	result["members"] = data.expandSecurityServiceGroupsMembersList(ctx, data.Members, diags)

	return &result
}

func (data *resourceSecurityServiceGroupsModel) getUpdateObjectSecurityServiceGroups(ctx context.Context, state resourceSecurityServiceGroupsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Proxy.IsNull() {
		result["proxy"] = data.Proxy.ValueBool()
	}

	if data.Members != nil {
		result["members"] = data.expandSecurityServiceGroupsMembersList(ctx, data.Members, diags)
	}

	return &result
}

func (data *resourceSecurityServiceGroupsModel) getURLObjectSecurityServiceGroups(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityServiceGroupsMembersModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityServiceGroupsMembersModel) flattenSecurityServiceGroupsMembers(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityServiceGroupsMembersModel {
	if input == nil {
		return &resourceSecurityServiceGroupsMembersModel{}
	}
	if m == nil {
		m = &resourceSecurityServiceGroupsMembersModel{}
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

func (s *resourceSecurityServiceGroupsModel) flattenSecurityServiceGroupsMembersList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityServiceGroupsMembersModel {
	if o == nil {
		return []resourceSecurityServiceGroupsMembersModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument members is not type of []interface{}.", "")
		return []resourceSecurityServiceGroupsMembersModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityServiceGroupsMembersModel{}
	}

	values := make([]resourceSecurityServiceGroupsMembersModel, len(l))
	for i, ele := range l {
		var m resourceSecurityServiceGroupsMembersModel
		values[i] = *m.flattenSecurityServiceGroupsMembers(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityServiceGroupsMembersModel) expandSecurityServiceGroupsMembers(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityServiceGroupsModel) expandSecurityServiceGroupsMembersList(ctx context.Context, l []resourceSecurityServiceGroupsMembersModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityServiceGroupsMembers(ctx, diags)
	}
	return result
}
