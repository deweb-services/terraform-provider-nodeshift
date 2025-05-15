package client

import (
	"context"
	"io"
	"net/http"
)

type mockedClient struct{}

func NewMockedClient() INodeshiftClient {
	return &mockedClient{}
}

func (m *mockedClient) DoRequest(ctx context.Context, req *http.Request) ([]byte, error) {
	return []byte{}, nil
}

func (m *mockedClient) DoSignedRequest(ctx context.Context, method string, endpoint string, body io.ReadSeeker) ([]byte, error) {
	return []byte{}, nil
}

func (m *mockedClient) SetGlobalTransactionNote(note string) {}

func (m *mockedClient) PollDeploymentTask(ctx context.Context, taskID string) (*AsyncAPIDeploymentResponse, error) {
	return &AsyncAPIDeploymentResponse{}, nil
}

func (m *mockedClient) CreateDeployment(ctx context.Context, r *DeploymentConfig) (*AsyncAPIDeploymentResponse, error) {
	return &AsyncAPIDeploymentResponse{
		Data: &DeploymentResponseData{
			ProviderPlan: nil,
		},
	}, nil
}

func (m *mockedClient) GetDeployment(ctx context.Context, id string) (*CreatedDeployment, error) {
	return &CreatedDeployment{}, nil
}

func (m *mockedClient) UpdateDeployment(ctx context.Context, id string, r *DeploymentConfig) (*AsyncAPIDeploymentResponse, error) {
	return &AsyncAPIDeploymentResponse{}, nil
}

func (m *mockedClient) DeleteDeployment(ctx context.Context, id string) error { return nil }

func (m *mockedClient) CreateGPU(ctx context.Context, gpu *GPUConfig) (*GPUConfigResponse, error) {
	return &GPUConfigResponse{}, nil
}

func (m *mockedClient) GetGPU(ctx context.Context, id string) (*RentedGpuInfoResponse, error) {
	return &RentedGpuInfoResponse{}, nil
}

func (m *mockedClient) UpdateGPU(ctx context.Context, id string, gpu *GPUConfig) (*GPUConfig, error) {
	return &GPUConfig{}, nil
}

func (m *mockedClient) DeleteGPU(ctx context.Context, id string) error {
	return nil
}

func (m *mockedClient) CreateVPC(ctx context.Context, vpc *VPCConfig) (*VPCConfig, error) {
	return &VPCConfig{}, nil
}

func (m *mockedClient) GetVPC(ctx context.Context, id string) (*VPCConfig, error) {
	return &VPCConfig{}, nil
}

func (m *mockedClient) UpdateVPC(ctx context.Context, id string, vpc *VPCConfig) (*VPCConfig, error) {
	return &VPCConfig{}, nil
}

func (m *mockedClient) DeleteVPC(ctx context.Context, id string) error {
	return nil
}

func (m *mockedClient) CreateBucket(ctx context.Context, bucket *S3BucketConfig) (*S3BucketConfig, error) {
	return &S3BucketConfig{}, nil
}

func (m *mockedClient) GetBucket(ctx context.Context, key string) (*S3BucketConfig, error) {
	return &S3BucketConfig{}, nil
}

func (m *mockedClient) UpdateBucket(ctx context.Context, bucket *S3BucketConfig) error {
	return nil
}

func (m *mockedClient) DeleteBucket(ctx context.Context, key string) error {
	return nil
}

func (m *mockedClient) ListRegions(ctx context.Context) ([]string, error) {
	return []string{}, nil
}
