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
var _ resource.Resource = &resourceNetworkImplicitDnsRule{}
var _ resource.ResourceWithMoveState = &resourceNetworkImplicitDnsRule{}

func newResourceNetworkImplicitDnsRule() resource.Resource {
	return &resourceNetworkImplicitDnsRule{}
}

type resourceNetworkImplicitDnsRule struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceNetworkImplicitDnsRuleModel describes the resource data model.
type resourceNetworkImplicitDnsRuleModel struct {
	ID         types.String `tfsdk:"id"`
	PrimaryKey types.String `tfsdk:"primary_key"`
	DnsServer  types.String `tfsdk:"dns_server"`
	DnsServer1 types.String `tfsdk:"dns_server1"`
	DnsServer2 types.String `tfsdk:"dns_server2"`
	Protocols  types.Set    `tfsdk:"protocols"`
	ForPrivate types.Bool   `tfsdk:"for_private"`
}

func (r *resourceNetworkImplicitDnsRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_implicit_dns_rule"
}

func (r *resourceNetworkImplicitDnsRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Implicit DNS Rule Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.OneOf("vpn", "other", "implicit_all"),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"dns_server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("fortiguard", "google", "quad9", "cloudflare", "endpoint", "custom"),
				},
				Computed: true,
				Optional: true,
			},
			"dns_server1": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"dns_server2": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"protocols": schema.SetAttribute{
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
		},
	}
}

func (r *resourceNetworkImplicitDnsRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_network_implicit_dns_rule"
}
func (r *resourceNetworkImplicitDnsRule) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_network_implicit_dns_rules" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceNetworkImplicitDnsRuleModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceNetworkImplicitDnsRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkImplicitDnsRules")
	lock.Lock()
	defer lock.Unlock()
	var data resourceNetworkImplicitDnsRuleModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectNetworkImplicitDnsRule(ctx, diags))
	input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRule(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateNetworkImplicitDnsRules(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRule(ctx, "read", diags))

	read_output, err := c.ReadNetworkImplicitDnsRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshNetworkImplicitDnsRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkImplicitDnsRule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkImplicitDnsRules")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceNetworkImplicitDnsRuleModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceNetworkImplicitDnsRuleModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectNetworkImplicitDnsRule(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRule(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateNetworkImplicitDnsRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRule(ctx, "read", diags))

	read_output, err := c.ReadNetworkImplicitDnsRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshNetworkImplicitDnsRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkImplicitDnsRule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceNetworkImplicitDnsRule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceNetworkImplicitDnsRuleModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkImplicitDnsRule(ctx, "read", diags))

	read_output, err := c.ReadNetworkImplicitDnsRules(&input_model)
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

	diags.Append(data.refreshNetworkImplicitDnsRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkImplicitDnsRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceNetworkImplicitDnsRuleModel) refreshNetworkImplicitDnsRule(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["dnsServer"]; ok {
		m.DnsServer = parseStringValue(v)
	}

	if v, ok := o["dnsServer1"]; ok {
		m.DnsServer1 = parseStringValue(v)
	}

	if v, ok := o["dnsServer2"]; ok {
		m.DnsServer2 = parseStringValue(v)
	}

	if v, ok := o["protocols"]; ok {
		m.Protocols = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Protocols = types.SetNull(types.StringType)
	}

	if v, ok := o["forPrivate"]; ok {
		m.ForPrivate = parseBoolValue(v)
	}

	return diags
}

func (data *resourceNetworkImplicitDnsRuleModel) getCreateObjectNetworkImplicitDnsRule(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.DnsServer.IsNull() && !data.DnsServer.IsUnknown() {
		result["dnsServer"] = data.DnsServer.ValueString()
	}

	if !data.DnsServer1.IsNull() && !data.DnsServer1.IsUnknown() {
		result["dnsServer1"] = data.DnsServer1.ValueString()
	}

	if !data.DnsServer2.IsNull() && !data.DnsServer2.IsUnknown() {
		result["dnsServer2"] = data.DnsServer2.ValueString()
	}

	if !data.Protocols.IsNull() && !data.Protocols.IsUnknown() {
		result["protocols"] = expandSetToStringList(data.Protocols)
	}

	if !data.ForPrivate.IsNull() && !data.ForPrivate.IsUnknown() {
		result["forPrivate"] = data.ForPrivate.ValueBool()
	}

	return &result
}

func (data *resourceNetworkImplicitDnsRuleModel) getUpdateObjectNetworkImplicitDnsRule(ctx context.Context, state resourceNetworkImplicitDnsRuleModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.DnsServer.IsNull() && !data.DnsServer.IsUnknown() {
		result["dnsServer"] = data.DnsServer.ValueString()
	}

	if !data.DnsServer1.IsNull() && !data.DnsServer1.IsUnknown() {
		result["dnsServer1"] = data.DnsServer1.ValueString()
	}

	if !data.DnsServer2.IsNull() && !data.DnsServer2.IsUnknown() {
		result["dnsServer2"] = data.DnsServer2.ValueString()
	}

	if !data.Protocols.IsNull() && !data.Protocols.IsUnknown() {
		result["protocols"] = expandSetToStringList(data.Protocols)
	}

	if !data.ForPrivate.IsNull() && !data.ForPrivate.IsUnknown() {
		result["forPrivate"] = data.ForPrivate.ValueBool()
	}

	return &result
}

func (data *resourceNetworkImplicitDnsRuleModel) getURLObjectNetworkImplicitDnsRule(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
