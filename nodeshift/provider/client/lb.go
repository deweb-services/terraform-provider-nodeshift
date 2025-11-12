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

// nolint: dupl
func (c *NodeshiftClient) CreateLB(ctx context.Context, lb *CreateLBRequest) (*GetLBResponse, error) {
	b, err := json.Marshal(lb)
	if err != nil {
		return nil, fmt.Errorf("failed to encode create load balancer request: %w", err)
	}

	u, err := url.JoinPath(c.url, LBEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to join create load balancer endpoint: %w", err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("failed to do signed request for create load balancer: %w", err)
	}

	tflog.Info(ctx, "created load balancer: "+string(responseBody))

	resp := new(createLBResponse)
	if err = json.Unmarshal(responseBody, resp); err != nil {
		return nil, fmt.Errorf("failed to decode create load balancer response: %w", err)
	}

	u, err = url.JoinPath(c.url, LBEndpoint, resp.UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to join gets load balancer endpoint: %w", err)
	}

	tctx, cancel := context.WithTimeout(ctx, fiveMinuteTimeout)
	defer cancel()

	glr, err := poll[GetLBResponse](tctx, c, u, func(glr *GetLBResponse) string { return glr.Status })
	if err != nil {
		return nil, fmt.Errorf("failed to poll create load balancer: %w", err)
	}

	return glr, nil
}

func (c *NodeshiftClient) GetLB(ctx context.Context, uuid string) (*GetLBResponse, error) {
	tflog.Debug(ctx, "Get load balancer by id: "+uuid)

	u, err := url.JoinPath(c.url, LBEndpoint, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to join get load balancer endpoint: %w", err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodGet, u, nil)
	tflog.Debug(ctx, "Get load balancer responseBody: "+string(responseBody))
	if err != nil {
		return nil, fmt.Errorf("failed to do signed request for get load balancer: %w", err)
	}

	lb := new(GetLBResponse)
	err = json.Unmarshal(responseBody, lb)
	if err != nil {
		return nil, fmt.Errorf("failed to decode get load balancer response body: %w, %s", err, string(responseBody))
	}

	return lb, nil
}

func (c *NodeshiftClient) UpdateLB(ctx context.Context, uuid string, lb *CreateLBRequest) (*GetLBResponse, error) {
	return nil, fmt.Errorf("failed to update load balancer: %w", errNotImplemented)
}

func (c *NodeshiftClient) DeleteLB(ctx context.Context, uuid string) error {
	tflog.Debug(ctx, "Delete load balancer by id: "+uuid)

	u, err := url.JoinPath(c.url, LBEndpoint, uuid)
	if err != nil {
		return fmt.Errorf("failed to join delete load balancer endpoint: %w", err)
	}

	if _, err := c.DoSignedRequest(ctx, http.MethodDelete, u, nil); err != nil {
		return fmt.Errorf("failed to do signed request for delete load balancer: %w", err)
	}

	return nil
}
