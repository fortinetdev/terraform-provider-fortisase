// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceEndpointConnectionProfiles{}

func newResourceEndpointConnectionProfiles() resource.Resource {
	return &resourceEndpointConnectionProfiles{}
}

type resourceEndpointConnectionProfiles struct {
	fortiClient *FortiClient
}

// resourceEndpointConnectionProfilesModel describes the resource data model.
type resourceEndpointConnectionProfilesModel struct {
	ID                             types.String                                                 `tfsdk:"id"`
	ConnectToFortiSase             types.String                                                 `tfsdk:"connect_to_forti_sase"`
	Lockdown                       *resourceEndpointConnectionProfilesLockdownModel             `tfsdk:"lockdown"`
	OnFabricRuleSet                *resourceEndpointConnectionProfilesOnFabricRuleSetModel      `tfsdk:"on_fabric_rule_set"`
	OffNetSplitTunnel              *resourceEndpointConnectionProfilesOffNetSplitTunnelModel    `tfsdk:"off_net_split_tunnel"`
	SplitTunnel                    *resourceEndpointConnectionProfilesSplitTunnelModel          `tfsdk:"split_tunnel"`
	AllowInvalidServerCertificate  types.String                                                 `tfsdk:"allow_invalid_server_certificate"`
	EndpointOnNetBypass            types.Bool                                                   `tfsdk:"endpoint_on_net_bypass"`
	AuthBeforeUserLogon            types.Bool                                                   `tfsdk:"auth_before_user_logon"`
	SecureInternetAccess           *resourceEndpointConnectionProfilesSecureInternetAccessModel `tfsdk:"secure_internet_access"`
	PreferredDtlsTunnel            types.String                                                 `tfsdk:"preferred_dtls_tunnel"`
	UseGuiSamlAuth                 types.String                                                 `tfsdk:"use_gui_saml_auth"`
	AllowPersonalVpns              types.Bool                                                   `tfsdk:"allow_personal_vpns"`
	MtuSize                        types.Float64                                                `tfsdk:"mtu_size"`
	AvailableVpNs                  []resourceEndpointConnectionProfilesAvailableVpNsModel       `tfsdk:"available_vp_ns"`
	ShowDisconnectBtn              types.String                                                 `tfsdk:"show_disconnect_btn"`
	EnableInvalidServerCertWarning types.String                                                 `tfsdk:"enable_invalid_server_cert_warning"`
	PreLogon                       *resourceEndpointConnectionProfilesPreLogonModel             `tfsdk:"pre_logon"`
	PrimaryKey                     types.String                                                 `tfsdk:"primary_key"`
}

func (r *resourceEndpointConnectionProfiles) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint_connection_profiles"
}

func (r *resourceEndpointConnectionProfiles) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"connect_to_forti_sase": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("automatically", "manually"),
				},
				Computed: true,
				Optional: true,
			},
			"allow_invalid_server_certificate": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"endpoint_on_net_bypass": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"auth_before_user_logon": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"preferred_dtls_tunnel": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"use_gui_saml_auth": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"allow_personal_vpns": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"mtu_size": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(576, 1500),
				},
				Computed: true,
				Optional: true,
			},
			"show_disconnect_btn": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"enable_invalid_server_cert_warning": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"primary_key": schema.StringAttribute{
				Description: "The primary key of the object. Can be found in the response from the get request.",
				Required:    true,
			},
			"lockdown": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"grace_period": schema.Float64Attribute{
						Computed: true,
						Optional: true,
					},
					"max_attempts": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.AtLeast(1),
						},
						Computed: true,
						Optional: true,
					},
					"ips": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"port": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"protocol": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("tcp", "udp", "icmp", ""),
									},
									Computed: true,
									Optional: true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
					"domains": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
					"detect_captive_portal": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("enable", "disable"),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"on_fabric_rule_set": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"primary_key": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"datasource": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("endpoint/on-net-rules"),
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"off_net_split_tunnel": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"local_apps": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"fqdns": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"subnets_ipsec": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"isdbs": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("network/basic-internet-services"),
									},
									Computed: true,
									Optional: true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
					"subnets": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("network/hosts", "network/host-groups"),
									},
									Computed: true,
									Optional: true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"split_tunnel": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"local_apps": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"fqdns": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"subnets_ipsec": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"isdbs": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("network/basic-internet-services"),
									},
									Computed: true,
									Optional: true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
					"subnets": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"primary_key": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"datasource": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("network/hosts", "network/host-groups"),
									},
									Computed: true,
									Optional: true,
								},
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"secure_internet_access": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"authenticate_with_sso": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"enable_local_lan": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"failover_sequence": schema.SetAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"external_browser_saml_login": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("enable", "disable"),
						},
						Computed: true,
						Optional: true,
					},
					"posture_check": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"tag": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
							"action": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("allow", "prohibit"),
								},
								Computed: true,
								Optional: true,
							},
							"check_failed_message": schema.StringAttribute{
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
				},
				Computed: true,
				Optional: true,
			},
			"available_vp_ns": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("sslVPN", "ipSecVPN"),
							},
							Computed: true,
							Optional: true,
						},
						"name": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"remote_gateway": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
							Computed: true,
							Optional: true,
						},
						"username_prompt": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"save_username": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_always_up": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_auto_connect": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"show_remember_password": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"authenticate_with_sso": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"enable_local_lan": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"external_browser_saml_login": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"port": schema.Float64Attribute{
							Validators: []validator.Float64{
								float64validator.AtMost(65535),
							},
							Computed: true,
							Optional: true,
						},
						"require_certificate": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"auth_method": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("preSharedKey", "smartCardCert", "systemStoreCert"),
							},
							Computed: true,
							Optional: true,
						},
						"show_passcode": schema.StringAttribute{
							Validators: []validator.String{
								stringvalidator.OneOf("enable", "disable"),
							},
							Computed: true,
							Optional: true,
						},
						"pre_shared_key": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
						"posture_check": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"tag": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
								"action": schema.StringAttribute{
									Validators: []validator.String{
										stringvalidator.OneOf("allow", "prohibit"),
									},
									Computed: true,
									Optional: true,
								},
								"check_failed_message": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
							},
							Computed: true,
							Optional: true,
						},
					},
				},
				Computed: true,
				Optional: true,
			},
			"pre_logon": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"vpn_type": schema.StringAttribute{
						Validators: []validator.String{
							stringvalidator.OneOf("sslVPN", "ipSecVPN"),
						},
						Computed: true,
						Optional: true,
					},
					"remote_gateway": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"port": schema.Float64Attribute{
						Validators: []validator.Float64{
							float64validator.AtMost(65535),
						},
						Computed: true,
						Optional: true,
					},
					"common_name": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"match_type": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("wildcard", "regex"),
								},
								Computed: true,
								Optional: true,
							},
							"pattern": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
								Computed: true,
								Optional: true,
							},
						},
						Computed: true,
						Optional: true,
					},
					"issuer": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"match_type": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.OneOf("wildcard", "regex"),
								},
								Computed: true,
								Optional: true,
							},
							"pattern": schema.StringAttribute{
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
								Computed: true,
								Optional: true,
							},
						},
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

func (r *resourceEndpointConnectionProfiles) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *resourceEndpointConnectionProfiles) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceEndpointConnectionProfilesModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = data.PrimaryKey.ValueString()
	input_model.BodyParams = *(data.getCreateObjectEndpointConnectionProfiles(ctx, diags))
	input_model.URLParams = *(data.getURLObjectEndpointConnectionProfiles(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.UpdateEndpointConnectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointConnectionProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointConnectionProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointConnectionProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointConnectionProfiles) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceEndpointConnectionProfilesModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceEndpointConnectionProfilesModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectEndpointConnectionProfiles(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectEndpointConnectionProfiles(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	_, err := c.UpdateEndpointConnectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectEndpointConnectionProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointConnectionProfiles(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointConnectionProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointConnectionProfiles) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourceEndpointConnectionProfiles) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceEndpointConnectionProfilesModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectEndpointConnectionProfiles(ctx, "read", diags))

	read_output, err := c.ReadEndpointConnectionProfiles(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource: %v", err),
			"",
		)
		return
	}

	diags.Append(data.refreshEndpointConnectionProfiles(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceEndpointConnectionProfiles) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceEndpointConnectionProfilesModel) refreshEndpointConnectionProfiles(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["connectToFortiSASE"]; ok {
		m.ConnectToFortiSase = parseStringValue(v)
	}

	if v, ok := o["lockdown"]; ok {
		m.Lockdown = m.Lockdown.flattenEndpointConnectionProfilesLockdown(ctx, v, &diags)
	}

	if v, ok := o["onFabricRuleSet"]; ok {
		m.OnFabricRuleSet = m.OnFabricRuleSet.flattenEndpointConnectionProfilesOnFabricRuleSet(ctx, v, &diags)
	}

	if v, ok := o["offNetSplitTunnel"]; ok {
		m.OffNetSplitTunnel = m.OffNetSplitTunnel.flattenEndpointConnectionProfilesOffNetSplitTunnel(ctx, v, &diags)
	}

	if v, ok := o["splitTunnel"]; ok {
		m.SplitTunnel = m.SplitTunnel.flattenEndpointConnectionProfilesSplitTunnel(ctx, v, &diags)
	}

	if v, ok := o["allowInvalidServerCertificate"]; ok {
		m.AllowInvalidServerCertificate = parseStringValue(v)
	}

	if v, ok := o["endpointOnNetBypass"]; ok {
		m.EndpointOnNetBypass = parseBoolValue(v)
	}

	if v, ok := o["authBeforeUserLogon"]; ok {
		m.AuthBeforeUserLogon = parseBoolValue(v)
	}

	if v, ok := o["secureInternetAccess"]; ok {
		m.SecureInternetAccess = m.SecureInternetAccess.flattenEndpointConnectionProfilesSecureInternetAccess(ctx, v, &diags)
	}

	if v, ok := o["preferredDTLSTunnel"]; ok {
		m.PreferredDtlsTunnel = parseStringValue(v)
	}

	if v, ok := o["useGuiSamlAuth"]; ok {
		m.UseGuiSamlAuth = parseStringValue(v)
	}

	if v, ok := o["allowPersonalVpns"]; ok {
		m.AllowPersonalVpns = parseBoolValue(v)
	}

	if v, ok := o["mtuSize"]; ok {
		m.MtuSize = parseFloat64Value(v)
	}

	if v, ok := o["availableVPNs"]; ok {
		m.AvailableVpNs = m.flattenEndpointConnectionProfilesAvailableVpNsList(ctx, v, &diags)
	}

	if v, ok := o["showDisconnectBtn"]; ok {
		m.ShowDisconnectBtn = parseStringValue(v)
	}

	if v, ok := o["enableInvalidServerCertWarning"]; ok {
		m.EnableInvalidServerCertWarning = parseStringValue(v)
	}

	if v, ok := o["preLogon"]; ok {
		m.PreLogon = m.PreLogon.flattenEndpointConnectionProfilesPreLogon(ctx, v, &diags)
	}

	return diags
}

func (data *resourceEndpointConnectionProfilesModel) getCreateObjectEndpointConnectionProfiles(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ConnectToFortiSase.IsNull() {
		result["connectToFortiSASE"] = data.ConnectToFortiSase.ValueString()
	}

	if data.Lockdown != nil && !isZeroStruct(*data.Lockdown) {
		result["lockdown"] = data.Lockdown.expandEndpointConnectionProfilesLockdown(ctx, diags)
	}

	if data.OnFabricRuleSet != nil && !isZeroStruct(*data.OnFabricRuleSet) {
		result["onFabricRuleSet"] = data.OnFabricRuleSet.expandEndpointConnectionProfilesOnFabricRuleSet(ctx, diags)
	}

	if data.OffNetSplitTunnel != nil && !isZeroStruct(*data.OffNetSplitTunnel) {
		result["offNetSplitTunnel"] = data.OffNetSplitTunnel.expandEndpointConnectionProfilesOffNetSplitTunnel(ctx, diags)
	}

	if data.SplitTunnel != nil && !isZeroStruct(*data.SplitTunnel) {
		result["splitTunnel"] = data.SplitTunnel.expandEndpointConnectionProfilesSplitTunnel(ctx, diags)
	}

	if !data.AllowInvalidServerCertificate.IsNull() {
		result["allowInvalidServerCertificate"] = data.AllowInvalidServerCertificate.ValueString()
	}

	if !data.EndpointOnNetBypass.IsNull() {
		result["endpointOnNetBypass"] = data.EndpointOnNetBypass.ValueBool()
	}

	if !data.AuthBeforeUserLogon.IsNull() {
		result["authBeforeUserLogon"] = data.AuthBeforeUserLogon.ValueBool()
	}

	if data.SecureInternetAccess != nil && !isZeroStruct(*data.SecureInternetAccess) {
		result["secureInternetAccess"] = data.SecureInternetAccess.expandEndpointConnectionProfilesSecureInternetAccess(ctx, diags)
	}

	if !data.PreferredDtlsTunnel.IsNull() {
		result["preferredDTLSTunnel"] = data.PreferredDtlsTunnel.ValueString()
	}

	if !data.UseGuiSamlAuth.IsNull() {
		result["useGuiSamlAuth"] = data.UseGuiSamlAuth.ValueString()
	}

	if !data.AllowPersonalVpns.IsNull() {
		result["allowPersonalVpns"] = data.AllowPersonalVpns.ValueBool()
	}

	if !data.MtuSize.IsNull() {
		result["mtuSize"] = data.MtuSize.ValueFloat64()
	}

	if len(data.AvailableVpNs) > 0 {
		result["availableVPNs"] = data.expandEndpointConnectionProfilesAvailableVpNsList(ctx, data.AvailableVpNs, diags)
	}

	if !data.ShowDisconnectBtn.IsNull() {
		result["showDisconnectBtn"] = data.ShowDisconnectBtn.ValueString()
	}

	if !data.EnableInvalidServerCertWarning.IsNull() {
		result["enableInvalidServerCertWarning"] = data.EnableInvalidServerCertWarning.ValueString()
	}

	if data.PreLogon != nil && !isZeroStruct(*data.PreLogon) {
		result["preLogon"] = data.PreLogon.expandEndpointConnectionProfilesPreLogon(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointConnectionProfilesModel) getUpdateObjectEndpointConnectionProfiles(ctx context.Context, state resourceEndpointConnectionProfilesModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.ConnectToFortiSase.IsNull() {
		result["connectToFortiSASE"] = data.ConnectToFortiSase.ValueString()
	}

	if data.Lockdown != nil && !isSameStruct(data.Lockdown, state.Lockdown) {
		result["lockdown"] = data.Lockdown.expandEndpointConnectionProfilesLockdown(ctx, diags)
	}

	if data.OnFabricRuleSet != nil && !isSameStruct(data.OnFabricRuleSet, state.OnFabricRuleSet) {
		result["onFabricRuleSet"] = data.OnFabricRuleSet.expandEndpointConnectionProfilesOnFabricRuleSet(ctx, diags)
	}

	if data.OffNetSplitTunnel != nil && !isSameStruct(data.OffNetSplitTunnel, state.OffNetSplitTunnel) {
		result["offNetSplitTunnel"] = data.OffNetSplitTunnel.expandEndpointConnectionProfilesOffNetSplitTunnel(ctx, diags)
	}

	if data.SplitTunnel != nil && !isSameStruct(data.SplitTunnel, state.SplitTunnel) {
		result["splitTunnel"] = data.SplitTunnel.expandEndpointConnectionProfilesSplitTunnel(ctx, diags)
	}

	if !data.AllowInvalidServerCertificate.IsNull() {
		result["allowInvalidServerCertificate"] = data.AllowInvalidServerCertificate.ValueString()
	}

	if !data.EndpointOnNetBypass.IsNull() {
		result["endpointOnNetBypass"] = data.EndpointOnNetBypass.ValueBool()
	}

	if !data.AuthBeforeUserLogon.IsNull() {
		result["authBeforeUserLogon"] = data.AuthBeforeUserLogon.ValueBool()
	}

	if data.SecureInternetAccess != nil && !isSameStruct(data.SecureInternetAccess, state.SecureInternetAccess) {
		result["secureInternetAccess"] = data.SecureInternetAccess.expandEndpointConnectionProfilesSecureInternetAccess(ctx, diags)
	}

	if !data.PreferredDtlsTunnel.IsNull() && !data.PreferredDtlsTunnel.Equal(state.PreferredDtlsTunnel) {
		result["preferredDTLSTunnel"] = data.PreferredDtlsTunnel.ValueString()
	}

	if !data.UseGuiSamlAuth.IsNull() && !data.UseGuiSamlAuth.Equal(state.UseGuiSamlAuth) {
		result["useGuiSamlAuth"] = data.UseGuiSamlAuth.ValueString()
	}

	if !data.AllowPersonalVpns.IsNull() {
		result["allowPersonalVpns"] = data.AllowPersonalVpns.ValueBool()
	}

	if !data.MtuSize.IsNull() && !data.MtuSize.Equal(state.MtuSize) {
		result["mtuSize"] = data.MtuSize.ValueFloat64()
	}

	if len(data.AvailableVpNs) > 0 || !isSameStruct(data.AvailableVpNs, state.AvailableVpNs) {
		result["availableVPNs"] = data.expandEndpointConnectionProfilesAvailableVpNsList(ctx, data.AvailableVpNs, diags)
	}

	if !data.ShowDisconnectBtn.IsNull() && !data.ShowDisconnectBtn.Equal(state.ShowDisconnectBtn) {
		result["showDisconnectBtn"] = data.ShowDisconnectBtn.ValueString()
	}

	if !data.EnableInvalidServerCertWarning.IsNull() && !data.EnableInvalidServerCertWarning.Equal(state.EnableInvalidServerCertWarning) {
		result["enableInvalidServerCertWarning"] = data.EnableInvalidServerCertWarning.ValueString()
	}

	if data.PreLogon != nil && !isSameStruct(data.PreLogon, state.PreLogon) {
		result["preLogon"] = data.PreLogon.expandEndpointConnectionProfilesPreLogon(ctx, diags)
	}

	return &result
}

func (data *resourceEndpointConnectionProfilesModel) getURLObjectEndpointConnectionProfiles(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}

type resourceEndpointConnectionProfilesLockdownModel struct {
	Status              types.String                                                        `tfsdk:"status"`
	GracePeriod         types.Float64                                                       `tfsdk:"grace_period"`
	MaxAttempts         types.Float64                                                       `tfsdk:"max_attempts"`
	Ips                 []resourceEndpointConnectionProfilesLockdownIpsModel                `tfsdk:"ips"`
	Domains             []resourceEndpointConnectionProfilesLockdownDomainsModel            `tfsdk:"domains"`
	DetectCaptivePortal *resourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel `tfsdk:"detect_captive_portal"`
}

type resourceEndpointConnectionProfilesLockdownIpsModel struct {
	Ip       types.String `tfsdk:"ip"`
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
}

type resourceEndpointConnectionProfilesLockdownDomainsModel struct {
	Address types.String `tfsdk:"address"`
}

type resourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel struct {
	Status types.String `tfsdk:"status"`
}

type resourceEndpointConnectionProfilesOnFabricRuleSetModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointConnectionProfilesOffNetSplitTunnelModel struct {
	LocalApps    types.Set                                                         `tfsdk:"local_apps"`
	Isdbs        []resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel   `tfsdk:"isdbs"`
	Fqdns        types.Set                                                         `tfsdk:"fqdns"`
	Subnets      []resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel `tfsdk:"subnets"`
	SubnetsIpsec types.Set                                                         `tfsdk:"subnets_ipsec"`
}

type resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointConnectionProfilesSplitTunnelModel struct {
	LocalApps    types.Set                                                   `tfsdk:"local_apps"`
	Isdbs        []resourceEndpointConnectionProfilesSplitTunnelIsdbsModel   `tfsdk:"isdbs"`
	Fqdns        types.Set                                                   `tfsdk:"fqdns"`
	Subnets      []resourceEndpointConnectionProfilesSplitTunnelSubnetsModel `tfsdk:"subnets"`
	SubnetsIpsec types.Set                                                   `tfsdk:"subnets_ipsec"`
}

type resourceEndpointConnectionProfilesSplitTunnelIsdbsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointConnectionProfilesSplitTunnelSubnetsModel struct {
	PrimaryKey types.String `tfsdk:"primary_key"`
	Datasource types.String `tfsdk:"datasource"`
}

type resourceEndpointConnectionProfilesSecureInternetAccessModel struct {
	AuthenticateWithSso      types.String                                                             `tfsdk:"authenticate_with_sso"`
	EnableLocalLan           types.String                                                             `tfsdk:"enable_local_lan"`
	FailoverSequence         types.Set                                                                `tfsdk:"failover_sequence"`
	PostureCheck             *resourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel `tfsdk:"posture_check"`
	ExternalBrowserSamlLogin types.String                                                             `tfsdk:"external_browser_saml_login"`
}

type resourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel struct {
	Tag                types.String `tfsdk:"tag"`
	Action             types.String `tfsdk:"action"`
	CheckFailedMessage types.String `tfsdk:"check_failed_message"`
}

type resourceEndpointConnectionProfilesAvailableVpNsModel struct {
	Type                     types.String                                                      `tfsdk:"type"`
	Name                     types.String                                                      `tfsdk:"name"`
	RemoteGateway            types.String                                                      `tfsdk:"remote_gateway"`
	UsernamePrompt           types.String                                                      `tfsdk:"username_prompt"`
	SaveUsername             types.String                                                      `tfsdk:"save_username"`
	ShowAlwaysUp             types.String                                                      `tfsdk:"show_always_up"`
	ShowAutoConnect          types.String                                                      `tfsdk:"show_auto_connect"`
	ShowRememberPassword     types.String                                                      `tfsdk:"show_remember_password"`
	AuthenticateWithSso      types.String                                                      `tfsdk:"authenticate_with_sso"`
	EnableLocalLan           types.String                                                      `tfsdk:"enable_local_lan"`
	ExternalBrowserSamlLogin types.String                                                      `tfsdk:"external_browser_saml_login"`
	Port                     types.Float64                                                     `tfsdk:"port"`
	RequireCertificate       types.String                                                      `tfsdk:"require_certificate"`
	AuthMethod               types.String                                                      `tfsdk:"auth_method"`
	ShowPasscode             types.String                                                      `tfsdk:"show_passcode"`
	PostureCheck             *resourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel `tfsdk:"posture_check"`
	PreSharedKey             types.String                                                      `tfsdk:"pre_shared_key"`
}

type resourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel struct {
	Tag                types.String `tfsdk:"tag"`
	Action             types.String `tfsdk:"action"`
	CheckFailedMessage types.String `tfsdk:"check_failed_message"`
}

type resourceEndpointConnectionProfilesPreLogonModel struct {
	VpnType       types.String                                               `tfsdk:"vpn_type"`
	RemoteGateway types.String                                               `tfsdk:"remote_gateway"`
	CommonName    *resourceEndpointConnectionProfilesPreLogonCommonNameModel `tfsdk:"common_name"`
	Issuer        *resourceEndpointConnectionProfilesPreLogonIssuerModel     `tfsdk:"issuer"`
	Port          types.Float64                                              `tfsdk:"port"`
}

type resourceEndpointConnectionProfilesPreLogonCommonNameModel struct {
	MatchType types.String `tfsdk:"match_type"`
	Pattern   types.String `tfsdk:"pattern"`
}

type resourceEndpointConnectionProfilesPreLogonIssuerModel struct {
	MatchType types.String `tfsdk:"match_type"`
	Pattern   types.String `tfsdk:"pattern"`
}

func (m *resourceEndpointConnectionProfilesLockdownModel) flattenEndpointConnectionProfilesLockdown(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesLockdownModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesLockdownModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesLockdownModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["gracePeriod"]; ok {
		m.GracePeriod = parseFloat64Value(v)
	}

	if v, ok := o["maxAttempts"]; ok {
		m.MaxAttempts = parseFloat64Value(v)
	}

	if v, ok := o["ips"]; ok {
		m.Ips = m.flattenEndpointConnectionProfilesLockdownIpsList(ctx, v, diags)
	}

	if v, ok := o["domains"]; ok {
		m.Domains = m.flattenEndpointConnectionProfilesLockdownDomainsList(ctx, v, diags)
	}

	if v, ok := o["detectCaptivePortal"]; ok {
		m.DetectCaptivePortal = m.DetectCaptivePortal.flattenEndpointConnectionProfilesLockdownDetectCaptivePortal(ctx, v, diags)
	}

	return m
}

func (m *resourceEndpointConnectionProfilesLockdownIpsModel) flattenEndpointConnectionProfilesLockdownIps(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesLockdownIpsModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesLockdownIpsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesLockdownIpsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["ip"]; ok {
		m.Ip = parseStringValue(v)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseStringValue(v)
	}

	if v, ok := o["protocol"]; ok {
		m.Protocol = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointConnectionProfilesLockdownModel) flattenEndpointConnectionProfilesLockdownIpsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfilesLockdownIpsModel {
	if o == nil {
		return []resourceEndpointConnectionProfilesLockdownIpsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument ips is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfilesLockdownIpsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfilesLockdownIpsModel{}
	}

	values := make([]resourceEndpointConnectionProfilesLockdownIpsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfilesLockdownIpsModel
		values[i] = *m.flattenEndpointConnectionProfilesLockdownIps(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfilesLockdownDomainsModel) flattenEndpointConnectionProfilesLockdownDomains(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesLockdownDomainsModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesLockdownDomainsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesLockdownDomainsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["address"]; ok {
		m.Address = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointConnectionProfilesLockdownModel) flattenEndpointConnectionProfilesLockdownDomainsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfilesLockdownDomainsModel {
	if o == nil {
		return []resourceEndpointConnectionProfilesLockdownDomainsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument domains is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfilesLockdownDomainsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfilesLockdownDomainsModel{}
	}

	values := make([]resourceEndpointConnectionProfilesLockdownDomainsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfilesLockdownDomainsModel
		values[i] = *m.flattenEndpointConnectionProfilesLockdownDomains(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel) flattenEndpointConnectionProfilesLockdownDetectCaptivePortal(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	return m
}

func (m *resourceEndpointConnectionProfilesOnFabricRuleSetModel) flattenEndpointConnectionProfilesOnFabricRuleSet(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesOnFabricRuleSetModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesOnFabricRuleSetModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesOnFabricRuleSetModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (m *resourceEndpointConnectionProfilesOffNetSplitTunnelModel) flattenEndpointConnectionProfilesOffNetSplitTunnel(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesOffNetSplitTunnelModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesOffNetSplitTunnelModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesOffNetSplitTunnelModel{}
	}
	o := input.(map[string]interface{})
	m.LocalApps = types.SetNull(types.StringType)
	m.Fqdns = types.SetNull(types.StringType)
	m.SubnetsIpsec = types.SetNull(types.StringType)

	if v, ok := o["localApps"]; ok {
		m.LocalApps = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["isdbs"]; ok {
		m.Isdbs = m.flattenEndpointConnectionProfilesOffNetSplitTunnelIsdbsList(ctx, v, diags)
	}

	if v, ok := o["fqdns"]; ok {
		m.Fqdns = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["subnets"]; ok {
		m.Subnets = m.flattenEndpointConnectionProfilesOffNetSplitTunnelSubnetsList(ctx, v, diags)
	}

	if v, ok := o["subnetsIpsec"]; ok {
		m.SubnetsIpsec = parseSetValue(ctx, v, types.StringType)
	}

	return m
}

func (m *resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel) flattenEndpointConnectionProfilesOffNetSplitTunnelIsdbs(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointConnectionProfilesOffNetSplitTunnelModel) flattenEndpointConnectionProfilesOffNetSplitTunnelIsdbsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel {
	if o == nil {
		return []resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument isdbs is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel{}
	}

	values := make([]resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel
		values[i] = *m.flattenEndpointConnectionProfilesOffNetSplitTunnelIsdbs(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel) flattenEndpointConnectionProfilesOffNetSplitTunnelSubnets(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointConnectionProfilesOffNetSplitTunnelModel) flattenEndpointConnectionProfilesOffNetSplitTunnelSubnetsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel {
	if o == nil {
		return []resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument subnets is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel{}
	}

	values := make([]resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel
		values[i] = *m.flattenEndpointConnectionProfilesOffNetSplitTunnelSubnets(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfilesSplitTunnelModel) flattenEndpointConnectionProfilesSplitTunnel(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesSplitTunnelModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesSplitTunnelModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesSplitTunnelModel{}
	}
	o := input.(map[string]interface{})
	m.LocalApps = types.SetNull(types.StringType)
	m.Fqdns = types.SetNull(types.StringType)
	m.SubnetsIpsec = types.SetNull(types.StringType)

	if v, ok := o["localApps"]; ok {
		m.LocalApps = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["isdbs"]; ok {
		m.Isdbs = m.flattenEndpointConnectionProfilesSplitTunnelIsdbsList(ctx, v, diags)
	}

	if v, ok := o["fqdns"]; ok {
		m.Fqdns = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["subnets"]; ok {
		m.Subnets = m.flattenEndpointConnectionProfilesSplitTunnelSubnetsList(ctx, v, diags)
	}

	if v, ok := o["subnetsIpsec"]; ok {
		m.SubnetsIpsec = parseSetValue(ctx, v, types.StringType)
	}

	return m
}

func (m *resourceEndpointConnectionProfilesSplitTunnelIsdbsModel) flattenEndpointConnectionProfilesSplitTunnelIsdbs(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesSplitTunnelIsdbsModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesSplitTunnelIsdbsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesSplitTunnelIsdbsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointConnectionProfilesSplitTunnelModel) flattenEndpointConnectionProfilesSplitTunnelIsdbsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfilesSplitTunnelIsdbsModel {
	if o == nil {
		return []resourceEndpointConnectionProfilesSplitTunnelIsdbsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument isdbs is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfilesSplitTunnelIsdbsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfilesSplitTunnelIsdbsModel{}
	}

	values := make([]resourceEndpointConnectionProfilesSplitTunnelIsdbsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfilesSplitTunnelIsdbsModel
		values[i] = *m.flattenEndpointConnectionProfilesSplitTunnelIsdbs(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfilesSplitTunnelSubnetsModel) flattenEndpointConnectionProfilesSplitTunnelSubnets(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesSplitTunnelSubnetsModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesSplitTunnelSubnetsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesSplitTunnelSubnetsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["primaryKey"]; ok {
		m.PrimaryKey = parseStringValue(v)
	}

	if v, ok := o["datasource"]; ok {
		m.Datasource = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointConnectionProfilesSplitTunnelModel) flattenEndpointConnectionProfilesSplitTunnelSubnetsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfilesSplitTunnelSubnetsModel {
	if o == nil {
		return []resourceEndpointConnectionProfilesSplitTunnelSubnetsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument subnets is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfilesSplitTunnelSubnetsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfilesSplitTunnelSubnetsModel{}
	}

	values := make([]resourceEndpointConnectionProfilesSplitTunnelSubnetsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfilesSplitTunnelSubnetsModel
		values[i] = *m.flattenEndpointConnectionProfilesSplitTunnelSubnets(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfilesSecureInternetAccessModel) flattenEndpointConnectionProfilesSecureInternetAccess(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesSecureInternetAccessModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesSecureInternetAccessModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesSecureInternetAccessModel{}
	}
	o := input.(map[string]interface{})
	m.FailoverSequence = types.SetNull(types.StringType)

	if v, ok := o["authenticateWithSSO"]; ok {
		m.AuthenticateWithSso = parseStringValue(v)
	}

	if v, ok := o["enableLocalLan"]; ok {
		m.EnableLocalLan = parseStringValue(v)
	}

	if v, ok := o["failoverSequence"]; ok {
		m.FailoverSequence = parseSetValue(ctx, v, types.StringType)
	}

	if v, ok := o["postureCheck"]; ok {
		m.PostureCheck = m.PostureCheck.flattenEndpointConnectionProfilesSecureInternetAccessPostureCheck(ctx, v, diags)
	}

	if v, ok := o["externalBrowserSamlLogin"]; ok {
		m.ExternalBrowserSamlLogin = parseStringValue(v)
	}

	return m
}

func (m *resourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel) flattenEndpointConnectionProfilesSecureInternetAccessPostureCheck(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["tag"]; ok {
		m.Tag = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["checkFailedMessage"]; ok {
		m.CheckFailedMessage = parseStringValue(v)
	}

	return m
}

func (m *resourceEndpointConnectionProfilesAvailableVpNsModel) flattenEndpointConnectionProfilesAvailableVpNs(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesAvailableVpNsModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesAvailableVpNsModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesAvailableVpNsModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["name"]; ok {
		m.Name = parseStringValue(v)
	}

	if v, ok := o["remoteGateway"]; ok {
		m.RemoteGateway = parseStringValue(v)
	}

	if v, ok := o["usernamePrompt"]; ok {
		m.UsernamePrompt = parseStringValue(v)
	}

	if v, ok := o["saveUsername"]; ok {
		m.SaveUsername = parseStringValue(v)
	}

	if v, ok := o["showAlwaysUp"]; ok {
		m.ShowAlwaysUp = parseStringValue(v)
	}

	if v, ok := o["showAutoConnect"]; ok {
		m.ShowAutoConnect = parseStringValue(v)
	}

	if v, ok := o["showRememberPassword"]; ok {
		m.ShowRememberPassword = parseStringValue(v)
	}

	if v, ok := o["authenticateWithSSO"]; ok {
		m.AuthenticateWithSso = parseStringValue(v)
	}

	if v, ok := o["enableLocalLan"]; ok {
		m.EnableLocalLan = parseStringValue(v)
	}

	if v, ok := o["externalBrowserSamlLogin"]; ok {
		m.ExternalBrowserSamlLogin = parseStringValue(v)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	if v, ok := o["requireCertificate"]; ok {
		m.RequireCertificate = parseStringValue(v)
	}

	if v, ok := o["authMethod"]; ok {
		m.AuthMethod = parseStringValue(v)
	}

	if v, ok := o["showPasscode"]; ok {
		m.ShowPasscode = parseStringValue(v)
	}

	if v, ok := o["postureCheck"]; ok {
		m.PostureCheck = m.PostureCheck.flattenEndpointConnectionProfilesAvailableVpNsPostureCheck(ctx, v, diags)
	}

	if v, ok := o["preSharedKey"]; ok {
		m.PreSharedKey = parseStringValue(v)
	}

	return m
}

func (s *resourceEndpointConnectionProfilesModel) flattenEndpointConnectionProfilesAvailableVpNsList(ctx context.Context, o interface{}, diags *diag.Diagnostics) []resourceEndpointConnectionProfilesAvailableVpNsModel {
	if o == nil {
		return []resourceEndpointConnectionProfilesAvailableVpNsModel{}
	}

	if _, ok := o.([]interface{}); !ok {
		diags.AddError("Argument available_vp_ns is not type of []interface{}.", "")
		return []resourceEndpointConnectionProfilesAvailableVpNsModel{}
	}

	l := o.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return []resourceEndpointConnectionProfilesAvailableVpNsModel{}
	}

	values := make([]resourceEndpointConnectionProfilesAvailableVpNsModel, len(l))
	for i, ele := range l {
		var m resourceEndpointConnectionProfilesAvailableVpNsModel
		values[i] = *m.flattenEndpointConnectionProfilesAvailableVpNs(ctx, ele, diags)
	}

	return values
}

func (m *resourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel) flattenEndpointConnectionProfilesAvailableVpNsPostureCheck(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["tag"]; ok {
		m.Tag = parseStringValue(v)
	}

	if v, ok := o["action"]; ok {
		m.Action = parseStringValue(v)
	}

	if v, ok := o["checkFailedMessage"]; ok {
		m.CheckFailedMessage = parseStringValue(v)
	}

	return m
}

func (m *resourceEndpointConnectionProfilesPreLogonModel) flattenEndpointConnectionProfilesPreLogon(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesPreLogonModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesPreLogonModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesPreLogonModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["vpnType"]; ok {
		m.VpnType = parseStringValue(v)
	}

	if v, ok := o["remoteGateway"]; ok {
		m.RemoteGateway = parseStringValue(v)
	}

	if v, ok := o["commonName"]; ok {
		m.CommonName = m.CommonName.flattenEndpointConnectionProfilesPreLogonCommonName(ctx, v, diags)
	}

	if v, ok := o["issuer"]; ok {
		m.Issuer = m.Issuer.flattenEndpointConnectionProfilesPreLogonIssuer(ctx, v, diags)
	}

	if v, ok := o["port"]; ok {
		m.Port = parseFloat64Value(v)
	}

	return m
}

func (m *resourceEndpointConnectionProfilesPreLogonCommonNameModel) flattenEndpointConnectionProfilesPreLogonCommonName(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesPreLogonCommonNameModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesPreLogonCommonNameModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesPreLogonCommonNameModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["matchType"]; ok {
		m.MatchType = parseStringValue(v)
	}

	if v, ok := o["pattern"]; ok {
		m.Pattern = parseStringValue(v)
	}

	return m
}

func (m *resourceEndpointConnectionProfilesPreLogonIssuerModel) flattenEndpointConnectionProfilesPreLogonIssuer(ctx context.Context, input interface{}, diags *diag.Diagnostics) *resourceEndpointConnectionProfilesPreLogonIssuerModel {
	if input == nil {
		return &resourceEndpointConnectionProfilesPreLogonIssuerModel{}
	}
	if m == nil {
		m = &resourceEndpointConnectionProfilesPreLogonIssuerModel{}
	}
	o := input.(map[string]interface{})
	if v, ok := o["matchType"]; ok {
		m.MatchType = parseStringValue(v)
	}

	if v, ok := o["pattern"]; ok {
		m.Pattern = parseStringValue(v)
	}

	return m
}

func (data *resourceEndpointConnectionProfilesLockdownModel) expandEndpointConnectionProfilesLockdown(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.GracePeriod.IsNull() {
		result["gracePeriod"] = data.GracePeriod.ValueFloat64()
	}

	if !data.MaxAttempts.IsNull() {
		result["maxAttempts"] = data.MaxAttempts.ValueFloat64()
	}

	if len(data.Ips) > 0 {
		result["ips"] = data.expandEndpointConnectionProfilesLockdownIpsList(ctx, data.Ips, diags)
	}

	if len(data.Domains) > 0 {
		result["domains"] = data.expandEndpointConnectionProfilesLockdownDomainsList(ctx, data.Domains, diags)
	}

	if data.DetectCaptivePortal != nil && !isZeroStruct(*data.DetectCaptivePortal) {
		result["detectCaptivePortal"] = data.DetectCaptivePortal.expandEndpointConnectionProfilesLockdownDetectCaptivePortal(ctx, diags)
	}

	return result
}

func (data *resourceEndpointConnectionProfilesLockdownIpsModel) expandEndpointConnectionProfilesLockdownIps(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Ip.IsNull() {
		result["ip"] = data.Ip.ValueString()
	}

	if !data.Port.IsNull() {
		result["port"] = data.Port.ValueString()
	}

	if !data.Protocol.IsNull() {
		result["protocol"] = data.Protocol.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfilesLockdownModel) expandEndpointConnectionProfilesLockdownIpsList(ctx context.Context, l []resourceEndpointConnectionProfilesLockdownIpsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfilesLockdownIps(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfilesLockdownDomainsModel) expandEndpointConnectionProfilesLockdownDomains(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Address.IsNull() {
		result["address"] = data.Address.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfilesLockdownModel) expandEndpointConnectionProfilesLockdownDomainsList(ctx context.Context, l []resourceEndpointConnectionProfilesLockdownDomainsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfilesLockdownDomains(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfilesLockdownDetectCaptivePortalModel) expandEndpointConnectionProfilesLockdownDetectCaptivePortal(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfilesOnFabricRuleSetModel) expandEndpointConnectionProfilesOnFabricRuleSet(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfilesOffNetSplitTunnelModel) expandEndpointConnectionProfilesOffNetSplitTunnel(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.LocalApps.IsNull() {
		result["localApps"] = expandSetToStringList(data.LocalApps)
	}

	result["isdbs"] = data.expandEndpointConnectionProfilesOffNetSplitTunnelIsdbsList(ctx, data.Isdbs, diags)

	if !data.Fqdns.IsNull() {
		result["fqdns"] = expandSetToStringList(data.Fqdns)
	}

	if len(data.Subnets) > 0 {
		result["subnets"] = data.expandEndpointConnectionProfilesOffNetSplitTunnelSubnetsList(ctx, data.Subnets, diags)
	}

	if !data.SubnetsIpsec.IsNull() {
		result["subnetsIpsec"] = expandSetToStringList(data.SubnetsIpsec)
	}

	return result
}

func (data *resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel) expandEndpointConnectionProfilesOffNetSplitTunnelIsdbs(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfilesOffNetSplitTunnelModel) expandEndpointConnectionProfilesOffNetSplitTunnelIsdbsList(ctx context.Context, l []resourceEndpointConnectionProfilesOffNetSplitTunnelIsdbsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfilesOffNetSplitTunnelIsdbs(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel) expandEndpointConnectionProfilesOffNetSplitTunnelSubnets(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfilesOffNetSplitTunnelModel) expandEndpointConnectionProfilesOffNetSplitTunnelSubnetsList(ctx context.Context, l []resourceEndpointConnectionProfilesOffNetSplitTunnelSubnetsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfilesOffNetSplitTunnelSubnets(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfilesSplitTunnelModel) expandEndpointConnectionProfilesSplitTunnel(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.LocalApps.IsNull() {
		result["localApps"] = expandSetToStringList(data.LocalApps)
	}

	result["isdbs"] = data.expandEndpointConnectionProfilesSplitTunnelIsdbsList(ctx, data.Isdbs, diags)

	if !data.Fqdns.IsNull() {
		result["fqdns"] = expandSetToStringList(data.Fqdns)
	}

	if len(data.Subnets) > 0 {
		result["subnets"] = data.expandEndpointConnectionProfilesSplitTunnelSubnetsList(ctx, data.Subnets, diags)
	}

	if !data.SubnetsIpsec.IsNull() {
		result["subnetsIpsec"] = expandSetToStringList(data.SubnetsIpsec)
	}

	return result
}

func (data *resourceEndpointConnectionProfilesSplitTunnelIsdbsModel) expandEndpointConnectionProfilesSplitTunnelIsdbs(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfilesSplitTunnelModel) expandEndpointConnectionProfilesSplitTunnelIsdbsList(ctx context.Context, l []resourceEndpointConnectionProfilesSplitTunnelIsdbsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfilesSplitTunnelIsdbs(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfilesSplitTunnelSubnetsModel) expandEndpointConnectionProfilesSplitTunnelSubnets(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Datasource.IsNull() {
		result["datasource"] = data.Datasource.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfilesSplitTunnelModel) expandEndpointConnectionProfilesSplitTunnelSubnetsList(ctx context.Context, l []resourceEndpointConnectionProfilesSplitTunnelSubnetsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfilesSplitTunnelSubnets(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfilesSecureInternetAccessModel) expandEndpointConnectionProfilesSecureInternetAccess(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.AuthenticateWithSso.IsNull() {
		result["authenticateWithSSO"] = data.AuthenticateWithSso.ValueString()
	}

	if !data.EnableLocalLan.IsNull() {
		result["enableLocalLan"] = data.EnableLocalLan.ValueString()
	}

	if !data.FailoverSequence.IsNull() {
		result["failoverSequence"] = expandSetToStringList(data.FailoverSequence)
	}

	if data.PostureCheck != nil && !isZeroStruct(*data.PostureCheck) {
		result["postureCheck"] = data.PostureCheck.expandEndpointConnectionProfilesSecureInternetAccessPostureCheck(ctx, diags)
	}

	if !data.ExternalBrowserSamlLogin.IsNull() {
		result["externalBrowserSamlLogin"] = data.ExternalBrowserSamlLogin.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfilesSecureInternetAccessPostureCheckModel) expandEndpointConnectionProfilesSecureInternetAccessPostureCheck(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Tag.IsNull() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.CheckFailedMessage.IsNull() {
		result["checkFailedMessage"] = data.CheckFailedMessage.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfilesAvailableVpNsModel) expandEndpointConnectionProfilesAvailableVpNs(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Type.IsNull() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Name.IsNull() {
		result["name"] = data.Name.ValueString()
	}

	if !data.RemoteGateway.IsNull() {
		result["remoteGateway"] = data.RemoteGateway.ValueString()
	}

	if !data.UsernamePrompt.IsNull() {
		result["usernamePrompt"] = data.UsernamePrompt.ValueString()
	}

	if !data.SaveUsername.IsNull() {
		result["saveUsername"] = data.SaveUsername.ValueString()
	}

	if !data.ShowAlwaysUp.IsNull() {
		result["showAlwaysUp"] = data.ShowAlwaysUp.ValueString()
	}

	if !data.ShowAutoConnect.IsNull() {
		result["showAutoConnect"] = data.ShowAutoConnect.ValueString()
	}

	if !data.ShowRememberPassword.IsNull() {
		result["showRememberPassword"] = data.ShowRememberPassword.ValueString()
	}

	if !data.AuthenticateWithSso.IsNull() {
		result["authenticateWithSSO"] = data.AuthenticateWithSso.ValueString()
	}

	if !data.EnableLocalLan.IsNull() {
		result["enableLocalLan"] = data.EnableLocalLan.ValueString()
	}

	if !data.ExternalBrowserSamlLogin.IsNull() {
		result["externalBrowserSamlLogin"] = data.ExternalBrowserSamlLogin.ValueString()
	}

	if !data.Port.IsNull() {
		result["port"] = data.Port.ValueFloat64()
	}

	if !data.RequireCertificate.IsNull() {
		result["requireCertificate"] = data.RequireCertificate.ValueString()
	}

	if !data.AuthMethod.IsNull() {
		result["authMethod"] = data.AuthMethod.ValueString()
	}

	if !data.ShowPasscode.IsNull() {
		result["showPasscode"] = data.ShowPasscode.ValueString()
	}

	if data.PostureCheck != nil && !isZeroStruct(*data.PostureCheck) {
		result["postureCheck"] = data.PostureCheck.expandEndpointConnectionProfilesAvailableVpNsPostureCheck(ctx, diags)
	}

	if !data.PreSharedKey.IsNull() {
		result["preSharedKey"] = data.PreSharedKey.ValueString()
	}

	return result
}

func (s *resourceEndpointConnectionProfilesModel) expandEndpointConnectionProfilesAvailableVpNsList(ctx context.Context, l []resourceEndpointConnectionProfilesAvailableVpNsModel, diags *diag.Diagnostics) []map[string]interface{} {
	result := make([]map[string]interface{}, len(l))
	for i, item := range l {
		result[i] = item.expandEndpointConnectionProfilesAvailableVpNs(ctx, diags)
	}
	return result
}

func (data *resourceEndpointConnectionProfilesAvailableVpNsPostureCheckModel) expandEndpointConnectionProfilesAvailableVpNsPostureCheck(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Tag.IsNull() {
		result["tag"] = data.Tag.ValueString()
	}

	if !data.Action.IsNull() {
		result["action"] = data.Action.ValueString()
	}

	if !data.CheckFailedMessage.IsNull() {
		result["checkFailedMessage"] = data.CheckFailedMessage.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfilesPreLogonModel) expandEndpointConnectionProfilesPreLogon(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.VpnType.IsNull() {
		result["vpnType"] = data.VpnType.ValueString()
	}

	if !data.RemoteGateway.IsNull() {
		result["remoteGateway"] = data.RemoteGateway.ValueString()
	}

	if data.CommonName != nil && !isZeroStruct(*data.CommonName) {
		result["commonName"] = data.CommonName.expandEndpointConnectionProfilesPreLogonCommonName(ctx, diags)
	}

	if data.Issuer != nil && !isZeroStruct(*data.Issuer) {
		result["issuer"] = data.Issuer.expandEndpointConnectionProfilesPreLogonIssuer(ctx, diags)
	}

	if !data.Port.IsNull() {
		result["port"] = data.Port.ValueFloat64()
	}

	return result
}

func (data *resourceEndpointConnectionProfilesPreLogonCommonNameModel) expandEndpointConnectionProfilesPreLogonCommonName(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.MatchType.IsNull() {
		result["matchType"] = data.MatchType.ValueString()
	}

	if !data.Pattern.IsNull() {
		result["pattern"] = data.Pattern.ValueString()
	}

	return result
}

func (data *resourceEndpointConnectionProfilesPreLogonIssuerModel) expandEndpointConnectionProfilesPreLogonIssuer(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	result := make(map[string]interface{})
	if !data.MatchType.IsNull() {
		result["matchType"] = data.MatchType.ValueString()
	}

	if !data.Pattern.IsNull() {
		result["pattern"] = data.Pattern.ValueString()
	}

	return result
}
