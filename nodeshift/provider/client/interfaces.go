package client

import (
	"context"
	"io"
	"net/http"
)

//go:generate mockgen -source interfaces.go -destination=./interfaces_mocks.go -package=client

type INodeshiftClient interface {
	DoRequest(ctx context.Context, req *http.Request) ([]byte, error)
	DoSignedRequest(ctx context.Context, method string, endpoint string, body io.ReadSeeker) ([]byte, error)
	SetGlobalTransactionNote(note string)

	CreateDeployment(ctx context.Context, r *CreateDeploymentRequest) (*GetDeploymentResponse, error)
	GetDeployment(ctx context.Context, id string) (*GetDeploymentResponse, error)
	UpdateDeployment(ctx context.Context, id string, r *CreateDeploymentRequest) (*GetDeploymentResponse, error)
	DeleteDeployment(ctx context.Context, id string) error

	CreateGPU(ctx context.Context, gpu *CreateGPURequest) (*GetGPUResponse, error)
	GetGPU(ctx context.Context, id string) (*GetGPUResponse, error)
	UpdateGPU(ctx context.Context, id string, gpu *CreateGPURequest) (*GetGPUResponse, error)
	DeleteGPU(ctx context.Context, id string) error

	CreateVPC(ctx context.Context, vpc *CreateVPCRequest) (*GetVPCResponse, error)
	GetVPC(ctx context.Context, id string) (*GetVPCResponse, error)
	UpdateVPC(ctx context.Context, id string, vpc *CreateVPCRequest) (*GetVPCResponse, error)
	DeleteVPC(ctx context.Context, id string) error

	CreateBucket(ctx context.Context, bucket *S3BucketConfig) (*S3BucketConfig, error)
	GetBucket(ctx context.Context, key string) (*S3BucketConfig, error)
	UpdateBucket(ctx context.Context, bucket *S3BucketConfig) error
	DeleteBucket(ctx context.Context, key string) error

	ListRegions(ctx context.Context) ([]string, error)

	CreateLB(ctx context.Context, lb *CreateLBRequest) (*GetLBResponse, error)
	GetLB(ctx context.Context, uuid string) (*GetLBResponse, error)
	UpdateLB(ctx context.Context, uuid string, lb *CreateLBRequest) (*GetLBResponse, error)
	DeleteLB(ctx context.Context, uuid string) error
}
