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
var _ datasource.DataSource = &datasourceEndpointSandboxProfiles{}

func newDatasourceEndpointSandboxProfiles() datasource.DataSource {
	return &datasourceEndpointSandboxProfiles{}
}

type datasourceEndpointSandboxProfiles struct {
	fortiClient *FortiClient
}

// datasourceEndpointSandboxProfilesModel describes the datasource data model.
type datasourceEndpointSandboxProfilesModel struct {
	SandboxMode                   types.String                                                 `tfsdk:"sandbox_mode"`
	NotificationType              types.String                                                 `tfsdk:"notification_type"`
	TimeoutAwaitingSandboxResults types.Float64                                                `tfsdk:"timeout_awaiting_sandbox_results"`
	FileSubmissionOptions         *datasourceEndpointSandboxProfilesFileSubmissionOptionsModel `tfsdk:"file_submission_options"`
	DetectionVerdictLevel         types.String                                                 `tfsdk:"detection_verdict_level"`
	Exceptions                    *datasourceEndpointSandboxProfilesExceptionsModel            `tfsdk:"exceptions"`
	RemediationActions            types.String                                                 `tfsdk:"remediation_actions"`
	HostName                      types.String                                                 `tfsdk:"host_name"`
	Authentication                types.Bool                                                   `tfsdk:"authentication"`
	Username                      types.String                                                 `tfsdk:"username"`
	Password                      types.String                                                 `tfsdk:"password"`
	PrimaryKey                    types.String                                                 `tfsdk:"primary_key"`
}

func (r *datasourceEndpointSandboxProfiles) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_sandbox_profiles"
}

func (r *datasourceEndpointSandboxProfiles) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"sandbox_mode": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("Disabled", "FortiSASE", "StandaloneFortiSandbox"),
				},
				Computed: true,
				Optional: true,
			},
			"notification_type": schema.StringAttribute{
				Description: "Integer representing how notifications should be handled on FortiSandbox file submission. 0 - display notification balloon when malware is detected in a submission. 1 - display a popup for all file submissions.",
				Validators: []validator.String{
					stringvalidator.OneOf("0", "1"),
				},
				Computed: true,
				Optional: true,
			},
			"timeout_awaiting_sandbox_results": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.AtMost(2147483647),
				},
				Computed: true,
				Optional: true,
			},
			"detection_verdict_level": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("Clean", "Malicious", "High", "Medium", "Low"),
				},
				Computed: true,
				Optional: true,
			},
			"remediation_actions": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("quarantine", "alert"),
				},
				Computed: true,
				Optional: true,
			},
			"host_name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
				Computed: true,
				Optional: true,
			},
			"authentication": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"username": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
				Computed: true,
				Optional: true,
			},
			"password": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				Description: "The primary key of the object. Can be found in the response from the get request.",
				Required:    true,
			},
			"file_submission_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"all_email_downloads": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"all_files_mapped_network_drives": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"all_files_removable_media": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"all_web_downloads": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"exceptions": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"exclude_files_from_trusted_sources": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
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
	}
}

func (r *datasourceEndpointSandboxProfiles) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceEndpointSandboxProfiles) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointSandboxProfilesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointSandboxProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointSandboxProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointSandboxProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointSandboxProfilesModel) refreshEndpointSandboxProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["sandboxMode"]; ok {
		m.SandboxMode = parseStringValue(v)
	}

	if v, ok := o["notificationType"]; ok {
		m.NotificationType = parseStringValue(v)
	}

	if v, ok := o["timeoutAwaitingSandboxResults"]; ok {
		m.TimeoutAwaitingSandboxResults = parseFloat64Value(v)
	}

	if v, ok := o["fileSubmissionOptions"]; ok {
		m.FileSubmissionOptions = m.FileSubmissionOptions.flattenEndpointSandboxProfilesFileSubmissionOptions(ctx, v, &diags)
	}

	if v, ok := o["detectionVerdictLevel"]; ok {
		m.DetectionVerdictLevel = parseStringValue(v)
	}

	if v, ok := o["exceptions"]; ok {
		m.Exceptions = m.Exceptions.flattenEndpointSandboxProfilesExceptions(ctx, v, &diags)
	}

	if v, ok := o["remediationActions"]; ok {
		m.RemediationActions = parseStringValue(v)
	}

	if v, ok := o["hostName"]; ok {
		m.HostName = parseStringValue(v)
	}

	if v, ok := o["authentication"]; ok {
		m.Authentication = parseBoolValue(v)
	}

	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	if v, ok := o["password"]; ok {
		m.Password = parseStringValue(v)
	}

	return diags
}

func (data *datasourceEndpointSandboxProfilesModel) getURLObjectEndpointSandboxProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointSandboxProfilesFileSubmissionOptionsModel struct {
	AllEmailDownloads           types.String `tfsdk:"all_email_downloads"`
	AllFilesMappedNetworkDrives types.String `tfsdk:"all_files_mapped_network_drives"`
	AllFilesRemovableMedia      types.String `tfsdk:"all_files_removable_media"`
	AllWebDownloads             types.String `tfsdk:"all_web_downloads"`
}

type datasourceEndpointSandboxProfilesExceptionsModel struct {
	ExcludeFilesFromTrustedSources types.String `tfsdk:"exclude_files_from_trusted_sources"`
	Files                          types.Set    `tfsdk:"files"`
	Folders                        types.Set    `tfsdk:"folders"`
}

func (m *datasourceEndpointSandboxProfilesFileSubmissionOptionsModel) flattenEndpointSandboxProfilesFileSubmissionOptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointSandboxProfilesFileSubmissionOptionsModel {
	if input == nil {
		return &datasourceEndpointSandboxProfilesFileSubmissionOptionsModel{}
	}
	if m == nil {
		m = &datasourceEndpointSandboxProfilesFileSubmissionOptionsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["allEmailDownloads"]; ok {
		m.AllEmailDownloads = parseStringValue(v)
	}

	if v, ok := o["allFilesMappedNetworkDrives"]; ok {
		m.AllFilesMappedNetworkDrives = parseStringValue(v)
	}

	if v, ok := o["allFilesRemovableMedia"]; ok {
		m.AllFilesRemovableMedia = parseStringValue(v)
	}

	if v, ok := o["allWebDownloads"]; ok {
		m.AllWebDownloads = parseStringValue(v)
	}

	return m
}

func (m *datasourceEndpointSandboxProfilesExceptionsModel) flattenEndpointSandboxProfilesExceptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointSandboxProfilesExceptionsModel {
	if input == nil {
		return &datasourceEndpointSandboxProfilesExceptionsModel{}
	}
	if m == nil {
		m = &datasourceEndpointSandboxProfilesExceptionsModel{}
	}
	o := input.(map[string]interface{})
	m.Files = types.SetNull(types.StringType)
	m.Folders = types.SetNull(types.StringType)

	if v, ok := o["excludeFilesFromTrustedSources"]; ok {
		m.ExcludeFilesFromTrustedSources = parseStringValue(v)
	}

	if v, ok := o["files"]; ok {
		m.Files = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["folders"]; ok {
		m.Folders = parseSetValue(ctx, v, types.StringType)
	}

	return m
}
