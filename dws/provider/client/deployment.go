package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	DeploymentEndpoint = "/api/terraform/deployment"
	TaskEndpoint       = "/api/task/%s"
)

func (c *DWSClient) CreateDeployment(ctx context.Context, r *VMConfig) (*VMResponse, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to encode deployment: %w", err)
	}

	body := bytes.NewReader(b)
	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, DeploymentEndpoint, body)
	if err != nil {
		return nil, err
	}

	taskResponse := VMResponseTask{}
	err = json.Unmarshal(responseBody, &taskResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to create deployment, unmarshal response error: %w", err)
	}

	ticker := time.NewTicker(5 * time.Second)
	deploymentResponse := new(VMResponse)

pollingCycle:
	for {
		select {
		case <-ticker.C:
			b, err := c.DoSignedRequest(ctx, http.MethodGet, fmt.Sprintf(TaskEndpoint, taskResponse.TaskID), nil)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(b, deploymentResponse)
			if err != nil {
				return nil, fmt.Errorf("failed to create deployment, unmarshal response error: %w", err)
			}

			if deploymentResponse.Data != nil && deploymentResponse.EndTime != nil {
				break pollingCycle
			}
		}
	}

	return deploymentResponse, nil

}

func (c *DWSClient) GetDeployment(ctx context.Context, id string) (*VMResponse, error) {
	// errPrefix := "failed to get deployment: %w"
	// b, err := c.DoSignedRequest(ctx, http.MethodGet, fmt.Sprintf(DeploymentEndpoint+"%s", id), nil)
	// if err != nil {
	// 	return nil, fmt.Errorf(errPrefix, err)
	// }

	// deployment := new(VMResponse)
	// err = json.Unmarshal(b, deployment)
	// if err != nil {
	// 	return nil, fmt.Errorf(errPrefix, err)
	// }

	return nil, errors.New("update not implemented")
}

func (c *DWSClient) UpdateDeployment(ctx context.Context, id string, r *VMConfig) (*VMResponse, error) {
	return nil, errors.New("update not implemented")
}

func (c *DWSClient) DeleteDeployment(ctx context.Context, id string) error {
	return errors.New("update not implemented")
}
