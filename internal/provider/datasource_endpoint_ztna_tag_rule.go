// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointZtnaTagRule{}

func newDatasourceEndpointZtnaTagRule() datasource.DataSource {
	return &datasourceEndpointZtnaTagRule{}
}

type datasourceEndpointZtnaTagRule struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointZtnaTagRuleModel describes the datasource data model.
type datasourceEndpointZtnaTagRuleModel struct {
	PrimaryKey  types.String                              `tfsdk:"primary_key"`
	Status      types.String                              `tfsdk:"status"`
	Description types.String                              `tfsdk:"description"`
	Comments    types.String                              `tfsdk:"comments"`
	Rules       []datasourceEndpointZtnaTagRuleRulesModel `tfsdk:"rules"`
	Logic       *datasourceEndpointZtnaTagRuleLogicModel  `tfsdk:"logic"`
}

func (r *datasourceEndpointZtnaTagRule) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_tag_rule"
}

func (r *datasourceEndpointZtnaTagRule) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ZTNA Tag Rule Resource API V2 for FortiSASE. This resource is restricted to EMS version: 7.4.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 128),
				},
				Required: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"description": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 512),
				},
				Computed: true,
			},
			"comments": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(1000),
				},
				Computed: true,
			},
			"rules": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtLeast(1),
							},
							Computed: true,
						},
						"os": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("windows", "macos", "linux", "ios", "android"),
							},
							Computed: true,
						},
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("ad-groups", "anti-virus", "certificate", "file", "logged-in-domain", "running-process", "registry-key", "os-version", "sandbox-detection", "vulnerable-devices", "windows-security", "user-identity", "ems-management", "security", "ip-range", "on-fabric-status", "fct-version", "security-status", "cve", "crowdstrike-zta-score"),
							},
							Computed: true,
						},
						"service": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("Google", "LinkedIn", "Salesforce", "Custom"),
							},
							Computed: true,
						},
						"account": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthBetween(1, 256),
							},
							Computed: true,
						},
						"match_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("simple", "regex", "wildcard"),
							},
							Computed: true,
						},
						"subject": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthBetween(1, 256),
							},
							Computed: true,
						},
						"issuer": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthBetween(1, 256),
							},
							Computed: true,
						},
						"content": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthBetween(1, 256),
							},
							Computed: true,
						},
						"path": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
							},
							Computed: true,
						},
						"negated": schema.BoolAttribute{
							Computed: true,
						},
						"enable_latest_update_check": schema.BoolAttribute{
							Computed: true,
						},
						"check_updates_within_days": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.Between(1, 3653),
							},
							Computed: true,
						},
						"comparator": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("=", ">", "<", ">=", "<="),
							},
							Computed: true,
						},
						"condition": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthBetween(1, 256),
									},
									Computed: true,
								},
								"is_dword": schema.BoolAttribute{
									Computed: true,
								},
								"comparator": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("=", "!=", ">", ">=", "<", "<="),
									},
									Computed: true,
								},
								"value": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthBetween(1, 256),
									},
									Computed: true,
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"logic": schema.SingleNestedAttribute{

				Computed: true,
			},
		},
	}
}

func (r *datasourceEndpointZtnaTagRule) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	support_versions := map[string][]string{
		"EMS": {"7.4"},
	}
	ok, err := checkVersionMatch(client.Client, support_versions)
	if !ok {
		resp.Diagnostics.AddError(
			"FortiSASE EMS version do not support this resource.",
			fmt.Sprintf("%v", err),
		)

		return
	}

	r.fortiClient = client
	r.resourceName = "fortisase_endpoint_ztna_tag_rule"
}

func (r *datasourceEndpointZtnaTagRule) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointZtnaTagRuleModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointZtnaTagRule(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaTagRule(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaTagRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointZtnaTagRuleModel) refreshEndpointZtnaTagRule(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["description"]; ok {
		m.Description = parseStringValue(v)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["rules"]; ok {
		m.Rules = m.flattenEndpointZtnaTagRuleRulesList(ctx, v, &diags)
	}

	if v, ok := o["logic"]; ok {
		m.Logic = m.Logic.flattenEndpointZtnaTagRuleLogic(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointZtnaTagRuleModel) getURLObjectEndpointZtnaTagRule(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointZtnaTagRuleRulesModel struct {
	Id                      types.Float64                                     `tfsdk:"id"`
	Os                      types.String                                      `tfsdk:"os"`
	Type                    types.String                                      `tfsdk:"type"`
	Service                 types.String                                      `tfsdk:"service"`
	Account                 types.String                                      `tfsdk:"account"`
	MatchType               types.String                                      `tfsdk:"match_type"`
	Subject                 types.String                                      `tfsdk:"subject"`
	Issuer                  types.String                                      `tfsdk:"issuer"`
	Content                 types.String                                      `tfsdk:"content"`
	Path                    types.String                                      `tfsdk:"path"`
	Negated                 types.Bool                                        `tfsdk:"negated"`
	EnableLatestUpdateCheck types.Bool                                        `tfsdk:"enable_latest_update_check"`
	CheckUpdatesWithinDays  types.Float64                                     `tfsdk:"check_updates_within_days"`
	Comparator              types.String                                      `tfsdk:"comparator"`
	Condition               *datasourceEndpointZtnaTagRuleRulesConditionModel `tfsdk:"condition"`
}

type datasourceEndpointZtnaTagRuleRulesConditionModel struct {
	Key        types.String `tfsdk:"key"`
	IsDword    types.Bool   `tfsdk:"is_dword"`
	Comparator types.String `tfsdk:"comparator"`
	Value      types.String `tfsdk:"value"`
}

type datasourceEndpointZtnaTagRuleLogicModel struct {
}

func (m *datasourceEndpointZtnaTagRuleRulesModel) flattenEndpointZtnaTagRuleRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaTagRuleRulesModel {
	if input == nil {
		return &datasourceEndpointZtnaTagRuleRulesModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaTagRuleRulesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	if v, ok := o["os"]; ok {
		m.Os = parseStringValue(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["service"]; ok {
		m.Service = parseStringValue(v)
	}

	if v, ok := o["account"]; ok {
		m.Account = parseStringValue(v)
	}

	if v, ok := o["matchType"]; ok {
		m.MatchType = parseStringValue(v)
	}

	if v, ok := o["subject"]; ok {
		m.Subject = parseStringValue(v)
	}

	if v, ok := o["issuer"]; ok {
		m.Issuer = parseStringValue(v)
	}

	if v, ok := o["content"]; ok {
		m.Content = parseStringValue(v)
	}

	if v, ok := o["path"]; ok {
		m.Path = parseStringValue(v)
	}

	if v, ok := o["negated"]; ok {
		m.Negated = parseBoolValue(v)
	}

	if v, ok := o["enableLatestUpdateCheck"]; ok {
		m.EnableLatestUpdateCheck = parseBoolValue(v)
	}

	if v, ok := o["checkUpdatesWithinDays"]; ok {
		m.CheckUpdatesWithinDays = parseFloat64Value(v)
	}

	if v, ok := o["comparator"]; ok {
		m.Comparator = parseStringValue(v)
	}

	if v, ok := o["condition"]; ok {
		m.Condition = m.Condition.flattenEndpointZtnaTagRuleRulesCondition(ctx, v, diags)
	}

	return m
}

func (s *datasourceEndpointZtnaTagRuleModel) flattenEndpointZtnaTagRuleRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointZtnaTagRuleRulesModel {
	if o == nil {
		return []datasourceEndpointZtnaTagRuleRulesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument rules is not type of []interface{}.", "")
		return []datasourceEndpointZtnaTagRuleRulesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointZtnaTagRuleRulesModel{}
	}

	values := make([]datasourceEndpointZtnaTagRuleRulesModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointZtnaTagRuleRulesModel
		if i < len(s.Rules) {
			m = s.Rules[i]
		}
		values[i] = *m.flattenEndpointZtnaTagRuleRules(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointZtnaTagRuleRulesConditionModel) flattenEndpointZtnaTagRuleRulesCondition(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaTagRuleRulesConditionModel {
	if input == nil {
		return &datasourceEndpointZtnaTagRuleRulesConditionModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaTagRuleRulesConditionModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["key"]; ok {
		m.Key = parseStringValue(v)
	}

	if v, ok := o["isDword"]; ok {
		m.IsDword = parseBoolValue(v)
	}

	if v, ok := o["comparator"]; ok {
		m.Comparator = parseStringValue(v)
	}

	if v, ok := o["value"]; ok {
		m.Value = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointZtnaTagRuleLogicModel) flattenEndpointZtnaTagRuleLogic(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaTagRuleLogicModel {
	if input == nil {
		return &datasourceEndpointZtnaTagRuleLogicModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaTagRuleLogicModel{}
	}

	return m
}
