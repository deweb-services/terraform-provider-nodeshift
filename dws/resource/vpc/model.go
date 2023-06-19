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
	vpc := client.VPCConfig{
		Name:        m.Name.ValueString(),
		Description: m.Description.ValueString(),
	}

	_, _, err := net.ParseCIDR(m.IPRange.ValueString())
	if err != nil {
		return nil, fmt.Errorf("failed to parse ip cidr: %s", m.IPRange.String())
	}

	vpc.IPRange = m.IPRange.ValueString()

	return &vpc, nil
}

func (m *VPCResourceModel) FromClientResponse(c *client.VPCConfig) error {
	m = &VPCResourceModel{
		ID:          types.StringValue(c.ID),
		Name:        types.StringValue(c.Name),
		Description: types.StringValue(c.Description),
		IPRange:     types.StringValue(c.IPRange),
	}

	return nil
}
