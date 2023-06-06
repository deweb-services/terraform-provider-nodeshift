package network

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type networkResourceModel struct {
	ID               types.String `tfsdk:"id"`
	LastUpdated      types.String `tfsdk:"last_updated"`
	IPRange          types.String `tfsdk:"ip_range"`
	Name             types.String `tfsdk:"name"`
	Nodes            types.List   `tfsdk:"nodes"`
	AddWgAccess      types.Bool   `tfsdk:"add_wg_access"`
	Description      types.String `tfsdk:"description"`
	NodesIPRange     types.Map    `tfsdk:"nodes_ip_range"`
	SolutionType     types.String `tfsdk:"solution_type"`
	ExternalIP       types.String `tfsdk:"external_ip"`
	ExternalSK       types.String `tfsdk:"external_sk"`
	NodeDeploymentID types.Map    `tfsdk:"node_deployment_id"`
	PublicNodeID     types.Int64  `tfsdk:"public_node_id"`
	AccessWGConfig   types.String `tfsdk:"access_wg_config"`
}

func (m *networkResourceModel) ToClientRequest() (*client.NetworkConfig, error) {
	n := client.NetworkConfig{
		Name:         m.Name.ValueString(),
		AddWGAccess:  m.AddWgAccess.ValueBool(),
		Description:  m.Description.ValueString(),
		SolutionType: m.SolutionType.ValueString(),
	}

	ip, ipNet, err := net.ParseCIDR(m.IPRange.ValueString())
	if err != nil {
		return nil, fmt.Errorf("failed to parse ip cidr: %s", m.IPRange.String())
	}
	ipNet.IP = ip
	n.IPRange = *ipNet

	nodes := make([]int64, 0, len(m.Nodes.Elements()))
	if diag := m.Nodes.ElementsAs(context.Background(), nodes, false); diag.HasError() {
		return nil, fmt.Errorf("failed to parse nodes: %s, errors: %v", m.Nodes.String(), diag.Errors())
	}
	for i := range nodes {
		n.Nodes = append(n.Nodes, uint32(nodes[i]))
	}

	nodesIPRange := make(map[string]string, len(m.Nodes.Elements()))
	if diag := m.NodesIPRange.ElementsAs(context.Background(), nodesIPRange, false); diag.HasError() {
		return nil, fmt.Errorf("failed to parse nodes ip range: %s, errors: %v", m.NodesIPRange.String(), diag.Errors())
	}
	for key, value := range nodesIPRange {
		ip, ipNet, err := net.ParseCIDR(value)
		if err != nil {
			return nil, fmt.Errorf("failed to parse nodes_ip_range: invalid CIDR found %s", value)
		}
		ipNet.IP = ip
		node, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse nodes_ip_range, invalid node key: %w", err)
		}
		n.NodesIPRange[uint32(node)] = *ipNet
	}

	return &n, nil
}

func (m *networkResourceModel) FromClientResponse(c *client.NetworkConfig) error {
	m = &networkResourceModel{
		ID:             types.StringValue(c.ID),
		Name:           types.StringValue(c.Name),
		Description:    types.StringValue(c.Description),
		SolutionType:   types.StringValue(c.SolutionType),
		AccessWGConfig: types.StringValue(c.AccessWGConfig),
		AddWgAccess:    types.BoolValue(c.AddWGAccess),
		PublicNodeID:   types.Int64Value(int64(c.PublicNodeID)),
		IPRange:        types.StringValue(c.IPRange.String()),
		ExternalIP:     types.StringValue(c.ExternalIP.String()),
		// base64 encoded
		ExternalSK: types.StringValue(c.ExternalSK.String()),
	}

	nodesInt64 := make([]int64, 0, len(c.Nodes))
	diag := diag.Diagnostics{}
	for i := range c.Nodes {
		nodesInt64 = append(nodesInt64, int64(c.Nodes[i]))
	}
	if m.Nodes, diag = types.ListValueFrom(context.Background(), m.Nodes.ElementType(context.Background()), nodesInt64); diag.HasError() {
		return errors.New("failed to parse response nodes list")
	}

	nodesIPRange := make(map[string]string, len(c.NodesIPRange))
	for key, value := range c.NodesIPRange {
		nodesIPRange[strconv.FormatUint(uint64(key), 10)] = value.String()
	}
	if m.NodesIPRange, diag = types.MapValueFrom(context.Background(), m.NodesIPRange.ElementType(context.Background()), nodesIPRange); diag.HasError() {
		return errors.New("failed to parse response nodes ip range")
	}

	nodeDeploymentId := make(map[string]string, len(c.NodeDeploymentID))
	for key, value := range c.NodeDeploymentID {
		nodeDeploymentId[strconv.Itoa(int(key))] = strconv.FormatUint(value, 10)
	}
	if m.NodeDeploymentID, diag = types.MapValueFrom(context.Background(), m.NodeDeploymentID.ElementType(context.Background()), nodeDeploymentId); diag.HasError() {
		return errors.New("failed to parse response node deployments id")
	}

	return nil
}
