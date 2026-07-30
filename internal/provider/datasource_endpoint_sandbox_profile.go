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
var _ datasource.DataSource = &datasourceEndpointSandboxProfile{}

func newDatasourceEndpointSandboxProfile() datasource.DataSource {
	return &datasourceEndpointSandboxProfile{}
}

type datasourceEndpointSandboxProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointSandboxProfileModel describes the datasource data model.
type datasourceEndpointSandboxProfileModel struct {
	SandboxMode                   types.String                                                `tfsdk:"sandbox_mode"`
	NotificationType              types.Float64                                               `tfsdk:"notification_type"`
	TimeoutAwaitingSandboxResults types.Float64                                               `tfsdk:"timeout_awaiting_sandbox_results"`
	FileSubmissionOptions         *datasourceEndpointSandboxProfileFileSubmissionOptionsModel `tfsdk:"file_submission_options"`
	DetectionVerdictLevel         types.String                                                `tfsdk:"detection_verdict_level"`
	Exceptions                    *datasourceEndpointSandboxProfileExceptionsModel            `tfsdk:"exceptions"`
	RemediationActions            types.String                                                `tfsdk:"remediation_actions"`
	HostName                      types.String                                                `tfsdk:"host_name"`
	Authentication                types.Bool                                                  `tfsdk:"authentication"`
	Username                      types.String                                                `tfsdk:"username"`
	Password                      types.String                                                `tfsdk:"password"`
	PrimaryKey                    types.String                                                `tfsdk:"primary_key"`
}

func (r *datasourceEndpointSandboxProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_sandbox_profile"
}

func (r *datasourceEndpointSandboxProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Sandbox Profile Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"sandbox_mode": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("Disabled", "FortiSASE", "StandaloneFortiSandbox"),
				},
				Computed: true,
			},
			"notification_type": schema.Float64Attribute{
				MarkdownDescription: "Integer representing how notifications should be handled on FortiSandbox file submission. 0 - display notification balloon when malware is detected in a submission. 1 - display a popup for all file submissions.",
				Computed:            true,
			},
			"timeout_awaiting_sandbox_results": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(2147483647),
				},
				Computed: true,
			},
			"detection_verdict_level": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("Clean", "Malicious", "High", "Medium", "Low"),
				},
				Computed: true,
			},
			"remediation_actions": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("quarantine", "alert"),
				},
				Computed: true,
			},
			"host_name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 128),
				},
				Computed: true,
			},
			"authentication": schema.BoolAttribute{
				Computed: true,
			},
			"username": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 128),
				},
				Computed: true,
			},
			"password": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
			},
			"file_submission_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"all_email_downloads": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
					},
					"all_files_mapped_network_drives": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
					},
					"all_files_removable_media": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
					},
					"all_web_downloads": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"exceptions": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"exclude_files_from_trusted_sources": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
					},
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
	}
}

func (r *datasourceEndpointSandboxProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_sandbox_profile"
}

func (r *datasourceEndpointSandboxProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointSandboxProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointSandboxProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointSandboxProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSandboxProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointSandboxProfileModel) refreshEndpointSandboxProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["sandboxMode"]; ok {
		m.SandboxMode = parseStringValue(v)
	}

	if v, ok := o["notificationType"]; ok {
		m.NotificationType = parseFloat64Value(v)
	}

	if v, ok := o["timeoutAwaitingSandboxResults"]; ok {
		m.TimeoutAwaitingSandboxResults = parseFloat64Value(v)
	}

	if v, ok := o["fileSubmissionOptions"]; ok {
		m.FileSubmissionOptions = m.FileSubmissionOptions.flattenEndpointSandboxProfileFileSubmissionOptions(ctx, v, &diags)
	}

	if v, ok := o["detectionVerdictLevel"]; ok {
		m.DetectionVerdictLevel = parseStringValue(v)
	}

	if v, ok := o["exceptions"]; ok {
		m.Exceptions = m.Exceptions.flattenEndpointSandboxProfileExceptions(ctx, v, &diags)
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

func (data *datasourceEndpointSandboxProfileModel) getURLObjectEndpointSandboxProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointSandboxProfileFileSubmissionOptionsModel struct {
	AllEmailDownloads           types.String `tfsdk:"all_email_downloads"`
	AllFilesMappedNetworkDrives types.String `tfsdk:"all_files_mapped_network_drives"`
	AllFilesRemovableMedia      types.String `tfsdk:"all_files_removable_media"`
	AllWebDownloads             types.String `tfsdk:"all_web_downloads"`
}

type datasourceEndpointSandboxProfileExceptionsModel struct {
	ExcludeFilesFromTrustedSources types.String `tfsdk:"exclude_files_from_trusted_sources"`
	Files                          types.Set    `tfsdk:"files"`
	Folders                        types.Set    `tfsdk:"folders"`
}

func (m *datasourceEndpointSandboxProfileFileSubmissionOptionsModel) flattenEndpointSandboxProfileFileSubmissionOptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointSandboxProfileFileSubmissionOptionsModel {
	if input == nil {
		return &datasourceEndpointSandboxProfileFileSubmissionOptionsModel{}
	}
	if m == nil {
		m = &datasourceEndpointSandboxProfileFileSubmissionOptionsModel{}
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

func (m *datasourceEndpointSandboxProfileExceptionsModel) flattenEndpointSandboxProfileExceptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointSandboxProfileExceptionsModel {
	if input == nil {
		return &datasourceEndpointSandboxProfileExceptionsModel{}
	}
	if m == nil {
		m = &datasourceEndpointSandboxProfileExceptionsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["excludeFilesFromTrustedSources"]; ok {
		m.ExcludeFilesFromTrustedSources = parseStringValue(v)
	}

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
