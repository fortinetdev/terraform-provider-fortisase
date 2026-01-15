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
var _ datasource.DataSource = &datasourceSecurityRecurringSchedules{}

func newDatasourceSecurityRecurringSchedules() datasource.DataSource {
	return &datasourceSecurityRecurringSchedules{}
}

type datasourceSecurityRecurringSchedules struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceSecurityRecurringSchedulesModel describes the datasource data model.
type datasourceSecurityRecurringSchedulesModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Days       types.Set    `tfsdk:"days"`
	StartTime  types.String `tfsdk:"start_time"`
	EndTime    types.String `tfsdk:"end_time"`
}

func (r *datasourceSecurityRecurringSchedules) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_recurring_schedules"
}

func (r *datasourceSecurityRecurringSchedules) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 31),
				},
				Required: true,
			},
			"days": schema.SetAttribute{
				Validators: []validator.Set{
					setvalidator.SizeBetween(1, 7),
				},
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"start_time": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"end_time": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *datasourceSecurityRecurringSchedules) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_recurring_schedules"
}

func (r *datasourceSecurityRecurringSchedules) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceSecurityRecurringSchedulesModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityRecurringSchedules(ctx, "read", diags))

	read_output, err := c.ReadSecurityRecurringSchedules(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityRecurringSchedules(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceSecurityRecurringSchedulesModel) refreshSecurityRecurringSchedules(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["days"]; ok {
		m.Days = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["startTime"]; ok {
		m.StartTime = parseStringValue(v)
	}

	if v, ok := o["endTime"]; ok {
		m.EndTime = parseStringValue(v)
	}

	return diags
}

func (data *datasourceSecurityRecurringSchedulesModel) getURLObjectSecurityRecurringSchedules(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
