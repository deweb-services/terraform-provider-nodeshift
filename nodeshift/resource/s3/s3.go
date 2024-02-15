package s3

import (
	"context"
	"fmt"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &bucketResource{}
	_ resource.ResourceWithConfigure   = &bucketResource{}
	_ resource.ResourceWithImportState = &bucketResource{}
)

// NewBucketResource is a helper function to simplify the provider implementation.
func NewBucketResource() resource.Resource {
	return &bucketResource{}
}

type bucketResource struct {
	client client.INodeshiftClient
}

// Metadata returns the resource type name.
func (r *bucketResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bucket"
}

func (r *bucketResource) Schema(c context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description: "Manages a s3 Bucket",
		Attributes: map[string]schema.Attribute{
			KeyBucketName: schema.StringAttribute{
				Description: DescriptionBucketName,
				Required:    true,
			},
		},
	}
}

func (r *bucketResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.NodeshiftClient)
}

// Create creates the resource and sets the initial Terraform state.
func (r *bucketResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan BucketResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Create new Bucket
	clientRequest, err := plan.ToClientRequest()
	if err != nil {
		tflog.Error(ctx, "failed to convert resource to client required type", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	bucket, err := r.client.CreateBucket(ctx, clientRequest)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating bucket",
			fmt.Sprintf("Could not create bucket, unexpected error: %s", err),
		)
		return
	}

	err = plan.FromClientResponse(bucket)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating bucket",
			fmt.Sprintf("Could not convert created bucket from client response, unexpected error: %s", err),
		)
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Bucket from client response: %+v", clientRequest))
	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors updating state", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *bucketResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state BucketResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Get refreshed order value from client
	bucket, err := r.client.GetBucket(ctx, state.Key.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading bucket state",
			fmt.Sprintf("Could not read bucket state key %s: %s", state.Key.ValueString(), err),
		)
		return
	}

	// Overwrite items with refreshed state
	err = state.FromClientResponse(bucket)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting bucket",
			fmt.Sprintf("Could not convert read bucket from client response, unexpected error: %s", err),
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
func (r *bucketResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan BucketResourceModel
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
	if err := r.client.UpdateBucket(ctx, clientRequest); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating bucket state",
			fmt.Sprintf("Could not update bucket state %s, unexpected error: %s", plan.Key.ValueString(), err),
		)
		return
	}

	// Fetch updated items from GetBucket as UpdateBucket items are not populated.
	bucket, err := r.client.GetBucket(ctx, plan.Key.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading bucket state",
			fmt.Sprintf("Could not read bucket name %s: %s", plan.Key.ValueString(), err),
		)
		return
	}

	err = plan.FromClientResponse(bucket)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating bucket",
			fmt.Sprintf("Could not convert updated Bucket from client response, unexpected error: %s", err),
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
func (r *bucketResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state BucketResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Errors getting current plan", map[string]interface{}{"count": resp.Diagnostics.ErrorsCount(), "errors": resp.Diagnostics.Errors()})
		return
	}

	// Delete existing bucket
	if err := r.client.DeleteBucket(ctx, state.Key.ValueString()); err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting bucket",
			fmt.Sprintf("Could not delete bucket, unexpected error: %s", err),
		)
		return
	}
}

func (r *bucketResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root(KeyBucketName), req, resp)
}
