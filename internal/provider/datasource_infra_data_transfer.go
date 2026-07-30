// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceInfraDataTransfer{}

func newDatasourceInfraDataTransfer() datasource.DataSource {
	return &datasourceInfraDataTransfer{}
}

type datasourceInfraDataTransfer struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceInfraDataTransferModel describes the datasource data model.
type datasourceInfraDataTransferModel struct {
	TenantId           types.String  `tfsdk:"tenant_id"`
	LicenseStart       types.Float64 `tfsdk:"license_start"`
	SnapshotEndMs      types.Float64 `tfsdk:"snapshot_end_ms"`
	AnnualAllotment    types.Float64 `tfsdk:"annual_allotment"`
	ConsumedBytes      types.Float64 `tfsdk:"consumed_bytes"`
	ConsumedPercent    types.Float64 `tfsdk:"consumed_percent"`
	RemainingAllotment types.Float64 `tfsdk:"remaining_allotment"`
}

func (r *datasourceInfraDataTransfer) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_infra_data_transfer"
}

func (r *datasourceInfraDataTransfer) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Get data transfer usage information.",
		Attributes: map[string]schema.Attribute{
			"tenant_id": schema.StringAttribute{
				MarkdownDescription: "Tenant ID",
				Computed:            true,
			},
			"license_start": schema.Float64Attribute{
				MarkdownDescription: "Most recent anniversary of the tenant's license start time as a UNIX timestamp in milliseconds.",
				Computed:            true,
			},
			"snapshot_end_ms": schema.Float64Attribute{
				MarkdownDescription: "Date up to which usage is calculated (exclusive of that date) as a UNIX timestamp in milliseconds.",
				Computed:            true,
			},
			"annual_allotment": schema.Float64Attribute{
				MarkdownDescription: "Number of bytes of data transfer allotted to the tenant annually.",
				Computed:            true,
			},
			"consumed_bytes": schema.Float64Attribute{
				MarkdownDescription: "Number of bytes the tenant has consumed since licenseStart.",
				Computed:            true,
			},
			"consumed_percent": schema.Float64Attribute{
				MarkdownDescription: "Percent of annualAllotment that the tenant has consumed.",
				Computed:            true,
			},
			"remaining_allotment": schema.Float64Attribute{
				MarkdownDescription: "Number of bytes remaining in the tenant's annual data transfer allotment.",
				Computed:            true,
			},
		},
	}
}

func (r *datasourceInfraDataTransfer) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_infra_data_transfer"
}

func (r *datasourceInfraDataTransfer) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceInfraDataTransferModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	var mkey interface{}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey

	read_output, err := c.ReadInfraDataTransfer(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshInfraDataTransfer(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceInfraDataTransferModel) refreshInfraDataTransfer(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["tenantId"]; ok {
		m.TenantId = parseStringValue(v)
	}

	if v, ok := o["licenseStart"]; ok {
		m.LicenseStart = parseFloat64Value(v)
	}

	if v, ok := o["snapshotEndMs"]; ok {
		m.SnapshotEndMs = parseFloat64Value(v)
	}

	if v, ok := o["annualAllotment"]; ok {
		m.AnnualAllotment = parseFloat64Value(v)
	}

	if v, ok := o["consumedBytes"]; ok {
		m.ConsumedBytes = parseFloat64Value(v)
	}

	if v, ok := o["consumedPercent"]; ok {
		m.ConsumedPercent = parseFloat64Value(v)
	}

	if v, ok := o["remainingAllotment"]; ok {
		m.RemainingAllotment = parseFloat64Value(v)
	}

	return diags
}
