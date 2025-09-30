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
var _ datasource.DataSource = &datasourceSecurityProfileGroups{}

func newDatasourceSecurityProfileGroups() datasource.DataSource {
	return &datasourceSecurityProfileGroups{}
}

type datasourceSecurityProfileGroups struct {
	fortiClient *FortiClient
}

// datasourceSecurityProfileGroupsModel describes the datasource data model.
type datasourceSecurityProfileGroupsModel struct {
	PrimaryKey                 types.String                                                    `tfsdk:"primary_key"`
	AntivirusProfile           *datasourceSecurityProfileGroupsAntivirusProfileModel           `tfsdk:"antivirus_profile"`
	WebFilterProfile           *datasourceSecurityProfileGroupsWebFilterProfileModel           `tfsdk:"web_filter_profile"`
	VideoFilterProfile         *datasourceSecurityProfileGroupsVideoFilterProfileModel         `tfsdk:"video_filter_profile"`
	DnsFilterProfile           *datasourceSecurityProfileGroupsDnsFilterProfileModel           `tfsdk:"dns_filter_profile"`
	ApplicationControlProfile  *datasourceSecurityProfileGroupsApplicationControlProfileModel  `tfsdk:"application_control_profile"`
	FileFilterProfile          *datasourceSecurityProfileGroupsFileFilterProfileModel          `tfsdk:"file_filter_profile"`
	DlpFilterProfile           *datasourceSecurityProfileGroupsDlpFilterProfileModel           `tfsdk:"dlp_filter_profile"`
	IntrusionPreventionProfile *datasourceSecurityProfileGroupsIntrusionPreventionProfileModel `tfsdk:"intrusion_prevention_profile"`
	SslSshProfile              *datasourceSecurityProfileGroupsSslSshProfileModel              `tfsdk:"ssl_ssh_profile"`
}

func (r *datasourceSecurityProfileGroups) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_profile_groups"
}

func (r *datasourceSecurityProfileGroups) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 79),
				},
				Required: true,
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

func (r *datasourceSecurityProfileGroups) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *datasourceSecurityProfileGroups) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityProfileGroupsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityProfileGroups(ctx, "read", diags))

	read_output, err := c.ReadSecurityProfileGroups(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityProfileGroups(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityProfileGroupsModel) refreshSecurityProfileGroups(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["antivirusProfile"]; ok {
		m.AntivirusProfile = m.AntivirusProfile.flattenSecurityProfileGroupsAntivirusProfile(ctx, v, &diags)
	}

	if v, ok := o["webFilterProfile"]; ok {
		m.WebFilterProfile = m.WebFilterProfile.flattenSecurityProfileGroupsWebFilterProfile(ctx, v, &diags)
	}

	if v, ok := o["videoFilterProfile"]; ok {
		m.VideoFilterProfile = m.VideoFilterProfile.flattenSecurityProfileGroupsVideoFilterProfile(ctx, v, &diags)
	}

	if v, ok := o["dnsFilterProfile"]; ok {
		m.DnsFilterProfile = m.DnsFilterProfile.flattenSecurityProfileGroupsDnsFilterProfile(ctx, v, &diags)
	}

	if v, ok := o["applicationControlProfile"]; ok {
		m.ApplicationControlProfile = m.ApplicationControlProfile.flattenSecurityProfileGroupsApplicationControlProfile(ctx, v, &diags)
	}

	if v, ok := o["fileFilterProfile"]; ok {
		m.FileFilterProfile = m.FileFilterProfile.flattenSecurityProfileGroupsFileFilterProfile(ctx, v, &diags)
	}

	if v, ok := o["dlpFilterProfile"]; ok {
		m.DlpFilterProfile = m.DlpFilterProfile.flattenSecurityProfileGroupsDlpFilterProfile(ctx, v, &diags)
	}

	if v, ok := o["intrusionPreventionProfile"]; ok {
		m.IntrusionPreventionProfile = m.IntrusionPreventionProfile.flattenSecurityProfileGroupsIntrusionPreventionProfile(ctx, v, &diags)
	}

	if v, ok := o["sslSshProfile"]; ok {
		m.SslSshProfile = m.SslSshProfile.flattenSecurityProfileGroupsSslSshProfile(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityProfileGroupsModel) getURLObjectSecurityProfileGroups(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityProfileGroupsAntivirusProfileModel struct {
	Status  types.String                                                 `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupsAntivirusProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupsAntivirusProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupsWebFilterProfileModel struct {
	Status  types.String                                                 `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupsWebFilterProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupsWebFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupsVideoFilterProfileModel struct {
	Status  types.String                                                   `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupsVideoFilterProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupsVideoFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupsDnsFilterProfileModel struct {
	Status  types.String                                                 `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupsDnsFilterProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupsDnsFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupsApplicationControlProfileModel struct {
	Status  types.String                                                          `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupsApplicationControlProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupsApplicationControlProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupsFileFilterProfileModel struct {
	Status  types.String                                                  `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupsFileFilterProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupsFileFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupsDlpFilterProfileModel struct {
	Status  types.String                                                 `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupsDlpFilterProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupsDlpFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupsIntrusionPreventionProfileModel struct {
	Status  types.String                                                           `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupsIntrusionPreventionProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupsIntrusionPreventionProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type datasourceSecurityProfileGroupsSslSshProfileModel struct {
	Status  types.String                                              `tfsdk:"status"`
	Profile *datasourceSecurityProfileGroupsSslSshProfileProfileModel `tfsdk:"profile"`
}

type datasourceSecurityProfileGroupsSslSshProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *datasourceSecurityProfileGroupsAntivirusProfileModel) flattenSecurityProfileGroupsAntivirusProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsAntivirusProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsAntivirusProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsAntivirusProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupsAntivirusProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupsAntivirusProfileProfileModel) flattenSecurityProfileGroupsAntivirusProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsAntivirusProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsAntivirusProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsAntivirusProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupsWebFilterProfileModel) flattenSecurityProfileGroupsWebFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsWebFilterProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsWebFilterProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsWebFilterProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupsWebFilterProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupsWebFilterProfileProfileModel) flattenSecurityProfileGroupsWebFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsWebFilterProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsWebFilterProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsWebFilterProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupsVideoFilterProfileModel) flattenSecurityProfileGroupsVideoFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsVideoFilterProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsVideoFilterProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsVideoFilterProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupsVideoFilterProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupsVideoFilterProfileProfileModel) flattenSecurityProfileGroupsVideoFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsVideoFilterProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsVideoFilterProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsVideoFilterProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupsDnsFilterProfileModel) flattenSecurityProfileGroupsDnsFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsDnsFilterProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsDnsFilterProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsDnsFilterProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupsDnsFilterProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupsDnsFilterProfileProfileModel) flattenSecurityProfileGroupsDnsFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsDnsFilterProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsDnsFilterProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsDnsFilterProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupsApplicationControlProfileModel) flattenSecurityProfileGroupsApplicationControlProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsApplicationControlProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsApplicationControlProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsApplicationControlProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupsApplicationControlProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupsApplicationControlProfileProfileModel) flattenSecurityProfileGroupsApplicationControlProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsApplicationControlProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsApplicationControlProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsApplicationControlProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupsFileFilterProfileModel) flattenSecurityProfileGroupsFileFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsFileFilterProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsFileFilterProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsFileFilterProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupsFileFilterProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupsFileFilterProfileProfileModel) flattenSecurityProfileGroupsFileFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsFileFilterProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsFileFilterProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsFileFilterProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupsDlpFilterProfileModel) flattenSecurityProfileGroupsDlpFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsDlpFilterProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsDlpFilterProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsDlpFilterProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupsDlpFilterProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupsDlpFilterProfileProfileModel) flattenSecurityProfileGroupsDlpFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsDlpFilterProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsDlpFilterProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsDlpFilterProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupsIntrusionPreventionProfileModel) flattenSecurityProfileGroupsIntrusionPreventionProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsIntrusionPreventionProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsIntrusionPreventionProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsIntrusionPreventionProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupsIntrusionPreventionProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupsIntrusionPreventionProfileProfileModel) flattenSecurityProfileGroupsIntrusionPreventionProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsIntrusionPreventionProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsIntrusionPreventionProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsIntrusionPreventionProfileProfileModel{}
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

func (m *datasourceSecurityProfileGroupsSslSshProfileModel) flattenSecurityProfileGroupsSslSshProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsSslSshProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsSslSshProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsSslSshProfileModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["profile"]; ok {
		m.Profile = m.Profile.flattenSecurityProfileGroupsSslSshProfileProfile(ctx, v, diags)
	}

	return m
}

func (m *datasourceSecurityProfileGroupsSslSshProfileProfileModel) flattenSecurityProfileGroupsSslSshProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityProfileGroupsSslSshProfileProfileModel {
	if input == nil {
		return &datasourceSecurityProfileGroupsSslSshProfileProfileModel{}
	}
	if m == nil {
		m = &datasourceSecurityProfileGroupsSslSshProfileProfileModel{}
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
