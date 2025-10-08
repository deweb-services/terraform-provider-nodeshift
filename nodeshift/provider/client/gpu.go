package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const errGPUPrefix = "failed to create gpu: %w"

// nolint: dupl
func (c *NodeshiftClient) CreateGPU(ctx context.Context, gpu *CreateGPURequest) (*GetGPUResponse, error) {
	b, err := json.Marshal(gpu)
	if err != nil {
		return nil, fmt.Errorf(errGPUPrefix, err)
	}

	u, err := url.JoinPath(c.url, GPUEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to join create GPU endpoint: %w", err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf(errGPUPrefix, err)
	}

	tflog.Info(ctx, "created GPU: %s"+string(responseBody))

	resp := new(CreateGPUResponse)
	if err = json.Unmarshal(responseBody, resp); err != nil {
		return nil, fmt.Errorf(errGPUPrefix, err)
	}

	u, err = url.JoinPath(c.url, GPUEndpoint, resp.UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to join get GPU endpoint: %w", err)
	}

	ggr, err := poll[GetGPUResponse](ctx, c, u, func(ggr *GetGPUResponse) string { return ggr.Status })
	if err != nil {
		return nil, fmt.Errorf("failed to poll create deployment: %w", err)
	}

	return ggr, nil
}

func (c *NodeshiftClient) GetGPU(ctx context.Context, id string) (*GetGPUResponse, error) {
	tflog.Debug(ctx, "Get GPU by id: %s"+id)

	u, err := url.JoinPath(c.url, GPUEndpoint, id)
	if err != nil {
		return nil, fmt.Errorf("failed to join get GPU endpoint: %w", err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodGet, u, nil)
	tflog.Debug(ctx, "Get GPU responseBody: %s"+string(responseBody))
	if err != nil {
		return nil, fmt.Errorf("failed to get GPU: %w", err)
	}

	gpu := new(GetGPUResponse)
	err = json.Unmarshal(responseBody, gpu)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get deployment response body: %w", err)
	}

	return gpu, nil
}

func (c *NodeshiftClient) UpdateGPU(ctx context.Context, id string, gpu *CreateGPURequest) (*GetGPUResponse, error) {
	return nil, fmt.Errorf("failed to update GPU: %w", errNotImplemented)
}

func (c *NodeshiftClient) DeleteGPU(ctx context.Context, id string) error {
	tflog.Debug(ctx, "Delete GPU by id: %s"+id)

	u, err := url.JoinPath(c.url, GPUEndpoint, id)
	if err != nil {
		return fmt.Errorf("failed to join delete GPU endpoint: %w", err)
	}

	if _, err := c.DoSignedRequest(ctx, http.MethodDelete, u, nil); err != nil {
		return fmt.Errorf("failed to delete GPU: %w", err)
	}

	return nil
}
