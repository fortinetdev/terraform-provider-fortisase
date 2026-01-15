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
var _ datasource.DataSource = &datasourceSecuritySslSshProfile{}

func newDatasourceSecuritySslSshProfile() datasource.DataSource {
	return &datasourceSecuritySslSshProfile{}
}

type datasourceSecuritySslSshProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecuritySslSshProfileModel describes the datasource data model.
type datasourceSecuritySslSshProfileModel struct {
	PrimaryKey                          types.String                                                `tfsdk:"primary_key"`
	InspectionMode                      types.String                                                `tfsdk:"inspection_mode"`
	ProfileProtocolOptions              *datasourceSecuritySslSshProfileProfileProtocolOptionsModel `tfsdk:"profile_protocol_options"`
	CaCertificate                       *datasourceSecuritySslSshProfileCaCertificateModel          `tfsdk:"ca_certificate"`
	ExpiredCertificateAction            types.String                                                `tfsdk:"expired_certificate_action"`
	RevokedCertificateAction            types.String                                                `tfsdk:"revoked_certificate_action"`
	TimedOutValidationCertificateAction types.String                                                `tfsdk:"timed_out_validation_certificate_action"`
	ValidationFailedCertificateAction   types.String                                                `tfsdk:"validation_failed_certificate_action"`
	CertProbeFailure                    types.String                                                `tfsdk:"cert_probe_failure"`
	Quic                                types.String                                                `tfsdk:"quic"`
	HostExemptions                      []datasourceSecuritySslSshProfileHostExemptionsModel        `tfsdk:"host_exemptions"`
	UrlCategoryExemptions               []datasourceSecuritySslSshProfileUrlCategoryExemptionsModel `tfsdk:"url_category_exemptions"`
	Direction                           types.String                                                `tfsdk:"direction"`
}

func (r *datasourceSecuritySslSshProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_ssl_ssh_profile"
}

func (r *datasourceSecuritySslSshProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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

func (r *datasourceSecuritySslSshProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecuritySslSshProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecuritySslSshProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecuritySslSshProfile(ctx, "read", diags))

	read_output, err := c.ReadSecuritySslSshProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
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

func (m *datasourceSecuritySslSshProfileModel) refreshSecuritySslSshProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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

func (data *datasourceSecuritySslSshProfileModel) getURLObjectSecuritySslSshProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
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

type datasourceSecuritySslSshProfileProfileProtocolOptionsModel struct {
	UnknownContentEncoding types.String  `tfsdk:"unknown_content_encoding"`
	OversizedAction        types.String  `tfsdk:"oversized_action"`
	CompressedLimit        types.Float64 `tfsdk:"compressed_limit"`
	UncompressedLimit      types.Float64 `tfsdk:"uncompressed_limit"`
}

type datasourceSecuritySslSshProfileCaCertificateModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecuritySslSshProfileHostExemptionsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecuritySslSshProfileUrlCategoryExemptionsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecuritySslSshProfileProfileProtocolOptionsModel) flattenSecuritySslSshProfileProfileProtocolOptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecuritySslSshProfileProfileProtocolOptionsModel {
	if input == nil {
		return &datasourceSecuritySslSshProfileProfileProtocolOptionsModel{}
	}
	if m == nil {
		m = &datasourceSecuritySslSshProfileProfileProtocolOptionsModel{}
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

func (m *datasourceSecuritySslSshProfileCaCertificateModel) flattenSecuritySslSshProfileCaCertificate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecuritySslSshProfileCaCertificateModel {
	if input == nil {
		return &datasourceSecuritySslSshProfileCaCertificateModel{}
	}
	if m == nil {
		m = &datasourceSecuritySslSshProfileCaCertificateModel{}
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

func (m *datasourceSecuritySslSshProfileHostExemptionsModel) flattenSecuritySslSshProfileHostExemptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecuritySslSshProfileHostExemptionsModel {
	if input == nil {
		return &datasourceSecuritySslSshProfileHostExemptionsModel{}
	}
	if m == nil {
		m = &datasourceSecuritySslSshProfileHostExemptionsModel{}
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

func (s *datasourceSecuritySslSshProfileModel) flattenSecuritySslSshProfileHostExemptionsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecuritySslSshProfileHostExemptionsModel {
	if o == nil {
		return []datasourceSecuritySslSshProfileHostExemptionsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument host_exemptions is not type of []interface{}.", "")
		return []datasourceSecuritySslSshProfileHostExemptionsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecuritySslSshProfileHostExemptionsModel{}
	}

	values := make([]datasourceSecuritySslSshProfileHostExemptionsModel, len(l))
	for i, ele := range l {
		var m datasourceSecuritySslSshProfileHostExemptionsModel
		values[i] = *m.flattenSecuritySslSshProfileHostExemptions(ctx, ele, diags)
	}

	return values
}

func (m *datasourceSecuritySslSshProfileUrlCategoryExemptionsModel) flattenSecuritySslSshProfileUrlCategoryExemptions(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecuritySslSshProfileUrlCategoryExemptionsModel {
	if input == nil {
		return &datasourceSecuritySslSshProfileUrlCategoryExemptionsModel{}
	}
	if m == nil {
		m = &datasourceSecuritySslSshProfileUrlCategoryExemptionsModel{}
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

func (s *datasourceSecuritySslSshProfileModel) flattenSecuritySslSshProfileUrlCategoryExemptionsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourceSecuritySslSshProfileUrlCategoryExemptionsModel {
	if o == nil {
		return []datasourceSecuritySslSshProfileUrlCategoryExemptionsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument url_category_exemptions is not type of []interface{}.", "")
		return []datasourceSecuritySslSshProfileUrlCategoryExemptionsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourceSecuritySslSshProfileUrlCategoryExemptionsModel{}
	}

	values := make([]datasourceSecuritySslSshProfileUrlCategoryExemptionsModel, len(l))
	for i, ele := range l {
		var m datasourceSecuritySslSshProfileUrlCategoryExemptionsModel
		values[i] = *m.flattenSecuritySslSshProfileUrlCategoryExemptions(ctx, ele, diags)
	}

	return values
}
