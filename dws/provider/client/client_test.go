package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

const (
	exampleUrlString    = "https://example.com/"
	exampleErrUrlString = "https://example.moc/"
)

var defaultSigner = NewSigner(WithStaticCredentials("access", "secret"))

func makeTwoClients() (*DWSClient, *DWSClient) {
	cli1 := NewClient(context.TODO(), DWSProviderConfiguration{})
	cli2 := &DWSClient{
		Config:          DWSProviderConfiguration{},
		transactionNote: cli1.transactionNote,
		client:          cli1.client,
		signer:          cli1.signer,
		url:             cli1.url,
		s3client:        cli1.s3client,
	}
	return cli1, cli2
}

var exampleComContent = `<!doctype html>
<html>
<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
        
    }
    div {
        width: 600px;
        margin: 5em auto;
        padding: 2em;
        background-color: #fdfdff;
        border-radius: 0.5em;
        box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);
    }
    a:link, a:visited {
        color: #38488f;
        text-decoration: none;
    }
    @media (max-width: 700px) {
        div {
            margin: 0 auto;
            width: auto;
        }
    }
    </style>    
</head>

<body>
<div>
    <h1>Example Domain</h1>
    <p>This domain is for use in illustrative examples in documents. You may use this
    domain in literature without prior coordination or asking for permission.</p>
    <p><a href="https://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>
`

func TestClientOptWithS3(t *testing.T) {
	tests := []struct {
		name string
		want ClientOpt
	}{
		{
			name: "client_opt_with_s3",
			want: ClientOpt(func(c *DWSClient) {}),
		},
		{
			name: "client_opt_with_s3_err",
			want: ClientOpt(func(c *DWSClient) {
				c.s3client = nil
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli1, cli2 := makeTwoClients()
			ClientOptWithS3()(cli1)
			tt.want(cli2)
			assert.Equalf(t, cli1, cli2, "ClientOptWithS3()")
		})
	}
}

func TestClientOptWithURL(t *testing.T) {
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
				url: "https://app.dws.sh",
			},
			want: ClientOpt(func(c *DWSClient) {
				c.url = "https://app.dws.sh"
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli1, cli2 := makeTwoClients()
			ClientOptWithURL(tt.args.url)(cli1)
			tt.want(cli2)
			assert.Equalf(t, cli1, cli2, tt.args.url)
		})
	}
}

func TestDWSClient_DoRequest(t *testing.T) {
	type fields struct {
		Config DWSProviderConfiguration
		client *http.Client
		signer *Signer
	}
	type args struct {
		req *http.Request
	}
	newUrl, _ := url.Parse(exampleUrlString)
	newErrUrl, _ := url.Parse(exampleErrUrlString)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr error
	}{
		{
			name: "do_request",
			fields: fields{
				Config: DWSProviderConfiguration{},
				client: &http.Client{},
				signer: &Signer{},
			},
			args: args{
				req: &http.Request{
					URL: newUrl,
				},
			},
			want:    []byte(exampleComContent),
			wantErr: nil,
		},
		{
			name: "do_request_err",
			fields: fields{
				Config: DWSProviderConfiguration{},
				client: &http.Client{},
				signer: &Signer{},
			},
			args: args{
				req: &http.Request{
					URL: newErrUrl,
				},
			},
			want:    nil,
			wantErr: fmt.Errorf("error making request: Get \"%s\": dial tcp: lookup example.moc: no such host", exampleErrUrlString),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DWSClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: tt.fields.signer,
			}
			got, err := c.DoRequest(tt.args.req)
			if tt.wantErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equalf(t, tt.want, got, "DoRequest(%v)", tt.args.req)
		})
	}
}

func TestDWSClient_DoSignedRequest(t *testing.T) {
	type fields struct {
		Config DWSProviderConfiguration
		client *http.Client
		signer *Signer
		url    string
	}
	type args struct {
		ctx      context.Context
		method   string
		endpoint string
		body     io.ReadSeeker
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr error
	}{
		{
			name: "do_signed_request",
			fields: fields{
				Config: DWSProviderConfiguration{},
				client: &http.Client{},
				signer: defaultSigner,
				url:    exampleUrlString,
			},
			args: args{
				ctx:      context.TODO(),
				method:   "GET",
				endpoint: exampleUrlString,
				body:     bytes.NewReader([]byte{}),
			},
			want: []byte(exampleComContent),
		},
		{
			name: "do_signed_request_err",
			fields: fields{
				Config: DWSProviderConfiguration{},
				client: &http.Client{},
				signer: defaultSigner,
			},
			args:    args{},
			want:    nil,
			wantErr: fmt.Errorf(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DWSClient{
				Config: tt.fields.Config,
				client: tt.fields.client,
				signer: tt.fields.signer,
				url:    tt.fields.url,
			}
			got, err := c.DoSignedRequest(tt.args.ctx, tt.args.method, tt.args.endpoint, tt.args.body)
			if tt.wantErr != nil {
				assert.NotNil(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equalf(t, tt.want, got, "DoSignedRequest(%v, %v, %v, %v)", tt.args.ctx, tt.args.method, tt.args.endpoint, tt.args.body)
		})
	}
}

func TestDWSClient_SetGlobalTransactionNote(t *testing.T) {
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
			c := &DWSClient{
				transactionNote: tt.fields.transactionNote,
			}
			c.SetGlobalTransactionNote(tt.args.note)
			assert.Equal(t, tt.args.note, c.transactionNote)
		})
	}
}

func TestDWSClient_newAwsClient(t *testing.T) {
	type fields struct {
		Config          DWSProviderConfiguration
		transactionNote string
		client          *http.Client
		signer          *Signer
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
				Config: DWSProviderConfiguration{
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
			c := &DWSClient{
				Config:          tt.fields.Config,
				transactionNote: tt.fields.transactionNote,
				client:          tt.fields.client,
				signer:          tt.fields.signer,
				url:             tt.fields.url,
				s3client:        tt.fields.s3client,
			}
			err := c.newAwsClient()
			assert.NoError(t, err)
			assert.NotNil(t, c.s3client)
		})
	}
}

func TestDWSProviderConfiguration_FromSlice(t *testing.T) {
	type fields struct {
		Timeout               time.Duration
		AccessKey             string
		SecretAccessKey       string
		SharedCredentialsFile string
		Profile               string
		S3Endpoint            string
		S3Region              string
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
			},
			args: args{
				values: []string{"a_key", "s_key", "s_file", "profile", "s3_endpoint", "s3_region"},
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
			},
			args: args{
				values: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &DWSProviderConfiguration{}
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
	type args struct {
		ctx           context.Context
		configuration DWSProviderConfiguration
		opts          []ClientOpt
	}
	tests := []struct {
		name string
		args args
		want *DWSClient
	}{
		{
			name: "new_client",
			args: args{
				ctx:           context.TODO(),
				configuration: DWSProviderConfiguration{},
				opts:          nil,
			},
		},
		{
			name: "new_client_with_configuration",
			args: args{
				ctx: context.TODO(),
				configuration: DWSProviderConfiguration{
					SharedCredentialsFile: "s_file",
					Profile:               "profile",
				},
				opts: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dwsCli := NewClient(tt.args.ctx, tt.args.configuration, tt.args.opts...)
			assert.NotNil(t, dwsCli)
		})
	}
}

func Test_checkResponse(t *testing.T) {
	type args struct {
		res *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "check_response",
			args: args{
				res: &http.Response{
					StatusCode: 200,
				},
			},
			wantErr: nil,
		},
		{
			name: "check_response_fail",
			args: args{
				res: &http.Response{
					StatusCode: 400,
				},
			},
			wantErr: fmt.Errorf("request failed, status code: %d", 400),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, checkResponse(tt.args.res), tt.wantErr)
		})
	}
}
