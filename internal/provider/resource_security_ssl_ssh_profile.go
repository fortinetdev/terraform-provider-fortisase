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
var _ resource.Resource = &resourceSecuritySslSshProfile{}

func newResourceSecuritySslSshProfile() resource.Resource {
	return &resourceSecuritySslSshProfile{}
}

type resourceSecuritySslSshProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecuritySslSshProfileModel describes the resource data model.
type resourceSecuritySslSshProfileModel struct {
	ID                                  types.String                                              `tfsdk:"id"`
	PrimaryKey                          types.String                                              `tfsdk:"primary_key"`
	InspectionMode                      types.String                                              `tfsdk:"inspection_mode"`
	ProfileProtocolOptions              *resourceSecuritySslSshProfileProfileProtocolOptionsModel `tfsdk:"profile_protocol_options"`
	CaCertificate                       *resourceSecuritySslSshProfileCaCertificateModel          `tfsdk:"ca_certificate"`
	ExpiredCertificateAction            types.String                                              `tfsdk:"expired_certificate_action"`
	RevokedCertificateAction            types.String                                              `tfsdk:"revoked_certificate_action"`
	TimedOutValidationCertificateAction types.String                                              `tfsdk:"timed_out_validation_certificate_action"`
	ValidationFailedCertificateAction   types.String                                              `tfsdk:"validation_failed_certificate_action"`
	CertProbeFailure                    types.String                                              `tfsdk:"cert_probe_failure"`
	Quic                                types.String                                              `tfsdk:"quic"`
	HostExemptions                      []resourceSecuritySslSshProfileHostExemptionsModel        `tfsdk:"host_exemptions"`
	UrlCategoryExemptions               []resourceSecuritySslSshProfileUrlCategoryExemptionsModel `tfsdk:"url_category_exemptions"`
	Direction                           types.String                                              `tfsdk:"direction"`
}

func (r *resourceSecuritySslSshProfile) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ssl_ssh_profile"
}

func (r *resourceSecuritySslSshProfile) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"inspection_mode": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("certificate-inspection", "no-inspection", "deep-inspection"),
				},
				Computed: true,
				Optional: true,
			},
			"expired_certificate_action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "block"),
				},
				Computed: true,
				Optional: true,
			},
			"revoked_certificate_action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "block"),
				},
				Computed: true,
				Optional: true,
			},
			"timed_out_validation_certificate_action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "block"),
				},
				Computed: true,
				Optional: true,
			},
			"validation_failed_certificate_action": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "block"),
				},
				Computed: true,
				Optional: true,
			},
			"cert_probe_failure": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "block"),
				},
				Computed: true,
				Optional: true,
			},
			"quic": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("inspect", "bypass", "block"),
				},
				Computed: true,
				Optional: true,
			},
			"direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				MarkdownDescription: "The direction of the target resource.\nSupported values: internal-profiles, outbound-profiles.",
				Computed:            true,
				Optional:            true,
			},
			"profile_protocol_options": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"unknown_content_encoding": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("block", "inspect", "bypass"),
						},
						Computed: true,
						Optional: true,
					},
					"oversized_action": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("allow", "block"),
						},
						Computed: true,
						Optional: true,
					},
					"compressed_limit": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.Between(10, 64),
						},
						Computed: true,
						Optional: true,
					},
					"uncompressed_limit": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.Between(10, 64),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"ca_certificate": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("system/certificate/ca-certificates"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"host_exemptions": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("network/hosts", "network/host-groups"),
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"url_category_exemptions": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"datasource": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("security/fortiguard-categories", "security/fortiguard-local-categories"),
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

func (r *resourceSecuritySslSshProfile) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_ssl_ssh_profile"
}

func (r *resourceSecuritySslSshProfile) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecuritySslSshProfile")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecuritySslSshProfileModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectSecuritySslSshProfile(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecuritySslSshProfile(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateSecuritySslSshProfile(&input_model)
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
	read_input_model.URLParams = *(data.getURLObjectSecuritySslSshProfile(ctx, "read", diags))

	read_output, err := c.ReadSecuritySslSshProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecuritySslSshProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecuritySslSshProfile) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecuritySslSshProfile")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecuritySslSshProfileModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecuritySslSshProfileModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecuritySslSshProfile(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecuritySslSshProfile(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecuritySslSshProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecuritySslSshProfile(ctx, "read", diags))

	read_output, err := c.ReadSecuritySslSshProfile(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecuritySslSshProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecuritySslSshProfile) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceSecuritySslSshProfile) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecuritySslSshProfileModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecuritySslSshProfile(ctx, "read", diags))

	read_output, err := c.ReadSecuritySslSshProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecuritySslSshProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecuritySslSshProfile) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecuritySslSshProfileModel) refreshSecuritySslSshProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["inspectionMode"]; ok {
		m.InspectionMode = parseStringValue(v)
	}

	if v, ok := o["profileProtocolOptions"]; ok {
		m.ProfileProtocolOptions = m.ProfileProtocolOptions.flattenSecuritySslSshProfileProfileProtocolOptions(ctx, v, &diags)
	}

	if v, ok := o["caCertificate"]; ok {
		m.CaCertificate = m.CaCertificate.flattenSecuritySslSshProfileCaCertificate(ctx, v, &diags)
	}

	if v, ok := o["expiredCertificateAction"]; ok {
		m.ExpiredCertificateAction = parseStringValue(v)
	}

	if v, ok := o["revokedCertificateAction"]; ok {
		m.RevokedCertificateAction = parseStringValue(v)
	}

	if v, ok := o["timedOutValidationCertificateAction"]; ok {
		m.TimedOutValidationCertificateAction = parseStringValue(v)
	}

	if v, ok := o["validationFailedCertificateAction"]; ok {
		m.ValidationFailedCertificateAction = parseStringValue(v)
	}

	if v, ok := o["certProbeFailure"]; ok {
		m.CertProbeFailure = parseStringValue(v)
	}

	if v, ok := o["quic"]; ok {
		m.Quic = parseStringValue(v)
	}

	if v, ok := o["hostExemptions"]; ok {
		m.HostExemptions = m.flattenSecuritySslSshProfileHostExemptionsList(ctx, v, &diags)
	}

	if v, ok := o["urlCategoryExemptions"]; ok {
		m.UrlCategoryExemptions = m.flattenSecuritySslSshProfileUrlCategoryExemptionsList(ctx, v, &diags)
	}

	return diags
}

func (data *resourceSecuritySslSshProfileModel) getCreateObjectSecuritySslSshProfile(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.InspectionMode.IsNull() {
		result["inspectionMode"] = data.InspectionMode.ValueString()
	}

	if data.ProfileProtocolOptions != nil && !isZeroStruct(*data.ProfileProtocolOptions) {
		result["profileProtocolOptions"] = data.ProfileProtocolOptions.expandSecuritySslSshProfileProfileProtocolOptions(ctx, diags)
	}

	if data.CaCertificate != nil && !isZeroStruct(*data.CaCertificate) {
		result["caCertificate"] = data.CaCertificate.expandSecuritySslSshProfileCaCertificate(ctx, diags)
	}

	if !data.ExpiredCertificateAction.IsNull() {
		result["expiredCertificateAction"] = data.ExpiredCertificateAction.ValueString()
	}

	if !data.RevokedCertificateAction.IsNull() {
		result["revokedCertificateAction"] = data.RevokedCertificateAction.ValueString()
	}

	if !data.TimedOutValidationCertificateAction.IsNull() {
		result["timedOutValidationCertificateAction"] = data.TimedOutValidationCertificateAction.ValueString()
	}

	if !data.ValidationFailedCertificateAction.IsNull() {
		result["validationFailedCertificateAction"] = data.ValidationFailedCertificateAction.ValueString()
	}

	if !data.CertProbeFailure.IsNull() {
		result["certProbeFailure"] = data.CertProbeFailure.ValueString()
	}

	if !data.Quic.IsNull() {
		result["quic"] = data.Quic.ValueString()
	}

	if data.HostExemptions != nil {
		result["hostExemptions"] = data.expandSecuritySslSshProfileHostExemptionsList(ctx, data.HostExemptions, diags)
	}

	if data.UrlCategoryExemptions != nil {
		result["urlCategoryExemptions"] = data.expandSecuritySslSshProfileUrlCategoryExemptionsList(ctx, data.UrlCategoryExemptions, diags)
	}

	return &result
}

func (data *resourceSecuritySslSshProfileModel) getUpdateObjectSecuritySslSshProfile(ctx context.Context, state resourceSecuritySslSshProfileModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.InspectionMode.IsNull() {
		result["inspectionMode"] = data.InspectionMode.ValueString()
	}

	if data.ProfileProtocolOptions != nil {
		result["profileProtocolOptions"] = data.ProfileProtocolOptions.expandSecuritySslSshProfileProfileProtocolOptions(ctx, diags)
	}

	if data.CaCertificate != nil {
		result["caCertificate"] = data.CaCertificate.expandSecuritySslSshProfileCaCertificate(ctx, diags)
	}

	if !data.ExpiredCertificateAction.IsNull() {
		result["expiredCertificateAction"] = data.ExpiredCertificateAction.ValueString()
	}

	if !data.RevokedCertificateAction.IsNull() {
		result["revokedCertificateAction"] = data.RevokedCertificateAction.ValueString()
	}

	if !data.TimedOutValidationCertificateAction.IsNull() {
		result["timedOutValidationCertificateAction"] = data.TimedOutValidationCertificateAction.ValueString()
	}

	if !data.ValidationFailedCertificateAction.IsNull() {
		result["validationFailedCertificateAction"] = data.ValidationFailedCertificateAction.ValueString()
	}

	if !data.CertProbeFailure.IsNull() {
		result["certProbeFailure"] = data.CertProbeFailure.ValueString()
	}

	if !data.Quic.IsNull() {
		result["quic"] = data.Quic.ValueString()
	}

	if data.HostExemptions != nil {
		result["hostExemptions"] = data.expandSecuritySslSshProfileHostExemptionsList(ctx, data.HostExemptions, diags)
	}

	if data.UrlCategoryExemptions != nil {
		result["urlCategoryExemptions"] = data.expandSecuritySslSshProfileUrlCategoryExemptionsList(ctx, data.UrlCategoryExemptions, diags)
	}

	return &result
}

func (data *resourceSecuritySslSshProfileModel) getURLObjectSecuritySslSshProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Direction.IsNull() {
		diags.AddWarning("\"direction\" is deprecated and may be removed in future.",
			"It is recommended to recreate the resource without \"direction\" to avoid unexpected behavior in future.",
		)
		result["direction"] = data.Direction.ValueString()
	}

	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceSecuritySslSshProfileProfileProtocolOptionsModel struct {
	UnknownContentEncoding types.String  `tfsdk:"unknown_content_encoding"`
	OversizedAction        types.String  `tfsdk:"oversized_action"`
	CompressedLimit        types.Float64 `tfsdk:"compressed_limit"`
	UncompressedLimit      types.Float64 `tfsdk:"uncompressed_limit"`
}

type resourceSecuritySslSshProfileCaCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecuritySslSshProfileHostExemptionsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecuritySslSshProfileUrlCategoryExemptionsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecuritySslSshProfileProfileProtocolOptionsModel) flattenSecuritySslSshProfileProfileProtocolOptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecuritySslSshProfileProfileProtocolOptionsModel {
	if input == nil {
		return &resourceSecuritySslSshProfileProfileProtocolOptionsModel{}
	}
	if m == nil {
		m = &resourceSecuritySslSshProfileProfileProtocolOptionsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["unknownContentEncoding"]; ok {
		m.UnknownContentEncoding = parseStringValue(v)
	}

	if v, ok := o["oversizedAction"]; ok {
		m.OversizedAction = parseStringValue(v)
	}

	if v, ok := o["compressedLimit"]; ok {
		m.CompressedLimit = parseFloat64Value(v)
	}

	if v, ok := o["uncompressedLimit"]; ok {
		m.UncompressedLimit = parseFloat64Value(v)
	}

	return m
}

func (m *resourceSecuritySslSshProfileCaCertificateModel) flattenSecuritySslSshProfileCaCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecuritySslSshProfileCaCertificateModel {
	if input == nil {
		return &resourceSecuritySslSshProfileCaCertificateModel{}
	}
	if m == nil {
		m = &resourceSecuritySslSshProfileCaCertificateModel{}
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

func (m *resourceSecuritySslSshProfileHostExemptionsModel) flattenSecuritySslSshProfileHostExemptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecuritySslSshProfileHostExemptionsModel {
	if input == nil {
		return &resourceSecuritySslSshProfileHostExemptionsModel{}
	}
	if m == nil {
		m = &resourceSecuritySslSshProfileHostExemptionsModel{}
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

func (s *resourceSecuritySslSshProfileModel) flattenSecuritySslSshProfileHostExemptionsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecuritySslSshProfileHostExemptionsModel {
	if o == nil {
		return []resourceSecuritySslSshProfileHostExemptionsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument host_exemptions is not type of []interface{}.", "")
		return []resourceSecuritySslSshProfileHostExemptionsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecuritySslSshProfileHostExemptionsModel{}
	}

	values := make([]resourceSecuritySslSshProfileHostExemptionsModel, len(l))
	for i, ele := range l {
		var m resourceSecuritySslSshProfileHostExemptionsModel
		values[i] = *m.flattenSecuritySslSshProfileHostExemptions(ctx, ele, diags)
	}

	return values
}

func (m *resourceSecuritySslSshProfileUrlCategoryExemptionsModel) flattenSecuritySslSshProfileUrlCategoryExemptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecuritySslSshProfileUrlCategoryExemptionsModel {
	if input == nil {
		return &resourceSecuritySslSshProfileUrlCategoryExemptionsModel{}
	}
	if m == nil {
		m = &resourceSecuritySslSshProfileUrlCategoryExemptionsModel{}
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

func (s *resourceSecuritySslSshProfileModel) flattenSecuritySslSshProfileUrlCategoryExemptionsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceSecuritySslSshProfileUrlCategoryExemptionsModel {
	if o == nil {
		return []resourceSecuritySslSshProfileUrlCategoryExemptionsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument url_category_exemptions is not type of []interface{}.", "")
		return []resourceSecuritySslSshProfileUrlCategoryExemptionsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceSecuritySslSshProfileUrlCategoryExemptionsModel{}
	}

	values := make([]resourceSecuritySslSshProfileUrlCategoryExemptionsModel, len(l))
	for i, ele := range l {
		var m resourceSecuritySslSshProfileUrlCategoryExemptionsModel
		values[i] = *m.flattenSecuritySslSshProfileUrlCategoryExemptions(ctx, ele, diags)
	}

	return values
}

func (data *resourceSecuritySslSshProfileProfileProtocolOptionsModel) expandSecuritySslSshProfileProfileProtocolOptions(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.UnknownContentEncoding.IsNull() {
		result["unknownContentEncoding"] = data.UnknownContentEncoding.ValueString()
	}

	if !data.OversizedAction.IsNull() {
		result["oversizedAction"] = data.OversizedAction.ValueString()
	}

	if !data.CompressedLimit.IsNull() {
		result["compressedLimit"] = data.CompressedLimit.ValueFloat64()
	}

	if !data.UncompressedLimit.IsNull() {
		result["uncompressedLimit"] = data.UncompressedLimit.ValueFloat64()
	}

	return result
}

func (data *resourceSecuritySslSshProfileCaCertificateModel) expandSecuritySslSshProfileCaCertificate(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecuritySslSshProfileHostExemptionsModel) expandSecuritySslSshProfileHostExemptions(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecuritySslSshProfileModel) expandSecuritySslSshProfileHostExemptionsList(ctx context.Context, l []resourceSecuritySslSshProfileHostExemptionsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecuritySslSshProfileHostExemptions(ctx, diags)
	}
	return result
}

func (data *resourceSecuritySslSshProfileUrlCategoryExemptionsModel) expandSecuritySslSshProfileUrlCategoryExemptions(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceSecuritySslSshProfileModel) expandSecuritySslSshProfileUrlCategoryExemptionsList(ctx context.Context, l []resourceSecuritySslSshProfileUrlCategoryExemptionsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandSecuritySslSshProfileUrlCategoryExemptions(ctx, diags)
	}
	return result
}
