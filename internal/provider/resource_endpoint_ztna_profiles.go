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
var _ resource.Resource = &resourceEndpointZtnaProfiles{}

func newResourceEndpointZtnaProfiles() resource.Resource {
	return &resourceEndpointZtnaProfiles{}
}

type resourceEndpointZtnaProfiles struct {
	fortiClient *FortiClient
}

// resourceEndpointZtnaProfilesModel describes the resource data model.
type resourceEndpointZtnaProfilesModel struct {
	ID                   types.String                                       `tfsdk:"id"`
	AllowAutomaticSignOn types.String                                       `tfsdk:"allow_automatic_sign_on"`
	ConnectionRules      []resourceEndpointZtnaProfilesConnectionRulesModel `tfsdk:"connection_rules"`
	EntraId              *resourceEndpointZtnaProfilesEntraIdModel          `tfsdk:"entra_id"`
	PrimaryKey           types.String                                       `tfsdk:"primary_key"`
}

func (r *resourceEndpointZtnaProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_profiles"
}

func (r *resourceEndpointZtnaProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"allow_automatic_sign_on": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				Description: "The primary key of the object. Can be found in the response from the get request.",
				Required:    true,
			},
			"connection_rules": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Float64Attribute{
							Computed: true,
							Optional: true,
						},
						"address": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"uid": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"encryption": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"gateways": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"alias": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"private_app_count": schema.Float64Attribute{
										Computed: true,
										Optional: true,
									},
									"vip": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"redirect": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("enable", "disable"),
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
				},
				Computed: true,
				Optional: true,
			},
			"entra_id": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"application_id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"domain_name": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceEndpointZtnaProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceEndpointZtnaProfiles) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointZtnaProfilesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointZtnaProfiles(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointZtnaProfiles(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateEndpointZtnaProfiles(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectEndpointZtnaProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointZtnaProfiles) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointZtnaProfilesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointZtnaProfilesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointZtnaProfiles(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointZtnaProfiles(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateEndpointZtnaProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointZtnaProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointZtnaProfiles) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointZtnaProfiles) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointZtnaProfilesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointZtnaProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointZtnaProfiles) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointZtnaProfilesModel) refreshEndpointZtnaProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["allowAutomaticSignOn"]; ok {
		m.AllowAutomaticSignOn = parseStringValue(v)
	}

	if v, ok := o["connectionRules"]; ok {
		m.ConnectionRules = m.flattenEndpointZtnaProfilesConnectionRulesList(ctx, v, &diags)
	}

	if v, ok := o["entraId"]; ok {
		m.EntraId = m.EntraId.flattenEndpointZtnaProfilesEntraId(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointZtnaProfilesModel) getCreateObjectEndpointZtnaProfiles(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AllowAutomaticSignOn.IsNull() {
		result["allowAutomaticSignOn"] = data.AllowAutomaticSignOn.ValueString()
	}

	if len(data.ConnectionRules) > 0 {
		result["connectionRules"] = data.expandEndpointZtnaProfilesConnectionRulesList(ctx, data.ConnectionRules, diags)
	}

	if data.EntraId != nil && !isZeroStruct(*data.EntraId) {
		result["entraId"] = data.EntraId.expandEndpointZtnaProfilesEntraId(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointZtnaProfilesModel) getUpdateObjectEndpointZtnaProfiles(ctx context.Context, state resourceEndpointZtnaProfilesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AllowAutomaticSignOn.IsNull() {
		result["allowAutomaticSignOn"] = data.AllowAutomaticSignOn.ValueString()
	}

	if len(data.ConnectionRules) > 0 || !isSameStruct(data.ConnectionRules, state.ConnectionRules) {
		result["connectionRules"] = data.expandEndpointZtnaProfilesConnectionRulesList(ctx, data.ConnectionRules, diags)
	}

	if data.EntraId != nil && !isSameStruct(data.EntraId, state.EntraId) {
		result["entraId"] = data.EntraId.expandEndpointZtnaProfilesEntraId(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointZtnaProfilesModel) getURLObjectEndpointZtnaProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointZtnaProfilesConnectionRulesModel struct {
	Id         types.Float64                                              `tfsdk:"id"`
	Address    types.String                                               `tfsdk:"address"`
	Uid        types.String                                               `tfsdk:"uid"`
	Gateways   []resourceEndpointZtnaProfilesConnectionRulesGatewaysModel `tfsdk:"gateways"`
	Name       types.String                                               `tfsdk:"name"`
	Encryption types.String                                               `tfsdk:"encryption"`
}

type resourceEndpointZtnaProfilesConnectionRulesGatewaysModel struct {
	Alias           types.String  `tfsdk:"alias"`
	PrivateAppCount types.Float64 `tfsdk:"private_app_count"`
	Vip             types.String  `tfsdk:"vip"`
	Redirect        types.String  `tfsdk:"redirect"`
}

type resourceEndpointZtnaProfilesEntraIdModel struct {
	ApplicationId types.String `tfsdk:"application_id"`
	DomainName    types.String `tfsdk:"domain_name"`
}

func (m *resourceEndpointZtnaProfilesConnectionRulesModel) flattenEndpointZtnaProfilesConnectionRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaProfilesConnectionRulesModel {
	if input == nil {
		return &resourceEndpointZtnaProfilesConnectionRulesModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaProfilesConnectionRulesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	if v, ok := o["address"]; ok {
		m.Address = parseStringValue(v)
	}

	if v, ok := o["uid"]; ok {
		m.Uid = parseStringValue(v)
	}

	if v, ok := o["gateways"]; ok {
		m.Gateways = m.flattenEndpointZtnaProfilesConnectionRulesGatewaysList(ctx, v, diags)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["encryption"]; ok {
		m.Encryption = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointZtnaProfilesModel) flattenEndpointZtnaProfilesConnectionRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointZtnaProfilesConnectionRulesModel {
	if o == nil {
		return []resourceEndpointZtnaProfilesConnectionRulesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument connection_rules is not type of []interface{}.", "")
		return []resourceEndpointZtnaProfilesConnectionRulesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointZtnaProfilesConnectionRulesModel{}
	}

	values := make([]resourceEndpointZtnaProfilesConnectionRulesModel, len(l))
	for i, ele := range l {
		var m resourceEndpointZtnaProfilesConnectionRulesModel
		values[i] = *m.flattenEndpointZtnaProfilesConnectionRules(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointZtnaProfilesConnectionRulesGatewaysModel) flattenEndpointZtnaProfilesConnectionRulesGateways(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaProfilesConnectionRulesGatewaysModel {
	if input == nil {
		return &resourceEndpointZtnaProfilesConnectionRulesGatewaysModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaProfilesConnectionRulesGatewaysModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["alias"]; ok {
		m.Alias = parseStringValue(v)
	}

	if v, ok := o["private_app_count"]; ok {
		m.PrivateAppCount = parseFloat64Value(v)
	}

	if v, ok := o["vip"]; ok {
		m.Vip = parseStringValue(v)
	}

	if v, ok := o["redirect"]; ok {
		m.Redirect = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointZtnaProfilesConnectionRulesModel) flattenEndpointZtnaProfilesConnectionRulesGatewaysList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointZtnaProfilesConnectionRulesGatewaysModel {
	if o == nil {
		return []resourceEndpointZtnaProfilesConnectionRulesGatewaysModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument gateways is not type of []interface{}.", "")
		return []resourceEndpointZtnaProfilesConnectionRulesGatewaysModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointZtnaProfilesConnectionRulesGatewaysModel{}
	}

	values := make([]resourceEndpointZtnaProfilesConnectionRulesGatewaysModel, len(l))
	for i, ele := range l {
		var m resourceEndpointZtnaProfilesConnectionRulesGatewaysModel
		values[i] = *m.flattenEndpointZtnaProfilesConnectionRulesGateways(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointZtnaProfilesEntraIdModel) flattenEndpointZtnaProfilesEntraId(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaProfilesEntraIdModel {
	if input == nil {
		return &resourceEndpointZtnaProfilesEntraIdModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaProfilesEntraIdModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["applicationId"]; ok {
		m.ApplicationId = parseStringValue(v)
	}

	if v, ok := o["domainName"]; ok {
		m.DomainName = parseStringValue(v)
	}

	return m
}

func (data *resourceEndpointZtnaProfilesConnectionRulesModel) expandEndpointZtnaProfilesConnectionRules(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() {
		result["id"] = data.Id.ValueFloat64()
	}

	if !data.Address.IsNull() {
		result["address"] = data.Address.ValueString()
	}

	if !data.Uid.IsNull() {
		result["uid"] = data.Uid.ValueString()
	}

	result["gateways"] = data.expandEndpointZtnaProfilesConnectionRulesGatewaysList(ctx, data.Gateways, diags)

	if !data.Name.IsNull() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Encryption.IsNull() {
		result["encryption"] = data.Encryption.ValueString()
	}

	return result
}

func (s *resourceEndpointZtnaProfilesModel) expandEndpointZtnaProfilesConnectionRulesList(ctx context.Context, l []resourceEndpointZtnaProfilesConnectionRulesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointZtnaProfilesConnectionRules(ctx, diags)
	}
	return result
}

func (data *resourceEndpointZtnaProfilesConnectionRulesGatewaysModel) expandEndpointZtnaProfilesConnectionRulesGateways(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Alias.IsNull() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.PrivateAppCount.IsNull() {
		result["private_app_count"] = data.PrivateAppCount.ValueFloat64()
	}

	if !data.Vip.IsNull() {
		result["vip"] = data.Vip.ValueString()
	}

	if !data.Redirect.IsNull() {
		result["redirect"] = data.Redirect.ValueString()
	}

	return result
}

func (s *resourceEndpointZtnaProfilesConnectionRulesModel) expandEndpointZtnaProfilesConnectionRulesGatewaysList(ctx context.Context, l []resourceEndpointZtnaProfilesConnectionRulesGatewaysModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointZtnaProfilesConnectionRulesGateways(ctx, diags)
	}
	return result
}

func (data *resourceEndpointZtnaProfilesEntraIdModel) expandEndpointZtnaProfilesEntraId(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ApplicationId.IsNull() {
		result["applicationId"] = data.ApplicationId.ValueString()
	}

	if !data.DomainName.IsNull() {
		result["domainName"] = data.DomainName.ValueString()
	}

	return result
}
