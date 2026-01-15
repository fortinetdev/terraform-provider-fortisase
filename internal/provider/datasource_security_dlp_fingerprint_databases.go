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
var _ datasource.DataSource = &datasourceSecurityDlpFingerprintDatabases{}

func newDatasourceSecurityDlpFingerprintDatabases() datasource.DataSource {
	return &datasourceSecurityDlpFingerprintDatabases{}
}

type datasourceSecurityDlpFingerprintDatabases struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityDlpFingerprintDatabasesModel describes the datasource data model.
type datasourceSecurityDlpFingerprintDatabasesModel struct {
	PrimaryKey                    types.String                                                  `tfsdk:"primary_key"`
	Server                        types.String                                                  `tfsdk:"server"`
	Sensitivity                   types.String                                                  `tfsdk:"sensitivity"`
	IncludeSubdirectories         types.String                                                  `tfsdk:"include_subdirectories"`
	ServerDirectory               types.String                                                  `tfsdk:"server_directory"`
	FilePattern                   types.String                                                  `tfsdk:"file_pattern"`
	Schedule                      *datasourceSecurityDlpFingerprintDatabasesScheduleModel       `tfsdk:"schedule"`
	RemoveDeletedFileFingerprints types.String                                                  `tfsdk:"remove_deleted_file_fingerprints"`
	KeepModified                  types.String                                                  `tfsdk:"keep_modified"`
	ScanOnCreation                types.String                                                  `tfsdk:"scan_on_creation"`
	Authentication                *datasourceSecurityDlpFingerprintDatabasesAuthenticationModel `tfsdk:"authentication"`
}

func (r *datasourceSecurityDlpFingerprintDatabases) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_dlp_fingerprint_databases"
}

func (r *datasourceSecurityDlpFingerprintDatabases) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Required: true,
			},
			"server": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(35),
				},
				Computed: true,
				Optional: true,
			},
			"sensitivity": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("Warning", "Private", "Critical"),
				},
				Computed: true,
				Optional: true,
			},
			"include_subdirectories": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"server_directory": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(119),
				},
				Computed: true,
				Optional: true,
			},
			"file_pattern": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(35),
				},
				Computed: true,
				Optional: true,
			},
			"remove_deleted_file_fingerprints": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"keep_modified": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"scan_on_creation": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"schedule": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"period": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("daily", "weekly", "monthly"),
						},
						Computed: true,
						Optional: true,
					},
					"sync_hour": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.AtMost(23),
						},
						Computed: true,
						Optional: true,
					},
					"sync_minute": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.AtMost(59),
						},
						Computed: true,
						Optional: true,
					},
					"weekday": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"),
						},
						Computed: true,
						Optional: true,
					},
					"sync_day_of_the_month": schema.Float64Attribute{
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
			"authentication": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"username": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.LengthAtMost(35),
						},
						Computed: true,
						Optional: true,
					},
					"password": schema.StringAttribute{
						Sensitive: true,
						Computed:  true,
						Optional:  true,
					},
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityDlpFingerprintDatabases) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_dlp_fingerprint_databases"
}

func (r *datasourceSecurityDlpFingerprintDatabases) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityDlpFingerprintDatabasesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDlpFingerprintDatabases(ctx, "read", diags))

	read_output, err := c.ReadSecurityDlpFingerprintDatabases(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDlpFingerprintDatabases(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityDlpFingerprintDatabasesModel) refreshSecurityDlpFingerprintDatabases(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["server"]; ok {
		m.Server = parseStringValue(v)
	}

	if v, ok := o["sensitivity"]; ok {
		m.Sensitivity = parseStringValue(v)
	}

	if v, ok := o["includeSubdirectories"]; ok {
		m.IncludeSubdirectories = parseStringValue(v)
	}

	if v, ok := o["serverDirectory"]; ok {
		m.ServerDirectory = parseStringValue(v)
	}

	if v, ok := o["filePattern"]; ok {
		m.FilePattern = parseStringValue(v)
	}

	if v, ok := o["schedule"]; ok {
		m.Schedule = m.Schedule.flattenSecurityDlpFingerprintDatabasesSchedule(ctx, v, &diags)
	}

	if v, ok := o["removeDeletedFileFingerprints"]; ok {
		m.RemoveDeletedFileFingerprints = parseStringValue(v)
	}

	if v, ok := o["keepModified"]; ok {
		m.KeepModified = parseStringValue(v)
	}

	if v, ok := o["scanOnCreation"]; ok {
		m.ScanOnCreation = parseStringValue(v)
	}

	if v, ok := o["authentication"]; ok {
		m.Authentication = m.Authentication.flattenSecurityDlpFingerprintDatabasesAuthentication(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceSecurityDlpFingerprintDatabasesModel) getURLObjectSecurityDlpFingerprintDatabases(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceSecurityDlpFingerprintDatabasesScheduleModel struct {
	Period            types.String  `tfsdk:"period"`
	SyncHour          types.Float64 `tfsdk:"sync_hour"`
	SyncMinute        types.Float64 `tfsdk:"sync_minute"`
	Weekday           types.String  `tfsdk:"weekday"`
	SyncDayOfTheMonth types.Float64 `tfsdk:"sync_day_of_the_month"`
}

type datasourceSecurityDlpFingerprintDatabasesAuthenticationModel struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (m *datasourceSecurityDlpFingerprintDatabasesScheduleModel) flattenSecurityDlpFingerprintDatabasesSchedule(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpFingerprintDatabasesScheduleModel {
	if input == nil {
		return &datasourceSecurityDlpFingerprintDatabasesScheduleModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpFingerprintDatabasesScheduleModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["period"]; ok {
		m.Period = parseStringValue(v)
	}

	if v, ok := o["syncHour"]; ok {
		m.SyncHour = parseFloat64Value(v)
	}

	if v, ok := o["syncMinute"]; ok {
		m.SyncMinute = parseFloat64Value(v)
	}

	if v, ok := o["weekday"]; ok {
		m.Weekday = parseStringValue(v)
	}

	if v, ok := o["syncDayOfTheMonth"]; ok {
		m.SyncDayOfTheMonth = parseFloat64Value(v)
	}

	return m
}

func (m *datasourceSecurityDlpFingerprintDatabasesAuthenticationModel) flattenSecurityDlpFingerprintDatabasesAuthentication(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceSecurityDlpFingerprintDatabasesAuthenticationModel {
	if input == nil {
		return &datasourceSecurityDlpFingerprintDatabasesAuthenticationModel{}
	}
	if m == nil {
		m = &datasourceSecurityDlpFingerprintDatabasesAuthenticationModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	return m
}
