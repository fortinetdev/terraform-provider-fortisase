// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityDlpProfile{}

func newDatasourceSecurityDlpProfile() datasource.DataSource {
	return &datasourceSecurityDlpProfile{}
}

type datasourceSecurityDlpProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityDlpProfileModel describes the datasource data model.
type datasourceSecurityDlpProfileModel struct {
	PrimaryKey types.String                                `tfsdk:"primary_key"`
	DlpRules   []datasourceSecurityDlpProfileDlpRulesModel `tfsdk:"dlp_rules"`
	Direction  types.String                                `tfsdk:"direction"`
}

func (r *datasourceSecurityDlpProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_profile"
}

func (r *datasourceSecurityDlpProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				MarkdownDescription: "The direction of the target resource.\nSupported values: internal-profiles, outbound-profiles.",
				Computed:            true,
				Optional:            true,
			},
			"dlp_rules": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("sensors", "mpip-label", "fingerprint", "none"),
							},
							Computed: true,
							Optional: true,
						},
						"severity": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("critical", "informational", "low", "medium", "high"),
							},
							Computed: true,
							Optional: true,
						},
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "monitor", "block"),
							},
							Computed: true,
							Optional: true,
						},
						"dlp_rule_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("file", "message"),
							},
							Computed: true,
							Optional: true,
						},
						"file_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("all", "specify"),
							},
							Computed: true,
							Optional: true,
						},
						"protocols": schema.SetAttribute{
							Validators: []validator.Set{
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf("smtp", "pop3", "imap", "http-get", "http-post", "ftp", "nntp", "cifs"),
								),
							},
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
						"sensitivities": schema.SetAttribute{
							Validators: []validator.Set{
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf("Warning", "Private", "Critical"),
								),
								setvalidator.SizeAtLeast(1),
							},
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
						"dlp_sensors": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"primary_key": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"datasource": schema.StringAttribute{
										Validators: []validator.String{
											stringvalidator.OneOf("security/dlp-sensors"),
										},
										Computed: true,
										Optional: true,
									},
								},
							},
							Computed: true,
							Optional: true,
						},
						"sensitivity_label": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("security/dlp-dictionaries"),
									},
									Computed: true,
									Optional: true,
								},
							},
							Computed: true,
							Optional: true,
						},
						"dlp_file_pattern": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("security/dlp-file-patterns"),
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
		},
	}
}

func (r *datasourceSecurityDlpProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_profile"
}

func (r *datasourceSecurityDlpProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityDlpProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityDlpProfileModel) refreshSecurityDlpProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["dlpRules"]; ok {
		m.DlpRules = m.flattenSecurityDlpProfileDlpRulesList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityDlpProfileModel) getURLObjectSecurityDlpProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		diags.AddWarning("\"direction\" is deprecated and may be removed in future.",
			"It is recommended to recreate the resource without \"direction\" to avoid unexpected behavior in future.",
		)
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityDlpProfileDlpRulesModel struct {
	PrimaryKey       types.String                                               `tfsdk:"primary_key"`
	DatasourceType   types.String                                               `tfsdk:"datasource_type"`
	Severity         types.String                                               `tfsdk:"severity"`
	Action           types.String                                               `tfsdk:"action"`
	DlpRuleType      types.String                                               `tfsdk:"dlp_rule_type"`
	FileType         types.String                                               `tfsdk:"file_type"`
	Protocols        types.Set                                                  `tfsdk:"protocols"`
	DlpSensors       []datasourceSecurityDlpProfileDlpRulesDlpSensorsModel      `tfsdk:"dlp_sensors"`
	SensitivityLabel *datasourceSecurityDlpProfileDlpRulesSensitivityLabelModel `tfsdk:"sensitivity_label"`
	Sensitivities    types.Set                                                  `tfsdk:"sensitivities"`
	DlpFilePattern   *datasourceSecurityDlpProfileDlpRulesDlpFilePatternModel   `tfsdk:"dlp_file_pattern"`
}

type datasourceSecurityDlpProfileDlpRulesDlpSensorsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityDlpProfileDlpRulesSensitivityLabelModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityDlpProfileDlpRulesDlpFilePatternModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityDlpProfileDlpRulesModel) flattenSecurityDlpProfileDlpRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpProfileDlpRulesModel {
	if input == nil {
		return &datasourceSecurityDlpProfileDlpRulesModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpProfileDlpRulesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasourceType"]; ok {
		m.DatasourceType = parseStringValue(v)
	}

	if v, ok := o["severity"]; ok {
		m.Severity = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["dlpRuleType"]; ok {
		m.DlpRuleType = parseStringValue(v)
	}

	if v, ok := o["fileType"]; ok {
		m.FileType = parseStringValue(v)
	}

	if v, ok := o["protocols"]; ok {
		m.Protocols = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["dlpSensors"]; ok {
		m.DlpSensors = m.flattenSecurityDlpProfileDlpRulesDlpSensorsList(ctx, v, diags)
	}

	if v, ok := o["sensitivityLabel"]; ok {
		m.SensitivityLabel = m.SensitivityLabel.flattenSecurityDlpProfileDlpRulesSensitivityLabel(ctx, v, diags)
	}

	if v, ok := o["sensitivities"]; ok {
		m.Sensitivities = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["dlpFilePattern"]; ok {
		m.DlpFilePattern = m.DlpFilePattern.flattenSecurityDlpProfileDlpRulesDlpFilePattern(ctx, v, diags)
	}

	return m
}

func (s *datasourceSecurityDlpProfileModel) flattenSecurityDlpProfileDlpRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDlpProfileDlpRulesModel {
	if o == nil {
		return []datasourceSecurityDlpProfileDlpRulesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument dlp_rules is not type of []interface{}.", "")
		return []datasourceSecurityDlpProfileDlpRulesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDlpProfileDlpRulesModel{}
	}

	values := make([]datasourceSecurityDlpProfileDlpRulesModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDlpProfileDlpRulesModel
		values[i] = *m.flattenSecurityDlpProfileDlpRules(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityDlpProfileDlpRulesDlpSensorsModel) flattenSecurityDlpProfileDlpRulesDlpSensors(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpProfileDlpRulesDlpSensorsModel {
	if input == nil {
		return &datasourceSecurityDlpProfileDlpRulesDlpSensorsModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpProfileDlpRulesDlpSensorsModel{}
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

func (s *datasourceSecurityDlpProfileDlpRulesModel) flattenSecurityDlpProfileDlpRulesDlpSensorsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecurityDlpProfileDlpRulesDlpSensorsModel {
	if o == nil {
		return []datasourceSecurityDlpProfileDlpRulesDlpSensorsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument dlp_sensors is not type of []interface{}.", "")
		return []datasourceSecurityDlpProfileDlpRulesDlpSensorsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecurityDlpProfileDlpRulesDlpSensorsModel{}
	}

	values := make([]datasourceSecurityDlpProfileDlpRulesDlpSensorsModel, len(l))
	for i, ele := range l {
		var m datasourceSecurityDlpProfileDlpRulesDlpSensorsModel
		values[i] = *m.flattenSecurityDlpProfileDlpRulesDlpSensors(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecurityDlpProfileDlpRulesSensitivityLabelModel) flattenSecurityDlpProfileDlpRulesSensitivityLabel(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpProfileDlpRulesSensitivityLabelModel {
	if input == nil {
		return &datasourceSecurityDlpProfileDlpRulesSensitivityLabelModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpProfileDlpRulesSensitivityLabelModel{}
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

func (m *datasourceSecurityDlpProfileDlpRulesDlpFilePatternModel) flattenSecurityDlpProfileDlpRulesDlpFilePattern(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpProfileDlpRulesDlpFilePatternModel {
	if input == nil {
		return &datasourceSecurityDlpProfileDlpRulesDlpFilePatternModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpProfileDlpRulesDlpFilePatternModel{}
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
