// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/setvalidatorwarning"
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
var _ resource.Resource = &resourceNetworkDnsRule{}
var _ resource.ResourceWithMoveState = &resourceNetworkDnsRule{}

func newResourceNetworkDnsRule() resource.Resource {
	return &resourceNetworkDnsRule{}
}

type resourceNetworkDnsRule struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceNetworkDnsRuleModel describes the resource data model.
type resourceNetworkDnsRuleModel struct {
	ID             types.String                                         `tfsdk:"id"`
	PrimaryKey     types.String                                         `tfsdk:"primary_key"`
	PrimaryDns     types.String                                         `tfsdk:"primary_dns"`
	SecondaryDns   types.String                                         `tfsdk:"secondary_dns"`
	Domains        types.Set                                            `tfsdk:"domains"`
	PopDnsOverride map[string]resourceNetworkDnsRulePopDnsOverrideModel `tfsdk:"pop_dns_override"`
	ForPrivate     types.Bool                                           `tfsdk:"for_private"`
}

func (r *resourceNetworkDnsRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_dns_rule"
}

func (r *resourceNetworkDnsRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DNS Rule Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.LengthAtMost(30),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"primary_dns": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"secondary_dns": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"domains": schema.SetAttribute{
				Validators: []validator.Set{
					setvalidatorwarning.SizeAtLeast(1),
				},
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"for_private": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"pop_dns_override": schema.MapNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pop": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"primary_dns": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"secondary_dns": schema.StringAttribute{
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

func (r *resourceNetworkDnsRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_network_dns_rule"
}
func (r *resourceNetworkDnsRule) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_network_dns_rules" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceNetworkDnsRuleModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceNetworkDnsRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkDnsRules")
	lock.Lock()
	defer lock.Unlock()
	var data resourceNetworkDnsRuleModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectNetworkDnsRule(ctx, diags))
	input_model.URLParams = *(data.getURLObjectNetworkDnsRule(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateNetworkDnsRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	if responseMkey, ok := getCreateResponseMkey(output, "primaryKey"); ok {
		mkey = responseMkey
	}
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectNetworkDnsRule(ctx, "read", diags))

	read_output, err := c.ReadNetworkDnsRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshNetworkDnsRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkDnsRule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkDnsRules")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceNetworkDnsRuleModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceNetworkDnsRuleModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectNetworkDnsRule(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectNetworkDnsRule(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateNetworkDnsRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectNetworkDnsRule(ctx, "read", diags))

	read_output, err := c.ReadNetworkDnsRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshNetworkDnsRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkDnsRule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkDnsRules")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceNetworkDnsRuleModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkDnsRule(ctx, "delete", diags))

	output, err := c.DeleteNetworkDnsRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceNetworkDnsRule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceNetworkDnsRuleModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkDnsRule(ctx, "read", diags))

	read_output, err := c.ReadNetworkDnsRules(&input_model)
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

	diags.Append(data.refreshNetworkDnsRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkDnsRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceNetworkDnsRuleModel) refreshNetworkDnsRule(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryDns"]; ok {
		m.PrimaryDns = parseStringValue(v)
	}

	if v, ok := o["secondaryDns"]; ok {
		m.SecondaryDns = parseStringValue(v)
	}

	if v, ok := o["domains"]; ok {
		m.Domains = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Domains = types.SetNull(types.StringType)
	}

	if v, ok := o["popDnsOverride"]; ok {
		m.PopDnsOverride = m.flattenNetworkDnsRulePopDnsOverrideMap(ctx, v.(map[string]interface{}), &diags)
	}

	if v, ok := o["forPrivate"]; ok {
		m.ForPrivate = parseBoolValue(v)
	}

	return diags
}

func (data *resourceNetworkDnsRuleModel) getCreateObjectNetworkDnsRule(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.PrimaryDns.IsNull() && !data.PrimaryDns.IsUnknown() {
		result["primaryDns"] = data.PrimaryDns.ValueString()
	}

	if !data.SecondaryDns.IsNull() && !data.SecondaryDns.IsUnknown() {
		result["secondaryDns"] = data.SecondaryDns.ValueString()
	}

	if !data.Domains.IsNull() && !data.Domains.IsUnknown() {
		result["domains"] = expandSetToStringList(data.Domains)
	}

	if data.PopDnsOverride != nil {
		result["popDnsOverride"] = data.expandNetworkDnsRulePopDnsOverrideMap(ctx, diags)
	}

	if !data.ForPrivate.IsNull() && !data.ForPrivate.IsUnknown() {
		result["forPrivate"] = data.ForPrivate.ValueBool()
	}

	return &result
}

func (data *resourceNetworkDnsRuleModel) getUpdateObjectNetworkDnsRule(ctx context.Context, state resourceNetworkDnsRuleModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.PrimaryDns.IsNull() && !data.PrimaryDns.IsUnknown() {
		result["primaryDns"] = data.PrimaryDns.ValueString()
	}

	if !data.SecondaryDns.IsNull() && !data.SecondaryDns.IsUnknown() {
		result["secondaryDns"] = data.SecondaryDns.ValueString()
	}

	if !data.Domains.IsNull() && !data.Domains.IsUnknown() {
		result["domains"] = expandSetToStringList(data.Domains)
	}

	if data.PopDnsOverride != nil {
		result["popDnsOverride"] = data.expandNetworkDnsRulePopDnsOverrideMap(ctx, diags)
	}

	if !data.ForPrivate.IsNull() && !data.ForPrivate.IsUnknown() {
		result["forPrivate"] = data.ForPrivate.ValueBool()
	}

	return &result
}

func (data *resourceNetworkDnsRuleModel) getURLObjectNetworkDnsRule(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceNetworkDnsRulePopDnsOverrideModel struct {
	Pop          types.String `tfsdk:"pop"`
	PrimaryDns   types.String `tfsdk:"primary_dns"`
	SecondaryDns types.String `tfsdk:"secondary_dns"`
}

func (m *resourceNetworkDnsRulePopDnsOverrideModel) flattenNetworkDnsRulePopDnsOverride(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceNetworkDnsRulePopDnsOverrideModel {
	if input == nil {
		return &resourceNetworkDnsRulePopDnsOverrideModel{}
	}
	if m == nil {
		m = &resourceNetworkDnsRulePopDnsOverrideModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["pop"]; ok {
		m.Pop = parseStringValue(v)
	}

	if v, ok := o["primaryDns"]; ok {
		m.PrimaryDns = parseStringValue(v)
	}

	if v, ok := o["secondaryDns"]; ok {
		m.SecondaryDns = parseStringValue(v)
	}

	return m
}

func (s *resourceNetworkDnsRuleModel) flattenNetworkDnsRulePopDnsOverrideMap(ctx context.Context, o map[string]interface{}, diags *diag.Diagnostics) map[string]resourceNetworkDnsRulePopDnsOverrideModel {
	result := make(map[string]resourceNetworkDnsRulePopDnsOverrideModel)
	for k, v := range o {
		var m resourceNetworkDnsRulePopDnsOverrideModel
		if existing, ok := s.PopDnsOverride[k]; ok {
			m = existing
		}
		m = *m.flattenNetworkDnsRulePopDnsOverride(ctx, v, diags)
		result[k] = m
	}
	return result
}

func (data *resourceNetworkDnsRulePopDnsOverrideModel) expandNetworkDnsRulePopDnsOverride(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Pop.IsNull() && !data.Pop.IsUnknown() {
		result["pop"] = data.Pop.ValueString()
	}

	if !data.PrimaryDns.IsNull() && !data.PrimaryDns.IsUnknown() {
		result["primaryDns"] = data.PrimaryDns.ValueString()
	}

	if !data.SecondaryDns.IsNull() && !data.SecondaryDns.IsUnknown() {
		result["secondaryDns"] = data.SecondaryDns.ValueString()
	}

	return result
}

func (s *resourceNetworkDnsRuleModel) expandNetworkDnsRulePopDnsOverrideMap(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	m := s.PopDnsOverride
	for k, v := range m {
		result[k] = v.expandNetworkDnsRulePopDnsOverride(ctx, diags)
	}
	return result
}
