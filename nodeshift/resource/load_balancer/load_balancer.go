package load_balancer

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &lbResource{}
	_ resource.ResourceWithConfigure   = &lbResource{}
	_ resource.ResourceWithImportState = &lbResource{}
)

// NewLBResource is a helper function to simplify the provider implementation.
func NewLBResource() resource.Resource {
	return &lbResource{}
}

type lbResource struct {
	client client.INodeshiftClient
}

// Metadata returns the resource type name.
func (r *lbResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_load_balancer"
}

func (r *lbResource) Schema(c context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description: "Manages a load balancer",
		Attributes: map[string]schema.Attribute{
			KeyName: schema.StringAttribute{
				Description: DescriptionName,
				Required:    true,
			},
			KeyReplicas: schema.MapAttribute{
				Description: DescriptionReplicas,
				Required:    true,
				ElementType: types.Int64Type,
			},
			KeyCPUUUIDs: schema.ListAttribute{
				Description: DescriptionCPUUUIDs,
				Required:    true,
				ElementType: types.StringType,
			},
			KeyForwardingRules: schema.ListAttribute{
				Description: DescriptionForwardingRules,
				Required:    true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"in": types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"protocol": types.StringType,
								"port":     types.Int64Type,
							},
						},
						"out": types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"protocol": types.StringType,
								"port":     types.Int64Type,
							},
						},
					},
				},
			},
			KeyVPCUUID: schema.StringAttribute{
				Description: DescriptionVPCUUID,
				Required:    true,
			},

			UUID: schema.StringAttribute{
				Description: DescriptionUUID,
				Computed:    true,
			},
			KeyStatus: schema.StringAttribute{
				Description: DescriptionStatus,
				Computed:    true,
			},
			KeyTaskId: schema.StringAttribute{
				Description: DescriptionTaskId,
				Computed:    true,
			},
		},
	}
}

func (r *lbResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.NodeshiftClient)
}

// Create creates the resource and sets the initial Terraform state.
func (r *lbResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan LBResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Create new LB
	clientRequest, err := plan.ToClientRequest()
	if err != nil {
		tflog.Error(ctx, "failed to convert resource to client required type", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}

	lb, err := r.client.CreateLB(ctx, clientRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating lb",
			fmt.Sprintf("Could not create lb, unexpected error: %s", err),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	err = plan.FromClientResponse(lb)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating lb",
			fmt.Sprintf("Could not convert created LB from client response, unexpected error: %s", err),
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("LB from client response: %+v", lb))
	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *lbResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state LBResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Get refreshed order value from client
	lb, err := r.client.GetLB(ctx, state.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading lb state",
			fmt.Sprintf("Could not read lb state UUID %s: %s", state.UUID.ValueString(), err),
		)
		return
	}

	// Overwrite items with refreshed state
	err = state.FromClientRentedLBResponse(lb)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting lb",
			fmt.Sprintf("Could not convert read LB from client response, unexpected error: %s", err),
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
func (r *lbResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan LBResourceModel
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
	_, err = r.client.UpdateLB(ctx, plan.UUID.ValueString(), clientRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating lb state",
			fmt.Sprintf("Could not update lb state: %s", err),
		)
		return
	}

	// Fetch updated items from GetLB as UpdateLB items are not populated.
	lb, err := r.client.GetLB(ctx, plan.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading lb state",
			fmt.Sprintf("Could not read lb name %s: %s", plan.UUID.ValueString(), err),
		)
		return
	}

	err = plan.FromClientRentedLBResponse(lb)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating lb",
			fmt.Sprintf("Could not convert updated LB from client response, unexpected error: %s", err),
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
func (r *lbResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state LBResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Delete existing lb
	if err := r.client.DeleteLB(ctx, state.UUID.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting lb",
			fmt.Sprintf("Could not delete lb, unexpected error: %s", err),
		)
		return
	}
}

func (r *lbResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import UUID and save to uuid attribute
	resource.ImportStatePassthroughID(ctx, path.Root(UUID), req, resp)
}
