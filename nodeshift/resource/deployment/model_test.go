package deployment

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

func Test_vmResourceModel_FromClientResponse(t *testing.T) {
	t.Parallel()

	type fields struct {
		ID          types.String
		Image       types.String
		Region      types.String
		CPU         types.Int64
		RAM         types.Int64
		Disk        types.Int64
		DiskType    types.String
		IPv4        types.Bool
		IPv6        types.Bool
		SSHKey      types.String
		SSHKeyName  types.String
		HostName    types.String
		NetworkUUID types.String
		PublicIPv4  types.String
		PublicIPv6  types.String
	}
	type args struct {
		c *client.GetDeploymentResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "vm resource from client response",
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue(""),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Value(0),
				RAM:         types.Int64Value(0),
				Disk:        types.Int64Value(0),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue(""),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue(""),
				PublicIPv6:  types.StringNull(),
			},
			args: args{
				c: &client.GetDeploymentResponse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			v := &ResourceModel{
				UUID:        tt.fields.ID,
				Image:       tt.fields.Image,
				Region:      tt.fields.Region,
				CPU:         tt.fields.CPU,
				RAM:         tt.fields.RAM,
				Disk:        tt.fields.Disk,
				DiskType:    tt.fields.DiskType,
				IPv4:        tt.fields.IPv4,
				IPv6:        tt.fields.IPv6,
				SSHKey:      tt.fields.SSHKey,
				SSHKeyName:  tt.fields.SSHKeyName,
				HostName:    tt.fields.HostName,
				NetworkUUID: tt.fields.NetworkUUID,
				PublicIPv4:  tt.fields.PublicIPv4,
				PublicIPv6:  tt.fields.PublicIPv6,
			}
			v.FromClientResponse(tt.args.c)
			assert.Equal(t, tt.fields.ID, v.UUID)
			assert.Equal(t, tt.fields.Image, v.Image)
			assert.Equal(t, tt.fields.Region, v.Region)
			assert.Equal(t, tt.fields.CPU, v.CPU)
			assert.Equal(t, tt.fields.RAM, v.RAM)
			assert.Equal(t, tt.fields.Disk, v.Disk)
			assert.Equal(t, tt.fields.DiskType, v.DiskType)
			assert.Equal(t, tt.fields.IPv4, v.IPv4)
			assert.Equal(t, tt.fields.IPv6, v.IPv6)
			assert.Equal(t, tt.fields.SSHKey, v.SSHKey)
			assert.Equal(t, tt.fields.SSHKeyName, v.SSHKeyName)
			assert.Equal(t, tt.fields.HostName, v.HostName)
			assert.Equal(t, tt.fields.NetworkUUID, v.NetworkUUID)
			assert.Equal(t, tt.fields.PublicIPv4, v.PublicIPv4)
			assert.Equal(t, tt.fields.PublicIPv6, v.PublicIPv6)
		})
	}
}
