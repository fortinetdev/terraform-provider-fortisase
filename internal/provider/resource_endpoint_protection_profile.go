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
var _ resource.Resource = &resourceEndpointProtectionProfile{}
var _ resource.ResourceWithMoveState = &resourceEndpointProtectionProfile{}

func newResourceEndpointProtectionProfile() resource.Resource {
	return &resourceEndpointProtectionProfile{}
}

type resourceEndpointProtectionProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointProtectionProfileModel describes the resource data model.
type resourceEndpointProtectionProfileModel struct {
	ID                                types.String                                                  `tfsdk:"id"`
	Antivirus                         types.String                                                  `tfsdk:"antivirus"`
	Antiransomware                    types.String                                                  `tfsdk:"antiransomware"`
	EventBasedScanning                types.String                                                  `tfsdk:"event_based_scanning"`
	VulnerabilityScan                 types.String                                                  `tfsdk:"vulnerability_scan"`
	AntivirusScan                     types.String                                                  `tfsdk:"antivirus_scan"`
	AutomaticallyPatchVulnerabilities types.String                                                  `tfsdk:"automatically_patch_vulnerabilities"`
	AutomaticVulnerabilityPatchLevel  types.String                                                  `tfsdk:"automatic_vulnerability_patch_level"`
	NotifyEndpointOfBlocks            types.String                                                  `tfsdk:"notify_endpoint_of_blocks"`
	DefaultAction                     types.String                                                  `tfsdk:"default_action"`
	Rules                             []resourceEndpointProtectionProfileRulesModel                 `tfsdk:"rules"`
	Exclusions                        *resourceEndpointProtectionProfileExclusionsModel             `tfsdk:"exclusions"`
	ShowVulnerabilityPopup            types.String                                                  `tfsdk:"show_vulnerability_popup"`
	ProtectedFoldersPath              types.Set                                                     `tfsdk:"protected_folders_path"`
	ScheduledScan                     *resourceEndpointProtectionProfileScheduledScanModel          `tfsdk:"scheduled_scan"`
	ScheduledAntivirusScan            *resourceEndpointProtectionProfileScheduledAntivirusScanModel `tfsdk:"scheduled_antivirus_scan"`
	PrimaryKey                        types.String                                                  `tfsdk:"primary_key"`
}

func (r *resourceEndpointProtectionProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_protection_profile"
}

func (r *resourceEndpointProtectionProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Protection Profile Resource API V2 for FortiSASE.",
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
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"antiransomware": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"event_based_scanning": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"vulnerability_scan": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"antivirus_scan": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"automatically_patch_vulnerabilities": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"automatic_vulnerability_patch_level": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("low", "medium", "high", "critical"),
				},
				Computed: true,
				Optional: true,
			},
			"notify_endpoint_of_blocks": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"default_action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("allow", "block", "monitor"),
				},
				Computed: true,
				Optional: true,
			},
			"show_vulnerability_popup": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
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
								stringvalidatorwarning.OneOf("allow", "block", "monitor"),
							},
							Computed: true,
							Optional: true,
						},
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("simple", "regex"),
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
								stringvalidatorwarning.OneOf("HID", "WPD", "Bluetooth", "CDROM", "SmartCardReader", "USBDevice", "Camera"),
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

					"antivirus": schema.SingleNestedAttribute{
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
					"antiransomware": schema.SingleNestedAttribute{
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
							stringvalidatorwarning.OneOf("daily", "weekly", "monthly"),
						},
						Computed: true,
						Optional: true,
					},
					"day": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validatorwarning.Between(1, 31),
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
							stringvalidatorwarning.OneOf("full", "quick"),
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
							stringvalidatorwarning.OneOf("daily", "weekly", "monthly"),
						},
						Computed: true,
						Optional: true,
					},
					"day": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validatorwarning.Between(1, 31),
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

func (r *resourceEndpointProtectionProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_protection_profile"
}
func (r *resourceEndpointProtectionProfile) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_endpoint_protection_profiles" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceEndpointProtectionProfileModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceEndpointProtectionProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointProtectionProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointProtectionProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointProtectionProfile(ctx, "create", diags))

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

	mkey := data.PrimaryKey.ValueString()
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointProtectionProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointProtectionProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointProtectionProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointProtectionProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointProtectionProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointProtectionProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointProtectionProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointProtectionProfile(ctx, "update", diags))

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
	read_input_model.URLParams = *(data.getURLObjectEndpointProtectionProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointProtectionProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointProtectionProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointProtectionProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointProtectionProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointProtectionProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointProtectionProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointProtectionProfiles(&input_model)
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

	diags.Append(data.refreshEndpointProtectionProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointProtectionProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceEndpointProtectionProfileModel) refreshEndpointProtectionProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.Rules = m.flattenEndpointProtectionProfileRulesList(ctx, v, &diags)
	}

	if v, ok := o["exclusions"]; ok {
		m.Exclusions = m.Exclusions.flattenEndpointProtectionProfileExclusions(ctx, v, &diags)
	}

	if v, ok := o["showVulnerabilityPopup"]; ok {
		m.ShowVulnerabilityPopup = parseStringValue(v)
	}

	if v, ok := o["protectedFoldersPath"]; ok {
		m.ProtectedFoldersPath = parseSetValue(ctx, v, types.StringType)
	} else {
		m.ProtectedFoldersPath = types.SetNull(types.StringType)
	}

	if v, ok := o["scheduledScan"]; ok {
		m.ScheduledScan = m.ScheduledScan.flattenEndpointProtectionProfileScheduledScan(ctx, v, &diags)
	}

	if v, ok := o["scheduledAntivirusScan"]; ok {
		m.ScheduledAntivirusScan = m.ScheduledAntivirusScan.flattenEndpointProtectionProfileScheduledAntivirusScan(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointProtectionProfileModel) getCreateObjectEndpointProtectionProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Antivirus.IsNull() && !data.Antivirus.IsUnknown() {
		result["antivirus"] = data.Antivirus.ValueString()
	}

	if !data.Antiransomware.IsNull() && !data.Antiransomware.IsUnknown() {
		result["antiransomware"] = data.Antiransomware.ValueString()
	}

	if !data.EventBasedScanning.IsNull() && !data.EventBasedScanning.IsUnknown() {
		result["eventBasedScanning"] = data.EventBasedScanning.ValueString()
	}

	if !data.VulnerabilityScan.IsNull() && !data.VulnerabilityScan.IsUnknown() {
		result["vulnerabilityScan"] = data.VulnerabilityScan.ValueString()
	}

	if !data.AntivirusScan.IsNull() && !data.AntivirusScan.IsUnknown() {
		result["antivirusScan"] = data.AntivirusScan.ValueString()
	}

	if !data.AutomaticallyPatchVulnerabilities.IsNull() && !data.AutomaticallyPatchVulnerabilities.IsUnknown() {
		result["automaticallyPatchVulnerabilities"] = data.AutomaticallyPatchVulnerabilities.ValueString()
	}

	if !data.AutomaticVulnerabilityPatchLevel.IsNull() && !data.AutomaticVulnerabilityPatchLevel.IsUnknown() {
		result["automaticVulnerabilityPatchLevel"] = data.AutomaticVulnerabilityPatchLevel.ValueString()
	}

	if !data.NotifyEndpointOfBlocks.IsNull() && !data.NotifyEndpointOfBlocks.IsUnknown() {
		result["notifyEndpointOfBlocks"] = data.NotifyEndpointOfBlocks.ValueString()
	}

	if !data.DefaultAction.IsNull() && !data.DefaultAction.IsUnknown() {
		result["defaultAction"] = data.DefaultAction.ValueString()
	}

	result["rules"] = data.expandEndpointProtectionProfileRulesList(ctx, data.Rules, diags)

	if data.Exclusions != nil && !isZeroStruct(*data.Exclusions) {
		result["exclusions"] = data.Exclusions.expandEndpointProtectionProfileExclusions(ctx, diags)
	}

	if !data.ShowVulnerabilityPopup.IsNull() && !data.ShowVulnerabilityPopup.IsUnknown() {
		result["showVulnerabilityPopup"] = data.ShowVulnerabilityPopup.ValueString()
	}

	if !data.ProtectedFoldersPath.IsNull() && !data.ProtectedFoldersPath.IsUnknown() {
		result["protectedFoldersPath"] = expandSetToStringList(data.ProtectedFoldersPath)
	}

	if data.ScheduledScan != nil && !isZeroStruct(*data.ScheduledScan) {
		result["scheduledScan"] = data.ScheduledScan.expandEndpointProtectionProfileScheduledScan(ctx, diags)
	}

	if data.ScheduledAntivirusScan != nil && !isZeroStruct(*data.ScheduledAntivirusScan) {
		result["scheduledAntivirusScan"] = data.ScheduledAntivirusScan.expandEndpointProtectionProfileScheduledAntivirusScan(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointProtectionProfileModel) getUpdateObjectEndpointProtectionProfile(ctx context.Context, state resourceEndpointProtectionProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Antivirus.IsNull() && !data.Antivirus.IsUnknown() {
		result["antivirus"] = data.Antivirus.ValueString()
	}

	if !data.Antiransomware.IsNull() && !data.Antiransomware.IsUnknown() {
		result["antiransomware"] = data.Antiransomware.ValueString()
	}

	if !data.EventBasedScanning.IsNull() && !data.EventBasedScanning.IsUnknown() {
		result["eventBasedScanning"] = data.EventBasedScanning.ValueString()
	}

	if !data.VulnerabilityScan.IsNull() && !data.VulnerabilityScan.IsUnknown() {
		result["vulnerabilityScan"] = data.VulnerabilityScan.ValueString()
	}

	if !data.AntivirusScan.IsNull() && !data.AntivirusScan.IsUnknown() {
		result["antivirusScan"] = data.AntivirusScan.ValueString()
	}

	if !data.AutomaticallyPatchVulnerabilities.IsNull() && !data.AutomaticallyPatchVulnerabilities.IsUnknown() {
		result["automaticallyPatchVulnerabilities"] = data.AutomaticallyPatchVulnerabilities.ValueString()
	}

	if !data.AutomaticVulnerabilityPatchLevel.IsNull() && !data.AutomaticVulnerabilityPatchLevel.IsUnknown() {
		result["automaticVulnerabilityPatchLevel"] = data.AutomaticVulnerabilityPatchLevel.ValueString()
	}

	if !data.NotifyEndpointOfBlocks.IsNull() && !data.NotifyEndpointOfBlocks.IsUnknown() {
		result["notifyEndpointOfBlocks"] = data.NotifyEndpointOfBlocks.ValueString()
	}

	if !data.DefaultAction.IsNull() && !data.DefaultAction.IsUnknown() {
		result["defaultAction"] = data.DefaultAction.ValueString()
	}

	if data.Rules != nil {
		result["rules"] = data.expandEndpointProtectionProfileRulesList(ctx, data.Rules, diags)
	}

	if data.Exclusions != nil {
		result["exclusions"] = data.Exclusions.expandEndpointProtectionProfileExclusions(ctx, diags)
	}

	if !data.ShowVulnerabilityPopup.IsNull() && !data.ShowVulnerabilityPopup.IsUnknown() {
		result["showVulnerabilityPopup"] = data.ShowVulnerabilityPopup.ValueString()
	}

	if !data.ProtectedFoldersPath.IsNull() && !data.ProtectedFoldersPath.IsUnknown() {
		result["protectedFoldersPath"] = expandSetToStringList(data.ProtectedFoldersPath)
	}

	if data.ScheduledScan != nil {
		result["scheduledScan"] = data.ScheduledScan.expandEndpointProtectionProfileScheduledScan(ctx, diags)
	}

	if data.ScheduledAntivirusScan != nil {
		result["scheduledAntivirusScan"] = data.ScheduledAntivirusScan.expandEndpointProtectionProfileScheduledAntivirusScan(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointProtectionProfileModel) getURLObjectEndpointProtectionProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointProtectionProfileRulesModel struct {
	Action       types.String `tfsdk:"action"`
	Type         types.String `tfsdk:"type"`
	Description  types.String `tfsdk:"description"`
	Class        types.String `tfsdk:"class"`
	Manufacturer types.String `tfsdk:"manufacturer"`
	VendorId     types.String `tfsdk:"vendor_id"`
	ProductId    types.String `tfsdk:"product_id"`
	Revision     types.String `tfsdk:"revision"`
}

type resourceEndpointProtectionProfileExclusionsModel struct {
	Antivirus      *resourceEndpointProtectionProfileExclusionsAntivirusModel      `tfsdk:"antivirus"`
	Antiransomware *resourceEndpointProtectionProfileExclusionsAntiransomwareModel `tfsdk:"antiransomware"`
}

type resourceEndpointProtectionProfileExclusionsAntivirusModel struct {
	Files   types.Set `tfsdk:"files"`
	Folders types.Set `tfsdk:"folders"`
}

type resourceEndpointProtectionProfileExclusionsAntiransomwareModel struct {
	Files   types.Set `tfsdk:"files"`
	Folders types.Set `tfsdk:"folders"`
}

type resourceEndpointProtectionProfileScheduledScanModel struct {
	Time   types.String  `tfsdk:"time"`
	Repeat types.String  `tfsdk:"repeat"`
	Day    types.Float64 `tfsdk:"day"`
}

type resourceEndpointProtectionProfileScheduledAntivirusScanModel struct {
	ScanType types.String  `tfsdk:"scan_type"`
	Time     types.String  `tfsdk:"time"`
	Repeat   types.String  `tfsdk:"repeat"`
	Day      types.Float64 `tfsdk:"day"`
}

func (m *resourceEndpointProtectionProfileRulesModel) flattenEndpointProtectionProfileRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointProtectionProfileRulesModel {
	if input == nil {
		return &resourceEndpointProtectionProfileRulesModel{}
	}
	if m == nil {
		m = &resourceEndpointProtectionProfileRulesModel{}
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

func (s *resourceEndpointProtectionProfileModel) flattenEndpointProtectionProfileRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointProtectionProfileRulesModel {
	if o == nil {
		return []resourceEndpointProtectionProfileRulesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument rules is not type of []interface{}.", "")
		return []resourceEndpointProtectionProfileRulesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointProtectionProfileRulesModel{}
	}

	values := make([]resourceEndpointProtectionProfileRulesModel, len(l))
	for i, ele := range l {
		var m resourceEndpointProtectionProfileRulesModel
		if i < len(s.Rules) {
			m = s.Rules[i]
		}
		values[i] = *m.flattenEndpointProtectionProfileRules(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointProtectionProfileExclusionsModel) flattenEndpointProtectionProfileExclusions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointProtectionProfileExclusionsModel {
	if input == nil {
		return &resourceEndpointProtectionProfileExclusionsModel{}
	}
	if m == nil {
		m = &resourceEndpointProtectionProfileExclusionsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["antivirus"]; ok {
		m.Antivirus = m.Antivirus.flattenEndpointProtectionProfileExclusionsAntivirus(ctx, v, diags)
	}

	if v, ok := o["antiransomware"]; ok {
		m.Antiransomware = m.Antiransomware.flattenEndpointProtectionProfileExclusionsAntiransomware(ctx, v, diags)
	}

	return m
}

func (m *resourceEndpointProtectionProfileExclusionsAntivirusModel) flattenEndpointProtectionProfileExclusionsAntivirus(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointProtectionProfileExclusionsAntivirusModel {
	if input == nil {
		return &resourceEndpointProtectionProfileExclusionsAntivirusModel{}
	}
	if m == nil {
		m = &resourceEndpointProtectionProfileExclusionsAntivirusModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["files"]; ok {
		m.Files = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Files = types.SetNull(types.StringType)
	}

	if v, ok := o["folders"]; ok {
		m.Folders = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Folders = types.SetNull(types.StringType)
	}

	return m
}

func (m *resourceEndpointProtectionProfileExclusionsAntiransomwareModel) flattenEndpointProtectionProfileExclusionsAntiransomware(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointProtectionProfileExclusionsAntiransomwareModel {
	if input == nil {
		return &resourceEndpointProtectionProfileExclusionsAntiransomwareModel{}
	}
	if m == nil {
		m = &resourceEndpointProtectionProfileExclusionsAntiransomwareModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["files"]; ok {
		m.Files = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Files = types.SetNull(types.StringType)
	}

	if v, ok := o["folders"]; ok {
		m.Folders = parseSetValue(ctx, v, types.StringType)
	} else {
		m.Folders = types.SetNull(types.StringType)
	}

	return m
}

func (m *resourceEndpointProtectionProfileScheduledScanModel) flattenEndpointProtectionProfileScheduledScan(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointProtectionProfileScheduledScanModel {
	if input == nil {
		return &resourceEndpointProtectionProfileScheduledScanModel{}
	}
	if m == nil {
		m = &resourceEndpointProtectionProfileScheduledScanModel{}
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

func (m *resourceEndpointProtectionProfileScheduledAntivirusScanModel) flattenEndpointProtectionProfileScheduledAntivirusScan(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointProtectionProfileScheduledAntivirusScanModel {
	if input == nil {
		return &resourceEndpointProtectionProfileScheduledAntivirusScanModel{}
	}
	if m == nil {
		m = &resourceEndpointProtectionProfileScheduledAntivirusScanModel{}
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

func (data *resourceEndpointProtectionProfileRulesModel) expandEndpointProtectionProfileRules(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		result["action"] = data.Action.ValueString()
	}

	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		result["description"] = data.Description.ValueString()
	}

	if !data.Class.IsNull() && !data.Class.IsUnknown() {
		result["class"] = data.Class.ValueString()
	}

	if !data.Manufacturer.IsNull() && !data.Manufacturer.IsUnknown() {
		result["manufacturer"] = data.Manufacturer.ValueString()
	}

	if !data.VendorId.IsNull() && !data.VendorId.IsUnknown() {
		result["vendorId"] = data.VendorId.ValueString()
	}

	if !data.ProductId.IsNull() && !data.ProductId.IsUnknown() {
		result["productId"] = data.ProductId.ValueString()
	}

	if !data.Revision.IsNull() && !data.Revision.IsUnknown() {
		result["revision"] = data.Revision.ValueString()
	}

	return result
}

func (s *resourceEndpointProtectionProfileModel) expandEndpointProtectionProfileRulesList(ctx context.Context, l []resourceEndpointProtectionProfileRulesModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointProtectionProfileRules(ctx, diags)
	}
	return result
}

func (data *resourceEndpointProtectionProfileExclusionsModel) expandEndpointProtectionProfileExclusions(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if data.Antivirus != nil && !isZeroStruct(*data.Antivirus) {
		result["antivirus"] = data.Antivirus.expandEndpointProtectionProfileExclusionsAntivirus(ctx, diags)
	}

	if data.Antiransomware != nil && !isZeroStruct(*data.Antiransomware) {
		result["antiransomware"] = data.Antiransomware.expandEndpointProtectionProfileExclusionsAntiransomware(ctx, diags)
	}

	return result
}

func (data *resourceEndpointProtectionProfileExclusionsAntivirusModel) expandEndpointProtectionProfileExclusionsAntivirus(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Files.IsNull() && !data.Files.IsUnknown() {
		result["files"] = expandSetToStringList(data.Files)
	}

	if !data.Folders.IsNull() && !data.Folders.IsUnknown() {
		result["folders"] = expandSetToStringList(data.Folders)
	}

	return result
}

func (data *resourceEndpointProtectionProfileExclusionsAntiransomwareModel) expandEndpointProtectionProfileExclusionsAntiransomware(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Files.IsNull() && !data.Files.IsUnknown() {
		result["files"] = expandSetToStringList(data.Files)
	}

	if !data.Folders.IsNull() && !data.Folders.IsUnknown() {
		result["folders"] = expandSetToStringList(data.Folders)
	}

	return result
}

func (data *resourceEndpointProtectionProfileScheduledScanModel) expandEndpointProtectionProfileScheduledScan(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Time.IsNull() && !data.Time.IsUnknown() {
		result["time"] = data.Time.ValueString()
	}

	if !data.Repeat.IsNull() && !data.Repeat.IsUnknown() {
		result["repeat"] = data.Repeat.ValueString()
	}

	if !data.Day.IsNull() && !data.Day.IsUnknown() {
		result["day"] = data.Day.ValueFloat64()
	}

	return result
}

func (data *resourceEndpointProtectionProfileScheduledAntivirusScanModel) expandEndpointProtectionProfileScheduledAntivirusScan(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ScanType.IsNull() && !data.ScanType.IsUnknown() {
		result["scanType"] = data.ScanType.ValueString()
	}

	if !data.Time.IsNull() && !data.Time.IsUnknown() {
		result["time"] = data.Time.ValueString()
	}

	if !data.Repeat.IsNull() && !data.Repeat.IsUnknown() {
		result["repeat"] = data.Repeat.ValueString()
	}

	if !data.Day.IsNull() && !data.Day.IsUnknown() {
		result["day"] = data.Day.ValueFloat64()
	}

	return result
}
