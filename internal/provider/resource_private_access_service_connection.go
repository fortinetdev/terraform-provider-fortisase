// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/validators/stringvalidatorwarning"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourcePrivateAccessServiceConnection{}
var _ resource.ResourceWithMoveState = &resourcePrivateAccessServiceConnection{}

func newResourcePrivateAccessServiceConnection() resource.Resource {
	return &resourcePrivateAccessServiceConnection{}
}

type resourcePrivateAccessServiceConnection struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourcePrivateAccessServiceConnectionModel describes the resource data model.
type resourcePrivateAccessServiceConnectionModel struct {
	ID                  types.String                                             `tfsdk:"id"`
	Alias               types.String                                             `tfsdk:"alias"`
	BgpPeerIp           types.String                                             `tfsdk:"bgp_peer_ip"`
	IpsecRemoteGw       types.String                                             `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId    types.String                                             `tfsdk:"overlay_network_id"`
	RouteMapTag         types.String                                             `tfsdk:"route_map_tag"`
	Auth                types.String                                             `tfsdk:"auth"`
	IpsecPreSharedKey   types.String                                             `tfsdk:"ipsec_pre_shared_key"`
	IpsecCertName       types.String                                             `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion     types.String                                             `tfsdk:"ipsec_ike_version"`
	IpsecPeerName       types.String                                             `tfsdk:"ipsec_peer_name"`
	BackupLinks         []resourcePrivateAccessServiceConnectionBackupLinksModel `tfsdk:"backup_links"`
	Ftntid              types.String                                             `tfsdk:"ftntid"`
	Type                types.String                                             `tfsdk:"type"`
	ConfigState         types.String                                             `tfsdk:"config_state"`
	SeqNum              types.Float64                                            `tfsdk:"seq_num"`
	FailedMessage       types.String                                             `tfsdk:"failed_message"`
	Config              *resourcePrivateAccessServiceConnectionConfigModel       `tfsdk:"config"`
	CommonConfig        *resourcePrivateAccessServiceConnectionCommonConfigModel `tfsdk:"common_config"`
	IpAssigned          []resourcePrivateAccessServiceConnectionIpAssignedModel  `tfsdk:"ip_assigned"`
	RegionCost          types.Map                                                `tfsdk:"region_cost"`
	ServiceConnectionId types.String                                             `tfsdk:"service_connection_id"`
}

func (r *resourcePrivateAccessServiceConnection) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_service_connection"
}

func (r *resourcePrivateAccessServiceConnection) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"alias": schema.StringAttribute{
				MarkdownDescription: "alias for serivce connection",
				Optional:            true,
			},
			"bgp_peer_ip": schema.StringAttribute{
				MarkdownDescription: "BGP Routing Peer IP.",
				Optional:            true,
			},
			"ipsec_remote_gw": schema.StringAttribute{
				MarkdownDescription: "IPSEC Remote Gateway IP",
				Optional:            true,
			},
			"overlay_network_id": schema.StringAttribute{
				MarkdownDescription: "integer id for overlay",
				Optional:            true,
			},
			"route_map_tag": schema.StringAttribute{
				MarkdownDescription: "route map tag",
				Optional:            true,
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
			"ipsec_cert_name": schema.StringAttribute{
				MarkdownDescription: "the name of IPSEC authentication certificate that uploaded to SASE",
				Optional:            true,
			},
			"ipsec_ike_version": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("2"),
				},
				MarkdownDescription: "IKE version for IPSEC.\nSupported values: 2.",
				Optional:            true,
			},
			"ipsec_peer_name": schema.StringAttribute{
				MarkdownDescription: "Peer PKI user name that created on SASE for IPSEC authentication",
				Optional:            true,
			},
			"ftntid": schema.StringAttribute{
				MarkdownDescription: "unique id for service connection",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("overlay", "loopback"),
				},
				MarkdownDescription: "BGP Routing Design. Must be same as network configuration.\nSupported values: overlay, loopback.",
				Computed:            true,
				Optional:            true,
			},
			"config_state": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("success", "failed", "creating", "updating", "deleting"),
				},
				MarkdownDescription: "Configuration state of service connection.\nSupported values: success, failed, creating, updating, deleting.",
				Computed:            true,
			},
			"seq_num": schema.Float64Attribute{
				MarkdownDescription: "sequential unique number for service connection",
				Computed:            true,
			},
			"failed_message": schema.StringAttribute{
				MarkdownDescription: "failure message while config service connection",
				Computed:            true,
			},
			"region_cost": schema.MapAttribute{
				MarkdownDescription: "Cost value to determine the priority of SASE spokes. Default cost is 5 if not provided through initial api request.",
				Optional:            true,
				ElementType:         types.Int64Type,
			},
			"service_connection_id": schema.StringAttribute{
				MarkdownDescription: "the unique uuid for service connection",
				Computed:            true,
			},
			"backup_links": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"alias": schema.StringAttribute{
							MarkdownDescription: "alias for serivce connection additional overlay",
							Optional:            true,
						},
						"auth": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("pki", "psk"),
							},
							MarkdownDescription: "IPSEC authentication method.\nSupported values: pki, psk.",
							Optional:            true,
						},
						"ipsec_cert_name": schema.StringAttribute{
							MarkdownDescription: "the name of IPSEC authentication certificate that uploaded to SASE",
							Optional:            true,
						},
						"ipsec_ike_version": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidatorwarning.OneOf("2"),
							},
							MarkdownDescription: "IKE version for IPSEC.\nSupported values: 2.",
							Optional:            true,
						},
						"ipsec_peer_name": schema.StringAttribute{
							MarkdownDescription: "Peer PKI user name that created on SASE for IPSEC authentication",
							Optional:            true,
						},
						"ipsec_remote_gw": schema.StringAttribute{
							MarkdownDescription: "IPSEC Remote Gateway IP",
							Optional:            true,
						},
						"overlay_network_id": schema.StringAttribute{
							MarkdownDescription: "integer id for overlay",
							Optional:            true,
						},
						"ipsec_pre_shared_key": schema.StringAttribute{
							MarkdownDescription: "IPSEC auth by pre shared key.",
							Optional:            true,
						},
					},
				},
				Optional: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
			},
			"config": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"alias": schema.StringAttribute{
						MarkdownDescription: "alias for serivce connection",
						Computed:            true,
					},
					"auth": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("pki", "psk"),
						},
						MarkdownDescription: "IPSEC authentication method.\nSupported values: pki, psk.",
						Computed:            true,
					},
					"bgp_peer_ip": schema.StringAttribute{
						MarkdownDescription: "BGP Routing Peer IP.",
						Computed:            true,
					},
					"ipsec_cert_name": schema.StringAttribute{
						MarkdownDescription: "the name of IPSEC authentication certificate that uploaded to SASE",
						Computed:            true,
					},
					"ipsec_ike_version": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("2"),
						},
						MarkdownDescription: "IKE version for IPSEC.\nSupported values: 2.",
						Computed:            true,
					},
					"ipsec_peer_name": schema.StringAttribute{
						MarkdownDescription: "Peer PKI user name that created on SASE for IPSEC authentication",
						Computed:            true,
					},
					"ipsec_remote_gw": schema.StringAttribute{
						MarkdownDescription: "IPSEC Remote Gateway IP",
						Computed:            true,
					},
					"overlay_network_id": schema.StringAttribute{
						MarkdownDescription: "integer id for overlay",
						Computed:            true,
					},
					"route_map_tag": schema.StringAttribute{
						MarkdownDescription: "route map tag",
						Computed:            true,
					},
					"region_cost": schema.MapAttribute{
						MarkdownDescription: "cost value to determine the priority of SASE spokes",
						Computed:            true,
						ElementType:         types.Int64Type,
					},
					"backup_links": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									MarkdownDescription: "unique id for additional IPsec overlays.",
									Computed:            true,
								},
								"alias": schema.StringAttribute{
									MarkdownDescription: "alias for serivce connection additional overlay",
									Computed:            true,
								},
								"auth": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("pki", "psk"),
									},
									MarkdownDescription: "IPSEC authentication method.\nSupported values: pki, psk.",
									Computed:            true,
								},
								"ipsec_cert_name": schema.StringAttribute{
									MarkdownDescription: "the name of IPSEC authentication certificate that uploaded to SASE",
									Computed:            true,
								},
								"ipsec_ike_version": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidatorwarning.OneOf("2"),
									},
									MarkdownDescription: "IKE version for IPSEC.\nSupported values: 2.",
									Computed:            true,
								},
								"ipsec_peer_name": schema.StringAttribute{
									MarkdownDescription: "Peer PKI user name that created on SASE for IPSEC authentication",
									Computed:            true,
								},
								"ipsec_remote_gw": schema.StringAttribute{
									MarkdownDescription: "IPSEC Remote Gateway IP",
									Computed:            true,
								},
								"overlay_network_id": schema.StringAttribute{
									MarkdownDescription: "integer id for overlay",
									Computed:            true,
								},
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
			"common_config": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"config_state": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("success", "failed", "creating", "updating", "deleting"),
						},
						MarkdownDescription: "Configuration state of network configuration.\nSupported values: success, failed, creating, updating, deleting.",
						Computed:            true,
					},
					"bgp_design": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidatorwarning.OneOf("overlay", "loopback"),
						},
						MarkdownDescription: "BGP Routing Design.\nSupported values: overlay, loopback.",
						Computed:            true,
					},
					"bgp_router_ids_subnet": schema.StringAttribute{
						MarkdownDescription: "Available/unused subnet that can be used to assign loopback interface IP addresses used for BGP router IDs parameter on the FortiSASE security PoPs. /28 is the minimum subnet size.",
						Computed:            true,
					},
					"as_number": schema.StringAttribute{
						MarkdownDescription: "Autonomous System Number (ASN).",
						Computed:            true,
					},
					"recursive_next_hop": schema.BoolAttribute{
						MarkdownDescription: "BGP Recursive Routing. Enabling this setting allows for interhub connectivity. When use BGP design on-loopback this has to be enabled.",
						Computed:            true,
					},
					"sdwan_rule_enable": schema.BoolAttribute{
						MarkdownDescription: "Hub Selection Method. Enabling this setting the highest priority service connection that meets minimum SLA requirements is selected. Otherwise BGP MED (Multi-Exit Discriminator) will be used",
						Computed:            true,
					},
					"sdwan_health_check_vm": schema.StringAttribute{
						MarkdownDescription: "Health Check IP. Must be provided when enable sdwan rule which used to obtain Jitter, latency and packet loss measurements.",
						Computed:            true,
					},
				},
				Computed: true,
			},
			"ip_assigned": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "unique id for bgp router id assignment",
							Computed:            true,
						},
						"sdwan_common_id": schema.StringAttribute{
							MarkdownDescription: "unique id related to network configuration",
							Computed:            true,
						},
						"bgp_router_id": schema.StringAttribute{
							MarkdownDescription: "BGP Router ID generated from Router ID Subnets",
							Computed:            true,
						},
						"site_id": schema.StringAttribute{
							MarkdownDescription: "id for SASE spoke",
							Computed:            true,
						},
						"region": schema.StringAttribute{
							MarkdownDescription: "air port code for SASE spoke physical region",
							Computed:            true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *resourcePrivateAccessServiceConnection) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_private_access_service_connection"
}
func (r *resourcePrivateAccessServiceConnection) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_private_access_service_connections" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourcePrivateAccessServiceConnectionModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourcePrivateAccessServiceConnection) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("PrivateAccessServiceConnections")
	lock.Lock()
	defer lock.Unlock()
	var data resourcePrivateAccessServiceConnectionModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectPrivateAccessServiceConnection(ctx, diags))
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnection(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreatePrivateAccessServiceConnections(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["id"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnection(ctx, "read", diags))

	read_output := make(map[string]interface{})
	for i := 0; i < 20; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadPrivateAccessServiceConnections(&read_input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
				getErrorDetail(&read_input_model, read_output),
			)
			return
		}
		if v, ok := read_output["config_state"]; ok {
			if v == "failed" {
				// // resend the request
				// input_model.Mkey = mkey
				// output, err = c.UpdatePrivateAccessServiceConnections(&input_model)
				// if err != nil {
				// 	diags.AddError(
				// 		fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
				// 		getErrorDetail(&input_model, output),
				// 	)
				// 	return
				// }
				// continue
				failedMessage := ""
				if v, ok := read_output["failed_message"]; ok {
					if msg, ok := v.(string); ok {
						failedMessage = msg
					} else if v != nil {
						failedMessage = fmt.Sprintf("%v", v)
					}
				}
				diags.AddWarning(
					"The configuration state is failed.",
					fmt.Sprintf("The resource has been created successfully, but its configuration state is failed. Please check the configuration. Error message: %s", failedMessage),
				)
				break
			}
			if v != "success" {
				continue
			}
		}
		break
	}

	diags.Append(data.refreshPrivateAccessServiceConnection(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnection) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("PrivateAccessServiceConnections")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourcePrivateAccessServiceConnectionModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourcePrivateAccessServiceConnectionModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectPrivateAccessServiceConnection(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnection(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdatePrivateAccessServiceConnections(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnection(ctx, "read", diags))

	read_output := make(map[string]interface{})
	for i := 0; i < 20; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadPrivateAccessServiceConnections(&read_input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
				getErrorDetail(&read_input_model, read_output),
			)
			return
		}
		if v, ok := read_output["config_state"]; ok {
			if v == "failed" {
				// // resend the request
				// output, err = c.UpdatePrivateAccessServiceConnections(&input_model)
				// if err != nil {
				// 	diags.AddError(
				// 		fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
				// 		getErrorDetail(&input_model, output),
				// 	)
				// 	return
				// }
				// continue
				failedMessage := ""
				if v, ok := read_output["failed_message"]; ok {
					if msg, ok := v.(string); ok {
						failedMessage = msg
					} else if v != nil {
						failedMessage = fmt.Sprintf("%v", v)
					}
				}
				diags.AddWarning(
					"The configuration state is failed.",
					fmt.Sprintf("The resource was updated successfully, but its configuration state is failed. Please check the configuration. Error message: %s", failedMessage),
				)
				break
			}
			if v != "success" {
				continue
			}
		}
		break
	}

	diags.Append(data.refreshPrivateAccessServiceConnection(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnection) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("PrivateAccessServiceConnections")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourcePrivateAccessServiceConnectionModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnection(ctx, "delete", diags))

	output, err := c.DeletePrivateAccessServiceConnections(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	read_output := make(map[string]interface{})
	for i := 0; i < 20; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadPrivateAccessServiceConnections(&input_model)
		if err != nil || len(read_output) == 0 {
			// Delete success
			return
		}
	}
	diags.AddError(
		fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
		fmt.Sprintf("The resource still exists %s: %v", r.resourceName, read_output),
	)
}

func (r *resourcePrivateAccessServiceConnection) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourcePrivateAccessServiceConnectionModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnection(ctx, "read", diags))

	read_output, err := c.ReadPrivateAccessServiceConnections(&input_model)
	if err != nil {
		if isNotFoundResponse(read_output) {
			resp.State.RemoveResource(ctx)
			return
		}
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshPrivateAccessServiceConnection(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnection) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("service_connection_id"), req.ID)...)
}

func (m *resourcePrivateAccessServiceConnectionModel) refreshPrivateAccessServiceConnection(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["id"]; ok {
		m.Ftntid = parseStringValue(v)
	}

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["config_state"]; ok {
		m.ConfigState = parseStringValue(v)
	}

	if v, ok := o["seq_num"]; ok {
		m.SeqNum = parseFloat64Value(v)
	}

	if v, ok := o["failed_message"]; ok {
		m.FailedMessage = parseStringValue(v)
	}

	if v, ok := o["config"]; ok {
		m.Config = m.Config.flattenPrivateAccessServiceConnectionConfig(ctx, v, &diags)
	}

	if v, ok := o["common_config"]; ok {
		m.CommonConfig = m.CommonConfig.flattenPrivateAccessServiceConnectionCommonConfig(ctx, v, &diags)
	}

	if v, ok := o["ip_assigned"]; ok {
		m.IpAssigned = m.flattenPrivateAccessServiceConnectionIpAssignedList(ctx, v, &diags)
	}

	return diags
}

func (data *resourcePrivateAccessServiceConnectionModel) getCreateObjectPrivateAccessServiceConnection(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Alias.IsNull() && !data.Alias.IsUnknown() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.BgpPeerIp.IsNull() && !data.BgpPeerIp.IsUnknown() {
		result["bgp_peer_ip"] = data.BgpPeerIp.ValueString()
	}

	if !data.IpsecRemoteGw.IsNull() && !data.IpsecRemoteGw.IsUnknown() {
		result["ipsec_remote_gw"] = data.IpsecRemoteGw.ValueString()
	}

	if !data.OverlayNetworkId.IsNull() && !data.OverlayNetworkId.IsUnknown() {
		result["overlay_network_id"] = data.OverlayNetworkId.ValueString()
	}

	if !data.RouteMapTag.IsNull() && !data.RouteMapTag.IsUnknown() {
		result["route_map_tag"] = data.RouteMapTag.ValueString()
	}

	if !data.Auth.IsNull() && !data.Auth.IsUnknown() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() && !data.IpsecPreSharedKey.IsUnknown() {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	if !data.IpsecCertName.IsNull() && !data.IpsecCertName.IsUnknown() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	if !data.IpsecIkeVersion.IsNull() && !data.IpsecIkeVersion.IsUnknown() {
		result["ipsec_ike_version"] = data.IpsecIkeVersion.ValueString()
	}

	if !data.IpsecPeerName.IsNull() && !data.IpsecPeerName.IsUnknown() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if data.BackupLinks != nil {
		result["backup_links"] = data.expandPrivateAccessServiceConnectionBackupLinksList(ctx, data.BackupLinks, diags)
	}

	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		result["type"] = data.Type.ValueString()
	}

	if !data.RegionCost.IsNull() && !data.RegionCost.IsUnknown() {
		result["region_cost"] = data.RegionCost.Elements()
	}

	return &result
}

func (data *resourcePrivateAccessServiceConnectionModel) getUpdateObjectPrivateAccessServiceConnection(ctx context.Context, state resourcePrivateAccessServiceConnectionModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Alias.IsNull() && !data.Alias.IsUnknown() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.BgpPeerIp.IsNull() && !data.BgpPeerIp.IsUnknown() {
		result["bgp_peer_ip"] = data.BgpPeerIp.ValueString()
	}

	if !data.IpsecRemoteGw.IsNull() && !data.IpsecRemoteGw.IsUnknown() {
		result["ipsec_remote_gw"] = data.IpsecRemoteGw.ValueString()
	}

	if !data.OverlayNetworkId.IsNull() && !data.OverlayNetworkId.IsUnknown() {
		result["overlay_network_id"] = data.OverlayNetworkId.ValueString()
	}

	if !data.RouteMapTag.IsNull() && !data.RouteMapTag.IsUnknown() {
		result["route_map_tag"] = data.RouteMapTag.ValueString()
	}

	if !data.Auth.IsNull() && !data.Auth.IsUnknown() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() && !data.IpsecPreSharedKey.IsUnknown() {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	if !data.IpsecCertName.IsNull() && !data.IpsecCertName.IsUnknown() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	if !data.IpsecIkeVersion.IsNull() && !data.IpsecIkeVersion.IsUnknown() {
		result["ipsec_ike_version"] = data.IpsecIkeVersion.ValueString()
	}

	if !data.IpsecPeerName.IsNull() && !data.IpsecPeerName.IsUnknown() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	return &result
}

func (data *resourcePrivateAccessServiceConnectionModel) getURLObjectPrivateAccessServiceConnection(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ServiceConnectionId.IsNull() && !data.ServiceConnectionId.IsUnknown() {
		result["service-connection-id"] = data.ServiceConnectionId.ValueString()
	}

	return &result
}

type resourcePrivateAccessServiceConnectionBackupLinksModel struct {
	Alias             types.String `tfsdk:"alias"`
	Auth              types.String `tfsdk:"auth"`
	IpsecCertName     types.String `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion   types.String `tfsdk:"ipsec_ike_version"`
	IpsecPeerName     types.String `tfsdk:"ipsec_peer_name"`
	IpsecRemoteGw     types.String `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId  types.String `tfsdk:"overlay_network_id"`
	IpsecPreSharedKey types.String `tfsdk:"ipsec_pre_shared_key"`
}

type resourcePrivateAccessServiceConnectionConfigModel struct {
	Alias            types.String                                                   `tfsdk:"alias"`
	Auth             types.String                                                   `tfsdk:"auth"`
	BgpPeerIp        types.String                                                   `tfsdk:"bgp_peer_ip"`
	IpsecCertName    types.String                                                   `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion  types.String                                                   `tfsdk:"ipsec_ike_version"`
	IpsecPeerName    types.String                                                   `tfsdk:"ipsec_peer_name"`
	IpsecRemoteGw    types.String                                                   `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId types.String                                                   `tfsdk:"overlay_network_id"`
	RouteMapTag      types.String                                                   `tfsdk:"route_map_tag"`
	RegionCost       types.Map                                                      `tfsdk:"region_cost"`
	BackupLinks      []resourcePrivateAccessServiceConnectionConfigBackupLinksModel `tfsdk:"backup_links"`
}

type resourcePrivateAccessServiceConnectionConfigBackupLinksModel struct {
	Id               types.String `tfsdk:"id"`
	Alias            types.String `tfsdk:"alias"`
	Auth             types.String `tfsdk:"auth"`
	IpsecCertName    types.String `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion  types.String `tfsdk:"ipsec_ike_version"`
	IpsecPeerName    types.String `tfsdk:"ipsec_peer_name"`
	IpsecRemoteGw    types.String `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId types.String `tfsdk:"overlay_network_id"`
}

type resourcePrivateAccessServiceConnectionCommonConfigModel struct {
	ConfigState        types.String `tfsdk:"config_state"`
	BgpDesign          types.String `tfsdk:"bgp_design"`
	BgpRouterIdsSubnet types.String `tfsdk:"bgp_router_ids_subnet"`
	AsNumber           types.String `tfsdk:"as_number"`
	RecursiveNextHop   types.Bool   `tfsdk:"recursive_next_hop"`
	SdwanRuleEnable    types.Bool   `tfsdk:"sdwan_rule_enable"`
	SdwanHealthCheckVm types.String `tfsdk:"sdwan_health_check_vm"`
}

type resourcePrivateAccessServiceConnectionIpAssignedModel struct {
	Id            types.String `tfsdk:"id"`
	SdwanCommonId types.String `tfsdk:"sdwan_common_id"`
	BgpRouterId   types.String `tfsdk:"bgp_router_id"`
	SiteId        types.String `tfsdk:"site_id"`
	Region        types.String `tfsdk:"region"`
}

func (m *resourcePrivateAccessServiceConnectionBackupLinksModel) flattenPrivateAccessServiceConnectionBackupLinks(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionBackupLinksModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionBackupLinksModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionBackupLinksModel{}
	}

	return m
}

func (s *resourcePrivateAccessServiceConnectionModel) flattenPrivateAccessServiceConnectionBackupLinksList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourcePrivateAccessServiceConnectionBackupLinksModel {
	if o == nil {
		return []resourcePrivateAccessServiceConnectionBackupLinksModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument backup_links is not type of []interface{}.", "")
		return []resourcePrivateAccessServiceConnectionBackupLinksModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourcePrivateAccessServiceConnectionBackupLinksModel{}
	}

	values := make([]resourcePrivateAccessServiceConnectionBackupLinksModel, len(l))
	for i, ele := range l {
		var m resourcePrivateAccessServiceConnectionBackupLinksModel
		if i < len(s.BackupLinks) {
			m = s.BackupLinks[i]
		}
		values[i] = *m.flattenPrivateAccessServiceConnectionBackupLinks(ctx, ele, diags)
	}

	return values
}

func (m *resourcePrivateAccessServiceConnectionConfigModel) flattenPrivateAccessServiceConnectionConfig(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionConfigModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionConfigModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionConfigModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["alias"]; ok {
		m.Alias = parseStringValue(v)
	}

	if v, ok := o["auth"]; ok {
		m.Auth = parseStringValue(v)
	}

	if v, ok := o["bgp_peer_ip"]; ok {
		m.BgpPeerIp = parseStringValue(v)
	}

	if v, ok := o["ipsec_cert_name"]; ok {
		m.IpsecCertName = parseStringValue(v)
	}

	if v, ok := o["ipsec_ike_version"]; ok {
		m.IpsecIkeVersion = parseStringValue(v)
	}

	if v, ok := o["ipsec_peer_name"]; ok {
		m.IpsecPeerName = parseStringValue(v)
	}

	if v, ok := o["ipsec_remote_gw"]; ok {
		m.IpsecRemoteGw = parseStringValue(v)
	}

	if v, ok := o["overlay_network_id"]; ok {
		m.OverlayNetworkId = parseStringValue(v)
	}

	if v, ok := o["route_map_tag"]; ok {
		m.RouteMapTag = parseStringValue(v)
	}

	if v, ok := o["region_cost"]; ok {
		m.RegionCost = parseMapValue(ctx, v, types.Int64Type)
	}

	if v, ok := o["backup_links"]; ok {
		m.BackupLinks = m.flattenPrivateAccessServiceConnectionConfigBackupLinksList(ctx, v, diags)
	}

	return m
}

func (m *resourcePrivateAccessServiceConnectionConfigBackupLinksModel) flattenPrivateAccessServiceConnectionConfigBackupLinks(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionConfigBackupLinksModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionConfigBackupLinksModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionConfigBackupLinksModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseStringValue(v)
	}

	if v, ok := o["alias"]; ok {
		m.Alias = parseStringValue(v)
	}

	if v, ok := o["auth"]; ok {
		m.Auth = parseStringValue(v)
	}

	if v, ok := o["ipsec_cert_name"]; ok {
		m.IpsecCertName = parseStringValue(v)
	}

	if v, ok := o["ipsec_ike_version"]; ok {
		m.IpsecIkeVersion = parseStringValue(v)
	}

	if v, ok := o["ipsec_peer_name"]; ok {
		m.IpsecPeerName = parseStringValue(v)
	}

	if v, ok := o["ipsec_remote_gw"]; ok {
		m.IpsecRemoteGw = parseStringValue(v)
	}

	if v, ok := o["overlay_network_id"]; ok {
		m.OverlayNetworkId = parseStringValue(v)
	}

	return m
}

func (s *resourcePrivateAccessServiceConnectionConfigModel) flattenPrivateAccessServiceConnectionConfigBackupLinksList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourcePrivateAccessServiceConnectionConfigBackupLinksModel {
	if o == nil {
		return []resourcePrivateAccessServiceConnectionConfigBackupLinksModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument backup_links is not type of []interface{}.", "")
		return []resourcePrivateAccessServiceConnectionConfigBackupLinksModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourcePrivateAccessServiceConnectionConfigBackupLinksModel{}
	}

	values := make([]resourcePrivateAccessServiceConnectionConfigBackupLinksModel, len(l))
	for i, ele := range l {
		var m resourcePrivateAccessServiceConnectionConfigBackupLinksModel
		if i < len(s.BackupLinks) {
			m = s.BackupLinks[i]
		}
		values[i] = *m.flattenPrivateAccessServiceConnectionConfigBackupLinks(ctx, ele, diags)
	}

	return values
}

func (m *resourcePrivateAccessServiceConnectionCommonConfigModel) flattenPrivateAccessServiceConnectionCommonConfig(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionCommonConfigModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionCommonConfigModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionCommonConfigModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["config_state"]; ok {
		m.ConfigState = parseStringValue(v)
	}

	if v, ok := o["bgp_design"]; ok {
		m.BgpDesign = parseStringValue(v)
	}

	if v, ok := o["bgp_router_ids_subnet"]; ok {
		m.BgpRouterIdsSubnet = parseStringValue(v)
	}

	if v, ok := o["as_number"]; ok {
		m.AsNumber = parseStringValue(v)
	}

	if v, ok := o["recursive_next_hop"]; ok {
		m.RecursiveNextHop = parseBoolValue(v)
	}

	if v, ok := o["sdwan_rule_enable"]; ok {
		m.SdwanRuleEnable = parseBoolValue(v)
	}

	if v, ok := o["sdwan_health_check_vm"]; ok {
		m.SdwanHealthCheckVm = parseStringValue(v)
	}

	return m
}

func (m *resourcePrivateAccessServiceConnectionIpAssignedModel) flattenPrivateAccessServiceConnectionIpAssigned(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionIpAssignedModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionIpAssignedModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionIpAssignedModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseStringValue(v)
	}

	if v, ok := o["sdwan_common_id"]; ok {
		m.SdwanCommonId = parseStringValue(v)
	}

	if v, ok := o["bgp_router_id"]; ok {
		m.BgpRouterId = parseStringValue(v)
	}

	if v, ok := o["site_id"]; ok {
		m.SiteId = parseStringValue(v)
	}

	if v, ok := o["region"]; ok {
		m.Region = parseStringValue(v)
	}

	return m
}

func (s *resourcePrivateAccessServiceConnectionModel) flattenPrivateAccessServiceConnectionIpAssignedList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourcePrivateAccessServiceConnectionIpAssignedModel {
	if o == nil {
		return []resourcePrivateAccessServiceConnectionIpAssignedModel{}
	}

	var l []interface{}
	switch v := o.(type) {
	case []interface{}:
		l = v
	case map[string]interface{}:
		l = []interface{}{v}
	default:
		diags.AddError("Argument ip_assigned is not type of []interface{}.", "")
		return []resourcePrivateAccessServiceConnectionIpAssignedModel{}
	}

	if len(l) == 0 || l[0] == nil {
		return []resourcePrivateAccessServiceConnectionIpAssignedModel{}
	}

	values := make([]resourcePrivateAccessServiceConnectionIpAssignedModel, len(l))
	for i, ele := range l {
		var m resourcePrivateAccessServiceConnectionIpAssignedModel
		if i < len(s.IpAssigned) {
			m = s.IpAssigned[i]
		}
		values[i] = *m.flattenPrivateAccessServiceConnectionIpAssigned(ctx, ele, diags)
	}

	return values
}

func (data *resourcePrivateAccessServiceConnectionBackupLinksModel) expandPrivateAccessServiceConnectionBackupLinks(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Alias.IsNull() && !data.Alias.IsUnknown() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.Auth.IsNull() && !data.Auth.IsUnknown() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecCertName.IsNull() && !data.IpsecCertName.IsUnknown() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	if !data.IpsecIkeVersion.IsNull() && !data.IpsecIkeVersion.IsUnknown() {
		result["ipsec_ike_version"] = data.IpsecIkeVersion.ValueString()
	}

	if !data.IpsecPeerName.IsNull() && !data.IpsecPeerName.IsUnknown() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if !data.IpsecRemoteGw.IsNull() && !data.IpsecRemoteGw.IsUnknown() {
		result["ipsec_remote_gw"] = data.IpsecRemoteGw.ValueString()
	}

	if !data.OverlayNetworkId.IsNull() && !data.OverlayNetworkId.IsUnknown() {
		result["overlay_network_id"] = data.OverlayNetworkId.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() && !data.IpsecPreSharedKey.IsUnknown() {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	return result
}

func (s *resourcePrivateAccessServiceConnectionModel) expandPrivateAccessServiceConnectionBackupLinksList(ctx context.Context, l []resourcePrivateAccessServiceConnectionBackupLinksModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandPrivateAccessServiceConnectionBackupLinks(ctx, diags)
	}
	return result
}

func (data *resourcePrivateAccessServiceConnectionConfigModel) expandPrivateAccessServiceConnectionConfig(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})

	return result
}

func (data *resourcePrivateAccessServiceConnectionConfigBackupLinksModel) expandPrivateAccessServiceConnectionConfigBackupLinks(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})

	return result
}

func (s *resourcePrivateAccessServiceConnectionConfigModel) expandPrivateAccessServiceConnectionConfigBackupLinksList(ctx context.Context, l []resourcePrivateAccessServiceConnectionConfigBackupLinksModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandPrivateAccessServiceConnectionConfigBackupLinks(ctx, diags)
	}
	return result
}

func (data *resourcePrivateAccessServiceConnectionCommonConfigModel) expandPrivateAccessServiceConnectionCommonConfig(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})

	return result
}

func (data *resourcePrivateAccessServiceConnectionIpAssignedModel) expandPrivateAccessServiceConnectionIpAssigned(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})

	return result
}

func (s *resourcePrivateAccessServiceConnectionModel) expandPrivateAccessServiceConnectionIpAssignedList(ctx context.Context, l []resourcePrivateAccessServiceConnectionIpAssignedModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandPrivateAccessServiceConnectionIpAssigned(ctx, diags)
	}
	return result
}
