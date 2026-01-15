// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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
var _ resource.Resource = &resourceNetworkDnsRules{}

func newResourceNetworkDnsRules() resource.Resource {
	return &resourceNetworkDnsRules{}
}

type resourceNetworkDnsRules struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceNetworkDnsRulesModel describes the resource data model.
type resourceNetworkDnsRulesModel struct {
	ID             types.String                                          `tfsdk:"id"`
	PrimaryKey     types.String                                          `tfsdk:"primary_key"`
	PrimaryDns     types.String                                          `tfsdk:"primary_dns"`
	SecondaryDns   types.String                                          `tfsdk:"secondary_dns"`
	Domains        types.Set                                             `tfsdk:"domains"`
	PopDnsOverride map[string]resourceNetworkDnsRulesPopDnsOverrideModel `tfsdk:"pop_dns_override"`
	ForPrivate     types.Bool                                            `tfsdk:"for_private"`
}

func (r *resourceNetworkDnsRules) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_dns_rules"
}

func (r *resourceNetworkDnsRules) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					stringvalidator.LengthAtMost(30),
				},
				Required: true,
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
					setvalidator.SizeAtLeast(1),
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

func (r *resourceNetworkDnsRules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_network_dns_rules"
}

func (r *resourceNetworkDnsRules) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkDnsRules")
	lock.Lock()
	defer lock.Unlock()
	var data resourceNetworkDnsRulesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectNetworkDnsRules(ctx, diags))
	input_model.URLParams = *(data.getURLObjectNetworkDnsRules(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateNetworkDnsRules(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectNetworkDnsRules(ctx, "read", diags))

	read_output, err := c.ReadNetworkDnsRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshNetworkDnsRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkDnsRules) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkDnsRules")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceNetworkDnsRulesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceNetworkDnsRulesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectNetworkDnsRules(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectNetworkDnsRules(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectNetworkDnsRules(ctx, "read", diags))

	read_output, err := c.ReadNetworkDnsRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshNetworkDnsRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkDnsRules) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkDnsRules")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceNetworkDnsRulesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkDnsRules(ctx, "delete", diags))

	output, err := c.DeleteNetworkDnsRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceNetworkDnsRules) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceNetworkDnsRulesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkDnsRules(ctx, "read", diags))

	read_output, err := c.ReadNetworkDnsRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshNetworkDnsRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkDnsRules) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceNetworkDnsRulesModel) refreshNetworkDnsRules(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
	}

	if v, ok := o["popDnsOverride"]; ok {
		m.PopDnsOverride = m.flattenNetworkDnsRulesPopDnsOverrideMap(ctx, v.(map[string]interface{}), &diags)
	}

	if v, ok := o["forPrivate"]; ok {
		m.ForPrivate = parseBoolValue(v)
	}

	return diags
}

func (data *resourceNetworkDnsRulesModel) getCreateObjectNetworkDnsRules(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.PrimaryDns.IsNull() {
		result["primaryDns"] = data.PrimaryDns.ValueString()
	}

	if !data.SecondaryDns.IsNull() {
		result["secondaryDns"] = data.SecondaryDns.ValueString()
	}

	if !data.Domains.IsNull() {
		result["domains"] = expandSetToStringList(data.Domains)
	}

	if data.PopDnsOverride != nil {
		result["popDnsOverride"] = data.expandNetworkDnsRulesPopDnsOverrideMap(ctx, diags)
	}

	if !data.ForPrivate.IsNull() {
		result["forPrivate"] = data.ForPrivate.ValueBool()
	}

	return &result
}

func (data *resourceNetworkDnsRulesModel) getUpdateObjectNetworkDnsRules(ctx context.Context, state resourceNetworkDnsRulesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.PrimaryDns.IsNull() {
		result["primaryDns"] = data.PrimaryDns.ValueString()
	}

	if !data.SecondaryDns.IsNull() {
		result["secondaryDns"] = data.SecondaryDns.ValueString()
	}

	if !data.Domains.IsNull() {
		result["domains"] = expandSetToStringList(data.Domains)
	}

	if data.PopDnsOverride != nil {
		result["popDnsOverride"] = data.expandNetworkDnsRulesPopDnsOverrideMap(ctx, diags)
	}

	if !data.ForPrivate.IsNull() {
		result["forPrivate"] = data.ForPrivate.ValueBool()
	}

	return &result
}

func (data *resourceNetworkDnsRulesModel) getURLObjectNetworkDnsRules(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceNetworkDnsRulesPopDnsOverrideModel struct {
	Pop          types.String `tfsdk:"pop"`
	PrimaryDns   types.String `tfsdk:"primary_dns"`
	SecondaryDns types.String `tfsdk:"secondary_dns"`
}

func (m *resourceNetworkDnsRulesPopDnsOverrideModel) flattenNetworkDnsRulesPopDnsOverride(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceNetworkDnsRulesPopDnsOverrideModel {
	if input == nil {
		return &resourceNetworkDnsRulesPopDnsOverrideModel{}
	}
	if m == nil {
		m = &resourceNetworkDnsRulesPopDnsOverrideModel{}
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

func (s *resourceNetworkDnsRulesModel) flattenNetworkDnsRulesPopDnsOverrideMap(ctx context.Context, o map[string]interface{}, diags *diag.Diagnostics) map[string]resourceNetworkDnsRulesPopDnsOverrideModel {
	result := make(map[string]resourceNetworkDnsRulesPopDnsOverrideModel)
	for k, v := range o {
		var m resourceNetworkDnsRulesPopDnsOverrideModel
		m = *m.flattenNetworkDnsRulesPopDnsOverride(ctx, v, diags)
		result[k] = m
	}
	return result
}

func (data *resourceNetworkDnsRulesPopDnsOverrideModel) expandNetworkDnsRulesPopDnsOverride(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Pop.IsNull() {
		result["pop"] = data.Pop.ValueString()
	}

	if !data.PrimaryDns.IsNull() {
		result["primaryDns"] = data.PrimaryDns.ValueString()
	}

	if !data.SecondaryDns.IsNull() {
		result["secondaryDns"] = data.SecondaryDns.ValueString()
	}

	return result
}

func (s *resourceNetworkDnsRulesModel) expandNetworkDnsRulesPopDnsOverrideMap(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	m := s.PopDnsOverride
	for k, v := range m {
		result[k] = v.expandNetworkDnsRulesPopDnsOverride(ctx, diags)
	}
	return result
}
