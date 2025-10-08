package deployment

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

type vmResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Image       types.String `tfsdk:"image"`
	Region      types.String `tfsdk:"region"`
	CPU         types.Int64  `tfsdk:"cpu"`
	RAM         types.Int64  `tfsdk:"ram"`
	Disk        types.Int64  `tfsdk:"disk_size"`
	DiskType    types.String `tfsdk:"disk_type"`
	IPv4        types.Bool   `tfsdk:"assign_public_ipv4"`
	IPv6        types.Bool   `tfsdk:"assign_public_ipv6"`
	Ygg         types.Bool   `tfsdk:"assign_ygg_ip"`
	SSHKey      types.String `tfsdk:"ssh_key"`
	SSHKeyName  types.String `tfsdk:"ssh_key_name"`
	HostName    types.String `tfsdk:"host_name"`
	NetworkUUID types.String `tfsdk:"network_uuid"`
	YggIP       types.String `tfsdk:"ygg_ip"`

	// Computed
	PublicIPv4 types.String `tfsdk:"public_ipv4"`
	PublicIPv6 types.String `tfsdk:"public_ipv6"`
}

func (v *vmResourceModel) ToClientRequest() (*client.CreateDeploymentRequest, error) {
	r := &client.CreateDeploymentRequest{
		Ipv4:        v.IPv4.ValueBool(),
		Ipv6:        v.IPv6.ValueBool(),
		Ygg:         v.Ygg.ValueBool(),
		NetworkUUID: v.NetworkUUID.ValueString(),
	}

	if v.Image.IsUnknown() || v.Image.IsNull() {
		return nil, fmt.Errorf("image is required: %w", client.ErrPropertyEmpty)
	}

	r.ImageVersion = v.Image.ValueString()

	if v.Region.IsUnknown() || v.Region.IsNull() {
		return nil, fmt.Errorf("region is required: %w", client.ErrPropertyEmpty)
	}

	r.Region = v.Region.ValueString()

	if v.CPU.IsUnknown() || v.CPU.IsNull() {
		return nil, fmt.Errorf("cpu is required: %w", client.ErrPropertyEmpty)
	}

	r.CPU = int(v.CPU.ValueInt64())

	if v.RAM.IsUnknown() || v.RAM.IsNull() {
		return nil, fmt.Errorf("ram is required: %w", client.ErrPropertyEmpty)
	}

	r.RAM = int(v.RAM.ValueInt64())

	if v.Disk.IsUnknown() || v.Disk.IsNull() {
		return nil, fmt.Errorf("disk is required: %w", client.ErrPropertyEmpty)
	}

	r.Hdd = int(v.Disk.ValueInt64())

	if v.DiskType.IsUnknown() || v.DiskType.IsNull() {
		return nil, fmt.Errorf("disk_type is required: %w", client.ErrPropertyEmpty)
	}

	r.HddType = v.DiskType.ValueString()

	if v.SSHKey.IsUnknown() || v.SSHKey.IsNull() {
		return nil, fmt.Errorf("ssh_key is required: %w", client.ErrPropertyEmpty)
	}

	r.SSHKey = v.SSHKey.ValueString()

	if v.SSHKeyName.IsUnknown() || v.SSHKeyName.IsNull() {
		return nil, fmt.Errorf("ssh_key_name is required: %w", client.ErrPropertyEmpty)
	}

	r.SSHKeyName = v.SSHKeyName.ValueString()

	if v.HostName.IsUnknown() || v.HostName.IsNull() {
		return nil, fmt.Errorf("host_name is required: %w", client.ErrPropertyEmpty)
	}

	r.HostName = v.HostName.ValueString()

	return r, nil
}

func (v *vmResourceModel) FromClientResponse(c *client.GetDeploymentResponse) {
	v.Image = types.StringValue(c.ImageVersion)
	v.CPU = types.Int64Value(int64(c.Cru))
	v.RAM = types.Int64Value(int64(c.Mru))
	v.Disk = types.Int64Value(int64(c.Sru))
	v.PublicIPv4 = types.StringValue(c.IP)
	v.HostName = types.StringValue(c.Hostname)
}
