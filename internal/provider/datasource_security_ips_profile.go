// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityIpsProfile{}

func newDatasourceSecurityIpsProfile() datasource.DataSource {
	return &datasourceSecurityIpsProfile{}
}

type datasourceSecurityIpsProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityIpsProfileModel describes the datasource data model.
type datasourceSecurityIpsProfileModel struct {
	PrimaryKey             types.String                                        `tfsdk:"primary_key"`
	ProfileType            types.String                                        `tfsdk:"profile_type"`
	CustomRuleGroups       []datasourceSecurityIpsProfileCustomRuleGroupsModel `tfsdk:"custom_rule_groups"`
	IsBlockingMaliciousUrl types.Bool                                          `tfsdk:"is_blocking_malicious_url"`
	BotnetScanning         types.String                                        `tfsdk:"botnet_scanning"`
	IsExtendedLogEnabled   types.Bool                                          `tfsdk:"is_extended_log_enabled"`
	Comment                types.String                                        `tfsdk:"comment"`
	Entries                []datasourceSecurityIpsProfileEntriesModel          `tfsdk:"entries"`
	Direction              types.String                                        `tfsdk:"direction"`
}

func (r *datasourceSecurityIpsProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ips_profile"
}

func (r *datasourceSecurityIpsProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "IPS Profile Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"profile_type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("recommended", "critical", "monitor", "custom"),
				},
				Computed: true,
			},
			"is_blocking_malicious_url": schema.BoolAttribute{
				Computed: true,
			},
			"botnet_scanning": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("block", "disable", "monitor"),
				},
				Computed: true,
			},
			"is_extended_log_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Computed: true,
			},
			"direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("internal-profiles", "outbound-profiles"),
				},
				MarkdownDescription: "The direction of the target resource.\nSupported values: internal-profiles, outbound-profiles.",
				Required:            true,
			},
			"custom_rule_groups": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("allow", "monitor", "block"),
							},
							Computed: true,
						},
						"signatures": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Computed: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidatorwarning.OneOf("security/ips-custom-signatures"),
										},
										Computed: true,
									},
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"entries": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"location": schema.StringAttribute{
							Computed: true,
						},
						"severity": schema.StringAttribute{
							Computed: true,
						},
						"protocol": schema.StringAttribute{
							Computed: true,
						},
						"os": schema.StringAttribute{
							Computed: true,
						},
						"application": schema.StringAttribute{
							Computed: true,
						},
						"cve": schema.SetAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable", "default"),
							},
							Computed: true,
						},
						"log": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"log_packet": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("enable", "disable"),
							},
							Computed: true,
						},
						"log_attack_context": schema.StringAttribute{
							Computed: true,
						},
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("pass", "block", "default"),
							},
							Computed: true,
						},
						"quarantine": schema.StringAttribute{
							Computed: true,
						},
						"default_action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("all", "pass", "block"),
							},
							Computed: true,
						},
						"default_status": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("all", "enable", "disable"),
							},
							Computed: true,
						},
						"rule": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Computed: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidatorwarning.OneOf("security/ips-rule", "security/ips-custom-signatures"),
										},
										Computed: true,
									},
								},
							},
							Computed: true,
						},
						"vuln_type": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.Float64Attribute{
										Computed: true,
									},
								},
							},
							Computed: true,
						},
						"exempt_ip": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.Float64Attribute{
										Computed: true,
									},
									"src_ip": schema.StringAttribute{
										Computed: true,
									},
									"dst_ip": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceSecurityIpsProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_ips_profile"
}

func (r *datasourceSecurityIpsProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityIpsProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityIpsProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityIpsProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityIpsProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityIpsProfileModel) refreshSecurityIpsProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
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

func (data *datasourceSecurityIpsProfileModel) getURLObjectSecurityIpsProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() && !data.Direction.IsUnknown() {
		diags.AddWarning("\"direction\" is deprecated and may be removed in future.",
			"It is recommended to recreate the resource without \"direction\" to avoid unexpected behavior in future.",
		)
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityIpsProfileCustomRuleGroupsModel struct {
	Action     types.String                                                  `tfsdk:"action"`
	Signatures []datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel `tfsdk:"signatures"`
}

type datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityIpsProfileEntriesModel struct {
	Rule             []datasourceSecurityIpsProfileEntriesRuleModel     `tfsdk:"rule"`
	Location         types.String                                       `tfsdk:"location"`
	Severity         types.String                                       `tfsdk:"severity"`
	Protocol         types.String                                       `tfsdk:"protocol"`
	Os               types.String                                       `tfsdk:"os"`
	Application      types.String                                       `tfsdk:"application"`
	Cve              types.Set                                          `tfsdk:"cve"`
	Status           types.String                                       `tfsdk:"status"`
	Log              types.String                                       `tfsdk:"log"`
	LogPacket        types.String                                       `tfsdk:"log_packet"`
	LogAttackContext types.String                                       `tfsdk:"log_attack_context"`
	Action           types.String                                       `tfsdk:"action"`
	VulnType         []datasourceSecurityIpsProfileEntriesVulnTypeModel `tfsdk:"vuln_type"`
	Quarantine       types.String                                       `tfsdk:"quarantine"`
	ExemptIp         []datasourceSecurityIpsProfileEntriesExemptIpModel `tfsdk:"exempt_ip"`
	DefaultAction    types.String                                       `tfsdk:"default_action"`
	DefaultStatus    types.String                                       `tfsdk:"default_status"`
}

type datasourceSecurityIpsProfileEntriesRuleModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityIpsProfileEntriesVulnTypeModel struct {
	Id types.Float64 `tfsdk:"id"`
}

type datasourceSecurityIpsProfileEntriesExemptIpModel struct {
	Id    types.Float64 `tfsdk:"id"`
	SrcIp types.String  `tfsdk:"src_ip"`
	DstIp types.String  `tfsdk:"dst_ip"`
}

func (m *datasourceSecurityIpsProfileCustomRuleGroupsModel) flattenSecurityIpsProfileCustomRuleGroups(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityIpsProfileCustomRuleGroupsModel {
	if input == nil {
		return &datasourceSecurityIpsProfileCustomRuleGroupsModel{}
	}
	if m == nil {
		m = &datasourceSecurityIpsProfileCustomRuleGroupsModel{}
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

func (s *datasourceSecurityIpsProfileModel) flattenSecurityIpsProfileCustomRuleGroupsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityIpsProfileCustomRuleGroupsModel {
	if o == nil {
		return []datasourceSecurityIpsProfileCustomRuleGroupsModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument custom_rule_groups is not type of []interface{}.", "")
		return []datasourceSecurityIpsProfileCustomRuleGroupsModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityIpsProfileCustomRuleGroupsModel{}
	}

	values := make([]datasourceSecurityIpsProfileCustomRuleGroupsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityIpsProfileCustomRuleGroupsModel
		if i < len(s.CustomRuleGroups) {
			m = s.CustomRuleGroups[i]
		}
		values[i] = *m.flattenSecurityIpsProfileCustomRuleGroups(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel) flattenSecurityIpsProfileCustomRuleGroupsSignatures(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel {
	if input == nil {
		return &datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel{}
	}
	if m == nil {
		m = &datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel{}
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

func (s *datasourceSecurityIpsProfileCustomRuleGroupsModel) flattenSecurityIpsProfileCustomRuleGroupsSignaturesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel {
	if o == nil {
		return []datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument signatures is not type of []interface{}.", "")
		return []datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel{}
	}

	values := make([]datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityIpsProfileCustomRuleGroupsSignaturesModel
		if i < len(s.Signatures) {
			m = s.Signatures[i]
		}
		values[i] = *m.flattenSecurityIpsProfileCustomRuleGroupsSignatures(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityIpsProfileEntriesModel) flattenSecurityIpsProfileEntries(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityIpsProfileEntriesModel {
	if input == nil {
		return &datasourceSecurityIpsProfileEntriesModel{}
	}
	if m == nil {
		m = &datasourceSecurityIpsProfileEntriesModel{}
	}
	o := input.(map[string]interface{})
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
	} else {
		m.Cve = types.SetNull(types.StringType)
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

func (s *datasourceSecurityIpsProfileModel) flattenSecurityIpsProfileEntriesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityIpsProfileEntriesModel {
	if o == nil {
		return []datasourceSecurityIpsProfileEntriesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument entries is not type of []interface{}.", "")
		return []datasourceSecurityIpsProfileEntriesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityIpsProfileEntriesModel{}
	}

	values := make([]datasourceSecurityIpsProfileEntriesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityIpsProfileEntriesModel
		if i < len(s.Entries) {
			m = s.Entries[i]
		}
		values[i] = *m.flattenSecurityIpsProfileEntries(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityIpsProfileEntriesRuleModel) flattenSecurityIpsProfileEntriesRule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityIpsProfileEntriesRuleModel {
	if input == nil {
		return &datasourceSecurityIpsProfileEntriesRuleModel{}
	}
	if m == nil {
		m = &datasourceSecurityIpsProfileEntriesRuleModel{}
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

func (s *datasourceSecurityIpsProfileEntriesModel) flattenSecurityIpsProfileEntriesRuleList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityIpsProfileEntriesRuleModel {
	if o == nil {
		return []datasourceSecurityIpsProfileEntriesRuleModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument rule is not type of []interface{}.", "")
		return []datasourceSecurityIpsProfileEntriesRuleModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityIpsProfileEntriesRuleModel{}
	}

	values := make([]datasourceSecurityIpsProfileEntriesRuleModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityIpsProfileEntriesRuleModel
		if i < len(s.Rule) {
			m = s.Rule[i]
		}
		values[i] = *m.flattenSecurityIpsProfileEntriesRule(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityIpsProfileEntriesVulnTypeModel) flattenSecurityIpsProfileEntriesVulnType(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityIpsProfileEntriesVulnTypeModel {
	if input == nil {
		return &datasourceSecurityIpsProfileEntriesVulnTypeModel{}
	}
	if m == nil {
		m = &datasourceSecurityIpsProfileEntriesVulnTypeModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	return m
}

func (s *datasourceSecurityIpsProfileEntriesModel) flattenSecurityIpsProfileEntriesVulnTypeList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityIpsProfileEntriesVulnTypeModel {
	if o == nil {
		return []datasourceSecurityIpsProfileEntriesVulnTypeModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument vuln_type is not type of []interface{}.", "")
		return []datasourceSecurityIpsProfileEntriesVulnTypeModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityIpsProfileEntriesVulnTypeModel{}
	}

	values := make([]datasourceSecurityIpsProfileEntriesVulnTypeModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityIpsProfileEntriesVulnTypeModel
		if i < len(s.VulnType) {
			m = s.VulnType[i]
		}
		values[i] = *m.flattenSecurityIpsProfileEntriesVulnType(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityIpsProfileEntriesExemptIpModel) flattenSecurityIpsProfileEntriesExemptIp(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityIpsProfileEntriesExemptIpModel {
	if input == nil {
		return &datasourceSecurityIpsProfileEntriesExemptIpModel{}
	}
	if m == nil {
		m = &datasourceSecurityIpsProfileEntriesExemptIpModel{}
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

func (s *datasourceSecurityIpsProfileEntriesModel) flattenSecurityIpsProfileEntriesExemptIpList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityIpsProfileEntriesExemptIpModel {
	if o == nil {
		return []datasourceSecurityIpsProfileEntriesExemptIpModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument exempt_ip is not type of []interface{}.", "")
		return []datasourceSecurityIpsProfileEntriesExemptIpModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityIpsProfileEntriesExemptIpModel{}
	}

	values := make([]datasourceSecurityIpsProfileEntriesExemptIpModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityIpsProfileEntriesExemptIpModel
		if i < len(s.ExemptIp) {
			m = s.ExemptIp[i]
		}
		values[i] = *m.flattenSecurityIpsProfileEntriesExemptIp(ctx, ele, diags)
	}

	return values
}
