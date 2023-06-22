package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	VPCEndpoint = "/terraform/vpc"
)

func (c *DWSClient) CreateVPC(ctx context.Context, vpc *VPCConfig) (*VPCConfig, error) {
	errPrefix := "failed to create vpc: %w"
	b, err := json.Marshal(vpc)
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, c.url+VPCEndpoint, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	if err = json.Unmarshal(responseBody, vpc); err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	return vpc, nil
}

func (c *DWSClient) GetVPC(ctx context.Context, id string) (*VPCConfig, error) {
	return nil, errors.New("not implemented")
}

func (c *DWSClient) UpdateVPC(ctx context.Context, id string, vpc *VPCConfig) (*VPCConfig, error) {
	return nil, errors.New("not implemented")
}

func (c *DWSClient) DeleteVPC(ctx context.Context, id string) error {
	return errors.New("not implemented")
}
