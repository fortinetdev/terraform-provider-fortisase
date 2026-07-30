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
var _ datasource.DataSource = &datasourceEndpointProtectionProfile{}

func newDatasourceEndpointProtectionProfile() datasource.DataSource {
	return &datasourceEndpointProtectionProfile{}
}

type datasourceEndpointProtectionProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointProtectionProfileModel describes the datasource data model.
type datasourceEndpointProtectionProfileModel struct {
	Antivirus                         types.String                                                    `tfsdk:"antivirus"`
	Antiransomware                    types.String                                                    `tfsdk:"antiransomware"`
	EventBasedScanning                types.String                                                    `tfsdk:"event_based_scanning"`
	VulnerabilityScan                 types.String                                                    `tfsdk:"vulnerability_scan"`
	AntivirusScan                     types.String                                                    `tfsdk:"antivirus_scan"`
	AutomaticallyPatchVulnerabilities types.String                                                    `tfsdk:"automatically_patch_vulnerabilities"`
	AutomaticVulnerabilityPatchLevel  types.String                                                    `tfsdk:"automatic_vulnerability_patch_level"`
	NotifyEndpointOfBlocks            types.String                                                    `tfsdk:"notify_endpoint_of_blocks"`
	DefaultAction                     types.String                                                    `tfsdk:"default_action"`
	Rules                             []datasourceEndpointProtectionProfileRulesModel                 `tfsdk:"rules"`
	Exclusions                        *datasourceEndpointProtectionProfileExclusionsModel             `tfsdk:"exclusions"`
	ShowVulnerabilityPopup            types.String                                                    `tfsdk:"show_vulnerability_popup"`
	ProtectedFoldersPath              types.Set                                                       `tfsdk:"protected_folders_path"`
	ScheduledScan                     *datasourceEndpointProtectionProfileScheduledScanModel          `tfsdk:"scheduled_scan"`
	ScheduledAntivirusScan            *datasourceEndpointProtectionProfileScheduledAntivirusScanModel `tfsdk:"scheduled_antivirus_scan"`
	PrimaryKey                        types.String                                                    `tfsdk:"primary_key"`
}

func (r *datasourceEndpointProtectionProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_protection_profile"
}

func (r *datasourceEndpointProtectionProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Protection Profile Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"antivirus": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"antiransomware": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"event_based_scanning": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"vulnerability_scan": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"antivirus_scan": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"automatically_patch_vulnerabilities": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"automatic_vulnerability_patch_level": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("low", "medium", "high", "critical"),
				},
				Computed: true,
			},
			"notify_endpoint_of_blocks": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"default_action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("allow", "block", "monitor"),
				},
				Computed: true,
			},
			"show_vulnerability_popup": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("enable", "disable"),
				},
				Computed: true,
			},
			"protected_folders_path": schema.SetAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
			},
			"rules": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"action": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("allow", "block", "monitor"),
							},
							Computed: true,
						},
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("simple", "regex"),
							},
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"class": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("HID", "WPD", "Bluetooth", "CDROM", "SmartCardReader", "USBDevice", "Camera"),
							},
							Computed: true,
						},
						"manufacturer": schema.StringAttribute{
							Computed: true,
						},
						"vendor_id": schema.StringAttribute{
							Computed: true,
						},
						"product_id": schema.StringAttribute{
							Computed: true,
						},
						"revision": schema.StringAttribute{
							Computed: true,
						},
					},
				},
				Computed: true,
			},
			"exclusions": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{

					"antivirus": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"files": schema.SetAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"folders": schema.SetAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
						Computed: true,
					},
					"antiransomware": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"files": schema.SetAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"folders": schema.SetAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"scheduled_scan": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"time": schema.StringAttribute{
						Computed: true,
					},
					"repeat": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("daily", "weekly", "monthly"),
						},
						Computed: true,
					},
					"day": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validatorwarning.Between(1, 31),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"scheduled_antivirus_scan": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"scan_type": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("full", "quick"),
						},
						Computed: true,
					},
					"time": schema.StringAttribute{
						Computed: true,
					},
					"repeat": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("daily", "weekly", "monthly"),
						},
						Computed: true,
					},
					"day": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validatorwarning.Between(1, 31),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceEndpointProtectionProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceEndpointProtectionProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointProtectionProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointProtectionProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointProtectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
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

func (m *datasourceEndpointProtectionProfileModel) refreshEndpointProtectionProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *datasourceEndpointProtectionProfileModel) getURLObjectEndpointProtectionProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointProtectionProfileRulesModel struct {
	Action       types.String `tfsdk:"action"`
	Type         types.String `tfsdk:"type"`
	Description  types.String `tfsdk:"description"`
	Class        types.String `tfsdk:"class"`
	Manufacturer types.String `tfsdk:"manufacturer"`
	VendorId     types.String `tfsdk:"vendor_id"`
	ProductId    types.String `tfsdk:"product_id"`
	Revision     types.String `tfsdk:"revision"`
}

type datasourceEndpointProtectionProfileExclusionsModel struct {
	Antivirus      *datasourceEndpointProtectionProfileExclusionsAntivirusModel      `tfsdk:"antivirus"`
	Antiransomware *datasourceEndpointProtectionProfileExclusionsAntiransomwareModel `tfsdk:"antiransomware"`
}

type datasourceEndpointProtectionProfileExclusionsAntivirusModel struct {
	Files   types.Set `tfsdk:"files"`
	Folders types.Set `tfsdk:"folders"`
}

type datasourceEndpointProtectionProfileExclusionsAntiransomwareModel struct {
	Files   types.Set `tfsdk:"files"`
	Folders types.Set `tfsdk:"folders"`
}

type datasourceEndpointProtectionProfileScheduledScanModel struct {
	Time   types.String  `tfsdk:"time"`
	Repeat types.String  `tfsdk:"repeat"`
	Day    types.Float64 `tfsdk:"day"`
}

type datasourceEndpointProtectionProfileScheduledAntivirusScanModel struct {
	ScanType types.String  `tfsdk:"scan_type"`
	Time     types.String  `tfsdk:"time"`
	Repeat   types.String  `tfsdk:"repeat"`
	Day      types.Float64 `tfsdk:"day"`
}

func (m *datasourceEndpointProtectionProfileRulesModel) flattenEndpointProtectionProfileRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointProtectionProfileRulesModel {
	if input == nil {
		return &datasourceEndpointProtectionProfileRulesModel{}
	}
	if m == nil {
		m = &datasourceEndpointProtectionProfileRulesModel{}
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

func (s *datasourceEndpointProtectionProfileModel) flattenEndpointProtectionProfileRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointProtectionProfileRulesModel {
	if o == nil {
		return []datasourceEndpointProtectionProfileRulesModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument rules is not type of []interface{}.", "")
		return []datasourceEndpointProtectionProfileRulesModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointProtectionProfileRulesModel{}
	}

	values := make([]datasourceEndpointProtectionProfileRulesModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointProtectionProfileRulesModel
		if i < len(s.Rules) {
			m = s.Rules[i]
		}
		values[i] = *m.flattenEndpointProtectionProfileRules(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointProtectionProfileExclusionsModel) flattenEndpointProtectionProfileExclusions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointProtectionProfileExclusionsModel {
	if input == nil {
		return &datasourceEndpointProtectionProfileExclusionsModel{}
	}
	if m == nil {
		m = &datasourceEndpointProtectionProfileExclusionsModel{}
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

func (m *datasourceEndpointProtectionProfileExclusionsAntivirusModel) flattenEndpointProtectionProfileExclusionsAntivirus(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointProtectionProfileExclusionsAntivirusModel {
	if input == nil {
		return &datasourceEndpointProtectionProfileExclusionsAntivirusModel{}
	}
	if m == nil {
		m = &datasourceEndpointProtectionProfileExclusionsAntivirusModel{}
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

func (m *datasourceEndpointProtectionProfileExclusionsAntiransomwareModel) flattenEndpointProtectionProfileExclusionsAntiransomware(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointProtectionProfileExclusionsAntiransomwareModel {
	if input == nil {
		return &datasourceEndpointProtectionProfileExclusionsAntiransomwareModel{}
	}
	if m == nil {
		m = &datasourceEndpointProtectionProfileExclusionsAntiransomwareModel{}
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

func (m *datasourceEndpointProtectionProfileScheduledScanModel) flattenEndpointProtectionProfileScheduledScan(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointProtectionProfileScheduledScanModel {
	if input == nil {
		return &datasourceEndpointProtectionProfileScheduledScanModel{}
	}
	if m == nil {
		m = &datasourceEndpointProtectionProfileScheduledScanModel{}
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

func (m *datasourceEndpointProtectionProfileScheduledAntivirusScanModel) flattenEndpointProtectionProfileScheduledAntivirusScan(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointProtectionProfileScheduledAntivirusScanModel {
	if input == nil {
		return &datasourceEndpointProtectionProfileScheduledAntivirusScanModel{}
	}
	if m == nil {
		m = &datasourceEndpointProtectionProfileScheduledAntivirusScanModel{}
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
