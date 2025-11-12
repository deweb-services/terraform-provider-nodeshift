package deployment

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &vmResource{}
	_ resource.ResourceWithConfigure   = &vmResource{}
	_ resource.ResourceWithImportState = &vmResource{}
)

// NewDeploymentResource is a helper function to simplify the provider implementation.
func NewDeploymentResource() resource.Resource {
	return &vmResource{}
}

// vmResource is the resource implementation.
type vmResource struct {
	client client.INodeshiftClient
}

// Metadata returns the resource type name.
func (r *vmResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_deployment"
}

// Schema defines the schema for the resource.
func (r *vmResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a deployment",
		Attributes: map[string]schema.Attribute{
			UUID: schema.StringAttribute{
				Description: "String UUID of the deployment, computed",
				Computed:    true,
			},
			DeploymentKeysImage: schema.StringAttribute{
				Required:    true,
				Description: ImageDescription,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			DeploymentKeysRegion: schema.StringAttribute{
				Required:    true,
				Description: RegionDescription,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			DeploymentKeysCPU: schema.Int64Attribute{
				Required:    true,
				Description: CPUDescription,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			DeploymentKeysRAM: schema.Int64Attribute{
				Required:    true,
				Description: RAMDescription,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			DeploymentKeysDiskSize: schema.Int64Attribute{
				Required:    true,
				Description: DiskSizeDescription,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			DeploymentKeysDiskType: schema.StringAttribute{
				Required:    true,
				Description: DiskTypeDescription,
				Validators: []validator.String{
					stringvalidator.OneOf(availableDiskTypes...),
				},
			},
			DeploymentKeysAssignPublicIPv4: schema.BoolAttribute{
				Optional:    true,
				Description: AssignPublicIPv4Description,
			},
			DeploymentKeysAssignPublicIPv6: schema.BoolAttribute{
				Optional:    true,
				Description: AssignPublicIPv6Description,
			},
			DeploymentKeysSSHKey: schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: SSHKeyDescription,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			DeploymentKeysSSHKeyName: schema.StringAttribute{
				Required:    true,
				Sensitive:   false,
				Description: SSHKeyNameDescription,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			DeploymentKeysHostName: schema.StringAttribute{
				Required:    true,
				Description: HostNameDescription,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			DeploymentKeysNetworkUUID: schema.StringAttribute{
				Optional:    true,
				Description: NetworkUUIDDescription,
			},
			DeploymentKeysPublicIPv4: schema.StringAttribute{
				Computed:    true,
				Description: PublicIPv4Description,
			},
			DeploymentKeysPublicIPv6: schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: PublicIPv6Description,
			},
		},
	}
}

func (r *vmResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	p, ok := req.ProviderData.(client.INodeshiftClient)
	if !ok {
		return
	}

	r.client = p
}

// Create creates the resource and sets the initial Terraform state.
func (r *vmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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

	// Create new Deployment
	vm, err := r.client.CreateDeployment(ctx, plan.ToClientRequest())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Deployment",
			fmt.Sprintf("Could not create Deployment, very unexpected error: %s", err),
		)
		if strings.Contains(err.Error(), "status code: 400") &&
			strings.Contains(err.Error(), "The selected region is not supported") {
			regions, err := r.client.ListRegions(ctx)
			if err != nil {
				tflog.Error(ctx, "failed to fetch regions: "+err.Error())

				return
			}
			resp.Diagnostics.AddError(
				"Invalid region",
				fmt.Sprintf("Supported regions: %#v", regions),
			)
		}

		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.FromClientResponse(vm)
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
func (r *vmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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
	vm, err := r.client.GetDeployment(ctx, state.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Deployment state",
			fmt.Sprintf("Could not read Deployment state UUID %s: %s", state.UUID.ValueString(), err),
		)

		return
	}

	// Overwrite items with refreshed state
	state.FromClientResponse(vm)
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
func (r *vmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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
	if _, err := r.client.UpdateDeployment(ctx, plan.UUID.ValueString(), plan.ToClientRequest()); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Deployment state",
			fmt.Sprintf("Could not update Deployment state %s, unexpected error: %s", plan.UUID.ValueString(), err),
		)

		return
	}

	// Fetch updated items from GetDeployment as UpdateDeployment items are not populated.
	vm, err := r.client.GetDeployment(ctx, plan.UUID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Deployment state",
			fmt.Sprintf("Could not read Deployment name %s: %s", plan.UUID.ValueString(), err),
		)

		return
	}

	plan.FromClientResponse(vm)

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
func (r *vmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

	// Delete existing Deployment
	if err := r.client.DeleteDeployment(ctx, state.UUID.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting VM",
			fmt.Sprintf("Could not delete vm, unexpected error: %s", err),
		)

		return
	}
}

func (r *vmResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import UUID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root(UUID), req, resp)
}
