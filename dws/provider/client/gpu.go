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

func (c *DWSClient) CreateGPU(ctx context.Context, gpu *GPUConfig) (*GPUConfigResponse, error) {
	errPrefix := "failed to create gpu: %w"
	b, err := json.Marshal(gpu)
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	tflog.Debug(ctx, fmt.Sprintf("GPU to create: %s", string(b)))

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, c.url+GPUEndpoint, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	tflog.Info(ctx, fmt.Sprintf("created GPU: %s", string(responseBody)))
	resp := GPUConfigResponse{}
	if err = json.Unmarshal(responseBody, &resp); err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	return &resp, nil
}

func (c *DWSClient) GetGPU(ctx context.Context, id string) (*RentedGpuInfoResponse, error) {
	tflog.Debug(ctx, fmt.Sprintf("Get GPU by id: %s", id))

	responseBody, err := c.DoSignedRequest(ctx, http.MethodGet, fmt.Sprintf(c.url+GPUEndpoint+"/%s", id), nil)
	tflog.Debug(ctx, fmt.Sprintf("Get GPU responseBody: %s", string(responseBody)))
	if err != nil {
		return nil, fmt.Errorf("failed to get GPU: %w", err)
	}

	gpu := RentedGpuInfoResponse{}
	err = json.Unmarshal(responseBody, &gpu)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get deployment response body: %w", err)
	}

	return &gpu, nil
}

func (c *DWSClient) UpdateGPU(ctx context.Context, id string, gpu *GPUConfig) (*GPUConfig, error) {
	return nil, errors.New("not implemented")
}

func (c *DWSClient) DeleteGPU(ctx context.Context, id string) error {
	tflog.Debug(ctx, fmt.Sprintf("Delete GPU by id: %s", id))

	_, err := c.DoSignedRequest(ctx, http.MethodDelete, fmt.Sprintf(c.url+GPUEndpoint+"/%s", id), nil)
	if err != nil {
		return fmt.Errorf("failed to delete GPU: %w", err)
	}

	return nil
}
