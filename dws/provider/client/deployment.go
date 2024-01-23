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

func (c *DWSClient) CreateDeployment(ctx context.Context, r *DeploymentConfig) (*AsyncAPIDeploymentResponse, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to encode deployment: %w", err)
	}

	body := bytes.NewReader(b)
	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, c.url+DeploymentEndpoint, body)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, fmt.Sprintf("Create deployment responseBody: %s", string(responseBody)))

	taskResponse := AsyncAPIDeploymentTask{}
	err = json.Unmarshal(responseBody, &taskResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to create deployment, unmarshal response error: %w", err)
	}

	deploymentResponse, err := c.PollDeploymentTask(ctx, taskResponse.TaskID)
	if err != nil {
		return nil, fmt.Errorf("failed to create deployment: %w", err)
	}

	deploymentResponse.ID = taskResponse.ID

	return deploymentResponse, nil

}

func (c *DWSClient) GetDeployment(ctx context.Context, id string) (*CreatedDeployment, error) {
	tflog.Debug(ctx, fmt.Sprintf("Get deployment by id: %s", id))

	responseBody, err := c.DoSignedRequest(ctx, http.MethodGet, fmt.Sprintf(c.url+DeploymentEndpoint+"/%s", id), nil)
	tflog.Debug(ctx, fmt.Sprintf("Get Deployment responseBody: %s", string(responseBody)))
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment: %w", err)
	}

	deployment := new(CreatedDeployment)

	err = json.Unmarshal(responseBody, deployment)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get deployment response: %w", err)
	}

	return deployment, nil
}

func (c *DWSClient) UpdateDeployment(ctx context.Context, id string, r *DeploymentConfig) (*AsyncAPIDeploymentResponse, error) {
	return nil, errors.New("update not implemented")
}

func (c *DWSClient) DeleteDeployment(ctx context.Context, id string) error {
	tflog.Debug(ctx, fmt.Sprintf("Delete deployment by id: %s", id))

	responseBody, err := c.DoSignedRequest(ctx, http.MethodDelete, fmt.Sprintf(c.url+DeploymentEndpoint+"/%s", id), nil)
	if err != nil {
		return fmt.Errorf("failed to delete deployment: %w", err)
	}

	taskResponse := new(AsyncAPIDeploymentTask)
	err = json.Unmarshal(responseBody, taskResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal delete deployment response: %w", err)
	}

	_, err = c.PollDeploymentTask(ctx, taskResponse.TaskID)
	if err != nil {
		return fmt.Errorf("failed to delete deployment: %w", err)
	}

	return nil
}

func (c *DWSClient) PollDeploymentTask(ctx context.Context, taskID string) (*AsyncAPIDeploymentResponse, error) {
	deploymentResponse := new(AsyncAPIDeploymentResponse)
	ticker := time.NewTicker(5 * time.Second)
pollingCycle:
	for {
		select {
		case <-ticker.C:
			tflog.Debug(ctx, fmt.Sprintf("polling deployment by taskId: %s", taskID))
			b, err := c.DoSignedRequest(ctx, http.MethodGet, c.url+fmt.Sprintf(TaskEndpoint, taskID), nil)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(b, deploymentResponse)
			if err != nil {
				return nil, fmt.Errorf("failed to poll deployment, unmarshal response error: %w", err)
			}

			tflog.Debug(ctx, fmt.Sprintf("polling deployment by taskId: %s", taskID), map[string]interface{}{
				"response": string(b),
			})

			if (deploymentResponse.Data != nil && deploymentResponse.EndTime != nil) || deploymentResponse.IsError {
				break pollingCycle
			}
		case <-ctx.Done():
			return nil, errors.New("failed to poll deployment: context deadline exceeded")
		}
	}

	if deploymentResponse.IsError {
		return nil, fmt.Errorf("failed to poll deployment: %s", deploymentResponse.FailedReason)
	}

	return deploymentResponse, nil
}
