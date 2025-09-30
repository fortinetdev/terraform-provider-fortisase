// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSecurityPkiUsers{}

func newResourceSecurityPkiUsers() resource.Resource {
	return &resourceSecurityPkiUsers{}
}

type resourceSecurityPkiUsers struct {
	fortiClient *FortiClient
}

// resourceSecurityPkiUsersModel describes the resource data model.
type resourceSecurityPkiUsersModel struct {
	ID             types.String                     `tfsdk:"id"`
	PrimaryKey     types.String                     `tfsdk:"primary_key"`
	Subject        types.String                     `tfsdk:"subject"`
	Ca             *resourceSecurityPkiUsersCaModel `tfsdk:"ca"`
	IsStaticObject types.Bool                       `tfsdk:"is_static_object"`
	References     types.Float64                    `tfsdk:"references"`
	IsGlobalEntry  types.Bool                       `tfsdk:"is_global_entry"`
}

func (r *resourceSecurityPkiUsers) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_pki_users"
}

func (r *resourceSecurityPkiUsers) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Description: "Primary Key of PKI User.",
				Required:    true,
			},
			"subject": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"is_static_object": schema.BoolAttribute{
				Computed: true,
			},
			"references": schema.Float64Attribute{
				Computed: true,
			},
			"is_global_entry": schema.BoolAttribute{
				Computed: true,
			},
			"ca": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "CA Cert Name",
						Computed:    true,
						Optional:    true,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceSecurityPkiUsers) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityPkiUsers) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityPkiUsersModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityPkiUsers(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityPkiUsers(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityPkiUsers(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityPkiUsers(ctx, "read", diags))

	read_output, err := c.ReadSecurityPkiUsers(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityPkiUsers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityPkiUsers) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityPkiUsersModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityPkiUsersModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityPkiUsers(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityPkiUsers(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityPkiUsers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityPkiUsers(ctx, "read", diags))

	read_output, err := c.ReadSecurityPkiUsers(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityPkiUsers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityPkiUsers) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityPkiUsersModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityPkiUsers(ctx, "delete", diags))

	err := c.DeleteSecurityPkiUsers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceSecurityPkiUsers) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityPkiUsersModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityPkiUsers(ctx, "read", diags))

	read_output, err := c.ReadSecurityPkiUsers(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityPkiUsers(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityPkiUsers) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityPkiUsersModel) refreshSecurityPkiUsers(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["subject"]; ok {
		m.Subject = parseStringValue(v)
	}

	if v, ok := o["ca"]; ok {
		m.Ca = m.Ca.flattenSecurityPkiUsersCa(ctx, v, &diags)
	}

	if v, ok := o["isStaticObject"]; ok {
		m.IsStaticObject = parseBoolValue(v)
	}

	if v, ok := o["references"]; ok {
		m.References = parseFloat64Value(v)
	}

	if v, ok := o["isGlobalEntry"]; ok {
		m.IsGlobalEntry = parseBoolValue(v)
	}

	return diags
}

func (data *resourceSecurityPkiUsersModel) getCreateObjectSecurityPkiUsers(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Subject.IsNull() {
		result["subject"] = data.Subject.ValueString()
	}

	if data.Ca != nil && !isZeroStruct(*data.Ca) {
		result["ca"] = data.Ca.expandSecurityPkiUsersCa(ctx, diags)
	}

	return &result
}

func (data *resourceSecurityPkiUsersModel) getUpdateObjectSecurityPkiUsers(ctx context.Context, state resourceSecurityPkiUsersModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Subject.IsNull() && !data.Subject.Equal(state.Subject) {
		result["subject"] = data.Subject.ValueString()
	}

	if data.Ca != nil && !isSameStruct(data.Ca, state.Ca) {
		result["ca"] = data.Ca.expandSecurityPkiUsersCa(ctx, diags)
	}

	return &result
}

func (data *resourceSecurityPkiUsersModel) getURLObjectSecurityPkiUsers(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityPkiUsersCaModel struct {
	Name types.String `tfsdk:"name"`
}

func (m *resourceSecurityPkiUsersCaModel) flattenSecurityPkiUsersCa(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityPkiUsersCaModel {
	if input == nil {
		return &resourceSecurityPkiUsersCaModel{}
	}
	if m == nil {
		m = &resourceSecurityPkiUsersCaModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	return m
}

func (data *resourceSecurityPkiUsersCaModel) expandSecurityPkiUsersCa(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Name.IsNull() {
		result["name"] = data.Name.ValueString()
	}

	return result
}
