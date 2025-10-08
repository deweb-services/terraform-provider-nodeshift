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

const errLBPrefix = "failed to create lb: %w"

// nolint: dupl
func (c *NodeshiftClient) CreateLB(ctx context.Context, lb *CreateLBRequest) (*GetLBResponse, error) {
	b, err := json.Marshal(lb)
	if err != nil {
		return nil, fmt.Errorf(errLBPrefix, err)
	}

	u, err := url.JoinPath(c.url, LBEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to join create LB endpoint: %w", err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf(errLBPrefix, err)
	}

	tflog.Info(ctx, "created LB: %s"+string(responseBody))

	resp := new(createLBResponse)
	if err = json.Unmarshal(responseBody, resp); err != nil {
		return nil, fmt.Errorf(errLBPrefix, err)
	}

	u, err = url.JoinPath(c.url, LBEndpoint, resp.UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to join gets LB endpoint: %w", err)
	}

	glr, err := poll[GetLBResponse](ctx, c, u, func(glr *GetLBResponse) string { return glr.Status })
	if err != nil {
		return nil, fmt.Errorf("failed to poll create LB: %w", err)
	}

	return glr, nil
}

func (c *NodeshiftClient) GetLB(ctx context.Context, uuid string) (*GetLBResponse, error) {
	tflog.Debug(ctx, "Get LB by id: %s"+uuid)

	u, err := url.JoinPath(c.url, LBEndpoint, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to join get LB endpoint: %w", err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodGet, u, nil)
	tflog.Debug(ctx, "Get LB responseBody: %s"+string(responseBody))
	if err != nil {
		return nil, fmt.Errorf("failed to get LB: %w", err)
	}

	lb := new(GetLBResponse)
	err = json.Unmarshal(responseBody, lb)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get LB response body: %w, %s", err, string(responseBody))
	}

	return lb, nil
}

func (c *NodeshiftClient) UpdateLB(ctx context.Context, uuid string, lb *CreateLBRequest) (*GetLBResponse, error) {
	return nil, fmt.Errorf("failed to update LB: %w", errNotImplemented)
}

func (c *NodeshiftClient) DeleteLB(ctx context.Context, uuid string) error {
	tflog.Debug(ctx, "Delete LB by id: %s"+uuid)

	u, err := url.JoinPath(c.url, LBEndpoint, uuid)
	if err != nil {
		return fmt.Errorf("failed to join delete LB endpoint: %w", err)
	}

	if _, err := c.DoSignedRequest(ctx, http.MethodDelete, u, nil); err != nil {
		return fmt.Errorf("failed to delete LB: %w", err)
	}

	return nil
}
