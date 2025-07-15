package gpu

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/deweb-services/terraform-provider-nodeshift/nodeshift/provider/client"
)

func TestGPUResourceModel_FromClientRentedGPUResponse(t *testing.T) {
	type fields struct {
		GPUName        types.String
		Image          types.String
		SSHKey         types.String
		GPUCount       types.Int64
		Region         types.String
		UUID           types.String
		DiskSizeGB     types.Int64
		MinCudaVersion types.String
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
				GPUName:        types.String{},
				Image:          types.String{},
				SSHKey:         types.String{},
				GPUCount:       types.Int64{},
				Region:         types.String{},
				UUID:           types.String{},
				DiskSizeGB:     types.Int64{},
				MinCudaVersion: types.String{},
			},
			args: args{
				c: &client.RentedGpuInfoResponse{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GPUResourceModel{
				GPUName:        tt.fields.GPUName,
				Image:          tt.fields.Image,
				SSHKey:         tt.fields.SSHKey,
				GPUCount:       tt.fields.GPUCount,
				Region:         tt.fields.Region,
				UUID:           tt.fields.UUID,
				DiskSize:       tt.fields.DiskSizeGB,
				MinCudaVersion: tt.fields.MinCudaVersion,
			}
			if err := m.FromClientRentedGPUResponse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FromClientRentedGPUResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGPUResourceModel_FromClientResponse(t *testing.T) {
	type fields struct {
		GPUName        types.String
		Image          types.String
		SSHKey         types.String
		GPUCount       types.Int64
		Region         types.String
		UUID           types.String
		DiskSizeGB     types.Int64
		MinCudaVersion types.String
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
				GPUName:        types.String{},
				Image:          types.String{},
				SSHKey:         types.String{},
				GPUCount:       types.Int64{},
				Region:         types.String{},
				UUID:           types.String{},
				DiskSizeGB:     types.Int64{},
				MinCudaVersion: types.String{},
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
				GPUName:        tt.fields.GPUName,
				Image:          tt.fields.Image,
				SSHKey:         tt.fields.SSHKey,
				GPUCount:       tt.fields.GPUCount,
				Region:         tt.fields.Region,
				UUID:           tt.fields.UUID,
				DiskSize:       tt.fields.DiskSizeGB,
				MinCudaVersion: tt.fields.MinCudaVersion,
			}
			if err := m.FromClientResponse(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FromClientResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGPUResourceModel_ToClientRequest(t *testing.T) {
	type fields struct {
		GPUName        types.String
		Image          types.String
		SSHKey         types.String
		GPUCount       types.Int64
		Region         types.String
		UUID           types.String
		DiskSizeGB     types.Int64
		MinCudaVersion types.String
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
				GPUName:        basetypes.NewStringValue("gpu name"),
				Image:          basetypes.NewStringValue("image"),
				SSHKey:         basetypes.NewStringValue("ssh key"),
				GPUCount:       basetypes.NewInt64Value(2),
				Region:         basetypes.NewStringValue("region"),
				UUID:           basetypes.NewStringValue("uuid"),
				DiskSizeGB:     basetypes.NewInt64Value(30),
				MinCudaVersion: basetypes.NewStringValue("12.6"),
			},
			want: &client.GPUConfig{
				GPUName:        "gpu name",
				Image:          "image",
				SSHKey:         "ssh key",
				GPUCount:       2,
				Region:         "region",
				Disk:           30,
				MinCudaVersion: "12.6",
			},
			wantErr: false,
		},
		{
			name: "gpu resource model to client request convert error gpu",
			fields: fields{
				GPUName:        basetypes.NewStringUnknown(),
				Image:          basetypes.NewStringValue("image"),
				SSHKey:         basetypes.NewStringValue("ssh key"),
				GPUCount:       basetypes.NewInt64Value(2),
				Region:         basetypes.NewStringValue("region"),
				UUID:           basetypes.NewStringValue("uuid"),
				DiskSizeGB:     basetypes.NewInt64Value(30),
				MinCudaVersion: basetypes.NewStringValue("12.6"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "gpu resource model to client request convert error image",
			fields: fields{
				GPUName:        basetypes.NewStringValue("gpu name"),
				Image:          basetypes.NewStringNull(),
				SSHKey:         basetypes.NewStringValue("ssh key"),
				GPUCount:       basetypes.NewInt64Value(2),
				Region:         basetypes.NewStringValue("region"),
				UUID:           basetypes.NewStringValue("uuid"),
				DiskSizeGB:     basetypes.NewInt64Value(30),
				MinCudaVersion: basetypes.NewStringValue("12.6"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "gpu resource model to client request convert error ssh key",
			fields: fields{
				GPUName:        basetypes.NewStringValue("gpu name"),
				Image:          basetypes.NewStringValue("image"),
				SSHKey:         basetypes.NewStringUnknown(),
				GPUCount:       basetypes.NewInt64Value(2),
				Region:         basetypes.NewStringValue("region"),
				UUID:           basetypes.NewStringValue("uuid"),
				DiskSizeGB:     basetypes.NewInt64Value(30),
				MinCudaVersion: basetypes.NewStringValue("12.6"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GPUResourceModel{
				GPUName:        tt.fields.GPUName,
				Image:          tt.fields.Image,
				SSHKey:         tt.fields.SSHKey,
				GPUCount:       tt.fields.GPUCount,
				Region:         tt.fields.Region,
				UUID:           tt.fields.UUID,
				DiskSize:       tt.fields.DiskSizeGB,
				MinCudaVersion: tt.fields.MinCudaVersion,
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
