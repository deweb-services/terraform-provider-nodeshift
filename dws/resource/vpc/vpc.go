package vpc

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
	_ resource.Resource                = &vpcResource{}
	_ resource.ResourceWithConfigure   = &vpcResource{}
	_ resource.ResourceWithImportState = &vpcResource{}
)

// NewDeploymentResource is a helper function to simplify the provider implementation.
func NewVPCResource() resource.Resource {
	return &vpcResource{}
}

type vpcResource struct {
	client *client.DWSClient
}

// Metadata returns the resource type name.
func (r *vpcResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpc"
}

func (r *vpcResource) Schema(c context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description: "Manages a VPC",
		Attributes: map[string]schema.Attribute{
			ID: schema.StringAttribute{
				Description: "String ID of the VPC, computed",
				Computed:    true,
			},
			VPCIPRangeKeys: schema.StringAttribute{
				Description: "IP range of the VPC",
				Required:    true,
			},
			VPCNameKeys: schema.StringAttribute{
				Description: "Name of the VPC",
				Required:    true,
			},
			VPCDescriptionKeys: schema.StringAttribute{
				Description: "Description of the VPC",
				Optional:    true,
			},
		},
	}
}

func (r *vpcResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.DWSClient)
}

// Create creates the resource and sets the initial Terraform state.
func (r *vpcResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan VPCResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Create new VPC
	clientRequest, err := plan.ToClientRequest()
	if err != nil {
		tflog.Error(ctx, "failed to convert resource to client required type", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}

	vpc, err := r.client.CreateVPC(ctx, clientRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating vpc",
			fmt.Sprintf("Could not create vpc, unexpected error: %s", err),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	err = plan.FromClientResponse(vpc)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating vpc",
			fmt.Sprintf("Could not convert created VPC from client response, unexpected error: %s", err),
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("VPC from client response: %+v", vpc))
	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *vpcResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state VPCResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Get refreshed order value from client
	vpc, err := r.client.GetVPC(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading vpc state",
			fmt.Sprintf("Could not read vpc state ID %s: %s", state.ID.ValueString(), err),
		)
		return
	}

	// Overwrite items with refreshed state
	err = state.FromClientResponse(vpc)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating vpc",
			fmt.Sprintf("Could not convert read VPC from client response, unexpected error: %s", err),
		)
		return
	}
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *vpcResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan VPCResourceModel
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
	_, err = r.client.UpdateVPC(ctx, plan.ID.ValueString(), clientRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating vpc state",
			fmt.Sprintf("Could not update vpc state %s, unexpected error: %s", plan.ID.ValueString(), err),
		)
		return
	}

	// Fetch updated items from GetVPC as UpdateVPC items are not populated.
	vpc, err := r.client.GetVPC(ctx, plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading vpc state",
			fmt.Sprintf("Could not read vpc name %s: %s", plan.ID.ValueString(), err),
		)
		return
	}

	err = plan.FromClientResponse(vpc)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating vpc",
			fmt.Sprintf("Could not convert updated VPC from client response, unexpected error: %s", err),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *vpcResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state VPCResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Delete existing vpc
	if err := r.client.DeleteVPC(ctx, state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting vpc",
			fmt.Sprintf("Could not delete vpc, unexpected error: %s", err),
		)
		return
	}
}

func (r *vpcResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root(ID), req, resp)
}
