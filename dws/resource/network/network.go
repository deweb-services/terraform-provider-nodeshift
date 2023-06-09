package network

import (
	"context"
	"fmt"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &networkResource{}
	_ resource.ResourceWithConfigure   = &networkResource{}
	_ resource.ResourceWithImportState = &networkResource{}
)

type networkResource struct {
	client *client.DWSClient
}

// Metadata returns the resource type name.
func (r *networkResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network"
}

func (r *networkResource) Schema(c context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			ID: schema.StringAttribute{
				Computed: true,
			},
			NetworkIPRangeKeys: schema.StringAttribute{
				Required: true,
			},
			NetworkNameKeys: schema.StringAttribute{
				Required: true,
			},
			NetworkDescriptionKeys: schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *networkResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.DWSClient)
}

// Create creates the resource and sets the initial Terraform state.
func (r *networkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan networkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Create new Network
	clientRequest, err := plan.ToClientRequest()
	if err != nil {
		tflog.Error(ctx, "failed to convert resource to client required type", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}

	network, err := r.client.CreateNetwork(ctx, clientRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating network",
			fmt.Sprintf("Could not create network, unexpected error: %s", err),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.FromClientResponse(network)
	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *networkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state networkResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Get refreshed order value from client
	network, err := r.client.GetNetwork(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading network state",
			fmt.Sprintf("Could not read network state ID %s: %s", state.ID.ValueString(), err),
		)
		return
	}

	// Overwrite items with refreshed state
	state.FromClientResponse(network)
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *networkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan networkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	clientRequest, err := plan.ToClientRequest()
	if err != nil {
		tflog.Error(ctx, "failed to convert resource to client required type", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}

	// Update existing order
	_, err = r.client.UpdateNetwork(ctx, plan.ID.ValueString(), clientRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating network state",
			fmt.Sprintf("Could not update network state %s, unexpected error: %s", plan.ID.ValueString(), err),
		)
		return
	}

	// Fetch updated items from GetNetwork as UpdateNetwork items are not populated.
	network, err := r.client.GetNetwork(ctx, plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading network state",
			fmt.Sprintf("Could not read network name %s: %s", plan.ID.ValueString(), err),
		)
		return
	}

	plan.FromClientResponse(network)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *networkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state networkResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Delete existing network
	if err := r.client.DeleteNetwork(ctx, state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting network",
			fmt.Sprintf("Could not delete network, unexpected error: %s", err),
		)
		return
	}
}

func (r *networkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root(ID), req, resp)
}
