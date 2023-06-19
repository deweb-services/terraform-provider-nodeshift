package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type DWSClient struct {
	Config          DWSProviderConfiguration
	transactionNote string
	client          *http.Client
	signer          *Signer
}

type DWSProviderConfiguration struct {
	Timeout               time.Duration
	AccessKey             string
	SecretAccessKey       string
	SharedCredentialsFile string
	Profile               string
}

type TaskResponse struct {
	StartTime   int64  `json:"startTime"`
	ServiceType string `json:"serviceType"`
	EndTime     any    `json:"endTime"`
	IsError     bool   `json:"isError"`
	Data        any    `json:"data"`
}

func (dc *DWSProviderConfiguration) FromSlice(values []string) {
	if len(values) < 5 {
		return
	}
	dc.AccessKey = values[0]
	dc.SecretAccessKey = values[1]
	dc.SharedCredentialsFile = values[2]
	dc.Profile = values[3]
}

func (c *DWSClient) SetGlobalTransactionNote(note string) {
	c.transactionNote = note
}

func NewClient(configuration DWSProviderConfiguration) *DWSClient {

	signerOpts := []SignerOpt{}

	if configuration.AccessKey != "" && configuration.SecretAccessKey != "" {
		signerOpts = append(signerOpts, WithStaticCredentials(configuration.AccessKey, configuration.SecretAccessKey))
	}

	if configuration.SharedCredentialsFile != "" && configuration.Profile != "" {
		signerOpts = append(signerOpts, WithSharedCredentials(configuration.SharedCredentialsFile, configuration.Profile))
	}

	signer := NewSigner(signerOpts[len(signerOpts)-1])

	return &DWSClient{
		Config: configuration,
		client: &http.Client{
			Timeout: configuration.Timeout,
		},
		signer: signer,
	}
}

func (c *DWSClient) DoRequest(req *http.Request) ([]byte, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	err = checkResponse(res)
	if err != nil {
		return nil, fmt.Errorf("external API returned an error code: %w", err)
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return b, nil
}

func (c *DWSClient) DoSignedRequest(ctx context.Context, method string, endpoint string, body io.ReadSeeker) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, DeploymentEndpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create \"create deployment\" request: %w", err)
	}

	if err = c.signer.SignRequest(req, body); err != nil {
		return nil, err
	}

	return c.DoRequest(req)
}

func checkResponse(res *http.Response) error {
	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		return fmt.Errorf("request failed, status code: %d", res.StatusCode)
	}

	return nil
}
