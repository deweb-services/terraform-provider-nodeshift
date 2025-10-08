package loadbalancer

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

type RuleEndpointModel struct {
	Protocol types.String `tfsdk:"protocol"`
	Port     types.Int64  `tfsdk:"port"`
}

type ForwardingRuleModel struct {
	In  RuleEndpointModel `tfsdk:"in"`
	Out RuleEndpointModel `tfsdk:"out"`
}

type LBResourceModel struct {
	Name            types.String `tfsdk:"name"`
	Replicas        types.Map    `tfsdk:"replicas"`
	CPUUUIDs        types.List   `tfsdk:"cpu_uuids"`
	ForwardingRules types.List   `tfsdk:"forwarding_rules"`
	VPCUUID         types.String `tfsdk:"vpc_uuid"`

	UUID   types.String `tfsdk:"uuid"`
	Status types.String `tfsdk:"status"`
	TaskID types.String `tfsdk:"task_id"`
}

func (m *LBResourceModel) ToClientRequest() (*client.CreateLBRequest, error) {
	replicas := make(map[string]int)
	for k, v := range m.Replicas.Elements() {
		if intVal, ok := v.(types.Int64); ok && !intVal.IsNull() {
			replicas[k] = int(intVal.ValueInt64())
		}
	}

	cpuUUIDs := make([]string, 0, len(m.CPUUUIDs.Elements()))
	for _, v := range m.CPUUUIDs.Elements() {
		if s, ok := v.(types.String); ok && !s.IsNull() {
			cpuUUIDs = append(cpuUUIDs, s.ValueString())
		}
	}

	forwardingRules := make([]client.ForwardingRule, 0, len(m.ForwardingRules.Elements()))
	for _, ruleAttr := range m.ForwardingRules.Elements() {
		rule, err := convertToForwardingRule(ruleAttr)
		if err != nil {
			return nil, fmt.Errorf("failed to convert resource to client required type: %w", err)
		}

		forwardingRules = append(forwardingRules, *rule)
	}

	return &client.CreateLBRequest{
		Name:            m.Name.ValueString(),
		Replicas:        replicas,
		CPUUUIDs:        cpuUUIDs,
		ForwardingRules: forwardingRules,
		VPCUUID:         m.VPCUUID.ValueString(),
	}, nil
}

func convertToForwardingRule(attr attr.Value) (*client.ForwardingRule, error) {
	if attr.IsNull() {
		return nil, fmt.Errorf("forward rule is required: %w", client.ErrPropertyEmpty)
	}

	ruleObj, ok := attr.(types.Object)
	if !ok {
		return nil, fmt.Errorf("forward rule should be Object type: %w", client.ErrPropertyType)
	}

	inAttr, ok := ruleObj.Attributes()["in"].(types.Object)
	if !ok {
		return nil, fmt.Errorf("failed to cast forward rule -> in: %w", client.ErrPropertyCast)
	}

	if inAttr.IsNull() {
		return nil, fmt.Errorf("forward rule -> in should be Object type: %w", client.ErrPropertyType)
	}

	outAttr, ok := ruleObj.Attributes()["out"].(types.Object)
	if !ok {
		return nil, fmt.Errorf("failed to cast forward rule -> out: %w", client.ErrPropertyCast)
	}

	if outAttr.IsNull() {
		return nil, fmt.Errorf("forward rule -> out should be Object type: %w", client.ErrPropertyType)
	}

	inProtocol, err := getProtocol(inAttr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert in resource to port: %w", err)
	}

	outProtocol, err := getProtocol(outAttr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert out resource to port: %w", err)
	}

	inPort, err := getPort(inAttr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert in resource to port: %w", err)
	}

	outPort, err := getPort(outAttr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert out resource to port: %w", err)
	}

	return &client.ForwardingRule{
		In: client.RuleEndpoint{
			Protocol: inProtocol,
			Port:     inPort,
		},
		Out: client.RuleEndpoint{
			Protocol: outProtocol,
			Port:     outPort,
		},
	}, nil
}

func getProtocol(attr types.Object) (string, error) {
	protocol := attr.Attributes()["protocol"]
	if protocol.IsNull() {
		return "", fmt.Errorf("forward rule -> protocol is required: %w", client.ErrPropertyEmpty)
	}

	pt, ok := protocol.(types.String)
	if !ok {
		return "", fmt.Errorf("failed to cast forward rule -> protocol: %w", client.ErrPropertyCast)
	}

	if pt.IsNull() {
		return "", fmt.Errorf("forward rule -> protocol is required: %w", client.ErrPropertyEmpty)
	}

	return pt.ValueString(), nil
}

func getPort(attr types.Object) (int, error) {
	port := attr.Attributes()["port"]
	if port.IsNull() {
		return 0, fmt.Errorf("forward rule -> port is required: %w", client.ErrPropertyEmpty)
	}

	pt, ok := port.(types.Int64)
	if !ok {
		return 0, fmt.Errorf("failed to cast forward rule -> port: %w", client.ErrPropertyCast)
	}

	if pt.IsNull() {
		return 0, fmt.Errorf("forward rule -> port is required: %w", client.ErrPropertyEmpty)
	}

	return int(pt.ValueInt64()), nil
}

func (m *LBResourceModel) FromClientResponse(c *client.GetLBResponse) error {
	m.UUID = types.StringValue(c.UUID)
	m.Status = types.StringValue(c.Status)

	return nil
}

func (m *LBResourceModel) FromClientRentedLBResponse(c *client.GetLBResponse) error {
	m.UUID = types.StringValue(c.UUID)
	m.Status = types.StringValue(c.Status)

	return nil
}
