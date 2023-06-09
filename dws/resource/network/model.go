package network

import (
	"fmt"
	"net"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type networkResourceModel struct {
	ID          types.String `tfsdk:"id"`
	IPRange     types.String `tfsdk:"ip_range"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

func (m *networkResourceModel) ToClientRequest() (*client.NetworkConfig, error) {
	n := client.NetworkConfig{
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

func (m *networkResourceModel) FromClientResponse(c *client.NetworkConfig) error {
	m = &networkResourceModel{
		ID:          types.StringValue(c.ID),
		Name:        types.StringValue(c.Name),
		Description: types.StringValue(c.Description),
		IPRange:     types.StringValue(c.IPRange.String()),
	}

	return nil
}
