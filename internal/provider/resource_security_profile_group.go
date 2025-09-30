// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceSecurityProfileGroup{}

func newResourceSecurityProfileGroup() resource.Resource {
	return &resourceSecurityProfileGroup{}
}

type resourceSecurityProfileGroup struct {
	fortiClient *FortiClient
}

// resourceSecurityProfileGroupModel describes the resource data model.
type resourceSecurityProfileGroupModel struct {
	ID                         types.String                                                 `tfsdk:"id"`
	PrimaryKey                 types.String                                                 `tfsdk:"primary_key"`
	AntivirusProfile           *resourceSecurityProfileGroupAntivirusProfileModel           `tfsdk:"antivirus_profile"`
	WebFilterProfile           *resourceSecurityProfileGroupWebFilterProfileModel           `tfsdk:"web_filter_profile"`
	VideoFilterProfile         *resourceSecurityProfileGroupVideoFilterProfileModel         `tfsdk:"video_filter_profile"`
	DnsFilterProfile           *resourceSecurityProfileGroupDnsFilterProfileModel           `tfsdk:"dns_filter_profile"`
	ApplicationControlProfile  *resourceSecurityProfileGroupApplicationControlProfileModel  `tfsdk:"application_control_profile"`
	FileFilterProfile          *resourceSecurityProfileGroupFileFilterProfileModel          `tfsdk:"file_filter_profile"`
	DlpFilterProfile           *resourceSecurityProfileGroupDlpFilterProfileModel           `tfsdk:"dlp_filter_profile"`
	IntrusionPreventionProfile *resourceSecurityProfileGroupIntrusionPreventionProfileModel `tfsdk:"intrusion_prevention_profile"`
	SslSshProfile              *resourceSecurityProfileGroupSslSshProfileModel              `tfsdk:"ssl_ssh_profile"`
	Direction                  types.String                                                 `tfsdk:"direction"`
}

func (r *resourceSecurityProfileGroup) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_profile_group"
}

func (r *resourceSecurityProfileGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 79),
				},
				Required: true,
			},
			"direction": schema.StringAttribute{
				Description: "The direction of the target resource.",
				Validators: []validator.String{
					stringvalidator.OneOf("internal-profiles", "outbound-profiles"),
				},
				Computed: true,
				Optional: true,
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

func (r *resourceSecurityProfileGroup) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceSecurityProfileGroup) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceSecurityProfileGroupModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityProfileGroup(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityProfileGroup(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityProfileGroup(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityProfileGroup(ctx, "read", diags))

	read_output, err := c.ReadSecurityProfileGroup(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityProfileGroup(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityProfileGroup) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityProfileGroupModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityProfileGroupModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityProfileGroup(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityProfileGroup(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateSecurityProfileGroup(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityProfileGroup(ctx, "read", diags))

	read_output, err := c.ReadSecurityProfileGroup(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityProfileGroup(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityProfileGroup) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityProfileGroupModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityProfileGroup(ctx, "delete", diags))

	err := c.DeleteSecurityProfileGroup(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
}

func (r *resourceSecurityProfileGroup) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityProfileGroupModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityProfileGroup(ctx, "read", diags))

	read_output, err := c.ReadSecurityProfileGroup(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshSecurityProfileGroup(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityProfileGroup) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected format: direction/primary_key, got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("direction"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), parts[1])...)
}

func (m *resourceSecurityProfileGroupModel) refreshSecurityProfileGroup(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
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

func (data *resourceSecurityProfileGroupModel) getCreateObjectSecurityProfileGroup(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if data.AntivirusProfile != nil && !isZeroStruct(*data.AntivirusProfile) {
		result["antivirusProfile"] = data.AntivirusProfile.expandSecurityProfileGroupAntivirusProfile(ctx, diags)
	}

	if data.WebFilterProfile != nil && !isZeroStruct(*data.WebFilterProfile) {
		result["webFilterProfile"] = data.WebFilterProfile.expandSecurityProfileGroupWebFilterProfile(ctx, diags)
	}

	if data.VideoFilterProfile != nil && !isZeroStruct(*data.VideoFilterProfile) {
		result["videoFilterProfile"] = data.VideoFilterProfile.expandSecurityProfileGroupVideoFilterProfile(ctx, diags)
	}

	if data.DnsFilterProfile != nil && !isZeroStruct(*data.DnsFilterProfile) {
		result["dnsFilterProfile"] = data.DnsFilterProfile.expandSecurityProfileGroupDnsFilterProfile(ctx, diags)
	}

	if data.ApplicationControlProfile != nil && !isZeroStruct(*data.ApplicationControlProfile) {
		result["applicationControlProfile"] = data.ApplicationControlProfile.expandSecurityProfileGroupApplicationControlProfile(ctx, diags)
	}

	if data.FileFilterProfile != nil && !isZeroStruct(*data.FileFilterProfile) {
		result["fileFilterProfile"] = data.FileFilterProfile.expandSecurityProfileGroupFileFilterProfile(ctx, diags)
	}

	if data.DlpFilterProfile != nil && !isZeroStruct(*data.DlpFilterProfile) {
		result["dlpFilterProfile"] = data.DlpFilterProfile.expandSecurityProfileGroupDlpFilterProfile(ctx, diags)
	}

	if data.IntrusionPreventionProfile != nil && !isZeroStruct(*data.IntrusionPreventionProfile) {
		result["intrusionPreventionProfile"] = data.IntrusionPreventionProfile.expandSecurityProfileGroupIntrusionPreventionProfile(ctx, diags)
	}

	if data.SslSshProfile != nil && !isZeroStruct(*data.SslSshProfile) {
		result["sslSshProfile"] = data.SslSshProfile.expandSecurityProfileGroupSslSshProfile(ctx, diags)
	}

	return &result
}

func (data *resourceSecurityProfileGroupModel) getUpdateObjectSecurityProfileGroup(ctx context.Context, state resourceSecurityProfileGroupModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.Equal(state.PrimaryKey) {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if data.AntivirusProfile != nil && !isSameStruct(data.AntivirusProfile, state.AntivirusProfile) {
		result["antivirusProfile"] = data.AntivirusProfile.expandSecurityProfileGroupAntivirusProfile(ctx, diags)
	}

	if data.WebFilterProfile != nil && !isSameStruct(data.WebFilterProfile, state.WebFilterProfile) {
		result["webFilterProfile"] = data.WebFilterProfile.expandSecurityProfileGroupWebFilterProfile(ctx, diags)
	}

	if data.VideoFilterProfile != nil && !isSameStruct(data.VideoFilterProfile, state.VideoFilterProfile) {
		result["videoFilterProfile"] = data.VideoFilterProfile.expandSecurityProfileGroupVideoFilterProfile(ctx, diags)
	}

	if data.DnsFilterProfile != nil && !isSameStruct(data.DnsFilterProfile, state.DnsFilterProfile) {
		result["dnsFilterProfile"] = data.DnsFilterProfile.expandSecurityProfileGroupDnsFilterProfile(ctx, diags)
	}

	if data.ApplicationControlProfile != nil && !isSameStruct(data.ApplicationControlProfile, state.ApplicationControlProfile) {
		result["applicationControlProfile"] = data.ApplicationControlProfile.expandSecurityProfileGroupApplicationControlProfile(ctx, diags)
	}

	if data.FileFilterProfile != nil && !isSameStruct(data.FileFilterProfile, state.FileFilterProfile) {
		result["fileFilterProfile"] = data.FileFilterProfile.expandSecurityProfileGroupFileFilterProfile(ctx, diags)
	}

	if data.DlpFilterProfile != nil && !isSameStruct(data.DlpFilterProfile, state.DlpFilterProfile) {
		result["dlpFilterProfile"] = data.DlpFilterProfile.expandSecurityProfileGroupDlpFilterProfile(ctx, diags)
	}

	if data.IntrusionPreventionProfile != nil && !isSameStruct(data.IntrusionPreventionProfile, state.IntrusionPreventionProfile) {
		result["intrusionPreventionProfile"] = data.IntrusionPreventionProfile.expandSecurityProfileGroupIntrusionPreventionProfile(ctx, diags)
	}

	if data.SslSshProfile != nil && !isSameStruct(data.SslSshProfile, state.SslSshProfile) {
		result["sslSshProfile"] = data.SslSshProfile.expandSecurityProfileGroupSslSshProfile(ctx, diags)
	}

	return &result
}

func (data *resourceSecurityProfileGroupModel) getURLObjectSecurityProfileGroup(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Direction.IsNull() {
		result["direction"] = data.Direction.ValueString()
	}

	return &result
}

type resourceSecurityProfileGroupAntivirusProfileModel struct {
	Status  types.String                                              `tfsdk:"status"`
	Profile *resourceSecurityProfileGroupAntivirusProfileProfileModel `tfsdk:"profile"`
}

type resourceSecurityProfileGroupAntivirusProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityProfileGroupWebFilterProfileModel struct {
	Status  types.String                                              `tfsdk:"status"`
	Profile *resourceSecurityProfileGroupWebFilterProfileProfileModel `tfsdk:"profile"`
}

type resourceSecurityProfileGroupWebFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityProfileGroupVideoFilterProfileModel struct {
	Status  types.String                                                `tfsdk:"status"`
	Profile *resourceSecurityProfileGroupVideoFilterProfileProfileModel `tfsdk:"profile"`
}

type resourceSecurityProfileGroupVideoFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityProfileGroupDnsFilterProfileModel struct {
	Status  types.String                                              `tfsdk:"status"`
	Profile *resourceSecurityProfileGroupDnsFilterProfileProfileModel `tfsdk:"profile"`
}

type resourceSecurityProfileGroupDnsFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityProfileGroupApplicationControlProfileModel struct {
	Status  types.String                                                       `tfsdk:"status"`
	Profile *resourceSecurityProfileGroupApplicationControlProfileProfileModel `tfsdk:"profile"`
}

type resourceSecurityProfileGroupApplicationControlProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityProfileGroupFileFilterProfileModel struct {
	Status  types.String                                               `tfsdk:"status"`
	Profile *resourceSecurityProfileGroupFileFilterProfileProfileModel `tfsdk:"profile"`
}

type resourceSecurityProfileGroupFileFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityProfileGroupDlpFilterProfileModel struct {
	Status  types.String                                              `tfsdk:"status"`
	Profile *resourceSecurityProfileGroupDlpFilterProfileProfileModel `tfsdk:"profile"`
}

type resourceSecurityProfileGroupDlpFilterProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityProfileGroupIntrusionPreventionProfileModel struct {
	Status  types.String                                                        `tfsdk:"status"`
	Profile *resourceSecurityProfileGroupIntrusionPreventionProfileProfileModel `tfsdk:"profile"`
}

type resourceSecurityProfileGroupIntrusionPreventionProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceSecurityProfileGroupSslSshProfileModel struct {
	Status  types.String                                           `tfsdk:"status"`
	Profile *resourceSecurityProfileGroupSslSshProfileProfileModel `tfsdk:"profile"`
}

type resourceSecurityProfileGroupSslSshProfileProfileModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

func (m *resourceSecurityProfileGroupAntivirusProfileModel) flattenSecurityProfileGroupAntivirusProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupAntivirusProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupAntivirusProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupAntivirusProfileModel{}
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

func (m *resourceSecurityProfileGroupAntivirusProfileProfileModel) flattenSecurityProfileGroupAntivirusProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupAntivirusProfileProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupAntivirusProfileProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupAntivirusProfileProfileModel{}
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

func (m *resourceSecurityProfileGroupWebFilterProfileModel) flattenSecurityProfileGroupWebFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupWebFilterProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupWebFilterProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupWebFilterProfileModel{}
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

func (m *resourceSecurityProfileGroupWebFilterProfileProfileModel) flattenSecurityProfileGroupWebFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupWebFilterProfileProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupWebFilterProfileProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupWebFilterProfileProfileModel{}
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

func (m *resourceSecurityProfileGroupVideoFilterProfileModel) flattenSecurityProfileGroupVideoFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupVideoFilterProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupVideoFilterProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupVideoFilterProfileModel{}
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

func (m *resourceSecurityProfileGroupVideoFilterProfileProfileModel) flattenSecurityProfileGroupVideoFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupVideoFilterProfileProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupVideoFilterProfileProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupVideoFilterProfileProfileModel{}
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

func (m *resourceSecurityProfileGroupDnsFilterProfileModel) flattenSecurityProfileGroupDnsFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupDnsFilterProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupDnsFilterProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupDnsFilterProfileModel{}
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

func (m *resourceSecurityProfileGroupDnsFilterProfileProfileModel) flattenSecurityProfileGroupDnsFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupDnsFilterProfileProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupDnsFilterProfileProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupDnsFilterProfileProfileModel{}
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

func (m *resourceSecurityProfileGroupApplicationControlProfileModel) flattenSecurityProfileGroupApplicationControlProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupApplicationControlProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupApplicationControlProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupApplicationControlProfileModel{}
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

func (m *resourceSecurityProfileGroupApplicationControlProfileProfileModel) flattenSecurityProfileGroupApplicationControlProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupApplicationControlProfileProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupApplicationControlProfileProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupApplicationControlProfileProfileModel{}
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

func (m *resourceSecurityProfileGroupFileFilterProfileModel) flattenSecurityProfileGroupFileFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupFileFilterProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupFileFilterProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupFileFilterProfileModel{}
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

func (m *resourceSecurityProfileGroupFileFilterProfileProfileModel) flattenSecurityProfileGroupFileFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupFileFilterProfileProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupFileFilterProfileProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupFileFilterProfileProfileModel{}
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

func (m *resourceSecurityProfileGroupDlpFilterProfileModel) flattenSecurityProfileGroupDlpFilterProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupDlpFilterProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupDlpFilterProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupDlpFilterProfileModel{}
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

func (m *resourceSecurityProfileGroupDlpFilterProfileProfileModel) flattenSecurityProfileGroupDlpFilterProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupDlpFilterProfileProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupDlpFilterProfileProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupDlpFilterProfileProfileModel{}
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

func (m *resourceSecurityProfileGroupIntrusionPreventionProfileModel) flattenSecurityProfileGroupIntrusionPreventionProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupIntrusionPreventionProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupIntrusionPreventionProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupIntrusionPreventionProfileModel{}
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

func (m *resourceSecurityProfileGroupIntrusionPreventionProfileProfileModel) flattenSecurityProfileGroupIntrusionPreventionProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupIntrusionPreventionProfileProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupIntrusionPreventionProfileProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupIntrusionPreventionProfileProfileModel{}
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

func (m *resourceSecurityProfileGroupSslSshProfileModel) flattenSecurityProfileGroupSslSshProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupSslSshProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupSslSshProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupSslSshProfileModel{}
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

func (m *resourceSecurityProfileGroupSslSshProfileProfileModel) flattenSecurityProfileGroupSslSshProfileProfile(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceSecurityProfileGroupSslSshProfileProfileModel {
	if input == nil {
		return &resourceSecurityProfileGroupSslSshProfileProfileModel{}
	}
	if m == nil {
		m = &resourceSecurityProfileGroupSslSshProfileProfileModel{}
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

func (data *resourceSecurityProfileGroupAntivirusProfileModel) expandSecurityProfileGroupAntivirusProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Profile != nil && !isZeroStruct(*data.Profile) {
		result["profile"] = data.Profile.expandSecurityProfileGroupAntivirusProfileProfile(ctx, diags)
	}

	return result
}

func (data *resourceSecurityProfileGroupAntivirusProfileProfileModel) expandSecurityProfileGroupAntivirusProfileProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityProfileGroupWebFilterProfileModel) expandSecurityProfileGroupWebFilterProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Profile != nil && !isZeroStruct(*data.Profile) {
		result["profile"] = data.Profile.expandSecurityProfileGroupWebFilterProfileProfile(ctx, diags)
	}

	return result
}

func (data *resourceSecurityProfileGroupWebFilterProfileProfileModel) expandSecurityProfileGroupWebFilterProfileProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityProfileGroupVideoFilterProfileModel) expandSecurityProfileGroupVideoFilterProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Profile != nil && !isZeroStruct(*data.Profile) {
		result["profile"] = data.Profile.expandSecurityProfileGroupVideoFilterProfileProfile(ctx, diags)
	}

	return result
}

func (data *resourceSecurityProfileGroupVideoFilterProfileProfileModel) expandSecurityProfileGroupVideoFilterProfileProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityProfileGroupDnsFilterProfileModel) expandSecurityProfileGroupDnsFilterProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Profile != nil && !isZeroStruct(*data.Profile) {
		result["profile"] = data.Profile.expandSecurityProfileGroupDnsFilterProfileProfile(ctx, diags)
	}

	return result
}

func (data *resourceSecurityProfileGroupDnsFilterProfileProfileModel) expandSecurityProfileGroupDnsFilterProfileProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityProfileGroupApplicationControlProfileModel) expandSecurityProfileGroupApplicationControlProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Profile != nil && !isZeroStruct(*data.Profile) {
		result["profile"] = data.Profile.expandSecurityProfileGroupApplicationControlProfileProfile(ctx, diags)
	}

	return result
}

func (data *resourceSecurityProfileGroupApplicationControlProfileProfileModel) expandSecurityProfileGroupApplicationControlProfileProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityProfileGroupFileFilterProfileModel) expandSecurityProfileGroupFileFilterProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Profile != nil && !isZeroStruct(*data.Profile) {
		result["profile"] = data.Profile.expandSecurityProfileGroupFileFilterProfileProfile(ctx, diags)
	}

	return result
}

func (data *resourceSecurityProfileGroupFileFilterProfileProfileModel) expandSecurityProfileGroupFileFilterProfileProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityProfileGroupDlpFilterProfileModel) expandSecurityProfileGroupDlpFilterProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Profile != nil && !isZeroStruct(*data.Profile) {
		result["profile"] = data.Profile.expandSecurityProfileGroupDlpFilterProfileProfile(ctx, diags)
	}

	return result
}

func (data *resourceSecurityProfileGroupDlpFilterProfileProfileModel) expandSecurityProfileGroupDlpFilterProfileProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityProfileGroupIntrusionPreventionProfileModel) expandSecurityProfileGroupIntrusionPreventionProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Profile != nil && !isZeroStruct(*data.Profile) {
		result["profile"] = data.Profile.expandSecurityProfileGroupIntrusionPreventionProfileProfile(ctx, diags)
	}

	return result
}

func (data *resourceSecurityProfileGroupIntrusionPreventionProfileProfileModel) expandSecurityProfileGroupIntrusionPreventionProfileProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceSecurityProfileGroupSslSshProfileModel) expandSecurityProfileGroupSslSshProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if data.Profile != nil && !isZeroStruct(*data.Profile) {
		result["profile"] = data.Profile.expandSecurityProfileGroupSslSshProfileProfile(ctx, diags)
	}

	return result
}

func (data *resourceSecurityProfileGroupSslSshProfileProfileModel) expandSecurityProfileGroupSslSshProfileProfile(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}
