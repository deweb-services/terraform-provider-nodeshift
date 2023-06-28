package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	DeploymentEndpoint = "/api/terraform/deployment"
	TaskEndpoint       = "/api/task/%s"
)

func (c *DWSClient) CreateDeployment(ctx context.Context, r *DeploymentConfig) (*CreatedDeployment, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to encode deployment: %w", err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Deployment to create json: %s", string(b)))

	body := bytes.NewReader(b)
	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, c.url+DeploymentEndpoint, body)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("Create deployment responseBody: %s", string(responseBody)))

	taskResponse := DeploymentCreateTask{}
	err = json.Unmarshal(responseBody, &taskResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to create deployment, unmarshal response error: %w", err)
	}

	ticker := time.NewTicker(5 * time.Second)
	deploymentResponse := new(CreatedDeployment)

pollingCycle:
	for {
		select {
		case <-ticker.C:
			tflog.Debug(ctx, fmt.Sprintf("polling deployment by taskId: %s", taskResponse.TaskID))
			b, err := c.DoSignedRequest(ctx, http.MethodGet, c.url+fmt.Sprintf(TaskEndpoint, taskResponse.TaskID), nil)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(b, deploymentResponse)
			if err != nil {
				return nil, fmt.Errorf("failed to create deployment, unmarshal response error: %w", err)
			}

			tflog.Debug(ctx, fmt.Sprintf("polling deployment by taskId: %s", taskResponse.TaskID), map[string]interface{}{
				"response": string(b),
			})

			if deploymentResponse.Data != nil && deploymentResponse.EndTime != nil {
				break pollingCycle
			}
		case <-ctx.Done():
			return nil, errors.New("failed to create deployment: context deadline exceeded")
		}
	}

	deploymentResponse.ID = taskResponse.ID

	return deploymentResponse, nil

}

func (c *DWSClient) GetDeployment(ctx context.Context, id string) (*CreatedDeployment, error) {
	return nil, errors.New("update not implemented")
}

func (c *DWSClient) UpdateDeployment(ctx context.Context, id string, r *DeploymentConfig) (*CreatedDeployment, error) {
	return nil, errors.New("update not implemented")
}

func (c *DWSClient) DeleteDeployment(ctx context.Context, id string) error {
	return errors.New("update not implemented")
}
