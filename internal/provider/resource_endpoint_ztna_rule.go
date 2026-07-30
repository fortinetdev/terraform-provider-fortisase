// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
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
var _ resource.Resource = &resourceEndpointZtnaRule{}
var _ resource.ResourceWithMoveState = &resourceEndpointZtnaRule{}

func newResourceEndpointZtnaRule() resource.Resource {
	return &resourceEndpointZtnaRule{}
}

type resourceEndpointZtnaRule struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointZtnaRuleModel describes the resource data model.
type resourceEndpointZtnaRuleModel struct {
	ID         types.String                         `tfsdk:"id"`
	PrimaryKey types.String                         `tfsdk:"primary_key"`
	Status     types.String                         `tfsdk:"status"`
	Tag        *resourceEndpointZtnaRuleTagModel    `tfsdk:"tag"`
	Comments   types.String                         `tfsdk:"comments"`
	Rules      []resourceEndpointZtnaRuleRulesModel `tfsdk:"rules"`
	Logic      *resourceEndpointZtnaRuleLogicModel  `tfsdk:"logic"`
}

func (r *resourceEndpointZtnaRule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_ztna_rule"
}

func (r *resourceEndpointZtnaRule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ZTNA Rule Resource API V2 for FortiSASE. This resource is restricted to EMS version: 7.2.",
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
					stringvalidatorwarning.LengthBetween(1, 128),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"comments": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtMost(1000),
				},
				Computed: true,
				Optional: true,
			},
			"tag": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("endpoint/ztna-tags"),
						},
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "The property 'logic' is required when 'rules' are modified; otherwise, 'logic' will be set to a default value. This resource is restricted to EMS version: 7.2.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validatorwarning.AtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"os": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("windows", "macos", "linux", "ios", "android"),
							},
							Computed: true,
							Optional: true,
						},
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("ad-groups", "anti-virus", "certificate", "file", "logged-in-domain", "running-process", "registry-key", "os-version", "sandbox-detection", "vulnerable-devices", "windows-security", "user-identity", "ems-management", "security", "ip-range", "on-fabric-status", "fct-version", "security-status", "cve", "crowdstrike-zta-score"),
							},
							Computed: true,
							Optional: true,
						},
						"service": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("Google", "LinkedIn", "Salesforce", "Custom"),
							},
							Computed: true,
							Optional: true,
						},
						"account": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthBetween(1, 256),
							},
							Computed: true,
							Optional: true,
						},
						"match_type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("simple", "regex", "wildcard"),
							},
							Computed: true,
							Optional: true,
						},
						"subject": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthBetween(1, 256),
							},
							Computed: true,
							Optional: true,
						},
						"issuer": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthBetween(1, 256),
							},
							Computed: true,
							Optional: true,
						},
						"path": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.LengthAtLeast(1),
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
								float64validatorwarning.Between(1, 3653),
							},
							Computed: true,
							Optional: true,
						},
						"comparator": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("=", ">", "<", ">=", "<="),
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
										stringvalidatorwarning.LengthBetween(1, 256),
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
										stringvalidatorwarning.OneOf("=", "!=", ">", ">=", "<", "<="),
									},
									Computed: true,
									Optional: true,
								},
								"value": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.LengthBetween(1, 256),
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
				MarkdownDescription: "The property 'logic' is required when 'rules' are modified; otherwise, 'logic' will be set to a default value. This resource is restricted to EMS version: 7.2.",
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

func (r *resourceEndpointZtnaRule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
		"EMS": {"7.2"},
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
	r.resourceName = "fortisase_endpoint_ztna_rule"
}
func (r *resourceEndpointZtnaRule) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_endpoint_ztna_rules" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceEndpointZtnaRuleModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceEndpointZtnaRule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointZtnaRules")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointZtnaRuleModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectEndpointZtnaRule(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointZtnaRule(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateEndpointZtnaRules(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectEndpointZtnaRule(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointZtnaRule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointZtnaRules")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointZtnaRuleModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointZtnaRuleModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointZtnaRule(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointZtnaRule(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointZtnaRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointZtnaRule(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaRules(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointZtnaRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointZtnaRule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointZtnaRules")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceEndpointZtnaRuleModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointZtnaRule(ctx, "delete", diags))

	output, err := c.DeleteEndpointZtnaRules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceEndpointZtnaRule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointZtnaRuleModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointZtnaRule(ctx, "read", diags))

	read_output, err := c.ReadEndpointZtnaRules(&input_model)
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

	diags.Append(data.refreshEndpointZtnaRule(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointZtnaRule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceEndpointZtnaRuleModel) refreshEndpointZtnaRule(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["tag"]; ok {
		m.Tag = m.Tag.flattenEndpointZtnaRuleTag(ctx, v, &diags)
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["rules"]; ok {
		m.Rules = m.flattenEndpointZtnaRuleRulesList(ctx, v, &diags)
	}

	if v, ok := o["logic"]; ok {
		m.Logic = m.Logic.flattenEndpointZtnaRuleLogic(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointZtnaRuleModel) getCreateObjectEndpointZtnaRule(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	result["tag"] = nil
	if data.Tag != nil && !isZeroStruct(*data.Tag) {
		result["tag"] = data.Tag.expandEndpointZtnaRuleTag(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	result["rules"] = data.expandEndpointZtnaRuleRulesList(ctx, data.Rules, diags)

	if data.Logic != nil && !isZeroStruct(*data.Logic) {
		result["logic"] = data.Logic.expandEndpointZtnaRuleLogic(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointZtnaRuleModel) getUpdateObjectEndpointZtnaRule(ctx context.Context, state resourceEndpointZtnaRuleModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Status.IsNull() && !data.Status.IsUnknown() {
		result["status"] = data.Status.ValueString()
	}

	result["tag"] = nil
	if data.Tag != nil && !isZeroStruct(*data.Tag) {
		result["tag"] = data.Tag.expandEndpointZtnaRuleTag(ctx, diags)
	}

	if !data.Comments.IsNull() && !data.Comments.IsUnknown() {
		result["comments"] = data.Comments.ValueString()
	}

	if data.Rules != nil {
		result["rules"] = data.expandEndpointZtnaRuleRulesList(ctx, data.Rules, diags)
	}

	if data.Logic != nil {
		result["logic"] = data.Logic.expandEndpointZtnaRuleLogic(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointZtnaRuleModel) getURLObjectEndpointZtnaRule(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointZtnaRuleTagModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointZtnaRuleRulesModel struct {
	Id                      types.Float64                                `tfsdk:"id"`
	Os                      types.String                                 `tfsdk:"os"`
	Type                    types.String                                 `tfsdk:"type"`
	Service                 types.String                                 `tfsdk:"service"`
	Account                 types.String                                 `tfsdk:"account"`
	MatchType               types.String                                 `tfsdk:"match_type"`
	Subject                 types.String                                 `tfsdk:"subject"`
	Issuer                  types.String                                 `tfsdk:"issuer"`
	Path                    types.String                                 `tfsdk:"path"`
	Negated                 types.Bool                                   `tfsdk:"negated"`
	EnableLatestUpdateCheck types.Bool                                   `tfsdk:"enable_latest_update_check"`
	CheckUpdatesWithinDays  types.Float64                                `tfsdk:"check_updates_within_days"`
	Comparator              types.String                                 `tfsdk:"comparator"`
	Condition               *resourceEndpointZtnaRuleRulesConditionModel `tfsdk:"condition"`
	Content                 types.String                                 `tfsdk:"content"`
}

type resourceEndpointZtnaRuleRulesConditionModel struct {
	Key        types.String `tfsdk:"key"`
	IsDword    types.Bool   `tfsdk:"is_dword"`
	Comparator types.String `tfsdk:"comparator"`
	Value      types.String `tfsdk:"value"`
}

type resourceEndpointZtnaRuleLogicModel struct {
	Windows types.String `tfsdk:"windows"`
	Macos   types.String `tfsdk:"macos"`
	Linux   types.String `tfsdk:"linux"`
	Ios     types.String `tfsdk:"ios"`
	Android types.String `tfsdk:"android"`
}

func (m *resourceEndpointZtnaRuleTagModel) flattenEndpointZtnaRuleTag(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaRuleTagModel {
	if input == nil {
		return &resourceEndpointZtnaRuleTagModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaRuleTagModel{}
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

func (m *resourceEndpointZtnaRuleRulesModel) flattenEndpointZtnaRuleRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaRuleRulesModel {
	if input == nil {
		return &resourceEndpointZtnaRuleRulesModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaRuleRulesModel{}
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
		m.Condition = m.Condition.flattenEndpointZtnaRuleRulesCondition(ctx, v, diags)
	}

	if v, ok := o["content"]; ok {
		m.Content = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointZtnaRuleModel) flattenEndpointZtnaRuleRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointZtnaRuleRulesModel {
	if o == nil {
		return []resourceEndpointZtnaRuleRulesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument rules is not type of []interface{}.", "")
		return []resourceEndpointZtnaRuleRulesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointZtnaRuleRulesModel{}
	}

	values := make([]resourceEndpointZtnaRuleRulesModel, len(l))
	for i, ele := range l {
		var m resourceEndpointZtnaRuleRulesModel
		if i < len(s.Rules) {
			m = s.Rules[i]
		}
		values[i] = *m.flattenEndpointZtnaRuleRules(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointZtnaRuleRulesConditionModel) flattenEndpointZtnaRuleRulesCondition(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaRuleRulesConditionModel {
	if input == nil {
		return &resourceEndpointZtnaRuleRulesConditionModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaRuleRulesConditionModel{}
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

func (m *resourceEndpointZtnaRuleLogicModel) flattenEndpointZtnaRuleLogic(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointZtnaRuleLogicModel {
	if input == nil {
		return &resourceEndpointZtnaRuleLogicModel{}
	}
	if m == nil {
		m = &resourceEndpointZtnaRuleLogicModel{}
	}

	return m
}

func (data *resourceEndpointZtnaRuleTagModel) expandEndpointZtnaRuleTag(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() && !data.Datasource.IsUnknown() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceEndpointZtnaRuleRulesModel) expandEndpointZtnaRuleRules(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		result["id"] = data.Id.ValueFloat64()
	}

	if !data.Os.IsNull() && !data.Os.IsUnknown() {
		result["os"] = data.Os.ValueString()
	}

	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Service.IsNull() && !data.Service.IsUnknown() {
		result["service"] = data.Service.ValueString()
	}

	if !data.Account.IsNull() && !data.Account.IsUnknown() {
		result["account"] = data.Account.ValueString()
	}

	if !data.MatchType.IsNull() && !data.MatchType.IsUnknown() {
		result["matchType"] = data.MatchType.ValueString()
	}

	if !data.Subject.IsNull() && !data.Subject.IsUnknown() {
		result["subject"] = data.Subject.ValueString()
	}

	if !data.Issuer.IsNull() && !data.Issuer.IsUnknown() {
		result["issuer"] = data.Issuer.ValueString()
	}

	if !data.Path.IsNull() && !data.Path.IsUnknown() {
		result["path"] = data.Path.ValueString()
	}

	if !data.Negated.IsNull() && !data.Negated.IsUnknown() {
		result["negated"] = data.Negated.ValueBool()
	}

	if !data.EnableLatestUpdateCheck.IsNull() && !data.EnableLatestUpdateCheck.IsUnknown() {
		result["enableLatestUpdateCheck"] = data.EnableLatestUpdateCheck.ValueBool()
	}

	if !data.CheckUpdatesWithinDays.IsNull() && !data.CheckUpdatesWithinDays.IsUnknown() {
		result["checkUpdatesWithinDays"] = data.CheckUpdatesWithinDays.ValueFloat64()
	}

	if !data.Comparator.IsNull() && !data.Comparator.IsUnknown() {
		result["comparator"] = data.Comparator.ValueString()
	}

	if data.Condition != nil && !isZeroStruct(*data.Condition) {
		result["condition"] = data.Condition.expandEndpointZtnaRuleRulesCondition(ctx, diags)
	}

	if !data.Content.IsNull() && !data.Content.IsUnknown() {
		result["content"] = data.Content.ValueString()
	}

	return result
}

func (s *resourceEndpointZtnaRuleModel) expandEndpointZtnaRuleRulesList(ctx context.Context, l []resourceEndpointZtnaRuleRulesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointZtnaRuleRules(ctx, diags)
	}
	return result
}

func (data *resourceEndpointZtnaRuleRulesConditionModel) expandEndpointZtnaRuleRulesCondition(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Key.IsNull() && !data.Key.IsUnknown() {
		result["key"] = data.Key.ValueString()
	}

	if !data.IsDword.IsNull() && !data.IsDword.IsUnknown() {
		result["isDword"] = data.IsDword.ValueBool()
	}

	if !data.Comparator.IsNull() && !data.Comparator.IsUnknown() {
		result["comparator"] = data.Comparator.ValueString()
	}

	if !data.Value.IsNull() && !data.Value.IsUnknown() {
		result["value"] = data.Value.ValueString()
	}

	return result
}

func (data *resourceEndpointZtnaRuleLogicModel) expandEndpointZtnaRuleLogic(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Windows.IsNull() && !data.Windows.IsUnknown() {
		result["windows"] = data.Windows.ValueString()
	}

	if !data.Macos.IsNull() && !data.Macos.IsUnknown() {
		result["macos"] = data.Macos.ValueString()
	}

	if !data.Linux.IsNull() && !data.Linux.IsUnknown() {
		result["linux"] = data.Linux.ValueString()
	}

	if !data.Ios.IsNull() && !data.Ios.IsUnknown() {
		result["ios"] = data.Ios.ValueString()
	}

	if !data.Android.IsNull() && !data.Android.IsUnknown() {
		result["android"] = data.Android.ValueString()
	}

	return result
}
