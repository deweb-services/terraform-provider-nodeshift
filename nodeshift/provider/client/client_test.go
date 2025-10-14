package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	exampleURLString    = "https://example.com/"
	exampleErrURLString = "https://example.moc/"
)

func makeTwoClients() (*NodeshiftClient, *NodeshiftClient) {
	cli1 := NewClient(context.Background(), NodeshiftProviderConfiguration{})
	cli2 := &NodeshiftClient{
		Config:          NodeshiftProviderConfiguration{},
		transactionNote: cli1.transactionNote,
		client:          cli1.client,
		signer:          cli1.signer,
		url:             cli1.url,
		s3client:        cli1.s3client,
	}

	return cli1, cli2
}

func TestClientOptWithS3(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want ClientOpt
	}{
		{
			name: "client_opt_with_s3",
			want: ClientOpt(func(c *NodeshiftClient) {}),
		},
		{
			name: "client_opt_with_s3_err",
			want: ClientOpt(func(c *NodeshiftClient) {
				c.s3client = nil
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cli1, cli2 := makeTwoClients()
			ClientOptWithS3()(cli1)
			tt.want(cli2)
			assert.Equal(t, cli1, cli2)
		})
	}
}

func TestClientOptWithURL(t *testing.T) {
	t.Parallel()

	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want ClientOpt
	}{
		{
			name: "client_opt_with_url",
			args: args{
				url: "https://app.nodeshift.sh",
			},
			want: ClientOpt(func(c *NodeshiftClient) {
				c.url = "https://app.nodeshift.sh"
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cli1, cli2 := makeTwoClients()
			ClientOptWithURL(tt.args.url)(cli1)
			tt.want(cli2)
			assert.Equalf(t, cli1, cli2, tt.args.url)
		})
	}
}

func TestNodeshiftClient_DoRequest(t *testing.T) {
	t.Parallel()

	type args struct {
		req *http.Request
	}
	newURL, err := url.Parse(exampleURLString)
	require.NoError(t, err)

	newErrURL, err := url.Parse(exampleErrURLString)
	require.NoError(t, err)

	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			name: "do_request",
			args: args{
				req: &http.Request{
					URL: newURL,
				},
			},
		},
		{
			name: "do_request_err",
			args: args{
				req: &http.Request{
					URL: newErrURL,
				},
			},
			wantErr: "error making request: Get \"https://example.moc/\": dial tcp: lookup example.moc: no such host",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				signer: &Signer{},
			}
			got, err := c.doRequest(tt.args.req)
			if tt.wantErr != "" {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.NotEmpty(t, got)
		})
	}
}

func TestNodeshiftClient_DoSignedRequest(t *testing.T) {
	t.Parallel()

	type fields struct {
		url string
	}
	type args struct {
		method   string
		endpoint string
		body     io.ReadSeeker
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "do_signed_request",
			fields: fields{
				url: exampleURLString,
			},
			args: args{
				method:   "GET",
				endpoint: exampleURLString,
				body:     bytes.NewReader([]byte{}),
			},
		},
		{
			name:    "do_signed_request_err",
			fields:  fields{},
			args:    args{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config: NodeshiftProviderConfiguration{},
				client: &http.Client{},
				signer: NewSigner(WithStaticCredentials("access", "secret")),
				url:    tt.fields.url,
			}
			got, err := c.DoSignedRequest(context.Background(), tt.args.method, tt.args.endpoint, tt.args.body)
			if tt.wantErr != nil {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.NotEmpty(t, got)
		})
	}
}

func TestNodeshiftClient_SetGlobalTransactionNote(t *testing.T) {
	t.Parallel()

	type fields struct {
		transactionNote string
	}
	type args struct {
		note string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "set_global_transaction_note",
			fields: fields{
				transactionNote: "",
			},
			args: args{
				note: "note",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				transactionNote: tt.fields.transactionNote,
			}
			c.SetGlobalTransactionNote(tt.args.note)
			assert.Equal(t, tt.args.note, c.transactionNote)
		})
	}
}

func TestNodeshiftClient_newAwsClient(t *testing.T) {
	t.Parallel()

	type fields struct {
		Config          NodeshiftProviderConfiguration
		transactionNote string
		client          *http.Client
		url             string
		s3client        *s3.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "new_aws_client",
			fields: fields{
				Config: NodeshiftProviderConfiguration{
					AccessKey:       "a_key",
					SecretAccessKey: "s_key",
					S3Endpoint:      "https://s3.ep.com",
					S3Region:        "us-west",
				},
				s3client: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &NodeshiftClient{
				Config:          tt.fields.Config,
				transactionNote: tt.fields.transactionNote,
				client:          tt.fields.client,
				signer:          &Signer{},
				url:             tt.fields.url,
				s3client:        tt.fields.s3client,
			}
			err := c.newAwsClient()

			require.NoError(t, err)
			assert.NotNil(t, c.s3client)
		})
	}
}

func TestNodeshiftProviderConfiguration_FromSlice(t *testing.T) {
	t.Parallel()

	type fields struct {
		Timeout               time.Duration
		AccessKey             string
		SecretAccessKey       string
		SharedCredentialsFile string
		Profile               string
		S3Endpoint            string
		S3Region              string
		APIEndpoint           string
	}
	type args struct {
		values []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "configuration_from_slice",
			fields: fields{
				AccessKey:             "a_key",
				SecretAccessKey:       "s_key",
				SharedCredentialsFile: "s_file",
				Profile:               "profile",
				S3Endpoint:            "s3_endpoint",
				S3Region:              "s3_region",
				APIEndpoint:           "api_endpoint",
			},
			args: args{
				values: []string{"a_key", "s_key", "s_file", "profile", "s3_endpoint", "s3_region", "api_endpoint"},
			},
		},
		{
			name: "configuration_from_slice - empty",
			fields: fields{
				AccessKey:             "a_key",
				SecretAccessKey:       "s_key",
				SharedCredentialsFile: "s_file",
				Profile:               "profile",
				S3Endpoint:            "s3_endpoint",
				S3Region:              "s3_region",
				APIEndpoint:           "api_url",
			},
			args: args{
				values: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dc := &NodeshiftProviderConfiguration{}
			dc.FromSlice(tt.args.values)
			if len(tt.args.values) != 0 {
				assert.Equal(t, tt.fields.AccessKey, tt.args.values[0])
				assert.Equal(t, tt.fields.SecretAccessKey, tt.args.values[1])
				assert.Equal(t, tt.fields.SharedCredentialsFile, tt.args.values[2])
				assert.Equal(t, tt.fields.Profile, tt.args.values[3])
				assert.Equal(t, tt.fields.S3Endpoint, tt.args.values[4])
				assert.Equal(t, tt.fields.S3Region, tt.args.values[5])
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	t.Parallel()

	type args struct {
		configuration NodeshiftProviderConfiguration
		opts          []ClientOpt
	}
	tests := []struct {
		name string
		args args
		want *NodeshiftClient
	}{
		{
			name: "new_client",
			args: args{
				configuration: NodeshiftProviderConfiguration{},
				opts:          nil,
			},
		},
		{
			name: "new_client_with_configuration",
			args: args{
				configuration: NodeshiftProviderConfiguration{
					SharedCredentialsFile: "s_file",
					Profile:               "profile",
				},
				opts: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			nodeshiftCli := NewClient(context.Background(), tt.args.configuration, tt.args.opts...)
			assert.NotNil(t, nodeshiftCli)
		})
	}
}

func Test_checkResponse(t *testing.T) {
	t.Parallel()

	type args struct {
		res *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "check_response",
			args: args{
				res: &http.Response{
					StatusCode: http.StatusOK,
				},
			},
		},
		{
			name: "check_response_fail",
			args: args{
				res: &http.Response{
					StatusCode: http.StatusBadRequest,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := checkResponse(tt.args.res)
			if tt.wantErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
		})
	}
}
