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
	"time"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourcePrivateAccessServiceConnections{}

func newResourcePrivateAccessServiceConnections() resource.Resource {
	return &resourcePrivateAccessServiceConnections{}
}

type resourcePrivateAccessServiceConnections struct {
	fortiClient *FortiClient
}

// resourcePrivateAccessServiceConnectionsModel describes the resource data model.
type resourcePrivateAccessServiceConnectionsModel struct {
	ID                  types.String                                              `tfsdk:"id"`
	Alias               types.String                                              `tfsdk:"alias"`
	BgpPeerIp           types.String                                              `tfsdk:"bgp_peer_ip"`
	IpsecRemoteGw       types.String                                              `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId    types.String                                              `tfsdk:"overlay_network_id"`
	RouteMapTag         types.String                                              `tfsdk:"route_map_tag"`
	Auth                types.String                                              `tfsdk:"auth"`
	IpsecPreSharedKey   types.String                                              `tfsdk:"ipsec_pre_shared_key"`
	IpsecCertName       types.String                                              `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion     types.String                                              `tfsdk:"ipsec_ike_version"`
	IpsecPeerName       types.String                                              `tfsdk:"ipsec_peer_name"`
	BackupLinks         []resourcePrivateAccessServiceConnectionsBackupLinksModel `tfsdk:"backup_links"`
	Ftntid              types.String                                              `tfsdk:"ftntid"`
	Type                types.String                                              `tfsdk:"type"`
	ConfigState         types.String                                              `tfsdk:"config_state"`
	SeqNum              types.Float64                                             `tfsdk:"seq_num"`
	FailedMessage       types.String                                              `tfsdk:"failed_message"`
	Config              *resourcePrivateAccessServiceConnectionsConfigModel       `tfsdk:"config"`
	CommonConfig        *resourcePrivateAccessServiceConnectionsCommonConfigModel `tfsdk:"common_config"`
	IpAssigned          []resourcePrivateAccessServiceConnectionsIpAssignedModel  `tfsdk:"ip_assigned"`
	RegionCost          *resourcePrivateAccessServiceConnectionsRegionCostModel   `tfsdk:"region_cost"`
	ServiceConnectionId types.String                                              `tfsdk:"service_connection_id"`
}

func (r *resourcePrivateAccessServiceConnections) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_service_connections"
}

func (r *resourcePrivateAccessServiceConnections) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"alias": schema.StringAttribute{
				Description: "alias for serivce connection",
				Optional:    true,
			},
			"bgp_peer_ip": schema.StringAttribute{
				Description: "BGP Routing Peer IP.",
				Optional:    true,
			},
			"ipsec_remote_gw": schema.StringAttribute{
				Description: "IPSEC Remote Gateway IP",
				Optional:    true,
			},
			"overlay_network_id": schema.StringAttribute{
				Description: "integer id for overlay",
				Optional:    true,
			},
			"route_map_tag": schema.StringAttribute{
				Description: "route map tag",
				Optional:    true,
			},
			"auth": schema.StringAttribute{
				Description: "IPSEC authentication method",
				Validators: []validator.String{
					stringvalidator.OneOf("pki", "psk"),
				},
				Optional: true,
			},
			"ipsec_pre_shared_key": schema.StringAttribute{
				Description: "IPSEC auth by pre shared key.",
				Optional:    true,
			},
			"ipsec_cert_name": schema.StringAttribute{
				Description: "the name of IPSEC authentication certificate that uploaded to SASE",
				Optional:    true,
			},
			"ipsec_ike_version": schema.StringAttribute{
				Description: "IKE version for IPSEC",
				Validators: []validator.String{
					stringvalidator.OneOf("2"),
				},
				Optional: true,
			},
			"ipsec_peer_name": schema.StringAttribute{
				Description: "Peer PKI user name that created on SASE for IPSEC authentication",
				Optional:    true,
			},
			"ftntid": schema.StringAttribute{
				Description: "unique id for service connection",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "BGP Routing Design. Must be same as network configuration.",
				Validators: []validator.String{
					stringvalidator.OneOf("overlay", "loopback"),
				},
				Computed: true,
				Optional: true,
			},
			"config_state": schema.StringAttribute{
				Description: "Configuration state of service connection.",
				Validators: []validator.String{
					stringvalidator.OneOf("success", "failed", "creating", "updating", "deleting"),
				},
				Computed: true,
			},
			"seq_num": schema.Float64Attribute{
				Description: "sequential unique number for service connection",
				Computed:    true,
			},
			"failed_message": schema.StringAttribute{
				Description: "failure message while config service connection",
				Computed:    true,
			},
			"service_connection_id": schema.StringAttribute{
				Description: "the unique uuid for service connection",
				Computed:    true,
			},
			"backup_links": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"delete": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "unique id for additional IPsec overlays.",
										Computed:    true,
										Optional:    true,
									},
								},
							},
							Optional: true,
						},
						"update": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "unique id for additional IPsec overlays.",
										Computed:    true,
										Optional:    true,
									},
									"alias": schema.StringAttribute{
										Description: "alias for serivce connection additional overlay",
										Computed:    true,
										Optional:    true,
									},
									"auth": schema.StringAttribute{
										Description: "IPSEC authentication method",
										Validators: []validator.String{
											stringvalidator.OneOf("pki", "psk"),
										},
										Computed: true,
										Optional: true,
									},
									"ipsec_cert_name": schema.StringAttribute{
										Description: "the name of IPSEC authentication certificate that uploaded to SASE",
										Computed:    true,
										Optional:    true,
									},
									"ipsec_ike_version": schema.StringAttribute{
										Description: "IKE version for IPSEC",
										Validators: []validator.String{
											stringvalidator.OneOf("2"),
										},
										Computed: true,
										Optional: true,
									},
									"ipsec_peer_name": schema.StringAttribute{
										Description: "Peer PKI user name that created on SASE for IPSEC authentication",
										Computed:    true,
										Optional:    true,
									},
									"ipsec_remote_gw": schema.StringAttribute{
										Description: "IPSEC Remote Gateway IP",
										Computed:    true,
										Optional:    true,
									},
									"overlay_network_id": schema.StringAttribute{
										Description: "integer id for overlay",
										Computed:    true,
										Optional:    true,
									},
									"ipsec_pre_shared_key": schema.StringAttribute{
										Description: "IPSEC auth by pre shared key.",
										Computed:    true,
										Optional:    true,
									},
								},
							},
							Optional: true,
						},
						"create": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"alias": schema.StringAttribute{
										Description: "alias for serivce connection additional overlay",
										Computed:    true,
										Optional:    true,
									},
									"auth": schema.StringAttribute{
										Description: "IPSEC authentication method",
										Validators: []validator.String{
											stringvalidator.OneOf("pki", "psk"),
										},
										Computed: true,
										Optional: true,
									},
									"ipsec_cert_name": schema.StringAttribute{
										Description: "the name of IPSEC authentication certificate that uploaded to SASE",
										Computed:    true,
										Optional:    true,
									},
									"ipsec_ike_version": schema.StringAttribute{
										Description: "IKE version for IPSEC",
										Validators: []validator.String{
											stringvalidator.OneOf("2"),
										},
										Computed: true,
										Optional: true,
									},
									"ipsec_peer_name": schema.StringAttribute{
										Description: "Peer PKI user name that created on SASE for IPSEC authentication",
										Computed:    true,
										Optional:    true,
									},
									"ipsec_remote_gw": schema.StringAttribute{
										Description: "IPSEC Remote Gateway IP",
										Computed:    true,
										Optional:    true,
									},
									"overlay_network_id": schema.StringAttribute{
										Description: "integer id for overlay",
										Computed:    true,
										Optional:    true,
									},
									"ipsec_pre_shared_key": schema.StringAttribute{
										Description: "IPSEC auth by pre shared key.",
										Computed:    true,
										Optional:    true,
									},
								},
							},
							Optional: true,
						},
					},
				},
				Optional: true,
			},
			"config": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"alias": schema.StringAttribute{
						Description: "alias for serivce connection",
						Computed:    true,
						Optional:    true,
					},
					"auth": schema.StringAttribute{
						Description: "IPSEC authentication method",
						Validators: []validator.String{
							stringvalidator.OneOf("pki", "psk"),
						},
						Computed: true,
						Optional: true,
					},
					"bgp_peer_ip": schema.StringAttribute{
						Description: "BGP Routing Peer IP.",
						Computed:    true,
						Optional:    true,
					},
					"ipsec_cert_name": schema.StringAttribute{
						Description: "the name of IPSEC authentication certificate that uploaded to SASE",
						Computed:    true,
						Optional:    true,
					},
					"ipsec_ike_version": schema.StringAttribute{
						Description: "IKE version for IPSEC",
						Validators: []validator.String{
							stringvalidator.OneOf("2"),
						},
						Computed: true,
						Optional: true,
					},
					"ipsec_peer_name": schema.StringAttribute{
						Description: "Peer PKI user name that created on SASE for IPSEC authentication",
						Computed:    true,
						Optional:    true,
					},
					"ipsec_remote_gw": schema.StringAttribute{
						Description: "IPSEC Remote Gateway IP",
						Computed:    true,
						Optional:    true,
					},
					"overlay_network_id": schema.StringAttribute{
						Description: "integer id for overlay",
						Computed:    true,
						Optional:    true,
					},
					"route_map_tag": schema.StringAttribute{
						Description: "route map tag",
						Computed:    true,
						Optional:    true,
					},
					"region_cost": schema.SingleNestedAttribute{
						Description: "cost value to determine the priority of SASE spokes",
						Attributes: map[string]schema.Attribute{
							"sjc_f1": schema.Float64Attribute{
								Computed: true,
								Optional: true,
							},
							"lon_f1": schema.Float64Attribute{
								Computed: true,
								Optional: true,
							},
							"fra_f1": schema.Float64Attribute{
								Computed: true,
								Optional: true,
							},
							"iad_f1": schema.Float64Attribute{
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
					"backup_links": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "unique id for additional IPsec overlays.",
									Computed:    true,
									Optional:    true,
								},
								"alias": schema.StringAttribute{
									Description: "alias for serivce connection additional overlay",
									Computed:    true,
									Optional:    true,
								},
								"auth": schema.StringAttribute{
									Description: "IPSEC authentication method",
									Validators: []validator.String{
										stringvalidator.OneOf("pki", "psk"),
									},
									Computed: true,
									Optional: true,
								},
								"ipsec_cert_name": schema.StringAttribute{
									Description: "the name of IPSEC authentication certificate that uploaded to SASE",
									Computed:    true,
									Optional:    true,
								},
								"ipsec_ike_version": schema.StringAttribute{
									Description: "IKE version for IPSEC",
									Validators: []validator.String{
										stringvalidator.OneOf("2"),
									},
									Computed: true,
									Optional: true,
								},
								"ipsec_peer_name": schema.StringAttribute{
									Description: "Peer PKI user name that created on SASE for IPSEC authentication",
									Computed:    true,
									Optional:    true,
								},
								"ipsec_remote_gw": schema.StringAttribute{
									Description: "IPSEC Remote Gateway IP",
									Computed:    true,
									Optional:    true,
								},
								"overlay_network_id": schema.StringAttribute{
									Description: "integer id for overlay",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
			},
			"common_config": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"config_state": schema.StringAttribute{
						Description: "Configuration state of network configuration.",
						Validators: []validator.String{
							stringvalidator.OneOf("success", "failed", "creating", "updating", "deleting"),
						},
						Computed: true,
						Optional: true,
					},
					"bgp_design": schema.StringAttribute{
						Description: "BGP Routing Design.",
						Validators: []validator.String{
							stringvalidator.OneOf("overlay", "loopback"),
						},
						Computed: true,
						Optional: true,
					},
					"bgp_router_ids_subnet": schema.StringAttribute{
						Description: "Available/unused subnet that can be used to assign loopback interface IP addresses used for BGP router IDs parameter on the FortiSASE security PoPs. /28 is the minimum subnet size.",
						Computed:    true,
						Optional:    true,
					},
					"as_number": schema.StringAttribute{
						Description: "Autonomous System Number (ASN).",
						Computed:    true,
						Optional:    true,
					},
					"recursive_next_hop": schema.BoolAttribute{
						Description: "BGP Recursive Routing. Enabling this setting allows for interhub connectivity. When use BGP design on-loopback this has to be enabled.",
						Computed:    true,
						Optional:    true,
					},
					"sdwan_rule_enable": schema.BoolAttribute{
						Description: "Hub Selection Method. Enabling this setting the highest priority service connection that meets minimum SLA requirements is selected. Otherwise BGP MED (Multi-Exit Discriminator) will be used",
						Computed:    true,
						Optional:    true,
					},
					"sdwan_health_check_vm": schema.StringAttribute{
						Description: "Health Check IP. Must be provided when enable sdwan rule which used to obtain Jitter, latency and packet loss measurements.",
						Computed:    true,
						Optional:    true,
					},
				},
				Computed: true,
			},
			"ip_assigned": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "unique id for bgp router id assignment",
							Computed:    true,
							Optional:    true,
						},
						"sdwan_common_id": schema.StringAttribute{
							Description: "unique id related to network configuration",
							Computed:    true,
							Optional:    true,
						},
						"bgp_router_id": schema.StringAttribute{
							Description: "BGP Router ID generated from Router ID Subnets",
							Computed:    true,
							Optional:    true,
						},
						"site_id": schema.StringAttribute{
							Description: "id for SASE spoke",
							Computed:    true,
							Optional:    true,
						},
						"region": schema.StringAttribute{
							Description: "air port code for SASE spoke physical region",
							Computed:    true,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
			"region_cost": schema.SingleNestedAttribute{
				Description: "Cost value to determine the priority of SASE spokes. Default cost is 5 if not provided through initial api request.",
				Attributes: map[string]schema.Attribute{
					"sjc_f1": schema.Float64Attribute{
						Computed: true,
						Optional: true,
					},
					"lon_f1": schema.Float64Attribute{
						Computed: true,
						Optional: true,
					},
					"fra_f1": schema.Float64Attribute{
						Computed: true,
						Optional: true,
					},
					"iad_f1": schema.Float64Attribute{
						Computed: true,
						Optional: true,
					},
				},
				Optional: true,
			},
		},
	}
}

func (r *resourcePrivateAccessServiceConnections) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourcePrivateAccessServiceConnections) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourcePrivateAccessServiceConnectionsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectPrivateAccessServiceConnections(ctx, diags))
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnections(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreatePrivateAccessServiceConnections(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["id"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnections(ctx, "read", diags))

	read_output := make(map[string]interface{})
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadPrivateAccessServiceConnections(&read_input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to read resource: %v", err),
				"",
			)
			return
		}
		if v, ok := read_output["config_state"]; ok {
			if v != "success" {
				continue
			}
		}
		break
	}

	diags.Append(data.refreshPrivateAccessServiceConnections(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnections) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourcePrivateAccessServiceConnectionsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourcePrivateAccessServiceConnectionsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectPrivateAccessServiceConnections(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnections(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdatePrivateAccessServiceConnections(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnections(ctx, "read", diags))

	read_output := make(map[string]interface{})
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadPrivateAccessServiceConnections(&read_input_model)
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Error to read resource: %v", err),
				"",
			)
			return
		}
		if v, ok := read_output["config_state"]; ok {
			if v != "success" {
				continue
			}
		}
		break
	}

	diags.Append(data.refreshPrivateAccessServiceConnections(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnections) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	diags := &resp.Diagnostics
	var data resourcePrivateAccessServiceConnectionsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnections(ctx, "delete", diags))

	err := c.DeletePrivateAccessServiceConnections(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource: %v", err),
			"",
		)
		return
	}
	read_output := make(map[string]interface{})
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		read_output, err = c.ReadPrivateAccessServiceConnections(&input_model)
		if err != nil || len(read_output) == 0 {
			// Delete success
			return
		}
	}
	diags.AddError(
		fmt.Sprintf("Error to delete resource: %v", err),
		fmt.Sprintf("The resource still exists: %v", read_output),
	)
}

func (r *resourcePrivateAccessServiceConnections) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourcePrivateAccessServiceConnectionsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnections(ctx, "read", diags))

	read_output, err := c.ReadPrivateAccessServiceConnections(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshPrivateAccessServiceConnections(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnections) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourcePrivateAccessServiceConnectionsModel) refreshPrivateAccessServiceConnections(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["alias"]; ok {
		m.Alias = parseStringValue(v)
	}

	if v, ok := o["bgp_peer_ip"]; ok {
		m.BgpPeerIp = parseStringValue(v)
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

	if v, ok := o["auth"]; ok {
		m.Auth = parseStringValue(v)
	}

	if v, ok := o["ipsec_pre_shared_key"]; ok {
		m.IpsecPreSharedKey = parseStringValue(v)
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

	if v, ok := o["backup_links"]; ok {
		m.BackupLinks = m.flattenPrivateAccessServiceConnectionsBackupLinksList(ctx, v, &diags)
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
		m.Config = m.Config.flattenPrivateAccessServiceConnectionsConfig(ctx, v, &diags)
	}

	if v, ok := o["common_config"]; ok {
		m.CommonConfig = m.CommonConfig.flattenPrivateAccessServiceConnectionsCommonConfig(ctx, v, &diags)
	}

	if v, ok := o["ip_assigned"]; ok {
		m.IpAssigned = m.flattenPrivateAccessServiceConnectionsIpAssignedList(ctx, v, &diags)
	}

	if v, ok := o["region_cost"]; ok {
		m.RegionCost = m.RegionCost.flattenPrivateAccessServiceConnectionsRegionCost(ctx, v, &diags)
	}

	return diags
}

func (data *resourcePrivateAccessServiceConnectionsModel) getCreateObjectPrivateAccessServiceConnections(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Alias.IsNull() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.BgpPeerIp.IsNull() {
		result["bgp_peer_ip"] = data.BgpPeerIp.ValueString()
	}

	if !data.IpsecRemoteGw.IsNull() {
		result["ipsec_remote_gw"] = data.IpsecRemoteGw.ValueString()
	}

	if !data.OverlayNetworkId.IsNull() {
		result["overlay_network_id"] = data.OverlayNetworkId.ValueString()
	}

	if !data.RouteMapTag.IsNull() {
		result["route_map_tag"] = data.RouteMapTag.ValueString()
	}

	if !data.Auth.IsNull() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	if !data.IpsecCertName.IsNull() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	if !data.IpsecIkeVersion.IsNull() {
		result["ipsec_ike_version"] = data.IpsecIkeVersion.ValueString()
	}

	if !data.IpsecPeerName.IsNull() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if len(data.BackupLinks) > 0 {
		result["backup_links"] = data.expandPrivateAccessServiceConnectionsBackupLinksList(ctx, data.BackupLinks, diags)
	}

	if !data.Type.IsNull() {
		result["type"] = data.Type.ValueString()
	}

	if data.RegionCost != nil && !isZeroStruct(*data.RegionCost) {
		result["region_cost"] = data.RegionCost.expandPrivateAccessServiceConnectionsRegionCost(ctx, diags)
	}

	return &result
}

func (data *resourcePrivateAccessServiceConnectionsModel) getUpdateObjectPrivateAccessServiceConnections(ctx context.Context, state resourcePrivateAccessServiceConnectionsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Alias.IsNull() && !data.Alias.Equal(state.Alias) {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.BgpPeerIp.IsNull() {
		result["bgp_peer_ip"] = data.BgpPeerIp.ValueString()
	}

	if !data.IpsecRemoteGw.IsNull() {
		result["ipsec_remote_gw"] = data.IpsecRemoteGw.ValueString()
	}

	if !data.OverlayNetworkId.IsNull() {
		result["overlay_network_id"] = data.OverlayNetworkId.ValueString()
	}

	if !data.RouteMapTag.IsNull() && !data.RouteMapTag.Equal(state.RouteMapTag) {
		result["route_map_tag"] = data.RouteMapTag.ValueString()
	}

	if !data.Auth.IsNull() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() && !data.IpsecPreSharedKey.Equal(state.IpsecPreSharedKey) {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	if !data.IpsecCertName.IsNull() && !data.IpsecCertName.Equal(state.IpsecCertName) {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	if !data.IpsecIkeVersion.IsNull() {
		result["ipsec_ike_version"] = data.IpsecIkeVersion.ValueString()
	}

	if !data.IpsecPeerName.IsNull() && !data.IpsecPeerName.Equal(state.IpsecPeerName) {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if len(data.BackupLinks) > 0 || !isSameStruct(data.BackupLinks, state.BackupLinks) {
		result["backup_links"] = data.expandPrivateAccessServiceConnectionsBackupLinksList(ctx, data.BackupLinks, diags)
	}

	return &result
}

func (data *resourcePrivateAccessServiceConnectionsModel) getURLObjectPrivateAccessServiceConnections(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ServiceConnectionId.IsNull() {
		result["service-connection-id"] = data.ServiceConnectionId.ValueString()
	}

	return &result
}

type resourcePrivateAccessServiceConnectionsBackupLinksModel struct {
	Delete []resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel `tfsdk:"delete"`
	Update []resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel `tfsdk:"update"`
	Create []resourcePrivateAccessServiceConnectionsBackupLinksCreateModel `tfsdk:"create"`
}

type resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel struct {
	Id types.String `tfsdk:"id"`
}

type resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel struct {
	Id                types.String `tfsdk:"id"`
	Alias             types.String `tfsdk:"alias"`
	Auth              types.String `tfsdk:"auth"`
	IpsecCertName     types.String `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion   types.String `tfsdk:"ipsec_ike_version"`
	IpsecPeerName     types.String `tfsdk:"ipsec_peer_name"`
	IpsecRemoteGw     types.String `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId  types.String `tfsdk:"overlay_network_id"`
	IpsecPreSharedKey types.String `tfsdk:"ipsec_pre_shared_key"`
}

type resourcePrivateAccessServiceConnectionsBackupLinksCreateModel struct {
	Alias             types.String `tfsdk:"alias"`
	Auth              types.String `tfsdk:"auth"`
	IpsecCertName     types.String `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion   types.String `tfsdk:"ipsec_ike_version"`
	IpsecPeerName     types.String `tfsdk:"ipsec_peer_name"`
	IpsecRemoteGw     types.String `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId  types.String `tfsdk:"overlay_network_id"`
	IpsecPreSharedKey types.String `tfsdk:"ipsec_pre_shared_key"`
}

type resourcePrivateAccessServiceConnectionsConfigModel struct {
	Alias            types.String                                                    `tfsdk:"alias"`
	Auth             types.String                                                    `tfsdk:"auth"`
	BgpPeerIp        types.String                                                    `tfsdk:"bgp_peer_ip"`
	IpsecCertName    types.String                                                    `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion  types.String                                                    `tfsdk:"ipsec_ike_version"`
	IpsecPeerName    types.String                                                    `tfsdk:"ipsec_peer_name"`
	IpsecRemoteGw    types.String                                                    `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId types.String                                                    `tfsdk:"overlay_network_id"`
	RouteMapTag      types.String                                                    `tfsdk:"route_map_tag"`
	RegionCost       *resourcePrivateAccessServiceConnectionsConfigRegionCostModel   `tfsdk:"region_cost"`
	BackupLinks      []resourcePrivateAccessServiceConnectionsConfigBackupLinksModel `tfsdk:"backup_links"`
}

type resourcePrivateAccessServiceConnectionsConfigRegionCostModel struct {
	SjcF1 types.Float64 `tfsdk:"sjc_f1"`
	LonF1 types.Float64 `tfsdk:"lon_f1"`
	FraF1 types.Float64 `tfsdk:"fra_f1"`
	IadF1 types.Float64 `tfsdk:"iad_f1"`
}

type resourcePrivateAccessServiceConnectionsConfigBackupLinksModel struct {
	Id               types.String `tfsdk:"id"`
	Alias            types.String `tfsdk:"alias"`
	Auth             types.String `tfsdk:"auth"`
	IpsecCertName    types.String `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion  types.String `tfsdk:"ipsec_ike_version"`
	IpsecPeerName    types.String `tfsdk:"ipsec_peer_name"`
	IpsecRemoteGw    types.String `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId types.String `tfsdk:"overlay_network_id"`
}

type resourcePrivateAccessServiceConnectionsCommonConfigModel struct {
	ConfigState        types.String `tfsdk:"config_state"`
	BgpDesign          types.String `tfsdk:"bgp_design"`
	BgpRouterIdsSubnet types.String `tfsdk:"bgp_router_ids_subnet"`
	AsNumber           types.String `tfsdk:"as_number"`
	RecursiveNextHop   types.Bool   `tfsdk:"recursive_next_hop"`
	SdwanRuleEnable    types.Bool   `tfsdk:"sdwan_rule_enable"`
	SdwanHealthCheckVm types.String `tfsdk:"sdwan_health_check_vm"`
}

type resourcePrivateAccessServiceConnectionsIpAssignedModel struct {
	Id            types.String `tfsdk:"id"`
	SdwanCommonId types.String `tfsdk:"sdwan_common_id"`
	BgpRouterId   types.String `tfsdk:"bgp_router_id"`
	SiteId        types.String `tfsdk:"site_id"`
	Region        types.String `tfsdk:"region"`
}

type resourcePrivateAccessServiceConnectionsRegionCostModel struct {
	SjcF1 types.Float64 `tfsdk:"sjc_f1"`
	LonF1 types.Float64 `tfsdk:"lon_f1"`
	FraF1 types.Float64 `tfsdk:"fra_f1"`
	IadF1 types.Float64 `tfsdk:"iad_f1"`
}

func (m *resourcePrivateAccessServiceConnectionsBackupLinksModel) flattenPrivateAccessServiceConnectionsBackupLinks(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionsBackupLinksModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionsBackupLinksModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionsBackupLinksModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["delete"]; ok {
		m.Delete = m.flattenPrivateAccessServiceConnectionsBackupLinksDeleteList(ctx, v, diags)
	}

	if v, ok := o["update"]; ok {
		m.Update = m.flattenPrivateAccessServiceConnectionsBackupLinksUpdateList(ctx, v, diags)
	}

	if v, ok := o["create"]; ok {
		m.Create = m.flattenPrivateAccessServiceConnectionsBackupLinksCreateList(ctx, v, diags)
	}

	return m
}

func (s *resourcePrivateAccessServiceConnectionsModel) flattenPrivateAccessServiceConnectionsBackupLinksList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourcePrivateAccessServiceConnectionsBackupLinksModel {
	if o == nil {
		return []resourcePrivateAccessServiceConnectionsBackupLinksModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument backup_links is not type of []interface{}.", "")
		return []resourcePrivateAccessServiceConnectionsBackupLinksModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourcePrivateAccessServiceConnectionsBackupLinksModel{}
	}

	values := make([]resourcePrivateAccessServiceConnectionsBackupLinksModel, len(l))
	for i, ele := range l {
		var m resourcePrivateAccessServiceConnectionsBackupLinksModel
		values[i] = *m.flattenPrivateAccessServiceConnectionsBackupLinks(ctx, ele, diags)
	}

	return values
}

func (m *resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel) flattenPrivateAccessServiceConnectionsBackupLinksDelete(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["id"]; ok {
		m.Id = parseStringValue(v)
	}

	return m
}

func (s *resourcePrivateAccessServiceConnectionsBackupLinksModel) flattenPrivateAccessServiceConnectionsBackupLinksDeleteList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel {
	if o == nil {
		return []resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument delete is not type of []interface{}.", "")
		return []resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel{}
	}

	values := make([]resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel, len(l))
	for i, ele := range l {
		var m resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel
		values[i] = *m.flattenPrivateAccessServiceConnectionsBackupLinksDelete(ctx, ele, diags)
	}

	return values
}

func (m *resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel) flattenPrivateAccessServiceConnectionsBackupLinksUpdate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel{}
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

	if v, ok := o["ipsec_pre_shared_key"]; ok {
		m.IpsecPreSharedKey = parseStringValue(v)
	}

	return m
}

func (s *resourcePrivateAccessServiceConnectionsBackupLinksModel) flattenPrivateAccessServiceConnectionsBackupLinksUpdateList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel {
	if o == nil {
		return []resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument update is not type of []interface{}.", "")
		return []resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel{}
	}

	values := make([]resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel, len(l))
	for i, ele := range l {
		var m resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel
		values[i] = *m.flattenPrivateAccessServiceConnectionsBackupLinksUpdate(ctx, ele, diags)
	}

	return values
}

func (m *resourcePrivateAccessServiceConnectionsBackupLinksCreateModel) flattenPrivateAccessServiceConnectionsBackupLinksCreate(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionsBackupLinksCreateModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionsBackupLinksCreateModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionsBackupLinksCreateModel{}
	}
	o := input.(map[string]interface{})
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

	if v, ok := o["ipsec_pre_shared_key"]; ok {
		m.IpsecPreSharedKey = parseStringValue(v)
	}

	return m
}

func (s *resourcePrivateAccessServiceConnectionsBackupLinksModel) flattenPrivateAccessServiceConnectionsBackupLinksCreateList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourcePrivateAccessServiceConnectionsBackupLinksCreateModel {
	if o == nil {
		return []resourcePrivateAccessServiceConnectionsBackupLinksCreateModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument create is not type of []interface{}.", "")
		return []resourcePrivateAccessServiceConnectionsBackupLinksCreateModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourcePrivateAccessServiceConnectionsBackupLinksCreateModel{}
	}

	values := make([]resourcePrivateAccessServiceConnectionsBackupLinksCreateModel, len(l))
	for i, ele := range l {
		var m resourcePrivateAccessServiceConnectionsBackupLinksCreateModel
		values[i] = *m.flattenPrivateAccessServiceConnectionsBackupLinksCreate(ctx, ele, diags)
	}

	return values
}

func (m *resourcePrivateAccessServiceConnectionsConfigModel) flattenPrivateAccessServiceConnectionsConfig(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionsConfigModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionsConfigModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionsConfigModel{}
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
		m.RegionCost = m.RegionCost.flattenPrivateAccessServiceConnectionsConfigRegionCost(ctx, v, diags)
	}

	if v, ok := o["backup_links"]; ok {
		m.BackupLinks = m.flattenPrivateAccessServiceConnectionsConfigBackupLinksList(ctx, v, diags)
	}

	return m
}

func (m *resourcePrivateAccessServiceConnectionsConfigRegionCostModel) flattenPrivateAccessServiceConnectionsConfigRegionCost(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionsConfigRegionCostModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionsConfigRegionCostModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionsConfigRegionCostModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["sjc-f1"]; ok {
		m.SjcF1 = parseFloat64Value(v)
	}

	if v, ok := o["lon-f1"]; ok {
		m.LonF1 = parseFloat64Value(v)
	}

	if v, ok := o["fra-f1"]; ok {
		m.FraF1 = parseFloat64Value(v)
	}

	if v, ok := o["iad-f1"]; ok {
		m.IadF1 = parseFloat64Value(v)
	}

	return m
}

func (m *resourcePrivateAccessServiceConnectionsConfigBackupLinksModel) flattenPrivateAccessServiceConnectionsConfigBackupLinks(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionsConfigBackupLinksModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionsConfigBackupLinksModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionsConfigBackupLinksModel{}
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

func (s *resourcePrivateAccessServiceConnectionsConfigModel) flattenPrivateAccessServiceConnectionsConfigBackupLinksList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourcePrivateAccessServiceConnectionsConfigBackupLinksModel {
	if o == nil {
		return []resourcePrivateAccessServiceConnectionsConfigBackupLinksModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument backup_links is not type of []interface{}.", "")
		return []resourcePrivateAccessServiceConnectionsConfigBackupLinksModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourcePrivateAccessServiceConnectionsConfigBackupLinksModel{}
	}

	values := make([]resourcePrivateAccessServiceConnectionsConfigBackupLinksModel, len(l))
	for i, ele := range l {
		var m resourcePrivateAccessServiceConnectionsConfigBackupLinksModel
		values[i] = *m.flattenPrivateAccessServiceConnectionsConfigBackupLinks(ctx, ele, diags)
	}

	return values
}

func (m *resourcePrivateAccessServiceConnectionsCommonConfigModel) flattenPrivateAccessServiceConnectionsCommonConfig(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionsCommonConfigModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionsCommonConfigModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionsCommonConfigModel{}
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

func (m *resourcePrivateAccessServiceConnectionsIpAssignedModel) flattenPrivateAccessServiceConnectionsIpAssigned(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionsIpAssignedModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionsIpAssignedModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionsIpAssignedModel{}
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

func (s *resourcePrivateAccessServiceConnectionsModel) flattenPrivateAccessServiceConnectionsIpAssignedList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourcePrivateAccessServiceConnectionsIpAssignedModel {
	if o == nil {
		return []resourcePrivateAccessServiceConnectionsIpAssignedModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument ip_assigned is not type of []interface{}.", "")
		return []resourcePrivateAccessServiceConnectionsIpAssignedModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourcePrivateAccessServiceConnectionsIpAssignedModel{}
	}

	values := make([]resourcePrivateAccessServiceConnectionsIpAssignedModel, len(l))
	for i, ele := range l {
		var m resourcePrivateAccessServiceConnectionsIpAssignedModel
		values[i] = *m.flattenPrivateAccessServiceConnectionsIpAssigned(ctx, ele, diags)
	}

	return values
}

func (m *resourcePrivateAccessServiceConnectionsRegionCostModel) flattenPrivateAccessServiceConnectionsRegionCost(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourcePrivateAccessServiceConnectionsRegionCostModel {
	if input == nil {
		return &resourcePrivateAccessServiceConnectionsRegionCostModel{}
	}
	if m == nil {
		m = &resourcePrivateAccessServiceConnectionsRegionCostModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["sjc-f1"]; ok {
		m.SjcF1 = parseFloat64Value(v)
	}

	if v, ok := o["lon-f1"]; ok {
		m.LonF1 = parseFloat64Value(v)
	}

	if v, ok := o["fra-f1"]; ok {
		m.FraF1 = parseFloat64Value(v)
	}

	if v, ok := o["iad-f1"]; ok {
		m.IadF1 = parseFloat64Value(v)
	}

	return m
}

func (data *resourcePrivateAccessServiceConnectionsBackupLinksModel) expandPrivateAccessServiceConnectionsBackupLinks(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if len(data.Delete) > 0 {
		result["delete"] = data.expandPrivateAccessServiceConnectionsBackupLinksDeleteList(ctx, data.Delete, diags)
	}

	if len(data.Update) > 0 {
		result["update"] = data.expandPrivateAccessServiceConnectionsBackupLinksUpdateList(ctx, data.Update, diags)
	}

	if len(data.Create) > 0 {
		result["create"] = data.expandPrivateAccessServiceConnectionsBackupLinksCreateList(ctx, data.Create, diags)
	}

	return result
}

func (s *resourcePrivateAccessServiceConnectionsModel) expandPrivateAccessServiceConnectionsBackupLinksList(ctx context.Context, l []resourcePrivateAccessServiceConnectionsBackupLinksModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandPrivateAccessServiceConnectionsBackupLinks(ctx, diags)
	}
	return result
}

func (data *resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel) expandPrivateAccessServiceConnectionsBackupLinksDelete(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() {
		result["id"] = data.Id.ValueString()
	}

	return result
}

func (s *resourcePrivateAccessServiceConnectionsBackupLinksModel) expandPrivateAccessServiceConnectionsBackupLinksDeleteList(ctx context.Context, l []resourcePrivateAccessServiceConnectionsBackupLinksDeleteModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandPrivateAccessServiceConnectionsBackupLinksDelete(ctx, diags)
	}
	return result
}

func (data *resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel) expandPrivateAccessServiceConnectionsBackupLinksUpdate(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() {
		result["id"] = data.Id.ValueString()
	}

	if !data.Alias.IsNull() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.Auth.IsNull() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecCertName.IsNull() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	if !data.IpsecIkeVersion.IsNull() {
		result["ipsec_ike_version"] = data.IpsecIkeVersion.ValueString()
	}

	if !data.IpsecPeerName.IsNull() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if !data.IpsecRemoteGw.IsNull() {
		result["ipsec_remote_gw"] = data.IpsecRemoteGw.ValueString()
	}

	if !data.OverlayNetworkId.IsNull() {
		result["overlay_network_id"] = data.OverlayNetworkId.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	return result
}

func (s *resourcePrivateAccessServiceConnectionsBackupLinksModel) expandPrivateAccessServiceConnectionsBackupLinksUpdateList(ctx context.Context, l []resourcePrivateAccessServiceConnectionsBackupLinksUpdateModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandPrivateAccessServiceConnectionsBackupLinksUpdate(ctx, diags)
	}
	return result
}

func (data *resourcePrivateAccessServiceConnectionsBackupLinksCreateModel) expandPrivateAccessServiceConnectionsBackupLinksCreate(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Alias.IsNull() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.Auth.IsNull() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecCertName.IsNull() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	if !data.IpsecIkeVersion.IsNull() {
		result["ipsec_ike_version"] = data.IpsecIkeVersion.ValueString()
	}

	if !data.IpsecPeerName.IsNull() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if !data.IpsecRemoteGw.IsNull() {
		result["ipsec_remote_gw"] = data.IpsecRemoteGw.ValueString()
	}

	if !data.OverlayNetworkId.IsNull() {
		result["overlay_network_id"] = data.OverlayNetworkId.ValueString()
	}

	if !data.IpsecPreSharedKey.IsNull() {
		result["ipsec_pre_shared_key"] = data.IpsecPreSharedKey.ValueString()
	}

	return result
}

func (s *resourcePrivateAccessServiceConnectionsBackupLinksModel) expandPrivateAccessServiceConnectionsBackupLinksCreateList(ctx context.Context, l []resourcePrivateAccessServiceConnectionsBackupLinksCreateModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandPrivateAccessServiceConnectionsBackupLinksCreate(ctx, diags)
	}
	return result
}

func (data *resourcePrivateAccessServiceConnectionsConfigModel) expandPrivateAccessServiceConnectionsConfig(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Alias.IsNull() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.Auth.IsNull() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.BgpPeerIp.IsNull() {
		result["bgp_peer_ip"] = data.BgpPeerIp.ValueString()
	}

	if !data.IpsecCertName.IsNull() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	if !data.IpsecIkeVersion.IsNull() {
		result["ipsec_ike_version"] = data.IpsecIkeVersion.ValueString()
	}

	if !data.IpsecPeerName.IsNull() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if !data.IpsecRemoteGw.IsNull() {
		result["ipsec_remote_gw"] = data.IpsecRemoteGw.ValueString()
	}

	if !data.OverlayNetworkId.IsNull() {
		result["overlay_network_id"] = data.OverlayNetworkId.ValueString()
	}

	if !data.RouteMapTag.IsNull() {
		result["route_map_tag"] = data.RouteMapTag.ValueString()
	}

	if data.RegionCost != nil && !isZeroStruct(*data.RegionCost) {
		result["region_cost"] = data.RegionCost.expandPrivateAccessServiceConnectionsConfigRegionCost(ctx, diags)
	}

	if len(data.BackupLinks) > 0 {
		result["backup_links"] = data.expandPrivateAccessServiceConnectionsConfigBackupLinksList(ctx, data.BackupLinks, diags)
	}

	return result
}

func (data *resourcePrivateAccessServiceConnectionsConfigRegionCostModel) expandPrivateAccessServiceConnectionsConfigRegionCost(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SjcF1.IsNull() {
		result["sjc-f1"] = data.SjcF1.ValueFloat64()
	}

	if !data.LonF1.IsNull() {
		result["lon-f1"] = data.LonF1.ValueFloat64()
	}

	if !data.FraF1.IsNull() {
		result["fra-f1"] = data.FraF1.ValueFloat64()
	}

	if !data.IadF1.IsNull() {
		result["iad-f1"] = data.IadF1.ValueFloat64()
	}

	return result
}

func (data *resourcePrivateAccessServiceConnectionsConfigBackupLinksModel) expandPrivateAccessServiceConnectionsConfigBackupLinks(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() {
		result["id"] = data.Id.ValueString()
	}

	if !data.Alias.IsNull() {
		result["alias"] = data.Alias.ValueString()
	}

	if !data.Auth.IsNull() {
		result["auth"] = data.Auth.ValueString()
	}

	if !data.IpsecCertName.IsNull() {
		result["ipsec_cert_name"] = data.IpsecCertName.ValueString()
	}

	if !data.IpsecIkeVersion.IsNull() {
		result["ipsec_ike_version"] = data.IpsecIkeVersion.ValueString()
	}

	if !data.IpsecPeerName.IsNull() {
		result["ipsec_peer_name"] = data.IpsecPeerName.ValueString()
	}

	if !data.IpsecRemoteGw.IsNull() {
		result["ipsec_remote_gw"] = data.IpsecRemoteGw.ValueString()
	}

	if !data.OverlayNetworkId.IsNull() {
		result["overlay_network_id"] = data.OverlayNetworkId.ValueString()
	}

	return result
}

func (s *resourcePrivateAccessServiceConnectionsConfigModel) expandPrivateAccessServiceConnectionsConfigBackupLinksList(ctx context.Context, l []resourcePrivateAccessServiceConnectionsConfigBackupLinksModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandPrivateAccessServiceConnectionsConfigBackupLinks(ctx, diags)
	}
	return result
}

func (data *resourcePrivateAccessServiceConnectionsCommonConfigModel) expandPrivateAccessServiceConnectionsCommonConfig(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ConfigState.IsNull() {
		result["config_state"] = data.ConfigState.ValueString()
	}

	if !data.BgpDesign.IsNull() {
		result["bgp_design"] = data.BgpDesign.ValueString()
	}

	if !data.BgpRouterIdsSubnet.IsNull() {
		result["bgp_router_ids_subnet"] = data.BgpRouterIdsSubnet.ValueString()
	}

	if !data.AsNumber.IsNull() {
		result["as_number"] = data.AsNumber.ValueString()
	}

	if !data.RecursiveNextHop.IsNull() {
		result["recursive_next_hop"] = data.RecursiveNextHop.ValueBool()
	}

	if !data.SdwanRuleEnable.IsNull() {
		result["sdwan_rule_enable"] = data.SdwanRuleEnable.ValueBool()
	}

	if !data.SdwanHealthCheckVm.IsNull() {
		result["sdwan_health_check_vm"] = data.SdwanHealthCheckVm.ValueString()
	}

	return result
}

func (data *resourcePrivateAccessServiceConnectionsIpAssignedModel) expandPrivateAccessServiceConnectionsIpAssigned(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Id.IsNull() {
		result["id"] = data.Id.ValueString()
	}

	if !data.SdwanCommonId.IsNull() {
		result["sdwan_common_id"] = data.SdwanCommonId.ValueString()
	}

	if !data.BgpRouterId.IsNull() {
		result["bgp_router_id"] = data.BgpRouterId.ValueString()
	}

	if !data.SiteId.IsNull() {
		result["site_id"] = data.SiteId.ValueString()
	}

	if !data.Region.IsNull() {
		result["region"] = data.Region.ValueString()
	}

	return result
}

func (s *resourcePrivateAccessServiceConnectionsModel) expandPrivateAccessServiceConnectionsIpAssignedList(ctx context.Context, l []resourcePrivateAccessServiceConnectionsIpAssignedModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandPrivateAccessServiceConnectionsIpAssigned(ctx, diags)
	}
	return result
}

func (data *resourcePrivateAccessServiceConnectionsRegionCostModel) expandPrivateAccessServiceConnectionsRegionCost(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.SjcF1.IsNull() {
		result["sjc-f1"] = data.SjcF1.ValueFloat64()
	}

	if !data.LonF1.IsNull() {
		result["lon-f1"] = data.LonF1.ValueFloat64()
	}

	if !data.FraF1.IsNull() {
		result["fra-f1"] = data.FraF1.ValueFloat64()
	}

	if !data.IadF1.IsNull() {
		result["iad-f1"] = data.IadF1.ValueFloat64()
	}

	return result
}
