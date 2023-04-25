package vm

import (
	"time"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vmResourceModel maps the resource schema data.
type vmResourceModel struct {
	ID          types.String              `tfsdk:"id"`
	LastUpdated types.String              `tfsdk:"last_updated"`
	Deployment  vmDeploymentResourceModel `tfsdk:"deployment"`
	CPU         []vmCPUResourceModel      `tfsdk:"cpu"`
	RAM         vmRAMResourceModel        `tfsdk:"ram"`
	Disk        []vmDiskResourceModel     `tfsdk:"disk"`
	Protocols   vmProtocolsResourceModel  `tfsdk:"protocols"`
}

type vmDeploymentResourceModel struct {
	Name    types.String `tfsdk:"name"`
	Image   types.String `tfsdk:"image"`
	Network types.String `tfsdk:"network"`
	Region  types.String `tfsdk:"region"`
}

type vmCPUResourceModel struct {
	Quantity types.Int64  `tfsdk:"image"`
	Type     types.String `tfsdk:"type"`
}

type vmRAMResourceModel struct {
	Volume types.Int64 `tfsdk:"image"`
}

type vmDiskResourceModel struct {
	Type   types.String `tfsdk:"type"`
	Volume types.Int64  `tfsdk:"image"`
}

type vmProtocolsResourceModel struct {
	IP vmProtocolsIPResourceModel `tfsdk:"ip"`
}

type vmProtocolsIPResourceModel struct {
	V4 types.Bool `tfsdk:"v4"`
	V6 types.Bool `tfsdk:"v6"`
}

func (v *vmResourceModel) ToClientRequest() *client.VMConfig {
	r := &client.VMConfig{
		Deployment: client.VMDeployment{
			Name:    v.Deployment.Name.ValueString(),
			Image:   v.Deployment.Image.ValueString(),
			Network: v.Deployment.Network.ValueString(),
			Region:  v.Deployment.Region.ValueString(),
		},
		CPU:  make([]client.VMCPU, 0),
		RAM:  client.VMRAM{Volume: v.RAM.Volume.ValueInt64()},
		Disk: make([]client.VMDisk, 0),
		Protocols: client.VMProtocols{
			IP: client.VMProtocolsIP{
				V4: v.Protocols.IP.V4.ValueBool(),
				V6: v.Protocols.IP.V6.ValueBool(),
			}},
	}
	for _, cpu := range v.CPU {
		if (cpu.Quantity.IsUnknown() || cpu.Quantity.IsNull() || cpu.Quantity.ValueInt64() == 0) &&
			(cpu.Type.IsNull() || cpu.Type.IsUnknown() || cpu.Type.ValueString() == "") {
			continue
		}
		r.CPU = append(r.CPU, client.VMCPU{
			Quantity: cpu.Quantity.ValueInt64(),
			Type:     cpu.Type.ValueString(),
		})
	}
	for _, dsk := range v.Disk {
		if (dsk.Type.IsNull() || dsk.Type.IsUnknown() || dsk.Type.ValueString() == "") &&
			(dsk.Volume.IsUnknown() || dsk.Volume.IsNull() || dsk.Volume.ValueInt64() == 0) {
			continue
		}
		r.Disk = append(r.Disk, client.VMDisk{
			Type:   dsk.Type.ValueString(),
			Volume: dsk.Volume.ValueInt64(),
		})
	}
	if !(v.Protocols.IP.V4.ValueBool() || v.Protocols.IP.V6.ValueBool()) {
		r.Protocols.IP.V4 = true
	}

	return r
}

func (v *vmResourceModel) FromClientResponse(c *client.VMConfig) {
	v.Deployment = vmDeploymentResourceModel{
		Name:    types.StringValue(c.Deployment.Name),
		Image:   types.StringValue(c.Deployment.Image),
		Network: types.StringValue(c.Deployment.Network),
		Region:  types.StringValue(c.Deployment.Region),
	}
	v.Protocols = vmProtocolsResourceModel{
		IP: vmProtocolsIPResourceModel{
			V4: types.BoolValue(c.Protocols.IP.V4),
			V6: types.BoolValue(c.Protocols.IP.V6),
		},
	}
	v.RAM = vmRAMResourceModel{Volume: types.Int64Value(c.RAM.Volume)}
	v.CPU = make([]vmCPUResourceModel, 0)
	for _, cpu := range c.CPU {
		v.CPU = append(v.CPU, vmCPUResourceModel{
			Quantity: types.Int64Value(cpu.Quantity),
			Type:     types.StringValue(cpu.Type),
		})
	}
	v.Disk = make([]vmDiskResourceModel, 0)
	for _, dsk := range c.Disk {
		v.Disk = append(v.Disk, vmDiskResourceModel{
			Type:   types.StringValue(dsk.Type),
			Volume: types.Int64Value(dsk.Volume),
		})
	}
	v.ID = types.StringValue(c.ID)
	v.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
}
