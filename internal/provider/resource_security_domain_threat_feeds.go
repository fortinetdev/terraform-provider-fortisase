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
var _ resource.Resource = &resourceSecurityDomainThreatFeeds{}

func newResourceSecurityDomainThreatFeeds() resource.Resource {
	return &resourceSecurityDomainThreatFeeds{}
}

type resourceSecurityDomainThreatFeeds struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourceSecurityDomainThreatFeedsModel describes the resource data model.
type resourceSecurityDomainThreatFeedsModel struct {
	ID                  types.String  `tfsdk:"id"`
	PrimaryKey          types.String  `tfsdk:"primary_key"`
	Comments            types.String  `tfsdk:"comments"`
	Status              types.String  `tfsdk:"status"`
	RefreshRate         types.Float64 `tfsdk:"refresh_rate"`
	Uri                 types.String  `tfsdk:"uri"`
	BasicAuthentication types.String  `tfsdk:"basic_authentication"`
	Username            types.String  `tfsdk:"username"`
	Password            types.String  `tfsdk:"password"`
}

func (r *resourceSecurityDomainThreatFeeds) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_domain_threat_feeds"
}

func (r *resourceSecurityDomainThreatFeeds) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
					stringvalidator.LengthBetween(1, 35),
				},
				Required: true,
			},
			"comments": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
				Computed: true,
				Optional: true,
			},
			"status": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"refresh_rate": schema.Float64Attribute{
				Validators: []validator.Float64{
					float64validator.Between(1, 43200),
				},
				Computed: true,
				Optional: true,
			},
			"uri": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
				Computed: true,
				Optional: true,
			},
			"basic_authentication": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
				Computed: true,
				Optional: true,
			},
			"username": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
				Computed: true,
				Optional: true,
			},
			"password": schema.StringAttribute{
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 128),
				},
				Sensitive: true,
				Computed:  true,
				Optional:  true,
			},
		},
	}
}

func (r *resourceSecurityDomainThreatFeeds) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_security_domain_threat_feeds"
}

func (r *resourceSecurityDomainThreatFeeds) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDomainThreatFeeds")
	lock.Lock()
	defer lock.Unlock()
	var data resourceSecurityDomainThreatFeedsModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectSecurityDomainThreatFeeds(ctx, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDomainThreatFeeds(ctx, "create", diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreateSecurityDomainThreatFeeds(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}

	mkey := fmt.Sprintf("%v", output["primaryKey"])
	data.ID = types.StringValue(mkey)
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDomainThreatFeeds(ctx, "read", diags))

	read_output, err := c.ReadSecurityDomainThreatFeeds(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDomainThreatFeeds(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDomainThreatFeeds) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDomainThreatFeeds")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourceSecurityDomainThreatFeedsModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourceSecurityDomainThreatFeedsModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectSecurityDomainThreatFeeds(ctx, state, diags))
	input_model.URLParams = *(data.getURLObjectSecurityDomainThreatFeeds(ctx, "update", diags))

	if diags.HasError() {
		return
	}

	output, err := c.UpdateSecurityDomainThreatFeeds(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
	var read_input_model forticlient.InputModel
	read_input_model.Mkey = mkey
	read_input_model.URLParams = *(data.getURLObjectSecurityDomainThreatFeeds(ctx, "read", diags))

	read_output, err := c.ReadSecurityDomainThreatFeeds(&read_input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&read_input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDomainThreatFeeds(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDomainThreatFeeds) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	lock := r.fortiClient.GetResourceLock("SecurityDomainThreatFeeds")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics
	var data resourceSecurityDomainThreatFeedsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDomainThreatFeeds(ctx, "delete", diags))

	output, err := c.DeleteSecurityDomainThreatFeeds(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to delete resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, output),
		)
		return
	}
}

func (r *resourceSecurityDomainThreatFeeds) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	diags := &resp.Diagnostics
	var data resourceSecurityDomainThreatFeedsModel

	// Read Terraform prior state data into the model
	diags.Append(req.State.Get(ctx, &data)...)

	if diags.HasError() {
		return
	}

	mkey := data.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.URLParams = *(data.getURLObjectSecurityDomainThreatFeeds(ctx, "read", diags))

	read_output, err := c.ReadSecurityDomainThreatFeeds(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to read resource %s: %v", r.resourceName, err),
			getErrorDetail(&input_model, read_output),
		)
		return
	}

	diags.Append(data.refreshSecurityDomainThreatFeeds(ctx, read_output)...)
	if diags.HasError() {
		return
	}

	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourceSecurityDomainThreatFeeds) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *resourceSecurityDomainThreatFeedsModel) refreshSecurityDomainThreatFeeds(ctx context.Context, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	if o == nil {
		return diags
	}

	if v, ok := o["comments"]; ok {
		m.Comments = parseStringValue(v)
	}

	if v, ok := o["status"]; ok {
		m.Status = parseStringValue(v)
	}

	if v, ok := o["refreshRate"]; ok {
		m.RefreshRate = parseFloat64Value(v)
	}

	if v, ok := o["uri"]; ok {
		m.Uri = parseStringValue(v)
	}

	if v, ok := o["basicAuthentication"]; ok {
		m.BasicAuthentication = parseStringValue(v)
	}

	if v, ok := o["username"]; ok {
		m.Username = parseStringValue(v)
	}

	return diags
}

func (data *resourceSecurityDomainThreatFeedsModel) getCreateObjectSecurityDomainThreatFeeds(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Comments.IsNull() {
		result["comments"] = data.Comments.ValueString()
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.RefreshRate.IsNull() {
		result["refreshRate"] = data.RefreshRate.ValueFloat64()
	}

	if !data.Uri.IsNull() {
		result["uri"] = data.Uri.ValueString()
	}

	if !data.BasicAuthentication.IsNull() {
		result["basicAuthentication"] = data.BasicAuthentication.ValueString()
	}

	if !data.Username.IsNull() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	return &result
}

func (data *resourceSecurityDomainThreatFeedsModel) getUpdateObjectSecurityDomainThreatFeeds(ctx context.Context, state resourceSecurityDomainThreatFeedsModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	if !data.Comments.IsNull() {
		result["comments"] = data.Comments.ValueString()
	}

	if !data.Status.IsNull() {
		result["status"] = data.Status.ValueString()
	}

	if !data.RefreshRate.IsNull() {
		result["refreshRate"] = data.RefreshRate.ValueFloat64()
	}

	if !data.Uri.IsNull() {
		result["uri"] = data.Uri.ValueString()
	}

	if !data.BasicAuthentication.IsNull() {
		result["basicAuthentication"] = data.BasicAuthentication.ValueString()
	}

	if !data.Username.IsNull() {
		result["username"] = data.Username.ValueString()
	}

	if !data.Password.IsNull() {
		result["password"] = data.Password.ValueString()
	}

	return &result
}

func (data *resourceSecurityDomainThreatFeedsModel) getURLObjectSecurityDomainThreatFeeds(ctx context.Context, ope string, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.PrimaryKey.IsNull() {
		result["primaryKey"] = data.PrimaryKey.ValueString()
	}

	return &result
}
