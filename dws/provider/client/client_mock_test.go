package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMockedClient(t *testing.T) {
	tests := []struct {
		name string
		want IDWSClient
	}{
		{
			want: NewMockedClient(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewMockedClient(), "NewMockedClient()")
		})
	}
}

func Test_mockedClient_CreateBucket(t *testing.T) {
	type args struct {
		ctx    context.Context
		bucket *S3BucketConfig
	}
	tests := []struct {
		name string
		args args
		want *S3BucketConfig
	}{
		{
			name: "mocked client create bucket",
			args: args{
				ctx:    context.TODO(),
				bucket: &S3BucketConfig{},
			},
			want: &S3BucketConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.CreateBucket(tt.args.ctx, tt.args.bucket)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "CreateBucket(%v, %v)", tt.args.ctx, tt.args.bucket)
		})
	}
}

func Test_mockedClient_CreateDeployment(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *DeploymentConfig
	}
	tests := []struct {
		name string
		args args
		want *AsyncAPIDeploymentResponse
	}{
		{
			name: "mocked client create deployment",
			args: args{},
			want: &AsyncAPIDeploymentResponse{
				Data: &DeploymentResponseData{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.CreateDeployment(tt.args.ctx, tt.args.r)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "CreateDeployment(%v, %v)", tt.args.ctx, tt.args.r)
		})
	}
}

func Test_mockedClient_CreateGPU(t *testing.T) {
	type args struct {
		ctx context.Context
		gpu *GPUConfig
	}
	tests := []struct {
		name string
		args args
		want *GPUConfigResponse
	}{
		{
			name: "mocked client create gpu",
			args: args{
				ctx: context.TODO(),
				gpu: &GPUConfig{},
			},
			want: &GPUConfigResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.CreateGPU(tt.args.ctx, tt.args.gpu)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "CreateGPU(%v, %v)", tt.args.ctx, tt.args.gpu)
		})
	}
}

func Test_mockedClient_CreateVPC(t *testing.T) {
	type args struct {
		ctx context.Context
		vpc *VPCConfig
	}
	tests := []struct {
		name string
		args args
		want *VPCConfig
	}{
		{
			name: "mocked client create vpc",
			args: args{
				ctx: context.TODO(),
				vpc: &VPCConfig{},
			},
			want: &VPCConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.CreateVPC(tt.args.ctx, tt.args.vpc)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "CreateVPC(%v, %v)", tt.args.ctx, tt.args.vpc)
		})
	}
}

func Test_mockedClient_DeleteBucket(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "mocked client delete bucket",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			err := m.DeleteBucket(tt.args.ctx, tt.args.key)
			assert.NoError(t, err)
		})
	}
}

func Test_mockedClient_DeleteDeployment(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "mocked client delete deployment",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			assert.NoError(t, m.DeleteDeployment(tt.args.ctx, tt.args.id))
		})
	}
}

func Test_mockedClient_DeleteGPU(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "mocked client delete gpu",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			assert.NoError(t, m.DeleteGPU(tt.args.ctx, tt.args.id), fmt.Sprintf("DeleteGPU(%v, %v)", tt.args.ctx, tt.args.id))
		})
	}
}

func Test_mockedClient_DeleteVPC(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "mocked client delete vpc",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			assert.NoError(t, m.DeleteVPC(tt.args.ctx, tt.args.id), fmt.Sprintf("DeleteVPC(%v, %v)", tt.args.ctx, tt.args.id))
		})
	}
}

func Test_mockedClient_DoRequest(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "mocked client do request",
			args: args{},
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.DoRequest(tt.args.req)
			assert.NoError(t, err, fmt.Sprintf("DoRequest(%v)", tt.args.req))
			assert.Equalf(t, tt.want, got, "DoRequest(%v)", tt.args.req)
		})
	}
}

func Test_mockedClient_DoSignedRequest(t *testing.T) {
	type args struct {
		ctx      context.Context
		method   string
		endpoint string
		body     io.ReadSeeker
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "mocked client do request",
			args: args{},
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.DoSignedRequest(tt.args.ctx, tt.args.method, tt.args.endpoint, tt.args.body)
			assert.NoError(t, err, fmt.Sprintf("DoSignedRequest(%v, %v, %v, %v)", tt.args.ctx, tt.args.method, tt.args.endpoint, tt.args.body))
			assert.Equalf(t, tt.want, got, "DoSignedRequest(%v, %v, %v, %v)", tt.args.ctx, tt.args.method, tt.args.endpoint, tt.args.body)
		})
	}
}

func Test_mockedClient_GetBucket(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name string
		args args
		want *S3BucketConfig
	}{
		{
			name: "mocked client get bucket",
			args: args{},
			want: &S3BucketConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.GetBucket(tt.args.ctx, tt.args.key)
			assert.NoError(t, err, fmt.Sprintf("GetBucket(%v, %v)", tt.args.ctx, tt.args.key))
			assert.Equalf(t, tt.want, got, "GetBucket(%v, %v)", tt.args.ctx, tt.args.key)
		})
	}
}

func Test_mockedClient_GetDeployment(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
		want *CreatedDeployment
	}{
		{
			name: "mocked client get deployment",
			args: args{},
			want: &CreatedDeployment{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.GetDeployment(tt.args.ctx, tt.args.id)
			assert.NoError(t, err, fmt.Sprintf("GetDeployment(%v, %v)", tt.args.ctx, tt.args.id))
			assert.Equalf(t, tt.want, got, "GetDeployment(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}

func Test_mockedClient_GetGPU(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
		want *RentedGpuInfoResponse
	}{
		{
			name: "mocked client get gpu",
			args: args{},
			want: &RentedGpuInfoResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.GetGPU(tt.args.ctx, tt.args.id)
			assert.NoError(t, err, fmt.Sprintf("GetGPU(%v, %v)", tt.args.ctx, tt.args.id))
			assert.Equalf(t, tt.want, got, "GetGPU(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}

func Test_mockedClient_GetVPC(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name string
		args args
		want *VPCConfig
	}{
		{
			name: "mocked client get vpc",
			args: args{},
			want: &VPCConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.GetVPC(tt.args.ctx, tt.args.id)
			assert.NoError(t, err, fmt.Sprintf("GetVPC(%v, %v)", tt.args.ctx, tt.args.id))
			assert.Equalf(t, tt.want, got, "GetVPC(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}

func Test_mockedClient_PollDeploymentTask(t *testing.T) {
	type args struct {
		ctx    context.Context
		taskID string
	}
	tests := []struct {
		name string
		args args
		want *AsyncAPIDeploymentResponse
	}{
		{
			name: "mocked client poll deployment task",
			args: args{},
			want: &AsyncAPIDeploymentResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.PollDeploymentTask(tt.args.ctx, tt.args.taskID)
			assert.NoError(t, err, fmt.Sprintf("PollDeploymentTask(%v, %v)", tt.args.ctx, tt.args.taskID))
			assert.Equalf(t, tt.want, got, "PollDeploymentTask(%v, %v)", tt.args.ctx, tt.args.taskID)
		})
	}
}

func Test_mockedClient_SetGlobalTransactionNote(t *testing.T) {
	type args struct {
		note string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "mocked client set global transaction note",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			m.SetGlobalTransactionNote(tt.args.note)
		})
	}
}

func Test_mockedClient_UpdateBucket(t *testing.T) {
	type args struct {
		ctx    context.Context
		bucket *S3BucketConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "mocked client update bucket",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			assert.NoError(t, m.UpdateBucket(tt.args.ctx, tt.args.bucket), fmt.Sprintf("UpdateBucket(%v, %v)", tt.args.ctx, tt.args.bucket))
		})
	}
}

func Test_mockedClient_UpdateDeployment(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
		r   *DeploymentConfig
	}
	tests := []struct {
		name string
		args args
		want *AsyncAPIDeploymentResponse
	}{
		{
			name: "mocked client update deployment",
			args: args{},
			want: &AsyncAPIDeploymentResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.UpdateDeployment(tt.args.ctx, tt.args.id, tt.args.r)
			assert.NoError(t, err, fmt.Sprintf("UpdateDeployment(%v, %v, %v)", tt.args.ctx, tt.args.id, tt.args.r))
			assert.Equalf(t, tt.want, got, "UpdateDeployment(%v, %v, %v)", tt.args.ctx, tt.args.id, tt.args.r)
		})
	}
}

func Test_mockedClient_UpdateGPU(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
		gpu *GPUConfig
	}
	tests := []struct {
		name string
		args args
		want *GPUConfig
	}{
		{
			name: "mocked client update gpu",
			args: args{},
			want: &GPUConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.UpdateGPU(tt.args.ctx, tt.args.id, tt.args.gpu)
			assert.NoError(t, err, fmt.Sprintf("UpdateGPU(%v, %v, %v)", tt.args.ctx, tt.args.id, tt.args.gpu))
			assert.Equalf(t, tt.want, got, "UpdateGPU(%v, %v, %v)", tt.args.ctx, tt.args.id, tt.args.gpu)
		})
	}
}

func Test_mockedClient_UpdateVPC(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
		vpc *VPCConfig
	}
	tests := []struct {
		name string
		args args
		want *VPCConfig
	}{
		{
			name: "mocked client update vpc",
			args: args{},
			want: &VPCConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockedClient{}
			got, err := m.UpdateVPC(tt.args.ctx, tt.args.id, tt.args.vpc)
			assert.NoError(t, err, fmt.Sprintf("UpdateVPC(%v, %v, %v)", tt.args.ctx, tt.args.id, tt.args.vpc))
			assert.Equalf(t, tt.want, got, "UpdateVPC(%v, %v, %v)", tt.args.ctx, tt.args.id, tt.args.vpc)
		})
	}
}
