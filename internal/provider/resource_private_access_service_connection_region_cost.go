// Copyright 2020 Fortinet, Inc. All rights reserved.
package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	forticlient "github.com/fortinetdev/terraform-provider-fortisase/internal/sdk/sdkcore"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &resourcePrivateAccessServiceConnectionRegionCost2Edl{}
var _ resource.ResourceWithMoveState = &resourcePrivateAccessServiceConnectionRegionCost2Edl{}

func newResourcePrivateAccessServiceConnectionRegionCost() resource.Resource {
	return &resourcePrivateAccessServiceConnectionRegionCost2Edl{}
}

type resourcePrivateAccessServiceConnectionRegionCost2Edl struct {
	fortiClient  *FortiClient
	resourceName string
}

// resourcePrivateAccessServiceConnectionRegionCost2EdlModel describes the resource data model.
type resourcePrivateAccessServiceConnectionRegionCost2EdlModel struct {
	ID      types.String `tfsdk:"id"`
	Entries types.Map    `tfsdk:"entries"`
}

func (r *resourcePrivateAccessServiceConnectionRegionCost2Edl) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_private_access_service_connection_region_cost"
}

func (r *resourcePrivateAccessServiceConnectionRegionCost2Edl) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identifier, required by Terraform, not configurable.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"entries": schema.MapAttribute{
				Description: "Arbitrary regions map. Key is string; value is a map of key:integer.",
				ElementType: types.MapType{ElemType: types.Int64Type},
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func (r *resourcePrivateAccessServiceConnectionRegionCost2Edl) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.resourceName = "fortisase_private_access_service_connection_region_cost"
}

func (r *resourcePrivateAccessServiceConnectionRegionCost2Edl) MoveState(ctx context.Context) []resource.StateMover {
	schemaResponse := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResponse)
	sourceSchema := &schemaResponse.Schema

	return []resource.StateMover{
		{
			SourceSchema: sourceSchema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "fortisase_private_access_service_connections_region_cost" || req.SourceSchemaVersion != 0 {
					return
				}

				var sourceState resourcePrivateAccessServiceConnectionRegionCost2EdlModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &sourceState)...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, &sourceState)...)
			},
		},
	}
}

func (r *resourcePrivateAccessServiceConnectionRegionCost2Edl) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Use the same lock as the resource_private_access_service_connection
	lock := r.fortiClient.GetResourceLock("PrivateAccessServiceConnections")
	lock.Lock()
	defer lock.Unlock()
	var data resourcePrivateAccessServiceConnectionRegionCost2EdlModel
	diags := &resp.Diagnostics

	// Read Terraform config data into the model
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.BodyParams = *(data.getCreateObjectPrivateAccessServiceConnectionRegionCost(ctx, diags))

	if diags.HasError() {
		return
	}
	output, err := c.CreatePrivateAccessServiceConnectionsRegionCost(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to create resource: %v", err),
			"",
		)
		return
	}

	mkey := "PrivateAccessServiceConnectionsRegionCost"
	data.ID = types.StringValue(mkey)

	var targetItems []string
	outputHubs, ok := output["hubs"].([]interface{})
	if !ok {
		outputHubs = []interface{}{}
	}
	for _, item := range outputHubs {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if id, exists := itemMap["id"]; exists {
				if idStr, ok := id.(string); ok {
					targetItems = append(targetItems, idStr)
				}
			}
		}
	}

	var pendingItems []string
	for i := 0; i < 20; i++ {
		time.Sleep(10 * time.Second)
		for _, item := range targetItems {
			var read_input_model forticlient.InputModel
			read_input_model.Mkey = item
			read_input_model.URLParams = map[string]interface{}{
				"service-connection-id": item,
			}
			read_output, err := c.ReadPrivateAccessServiceConnections(&read_input_model)
			if err != nil {
				diags.AddError(
					fmt.Sprintf("Error to read resource: %v", err),
					"",
				)
				return
			}
			if v, ok := read_output["config_state"]; ok {
				if configState, ok := v.(string); ok {
					if configState != "success" {
						pendingItems = append(pendingItems, item)
					}
				}
			}
		}
		targetItems = pendingItems
		if len(pendingItems) == 0 {
			break
		}
		pendingItems = []string{}
	}
	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnectionRegionCost2Edl) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Use the same lock as the resource_private_access_service_connection
	lock := r.fortiClient.GetResourceLock("PrivateAccessServiceConnections")
	lock.Lock()
	defer lock.Unlock()
	diags := &resp.Diagnostics

	// Read Terraform plan data into the model
	var state resourcePrivateAccessServiceConnectionRegionCost2EdlModel
	diags.Append(req.State.Get(ctx, &state)...)
	if diags.HasError() {
		return
	}

	var data resourcePrivateAccessServiceConnectionRegionCost2EdlModel
	diags.Append(req.Config.Get(ctx, &data)...)
	if diags.HasError() {
		return
	}
	data.ID = state.ID

	mkey := state.ID.ValueString()

	c := r.fortiClient.Client
	var input_model forticlient.InputModel
	input_model.Mkey = mkey
	input_model.BodyParams = *(data.getUpdateObjectPrivateAccessServiceConnectionRegionCost(ctx, state, diags))

	if diags.HasError() {
		return
	}

	output, err := c.CreatePrivateAccessServiceConnectionsRegionCost(&input_model)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Error to update resource: %v", err),
			"",
		)
		return
	}

	var targetItems []string
	outputHubs, ok := output["hubs"].([]interface{})
	if !ok {
		outputHubs = []interface{}{}
	}
	for _, item := range outputHubs {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if id, exists := itemMap["id"]; exists {
				if idStr, ok := id.(string); ok {
					targetItems = append(targetItems, idStr)
				}
			}
		}
	}

	var pendingItems []string
	for i := 0; i < 20; i++ {
		time.Sleep(10 * time.Second)
		for _, item := range targetItems {
			var read_input_model forticlient.InputModel
			read_input_model.Mkey = item
			read_input_model.URLParams = map[string]interface{}{
				"service-connection-id": item,
			}
			read_output, err := c.ReadPrivateAccessServiceConnections(&read_input_model)
			if err != nil {
				diags.AddError(
					fmt.Sprintf("Error to read resource: %v", err),
					"",
				)
				return
			}
			if v, ok := read_output["config_state"]; ok {
				if configState, ok := v.(string); ok {
					if configState != "success" {
						pendingItems = append(pendingItems, item)
					}
				}
			}
		}
		targetItems = pendingItems
		if len(pendingItems) == 0 {
			break
		}
		pendingItems = []string{}
	}
	diags.Append(resp.State.Set(ctx, &data)...)
}

func (r *resourcePrivateAccessServiceConnectionRegionCost2Edl) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// No delete operation for this resource
}

func (r *resourcePrivateAccessServiceConnectionRegionCost2Edl) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No read operation for this resource
}

func (data *resourcePrivateAccessServiceConnectionRegionCost2EdlModel) getCreateObjectPrivateAccessServiceConnectionRegionCost(ctx context.Context, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Entries.IsNull() {
		var entries map[string]types.Map
		diags.Append(data.Entries.ElementsAs(ctx, &entries, false)...)
		if diags.HasError() {
			return nil
		}
		for k, v := range entries {
			var inner map[string]int64
			diags.Append(v.ElementsAs(ctx, &inner, false)...)
			if diags.HasError() {
				return nil
			}
			hyphenKey := strings.ReplaceAll(k, "_", "-")
			result[hyphenKey] = inner
		}
	}

	return &result
}

func (data *resourcePrivateAccessServiceConnectionRegionCost2EdlModel) getUpdateObjectPrivateAccessServiceConnectionRegionCost(ctx context.Context, state resourcePrivateAccessServiceConnectionRegionCost2EdlModel, diags *diag.Diagnostics) *map[string]interface{} {
	result := make(map[string]interface{})
	if !data.Entries.IsNull() {
		var entries map[string]types.Map
		diags.Append(data.Entries.ElementsAs(ctx, &entries, false)...)
		if diags.HasError() {
			return nil
		}
		for k, v := range entries {
			var inner map[string]int64
			diags.Append(v.ElementsAs(ctx, &inner, false)...)
			if diags.HasError() {
				return nil
			}
			hyphenKey := strings.ReplaceAll(k, "_", "-")
			result[hyphenKey] = inner
		}
	}

	return &result
}
