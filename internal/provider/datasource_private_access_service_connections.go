// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &datasourcePrivateAccessServiceConnections{}

func newDatasourcePrivateAccessServiceConnections() datasource.DataSource {
	return &datasourcePrivateAccessServiceConnections{}
}

type datasourcePrivateAccessServiceConnections struct {
	fortiClient  *FortiClient
	resourceName string
}

// datasourcePrivateAccessServiceConnectionsModel describes the datasource data model.
type datasourcePrivateAccessServiceConnectionsModel struct {
	Alias               types.String                                                `tfsdk:"alias"`
	BgpPeerIp           types.String                                                `tfsdk:"bgp_peer_ip"`
	IpsecRemoteGw       types.String                                                `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId    types.String                                                `tfsdk:"overlay_network_id"`
	RouteMapTag         types.String                                                `tfsdk:"route_map_tag"`
	Auth                types.String                                                `tfsdk:"auth"`
	IpsecPreSharedKey   types.String                                                `tfsdk:"ipsec_pre_shared_key"`
	IpsecCertName       types.String                                                `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion     types.String                                                `tfsdk:"ipsec_ike_version"`
	IpsecPeerName       types.String                                                `tfsdk:"ipsec_peer_name"`
	BackupLinks         []datasourcePrivateAccessServiceConnectionsBackupLinksModel `tfsdk:"backup_links"`
	Ftntid              types.String                                                `tfsdk:"ftntid"`
	Type                types.String                                                `tfsdk:"type"`
	ConfigState         types.String                                                `tfsdk:"config_state"`
	SeqNum              types.Float64                                               `tfsdk:"seq_num"`
	FailedMessage       types.String                                                `tfsdk:"failed_message"`
	Config              *datasourcePrivateAccessServiceConnectionsConfigModel       `tfsdk:"config"`
	CommonConfig        *datasourcePrivateAccessServiceConnectionsCommonConfigModel `tfsdk:"common_config"`
	IpAssigned          []datasourcePrivateAccessServiceConnectionsIpAssignedModel  `tfsdk:"ip_assigned"`
	RegionCost          types.Map                                                   `tfsdk:"region_cost"`
	ServiceConnectionId types.String                                                `tfsdk:"service_connection_id"`
}

func (r *datasourcePrivateAccessServiceConnections) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_service_connections"
}

func (r *datasourcePrivateAccessServiceConnections) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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
					stringvalidator.OneOf("pki", "psk"),
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
					stringvalidator.OneOf("2"),
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
					stringvalidator.OneOf("overlay", "loopback"),
				},
				MarkdownDescription: "BGP Routing Design. Must be same as network configuration.\nSupported values: overlay, loopback.",
				Computed:            true,
				Optional:            true,
			},
			"config_state": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("success", "failed", "creating", "updating", "deleting"),
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
				Required:            true,
			},
			"backup_links": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"alias": schema.StringAttribute{
							MarkdownDescription: "alias for serivce connection additional overlay",
							Computed:            true,
							Optional:            true,
						},
						"auth": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("pki", "psk"),
							},
							MarkdownDescription: "IPSEC authentication method.\nSupported values: pki, psk.",
							Computed:            true,
							Optional:            true,
						},
						"ipsec_cert_name": schema.StringAttribute{
							MarkdownDescription: "the name of IPSEC authentication certificate that uploaded to SASE",
							Computed:            true,
							Optional:            true,
						},
						"ipsec_ike_version": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("2"),
							},
							MarkdownDescription: "IKE version for IPSEC.\nSupported values: 2.",
							Computed:            true,
							Optional:            true,
						},
						"ipsec_peer_name": schema.StringAttribute{
							MarkdownDescription: "Peer PKI user name that created on SASE for IPSEC authentication",
							Computed:            true,
							Optional:            true,
						},
						"ipsec_remote_gw": schema.StringAttribute{
							MarkdownDescription: "IPSEC Remote Gateway IP",
							Computed:            true,
							Optional:            true,
						},
						"overlay_network_id": schema.StringAttribute{
							MarkdownDescription: "integer id for overlay",
							Computed:            true,
							Optional:            true,
						},
						"ipsec_pre_shared_key": schema.StringAttribute{
							MarkdownDescription: "IPSEC auth by pre shared key.",
							Computed:            true,
							Optional:            true,
						},
					},
				},
				Optional: true,
			},
			"config": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"alias": schema.StringAttribute{
						MarkdownDescription: "alias for serivce connection",
						Computed:            true,
						Optional:            true,
					},
					"auth": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("pki", "psk"),
						},
						MarkdownDescription: "IPSEC authentication method.\nSupported values: pki, psk.",
						Computed:            true,
						Optional:            true,
					},
					"bgp_peer_ip": schema.StringAttribute{
						MarkdownDescription: "BGP Routing Peer IP.",
						Computed:            true,
						Optional:            true,
					},
					"ipsec_cert_name": schema.StringAttribute{
						MarkdownDescription: "the name of IPSEC authentication certificate that uploaded to SASE",
						Computed:            true,
						Optional:            true,
					},
					"ipsec_ike_version": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("2"),
						},
						MarkdownDescription: "IKE version for IPSEC.\nSupported values: 2.",
						Computed:            true,
						Optional:            true,
					},
					"ipsec_peer_name": schema.StringAttribute{
						MarkdownDescription: "Peer PKI user name that created on SASE for IPSEC authentication",
						Computed:            true,
						Optional:            true,
					},
					"ipsec_remote_gw": schema.StringAttribute{
						MarkdownDescription: "IPSEC Remote Gateway IP",
						Computed:            true,
						Optional:            true,
					},
					"overlay_network_id": schema.StringAttribute{
						MarkdownDescription: "integer id for overlay",
						Computed:            true,
						Optional:            true,
					},
					"route_map_tag": schema.StringAttribute{
						MarkdownDescription: "route map tag",
						Computed:            true,
						Optional:            true,
					},
					"region_cost": schema.MapAttribute{
						MarkdownDescription: "cost value to determine the priority of SASE spokes",
						Computed:            true,
						Optional:            true,
						ElementType:         types.Int64Type,
					},
					"backup_links": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									MarkdownDescription: "unique id for additional IPsec overlays.",
									Computed:            true,
									Optional:            true,
								},
								"alias": schema.StringAttribute{
									MarkdownDescription: "alias for serivce connection additional overlay",
									Computed:            true,
									Optional:            true,
								},
								"auth": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("pki", "psk"),
									},
									MarkdownDescription: "IPSEC authentication method.\nSupported values: pki, psk.",
									Computed:            true,
									Optional:            true,
								},
								"ipsec_cert_name": schema.StringAttribute{
									MarkdownDescription: "the name of IPSEC authentication certificate that uploaded to SASE",
									Computed:            true,
									Optional:            true,
								},
								"ipsec_ike_version": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("2"),
									},
									MarkdownDescription: "IKE version for IPSEC.\nSupported values: 2.",
									Computed:            true,
									Optional:            true,
								},
								"ipsec_peer_name": schema.StringAttribute{
									MarkdownDescription: "Peer PKI user name that created on SASE for IPSEC authentication",
									Computed:            true,
									Optional:            true,
								},
								"ipsec_remote_gw": schema.StringAttribute{
									MarkdownDescription: "IPSEC Remote Gateway IP",
									Computed:            true,
									Optional:            true,
								},
								"overlay_network_id": schema.StringAttribute{
									MarkdownDescription: "integer id for overlay",
									Computed:            true,
									Optional:            true,
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
						Validators: []validator.String{
							stringvalidator.OneOf("success", "failed", "creating", "updating", "deleting"),
						},
						MarkdownDescription: "Configuration state of network configuration.\nSupported values: success, failed, creating, updating, deleting.",
						Computed:            true,
						Optional:            true,
					},
					"bgp_design": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("overlay", "loopback"),
						},
						MarkdownDescription: "BGP Routing Design.\nSupported values: overlay, loopback.",
						Computed:            true,
						Optional:            true,
					},
					"bgp_router_ids_subnet": schema.StringAttribute{
						MarkdownDescription: "Available/unused subnet that can be used to assign loopback interface IP addresses used for BGP router IDs parameter on the FortiSASE security PoPs. /28 is the minimum subnet size.",
						Computed:            true,
						Optional:            true,
					},
					"as_number": schema.StringAttribute{
						MarkdownDescription: "Autonomous System Number (ASN).",
						Computed:            true,
						Optional:            true,
					},
					"recursive_next_hop": schema.BoolAttribute{
						MarkdownDescription: "BGP Recursive Routing. Enabling this setting allows for interhub connectivity. When use BGP design on-loopback this has to be enabled.",
						Computed:            true,
						Optional:            true,
					},
					"sdwan_rule_enable": schema.BoolAttribute{
						MarkdownDescription: "Hub Selection Method. Enabling this setting the highest priority service connection that meets minimum SLA requirements is selected. Otherwise BGP MED (Multi-Exit Discriminator) will be used",
						Computed:            true,
						Optional:            true,
					},
					"sdwan_health_check_vm": schema.StringAttribute{
						MarkdownDescription: "Health Check IP. Must be provided when enable sdwan rule which used to obtain Jitter, latency and packet loss measurements.",
						Computed:            true,
						Optional:            true,
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
							Optional:            true,
						},
						"sdwan_common_id": schema.StringAttribute{
							MarkdownDescription: "unique id related to network configuration",
							Computed:            true,
							Optional:            true,
						},
						"bgp_router_id": schema.StringAttribute{
							MarkdownDescription: "BGP Router ID generated from Router ID Subnets",
							Computed:            true,
							Optional:            true,
						},
						"site_id": schema.StringAttribute{
							MarkdownDescription: "id for SASE spoke",
							Computed:            true,
							Optional:            true,
						},
						"region": schema.StringAttribute{
							MarkdownDescription: "air port code for SASE spoke physical region",
							Computed:            true,
							Optional:            true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (r *datasourcePrivateAccessServiceConnections) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	r.resourceName = "fortisase_private_access_service_connections"
}

func (r *datasourcePrivateAccessServiceConnections) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	diags := &resp.Diagnostics
	var data datasourcePrivateAccessServiceConnectionsModel

	// Read Terraform prior config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ServiceConnectionId.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectPrivateAccessServiceConnections(ctx, "read", diags))

	read_output, err := c.ReadPrivateAccessServiceConnections(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read data source %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshPrivateAccessServiceConnections(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (m *datasourcePrivateAccessServiceConnectionsModel) refreshPrivateAccessServiceConnections(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
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
		m.Config = m.Config.flattenPrivateAccessServiceConnectionsConfig(ctx, v, &diags)
	}

	if v, ok := o["common_config"]; ok {
		m.CommonConfig = m.CommonConfig.flattenPrivateAccessServiceConnectionsCommonConfig(ctx, v, &diags)
	}

	if v, ok := o["ip_assigned"]; ok {
		m.IpAssigned = m.flattenPrivateAccessServiceConnectionsIpAssignedList(ctx, v, &diags)
	}

	return diags
}

func (data *datasourcePrivateAccessServiceConnectionsModel) getURLObjectPrivateAccessServiceConnections(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ServiceConnectionId.IsNull() {
		result["service-connection-id"] = data.ServiceConnectionId.ValueString()
	}

	return &result
}

type datasourcePrivateAccessServiceConnectionsBackupLinksModel struct {
	Alias             types.String `tfsdk:"alias"`
	Auth              types.String `tfsdk:"auth"`
	IpsecCertName     types.String `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion   types.String `tfsdk:"ipsec_ike_version"`
	IpsecPeerName     types.String `tfsdk:"ipsec_peer_name"`
	IpsecRemoteGw     types.String `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId  types.String `tfsdk:"overlay_network_id"`
	IpsecPreSharedKey types.String `tfsdk:"ipsec_pre_shared_key"`
}

type datasourcePrivateAccessServiceConnectionsConfigModel struct {
	Alias            types.String                                                      `tfsdk:"alias"`
	Auth             types.String                                                      `tfsdk:"auth"`
	BgpPeerIp        types.String                                                      `tfsdk:"bgp_peer_ip"`
	IpsecCertName    types.String                                                      `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion  types.String                                                      `tfsdk:"ipsec_ike_version"`
	IpsecPeerName    types.String                                                      `tfsdk:"ipsec_peer_name"`
	IpsecRemoteGw    types.String                                                      `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId types.String                                                      `tfsdk:"overlay_network_id"`
	RouteMapTag      types.String                                                      `tfsdk:"route_map_tag"`
	RegionCost       types.Map                                                         `tfsdk:"region_cost"`
	BackupLinks      []datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel `tfsdk:"backup_links"`
}

type datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel struct {
	Id               types.String `tfsdk:"id"`
	Alias            types.String `tfsdk:"alias"`
	Auth             types.String `tfsdk:"auth"`
	IpsecCertName    types.String `tfsdk:"ipsec_cert_name"`
	IpsecIkeVersion  types.String `tfsdk:"ipsec_ike_version"`
	IpsecPeerName    types.String `tfsdk:"ipsec_peer_name"`
	IpsecRemoteGw    types.String `tfsdk:"ipsec_remote_gw"`
	OverlayNetworkId types.String `tfsdk:"overlay_network_id"`
}

type datasourcePrivateAccessServiceConnectionsCommonConfigModel struct {
	ConfigState        types.String `tfsdk:"config_state"`
	BgpDesign          types.String `tfsdk:"bgp_design"`
	BgpRouterIdsSubnet types.String `tfsdk:"bgp_router_ids_subnet"`
	AsNumber           types.String `tfsdk:"as_number"`
	RecursiveNextHop   types.Bool   `tfsdk:"recursive_next_hop"`
	SdwanRuleEnable    types.Bool   `tfsdk:"sdwan_rule_enable"`
	SdwanHealthCheckVm types.String `tfsdk:"sdwan_health_check_vm"`
}

type datasourcePrivateAccessServiceConnectionsIpAssignedModel struct {
	Id            types.String `tfsdk:"id"`
	SdwanCommonId types.String `tfsdk:"sdwan_common_id"`
	BgpRouterId   types.String `tfsdk:"bgp_router_id"`
	SiteId        types.String `tfsdk:"site_id"`
	Region        types.String `tfsdk:"region"`
}

func (m *datasourcePrivateAccessServiceConnectionsBackupLinksModel) flattenPrivateAccessServiceConnectionsBackupLinks(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourcePrivateAccessServiceConnectionsBackupLinksModel {
	if input == nil {
		return &datasourcePrivateAccessServiceConnectionsBackupLinksModel{}
	}
	if m == nil {
		m = &datasourcePrivateAccessServiceConnectionsBackupLinksModel{}
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

func (s *datasourcePrivateAccessServiceConnectionsModel) flattenPrivateAccessServiceConnectionsBackupLinksList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourcePrivateAccessServiceConnectionsBackupLinksModel {
	if o == nil {
		return []datasourcePrivateAccessServiceConnectionsBackupLinksModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument backup_links is not type of []interface{}.", "")
		return []datasourcePrivateAccessServiceConnectionsBackupLinksModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourcePrivateAccessServiceConnectionsBackupLinksModel{}
	}

	values := make([]datasourcePrivateAccessServiceConnectionsBackupLinksModel, len(l))
	for i, ele := range l {
		var m datasourcePrivateAccessServiceConnectionsBackupLinksModel
		values[i] = *m.flattenPrivateAccessServiceConnectionsBackupLinks(ctx, ele, diags)
	}

	return values
}

func (m *datasourcePrivateAccessServiceConnectionsConfigModel) flattenPrivateAccessServiceConnectionsConfig(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourcePrivateAccessServiceConnectionsConfigModel {
	if input == nil {
		return &datasourcePrivateAccessServiceConnectionsConfigModel{}
	}
	if m == nil {
		m = &datasourcePrivateAccessServiceConnectionsConfigModel{}
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
		m.BackupLinks = m.flattenPrivateAccessServiceConnectionsConfigBackupLinksList(ctx, v, diags)
	}

	return m
}

func (m *datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel) flattenPrivateAccessServiceConnectionsConfigBackupLinks(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel {
	if input == nil {
		return &datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel{}
	}
	if m == nil {
		m = &datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel{}
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

func (s *datasourcePrivateAccessServiceConnectionsConfigModel) flattenPrivateAccessServiceConnectionsConfigBackupLinksList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel {
	if o == nil {
		return []datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument backup_links is not type of []interface{}.", "")
		return []datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel{}
	}

	values := make([]datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel, len(l))
	for i, ele := range l {
		var m datasourcePrivateAccessServiceConnectionsConfigBackupLinksModel
		values[i] = *m.flattenPrivateAccessServiceConnectionsConfigBackupLinks(ctx, ele, diags)
	}

	return values
}

func (m *datasourcePrivateAccessServiceConnectionsCommonConfigModel) flattenPrivateAccessServiceConnectionsCommonConfig(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourcePrivateAccessServiceConnectionsCommonConfigModel {
	if input == nil {
		return &datasourcePrivateAccessServiceConnectionsCommonConfigModel{}
	}
	if m == nil {
		m = &datasourcePrivateAccessServiceConnectionsCommonConfigModel{}
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

func (m *datasourcePrivateAccessServiceConnectionsIpAssignedModel) flattenPrivateAccessServiceConnectionsIpAssigned(ctx context.Context, input interface{}, diags *diag.Diagnostics) *datasourcePrivateAccessServiceConnectionsIpAssignedModel {
	if input == nil {
		return &datasourcePrivateAccessServiceConnectionsIpAssignedModel{}
	}
	if m == nil {
		m = &datasourcePrivateAccessServiceConnectionsIpAssignedModel{}
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

func (s *datasourcePrivateAccessServiceConnectionsModel) flattenPrivateAccessServiceConnectionsIpAssignedList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []datasourcePrivateAccessServiceConnectionsIpAssignedModel {
	if o == nil {
		return []datasourcePrivateAccessServiceConnectionsIpAssignedModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument ip_assigned is not type of []interface{}.", "")
		return []datasourcePrivateAccessServiceConnectionsIpAssignedModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []datasourcePrivateAccessServiceConnectionsIpAssignedModel{}
	}

	values := make([]datasourcePrivateAccessServiceConnectionsIpAssignedModel, len(l))
	for i, ele := range l {
		var m datasourcePrivateAccessServiceConnectionsIpAssignedModel
		values[i] = *m.flattenPrivateAccessServiceConnectionsIpAssigned(ctx, ele, diags)
	}

	return values
}
