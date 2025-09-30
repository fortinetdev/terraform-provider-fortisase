// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourcePrivateAccessServiceConnectionsAuth2Edl{}

func newResourcePrivateAccessServiceConnectionsAuth() resource.Resource {
	return &resourcePrivateAccessServiceConnectionsAuth2Edl{}
}

type resourcePrivateAccessServiceConnectionsAuth2Edl struct {
	fortiClient *FortiClient
}

// resourcePrivateAccessServiceConnectionsAuth2EdlModel describes the resource data model.
type resourcePrivateAccessServiceConnectionsAuth2EdlModel struct {
	ID                  types.String `tfsdk:"id"`
	Auth                types.String `tfsdk:"auth"`
	IpsecPreSharedKey   types.String `tfsdk:"ipsec_pre_shared_key"`
	IpsecPeerName       types.String `tfsdk:"ipsec_peer_name"`
	IpsecCertName       types.String `tfsdk:"ipsec_cert_name"`
	ServiceConnectionId types.String `tfsdk:"service_connection_id"`
}

func (r *resourcePrivateAccessServiceConnectionsAuth2Edl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_service_connections_auth"
}

func (r *resourcePrivateAccessServiceConnectionsAuth2Edl) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"auth": schema.StringAttribute{
				Description: "IPSEC authentication method",
				Validators: []validator.String{
					stringvalidator.OneOf("pki", "psk"),
				},
				Computed: true,
				Optional: true,
			},
			"ipsec_pre_shared_key": schema.StringAttribute{
				Description: "IPSEC auth by pre shared key.",
				Computed:    true,
				Optional:    true,
			},
			"ipsec_peer_name": schema.StringAttribute{
				Description: "Peer PKI user name that created on SASE for IPSEC authentication",
				Computed:    true,
				Optional:    true,
			},
			"ipsec_cert_name": schema.StringAttribute{
				Description: "the name of IPSEC authentication certificate that uploaded to SASE",
				Computed:    true,
				Optional:    true,
			},
			"service_connection_id": schema.StringAttribute{
				Description: "the unique uuid for service connection",
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func (r *resourcePrivateAccessServiceConnectionsAuth2Edl) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourcePrivateAccessServiceConnectionsAuth2Edl) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourcePrivateAccessServiceConnectionsAuth2EdlModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectPrivateAccessServiceConnectionsAuth(ctx, diags))
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnectionsAuth(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	_, err := c.CreatePrivateAccessServiceConnectionsAuth(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := "PrivateAccessServiceConnectionsAuth"
	data.ID = types.StringValue(mkey)

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnectionsAuth2Edl) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourcePrivateAccessServiceConnectionsAuth2EdlModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourcePrivateAccessServiceConnectionsAuth2EdlModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectPrivateAccessServiceConnectionsAuth(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnectionsAuth(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.CreatePrivateAccessServiceConnectionsAuth(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnectionsAuth2Edl) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourcePrivateAccessServiceConnectionsAuth2Edl) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourcePrivateAccessServiceConnectionsAuth2EdlModel) getCreateObjectPrivateAccessServiceConnectionsAuth(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Auth.IsNull() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	if !data.IpsecPeerName.IsNull() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if !data.IpsecCertName.IsNull() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	return &result
}

func (data *resourcePrivateAccessServiceConnectionsAuth2EdlModel) getUpdateObjectPrivateAccessServiceConnectionsAuth(ctx context.Context, state resourcePrivateAccessServiceConnectionsAuth2EdlModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Auth.IsNull() && !data.Auth.Equal(state.Auth) {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() && !data.IpsecPreSharedKey.Equal(state.IpsecPreSharedKey) {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	if !data.IpsecPeerName.IsNull() && !data.IpsecPeerName.Equal(state.IpsecPeerName) {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if !data.IpsecCertName.IsNull() && !data.IpsecCertName.Equal(state.IpsecCertName) {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	return &result
}

func (data *resourcePrivateAccessServiceConnectionsAuth2EdlModel) getURLObjectPrivateAccessServiceConnectionsAuth(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ServiceConnectionId.IsNull() {
		result["service-connection-id"] = data.ServiceConnectionId.ValueString()
	}

	return &result
}
