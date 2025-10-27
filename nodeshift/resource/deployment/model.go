package deployment

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

type ResourceModel struct {
	UUID        types.String `tfsdk:"uuid"`
	Image       types.String `tfsdk:"image"`
	Region      types.String `tfsdk:"region"`
	CPU         types.Int64  `tfsdk:"cpu"`
	RAM         types.Int64  `tfsdk:"ram"`
	Disk        types.Int64  `tfsdk:"disk_size"`
	DiskType    types.String `tfsdk:"disk_type"`
	IPv4        types.Bool   `tfsdk:"assign_public_ipv4"`
	IPv6        types.Bool   `tfsdk:"assign_public_ipv6"`
	SSHKey      types.String `tfsdk:"ssh_key"`
	SSHKeyName  types.String `tfsdk:"ssh_key_name"`
	HostName    types.String `tfsdk:"host_name"`
	NetworkUUID types.String `tfsdk:"network_uuid"`

	// Computed
	PublicIPv4 types.String `tfsdk:"public_ipv4"`
	PublicIPv6 types.String `tfsdk:"public_ipv6"`
}

func (v *ResourceModel) ToClientRequest() *client.CreateDeploymentRequest {
	return &client.CreateDeploymentRequest{
		Ipv4:         v.IPv4.ValueBool(),
		Ipv6:         v.IPv6.ValueBool(),
		NetworkUUID:  v.NetworkUUID.ValueString(),
		ImageVersion: v.Image.ValueString(),
		Region:       v.Region.ValueString(),
		CPU:          v.CPU.ValueInt64(),
		RAM:          v.RAM.ValueInt64(),
		Hdd:          v.Disk.ValueInt64(),
		HddType:      v.DiskType.ValueString(),
		SSHKey:       v.SSHKey.ValueString(),
		SSHKeyName:   v.SSHKeyName.ValueString(),
		HostName:     v.HostName.ValueString(),
	}
}

func (v *ResourceModel) FromClientResponse(c *client.GetDeploymentResponse) {
	v.UUID = types.StringValue(c.UUID)
	v.Image = types.StringValue(c.ImageVersion)
	v.CPU = types.Int64Value(int64(c.Cru))
	v.RAM = types.Int64Value(int64(c.Mru))
	v.Disk = types.Int64Value(int64(c.Sru))
	v.PublicIPv4 = types.StringValue(c.IP)
	v.PublicIPv6 = types.StringNull()
	v.HostName = types.StringValue(c.Hostname)
}
