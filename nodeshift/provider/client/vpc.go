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

const errVPCPrefix = "failed to create vpc: %w"

func (c *NodeshiftClient) CreateVPC(ctx context.Context, vpc *CreateVPCRequest) (*GetVPCResponse, error) {
	b, err := json.Marshal(vpc)
	if err != nil {
		return nil, fmt.Errorf(errVPCPrefix, err)
	}

	tflog.Info(ctx, "VPC to create: "+string(b))

	u, err := url.JoinPath(c.url, VPCEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to join create VPC endpoint: %w", err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf(errVPCPrefix, err)
	}

	tflog.Info(ctx, "created VPC: "+string(responseBody))

	cr := new(createVPCResponse)
	if err = json.Unmarshal(responseBody, cr); err != nil {
		return nil, fmt.Errorf(errVPCPrefix, err)
	}

	u, err = url.JoinPath(c.url, VPCEndpoint, cr.UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to join get VPC endpoint: %w", err)
	}

	gvr, err := poll[GetVPCResponse](ctx, c, u, func(gvr *GetVPCResponse) string {
		for _, r := range gvr.Resources {
			if r.Status != "running" {
				return ""
			}
		}

		return "running"
	})
	if err != nil {
		return nil, fmt.Errorf("failed to poll create deployment: %w", err)
	}

	return gvr, nil
}

func (c *NodeshiftClient) GetVPC(ctx context.Context, id string) (*GetVPCResponse, error) {
	tflog.Debug(ctx, "Get VPC by id: "+id)

	u, err := url.JoinPath(c.url, VPCEndpoint, id)
	if err != nil {
		return nil, fmt.Errorf("failed to join get VPC endpoint: %w", err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get VPC: %w", err)
	}

	tflog.Debug(ctx, "Get VPC responseBody: %s"+string(responseBody))

	gr := new(GetVPCResponse)
	err = json.Unmarshal(responseBody, gr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get deployment response body: %w", err)
	}

	return gr, nil
}

func (c *NodeshiftClient) UpdateVPC(ctx context.Context, id string, vpc *CreateVPCRequest) (*GetVPCResponse, error) {
	return nil, fmt.Errorf("failed to update VPC: %w", errNotImplemented)
}

func (c *NodeshiftClient) DeleteVPC(ctx context.Context, id string) error {
	tflog.Debug(ctx, "Delete VPC by id: "+id)

	u, err := url.JoinPath(c.url, VPCEndpoint, id)
	if err != nil {
		return fmt.Errorf("failed to join delete VPC endpoint: %w", err)
	}

	if _, err := c.DoSignedRequest(ctx, http.MethodDelete, u, nil); err != nil {
		return fmt.Errorf("failed to delete VPC: %w", err)
	}

	return nil
}
