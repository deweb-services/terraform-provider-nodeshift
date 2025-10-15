package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	DefaultAPIEndpoint = "https://app.nodeshift.com"

	defaultTimeoutSeconds = 120

	deploymentEndpoint = "/api/terraform/deployment"

	VPCEndpoint = "/api/terraform/vpc"

	GPUEndpoint = "/api/terraform/gpu"

	LBEndpoint = "/api/terraform/load-balancer"

	regionsEndpoint = "/api/terraform/deployment/all-countries"

	providerParametersCount = 6
	providerOptionsCount    = 4
)

type NodeshiftClient struct {
	Config          NodeshiftProviderConfiguration
	transactionNote string
	client          *http.Client
	signer          *Signer
	url             string
	s3client        *s3.Client
}

type NodeshiftProviderConfiguration struct {
	Timeout               time.Duration
	AccessKey             string
	SecretAccessKey       string
	SharedCredentialsFile string
	Profile               string
	S3Endpoint            string
	S3Region              string
	APIEndpoint           string
	Insecure              bool
}

type TaskResponse struct {
	StartTime   int64  `json:"startTime"`
	ServiceType string `json:"serviceType"`
	EndTime     any    `json:"endTime"`
	IsError     bool   `json:"isError"`
	Data        any    `json:"data"`
}

type ClientOpt func(c *NodeshiftClient)

func (dc *NodeshiftProviderConfiguration) FromSlice(values []string) {
	if len(values) < providerParametersCount {
		return
	}

	dc.AccessKey = values[0]
	dc.SecretAccessKey = values[1]
	dc.SharedCredentialsFile = values[2]
	dc.Profile = values[3]
	dc.S3Endpoint = values[4]
	dc.S3Region = values[5]
	dc.APIEndpoint = values[6]
}

func (c *NodeshiftClient) SetGlobalTransactionNote(note string) {
	c.transactionNote = note
}

func NewClient(ctx context.Context, configuration NodeshiftProviderConfiguration, opts ...ClientOpt) *NodeshiftClient {
	signerOpts := make([]CredentialsOpt, 0, providerOptionsCount)

	if configuration.AccessKey != "" && configuration.SecretAccessKey != "" {
		signerOpts = append(signerOpts, WithStaticCredentials(configuration.AccessKey, configuration.SecretAccessKey))
	}

	if configuration.SharedCredentialsFile != "" && configuration.Profile != "" {
		signerOpts = append(signerOpts, WithSharedCredentials(configuration.SharedCredentialsFile, configuration.Profile))
	}

	if len(signerOpts) == 0 {
		signerOpts = append(signerOpts, WithAnonymousCredentials())
	}

	signer := NewSigner(signerOpts[len(signerOpts)-1], WithDebugLogger(&DebugLogger{ctx}))

	c := &NodeshiftClient{
		Config: configuration,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
		},
		signer: signer,
		url:    DefaultAPIEndpoint,
	}

	if configuration.Timeout == 0 {
		c.client.Timeout = defaultTimeoutSeconds * time.Second
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *NodeshiftClient) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	err = checkResponse(res)
	if err != nil {
		return nil, fmt.Errorf("external API returned an error code: %w, response body: %s", err, string(b))
	}

	return b, nil
}

func (c *NodeshiftClient) DoSignedRequest(ctx context.Context, method string, endpoint string, body io.ReadSeeker) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	tflog.Debug(ctx, "signing request")
	if err = c.signer.SignRequest(req, body); err != nil {
		return nil, err
	}

	return c.doRequest(req)
}

func checkResponse(res *http.Response) error {
	if res.StatusCode >= 400 && res.StatusCode <= 599 {
		return fmt.Errorf("%w: status code %d", errRequestFailed, res.StatusCode)
	}

	return nil
}

func ClientOptWithURL(url string) ClientOpt {
	return func(c *NodeshiftClient) {
		c.url = url
	}
}

func ClientOptWithS3() ClientOpt {
	return func(c *NodeshiftClient) {
		if c.Config.S3Endpoint == "" {
			return
		}
		if err := c.newAwsClient(); err != nil {
			tflog.Error(context.Background(), err.Error())
		}
	}
}

func ClientOptWithInsecure() ClientOpt {
	return func(c *NodeshiftClient) {
		if c.client == nil {
			return
		}

		c.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				// nolint: gosec
				InsecureSkipVerify: true,
			},
		}
	}
}

func (c *NodeshiftClient) newAwsClient() error {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return fmt.Errorf("s3 load deafault config error: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = c.Config.S3Region
		o.HTTPClient = c.client
		o.Credentials = credentials.NewStaticCredentialsProvider(c.Config.AccessKey,
			c.Config.SecretAccessKey, "")
		o.BaseEndpoint = aws.String(c.Config.S3Endpoint)
		o.UsePathStyle = true
	})
	c.s3client = client

	return nil
}
