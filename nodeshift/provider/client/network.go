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

func (c *NodeshiftClient) CreateVPC(ctx context.Context, vpc *VPCConfig) (*VPCConfig, error) {
	errPrefix := "failed to create vpc: %w"
	b, err := json.Marshal(vpc)
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	tflog.Info(ctx, fmt.Sprintf("VPC to create: %s", string(b)))

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, c.url+VPCEndpoint, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	tflog.Info(ctx, fmt.Sprintf("created VPC: %s", string(responseBody)))

	if err = json.Unmarshal(responseBody, vpc); err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	return vpc, nil
}

func (c *NodeshiftClient) GetVPC(ctx context.Context, id string) (*VPCConfig, error) {
	tflog.Debug(ctx, fmt.Sprintf("Get VPC by id: %s", id))

	responseBody, err := c.DoSignedRequest(ctx, http.MethodGet, fmt.Sprintf(c.url+VPCEndpoint+"/%s", id), nil)
	tflog.Debug(ctx, fmt.Sprintf("Get VPC responseBody: %s", string(responseBody)))
	if err != nil {
		return nil, fmt.Errorf("failed to get VPC: %w", err)
	}

	vpc := VPCConfigResponse{}
	err = json.Unmarshal(responseBody, &vpc)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get deployment response body: %w", err)
	}

	return &VPCConfig{
		Name:        vpc.Name,
		Description: vpc.Description,
		IPRange:     vpc.IPRange,
		ID:          vpc.ID,
	}, nil
}

func (c *NodeshiftClient) UpdateVPC(ctx context.Context, id string, vpc *VPCConfig) (*VPCConfig, error) {
	return nil, errors.New("not implemented")
}

func (c *NodeshiftClient) DeleteVPC(ctx context.Context, id string) error {
	tflog.Debug(ctx, fmt.Sprintf("Delete VPC by id: %s", id))

	_, err := c.DoSignedRequest(ctx, http.MethodDelete, fmt.Sprintf(c.url+VPCEndpoint+"/%s", id), nil)
	if err != nil {
		return fmt.Errorf("failed to delete VPC: %w", err)
	}

	return nil
}
