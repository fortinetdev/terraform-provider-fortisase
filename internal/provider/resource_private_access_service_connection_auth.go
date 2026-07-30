// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourcePrivateAccessServiceConnectionAuth2Edl{}
var _ resource.ResourceWithMoveState = &resourcePrivateAccessServiceConnectionAuth2Edl{}

func newResourcePrivateAccessServiceConnectionAuth() resource.Resource {
	return &resourcePrivateAccessServiceConnectionAuth2Edl{}
}

type resourcePrivateAccessServiceConnectionAuth2Edl struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourcePrivateAccessServiceConnectionAuth2EdlModel describes the resource data model.
type resourcePrivateAccessServiceConnectionAuth2EdlModel struct {
	ID                  types.String `tfsdk:"id"`
	Auth                types.String `tfsdk:"auth"`
	IpsecPreSharedKey   types.String `tfsdk:"ipsec_pre_shared_key"`
	IpsecPeerName       types.String `tfsdk:"ipsec_peer_name"`
	IpsecCertName       types.String `tfsdk:"ipsec_cert_name"`
	ServiceConnectionId types.String `tfsdk:"service_connection_id"`
}

func (r *resourcePrivateAccessServiceConnectionAuth2Edl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_service_connection_auth"
}

func (r *resourcePrivateAccessServiceConnectionAuth2Edl) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Secure Private Access Resource API for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("pki", "psk"),
				},
				MarkdownDescription: "IPSEC authentication method.\nSupported values: pki, psk.",
				Optional:            true,
			},
			"ipsec_pre_shared_key": schema.StringAttribute{
				MarkdownDescription: "IPSEC auth by pre shared key.",
				Optional:            true,
			},
			"ipsec_peer_name": schema.StringAttribute{
				MarkdownDescription: "Peer PKI user name that created on SASE for IPSEC authentication",
				Optional:            true,
			},
			"ipsec_cert_name": schema.StringAttribute{
				MarkdownDescription: "the name of IPSEC authentication certificate that uploaded to SASE",
				Optional:            true,
			},
			"service_connection_id": schema.StringAttribute{
				MarkdownDescription: "the unique uuid for service connection",
				Optional:            true,
			},
		},
	}
}

func (r *resourcePrivateAccessServiceConnectionAuth2Edl) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_private_access_service_connection_auth"
}
func (r *resourcePrivateAccessServiceConnectionAuth2Edl) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_private_access_service_connections_auth" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourcePrivateAccessServiceConnectionAuth2EdlModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourcePrivateAccessServiceConnectionAuth2Edl) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourcePrivateAccessServiceConnectionAuth2EdlModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectPrivateAccessServiceConnectionAuth(ctx, diags))
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnectionAuth(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreatePrivateAccessServiceConnectionsAuth(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := "PrivateAccessServiceConnectionsAuth"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnectionAuth2Edl) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourcePrivateAccessServiceConnectionAuth2EdlModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourcePrivateAccessServiceConnectionAuth2EdlModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectPrivateAccessServiceConnectionAuth(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnectionAuth(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.CreatePrivateAccessServiceConnectionsAuth(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnectionAuth2Edl) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourcePrivateAccessServiceConnectionAuth2Edl) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourcePrivateAccessServiceConnectionAuth2EdlModel) getCreateObjectPrivateAccessServiceConnectionAuth(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Auth.IsNull() && !data.Auth.IsUnknown() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() && !data.IpsecPreSharedKey.IsUnknown() {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	if !data.IpsecPeerName.IsNull() && !data.IpsecPeerName.IsUnknown() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if !data.IpsecCertName.IsNull() && !data.IpsecCertName.IsUnknown() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	return &result
}

func (data *resourcePrivateAccessServiceConnectionAuth2EdlModel) getUpdateObjectPrivateAccessServiceConnectionAuth(ctx context.Context, state resourcePrivateAccessServiceConnectionAuth2EdlModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Auth.IsNull() && !data.Auth.IsUnknown() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() && !data.IpsecPreSharedKey.IsUnknown() {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	if !data.IpsecPeerName.IsNull() && !data.IpsecPeerName.IsUnknown() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if !data.IpsecCertName.IsNull() && !data.IpsecCertName.IsUnknown() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	return &result
}

func (data *resourcePrivateAccessServiceConnectionAuth2EdlModel) getURLObjectPrivateAccessServiceConnectionAuth(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ServiceConnectionId.IsNull() && !data.ServiceConnectionId.IsUnknown() {
		result["service-connection-id"] = data.ServiceConnectionId.ValueString()
	}

	return &result
}
