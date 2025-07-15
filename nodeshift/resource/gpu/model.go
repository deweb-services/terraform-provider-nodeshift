package gpu

import (
	"errors"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

type GPUResourceModel struct {
	GPUName        types.String `tfsdk:"gpu_name"`
	Image          types.String `tfsdk:"image"`
	SSHKey         types.String `tfsdk:"ssh_key"`
	GPUCount       types.Int64  `tfsdk:"gpu_count"`
	Region         types.String `tfsdk:"region"`
	UUID           types.String `tfsdk:"uuid"`
	DiskSize       types.Int64  `tfsdk:"disk_size_gb"`
	MinCudaVersion types.String `tfsdk:"min_cuda_version"`
}

func (m *GPUResourceModel) ToClientRequest() (*client.GPUConfig, error) {
	if m.GPUName.IsUnknown() || m.GPUName.IsNull() {
		return nil, errors.New("gpu name property is required and cannot be empty")
	}
	if m.Image.IsUnknown() || m.Image.IsNull() {
		return nil, errors.New("image property is required and cannot be empty")
	}
	if m.SSHKey.IsUnknown() || m.SSHKey.IsNull() {
		return nil, errors.New("ssh key property is required and cannot be empty")
	}

	return &client.GPUConfig{
		GPUName:        strings.TrimSpace(m.GPUName.ValueString()),
		Image:          m.Image.ValueString(),
		SSHKey:         m.SSHKey.ValueString(),
		GPUCount:       m.GPUCount.ValueInt64(),
		Region:         m.Region.ValueString(),
		Disk:           m.DiskSize.ValueInt64(),
		MinCudaVersion: m.MinCudaVersion.ValueString(),
	}, nil
}

func (m *GPUResourceModel) FromClientResponse(c *client.GPUConfigResponse) error {
	m.GPUName = types.StringValue(c.GPUName)
	m.Image = types.StringValue(c.Image)
	m.Region = types.StringValue(c.Region)
	m.UUID = types.StringValue(c.UUID)
	return nil
}

func (m *GPUResourceModel) FromClientRentedGPUResponse(c *client.RentedGpuInfoResponse) error {
	m.GPUName = types.StringValue(c.GpuName)
	m.GPUCount = types.Int64Value(c.NumGpus)
	return nil
}
