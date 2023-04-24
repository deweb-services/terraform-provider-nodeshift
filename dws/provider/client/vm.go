package client

import "context"

func (c *DWSClient) CreateVM(ctx context.Context, r *VMConfig) (*VMConfig, error) {
	return r, nil
}

func (c *DWSClient) GetVM(ctx context.Context, id string) (*VMConfig, error) {
	return &VMConfig{}, nil
}

func (c *DWSClient) UpdateVM(ctx context.Context, id string, r *VMConfig) (*VMConfig, error) {
	return &VMConfig{}, nil
}

func (c *DWSClient) DeleteVM(ctx context.Context, id string) error {
	return nil
}
