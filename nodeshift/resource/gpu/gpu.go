package gpu

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &gpuResource{}
	_ resource.ResourceWithConfigure   = &gpuResource{}
	_ resource.ResourceWithImportState = &gpuResource{}
)

// NewGPUResource is a helper function to simplify the provider implementation.
func NewGPUResource() resource.Resource {
	return &gpuResource{}
}

type gpuResource struct {
	client client.INodeshiftClient
}

// Metadata returns the resource type name.
func (r *gpuResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gpu"
}

func (r *gpuResource) Schema(c context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description: "Manages a GPU",
		Attributes: map[string]schema.Attribute{
			UUID: schema.StringAttribute{
				Computed:    true,
				Description: DescriptionUUID,
			},
			KeyGPUName: schema.StringAttribute{
				Required:    true,
				Description: DescriptionGPUName,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			KeyImage: schema.StringAttribute{
				Required:    true,
				Description: DescriptionImage,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			KeySSHKey: schema.StringAttribute{
				Required:    true,
				Description: DescriptionSSHKey,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			KeyGPUCount: schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: DescriptionGPUCount,
				Default:     int64default.StaticInt64(1),
			},
			KeyRegion: schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: DescriptionRegion,
			},
			KeyDiskSizeGB: schema.Int64Attribute{
				Required:    true,
				Description: DescriptionDiskSizeGB,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			KeyMinCudaVersion: schema.StringAttribute{
				Optional:    true,
				Description: DescriptionMinCudaVersion,
			},
			KeyMachineTypeVersion: schema.StringAttribute{
				Optional:    true,
				Description: DescriptionMachineTypeVersion,
				Validators: []validator.String{
					stringvalidator.OneOf(availableMachineTypes...),
				},
			},
		},
	}
}

func (r *gpuResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	p, ok := req.ProviderData.(*client.NodeshiftClient)
	if !ok {
		return
	}

	r.client = p
}

// Create creates the resource and sets the initial Terraform state.
func (r *gpuResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan ResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(
			ctx,
			"Errors getting current plan",
			map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()},
		)

		return
	}

	gpu, err := r.client.CreateGPU(ctx, plan.ToClientRequest())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating gpu",
			fmt.Sprintf("Could not create gpu, unexpected error: %s", err),
		)

		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.FromClientResponse(gpu)
	tflog.Info(ctx, fmt.Sprintf("GPU from client response: %+v", gpu))
	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(
			ctx,
			"Errors updating state",
			map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()},
		)
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *gpuResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state ResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(
			ctx,
			"Errors getting current plan",
			map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()},
		)

		return
	}

	// Get refreshed order value from client
	gpu, err := r.client.GetGPU(ctx, state.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading gpu state",
			fmt.Sprintf("Could not read gpu state UUID %s: %s", state.UUID.ValueString(), err),
		)

		return
	}

	// Overwrite items with refreshed state
	state.FromClientResponse(gpu)
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(
			ctx,
			"Errors updating state",
			map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()},
		)
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *gpuResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan ResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(
			ctx,
			"Errors getting current plan",
			map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()},
		)

		return
	}

	// Update existing order
	if _, err := r.client.UpdateGPU(ctx, plan.UUID.ValueString(), plan.ToClientRequest()); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating gpu state",
			fmt.Sprintf("Could not update gpu state %s, unexpected error: %s", plan.UUID.ValueString(), err),
		)

		return
	}

	// Fetch updated items from GetGPU as UpdateGPU items are not populated.
	gpu, err := r.client.GetGPU(ctx, plan.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading gpu state",
			fmt.Sprintf("Could not read gpu name %s: %s", plan.UUID.ValueString(), err),
		)

		return
	}

	plan.FromClientResponse(gpu)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(
			ctx,
			"Errors updating state",
			map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()},
		)
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *gpuResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state ResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(
			ctx,
			"Errors getting current plan",
			map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()},
		)

		return
	}

	// Delete existing gpu
	if err := r.client.DeleteGPU(ctx, state.UUID.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting gpu",
			fmt.Sprintf("Could not delete gpu, unexpected error: %s", err),
		)

		return
	}
}

func (r *gpuResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import UUID and save to uuid attribute
	resource.ImportStatePassthroughID(ctx, path.Root(UUID), req, resp)
}
