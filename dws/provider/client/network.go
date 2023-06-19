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
	NetworkEndpoint = "/api/terraform/vpc"
)

func (c *DWSClient) CreateNetwork(ctx context.Context, network *NetworkConfig) (*NetworkConfig, error) {
	errPrefix := "failed to create network: %w"
	b, err := json.Marshal(network)
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, NetworkEndpoint, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	if err = json.Unmarshal(responseBody, network); err != nil {
		return nil, fmt.Errorf(errPrefix, err)
	}

	return network, nil
}

func (c *DWSClient) GetNetwork(ctx context.Context, id string) (*NetworkConfig, error) {
	return nil, errors.New("not implemented")
}

func (c *DWSClient) UpdateNetwork(ctx context.Context, id string, network *NetworkConfig) (*NetworkConfig, error) {
	return nil, errors.New("not implemented")
}

func (c *DWSClient) DeleteNetwork(ctx context.Context, id string) error {
	return errors.New("not implemented")
}
