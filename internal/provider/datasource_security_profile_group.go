// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityProfileGroup{}

func newDatasourceSecurityProfileGroup() datasource.DataSource {
	return &datasourceSecurityProfileGroup{}
}

type datasourceSecurityProfileGroup struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityProfileGroupModel describes the datasource data model.
type datasourceSecurityProfileGroupModel struct {
	PrimaryKey                 types.String                                                   `tfsdk:"primary_key"`
	AntivirusProfile           *datasourceSecurityProfileGroupAntivirusProfileModel           `tfsdk:"antivirus_profile"`
	WebFilterProfile           *datasourceSecurityProfileGroupWebFilterProfileModel           `tfsdk:"web_filter_profile"`
	VideoFilterProfile         *datasourceSecurityProfileGroupVideoFilterProfileModel         `tfsdk:"video_filter_profile"`
	DnsFilterProfile           *datasourceSecurityProfileGroupDnsFilterProfileModel           `tfsdk:"dns_filter_profile"`
	ApplicationControlProfile  *datasourceSecurityProfileGroupApplicationControlProfileModel  `tfsdk:"application_control_profile"`
	FileFilterProfile          *datasourceSecurityProfileGroupFileFilterProfileModel          `tfsdk:"file_filter_profile"`
	DlpFilterProfile           *datasourceSecurityProfileGroupDlpFilterProfileModel           `tfsdk:"dlp_filter_profile"`
	IntrusionPreventionProfile *datasourceSecurityProfileGroupIntrusionPreventionProfileModel `tfsdk:"intrusion_prevention_profile"`
	SslSshProfile              *datasourceSecurityProfileGroupSslSshProfileModel              `tfsdk:"ssl_ssh_profile"`
	Direction                  types.String                                                   `tfsdk:"direction"`
}

func (r *datasourceSecurityProfileGroup) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_profile_group"
}

func (r *datasourceSecurityProfileGroup) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 79),
				},
				Required: true,
			},
			"direction": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				MarkdownDescription: "The direction of the target resource.\nSupported values: internal-profiles, outbound-profiles.",
				Computed:            true,
				Optional:            true,
			},
			"antivirus_profile": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"profile": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("security/antivirus-profiles"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"web_filter_profile": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"profile": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("security/web-filter-profiles"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"video_filter_profile": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"profile": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("security/video-filter-profiles"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"dns_filter_profile": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"profile": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("security/dns-filter-profiles"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"application_control_profile": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"profile": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("security/application-control-profiles"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"file_filter_profile": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"profile": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("security/file-filter-profiles"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"dlp_filter_profile": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"profile": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("security/dlp-profiles"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"intrusion_prevention_profile": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"profile": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("security/ips-profiles"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"ssl_ssh_profile": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable"),
						},
						Computed: true,
						Optional: true,
					},
					"profile": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"primary_key": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"datasource": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("security/ssl-ssh-profiles"),
								},
								Computed: true,
								Optional: true,
							},
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

func (r *datasourceSecurityProfileGroup) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_profile_group"
}

func (r *datasourceSecurityProfileGroup) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityProfileGroupModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityProfileGroup(ctx, "read", diags))

	read_output, err := c.ReadSecurityProfileGroup(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityProfileGroup(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityProfileGroupModel) refreshSecurityProfileGroup(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["antivirusProfile"]; ok {
		m.AntivirusProfile = m.AntivirusProfile.flattenSecurityProfileGroupAntivirusProfile(ctx, v, &diags)
	}

	if v, ok := o["webFilterProfile"]; ok {
		m.WebFilterProfile = m.WebFilterProfile.flattenSecurityProfileGroupWebFilterProfile(ctx, v, &diags)
	}

	if v, ok := o["videoFilterProfile"]; ok {
		m.VideoFilterProfile = m.VideoFilterProfile.flattenSecurityProfileGroupVideoFilterProfile(ctx, v, &diags)
	}

	if v, ok := o["dnsFilterProfile"]; ok {
		m.DnsFilterProfile = m.DnsFilterProfile.flattenSecurityProfileGroupDnsFilterProfile(ctx, v, &diags)
	}

	if v, ok := o["applicationControlProfile"]; ok {
		m.ApplicationControlProfile = m.ApplicationControlProfile.flattenSecurityProfileGroupApplicationControlProfile(ctx, v, &diags)
	}

	if v, ok := o["fileFilterProfile"]; ok {
		m.FileFilterProfile = m.FileFilterProfile.flattenSecurityProfileGroupFileFilterProfile(ctx, v, &diags)
	}

	if v, ok := o["dlpFilterProfile"]; ok {
		m.DlpFilterProfile = m.DlpFilterProfile.flattenSecurityProfileGroupDlpFilterProfile(ctx, v, &diags)
	}

	if v, ok := o["intrusionPreventionProfile"]; ok {
		m.IntrusionPreventionProfile = m.IntrusionPreventionProfile.flattenSecurityProfileGroupIntrusionPreventionProfile(ctx, v, &diags)
	}

	if v, ok := o["sslSshProfile"]; ok {
		m.SslSshProfile = m.SslSshProfile.flattenSecurityProfileGroupSslSshProfile(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityProfileGroupModel) getURLObjectSecurityProfileGroup(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Direction.IsNull() {
		diags.AddWarning("\"direction\" is deprecated and may be removed in future.",
			"It is recommended to recreate the resource without \"direction\" to avoid unexpected behavior in future.",
		)
		result["direction"] = data.Direction.ValueString()
	}

	return &result
}

type datasourceSecurityProfileGroupAntivirusProfileModel struct {
	Status  types.String                                                `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupAntivirusProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupAntivirusProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupWebFilterProfileModel struct {
	Status  types.String                                                `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupWebFilterProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupWebFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupVideoFilterProfileModel struct {
	Status  types.String                                                  `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupVideoFilterProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupVideoFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupDnsFilterProfileModel struct {
	Status  types.String                                                `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupDnsFilterProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupDnsFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupApplicationControlProfileModel struct {
	Status  types.String                                                         `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupApplicationControlProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupApplicationControlProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupFileFilterProfileModel struct {
	Status  types.String                                                 `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupFileFilterProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupFileFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupDlpFilterProfileModel struct {
	Status  types.String                                                `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupDlpFilterProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupDlpFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupIntrusionPreventionProfileModel struct {
	Status  types.String                                                          `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupIntrusionPreventionProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupIntrusionPreventionProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupSslSshProfileModel struct {
	Status  types.String                                             `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupSslSshProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupSslSshProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityProfileGroupAntivirusProfileModel) flattenSecurityProfileGroupAntivirusProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupAntivirusProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupAntivirusProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupAntivirusProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupAntivirusProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupAntivirusProfileProfileModel) flattenSecurityProfileGroupAntivirusProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupAntivirusProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupAntivirusProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupAntivirusProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupWebFilterProfileModel) flattenSecurityProfileGroupWebFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupWebFilterProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupWebFilterProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupWebFilterProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupWebFilterProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupWebFilterProfileProfileModel) flattenSecurityProfileGroupWebFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupWebFilterProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupWebFilterProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupWebFilterProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupVideoFilterProfileModel) flattenSecurityProfileGroupVideoFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupVideoFilterProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupVideoFilterProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupVideoFilterProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupVideoFilterProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupVideoFilterProfileProfileModel) flattenSecurityProfileGroupVideoFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupVideoFilterProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupVideoFilterProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupVideoFilterProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupDnsFilterProfileModel) flattenSecurityProfileGroupDnsFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupDnsFilterProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupDnsFilterProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupDnsFilterProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupDnsFilterProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupDnsFilterProfileProfileModel) flattenSecurityProfileGroupDnsFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupDnsFilterProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupDnsFilterProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupDnsFilterProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupApplicationControlProfileModel) flattenSecurityProfileGroupApplicationControlProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupApplicationControlProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupApplicationControlProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupApplicationControlProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupApplicationControlProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupApplicationControlProfileProfileModel) flattenSecurityProfileGroupApplicationControlProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupApplicationControlProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupApplicationControlProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupApplicationControlProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupFileFilterProfileModel) flattenSecurityProfileGroupFileFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupFileFilterProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupFileFilterProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupFileFilterProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupFileFilterProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupFileFilterProfileProfileModel) flattenSecurityProfileGroupFileFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupFileFilterProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupFileFilterProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupFileFilterProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupDlpFilterProfileModel) flattenSecurityProfileGroupDlpFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupDlpFilterProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupDlpFilterProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupDlpFilterProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupDlpFilterProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupDlpFilterProfileProfileModel) flattenSecurityProfileGroupDlpFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupDlpFilterProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupDlpFilterProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupDlpFilterProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupIntrusionPreventionProfileModel) flattenSecurityProfileGroupIntrusionPreventionProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupIntrusionPreventionProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupIntrusionPreventionProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupIntrusionPreventionProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupIntrusionPreventionProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupIntrusionPreventionProfileProfileModel) flattenSecurityProfileGroupIntrusionPreventionProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupIntrusionPreventionProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupIntrusionPreventionProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupIntrusionPreventionProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupSslSshProfileModel) flattenSecurityProfileGroupSslSshProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupSslSshProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupSslSshProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupSslSshProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupSslSshProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupSslSshProfileProfileModel) flattenSecurityProfileGroupSslSshProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupSslSshProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupSslSshProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupSslSshProfileProfileModel{}
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
