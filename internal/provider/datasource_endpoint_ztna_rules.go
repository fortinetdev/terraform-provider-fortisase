// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointZtnaRules{}

func newDatasourceEndpointZtnaRules() datasource.DataSource {
	return &datasourceEndpointZtnaRules{}
}

type datasourceEndpointZtnaRules struct {
	fortiClient *FortiClient
}

// datasourceEndpointZtnaRulesModel describes the datasource data model.
type datasourceEndpointZtnaRulesModel struct {
	PrimaryKey types.String                            `tfsdk:"primary_key"`
	Status     types.String                            `tfsdk:"status"`
	Tag        *datasourceEndpointZtnaRulesTagModel    `tfsdk:"tag"`
	Comments   types.String                            `tfsdk:"comments"`
	Rules      []datasourceEndpointZtnaRulesRulesModel `tfsdk:"rules"`
	Logic      *datasourceEndpointZtnaRulesLogicModel  `tfsdk:"logic"`
}

func (r *datasourceEndpointZtnaRules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_rules"
}

func (r *datasourceEndpointZtnaRules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
				Required: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"comments": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(1000),
				},
				Computed: true,
				Optional: true,
			},
			"tag": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("endpoint/ztna-tags"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"rules": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validator.AtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"os": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("windows", "macos", "linux", "ios", "android"),
							},
							Computed: true,
							Optional: true,
						},
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("ad-groups", "anti-virus", "certificate", "file", "logged-in-domain", "running-process", "registry-key", "os-version", "sandbox-detection", "vulnerable-devices", "windows-security", "user-identity", "ems-management", "security", "ip-range", "on-fabric-status", "fct-version", "security-status", "cve", "crowdstrike-zta-score"),
							},
							Computed: true,
							Optional: true,
						},
						"service": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("Google", "LinkedIn", "Salesforce", "Custom"),
							},
							Computed: true,
							Optional: true,
						},
						"account": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 256),
							},
							Computed: true,
							Optional: true,
						},
						"match_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("simple", "regex", "wildcard"),
							},
							Computed: true,
							Optional: true,
						},
						"subject": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 256),
							},
							Computed: true,
							Optional: true,
						},
						"issuer": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 256),
							},
							Computed: true,
							Optional: true,
						},
						"path": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"negated": schema.BoolAttribute{
							Computed: true,
							Optional: true,
						},
						"enable_latest_update_check": schema.BoolAttribute{
							Computed: true,
							Optional: true,
						},
						"check_updates_within_days": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validator.Between(1, 3653),
							},
							Computed: true,
							Optional: true,
						},
						"comparator": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("=", ">", "<", ">=", "<="),
							},
							Computed: true,
							Optional: true,
						},
						"content": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"condition": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthBetween(1, 256),
									},
									Computed: true,
									Optional: true,
								},
								"is_dword": schema.BoolAttribute{
									Computed: true,
									Optional: true,
								},
								"comparator": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("=", "!=", ">", ">=", "<", "<="),
									},
									Computed: true,
									Optional: true,
								},
								"value": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.LengthBetween(1, 256),
									},
									Computed: true,
									Optional: true,
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
			"logic": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"windows": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"macos": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"linux": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"ios": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"android": schema.StringAttribute{
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

func (r *datasourceEndpointZtnaRules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceEndpointZtnaRules) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointZtnaRulesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointZtnaRules(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaRules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointZtnaRulesModel) refreshEndpointZtnaRules(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["tag"]; ok {
		m.Tag = m.Tag.flattenEndpointZtnaRulesTag(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["rules"]; ok {
		m.Rules = m.flattenEndpointZtnaRulesRulesList(ctx, v, &diags)
	}

	if v, ok := o["logic"]; ok {
		m.Logic = m.Logic.flattenEndpointZtnaRulesLogic(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointZtnaRulesModel) getURLObjectEndpointZtnaRules(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointZtnaRulesTagModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceEndpointZtnaRulesRulesModel struct {
	Id                      types.Float64                                   `tfsdk:"id"`
	Os                      types.String                                    `tfsdk:"os"`
	Type                    types.String                                    `tfsdk:"type"`
	Service                 types.String                                    `tfsdk:"service"`
	Account                 types.String                                    `tfsdk:"account"`
	MatchType               types.String                                    `tfsdk:"match_type"`
	Subject                 types.String                                    `tfsdk:"subject"`
	Issuer                  types.String                                    `tfsdk:"issuer"`
	Path                    types.String                                    `tfsdk:"path"`
	Negated                 types.Bool                                      `tfsdk:"negated"`
	EnableLatestUpdateCheck types.Bool                                      `tfsdk:"enable_latest_update_check"`
	CheckUpdatesWithinDays  types.Float64                                   `tfsdk:"check_updates_within_days"`
	Comparator              types.String                                    `tfsdk:"comparator"`
	Condition               *datasourceEndpointZtnaRulesRulesConditionModel `tfsdk:"condition"`
	Content                 types.String                                    `tfsdk:"content"`
}

type datasourceEndpointZtnaRulesRulesConditionModel struct {
	Key        types.String `tfsdk:"key"`
	IsDword    types.Bool   `tfsdk:"is_dword"`
	Comparator types.String `tfsdk:"comparator"`
	Value      types.String `tfsdk:"value"`
}

type datasourceEndpointZtnaRulesLogicModel struct {
	Windows types.String `tfsdk:"windows"`
	Macos   types.String `tfsdk:"macos"`
	Linux   types.String `tfsdk:"linux"`
	Ios     types.String `tfsdk:"ios"`
	Android types.String `tfsdk:"android"`
}

func (m *datasourceEndpointZtnaRulesTagModel) flattenEndpointZtnaRulesTag(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaRulesTagModel {
	if input == nil {
		return &datasourceEndpointZtnaRulesTagModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaRulesTagModel{}
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

func (m *datasourceEndpointZtnaRulesRulesModel) flattenEndpointZtnaRulesRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaRulesRulesModel {
	if input == nil {
		return &datasourceEndpointZtnaRulesRulesModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaRulesRulesModel{}
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
		m.Condition = m.Condition.flattenEndpointZtnaRulesRulesCondition(ctx, v, diags)
	}

	if v, ok := o["content"]; ok {
		m.Content = parseStringValue(v)
	}

	return m
}

func (s *datasourceEndpointZtnaRulesModel) flattenEndpointZtnaRulesRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointZtnaRulesRulesModel {
	if o == nil {
		return []datasourceEndpointZtnaRulesRulesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument rules is not type of []interface{}.", "")
		return []datasourceEndpointZtnaRulesRulesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointZtnaRulesRulesModel{}
	}

	values := make([]datasourceEndpointZtnaRulesRulesModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointZtnaRulesRulesModel
		values[i] = *m.flattenEndpointZtnaRulesRules(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointZtnaRulesRulesConditionModel) flattenEndpointZtnaRulesRulesCondition(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaRulesRulesConditionModel {
	if input == nil {
		return &datasourceEndpointZtnaRulesRulesConditionModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaRulesRulesConditionModel{}
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

func (m *datasourceEndpointZtnaRulesLogicModel) flattenEndpointZtnaRulesLogic(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointZtnaRulesLogicModel {
	if input == nil {
		return &datasourceEndpointZtnaRulesLogicModel{}
	}
	if m == nil {
		m = &datasourceEndpointZtnaRulesLogicModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["windows"]; ok {
		m.Windows = parseStringValue(v)
	}

	if v, ok := o["macos"]; ok {
		m.Macos = parseStringValue(v)
	}

	if v, ok := o["linux"]; ok {
		m.Linux = parseStringValue(v)
	}

	if v, ok := o["ios"]; ok {
		m.Ios = parseStringValue(v)
	}

	if v, ok := o["android"]; ok {
		m.Android = parseStringValue(v)
	}

	return m
}
