package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (c *DWSClient) CreateBucket(ctx context.Context, bucket *BucketConfig) (*BucketConfigResponse, error) {
	errPrefix := "failed to create bucket: %w"
	b, err := json.Marshal(bucket)
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	tflog.Info(ctx, fmt.Sprintf("Bucket to create: %s", string(b)))

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, c.url+BucketEndpoint, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	tflog.Info(ctx, fmt.Sprintf("created Bucket: %s", string(responseBody)))
	resp := BucketConfigResponse{}
	if err = json.Unmarshal(responseBody, &resp); err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	return &resp, nil
}

func (c *DWSClient) GetBucket(ctx context.Context, id string) (*RentedGpuInfoResponse, error) {
	tflog.Debug(ctx, fmt.Sprintf("Get Bucket by id: %s", id))

	responseBody, err := c.DoSignedRequest(ctx, http.MethodGet, fmt.Sprintf(c.url+BucketEndpoint+"/%s", id), nil)
	tflog.Debug(ctx, fmt.Sprintf("Get Bucket responseBody: %s", string(responseBody)))
	if err != nil {
		return nil, fmt.Errorf("failed to get Bucket: %w", err)
	}

	bucket := RentedGpuInfoResponse{}
	err = json.Unmarshal(responseBody, &bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get deployment response body: %w", err)
	}

	return &bucket, nil
}

func (c *DWSClient) UpdateBucket(ctx context.Context, id string, bucket *BucketConfig) (*BucketConfig, error) {
	return nil, errors.New("not implemented")
}

func (c *DWSClient) DeleteBucket(ctx context.Context, id string) error {
	tflog.Debug(ctx, fmt.Sprintf("Delete Bucket by id: %s", id))

	_, err := c.DoSignedRequest(ctx, http.MethodDelete, fmt.Sprintf(c.url+BucketEndpoint+"/%s", id), nil)
	if err != nil {
		return fmt.Errorf("failed to delete Bucket: %w", err)
	}

	return nil
}
