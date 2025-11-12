package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const fiveMinuteTimeout = 5 * time.Minute

func (c *NodeshiftClient) CreateVPC(ctx context.Context, vpc *CreateVPCRequest) (*GetVPCResponse, error) {
	b, err := json.Marshal(vpc)
	if err != nil {
		return nil, fmt.Errorf("failed to encode create vpc request: %w", err)
	}

	tflog.Info(ctx, "VPC to create: "+string(b))

	u, err := url.JoinPath(c.url, VPCEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to join create VPC endpoint: %w", err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, u, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("failed to do signed request for vpc create: %w", err)
	}

	tflog.Info(ctx, "created VPC: "+string(responseBody))

	cr := new(createVPCResponse)
	if err = json.Unmarshal(responseBody, cr); err != nil {
		return nil, fmt.Errorf("failed to decode vpc create response: %w", err)
	}

	u, err = url.JoinPath(c.url, VPCEndpoint, cr.UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to join get VPC endpoint: %w", err)
	}

	tctx, cancel := context.WithTimeout(ctx, fiveMinuteTimeout)
	defer cancel()

	gvr, err := poll[GetVPCResponse](tctx, c, u, func(gvr *GetVPCResponse) string { return gvr.Status })
	if err != nil {
		return nil, fmt.Errorf("failed to poll create VPC: %w", err)
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
		return nil, fmt.Errorf("failed to do signed request for vpc get: %w", err)
	}

	tflog.Debug(ctx, "Get VPC responseBody: "+string(responseBody))

	gr := new(GetVPCResponse)
	err = json.Unmarshal(responseBody, gr)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get vpc response: %w", err)
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
