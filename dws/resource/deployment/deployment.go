package deployment

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
	client *client.DWSClient
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
			ID: schema.StringAttribute{
				Description: "String ID of the deployment, computed",
				Computed:    true,
			},
			DeploymentKeysImage: schema.StringAttribute{
				Required:    true,
				Description: ImageDescription,
			},
			DeploymentKeysRegion: schema.StringAttribute{
				Required:    true,
				Description: RegionDescription,
			},
			DeploymentKeysCPU: schema.Int64Attribute{
				Required:    true,
				Description: CPUDescription,
			},
			DeploymentKeysRAM: schema.Int64Attribute{
				Required:    true,
				Description: RAMDescription,
			},
			DeploymentKeysDiskSize: schema.Int64Attribute{
				Required:    true,
				Description: DiskSizeDescription,
			},
			DeploymentKeysDiskType: schema.StringAttribute{
				Required:    true,
				Description: DiskTypeDescription,
			},
			DeploymentKeysAssignPublicIPv4: schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: AssignPublicIPv4Description,
			},
			DeploymentKeysAssignPublicIPv6: schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: AssignPublicIPv6Description,
			},
			DeploymentKeysAssignYggIP: schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: AssignYggIPDescription,
			},
			DeploymentKeysSSHKey: schema.StringAttribute{
				Required:    true,
				Description: SSHKeyDescription,
				Sensitive:   true,
			},
			DeploymentKeysHostName: schema.StringAttribute{
				Required:    true,
				Description: HostNameDescription,
			},
			DeploymentKeysVPCID: schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: VPCIDDescription,
			},
			DeploymentKeysPublicIPv4: schema.StringAttribute{
				Computed:    true,
				Description: PublicIPv4Description,
			},
			DeploymentKeysPublicIPv6: schema.StringAttribute{
				Computed:    true,
				Description: PublicIPv4Description,
			},
			DeploymentKeysYggIP: schema.StringAttribute{
				Computed:    true,
				Description: YggIPDescription,
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
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	requestData, err := plan.ToClientRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Deployment",
			fmt.Sprintf("Could not create Deployment, unexpected error: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Deployment to create: %+v", requestData))

	// Create new Deployment
	vm, err := r.client.CreateDeployment(ctx, requestData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Deployment",
			fmt.Sprintf("Could not create Deployment, unexpected error: %s", err),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.FromAsyncAPIResponse(vm)
	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *vmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state vmResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Get refreshed order value from client
	vm, err := r.client.GetDeployment(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Deployment state",
			fmt.Sprintf("Could not read Deployment state ID %s: %s", state.ID.ValueString(), err),
		)
		return
	}

	// Overwrite items with refreshed state
	state.FromClientResponse(vm)
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *vmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan vmResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	requestData, err := plan.ToClientRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Deployment",
			fmt.Sprintf("Could not update Deployment, unexpected error: %s", err.Error()),
		)
		return
	}

	// Update existing order
	_, err = r.client.UpdateDeployment(ctx, plan.ID.ValueString(), requestData)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Deployment state",
			fmt.Sprintf("Could not update Deployment state %s, unexpected error: %s", plan.ID.ValueString(), err),
		)
		return
	}

	// Fetch updated items from GetDeployment as UpdateDeployment items are not populated.
	vm, err := r.client.GetDeployment(ctx, plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Deployment state",
			fmt.Sprintf("Could not read Deployment name %s: %s", plan.ID.ValueString(), err),
		)
		return
	}

	plan.FromClientResponse(vm)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *vmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state vmResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Delete existing Deployment
	if err := r.client.DeleteDeployment(ctx, state.ID.ValueString()); err != nil {
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
