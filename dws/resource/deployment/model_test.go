package deployment

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Test_vmResourceModel_FromAsyncAPIResponse(t *testing.T) {
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
		Ygg         types.Bool
		SSHKey      types.String
		SSHKeyName  types.String
		HostName    types.String
		NetworkUUID types.String
		PublicIPv4  types.String
		PublicIPv6  types.String
		YggIP       types.String
	}
	type args struct {
		c *client.AsyncAPIDeploymentResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "vm resource from async api response",
			args: args{
				c: &client.AsyncAPIDeploymentResponse{
					Data: &client.DeploymentResponseData{
						IP:           "public_ipv4",
						IPv6:         "",
						Ygg:          "",
						ProviderPlan: "",
					},
				},
			},
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue("image"),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Value(1),
				RAM:         types.Int64Value(2),
				Disk:        types.Int64Value(3),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue("host_name"),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue("public_ipv4"),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
		},
		{
			name: "vm resource from async api response error",
			args: args{
				c: &client.AsyncAPIDeploymentResponse{
					Data: &client.DeploymentResponseData{
						IP:           "",
						IPv6:         "",
						Ygg:          "",
						ProviderPlan: "",
					},
				},
			},
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue("image"),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Value(1),
				RAM:         types.Int64Value(2),
				Disk:        types.Int64Value(3),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(false),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue("host_name"),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue(""),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vmResourceModel{
				ID:          tt.fields.ID,
				Image:       tt.fields.Image,
				Region:      tt.fields.Region,
				CPU:         tt.fields.CPU,
				RAM:         tt.fields.RAM,
				Disk:        tt.fields.Disk,
				DiskType:    tt.fields.DiskType,
				IPv4:        tt.fields.IPv4,
				IPv6:        tt.fields.IPv6,
				Ygg:         tt.fields.Ygg,
				SSHKey:      tt.fields.SSHKey,
				SSHKeyName:  tt.fields.SSHKeyName,
				HostName:    tt.fields.HostName,
				NetworkUUID: tt.fields.NetworkUUID,
				PublicIPv4:  tt.fields.PublicIPv4,
				PublicIPv6:  tt.fields.PublicIPv6,
				YggIP:       tt.fields.YggIP,
			}
			v.FromAsyncAPIResponse(tt.args.c)
			assert.Equal(t, tt.fields.ID, v.ID)
			assert.Equal(t, tt.fields.Image, v.Image)
			assert.Equal(t, tt.fields.Region, v.Region)
			assert.Equal(t, tt.fields.CPU, v.CPU)
			assert.Equal(t, tt.fields.RAM, v.RAM)
			assert.Equal(t, tt.fields.Disk, v.Disk)
			assert.Equal(t, tt.fields.DiskType, v.DiskType)
			assert.Equal(t, tt.fields.IPv4, v.IPv4)
			assert.Equal(t, tt.fields.IPv6, v.IPv6)
			assert.Equal(t, tt.fields.Ygg, v.Ygg)
			assert.Equal(t, tt.fields.SSHKey, v.SSHKey)
			assert.Equal(t, tt.fields.SSHKeyName, v.SSHKeyName)
			assert.Equal(t, tt.fields.HostName, v.HostName)
			assert.Equal(t, tt.fields.NetworkUUID, v.NetworkUUID)
			assert.Equal(t, tt.fields.PublicIPv4, v.PublicIPv4)
			assert.Equal(t, tt.fields.PublicIPv6, v.PublicIPv6)
			assert.Equal(t, tt.fields.YggIP, v.YggIP)
		})
	}
}

func Test_vmResourceModel_FromClientResponse(t *testing.T) {
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
		Ygg         types.Bool
		SSHKey      types.String
		SSHKeyName  types.String
		HostName    types.String
		NetworkUUID types.String
		PublicIPv4  types.String
		PublicIPv6  types.String
		YggIP       types.String
	}
	type args struct {
		c *client.CreatedDeployment
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
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue(""),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue(""),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
			args: args{
				c: &client.CreatedDeployment{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vmResourceModel{
				ID:          tt.fields.ID,
				Image:       tt.fields.Image,
				Region:      tt.fields.Region,
				CPU:         tt.fields.CPU,
				RAM:         tt.fields.RAM,
				Disk:        tt.fields.Disk,
				DiskType:    tt.fields.DiskType,
				IPv4:        tt.fields.IPv4,
				IPv6:        tt.fields.IPv6,
				Ygg:         tt.fields.Ygg,
				SSHKey:      tt.fields.SSHKey,
				SSHKeyName:  tt.fields.SSHKeyName,
				HostName:    tt.fields.HostName,
				NetworkUUID: tt.fields.NetworkUUID,
				PublicIPv4:  tt.fields.PublicIPv4,
				PublicIPv6:  tt.fields.PublicIPv6,
				YggIP:       tt.fields.YggIP,
			}
			v.FromClientResponse(tt.args.c)
			assert.Equal(t, tt.fields.ID, v.ID)
			assert.Equal(t, tt.fields.Image, v.Image)
			assert.Equal(t, tt.fields.Region, v.Region)
			assert.Equal(t, tt.fields.CPU, v.CPU)
			assert.Equal(t, tt.fields.RAM, v.RAM)
			assert.Equal(t, tt.fields.Disk, v.Disk)
			assert.Equal(t, tt.fields.DiskType, v.DiskType)
			assert.Equal(t, tt.fields.IPv4, v.IPv4)
			assert.Equal(t, tt.fields.IPv6, v.IPv6)
			assert.Equal(t, tt.fields.Ygg, v.Ygg)
			assert.Equal(t, tt.fields.SSHKey, v.SSHKey)
			assert.Equal(t, tt.fields.SSHKeyName, v.SSHKeyName)
			assert.Equal(t, tt.fields.HostName, v.HostName)
			assert.Equal(t, tt.fields.NetworkUUID, v.NetworkUUID)
			assert.Equal(t, tt.fields.PublicIPv4, v.PublicIPv4)
			assert.Equal(t, tt.fields.PublicIPv6, v.PublicIPv6)
			assert.Equal(t, tt.fields.YggIP, v.YggIP)
		})
	}
}

func Test_vmResourceModel_ToClientRequest(t *testing.T) {
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
		Ygg         types.Bool
		SSHKey      types.String
		SSHKeyName  types.String
		HostName    types.String
		NetworkUUID types.String
		PublicIPv4  types.String
		PublicIPv6  types.String
		YggIP       types.String
	}
	tests := []struct {
		name    string
		fields  fields
		want    *client.DeploymentConfig
		wantErr bool
	}{
		{
			name: "vm resource to client request",
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue("image"),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Value(1),
				RAM:         types.Int64Value(2),
				Disk:        types.Int64Value(3),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue("host_name"),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue("public_ipv4"),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
			want: &client.DeploymentConfig{
				ImageVersion: "image",
				Region:       "region",
				CPU:          1,
				RAM:          2,
				Hdd:          3,
				HddType:      "disk_type",
				Ipv4:         true,
				Ipv6:         false,
				Ygg:          false,
				SSHKey:       "ssh_key",
				SSHKeyName:   "ssh_key_name",
				HostName:     "host_name",
				NetworkUUID:  "network_uuid",
			},
		},
		{
			name: "vm resource to client request error image",
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringNull(),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Value(1),
				RAM:         types.Int64Value(2),
				Disk:        types.Int64Value(3),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue("host_name"),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue("public_ipv4"),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "vm resource to client request error region",
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue("image"),
				Region:      types.StringNull(),
				CPU:         types.Int64Value(1),
				RAM:         types.Int64Value(2),
				Disk:        types.Int64Value(3),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue("host_name"),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue("public_ipv4"),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "vm resource to client request error cpu",
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue("image"),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Null(),
				RAM:         types.Int64Value(2),
				Disk:        types.Int64Value(3),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue("host_name"),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue("public_ipv4"),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "vm resource to client request error ram",
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue("image"),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Value(1),
				RAM:         types.Int64Null(),
				Disk:        types.Int64Value(3),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue("host_name"),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue("public_ipv4"),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "vm resource to client request error disk",
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue("image"),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Value(1),
				RAM:         types.Int64Value(2),
				Disk:        types.Int64Null(),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue("host_name"),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue("public_ipv4"),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "vm resource to client request error disktype",
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue("image"),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Value(1),
				RAM:         types.Int64Value(2),
				Disk:        types.Int64Value(3),
				DiskType:    types.StringNull(),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue("host_name"),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue("public_ipv4"),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "vm resource to client request error ssh key",
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue("image"),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Value(1),
				RAM:         types.Int64Value(2),
				Disk:        types.Int64Value(3),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringNull(),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringValue("host_name"),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue("public_ipv4"),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "vm resource to client request error host name",
			fields: fields{
				ID:          types.StringValue(""),
				Image:       types.StringValue("image"),
				Region:      types.StringValue("region"),
				CPU:         types.Int64Value(1),
				RAM:         types.Int64Value(2),
				Disk:        types.Int64Value(3),
				DiskType:    types.StringValue("disk_type"),
				IPv4:        types.BoolValue(true),
				IPv6:        types.BoolValue(false),
				Ygg:         types.BoolValue(false),
				SSHKey:      types.StringValue("ssh_key"),
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    types.StringNull(),
				NetworkUUID: types.StringValue("network_uuid"),
				PublicIPv4:  types.StringValue("public_ipv4"),
				PublicIPv6:  types.StringValue(""),
				YggIP:       types.StringValue(""),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vmResourceModel{
				ID:          tt.fields.ID,
				Image:       tt.fields.Image,
				Region:      tt.fields.Region,
				CPU:         tt.fields.CPU,
				RAM:         tt.fields.RAM,
				Disk:        tt.fields.Disk,
				DiskType:    tt.fields.DiskType,
				IPv4:        tt.fields.IPv4,
				IPv6:        tt.fields.IPv6,
				Ygg:         tt.fields.Ygg,
				SSHKey:      tt.fields.SSHKey,
				SSHKeyName:  types.StringValue("ssh_key_name"),
				HostName:    tt.fields.HostName,
				NetworkUUID: tt.fields.NetworkUUID,
				PublicIPv4:  tt.fields.PublicIPv4,
				PublicIPv6:  tt.fields.PublicIPv6,
				YggIP:       tt.fields.YggIP,
			}
			got, err := v.ToClientRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToClientRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToClientRequest() got = %v, want %v", got, tt.want)
			}
			assert.Equal(t, tt.fields.ID, v.ID)
			assert.Equal(t, tt.fields.Image, v.Image)
			assert.Equal(t, tt.fields.Region, v.Region)
			assert.Equal(t, tt.fields.CPU, v.CPU)
			assert.Equal(t, tt.fields.RAM, v.RAM)
			assert.Equal(t, tt.fields.Disk, v.Disk)
			assert.Equal(t, tt.fields.DiskType, v.DiskType)
			assert.Equal(t, tt.fields.IPv4, v.IPv4)
			assert.Equal(t, tt.fields.IPv6, v.IPv6)
			assert.Equal(t, tt.fields.Ygg, v.Ygg)
			assert.Equal(t, tt.fields.SSHKey, v.SSHKey)
			assert.Equal(t, tt.fields.SSHKeyName, v.SSHKeyName)
			assert.Equal(t, tt.fields.HostName, v.HostName)
			assert.Equal(t, tt.fields.NetworkUUID, v.NetworkUUID)
			assert.Equal(t, tt.fields.PublicIPv4, v.PublicIPv4)
			assert.Equal(t, tt.fields.PublicIPv6, v.PublicIPv6)
			assert.Equal(t, tt.fields.YggIP, v.YggIP)
		})
	}
}
