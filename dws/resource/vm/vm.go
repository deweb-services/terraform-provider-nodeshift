package vm

import (
	"context"
	"fmt"

	"github.com/deweb-services/dws-terraform-provider/dws/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &vmResource{}
	_ resource.ResourceWithConfigure   = &vmResource{}
	_ resource.ResourceWithImportState = &vmResource{}
)

// NewVMResource is a helper function to simplify the provider implementation.
func NewVMResource() resource.Resource {
	return &vmResource{}
}

// vmResource is the resource implementation.
type vmResource struct {
	client *client.DWSClient
}

// Metadata returns the resource type name.
func (r *vmResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vm"
}

// Schema defines the schema for the resource.
func (r *vmResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			ID: schema.StringAttribute{
				Computed: true,
			},
			LastUpdated: schema.StringAttribute{
				Computed: true,
			},
			VmKeysDeployment: schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					VmKeysDeploymentName: schema.StringAttribute{
						Optional: true,
					},
					VmKeysDeploymentImage: schema.StringAttribute{
						Optional: true,
					},
					VmKeysDeploymentNetwork: schema.StringAttribute{
						Optional: true,
					},
					VmKeysDeploymentRegion: schema.StringAttribute{
						Optional: true,
					},
				},
			},
			VmKeysCPU: schema.ListNestedAttribute{
				Optional: true,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						VmKeysCPUQuantity: schema.Int64Attribute{
							Optional: true,
						},
						VmKeysCPUType: schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			VmKeysRam: schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					VmKeysRamVolume: schema.Int64Attribute{
						Optional: true,
						Computed: true,
					},
				},
			},
			VmKeysDisk: schema.ListNestedAttribute{
				Optional: true,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						VmKeysDiskType: schema.StringAttribute{
							Optional: true,
						},
						VmKeysDiskVolume: schema.Int64Attribute{
							Optional: true,
						},
					},
				},
			},
			VmKeysProtocols: schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					VmKeysProtocolsIP: schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							VmKeysProtocolsIPV4: schema.BoolAttribute{
								Optional: true,
							},
							VmKeysProtocolsIPV6: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
	}
}

func (r *vmResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.DWSClient)
}

// Create creates the resource and sets the initial Terraform state.
func (r *vmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan vmResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new VM
	vm, err := r.client.CreateVM(ctx, plan.ToClientRequest())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating VM",
			fmt.Sprintf("Could not create VM, unexpected error: %s", err),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.FromClientResponse(vm)
	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *vmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state vmResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from client
	vm, err := r.client.GetVM(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading VM state",
			fmt.Sprintf("Could not read VM state ID %s: %s", state.ID.ValueString(), err),
		)
		return
	}

	// Overwrite items with refreshed state
	state.FromClientResponse(vm)
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *vmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan vmResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing order
	_, err := r.client.UpdateVM(ctx, plan.ID.ValueString(), plan.ToClientRequest())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating VM state",
			fmt.Sprintf("Could not update VM state %s, unexpected error: %s", plan.ID.ValueString(), err),
		)
		return
	}

	// Fetch updated items from GetVM as UpdateVM items are not populated.
	vm, err := r.client.GetVM(ctx, plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading VM state",
			fmt.Sprintf("Could not read VM name %s: %s", plan.ID.ValueString(), err),
		)
		return
	}

	plan.FromClientResponse(vm)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *vmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state vmResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing VM
	if err := r.client.DeleteVM(ctx, state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting VM",
			fmt.Sprintf("Could not delete vm, unexpected error: %s", err),
		)
		return
	}
}

func (r *vmResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root(ID), req, resp)
}
