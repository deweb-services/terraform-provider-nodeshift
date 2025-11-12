package gpu

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

type ResourceModel struct {
	GPUName        types.String `tfsdk:"gpu_name"`
	Image          types.String `tfsdk:"image"`
	SSHKey         types.String `tfsdk:"ssh_key"`
	GPUCount       types.Int64  `tfsdk:"gpu_count"`
	Region         types.String `tfsdk:"region"`
	UUID           types.String `tfsdk:"uuid"`
	DiskSize       types.Int64  `tfsdk:"disk_size_gb"`
	MinCudaVersion types.String `tfsdk:"min_cuda_version"`
	MachineType    types.String `tfsdk:"machine_type"`
}

func (m *ResourceModel) ToClientRequest() *client.CreateGPURequest {
	return &client.CreateGPURequest{
		GPUName:        strings.TrimSpace(m.GPUName.ValueString()),
		Image:          m.Image.ValueString(),
		SSHKey:         m.SSHKey.ValueString(),
		GPUCount:       m.GPUCount.ValueInt64(),
		Region:         m.Region.ValueString(),
		Disk:           m.DiskSize.ValueInt64(),
		MinCudaVersion: m.MinCudaVersion.ValueString(),
		MachineType:    m.MachineType.ValueString(),
	}
}

func (m *ResourceModel) FromClientResponse(c *client.GetGPUResponse) {
	m.UUID = types.StringValue(c.UUID)
	m.GPUName = types.StringValue(c.GpuName)
	m.GPUCount = types.Int64Value(c.NumGpus)
}
