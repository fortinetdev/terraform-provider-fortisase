// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/float64validatorwarning"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourceEndpointGroupInvitationCode{}

func newDatasourceEndpointGroupInvitationCode() datasource.DataSource {
	return &datasourceEndpointGroupInvitationCode{}
}

type datasourceEndpointGroupInvitationCode struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourceEndpointGroupInvitationCodeModel describes the datasource data model.
type datasourceEndpointGroupInvitationCodeModel struct {
	PrimaryKey      types.String                                               `tfsdk:"primary_key"`
	ExpireDate      types.String                                               `tfsdk:"expire_date"`
	GroupAssignment *datasourceEndpointGroupInvitationCodeGroupAssignmentModel `tfsdk:"group_assignment"`
}

func (r *datasourceEndpointGroupInvitationCode) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_group_invitation_code"
}

func (r *datasourceEndpointGroupInvitationCode) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Group-Based Invitation Code Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 128),
				},
				Required: true,
			},
			"expire_date": schema.StringAttribute{
				Computed: true,
			},
			"group_assignment": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Computed: true,
					},
					"group": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"id": schema.Float64Attribute{
								Validators: []validator.Float64{
									float64validatorwarning.AtLeast(1),
								},
								Computed: true,
							},
							"path": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidatorwarning.LengthAtLeast(1),
								},
								Computed: true,
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourceEndpointGroupInvitationCode) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_endpoint_group_invitation_code"
}

func (r *datasourceEndpointGroupInvitationCode) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourceEndpointGroupInvitationCodeModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.PrimaryKey.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointGroupInvitationCode(ctx, "read", diags))

	read_output, err := c.ReadEndpointGroupInvitationCodes(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshEndpointGroupInvitationCode(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourceEndpointGroupInvitationCodeModel) refreshEndpointGroupInvitationCode(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["expireDate"]; ok {
		m.ExpireDate = parseStringValue(v)
	}

	if v, ok := o["groupAssignment"]; ok {
		m.GroupAssignment = m.GroupAssignment.flattenEndpointGroupInvitationCodeGroupAssignment(ctx, v, &diags)
	}

	return diags
}

func (data *datasourceEndpointGroupInvitationCodeModel) getURLObjectEndpointGroupInvitationCode(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type datasourceEndpointGroupInvitationCodeGroupAssignmentModel struct {
	Enabled types.Bool                                                      `tfsdk:"enabled"`
	Group   *datasourceEndpointGroupInvitationCodeGroupAssignmentGroupModel `tfsdk:"group"`
}

type datasourceEndpointGroupInvitationCodeGroupAssignmentGroupModel struct {
	Id   types.Float64 `tfsdk:"id"`
	Path types.String  `tfsdk:"path"`
}

func (m *datasourceEndpointGroupInvitationCodeGroupAssignmentModel) flattenEndpointGroupInvitationCodeGroupAssignment(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointGroupInvitationCodeGroupAssignmentModel {
	if input == nil {
		return &datasourceEndpointGroupInvitationCodeGroupAssignmentModel{}
	}
	if m == nil {
		m = &datasourceEndpointGroupInvitationCodeGroupAssignmentModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["enabled"]; ok {
		m.Enabled = parseBoolValue(v)
	}

	if v, ok := o["group"]; ok {
		m.Group = m.Group.flattenEndpointGroupInvitationCodeGroupAssignmentGroup(ctx, v, diags)
	}

	return m
}

func (m *datasourceEndpointGroupInvitationCodeGroupAssignmentGroupModel) flattenEndpointGroupInvitationCodeGroupAssignmentGroup(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourceEndpointGroupInvitationCodeGroupAssignmentGroupModel {
	if input == nil {
		return &datasourceEndpointGroupInvitationCodeGroupAssignmentGroupModel{}
	}
	if m == nil {
		m = &datasourceEndpointGroupInvitationCodeGroupAssignmentGroupModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseFloat64Value(v)
	}

	if v, ok := o["path"]; ok {
		m.Path = parseStringValue(v)
	}

	return m
}
