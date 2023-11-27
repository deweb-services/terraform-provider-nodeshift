package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewSigner_WithStaticCredentials(t *testing.T) {
	accessKey := "ACCESS_KEY"
	secretKey := "SECRET_KEY"

	signer := NewSigner(WithStaticCredentials(accessKey, secretKey))

	creds, _ := signer.v4.Credentials.Get()

	assert.NotNil(t, signer)
	assert.Equal(t, accessKey, creds.AccessKeyID)
	assert.Equal(t, secretKey, creds.SecretAccessKey)
}

func TestNewSigner_WithSharedCredentials(t *testing.T) {
	filename := "credentials"
	profile := "default"

	signer := NewSigner(WithSharedCredentials(filename, profile))

	creds, _ := signer.v4.Credentials.Get()

	assert.NotNil(t, signer)
	assert.Equal(t, credentials.SharedCredsProviderName, creds.ProviderName)
}

func TestSigner_SignRequest(t *testing.T) {
	accessKey := "ACCESS_KEY"
	secretKey := "SECRET_KEY"

	signer := NewSigner(WithStaticCredentials(accessKey, secretKey), WithDebugLogger(t))

	// Create a sample HTTP request
	req, err := http.NewRequest(http.MethodGet, "https://localhost:6005", nil)
	assert.NoError(t, err)

	err = signer.SignRequest(req, nil)
	assert.NoError(t, err)

	for header, value := range req.Header {
		t.Logf("header: %s, value: %s", header, value)
	}

	assert.NoError(t, err)

	authHeader := req.Header.Get("Authorization")
	assert.NotEmpty(t, authHeader)

	authHeaderValues := strings.Split(strings.Replace(authHeader, ",", "", -1), " ")

	assert.Equal(t, "AWS4-HMAC-SHA256", authHeaderValues[0])
	assert.Equal(t, "SignedHeaders=host;x-amz-date", authHeaderValues[2])
}

func TestDebugLogger_Log(t *testing.T) {
	type fields struct {
		Context context.Context
	}
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "debug logger",
			fields: fields{
				Context: context.TODO(),
			},
			args: args{
				values: []any{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DebugLogger{
				Context: tt.fields.Context,
			}
			l.Log(tt.args.values...)
		})
	}
}

func TestNewSigner(t *testing.T) {
	type args struct {
		credentialsOpt CredentialsOpt
		opts           []SignerOpt
	}
	opt := WithStaticCredentials("", "")
	signer := &v4.Signer{
		Credentials: opt(),
	}
	tests := []struct {
		name string
		args args
		want *Signer
	}{
		{
			name: "new signer",
			args: args{
				credentialsOpt: opt,
				opts:           []SignerOpt{},
			},
			want: &Signer{
				v4: signer,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewSigner(tt.args.credentialsOpt, tt.args.opts...), "NewSigner(%v, %v)", tt.args.credentialsOpt)
		})
	}
}

func TestSigner_SignRequest1(t *testing.T) {
	type fields struct {
		v4 *v4.Signer
	}
	type args struct {
		req     *http.Request
		body    io.ReadSeeker
		wantErr bool
	}
	newUrl, _ := url.Parse(exampleUrlString)
	rq := &http.Request{
		Header: make(http.Header, 0),
		URL:    newUrl,
	}
	ctx, cls := context.WithCancel(context.TODO())
	rq = rq.WithContext(ctx)
	cls()
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "sign request",
			fields: fields{
				v4: defaultSigner.v4,
			},
			args: args{
				req: &http.Request{
					URL:    newUrl,
					Header: make(http.Header, 0),
				},
				body: bytes.NewReader([]byte{}),
			},
		},
		{
			name: "sign request error",
			fields: fields{
				v4: defaultSigner.v4,
			},
			args: args{
				req:     rq,
				wantErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Signer{
				v4: tt.fields.v4,
			}
			err := s.SignRequest(tt.args.req, tt.args.body)
			if !tt.args.wantErr {
				assert.NoError(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestSigner_signRequest(t *testing.T) {
	type fields struct {
		v4 *v4.Signer
	}
	type args struct {
		req     *http.Request
		body    io.ReadSeeker
		wantErr bool
	}
	newUrl, _ := url.Parse(exampleUrlString)
	rq := &http.Request{
		Header: make(http.Header, 0),
		URL:    newUrl,
	}
	ctx, cls := context.WithCancel(context.TODO())
	rq = rq.WithContext(ctx)
	cls()
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "sign request",
			fields: fields{
				v4: defaultSigner.v4,
			},
			args: args{
				req: &http.Request{
					URL:    newUrl,
					Header: make(http.Header, 0),
				},
				body: bytes.NewReader([]byte{}),
			},
		},
		{
			name: "sign request error",
			fields: fields{
				v4: defaultSigner.v4,
			},
			args: args{
				req:     rq,
				body:    bytes.NewReader([]byte{}),
				wantErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Signer{
				v4: tt.fields.v4,
			}
			err := s.signRequest(tt.args.req, tt.args.body)
			if !tt.args.wantErr {
				assert.NoError(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestWithAnonymousCredentials(t *testing.T) {
	tests := []struct {
		name string
		want *credentials.Credentials
	}{
		{
			name: "anonymous credentials",
			want: credentials.AnonymousCredentials,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, WithAnonymousCredentials()(), "WithAnonymousCredentials()")
		})
	}
}

func TestWithDebugLogger(t *testing.T) {
	type args struct {
		logger aws.Logger
	}
	tests := []struct {
		name string
		args args
		want SignerOpt
	}{
		{
			name: "debug logger",
			args: args{
				logger: &DebugLogger{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := defaultSigner
			WithDebugLogger(tt.args.logger)(d)

			assert.NotNil(t, d.v4.Logger)
		})
	}
}

func TestWithSharedCredentials(t *testing.T) {
	type args struct {
		filename string
		profile  string
	}
	tests := []struct {
		name string
		args args
		want *credentials.Credentials
	}{
		{
			name: "shared credentials",
			args: args{
				filename: "filename-1",
				profile:  "profile-1",
			},
			want: credentials.NewSharedCredentials("filename-1", "profile-1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := WithSharedCredentials(tt.args.filename, tt.args.profile)()
			assert.Equal(t, tt.want, c)
		})
	}
}

func TestWithStaticCredentials(t *testing.T) {
	type args struct {
		accessKey string
		secretKey string
	}
	tests := []struct {
		name string
		args args
		want *credentials.Credentials
	}{
		{
			name: "static credentials",
			args: args{
				accessKey: "access",
				secretKey: "secret",
			},
			want: credentials.NewStaticCredentials("access", "secret", ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, WithStaticCredentials(tt.args.accessKey, tt.args.secretKey)())
		})
	}
}
