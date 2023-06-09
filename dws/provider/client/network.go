package client

import (
	"context"
	"errors"
)

func (c *DWSClient) CreateNetwork(ctx context.Context, network *NetworkConfig) (*NetworkConfig, error) {
	return nil, errors.New("not implemented")
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
