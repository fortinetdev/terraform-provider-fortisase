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
	"strings"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSecurityDlpProfile{}

func newResourceSecurityDlpProfile() resource.Resource {
	return &resourceSecurityDlpProfile{}
}

type resourceSecurityDlpProfile struct {
	fortiClient *FortiClient
}

// resourceSecurityDlpProfileModel describes the resource data model.
type resourceSecurityDlpProfileModel struct {
	ID         types.String                              `tfsdk:"id"`
	PrimaryKey types.String                              `tfsdk:"primary_key"`
	DlpRules   []resourceSecurityDlpProfileDlpRulesModel `tfsdk:"dlp_rules"`
	Direction  types.String                              `tfsdk:"direction"`
}

func (r *resourceSecurityDlpProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_profile"
}

func (r *resourceSecurityDlpProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"direction": schema.StringAttribute{
				Description: "The direction of the target resource.",
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				Computed: true,
				Optional: true,
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

func (r *resourceSecurityDlpProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityDlpProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityDlpProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecurityDlpProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecurityDlpProfile(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDlpProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDlpProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDlpProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDlpProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityDlpProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDlpProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecurityDlpProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDlpProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityDlpProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDlpProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (m *resourceSecurityDlpProfileModel) refreshSecurityDlpProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["dlpRules"]; ok {
		m.DlpRules = m.flattenSecurityDlpProfileDlpRulesList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecurityDlpProfileModel) getCreateObjectSecurityDlpProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	result["dlpRules"] = data.expandSecurityDlpProfileDlpRulesList(ctx, data.DlpRules, diags)

	return &result
}

func (data *resourceSecurityDlpProfileModel) getUpdateObjectSecurityDlpProfile(ctx context.Context, state resourceSecurityDlpProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if len(data.DlpRules) > 0 || !isSameStruct(data.DlpRules, state.DlpRules) {
		result["dlpRules"] = data.expandSecurityDlpProfileDlpRulesList(ctx, data.DlpRules, diags)
	}

	return &result
}

func (data *resourceSecurityDlpProfileModel) getURLObjectSecurityDlpProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecurityDlpProfileDlpRulesModel struct {
	PrimaryKey       types.String                                             `tfsdk:"primary_key"`
	DatasourceType   types.String                                             `tfsdk:"datasource_type"`
	Severity         types.String                                             `tfsdk:"severity"`
	Action           types.String                                             `tfsdk:"action"`
	DlpRuleType      types.String                                             `tfsdk:"dlp_rule_type"`
	FileType         types.String                                             `tfsdk:"file_type"`
	Protocols        types.Set                                                `tfsdk:"protocols"`
	DlpSensors       []resourceSecurityDlpProfileDlpRulesDlpSensorsModel      `tfsdk:"dlp_sensors"`
	SensitivityLabel *resourceSecurityDlpProfileDlpRulesSensitivityLabelModel `tfsdk:"sensitivity_label"`
	Sensitivities    types.Set                                                `tfsdk:"sensitivities"`
	DlpFilePattern   *resourceSecurityDlpProfileDlpRulesDlpFilePatternModel   `tfsdk:"dlp_file_pattern"`
}

type resourceSecurityDlpProfileDlpRulesDlpSensorsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityDlpProfileDlpRulesSensitivityLabelModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityDlpProfileDlpRulesDlpFilePatternModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityDlpProfileDlpRulesModel) flattenSecurityDlpProfileDlpRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpProfileDlpRulesModel {
	if input == nil {
		return &resourceSecurityDlpProfileDlpRulesModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpProfileDlpRulesModel{}
	}
	o := input.(map[string]interface{})
	m.Protocols = types.SetNull(types.StringType)
	m.Sensitivities = types.SetNull(types.StringType)

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

func (s *resourceSecurityDlpProfileModel) flattenSecurityDlpProfileDlpRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDlpProfileDlpRulesModel {
	if o == nil {
		return []resourceSecurityDlpProfileDlpRulesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument dlp_rules is not type of []interface{}.", "")
		return []resourceSecurityDlpProfileDlpRulesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDlpProfileDlpRulesModel{}
	}

	values := make([]resourceSecurityDlpProfileDlpRulesModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDlpProfileDlpRulesModel
		values[i] = *m.flattenSecurityDlpProfileDlpRules(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDlpProfileDlpRulesDlpSensorsModel) flattenSecurityDlpProfileDlpRulesDlpSensors(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpProfileDlpRulesDlpSensorsModel {
	if input == nil {
		return &resourceSecurityDlpProfileDlpRulesDlpSensorsModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpProfileDlpRulesDlpSensorsModel{}
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

func (s *resourceSecurityDlpProfileDlpRulesModel) flattenSecurityDlpProfileDlpRulesDlpSensorsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecurityDlpProfileDlpRulesDlpSensorsModel {
	if o == nil {
		return []resourceSecurityDlpProfileDlpRulesDlpSensorsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument dlp_sensors is not type of []interface{}.", "")
		return []resourceSecurityDlpProfileDlpRulesDlpSensorsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecurityDlpProfileDlpRulesDlpSensorsModel{}
	}

	values := make([]resourceSecurityDlpProfileDlpRulesDlpSensorsModel, len(l))
	for i, ele := range l {
		var m resourceSecurityDlpProfileDlpRulesDlpSensorsModel
		values[i] = *m.flattenSecurityDlpProfileDlpRulesDlpSensors(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecurityDlpProfileDlpRulesSensitivityLabelModel) flattenSecurityDlpProfileDlpRulesSensitivityLabel(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpProfileDlpRulesSensitivityLabelModel {
	if input == nil {
		return &resourceSecurityDlpProfileDlpRulesSensitivityLabelModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpProfileDlpRulesSensitivityLabelModel{}
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

func (m *resourceSecurityDlpProfileDlpRulesDlpFilePatternModel) flattenSecurityDlpProfileDlpRulesDlpFilePattern(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityDlpProfileDlpRulesDlpFilePatternModel {
	if input == nil {
		return &resourceSecurityDlpProfileDlpRulesDlpFilePatternModel{}
	}
	if m == nil {
		m = &resourceSecurityDlpProfileDlpRulesDlpFilePatternModel{}
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

func (data *resourceSecurityDlpProfileDlpRulesModel) expandSecurityDlpProfileDlpRules(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.DatasourceType.IsNull() {
		result["datasourceType"] = data.DatasourceType.ValueString()
	}

	if !data.Severity.IsNull() {
		result["severity"] = data.Severity.ValueString()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.DlpRuleType.IsNull() {
		result["dlpRuleType"] = data.DlpRuleType.ValueString()
	}

	if !data.FileType.IsNull() {
		result["fileType"] = data.FileType.ValueString()
	}

	if !data.Protocols.IsNull() {
		result["protocols"] = expandSetToStringList(data.Protocols)
	}

	if len(data.DlpSensors) > 0 {
		result["dlpSensors"] = data.expandSecurityDlpProfileDlpRulesDlpSensorsList(ctx, data.DlpSensors, diags)
	}

	if data.SensitivityLabel != nil && !isZeroStruct(*data.SensitivityLabel) {
		result["sensitivityLabel"] = data.SensitivityLabel.expandSecurityDlpProfileDlpRulesSensitivityLabel(ctx, diags)
	}

	if !data.Sensitivities.IsNull() {
		result["sensitivities"] = expandSetToStringList(data.Sensitivities)
	}

	if data.DlpFilePattern != nil && !isZeroStruct(*data.DlpFilePattern) {
		result["dlpFilePattern"] = data.DlpFilePattern.expandSecurityDlpProfileDlpRulesDlpFilePattern(ctx, diags)
	}

	return result
}

func (s *resourceSecurityDlpProfileModel) expandSecurityDlpProfileDlpRulesList(ctx context.Context, l []resourceSecurityDlpProfileDlpRulesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDlpProfileDlpRules(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDlpProfileDlpRulesDlpSensorsModel) expandSecurityDlpProfileDlpRulesDlpSensors(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecurityDlpProfileDlpRulesModel) expandSecurityDlpProfileDlpRulesDlpSensorsList(ctx context.Context, l []resourceSecurityDlpProfileDlpRulesDlpSensorsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecurityDlpProfileDlpRulesDlpSensors(ctx, diags)
	}
	return result
}

func (data *resourceSecurityDlpProfileDlpRulesSensitivityLabelModel) expandSecurityDlpProfileDlpRulesSensitivityLabel(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityDlpProfileDlpRulesDlpFilePatternModel) expandSecurityDlpProfileDlpRulesDlpFilePattern(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
