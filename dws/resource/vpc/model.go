package vpc

import (
	"fmt"
	"net"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type VPCResourceModel struct {
	ID          types.String `tfsdk:"id"`
	IPRange     types.String `tfsdk:"ip_range"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

func (m *VPCResourceModel) ToClientRequest() (*client.VPCConfig, error) {
	n := client.VPCConfig{
		Name:        m.Name.ValueString(),
		Description: m.Description.ValueString(),
	}

	ip, ipNet, err := net.ParseCIDR(m.IPRange.ValueString())
	if err != nil {
		return nil, fmt.Errorf("failed to parse ip cidr: %s", m.IPRange.String())
	}
	ipNet.IP = ip
	n.IPRange = *ipNet

	return &n, nil
}

func (m *VPCResourceModel) FromClientResponse(c *client.VPCConfig) error {
	m = &VPCResourceModel{
		ID:          types.StringValue(c.ID),
		Name:        types.StringValue(c.Name),
		Description: types.StringValue(c.Description),
		IPRange:     types.StringValue(c.IPRange.String()),
	}

	return nil
}
