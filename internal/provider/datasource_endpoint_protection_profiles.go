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
var _ datasource.DataSource = &datasourceEndpointProtectionProfiles{}

func newDatasourceEndpointProtectionProfiles() datasource.DataSource {
	return &datasourceEndpointProtectionProfiles{}
}

type datasourceEndpointProtectionProfiles struct {
	fortiClient *FortiClient
}

// datasourceEndpointProtectionProfilesModel describes the datasource data model.
type datasourceEndpointProtectionProfilesModel struct {
	Antivirus                         types.String                                            `tfsdk:"antivirus"`
	Antiransomware                    types.String                                            `tfsdk:"antiransomware"`
	EventBasedScanning                types.String                                            `tfsdk:"event_based_scanning"`
	VulnerabilityScan                 types.String                                            `tfsdk:"vulnerability_scan"`
	AutomaticallyPatchVulnerabilities types.String                                            `tfsdk:"automatically_patch_vulnerabilities"`
	AutomaticVulnerabilityPatchLevel  types.String                                            `tfsdk:"automatic_vulnerability_patch_level"`
	NotifyEndpointOfBlocks            types.String                                            `tfsdk:"notify_endpoint_of_blocks"`
	DefaultAction                     types.String                                            `tfsdk:"default_action"`
	Rules                             []datasourceEndpointProtectionProfilesRulesModel        `tfsdk:"rules"`
	Exclusions                        *datasourceEndpointProtectionProfilesExclusionsModel    `tfsdk:"exclusions"`
	ProtectedFoldersPath              types.Set                                               `tfsdk:"protected_folders_path"`
	ScheduledScan                     *datasourceEndpointProtectionProfilesScheduledScanModel `tfsdk:"scheduled_scan"`
	PrimaryKey                        types.String                                            `tfsdk:"primary_key"`
}

func (r *datasourceEndpointProtectionProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_protection_profiles"
}

func (r *datasourceEndpointProtectionProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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
				Description: "The primary key of the object. Can be found in the response from the get request.",
				Required:    true,
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
		},
	}
}

func (r *datasourceEndpointProtectionProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceEndpointProtectionProfiles) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointProtectionProfilesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointProtectionProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointProtectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointProtectionProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointProtectionProfilesModel) refreshEndpointProtectionProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

	return diags
}

func (data *datasourceEndpointProtectionProfilesModel) getURLObjectEndpointProtectionProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointProtectionProfilesRulesModel struct {
	Action       types.String `tfsdk:"action"`
	Type         types.String `tfsdk:"type"`
	Description  types.String `tfsdk:"description"`
	Class        types.String `tfsdk:"class"`
	Manufacturer types.String `tfsdk:"manufacturer"`
	VendorId     types.String `tfsdk:"vendor_id"`
	ProductId    types.String `tfsdk:"product_id"`
	Revision     types.String `tfsdk:"revision"`
}

type datasourceEndpointProtectionProfilesExclusionsModel struct {
	Files   types.Set `tfsdk:"files"`
	Folders types.Set `tfsdk:"folders"`
}

type datasourceEndpointProtectionProfilesScheduledScanModel struct {
	Time   types.String  `tfsdk:"time"`
	Repeat types.String  `tfsdk:"repeat"`
	Day    types.Float64 `tfsdk:"day"`
}

func (m *datasourceEndpointProtectionProfilesRulesModel) flattenEndpointProtectionProfilesRules(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointProtectionProfilesRulesModel {
	if input == nil {
		return &datasourceEndpointProtectionProfilesRulesModel{}
	}
	if m == nil {
		m = &datasourceEndpointProtectionProfilesRulesModel{}
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

func (s *datasourceEndpointProtectionProfilesModel) flattenEndpointProtectionProfilesRulesList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceEndpointProtectionProfilesRulesModel {
	if o == nil {
		return []datasourceEndpointProtectionProfilesRulesModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument rules is not type of []interface{}.", "")
		return []datasourceEndpointProtectionProfilesRulesModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceEndpointProtectionProfilesRulesModel{}
	}

	values := make([]datasourceEndpointProtectionProfilesRulesModel, len(l))
	for i, ele := range l {
		var m datasourceEndpointProtectionProfilesRulesModel
		values[i] = *m.flattenEndpointProtectionProfilesRules(ctx, ele, diags)
	}

	return values
}

func (m *datasourceEndpointProtectionProfilesExclusionsModel) flattenEndpointProtectionProfilesExclusions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointProtectionProfilesExclusionsModel {
	if input == nil {
		return &datasourceEndpointProtectionProfilesExclusionsModel{}
	}
	if m == nil {
		m = &datasourceEndpointProtectionProfilesExclusionsModel{}
	}
	o := input.(map[string]interface{})
	m.Files = types.SetNull(types.StringType)
	m.Folders = types.SetNull(types.StringType)

	if v, ok := o["files"]; ok {
		m.Files = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["folders"]; ok {
		m.Folders = parseSetValue(ctx, v, types.StringType)
	}

	return m
}

func (m *datasourceEndpointProtectionProfilesScheduledScanModel) flattenEndpointProtectionProfilesScheduledScan(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointProtectionProfilesScheduledScanModel {
	if input == nil {
		return &datasourceEndpointProtectionProfilesScheduledScanModel{}
	}
	if m == nil {
		m = &datasourceEndpointProtectionProfilesScheduledScanModel{}
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
