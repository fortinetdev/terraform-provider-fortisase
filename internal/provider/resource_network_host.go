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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourceNetworkHost{}
var _ resource.ResourceWithMoveState = &resourceNetworkHost{}

func newResourceNetworkHost() resource.Resource {
	return &resourceNetworkHost{}
}

type resourceNetworkHost struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceNetworkHostModel describes the resource data model.
type resourceNetworkHostModel struct {
	ID         types.String `tfsdk:"id"`
	PrimaryKey types.String `tfsdk:"primary_key"`
	Type       types.String `tfsdk:"type"`
	Location   types.String `tfsdk:"location"`
	Subnet     types.String `tfsdk:"subnet"`
	StartIp    types.String `tfsdk:"start_ip"`
	EndIp      types.String `tfsdk:"end_ip"`
	Fqdn       types.String `tfsdk:"fqdn"`
	CountryId  types.String `tfsdk:"country_id"`
}

func (r *resourceNetworkHost) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_host"
}

func (r *resourceNetworkHost) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Host Resource API V2 for FortiSASE.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"primary_key": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 79),
				},
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("ipmask", "iprange", "fqdn", "geography"),
				},
				Computed: true,
				Optional: true,
			},
			"location": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.OneOf("internal", "external", "private-access", "unspecified"),
				},
				Computed: true,
				Optional: true,
			},
			"subnet": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"start_ip": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"end_ip": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"fqdn": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(1, 255),
				},
				Computed: true,
				Optional: true,
			},
			"country_id": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidatorwarning.LengthBetween(2, 2),
				},
				Computed: true,
				Optional: true,
			},
		},
	}
}

func (r *resourceNetworkHost) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_network_host"
}
func (r *resourceNetworkHost) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_network_hosts" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourceNetworkHostModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourceNetworkHost) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkHosts")
	lock.Lock()
	defer lock.Unlock()
	var data resourceNetworkHostModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectNetworkHost(ctx, diags))
	input_model.URLParams = *(data.getURLObjectNetworkHost(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	mkey := data.PrimaryKey.ValueString()
	output, err := c.CreateNetworkHosts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	if responseMkey, ok := getCreateResponseMkey(output, "primaryKey"); ok {
		mkey = responseMkey
	}
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectNetworkHost(ctx, "read", diags))

	read_output, err := c.ReadNetworkHosts(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshNetworkHost(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkHost) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkHosts")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceNetworkHostModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceNetworkHostModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectNetworkHost(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectNetworkHost(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateNetworkHosts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectNetworkHost(ctx, "read", diags))

	read_output, err := c.ReadNetworkHosts(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshNetworkHost(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkHost) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("NetworkHosts")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceNetworkHostModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkHost(ctx, "delete", diags))

	output, err := c.DeleteNetworkHosts(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceNetworkHost) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceNetworkHostModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectNetworkHost(ctx, "read", diags))

	read_output, err := c.ReadNetworkHosts(&input_model)
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

	diags.Append(data.refreshNetworkHost(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceNetworkHost) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("primary_key"), req.ID)...)
}

func (m *resourceNetworkHostModel) refreshNetworkHost(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["type"]; ok {
		m.Type = parseStringValue(v)
	}

	if v, ok := o["location"]; ok {
		m.Location = parseStringValue(v)
	}

	if v, ok := o["subnet"]; ok {
		m.Subnet = parseStringValue(v)
	}

	if v, ok := o["startIp"]; ok {
		m.StartIp = parseStringValue(v)
	}

	if v, ok := o["endIp"]; ok {
		m.EndIp = parseStringValue(v)
	}

	if v, ok := o["fqdn"]; ok {
		m.Fqdn = parseStringValue(v)
	}

	if v, ok := o["countryId"]; ok {
		m.CountryId = parseStringValue(v)
	}

	return diags
}

func (data *resourceNetworkHostModel) getCreateObjectNetworkHost(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Location.IsNull() && !data.Location.IsUnknown() {
		result["location"] = data.Location.ValueString()
	}

	if !data.Subnet.IsNull() && !data.Subnet.IsUnknown() {
		result["subnet"] = data.Subnet.ValueString()
	}

	if !data.StartIp.IsNull() && !data.StartIp.IsUnknown() {
		result["startIp"] = data.StartIp.ValueString()
	}

	if !data.EndIp.IsNull() && !data.EndIp.IsUnknown() {
		result["endIp"] = data.EndIp.ValueString()
	}

	if !data.Fqdn.IsNull() && !data.Fqdn.IsUnknown() {
		result["fqdn"] = data.Fqdn.ValueString()
	}

	if !data.CountryId.IsNull() && !data.CountryId.IsUnknown() {
		result["countryId"] = data.CountryId.ValueString()
	}

	return &result
}

func (data *resourceNetworkHostModel) getUpdateObjectNetworkHost(ctx context.Context, state resourceNetworkHostModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		result["type"] = data.Type.ValueString()
	}

	if !data.Location.IsNull() && !data.Location.IsUnknown() {
		result["location"] = data.Location.ValueString()
	}

	if !data.Subnet.IsNull() && !data.Subnet.IsUnknown() {
		result["subnet"] = data.Subnet.ValueString()
	}

	if !data.StartIp.IsNull() && !data.StartIp.IsUnknown() {
		result["startIp"] = data.StartIp.ValueString()
	}

	if !data.EndIp.IsNull() && !data.EndIp.IsUnknown() {
		result["endIp"] = data.EndIp.ValueString()
	}

	if !data.Fqdn.IsNull() && !data.Fqdn.IsUnknown() {
		result["fqdn"] = data.Fqdn.ValueString()
	}

	if !data.CountryId.IsNull() && !data.CountryId.IsUnknown() {
		result["countryId"] = data.CountryId.ValueString()
	}

	return &result
}

func (data *resourceNetworkHostModel) getURLObjectNetworkHost(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() && !data.PrimaryKey.IsUnknown() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
