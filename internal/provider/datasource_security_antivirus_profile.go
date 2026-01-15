// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceSecurityAntivirusProfile{}

func newDatasourceSecurityAntivirusProfile() datasource.DataSource {
	return &datasourceSecurityAntivirusProfile{}
}

type datasourceSecurityAntivirusProfile struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityAntivirusProfileModel describes the datasource data model.
type datasourceSecurityAntivirusProfileModel struct {
	PrimaryKey types.String                                `tfsdk:"primary_key"`
	Http       types.String                                `tfsdk:"http"`
	Smtp       types.String                                `tfsdk:"smtp"`
	Pop3       types.String                                `tfsdk:"pop3"`
	Imap       types.String                                `tfsdk:"imap"`
	Ftp        types.String                                `tfsdk:"ftp"`
	Cifs       types.String                                `tfsdk:"cifs"`
	Cdr        *datasourceSecurityAntivirusProfileCdrModel `tfsdk:"cdr"`
	Direction  types.String                                `tfsdk:"direction"`
}

func (r *datasourceSecurityAntivirusProfile) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_antivirus_profile"
}

func (r *datasourceSecurityAntivirusProfile) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"http": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"smtp": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"pop3": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"imap": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"ftp": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"cifs": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
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
			"cdr": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"enable": schema.BoolAttribute{
						Computed: true,
						Optional: true,
					},
					"file_types": schema.SetAttribute{
						Validators: []validator.Set{
							setvalidator.ValueStringsAre(
								stringvalidator.OneOf("pdf", "office"),
							),
						},
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"allow_error_transmission": schema.BoolAttribute{
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

func (r *datasourceSecurityAntivirusProfile) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_antivirus_profile"
}

func (r *datasourceSecurityAntivirusProfile) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityAntivirusProfileModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityAntivirusProfile(ctx, "read", diags))

	read_output, err := c.ReadSecurityAntivirusProfile(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityAntivirusProfile(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityAntivirusProfileModel) refreshSecurityAntivirusProfile(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["http"]; ok {
		m.Http = parseStringValue(v)
	}

	if v, ok := o["smtp"]; ok {
		m.Smtp = parseStringValue(v)
	}

	if v, ok := o["pop3"]; ok {
		m.Pop3 = parseStringValue(v)
	}

	if v, ok := o["imap"]; ok {
		m.Imap = parseStringValue(v)
	}

	if v, ok := o["ftp"]; ok {
		m.Ftp = parseStringValue(v)
	}

	if v, ok := o["cifs"]; ok {
		m.Cifs = parseStringValue(v)
	}

	if v, ok := o["cdr"]; ok {
		m.Cdr = m.Cdr.flattenSecurityAntivirusProfileCdr(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityAntivirusProfileModel) getURLObjectSecurityAntivirusProfile(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
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

type datasourceSecurityAntivirusProfileCdrModel struct {
	Enable                 types.Bool `tfsdk:"enable"`
	FileTypes              types.Set  `tfsdk:"file_types"`
	AllowErrorTransmission types.Bool `tfsdk:"allow_error_transmission"`
}

func (m *datasourceSecurityAntivirusProfileCdrModel) flattenSecurityAntivirusProfileCdr(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityAntivirusProfileCdrModel {
	if input == nil {
		return &datasourceSecurityAntivirusProfileCdrModel{}
	}
	if m == nil {
		m = &datasourceSecurityAntivirusProfileCdrModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["enable"]; ok {
		m.Enable = parseBoolValue(v)
	}

	if v, ok := o["fileTypes"]; ok {
		m.FileTypes = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["allowErrorTransmission"]; ok {
		m.AllowErrorTransmission = parseBoolValue(v)
	}

	return m
}
