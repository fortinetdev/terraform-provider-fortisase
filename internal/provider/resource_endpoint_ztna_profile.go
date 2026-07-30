// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
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
var _ resource.Resource = &resourceEndpointZtnaProfile{}
var _ resource.ResourceWithMoveState = &resourceEndpointZtnaProfile{}

func newResourceEndpointZtnaProfile() resource.Resource {
	return &resourceEndpointZtnaProfile{}
}

type resourceEndpointZtnaProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointZtnaProfileModel describes the resource data model.
type resourceEndpointZtnaProfileModel struct {
	ID                   types.String                                      `tfsdk:"id"`
	AllowAutomaticSignOn types.String                                      `tfsdk:"allow_automatic_sign_on"`
	Status               types.String                                      `tfsdk:"status"`
	ConnectionRules      []resourceEndpointZtnaProfileConnectionRulesModel `tfsdk:"connection_rules"`
	EntraId              *resourceEndpointZtnaProfileEntraIdModel          `tfsdk:"entra_id"`
	PrimaryKey           types.String                                      `tfsdk:"primary_key"`
}

func (r *resourceEndpointZtnaProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_profile"
}

func (r *resourceEndpointZtnaProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ZTNA Profile Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
						"mask": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"port": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"encryption": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"gateways": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.Float64Attribute{
										Computed: true,
										Optional: true,
									},
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
											stringvalidatorwarning.OneOf("enable", "disable"),
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

func (r *resourceEndpointZtnaProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_ztna_profile"
}
func (r *resourceEndpointZtnaProfile) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_endpoint_ztna_profiles" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceEndpointZtnaProfileModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceEndpointZtnaProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointZtnaProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointZtnaProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointZtnaProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateEndpointZtnaProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := data.PrimaryKey.ValueString()
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointZtnaProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointZtnaProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointZtnaProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointZtnaProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointZtnaProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointZtnaProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointZtnaProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointZtnaProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointZtnaProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointZtnaProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointZtnaProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointZtnaProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaProfiles(&input_model)
	if err != nil {
		if isNotFoundResponse(read_output) {
			resp.State.RemoveResource(ctx)
			return
		}
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointZtnaProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceEndpointZtnaProfileModel) refreshEndpointZtnaProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["allowAutomaticSignOn"]; ok {
		m.AllowAutomaticSignOn = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["connectionRules"]; ok {
		m.ConnectionRules = m.flattenEndpointZtnaProfileConnectionRulesList(ctx, v, &diags)
	}

	if v, ok := o["entraId"]; ok {
		m.EntraId = m.EntraId.flattenEndpointZtnaProfileEntraId(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointZtnaProfileModel) getCreateObjectEndpointZtnaProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AllowAutomaticSignOn.IsNull() && !data.AllowAutomaticSignOn.IsUnknown() {
		result["allowAutomaticSignOn"] = data.AllowAutomaticSignOn.ValueString()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if data.ConnectionRules != nil {
		result["connectionRules"] = data.expandEndpointZtnaProfileConnectionRulesList(ctx, data.ConnectionRules, diags)
	}

	if data.EntraId != nil && !isZeroStruct(*data.EntraId) {
		result["entraId"] = data.EntraId.expandEndpointZtnaProfileEntraId(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointZtnaProfileModel) getUpdateObjectEndpointZtnaProfile(ctx context.Context, state resourceEndpointZtnaProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AllowAutomaticSignOn.IsNull() && !data.AllowAutomaticSignOn.IsUnknown() {
		result["allowAutomaticSignOn"] = data.AllowAutomaticSignOn.ValueString()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	if data.ConnectionRules != nil {
		result["connectionRules"] = data.expandEndpointZtnaProfileConnectionRulesList(ctx, data.ConnectionRules, diags)
	}

	if data.EntraId != nil {
		result["entraId"] = data.EntraId.expandEndpointZtnaProfileEntraId(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointZtnaProfileModel) getURLObjectEndpointZtnaProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointZtnaProfileConnectionRulesModel struct {
	Id         types.Float64                                             `tfsdk:"id"`
	Address    types.String                                              `tfsdk:"address"`
	Uid        types.String                                              `tfsdk:"uid"`
	Gateways   []resourceEndpointZtnaProfileConnectionRulesGatewaysModel `tfsdk:"gateways"`
	Mask       types.String                                              `tfsdk:"mask"`
	Port       types.String                                              `tfsdk:"port"`
	Name       types.String                                              `tfsdk:"name"`
	Encryption types.String                                              `tfsdk:"encryption"`
}

type resourceEndpointZtnaProfileConnectionRulesGatewaysModel struct {
	Id              types.Float64 `tfsdk:"id"`
	Alias           types.String  `tfsdk:"alias"`
	PrivateAppCount types.Float64 `tfsdk:"private_app_count"`
	Vip             types.String  `tfsdk:"vip"`
	Redirect        types.String  `tfsdk:"redirect"`
}

type resourceEndpointZtnaProfileEntraIdModel struct {
	ApplicationId types.String `tfsdk:"application_id"`
	DomainName    types.String `tfsdk:"domain_name"`
}

func (m *resourceEndpointZtnaProfileConnectionRulesModel) flattenEndpointZtnaProfileConnectionRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaProfileConnectionRulesModel {
	if input == nil {
		return &resourceEndpointZtnaProfileConnectionRulesModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaProfileConnectionRulesModel{}
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
		m.Gateways = m.flattenEndpointZtnaProfileConnectionRulesGatewaysList(ctx, v, diags)
	}

	if v, ok := o["mask"]; ok {
		m.Mask = parseStringValue(v)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["encryption"]; ok {
		m.Encryption = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointZtnaProfileModel) flattenEndpointZtnaProfileConnectionRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointZtnaProfileConnectionRulesModel {
	if o == nil {
		return []resourceEndpointZtnaProfileConnectionRulesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument connection_rules is not type of []interface{}.", "")
		return []resourceEndpointZtnaProfileConnectionRulesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointZtnaProfileConnectionRulesModel{}
	}

	values := make([]resourceEndpointZtnaProfileConnectionRulesModel, len(l))
	for i, ele := range l {
		var m resourceEndpointZtnaProfileConnectionRulesModel
		if i < len(s.ConnectionRules) {
			m = s.ConnectionRules[i]
		}
		values[i] = *m.flattenEndpointZtnaProfileConnectionRules(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointZtnaProfileConnectionRulesGatewaysModel) flattenEndpointZtnaProfileConnectionRulesGateways(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaProfileConnectionRulesGatewaysModel {
	if input == nil {
		return &resourceEndpointZtnaProfileConnectionRulesGatewaysModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaProfileConnectionRulesGatewaysModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

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

func (s *resourceEndpointZtnaProfileConnectionRulesModel) flattenEndpointZtnaProfileConnectionRulesGatewaysList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointZtnaProfileConnectionRulesGatewaysModel {
	if o == nil {
		return []resourceEndpointZtnaProfileConnectionRulesGatewaysModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument gateways is not type of []interface{}.", "")
		return []resourceEndpointZtnaProfileConnectionRulesGatewaysModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointZtnaProfileConnectionRulesGatewaysModel{}
	}

	values := make([]resourceEndpointZtnaProfileConnectionRulesGatewaysModel, len(l))
	for i, ele := range l {
		var m resourceEndpointZtnaProfileConnectionRulesGatewaysModel
		if i < len(s.Gateways) {
			m = s.Gateways[i]
		}
		values[i] = *m.flattenEndpointZtnaProfileConnectionRulesGateways(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointZtnaProfileEntraIdModel) flattenEndpointZtnaProfileEntraId(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaProfileEntraIdModel {
	if input == nil {
		return &resourceEndpointZtnaProfileEntraIdModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaProfileEntraIdModel{}
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

func (data *resourceEndpointZtnaProfileConnectionRulesModel) expandEndpointZtnaProfileConnectionRules(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		result["id"] = data.Id.ValueFloat64()
	}

	if !data.Address.IsNull() && !data.Address.IsUnknown() {
		result["address"] = data.Address.ValueString()
	}

	if !data.Uid.IsNull() && !data.Uid.IsUnknown() {
		result["uid"] = data.Uid.ValueString()
	}

	result["gateways"] = data.expandEndpointZtnaProfileConnectionRulesGatewaysList(ctx, data.Gateways, diags)

	if !data.Mask.IsNull() && !data.Mask.IsUnknown() {
		result["mask"] = data.Mask.ValueString()
	}

	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		result["port"] = data.Port.ValueString()
	}

	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		result["name"] = data.Name.ValueString()
	}

	if !data.Encryption.IsNull() && !data.Encryption.IsUnknown() {
		result["encryption"] = data.Encryption.ValueString()
	}

	return result
}

func (s *resourceEndpointZtnaProfileModel) expandEndpointZtnaProfileConnectionRulesList(ctx context.Context, l []resourceEndpointZtnaProfileConnectionRulesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointZtnaProfileConnectionRules(ctx, diags)
	}
	return result
}

func (data *resourceEndpointZtnaProfileConnectionRulesGatewaysModel) expandEndpointZtnaProfileConnectionRulesGateways(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		result["id"] = data.Id.ValueFloat64()
	}

	if !data.Alias.IsNull() && !data.Alias.IsUnknown() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.PrivateAppCount.IsNull() && !data.PrivateAppCount.IsUnknown() {
		result["private_app_count"] = data.PrivateAppCount.ValueFloat64()
	}

	if !data.Vip.IsNull() && !data.Vip.IsUnknown() {
		result["vip"] = data.Vip.ValueString()
	}

	if !data.Redirect.IsNull() && !data.Redirect.IsUnknown() {
		result["redirect"] = data.Redirect.ValueString()
	}

	return result
}

func (s *resourceEndpointZtnaProfileConnectionRulesModel) expandEndpointZtnaProfileConnectionRulesGatewaysList(ctx context.Context, l []resourceEndpointZtnaProfileConnectionRulesGatewaysModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointZtnaProfileConnectionRulesGateways(ctx, diags)
	}
	return result
}

func (data *resourceEndpointZtnaProfileEntraIdModel) expandEndpointZtnaProfileEntraId(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ApplicationId.IsNull() && !data.ApplicationId.IsUnknown() {
		result["applicationId"] = data.ApplicationId.ValueString()
	}

	if !data.DomainName.IsNull() && !data.DomainName.IsUnknown() {
		result["domainName"] = data.DomainName.ValueString()
	}

	return result
}
