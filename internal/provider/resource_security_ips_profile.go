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
	"strings"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSecurityIpsProfile{}

func newResourceSecurityIpsProfile() resource.Resource {
	return &resourceSecurityIpsProfile{}
}

type resourceSecurityIpsProfile struct {
	fortiClient *FortiClient
}

// resourceSecurityIpsProfileModel describes the resource data model.
type resourceSecurityIpsProfileModel struct {
	ID                     types.String                                      `tfsdk:"id"`
	PrimaryKey             types.String                                      `tfsdk:"primary_key"`
	ProfileType            types.String                                      `tfsdk:"profile_type"`
	CustomRuleGroups       []resourceSecurityIpsProfileCustomRuleGroupsModel `tfsdk:"custom_rule_groups"`
	IsBlockingMaliciousUrl types.Bool                                        `tfsdk:"is_blocking_malicious_url"`
	BotnetScanning         types.String                                      `tfsdk:"botnet_scanning"`
	IsExtendedLogEnabled   types.Bool                                        `tfsdk:"is_extended_log_enabled"`
	Comment                types.String                                      `tfsdk:"comment"`
	Entries                []resourceSecurityIpsProfileEntriesModel          `tfsdk:"entries"`
	Direction              types.String                                      `tfsdk:"direction"`
}

func (r *resourceSecurityIpsProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ips_profile"
}

func (r *resourceSecurityIpsProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Required: true,
			},
			"profile_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("recommended", "critical", "monitor", "custom"),
				},
				Computed: true,
				Optional: true,
			},
			"is_blocking_malicious_url": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"botnet_scanning": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("block", "disable", "monitor"),
				},
				Computed: true,
				Optional: true,
			},
			"is_extended_log_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"comment": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"direction": schema.StringAttribute{
				Description: "The direction of the target resource.",
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				Computed: true,
				Optional: true,
			},
			"custom_rule_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "monitor", "block"),
							},
							Computed: true,
							Optional: true,
						},
						"signatures": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("security/ips-custom-signatures"),
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
			"entries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"location": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"severity": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"protocol": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"os": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"application": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"cve": schema.SetAttribute{
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable", "default"),
							},
							Computed: true,
							Optional: true,
						},
						"log": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"log_packet": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"log_attack_context": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("pass", "block", "default"),
							},
							Computed: true,
							Optional: true,
						},
						"quarantine": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"default_action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("all", "pass", "block"),
							},
							Computed: true,
							Optional: true,
						},
						"default_status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("all", "enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"rule": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("security/ips-rule", "security/ips-custom-signatures"),
										},
										Computed: true,
										Optional: true,
									},
								},
							},
							Computed: true,
							Optional: true,
						},
						"vuln_type": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.Float64Attribute{
										Computed: true,
										Optional: true,
									},
								},
							},
							Computed: true,
							Optional: true,
						},
						"exempt_ip": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.Float64Attribute{
										Computed: true,
										Optional: true,
									},
									"src_ip": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"dst_ip": schema.StringAttribute{
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

func (r *resourceSecurityIpsProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityIpsProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityIpsProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecurityIpsProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityIpsProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecurityIpsProfile(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityIpsProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityIpsProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityIpsProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityIpsProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityIpsProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityIpsProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityIpsProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityIpsProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityIpsProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityIpsProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityIpsProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityIpsProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityIpsProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityIpsProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityIpsProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityIpsProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected format: direction/primary_key, got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("direction"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), parts[1])...)
}

func (m *resourceSecurityIpsProfileModel) refreshSecurityIpsProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["profileType"]; ok {
		m.ProfileType = parseStringValue(v)
	}

	if v, ok := o["customRuleGroups"]; ok {
		m.CustomRuleGroups = m.flattenSecurityIpsProfileCustomRuleGroupsList(ctx, v, &diags)
	}

	if v, ok := o["isBlockingMaliciousUrl"]; ok {
		m.IsBlockingMaliciousUrl = parseBoolValue(v)
	}

	if v, ok := o["botnetScanning"]; ok {
		m.BotnetScanning = parseStringValue(v)
	}

	if v, ok := o["isExtendedLogEnabled"]; ok {
		m.IsExtendedLogEnabled = parseBoolValue(v)
	}

	if v, ok := o["comment"]; ok {
		m.Comment = parseStringValue(v)
	}

	if v, ok := o["entries"]; ok {
		m.Entries = m.flattenSecurityIpsProfileEntriesList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityIpsProfileModel) getCreateObjectSecurityIpsProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ProfileType.IsNull() {
		result["profileType"] = data.ProfileType.ValueString()
	}

	if len(data.CustomRuleGroups) > 0 {
		result["customRuleGroups"] = data.expandSecurityIpsProfileCustomRuleGroupsList(ctx, data.CustomRuleGroups, diags)
	}

	if !data.IsBlockingMaliciousUrl.IsNull() {
		result["isBlockingMaliciousUrl"] = data.IsBlockingMaliciousUrl.ValueBool()
	}

	if !data.BotnetScanning.IsNull() {
		result["botnetScanning"] = data.BotnetScanning.ValueString()
	}

	if !data.IsExtendedLogEnabled.IsNull() {
		result["isExtendedLogEnabled"] = data.IsExtendedLogEnabled.ValueBool()
	}

	if !data.Comment.IsNull() {
		result["comment"] = data.Comment.ValueString()
	}

	result["entries"] = data.expandSecurityIpsProfileEntriesList(ctx, data.Entries, diags)

	return &result
}

func (data *resourceSecurityIpsProfileModel) getUpdateObjectSecurityIpsProfile(ctx context.Context, state resourceSecurityIpsProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.ProfileType.IsNull() && !data.ProfileType.Equal(state.ProfileType) {
		result["profileType"] = data.ProfileType.ValueString()
	}

	if len(data.CustomRuleGroups) > 0 || !isSameStruct(data.CustomRuleGroups, state.CustomRuleGroups) {
		result["customRuleGroups"] = data.expandSecurityIpsProfileCustomRuleGroupsList(ctx, data.CustomRuleGroups, diags)
	}

	if !data.IsBlockingMaliciousUrl.IsNull() {
		result["isBlockingMaliciousUrl"] = data.IsBlockingMaliciousUrl.ValueBool()
	}

	if !data.BotnetScanning.IsNull() {
		result["botnetScanning"] = data.BotnetScanning.ValueString()
	}

	if !data.IsExtendedLogEnabled.IsNull() {
		result["isExtendedLogEnabled"] = data.IsExtendedLogEnabled.ValueBool()
	}

	if !data.Comment.IsNull() {
		result["comment"] = data.Comment.ValueString()
	}

	if len(data.Entries) > 0 || !isSameStruct(data.Entries, state.Entries) {
		result["entries"] = data.expandSecurityIpsProfileEntriesList(ctx, data.Entries, diags)
	}

	return &result
}

func (data *resourceSecurityIpsProfileModel) getURLObjectSecurityIpsProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityIpsProfileCustomRuleGroupsModel struct {
	Action     types.String                                                `tfsdk:"action"`
	Signatures []resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel `tfsdk:"signatures"`
}

type resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityIpsProfileEntriesModel struct {
	Rule             []resourceSecurityIpsProfileEntriesRuleModel     `tfsdk:"rule"`
	Location         types.String                                     `tfsdk:"location"`
	Severity         types.String                                     `tfsdk:"severity"`
	Protocol         types.String                                     `tfsdk:"protocol"`
	Os               types.String                                     `tfsdk:"os"`
	Application      types.String                                     `tfsdk:"application"`
	Cve              types.Set                                        `tfsdk:"cve"`
	Status           types.String                                     `tfsdk:"status"`
	Log              types.String                                     `tfsdk:"log"`
	LogPacket        types.String                                     `tfsdk:"log_packet"`
	LogAttackContext types.String                                     `tfsdk:"log_attack_context"`
	Action           types.String                                     `tfsdk:"action"`
	VulnType         []resourceSecurityIpsProfileEntriesVulnTypeModel `tfsdk:"vuln_type"`
	Quarantine       types.String                                     `tfsdk:"quarantine"`
	ExemptIp         []resourceSecurityIpsProfileEntriesExemptIpModel `tfsdk:"exempt_ip"`
	DefaultAction    types.String                                     `tfsdk:"default_action"`
	DefaultStatus    types.String                                     `tfsdk:"default_status"`
}

type resourceSecurityIpsProfileEntriesRuleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityIpsProfileEntriesVulnTypeModel struct {
	Id types.Float64 `tfsdk:"id"`
}

type resourceSecurityIpsProfileEntriesExemptIpModel struct {
	Id    types.Float64 `tfsdk:"id"`
	SrcIp types.String  `tfsdk:"src_ip"`
	DstIp types.String  `tfsdk:"dst_ip"`
}

func (m *resourceSecurityIpsProfileCustomRuleGroupsModel) flattenSecurityIpsProfileCustomRuleGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityIpsProfileCustomRuleGroupsModel {
	if input == nil {
		return &resourceSecurityIpsProfileCustomRuleGroupsModel{}
	}
	if m == nil {
		m = &resourceSecurityIpsProfileCustomRuleGroupsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["signatures"]; ok {
		m.Signatures = m.flattenSecurityIpsProfileCustomRuleGroupsSignaturesList(ctx, v, diags)
	}

	return m
}

func (s *resourceSecurityIpsProfileModel) flattenSecurityIpsProfileCustomRuleGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityIpsProfileCustomRuleGroupsModel {
	if o == nil {
		return []resourceSecurityIpsProfileCustomRuleGroupsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument custom_rule_groups is not type of []interface{}.", "")
		return []resourceSecurityIpsProfileCustomRuleGroupsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityIpsProfileCustomRuleGroupsModel{}
	}

	values := make([]resourceSecurityIpsProfileCustomRuleGroupsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityIpsProfileCustomRuleGroupsModel
		values[i] = *m.flattenSecurityIpsProfileCustomRuleGroups(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel) flattenSecurityIpsProfileCustomRuleGroupsSignatures(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel {
	if input == nil {
		return &resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel{}
	}
	if m == nil {
		m = &resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel{}
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

func (s *resourceSecurityIpsProfileCustomRuleGroupsModel) flattenSecurityIpsProfileCustomRuleGroupsSignaturesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel {
	if o == nil {
		return []resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument signatures is not type of []interface{}.", "")
		return []resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel{}
	}

	values := make([]resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel
		values[i] = *m.flattenSecurityIpsProfileCustomRuleGroupsSignatures(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityIpsProfileEntriesModel) flattenSecurityIpsProfileEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityIpsProfileEntriesModel {
	if input == nil {
		return &resourceSecurityIpsProfileEntriesModel{}
	}
	if m == nil {
		m = &resourceSecurityIpsProfileEntriesModel{}
	}
	o := input.(map[string]interface{})
	m.Cve = types.SetNull(types.StringType)

	if v, ok := o["rule"]; ok {
		m.Rule = m.flattenSecurityIpsProfileEntriesRuleList(ctx, v, diags)
	}

	if v, ok := o["location"]; ok {
		m.Location = parseStringValue(v)
	}

	if v, ok := o["severity"]; ok {
		m.Severity = parseStringValue(v)
	}

	if v, ok := o["protocol"]; ok {
		m.Protocol = parseStringValue(v)
	}

	if v, ok := o["os"]; ok {
		m.Os = parseStringValue(v)
	}

	if v, ok := o["application"]; ok {
		m.Application = parseStringValue(v)
	}

	if v, ok := o["cve"]; ok {
		m.Cve = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["log"]; ok {
		m.Log = parseStringValue(v)
	}

	if v, ok := o["logPacket"]; ok {
		m.LogPacket = parseStringValue(v)
	}

	if v, ok := o["logAttackContext"]; ok {
		m.LogAttackContext = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["vulnType"]; ok {
		m.VulnType = m.flattenSecurityIpsProfileEntriesVulnTypeList(ctx, v, diags)
	}

	if v, ok := o["quarantine"]; ok {
		m.Quarantine = parseStringValue(v)
	}

	if v, ok := o["exempt-ip"]; ok {
		m.ExemptIp = m.flattenSecurityIpsProfileEntriesExemptIpList(ctx, v, diags)
	}

	if v, ok := o["defaultAction"]; ok {
		m.DefaultAction = parseStringValue(v)
	}

	if v, ok := o["defaultStatus"]; ok {
		m.DefaultStatus = parseStringValue(v)
	}

	return m
}

func (s *resourceSecurityIpsProfileModel) flattenSecurityIpsProfileEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityIpsProfileEntriesModel {
	if o == nil {
		return []resourceSecurityIpsProfileEntriesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument entries is not type of []interface{}.", "")
		return []resourceSecurityIpsProfileEntriesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityIpsProfileEntriesModel{}
	}

	values := make([]resourceSecurityIpsProfileEntriesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityIpsProfileEntriesModel
		values[i] = *m.flattenSecurityIpsProfileEntries(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityIpsProfileEntriesRuleModel) flattenSecurityIpsProfileEntriesRule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityIpsProfileEntriesRuleModel {
	if input == nil {
		return &resourceSecurityIpsProfileEntriesRuleModel{}
	}
	if m == nil {
		m = &resourceSecurityIpsProfileEntriesRuleModel{}
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

func (s *resourceSecurityIpsProfileEntriesModel) flattenSecurityIpsProfileEntriesRuleList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityIpsProfileEntriesRuleModel {
	if o == nil {
		return []resourceSecurityIpsProfileEntriesRuleModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument rule is not type of []interface{}.", "")
		return []resourceSecurityIpsProfileEntriesRuleModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityIpsProfileEntriesRuleModel{}
	}

	values := make([]resourceSecurityIpsProfileEntriesRuleModel, len(l))
	for i, ele := range l {
		var m resourceSecurityIpsProfileEntriesRuleModel
		values[i] = *m.flattenSecurityIpsProfileEntriesRule(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityIpsProfileEntriesVulnTypeModel) flattenSecurityIpsProfileEntriesVulnType(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityIpsProfileEntriesVulnTypeModel {
	if input == nil {
		return &resourceSecurityIpsProfileEntriesVulnTypeModel{}
	}
	if m == nil {
		m = &resourceSecurityIpsProfileEntriesVulnTypeModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	return m
}

func (s *resourceSecurityIpsProfileEntriesModel) flattenSecurityIpsProfileEntriesVulnTypeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityIpsProfileEntriesVulnTypeModel {
	if o == nil {
		return []resourceSecurityIpsProfileEntriesVulnTypeModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument vuln_type is not type of []interface{}.", "")
		return []resourceSecurityIpsProfileEntriesVulnTypeModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityIpsProfileEntriesVulnTypeModel{}
	}

	values := make([]resourceSecurityIpsProfileEntriesVulnTypeModel, len(l))
	for i, ele := range l {
		var m resourceSecurityIpsProfileEntriesVulnTypeModel
		values[i] = *m.flattenSecurityIpsProfileEntriesVulnType(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityIpsProfileEntriesExemptIpModel) flattenSecurityIpsProfileEntriesExemptIp(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityIpsProfileEntriesExemptIpModel {
	if input == nil {
		return &resourceSecurityIpsProfileEntriesExemptIpModel{}
	}
	if m == nil {
		m = &resourceSecurityIpsProfileEntriesExemptIpModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	if v, ok := o["src-ip"]; ok {
		m.SrcIp = parseStringValue(v)
	}

	if v, ok := o["dst-ip"]; ok {
		m.DstIp = parseStringValue(v)
	}

	return m
}

func (s *resourceSecurityIpsProfileEntriesModel) flattenSecurityIpsProfileEntriesExemptIpList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityIpsProfileEntriesExemptIpModel {
	if o == nil {
		return []resourceSecurityIpsProfileEntriesExemptIpModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument exempt_ip is not type of []interface{}.", "")
		return []resourceSecurityIpsProfileEntriesExemptIpModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityIpsProfileEntriesExemptIpModel{}
	}

	values := make([]resourceSecurityIpsProfileEntriesExemptIpModel, len(l))
	for i, ele := range l {
		var m resourceSecurityIpsProfileEntriesExemptIpModel
		values[i] = *m.flattenSecurityIpsProfileEntriesExemptIp(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecurityIpsProfileCustomRuleGroupsModel) expandSecurityIpsProfileCustomRuleGroups(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	result["signatures"] = data.expandSecurityIpsProfileCustomRuleGroupsSignaturesList(ctx, data.Signatures, diags)

	return result
}

func (s *resourceSecurityIpsProfileModel) expandSecurityIpsProfileCustomRuleGroupsList(ctx context.Context, l []resourceSecurityIpsProfileCustomRuleGroupsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityIpsProfileCustomRuleGroups(ctx, diags)
	}
	return result
}

func (data *resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel) expandSecurityIpsProfileCustomRuleGroupsSignatures(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityIpsProfileCustomRuleGroupsModel) expandSecurityIpsProfileCustomRuleGroupsSignaturesList(ctx context.Context, l []resourceSecurityIpsProfileCustomRuleGroupsSignaturesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityIpsProfileCustomRuleGroupsSignatures(ctx, diags)
	}
	return result
}

func (data *resourceSecurityIpsProfileEntriesModel) expandSecurityIpsProfileEntries(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if len(data.Rule) > 0 {
		result["rule"] = data.expandSecurityIpsProfileEntriesRuleList(ctx, data.Rule, diags)
	}

	if !data.Location.IsNull() {
		result["location"] = data.Location.ValueString()
	}

	if !data.Severity.IsNull() {
		result["severity"] = data.Severity.ValueString()
	}

	if !data.Protocol.IsNull() {
		result["protocol"] = data.Protocol.ValueString()
	}

	if !data.Os.IsNull() {
		result["os"] = data.Os.ValueString()
	}

	if !data.Application.IsNull() {
		result["application"] = data.Application.ValueString()
	}

	if !data.Cve.IsNull() {
		result["cve"] = expandSetToStringList(data.Cve)
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.Log.IsNull() {
		result["log"] = data.Log.ValueString()
	}

	if !data.LogPacket.IsNull() {
		result["logPacket"] = data.LogPacket.ValueString()
	}

	if !data.LogAttackContext.IsNull() {
		result["logAttackContext"] = data.LogAttackContext.ValueString()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if len(data.VulnType) > 0 {
		result["vulnType"] = data.expandSecurityIpsProfileEntriesVulnTypeList(ctx, data.VulnType, diags)
	}

	if !data.Quarantine.IsNull() {
		result["quarantine"] = data.Quarantine.ValueString()
	}

	if len(data.ExemptIp) > 0 {
		result["exempt-ip"] = data.expandSecurityIpsProfileEntriesExemptIpList(ctx, data.ExemptIp, diags)
	}

	if !data.DefaultAction.IsNull() {
		result["defaultAction"] = data.DefaultAction.ValueString()
	}

	if !data.DefaultStatus.IsNull() {
		result["defaultStatus"] = data.DefaultStatus.ValueString()
	}

	return result
}

func (s *resourceSecurityIpsProfileModel) expandSecurityIpsProfileEntriesList(ctx context.Context, l []resourceSecurityIpsProfileEntriesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityIpsProfileEntries(ctx, diags)
	}
	return result
}

func (data *resourceSecurityIpsProfileEntriesRuleModel) expandSecurityIpsProfileEntriesRule(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityIpsProfileEntriesModel) expandSecurityIpsProfileEntriesRuleList(ctx context.Context, l []resourceSecurityIpsProfileEntriesRuleModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityIpsProfileEntriesRule(ctx, diags)
	}
	return result
}

func (data *resourceSecurityIpsProfileEntriesVulnTypeModel) expandSecurityIpsProfileEntriesVulnType(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() {
		result["id"] = data.Id.ValueFloat64()
	}

	return result
}

func (s *resourceSecurityIpsProfileEntriesModel) expandSecurityIpsProfileEntriesVulnTypeList(ctx context.Context, l []resourceSecurityIpsProfileEntriesVulnTypeModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityIpsProfileEntriesVulnType(ctx, diags)
	}
	return result
}

func (data *resourceSecurityIpsProfileEntriesExemptIpModel) expandSecurityIpsProfileEntriesExemptIp(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() {
		result["id"] = data.Id.ValueFloat64()
	}

	if !data.SrcIp.IsNull() {
		result["src-ip"] = data.SrcIp.ValueString()
	}

	if !data.DstIp.IsNull() {
		result["dst-ip"] = data.DstIp.ValueString()
	}

	return result
}

func (s *resourceSecurityIpsProfileEntriesModel) expandSecurityIpsProfileEntriesExemptIpList(ctx context.Context, l []resourceSecurityIpsProfileEntriesExemptIpModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityIpsProfileEntriesExemptIp(ctx, diags)
	}
	return result
}
