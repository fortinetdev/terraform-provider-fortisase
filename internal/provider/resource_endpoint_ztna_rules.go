// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
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
var _ resource.Resource = &resourceEndpointZtnaRules{}

func newResourceEndpointZtnaRules() resource.Resource {
	return &resourceEndpointZtnaRules{}
}

type resourceEndpointZtnaRules struct {
	fortiClient *FortiClient
}

// resourceEndpointZtnaRulesModel describes the resource data model.
type resourceEndpointZtnaRulesModel struct {
	ID         types.String                          `tfsdk:"id"`
	PrimaryKey types.String                          `tfsdk:"primary_key"`
	Status     types.String                          `tfsdk:"status"`
	Tag        *resourceEndpointZtnaRulesTagModel    `tfsdk:"tag"`
	Comments   types.String                          `tfsdk:"comments"`
	Rules      []resourceEndpointZtnaRulesRulesModel `tfsdk:"rules"`
	Logic      *resourceEndpointZtnaRulesLogicModel  `tfsdk:"logic"`
}

func (r *resourceEndpointZtnaRules) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_rules"
}

func (r *resourceEndpointZtnaRules) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *resourceEndpointZtnaRules) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceEndpointZtnaRules) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointZtnaRulesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointZtnaRules(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointZtnaRules(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateEndpointZtnaRules(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectEndpointZtnaRules(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
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

func (r *resourceEndpointZtnaRules) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointZtnaRulesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointZtnaRulesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointZtnaRules(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointZtnaRules(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateEndpointZtnaRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointZtnaRules(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
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

func (r *resourceEndpointZtnaRules) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointZtnaRulesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointZtnaRules(ctx, "delete", diags))

	err := c.DeleteEndpointZtnaRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceEndpointZtnaRules) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointZtnaRulesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointZtnaRules(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
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

func (r *resourceEndpointZtnaRules) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointZtnaRulesModel) refreshEndpointZtnaRules(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceEndpointZtnaRulesModel) getCreateObjectEndpointZtnaRules(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Tag != nil && !isZeroStruct(*data.Tag) {
		result["tag"] = data.Tag.expandEndpointZtnaRulesTag(ctx, diags)
	}

	if !data.Comments.IsNull() {
		result["comments"] = data.Comments.ValueString()
	}

	result["rules"] = data.expandEndpointZtnaRulesRulesList(ctx, data.Rules, diags)

	if data.Logic != nil && !isZeroStruct(*data.Logic) {
		result["logic"] = data.Logic.expandEndpointZtnaRulesLogic(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointZtnaRulesModel) getUpdateObjectEndpointZtnaRules(ctx context.Context, state resourceEndpointZtnaRulesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Tag != nil && !isSameStruct(data.Tag, state.Tag) {
		result["tag"] = data.Tag.expandEndpointZtnaRulesTag(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.Equal(state.Comments) {
		result["comments"] = data.Comments.ValueString()
	}

	if len(data.Rules) > 0 || !isSameStruct(data.Rules, state.Rules) {
		result["rules"] = data.expandEndpointZtnaRulesRulesList(ctx, data.Rules, diags)
	}

	if data.Logic != nil && !isSameStruct(data.Logic, state.Logic) {
		result["logic"] = data.Logic.expandEndpointZtnaRulesLogic(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointZtnaRulesModel) getURLObjectEndpointZtnaRules(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointZtnaRulesTagModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointZtnaRulesRulesModel struct {
	Id                      types.Float64                                 `tfsdk:"id"`
	Os                      types.String                                  `tfsdk:"os"`
	Type                    types.String                                  `tfsdk:"type"`
	Service                 types.String                                  `tfsdk:"service"`
	Account                 types.String                                  `tfsdk:"account"`
	MatchType               types.String                                  `tfsdk:"match_type"`
	Subject                 types.String                                  `tfsdk:"subject"`
	Issuer                  types.String                                  `tfsdk:"issuer"`
	Path                    types.String                                  `tfsdk:"path"`
	Negated                 types.Bool                                    `tfsdk:"negated"`
	EnableLatestUpdateCheck types.Bool                                    `tfsdk:"enable_latest_update_check"`
	CheckUpdatesWithinDays  types.Float64                                 `tfsdk:"check_updates_within_days"`
	Comparator              types.String                                  `tfsdk:"comparator"`
	Condition               *resourceEndpointZtnaRulesRulesConditionModel `tfsdk:"condition"`
	Content                 types.String                                  `tfsdk:"content"`
}

type resourceEndpointZtnaRulesRulesConditionModel struct {
	Key        types.String `tfsdk:"key"`
	IsDword    types.Bool   `tfsdk:"is_dword"`
	Comparator types.String `tfsdk:"comparator"`
	Value      types.String `tfsdk:"value"`
}

type resourceEndpointZtnaRulesLogicModel struct {
	Windows types.String `tfsdk:"windows"`
	Macos   types.String `tfsdk:"macos"`
	Linux   types.String `tfsdk:"linux"`
	Ios     types.String `tfsdk:"ios"`
	Android types.String `tfsdk:"android"`
}

func (m *resourceEndpointZtnaRulesTagModel) flattenEndpointZtnaRulesTag(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaRulesTagModel {
	if input == nil {
		return &resourceEndpointZtnaRulesTagModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaRulesTagModel{}
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

func (m *resourceEndpointZtnaRulesRulesModel) flattenEndpointZtnaRulesRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaRulesRulesModel {
	if input == nil {
		return &resourceEndpointZtnaRulesRulesModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaRulesRulesModel{}
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

func (s *resourceEndpointZtnaRulesModel) flattenEndpointZtnaRulesRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointZtnaRulesRulesModel {
	if o == nil {
		return []resourceEndpointZtnaRulesRulesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument rules is not type of []interface{}.", "")
		return []resourceEndpointZtnaRulesRulesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointZtnaRulesRulesModel{}
	}

	values := make([]resourceEndpointZtnaRulesRulesModel, len(l))
	for i, ele := range l {
		var m resourceEndpointZtnaRulesRulesModel
		values[i] = *m.flattenEndpointZtnaRulesRules(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointZtnaRulesRulesConditionModel) flattenEndpointZtnaRulesRulesCondition(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaRulesRulesConditionModel {
	if input == nil {
		return &resourceEndpointZtnaRulesRulesConditionModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaRulesRulesConditionModel{}
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

func (m *resourceEndpointZtnaRulesLogicModel) flattenEndpointZtnaRulesLogic(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaRulesLogicModel {
	if input == nil {
		return &resourceEndpointZtnaRulesLogicModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaRulesLogicModel{}
	}

	return m
}

func (data *resourceEndpointZtnaRulesTagModel) expandEndpointZtnaRulesTag(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceEndpointZtnaRulesRulesModel) expandEndpointZtnaRulesRules(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() {
		result["id"] = data.Id.ValueFloat64()
	}

	if !data.Os.IsNull() {
		result["os"] = data.Os.ValueString()
	}

	if !data.Type.IsNull() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Service.IsNull() {
		result["service"] = data.Service.ValueString()
	}

	if !data.Account.IsNull() {
		result["account"] = data.Account.ValueString()
	}

	if !data.MatchType.IsNull() {
		result["matchType"] = data.MatchType.ValueString()
	}

	if !data.Subject.IsNull() {
		result["subject"] = data.Subject.ValueString()
	}

	if !data.Issuer.IsNull() {
		result["issuer"] = data.Issuer.ValueString()
	}

	if !data.Path.IsNull() {
		result["path"] = data.Path.ValueString()
	}

	if !data.Negated.IsNull() {
		result["negated"] = data.Negated.ValueBool()
	}

	if !data.EnableLatestUpdateCheck.IsNull() {
		result["enableLatestUpdateCheck"] = data.EnableLatestUpdateCheck.ValueBool()
	}

	if !data.CheckUpdatesWithinDays.IsNull() {
		result["checkUpdatesWithinDays"] = data.CheckUpdatesWithinDays.ValueFloat64()
	}

	if !data.Comparator.IsNull() {
		result["comparator"] = data.Comparator.ValueString()
	}

	if data.Condition != nil && !isZeroStruct(*data.Condition) {
		result["condition"] = data.Condition.expandEndpointZtnaRulesRulesCondition(ctx, diags)
	}

	if !data.Content.IsNull() {
		result["content"] = data.Content.ValueString()
	}

	return result
}

func (s *resourceEndpointZtnaRulesModel) expandEndpointZtnaRulesRulesList(ctx context.Context, l []resourceEndpointZtnaRulesRulesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointZtnaRulesRules(ctx, diags)
	}
	return result
}

func (data *resourceEndpointZtnaRulesRulesConditionModel) expandEndpointZtnaRulesRulesCondition(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Key.IsNull() {
		result["key"] = data.Key.ValueString()
	}

	if !data.IsDword.IsNull() {
		result["isDword"] = data.IsDword.ValueBool()
	}

	if !data.Comparator.IsNull() {
		result["comparator"] = data.Comparator.ValueString()
	}

	if !data.Value.IsNull() {
		result["value"] = data.Value.ValueString()
	}

	return result
}

func (data *resourceEndpointZtnaRulesLogicModel) expandEndpointZtnaRulesLogic(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Windows.IsNull() {
		result["windows"] = data.Windows.ValueString()
	}

	if !data.Macos.IsNull() {
		result["macos"] = data.Macos.ValueString()
	}

	if !data.Linux.IsNull() {
		result["linux"] = data.Linux.ValueString()
	}

	if !data.Ios.IsNull() {
		result["ios"] = data.Ios.ValueString()
	}

	if !data.Android.IsNull() {
		result["android"] = data.Android.ValueString()
	}

	return result
}
