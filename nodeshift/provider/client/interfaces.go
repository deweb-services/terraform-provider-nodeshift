package client

import (
	"context"
	"io"
	"net/http"
)

type INodeshiftClient interface {
	DoRequest(ctx context.Context, req *http.Request) ([]byte, error)
	DoSignedRequest(ctx context.Context, method string, endpoint string, body io.ReadSeeker) ([]byte, error)
	SetGlobalTransactionNote(note string)
	PollDeploymentTask(ctx context.Context, taskID string) (*AsyncAPIDeploymentResponse, error)

	CreateDeployment(ctx context.Context, r *DeploymentConfig) (*AsyncAPIDeploymentResponse, error)
	GetDeployment(ctx context.Context, id string) (*CreatedDeployment, error)
	UpdateDeployment(ctx context.Context, id string, r *DeploymentConfig) (*AsyncAPIDeploymentResponse, error)
	DeleteDeployment(ctx context.Context, id string) error

	CreateGPU(ctx context.Context, gpu *GPUConfig) (*GPUConfigResponse, error)
	GetGPU(ctx context.Context, id string) (*RentedGpuInfoResponse, error)
	UpdateGPU(ctx context.Context, id string, gpu *GPUConfig) (*GPUConfig, error)
	DeleteGPU(ctx context.Context, id string) error

	CreateVPC(ctx context.Context, vpc *VPCConfig) (*VPCConfig, error)
	GetVPC(ctx context.Context, id string) (*VPCConfig, error)
	UpdateVPC(ctx context.Context, id string, vpc *VPCConfig) (*VPCConfig, error)
	DeleteVPC(ctx context.Context, id string) error

	CreateBucket(ctx context.Context, bucket *S3BucketConfig) (*S3BucketConfig, error)
	GetBucket(ctx context.Context, key string) (*S3BucketConfig, error)
	UpdateBucket(ctx context.Context, bucket *S3BucketConfig) error
	DeleteBucket(ctx context.Context, key string) error
}
