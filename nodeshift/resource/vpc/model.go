package vpc

import (
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

type VPCResourceModel struct {
	UUID        types.String `tfsdk:"uuid"`
	IPRange     types.String `tfsdk:"ip_range"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

func (m *VPCResourceModel) ToClientRequest() (*client.CreateVPCRequest, error) {
	vpc := client.CreateVPCRequest{
		Name:        m.Name.ValueString(),
		Description: m.Description.ValueString(),
	}

	_, _, err := net.ParseCIDR(m.IPRange.ValueString())
	if err != nil {
		return nil, fmt.Errorf("failed to parse ip cidr: %w", err)
	}

	vpc.IPRange = m.IPRange.ValueString()

	return &vpc, nil
}

func (m *VPCResourceModel) FromClientResponse(c *client.GetVPCResponse) error {
	m.UUID = types.StringValue(c.UUID)
	m.Name = types.StringValue(c.Name)
	m.Description = types.StringValue(c.Description)
	m.IPRange = types.StringValue(c.IPRange)

	return nil
}
