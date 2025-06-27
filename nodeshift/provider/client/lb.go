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

func (c *NodeshiftClient) CreateLB(ctx context.Context, lb *LoadBalancerConfig) (*LoadBalancerConfigResponse, error) {
	errPrefix := "failed to create lb: %w"
	b, err := json.Marshal(lb)
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, c.url+LBEndpoint, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	tflog.Info(ctx, fmt.Sprintf("created LB: %s", string(responseBody)))
	resp := LoadBalancerConfigResponse{}
	if err = json.Unmarshal(responseBody, &resp); err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	return &resp, nil
}

func (c *NodeshiftClient) GetLB(ctx context.Context, uuid string) (*GetLBResponse, error) {
	tflog.Debug(ctx, fmt.Sprintf("Get LB by id: %s", uuid))

	responseBody, err := c.DoSignedRequest(ctx, http.MethodGet, fmt.Sprintf(c.url+LBEndpoint+"/%s", uuid), nil)
	tflog.Debug(ctx, fmt.Sprintf("Get LB responseBody: %s", string(responseBody)))
	if err != nil {
		return nil, fmt.Errorf("failed to get LB: %w", err)
	}

	lb := GetLBResponse{}
	err = json.Unmarshal(responseBody, &lb)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get load_balancer response body: %w, %s", err, string(responseBody))
	}

	return &lb, nil
}

func (c *NodeshiftClient) UpdateLB(ctx context.Context, uuid string, lb *LoadBalancerConfig) (*LoadBalancerConfigResponse, error) {
	return nil, errors.New("not implemented")
}

func (c *NodeshiftClient) DeleteLB(ctx context.Context, uuid string) error {
	tflog.Debug(ctx, fmt.Sprintf("Delete LB by id: %s", uuid))

	_, err := c.DoSignedRequest(ctx, http.MethodDelete, fmt.Sprintf(c.url+LBEndpoint+"/%s", uuid), nil)
	if err != nil {
		return fmt.Errorf("failed to delete LB: %w", err)
	}

	return nil
}
