package gpu

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/deweb-services/terraform-provider-dws/dws/provider/client"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGPUResourceModel_FromClientRentedGPUResponse(t *testing.T) {
	type fields struct {
		GPUName  types.String
		Image    types.String
		SSHKey   types.String
		GPUCount types.Int64
		Region   types.String
		UUID     types.String
	}
	type args struct {
		c *client.RentedGpuInfoResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "gpu resource model from client rented gpu response",
			fields: fields{
				GPUName:  types.String{},
				Image:    types.String{},
				SSHKey:   types.String{},
				GPUCount: types.Int64{},
				Region:   types.String{},
				UUID:     types.String{},
			},
			args: args{
				c: &client.RentedGpuInfoResponse{
					BundleID:          0,
					BundledResults:    0,
					BwNvlink:          0,
					ComputeCap:        0,
					CpuCores:          0,
					CpuCoresEffective: 0,
					CpuName:           "",
					CpuRam:            0,
					CudaMaxGood:       0,
					DirectPortCount:   0,
					DiskBw:            0,
					DiskName:          "",
					DiskSpace:         0,
					Dlperf:            0,
					DlperfPerDphtotal: 0,
					DphBase:           0,
					DphTotal:          0,
					DriverVersion:     "",
					Duration:          0,
					EndDate:           0,
					External:          false,
					FlopsPerDphtotal:  0,
					Geolocation:       "",
					GpuDisplayActive:  false,
					GpuFrac:           0,
					GpuLanes:          0,
					GpuMemBw:          0,
					GpuName:           "",
					GpuRam:            0,
					HasAvx:            0,
					HostId:            0,
					HostRunTime:       0,
					HostingType:       "",
					Id:                0,
					InetDown:          0,
					InetDownBilled:    0,
					InetDownCost:      0,
					InetUp:            0,
					InetUpBilled:      0,
					InetUpCost:        0,
					IsBid:             false,
					MachineId:         0,
					MinBid:            0,
					MoboName:          "",
					NumGpus:           0,
					PciGen:            0,
					PcieBw:            0,
					PendingCount:      0,
					PublicIpaddr:      "",
					Reliability2:      0,
					Rentable:          false,
					Rented:            false,
					Score:             0,
					StartDate:         0,
					StorageCost:       0,
					StorageTotalCost:  0,
					TotalFlops:        0,
					Verification:      "",
					Webpage:           "",
					ActualStatus:      "",
					CurState:          "",
					DirectPortEnd:     0,
					DirectPortStart:   0,
					DiskUtil:          0,
					ExtraEnv:          nil,
					GpuTemp:           0,
					GpuUtil:           0,
					ImageArgs:         nil,
					ImageRuntype:      "",
					ImageUuid:         "",
					IntendedStatus:    "",
					JupyterToken:      "",
					Label:             "",
					LocalIpaddrs:      "",
					Logo:              "",
					MachineDirSshPort: 0,
					MemLimit:          0,
					MemUsage:          0,
					NextState:         "",
					Onstart:           "",
					SshHost:           "",
					SshIdx:            "",
					SshPort:           0,
					StatusMsg:         "",
					VmemUsage:         0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GPUResourceModel{
				GPUName:  tt.fields.GPUName,
				Image:    tt.fields.Image,
				SSHKey:   tt.fields.SSHKey,
				GPUCount: tt.fields.GPUCount,
				Region:   tt.fields.Region,
				UUID:     tt.fields.UUID,
			}
			if err := m.FromClientRentedGPUResponse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FromClientRentedGPUResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGPUResourceModel_FromClientResponse(t *testing.T) {
	type fields struct {
		GPUName  types.String
		Image    types.String
		SSHKey   types.String
		GPUCount types.Int64
		Region   types.String
		UUID     types.String
	}
	type args struct {
		c *client.GPUConfigResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "gpu resource model from client response",
			fields: fields{
				GPUName:  types.String{},
				Image:    types.String{},
				SSHKey:   types.String{},
				GPUCount: types.Int64{},
				Region:   types.String{},
				UUID:     types.String{},
			},
			args: args{
				c: &client.GPUConfigResponse{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GPUResourceModel{
				GPUName:  tt.fields.GPUName,
				Image:    tt.fields.Image,
				SSHKey:   tt.fields.SSHKey,
				GPUCount: tt.fields.GPUCount,
				Region:   tt.fields.Region,
				UUID:     tt.fields.UUID,
			}
			if err := m.FromClientResponse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FromClientResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGPUResourceModel_ToClientRequest(t *testing.T) {
	type fields struct {
		GPUName  types.String
		Image    types.String
		SSHKey   types.String
		GPUCount types.Int64
		Region   types.String
		UUID     types.String
	}
	tests := []struct {
		name    string
		fields  fields
		want    *client.GPUConfig
		wantErr bool
	}{
		{
			name: "gpu resource model to client request",
			fields: fields{
				GPUName:  basetypes.NewStringValue("gpu name"),
				Image:    basetypes.NewStringValue("image"),
				SSHKey:   basetypes.NewStringValue("ssh key"),
				GPUCount: basetypes.NewInt64Value(2),
				Region:   basetypes.NewStringValue("region"),
				UUID:     basetypes.NewStringValue("uuid"),
			},
			want: &client.GPUConfig{
				GPUName:  "gpu_name",
				Image:    "image",
				SSHKey:   "ssh key",
				GPUCount: 2,
				Region:   "region",
			},
			wantErr: false,
		},
		{
			name: "gpu resource model to client request convert error gpu",
			fields: fields{
				GPUName:  basetypes.NewStringUnknown(),
				Image:    basetypes.NewStringValue("image"),
				SSHKey:   basetypes.NewStringValue("ssh key"),
				GPUCount: basetypes.NewInt64Value(2),
				Region:   basetypes.NewStringValue("region"),
				UUID:     basetypes.NewStringValue("uuid"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "gpu resource model to client request convert error image",
			fields: fields{
				GPUName:  basetypes.NewStringValue("gpu name"),
				Image:    basetypes.NewStringNull(),
				SSHKey:   basetypes.NewStringValue("ssh key"),
				GPUCount: basetypes.NewInt64Value(2),
				Region:   basetypes.NewStringValue("region"),
				UUID:     basetypes.NewStringValue("uuid"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "gpu resource model to client request convert error ssh key",
			fields: fields{
				GPUName:  basetypes.NewStringValue("gpu name"),
				Image:    basetypes.NewStringValue("image"),
				SSHKey:   basetypes.NewStringUnknown(),
				GPUCount: basetypes.NewInt64Value(2),
				Region:   basetypes.NewStringValue("region"),
				UUID:     basetypes.NewStringValue("uuid"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GPUResourceModel{
				GPUName:  tt.fields.GPUName,
				Image:    tt.fields.Image,
				SSHKey:   tt.fields.SSHKey,
				GPUCount: tt.fields.GPUCount,
				Region:   tt.fields.Region,
				UUID:     tt.fields.UUID,
			}
			got, err := m.ToClientRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToClientRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToClientRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
