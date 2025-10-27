package vpc

import (
	"fmt"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

type ResourceModel struct {
	UUID        types.String `tfsdk:"uuid"`
	IPRange     types.String `tfsdk:"ip_range"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

func (m *ResourceModel) ToClientRequest() (*client.CreateVPCRequest, error) {
	vpc := client.CreateVPCRequest{
		Name:        m.Name.ValueString(),
		Description: m.Description.ValueString(),
	}

	if m.IPRange.IsNull() || m.IPRange.IsUnknown() {
		return &vpc, nil
	}

	ipRangePrefix := m.IPRange.ValueString()
	if strings.Count(ipRangePrefix, ".") != 1 {
		return nil, fmt.Errorf("incorrect ip range prefix: %w", errIncorrectOctet)
	}

	_, _, err := net.ParseCIDR(ipRangePrefix + ".0.0/16")
	if err != nil {
		return nil, fmt.Errorf("failed to parse ip cidr: %w", err)
	}

	vpc.IPRange = m.IPRange.ValueString()

	return &vpc, nil
}

func (m *ResourceModel) FromClientResponse(c *client.GetVPCResponse) {
	m.UUID = types.StringValue(c.UUID)
	m.Name = types.StringValue(c.Name)
	m.Description = types.StringValue(c.Description)
	m.IPRange = types.StringValue(c.IPRange)
}
