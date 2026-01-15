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
var _ resource.Resource = &resourceEndpointProtectionProfiles{}

func newResourceEndpointProtectionProfiles() resource.Resource {
	return &resourceEndpointProtectionProfiles{}
}

type resourceEndpointProtectionProfiles struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointProtectionProfilesModel describes the resource data model.
type resourceEndpointProtectionProfilesModel struct {
	ID                                types.String                                                   `tfsdk:"id"`
	Antivirus                         types.String                                                   `tfsdk:"antivirus"`
	Antiransomware                    types.String                                                   `tfsdk:"antiransomware"`
	EventBasedScanning                types.String                                                   `tfsdk:"event_based_scanning"`
	VulnerabilityScan                 types.String                                                   `tfsdk:"vulnerability_scan"`
	AntivirusScan                     types.String                                                   `tfsdk:"antivirus_scan"`
	AutomaticallyPatchVulnerabilities types.String                                                   `tfsdk:"automatically_patch_vulnerabilities"`
	AutomaticVulnerabilityPatchLevel  types.String                                                   `tfsdk:"automatic_vulnerability_patch_level"`
	NotifyEndpointOfBlocks            types.String                                                   `tfsdk:"notify_endpoint_of_blocks"`
	DefaultAction                     types.String                                                   `tfsdk:"default_action"`
	Rules                             []resourceEndpointProtectionProfilesRulesModel                 `tfsdk:"rules"`
	Exclusions                        *resourceEndpointProtectionProfilesExclusionsModel             `tfsdk:"exclusions"`
	ProtectedFoldersPath              types.Set                                                      `tfsdk:"protected_folders_path"`
	ScheduledScan                     *resourceEndpointProtectionProfilesScheduledScanModel          `tfsdk:"scheduled_scan"`
	ScheduledAntivirusScan            *resourceEndpointProtectionProfilesScheduledAntivirusScanModel `tfsdk:"scheduled_antivirus_scan"`
	PrimaryKey                        types.String                                                   `tfsdk:"primary_key"`
}

func (r *resourceEndpointProtectionProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_protection_profiles"
}

func (r *resourceEndpointProtectionProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"antivirus": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"antiransomware": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"event_based_scanning": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"vulnerability_scan": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"antivirus_scan": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"automatically_patch_vulnerabilities": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"automatic_vulnerability_patch_level": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("low", "medium", "high", "critical"),
				},
				Computed: true,
				Optional: true,
			},
			"notify_endpoint_of_blocks": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"default_action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "block", "monitor"),
				},
				Computed: true,
				Optional: true,
			},
			"protected_folders_path": schema.SetAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"rules": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "block", "monitor"),
							},
							Computed: true,
							Optional: true,
						},
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("simple", "regex"),
							},
							Computed: true,
							Optional: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"class": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("HID", "WPD", "Bluetooth", "CDROM", "SmartCardReader", "USBDevice", "Camera"),
							},
							Computed: true,
							Optional: true,
						},
						"manufacturer": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"vendor_id": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"product_id": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"revision": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"exclusions": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"files": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"folders": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
				},
				Computed: true,
				Optional: true,
			},
			"scheduled_scan": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"time": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"repeat": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("daily", "weekly", "monthly"),
						},
						Computed: true,
						Optional: true,
					},
					"day": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.Between(1, 31),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"scheduled_antivirus_scan": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"scan_type": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("full", "quick"),
						},
						Computed: true,
						Optional: true,
					},
					"time": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"repeat": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("daily", "weekly", "monthly"),
						},
						Computed: true,
						Optional: true,
					},
					"day": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.Between(1, 31),
						},
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

func (r *resourceEndpointProtectionProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_protection_profiles"
}

func (r *resourceEndpointProtectionProfiles) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointProtectionProfiles")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointProtectionProfilesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointProtectionProfiles(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointProtectionProfiles(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateEndpointProtectionProfiles(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectEndpointProtectionProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointProtectionProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointProtectionProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointProtectionProfiles) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("EndpointProtectionProfiles")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointProtectionProfilesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointProtectionProfilesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointProtectionProfiles(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointProtectionProfiles(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointProtectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointProtectionProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointProtectionProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointProtectionProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointProtectionProfiles) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointProtectionProfiles) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointProtectionProfilesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointProtectionProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointProtectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointProtectionProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointProtectionProfiles) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointProtectionProfilesModel) refreshEndpointProtectionProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["antivirus"]; ok {
		m.Antivirus = parseStringValue(v)
	}

	if v, ok := o["antiransomware"]; ok {
		m.Antiransomware = parseStringValue(v)
	}

	if v, ok := o["eventBasedScanning"]; ok {
		m.EventBasedScanning = parseStringValue(v)
	}

	if v, ok := o["vulnerabilityScan"]; ok {
		m.VulnerabilityScan = parseStringValue(v)
	}

	if v, ok := o["antivirusScan"]; ok {
		m.AntivirusScan = parseStringValue(v)
	}

	if v, ok := o["automaticallyPatchVulnerabilities"]; ok {
		m.AutomaticallyPatchVulnerabilities = parseStringValue(v)
	}

	if v, ok := o["automaticVulnerabilityPatchLevel"]; ok {
		m.AutomaticVulnerabilityPatchLevel = parseStringValue(v)
	}

	if v, ok := o["notifyEndpointOfBlocks"]; ok {
		m.NotifyEndpointOfBlocks = parseStringValue(v)
	}

	if v, ok := o["defaultAction"]; ok {
		m.DefaultAction = parseStringValue(v)
	}

	if v, ok := o["rules"]; ok {
		m.Rules = m.flattenEndpointProtectionProfilesRulesList(ctx, v, &diags)
	}

	if v, ok := o["exclusions"]; ok {
		m.Exclusions = m.Exclusions.flattenEndpointProtectionProfilesExclusions(ctx, v, &diags)
	}

	if v, ok := o["protectedFoldersPath"]; ok {
		m.ProtectedFoldersPath = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["scheduledScan"]; ok {
		m.ScheduledScan = m.ScheduledScan.flattenEndpointProtectionProfilesScheduledScan(ctx, v, &diags)
	}

	if v, ok := o["scheduledAntivirusScan"]; ok {
		m.ScheduledAntivirusScan = m.ScheduledAntivirusScan.flattenEndpointProtectionProfilesScheduledAntivirusScan(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointProtectionProfilesModel) getCreateObjectEndpointProtectionProfiles(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Antivirus.IsNull() {
		result["antivirus"] = data.Antivirus.ValueString()
	}

	if !data.Antiransomware.IsNull() {
		result["antiransomware"] = data.Antiransomware.ValueString()
	}

	if !data.EventBasedScanning.IsNull() {
		result["eventBasedScanning"] = data.EventBasedScanning.ValueString()
	}

	if !data.VulnerabilityScan.IsNull() {
		result["vulnerabilityScan"] = data.VulnerabilityScan.ValueString()
	}

	if !data.AntivirusScan.IsNull() {
		result["antivirusScan"] = data.AntivirusScan.ValueString()
	}

	if !data.AutomaticallyPatchVulnerabilities.IsNull() {
		result["automaticallyPatchVulnerabilities"] = data.AutomaticallyPatchVulnerabilities.ValueString()
	}

	if !data.AutomaticVulnerabilityPatchLevel.IsNull() {
		result["automaticVulnerabilityPatchLevel"] = data.AutomaticVulnerabilityPatchLevel.ValueString()
	}

	if !data.NotifyEndpointOfBlocks.IsNull() {
		result["notifyEndpointOfBlocks"] = data.NotifyEndpointOfBlocks.ValueString()
	}

	if !data.DefaultAction.IsNull() {
		result["defaultAction"] = data.DefaultAction.ValueString()
	}

	result["rules"] = data.expandEndpointProtectionProfilesRulesList(ctx, data.Rules, diags)

	if data.Exclusions != nil && !isZeroStruct(*data.Exclusions) {
		result["exclusions"] = data.Exclusions.expandEndpointProtectionProfilesExclusions(ctx, diags)
	}

	if !data.ProtectedFoldersPath.IsNull() {
		result["protectedFoldersPath"] = expandSetToStringList(data.ProtectedFoldersPath)
	}

	if data.ScheduledScan != nil && !isZeroStruct(*data.ScheduledScan) {
		result["scheduledScan"] = data.ScheduledScan.expandEndpointProtectionProfilesScheduledScan(ctx, diags)
	}

	if data.ScheduledAntivirusScan != nil && !isZeroStruct(*data.ScheduledAntivirusScan) {
		result["scheduledAntivirusScan"] = data.ScheduledAntivirusScan.expandEndpointProtectionProfilesScheduledAntivirusScan(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointProtectionProfilesModel) getUpdateObjectEndpointProtectionProfiles(ctx context.Context, state resourceEndpointProtectionProfilesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Antivirus.IsNull() {
		result["antivirus"] = data.Antivirus.ValueString()
	}

	if !data.Antiransomware.IsNull() {
		result["antiransomware"] = data.Antiransomware.ValueString()
	}

	if !data.EventBasedScanning.IsNull() {
		result["eventBasedScanning"] = data.EventBasedScanning.ValueString()
	}

	if !data.VulnerabilityScan.IsNull() {
		result["vulnerabilityScan"] = data.VulnerabilityScan.ValueString()
	}

	if !data.AntivirusScan.IsNull() {
		result["antivirusScan"] = data.AntivirusScan.ValueString()
	}

	if !data.AutomaticallyPatchVulnerabilities.IsNull() {
		result["automaticallyPatchVulnerabilities"] = data.AutomaticallyPatchVulnerabilities.ValueString()
	}

	if !data.AutomaticVulnerabilityPatchLevel.IsNull() {
		result["automaticVulnerabilityPatchLevel"] = data.AutomaticVulnerabilityPatchLevel.ValueString()
	}

	if !data.NotifyEndpointOfBlocks.IsNull() {
		result["notifyEndpointOfBlocks"] = data.NotifyEndpointOfBlocks.ValueString()
	}

	if !data.DefaultAction.IsNull() {
		result["defaultAction"] = data.DefaultAction.ValueString()
	}

	if data.Rules != nil {
		result["rules"] = data.expandEndpointProtectionProfilesRulesList(ctx, data.Rules, diags)
	}

	if data.Exclusions != nil {
		result["exclusions"] = data.Exclusions.expandEndpointProtectionProfilesExclusions(ctx, diags)
	}

	if !data.ProtectedFoldersPath.IsNull() {
		result["protectedFoldersPath"] = expandSetToStringList(data.ProtectedFoldersPath)
	}

	if data.ScheduledScan != nil {
		result["scheduledScan"] = data.ScheduledScan.expandEndpointProtectionProfilesScheduledScan(ctx, diags)
	}

	if data.ScheduledAntivirusScan != nil {
		result["scheduledAntivirusScan"] = data.ScheduledAntivirusScan.expandEndpointProtectionProfilesScheduledAntivirusScan(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointProtectionProfilesModel) getURLObjectEndpointProtectionProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointProtectionProfilesRulesModel struct {
	Action       types.String `tfsdk:"action"`
	Type         types.String `tfsdk:"type"`
	Description  types.String `tfsdk:"description"`
	Class        types.String `tfsdk:"class"`
	Manufacturer types.String `tfsdk:"manufacturer"`
	VendorId     types.String `tfsdk:"vendor_id"`
	ProductId    types.String `tfsdk:"product_id"`
	Revision     types.String `tfsdk:"revision"`
}

type resourceEndpointProtectionProfilesExclusionsModel struct {
	Files   types.Set `tfsdk:"files"`
	Folders types.Set `tfsdk:"folders"`
}

type resourceEndpointProtectionProfilesScheduledScanModel struct {
	Time   types.String  `tfsdk:"time"`
	Repeat types.String  `tfsdk:"repeat"`
	Day    types.Float64 `tfsdk:"day"`
}

type resourceEndpointProtectionProfilesScheduledAntivirusScanModel struct {
	ScanType types.String  `tfsdk:"scan_type"`
	Time     types.String  `tfsdk:"time"`
	Repeat   types.String  `tfsdk:"repeat"`
	Day      types.Float64 `tfsdk:"day"`
}

func (m *resourceEndpointProtectionProfilesRulesModel) flattenEndpointProtectionProfilesRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointProtectionProfilesRulesModel {
	if input == nil {
		return &resourceEndpointProtectionProfilesRulesModel{}
	}
	if m == nil {
		m = &resourceEndpointProtectionProfilesRulesModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["description"]; ok {
		m.Description = parseStringValue(v)
	}

	if v, ok := o["class"]; ok {
		m.Class = parseStringValue(v)
	}

	if v, ok := o["manufacturer"]; ok {
		m.Manufacturer = parseStringValue(v)
	}

	if v, ok := o["vendorId"]; ok {
		m.VendorId = parseStringValue(v)
	}

	if v, ok := o["productId"]; ok {
		m.ProductId = parseStringValue(v)
	}

	if v, ok := o["revision"]; ok {
		m.Revision = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointProtectionProfilesModel) flattenEndpointProtectionProfilesRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointProtectionProfilesRulesModel {
	if o == nil {
		return []resourceEndpointProtectionProfilesRulesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument rules is not type of []interface{}.", "")
		return []resourceEndpointProtectionProfilesRulesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointProtectionProfilesRulesModel{}
	}

	values := make([]resourceEndpointProtectionProfilesRulesModel, len(l))
	for i, ele := range l {
		var m resourceEndpointProtectionProfilesRulesModel
		values[i] = *m.flattenEndpointProtectionProfilesRules(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointProtectionProfilesExclusionsModel) flattenEndpointProtectionProfilesExclusions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointProtectionProfilesExclusionsModel {
	if input == nil {
		return &resourceEndpointProtectionProfilesExclusionsModel{}
	}
	if m == nil {
		m = &resourceEndpointProtectionProfilesExclusionsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["files"]; ok {
		m.Files = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["folders"]; ok {
		m.Folders = parseSetValue(ctx, v, types.StringType)
	}

	return m
}

func (m *resourceEndpointProtectionProfilesScheduledScanModel) flattenEndpointProtectionProfilesScheduledScan(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointProtectionProfilesScheduledScanModel {
	if input == nil {
		return &resourceEndpointProtectionProfilesScheduledScanModel{}
	}
	if m == nil {
		m = &resourceEndpointProtectionProfilesScheduledScanModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["time"]; ok {
		m.Time = parseStringValue(v)
	}

	if v, ok := o["repeat"]; ok {
		m.Repeat = parseStringValue(v)
	}

	if v, ok := o["day"]; ok {
		m.Day = parseFloat64Value(v)
	}

	return m
}

func (m *resourceEndpointProtectionProfilesScheduledAntivirusScanModel) flattenEndpointProtectionProfilesScheduledAntivirusScan(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointProtectionProfilesScheduledAntivirusScanModel {
	if input == nil {
		return &resourceEndpointProtectionProfilesScheduledAntivirusScanModel{}
	}
	if m == nil {
		m = &resourceEndpointProtectionProfilesScheduledAntivirusScanModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["scanType"]; ok {
		m.ScanType = parseStringValue(v)
	}

	if v, ok := o["time"]; ok {
		m.Time = parseStringValue(v)
	}

	if v, ok := o["repeat"]; ok {
		m.Repeat = parseStringValue(v)
	}

	if v, ok := o["day"]; ok {
		m.Day = parseFloat64Value(v)
	}

	return m
}

func (data *resourceEndpointProtectionProfilesRulesModel) expandEndpointProtectionProfilesRules(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Type.IsNull() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Description.IsNull() {
		result["description"] = data.Description.ValueString()
	}

	if !data.Class.IsNull() {
		result["class"] = data.Class.ValueString()
	}

	if !data.Manufacturer.IsNull() {
		result["manufacturer"] = data.Manufacturer.ValueString()
	}

	if !data.VendorId.IsNull() {
		result["vendorId"] = data.VendorId.ValueString()
	}

	if !data.ProductId.IsNull() {
		result["productId"] = data.ProductId.ValueString()
	}

	if !data.Revision.IsNull() {
		result["revision"] = data.Revision.ValueString()
	}

	return result
}

func (s *resourceEndpointProtectionProfilesModel) expandEndpointProtectionProfilesRulesList(ctx context.Context, l []resourceEndpointProtectionProfilesRulesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointProtectionProfilesRules(ctx, diags)
	}
	return result
}

func (data *resourceEndpointProtectionProfilesExclusionsModel) expandEndpointProtectionProfilesExclusions(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Files.IsNull() {
		result["files"] = expandSetToStringList(data.Files)
	}

	if !data.Folders.IsNull() {
		result["folders"] = expandSetToStringList(data.Folders)
	}

	return result
}

func (data *resourceEndpointProtectionProfilesScheduledScanModel) expandEndpointProtectionProfilesScheduledScan(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Time.IsNull() {
		result["time"] = data.Time.ValueString()
	}

	if !data.Repeat.IsNull() {
		result["repeat"] = data.Repeat.ValueString()
	}

	if !data.Day.IsNull() {
		result["day"] = data.Day.ValueFloat64()
	}

	return result
}

func (data *resourceEndpointProtectionProfilesScheduledAntivirusScanModel) expandEndpointProtectionProfilesScheduledAntivirusScan(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ScanType.IsNull() {
		result["scanType"] = data.ScanType.ValueString()
	}

	if !data.Time.IsNull() {
		result["time"] = data.Time.ValueString()
	}

	if !data.Repeat.IsNull() {
		result["repeat"] = data.Repeat.ValueString()
	}

	if !data.Day.IsNull() {
		result["day"] = data.Day.ValueFloat64()
	}

	return result
}
