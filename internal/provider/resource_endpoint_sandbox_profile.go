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
var _ resource.Resource = &resourceEndpointSandboxProfile{}
var _ resource.ResourceWithMoveState = &resourceEndpointSandboxProfile{}

func newResourceEndpointSandboxProfile() resource.Resource {
	return &resourceEndpointSandboxProfile{}
}

type resourceEndpointSandboxProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceEndpointSandboxProfileModel describes the resource data model.
type resourceEndpointSandboxProfileModel struct {
	ID                            types.String                                              `tfsdk:"id"`
	SandboxMode                   types.String                                              `tfsdk:"sandbox_mode"`
	NotificationType              types.Float64                                             `tfsdk:"notification_type"`
	TimeoutAwaitingSandboxResults types.Float64                                             `tfsdk:"timeout_awaiting_sandbox_results"`
	FileSubmissionOptions         *resourceEndpointSandboxProfileFileSubmissionOptionsModel `tfsdk:"file_submission_options"`
	DetectionVerdictLevel         types.String                                              `tfsdk:"detection_verdict_level"`
	Exceptions                    *resourceEndpointSandboxProfileExceptionsModel            `tfsdk:"exceptions"`
	RemediationActions            types.String                                              `tfsdk:"remediation_actions"`
	HostName                      types.String                                              `tfsdk:"host_name"`
	Authentication                types.Bool                                                `tfsdk:"authentication"`
	Username                      types.String                                              `tfsdk:"username"`
	Password                      types.String                                              `tfsdk:"password"`
	PrimaryKey                    types.String                                              `tfsdk:"primary_key"`
}

func (r *resourceEndpointSandboxProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_sandbox_profile"
}

func (r *resourceEndpointSandboxProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Sandbox Profile Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"sandbox_mode": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("Disabled", "FortiSASE", "StandaloneFortiSandbox"),
				},
				Computed: true,
				Optional: true,
			},
			"notification_type": schema.Float64Attribute{
				MarkdownDescription: "Integer representing how notifications should be handled on FortiSandbox file submission. 0 - display notification balloon when malware is detected in a submission. 1 - display a popup for all file submissions.",
				Computed:            true,
				Optional:            true,
			},
			"timeout_awaiting_sandbox_results": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validatorwarning.AtMost(2147483647),
				},
				Computed: true,
				Optional: true,
			},
			"detection_verdict_level": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("Clean", "Malicious", "High", "Medium", "Low"),
				},
				Computed: true,
				Optional: true,
			},
			"remediation_actions": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("quarantine", "alert"),
				},
				Computed: true,
				Optional: true,
			},
			"host_name": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 128),
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
					stringvalidatorwarning.LengthBetween(1, 128),
				},
				Computed: true,
				Optional: true,
			},
			"password": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				MarkdownDescription: "The primary key of the object. Can be found in the response from the get request.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"file_submission_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"all_email_downloads": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"all_files_mapped_network_drives": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"all_files_removable_media": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"all_web_downloads": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("enable", "disable"),
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
							stringvalidatorwarning.OneOf("enable", "disable"),
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

func (r *resourceEndpointSandboxProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *resourceEndpointSandboxProfile) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_endpoint_sandbox_profiles" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceEndpointSandboxProfileModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceEndpointSandboxProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	var data resourceEndpointSandboxProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointSandboxProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointSandboxProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateEndpointSandboxProfiles(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectEndpointSandboxProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointSandboxProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSandboxProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSandboxProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("endpoint-profile")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointSandboxProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointSandboxProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointSandboxProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointSandboxProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateEndpointSandboxProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointSandboxProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointSandboxProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointSandboxProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSandboxProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointSandboxProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointSandboxProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointSandboxProfile(ctx, "read", diags))

	read_output, err := c.ReadEndpointSandboxProfiles(&input_model)
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

	diags.Append(data.refreshEndpointSandboxProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointSandboxProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceEndpointSandboxProfileModel) refreshEndpointSandboxProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *resourceEndpointSandboxProfileModel) getCreateObjectEndpointSandboxProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SandboxMode.IsNull() && !data.SandboxMode.IsUnknown() {
		result["sandboxMode"] = data.SandboxMode.ValueString()
	}

	if !data.NotificationType.IsNull() && !data.NotificationType.IsUnknown() {
		result["notificationType"] = data.NotificationType.ValueFloat64()
	}

	if !data.TimeoutAwaitingSandboxResults.IsNull() && !data.TimeoutAwaitingSandboxResults.IsUnknown() {
		result["timeoutAwaitingSandboxResults"] = data.TimeoutAwaitingSandboxResults.ValueFloat64()
	}

	if data.FileSubmissionOptions != nil && !isZeroStruct(*data.FileSubmissionOptions) {
		result["fileSubmissionOptions"] = data.FileSubmissionOptions.expandEndpointSandboxProfileFileSubmissionOptions(ctx, diags)
	}

	if !data.DetectionVerdictLevel.IsNull() && !data.DetectionVerdictLevel.IsUnknown() {
		result["detectionVerdictLevel"] = data.DetectionVerdictLevel.ValueString()
	}

	if data.Exceptions != nil && !isZeroStruct(*data.Exceptions) {
		result["exceptions"] = data.Exceptions.expandEndpointSandboxProfileExceptions(ctx, diags)
	}

	if !data.RemediationActions.IsNull() && !data.RemediationActions.IsUnknown() {
		result["remediationActions"] = data.RemediationActions.ValueString()
	}

	if !data.HostName.IsNull() && !data.HostName.IsUnknown() {
		result["hostName"] = data.HostName.ValueString()
	}

	if !data.Authentication.IsNull() && !data.Authentication.IsUnknown() {
		result["authentication"] = data.Authentication.ValueBool()
	}

	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		result["password"] = data.Password.ValueString()
	}

	return &result
}

func (data *resourceEndpointSandboxProfileModel) getUpdateObjectEndpointSandboxProfile(ctx context.Context, state resourceEndpointSandboxProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SandboxMode.IsNull() && !data.SandboxMode.IsUnknown() {
		result["sandboxMode"] = data.SandboxMode.ValueString()
	}

	if !data.NotificationType.IsNull() && !data.NotificationType.IsUnknown() {
		result["notificationType"] = data.NotificationType.ValueFloat64()
	}

	if !data.TimeoutAwaitingSandboxResults.IsNull() && !data.TimeoutAwaitingSandboxResults.IsUnknown() {
		result["timeoutAwaitingSandboxResults"] = data.TimeoutAwaitingSandboxResults.ValueFloat64()
	}

	if data.FileSubmissionOptions != nil {
		result["fileSubmissionOptions"] = data.FileSubmissionOptions.expandEndpointSandboxProfileFileSubmissionOptions(ctx, diags)
	}

	if !data.DetectionVerdictLevel.IsNull() && !data.DetectionVerdictLevel.IsUnknown() {
		result["detectionVerdictLevel"] = data.DetectionVerdictLevel.ValueString()
	}

	if data.Exceptions != nil {
		result["exceptions"] = data.Exceptions.expandEndpointSandboxProfileExceptions(ctx, diags)
	}

	if !data.RemediationActions.IsNull() && !data.RemediationActions.IsUnknown() {
		result["remediationActions"] = data.RemediationActions.ValueString()
	}

	if !data.HostName.IsNull() && !data.HostName.IsUnknown() {
		result["hostName"] = data.HostName.ValueString()
	}

	if !data.Authentication.IsNull() && !data.Authentication.IsUnknown() {
		result["authentication"] = data.Authentication.ValueBool()
	}

	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		result["password"] = data.Password.ValueString()
	}

	return &result
}

func (data *resourceEndpointSandboxProfileModel) getURLObjectEndpointSandboxProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointSandboxProfileFileSubmissionOptionsModel struct {
	AllEmailDownloads           types.String `tfsdk:"all_email_downloads"`
	AllFilesMappedNetworkDrives types.String `tfsdk:"all_files_mapped_network_drives"`
	AllFilesRemovableMedia      types.String `tfsdk:"all_files_removable_media"`
	AllWebDownloads             types.String `tfsdk:"all_web_downloads"`
}

type resourceEndpointSandboxProfileExceptionsModel struct {
	ExcludeFilesFromTrustedSources types.String `tfsdk:"exclude_files_from_trusted_sources"`
	Files                          types.Set    `tfsdk:"files"`
	Folders                        types.Set    `tfsdk:"folders"`
}

func (m *resourceEndpointSandboxProfileFileSubmissionOptionsModel) flattenEndpointSandboxProfileFileSubmissionOptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointSandboxProfileFileSubmissionOptionsModel {
	if input == nil {
		return &resourceEndpointSandboxProfileFileSubmissionOptionsModel{}
	}
	if m == nil {
		m = &resourceEndpointSandboxProfileFileSubmissionOptionsModel{}
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

func (m *resourceEndpointSandboxProfileExceptionsModel) flattenEndpointSandboxProfileExceptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointSandboxProfileExceptionsModel {
	if input == nil {
		return &resourceEndpointSandboxProfileExceptionsModel{}
	}
	if m == nil {
		m = &resourceEndpointSandboxProfileExceptionsModel{}
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

func (data *resourceEndpointSandboxProfileFileSubmissionOptionsModel) expandEndpointSandboxProfileFileSubmissionOptions(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AllEmailDownloads.IsNull() && !data.AllEmailDownloads.IsUnknown() {
		result["allEmailDownloads"] = data.AllEmailDownloads.ValueString()
	}

	if !data.AllFilesMappedNetworkDrives.IsNull() && !data.AllFilesMappedNetworkDrives.IsUnknown() {
		result["allFilesMappedNetworkDrives"] = data.AllFilesMappedNetworkDrives.ValueString()
	}

	if !data.AllFilesRemovableMedia.IsNull() && !data.AllFilesRemovableMedia.IsUnknown() {
		result["allFilesRemovableMedia"] = data.AllFilesRemovableMedia.ValueString()
	}

	if !data.AllWebDownloads.IsNull() && !data.AllWebDownloads.IsUnknown() {
		result["allWebDownloads"] = data.AllWebDownloads.ValueString()
	}

	return result
}

func (data *resourceEndpointSandboxProfileExceptionsModel) expandEndpointSandboxProfileExceptions(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ExcludeFilesFromTrustedSources.IsNull() && !data.ExcludeFilesFromTrustedSources.IsUnknown() {
		result["excludeFilesFromTrustedSources"] = data.ExcludeFilesFromTrustedSources.ValueString()
	}

	if !data.Files.IsNull() && !data.Files.IsUnknown() {
		result["files"] = expandSetToStringList(data.Files)
	}

	if !data.Folders.IsNull() && !data.Folders.IsUnknown() {
		result["folders"] = expandSetToStringList(data.Folders)
	}

	return result
}
