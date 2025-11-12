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

func (c *NodeshiftClient) CreateDeployment(ctx context.Context, r *CreateDeploymentRequest) (*GetDeploymentResponse, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to encode deployment: %w", err)
	}

	u, err := url.JoinPath(c.url, deploymentEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to join create deployment endpoint: %w", err)
	}

	body := bytes.NewReader(b)
	responseBody, err := c.DoSignedRequest(ctx, http.MethodPost, u, body)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "created Deployment: "+string(responseBody))

	cr := new(createDeploymentResponse)
	err = json.Unmarshal(responseBody, cr)
	if err != nil {
		return nil, fmt.Errorf("failed to decosde create deployment: %w", err)
	}

	u, err = url.JoinPath(c.url, deploymentEndpoint, cr.UUID)
	if err != nil {
		return nil, fmt.Errorf("failed to join get deployment endpoint: %w", err)
	}

	tctx, cancel := context.WithTimeout(ctx, fiveMinuteTimeout)
	defer cancel()

	gdr, err := poll[GetDeploymentResponse](tctx, c, u, func(gdr *GetDeploymentResponse) string { return gdr.Status })
	if err != nil {
		return nil, fmt.Errorf("failed to poll create deployment: %w", err)
	}

	return gdr, nil
}

func (c *NodeshiftClient) GetDeployment(ctx context.Context, id string) (*GetDeploymentResponse, error) {
	tflog.Debug(ctx, "Get deployment by id: "+id)

	u, err := url.JoinPath(c.url, deploymentEndpoint, id)
	if err != nil {
		return nil, fmt.Errorf("failed to join get deployment endpoint: %w", err)
	}

	responseBody, err := c.DoSignedRequest(ctx, http.MethodGet, u, nil)
	tflog.Debug(ctx, "Get Deployment responseBody: "+string(responseBody))
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment: %w", err)
	}

	deployment := new(GetDeploymentResponse)
	err = json.Unmarshal(responseBody, deployment)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal get deployment response: %w", err)
	}

	return deployment, nil
}

func (c *NodeshiftClient) UpdateDeployment(ctx context.Context, id string, r *CreateDeploymentRequest) (*GetDeploymentResponse, error) {
	return nil, fmt.Errorf("failed to update deployment: %w", errNotImplemented)
}

func (c *NodeshiftClient) DeleteDeployment(ctx context.Context, id string) error {
	tflog.Debug(ctx, "Delete deployment by id: "+id)

	u, err := url.JoinPath(c.url, deploymentEndpoint, id)
	if err != nil {
		return fmt.Errorf("failed to join delete deployment endpoint: %w", err)
	}

	if _, err := c.DoSignedRequest(ctx, http.MethodDelete, u, nil); err != nil {
		return fmt.Errorf("failed to delete deployment: %w", err)
	}

	return nil
}

func (c *NodeshiftClient) ListRegions(ctx context.Context) ([]string, error) {
	u, err := url.JoinPath(c.url, regionsEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to join list regions endpoint: %w", err)
	}

	b, err := c.DoSignedRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list regions: %w", err)
	}

	regions := make([]string, 0)
	if err := json.Unmarshal(b, &regions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal list regions response: %w", err)
	}

	return regions, nil
}
