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
	"github.com/stretchr/testify/require"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/stretchr/testify/assert"
)

func TestNewSigner_WithStaticCredentials(t *testing.T) {
	t.Parallel()

	accessKey := "ACCESS_KEY"
	secretKey := "SECRET_KEY"

	signer := NewSigner(WithStaticCredentials(accessKey, secretKey))

	creds, _ := signer.v4.Credentials.Get()

	assert.NotNil(t, signer)
	assert.Equal(t, accessKey, creds.AccessKeyID)
	assert.Equal(t, secretKey, creds.SecretAccessKey)
}

func TestNewSigner_WithSharedCredentials(t *testing.T) {
	t.Parallel()

	filename := "credentials"
	profile := "default"

	signer := NewSigner(WithSharedCredentials(filename, profile))

	creds, _ := signer.v4.Credentials.Get()

	assert.NotNil(t, signer)
	assert.Equal(t, credentials.SharedCredsProviderName, creds.ProviderName)
}

func TestSigner_SignRequest(t *testing.T) {
	t.Parallel()

	accessKey := "ACCESS_KEY"
	secretKey := "SECRET_KEY"

	signer := NewSigner(WithStaticCredentials(accessKey, secretKey), WithDebugLogger(t))

	// Create a sample HTTP request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://localhost:6005", nil)
	require.NoError(t, err)

	err = signer.SignRequest(req, nil)
	require.NoError(t, err)

	authHeader := req.Header.Get("Authorization")
	assert.NotEmpty(t, authHeader)

	authHeaderValues := strings.Split(strings.ReplaceAll(authHeader, ",", ""), " ")

	assert.Equal(t, "AWS4-HMAC-SHA256", authHeaderValues[0])
	assert.Equal(t, "SignedHeaders=host;x-amz-date", authHeaderValues[2])
}

func TestDebugLogger_Log(t *testing.T) {
	t.Parallel()

	type args struct {
		values []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "debug logger",
			args: args{
				values: []any{"test"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l := &DebugLogger{
				Context: context.Background(),
			}
			l.Log(tt.args.values...)
		})
	}
}

func TestSigner_SignRequest1(t *testing.T) {
	t.Parallel()

	type args struct {
		req     *http.Request
		body    io.ReadSeeker
		wantErr bool
	}
	newURL, err := url.Parse(exampleURLString)
	require.NoError(t, err)

	rq := &http.Request{
		Header: make(http.Header, 0),
		URL:    newURL,
	}
	ctx, cls := context.WithCancel(context.Background())
	rq = rq.WithContext(ctx)
	cls()
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sign request",
			args: args{
				req: &http.Request{
					URL:    newURL,
					Header: make(http.Header, 0),
				},
				body: bytes.NewReader([]byte{}),
			},
		},
		{
			name: "sign request error",
			args: args{
				req:     rq,
				wantErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Signer{
				v4: NewSigner(WithStaticCredentials("access", "secret")).v4,
			}
			err := s.SignRequest(tt.args.req, tt.args.body)
			if !tt.args.wantErr {
				assert.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestSigner_signRequest(t *testing.T) {
	t.Parallel()

	type args struct {
		req     *http.Request
		body    io.ReadSeeker
		wantErr bool
	}

	newURL, err := url.Parse(exampleURLString)
	require.NoError(t, err)

	rq := &http.Request{
		Header: make(http.Header, 0),
		URL:    newURL,
	}
	ctx, cls := context.WithCancel(context.Background())
	rq = rq.WithContext(ctx)
	cls()
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sign request",
			args: args{
				req: &http.Request{
					URL:    newURL,
					Header: make(http.Header, 0),
				},
				body: bytes.NewReader([]byte{}),
			},
		},
		{
			name: "sign request error",
			args: args{
				req:     rq,
				body:    bytes.NewReader([]byte{}),
				wantErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Signer{
				v4: NewSigner(WithStaticCredentials("access", "secret")).v4,
			}
			err := s.signRequest(tt.args.req, tt.args.body)
			if !tt.args.wantErr {
				assert.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestWithAnonymousCredentials(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			assert.Equalf(t, tt.want, WithAnonymousCredentials()(), "WithAnonymousCredentials()")
		})
	}
}

func TestWithDebugLogger(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			d := NewSigner(WithStaticCredentials("access", "secret"))
			WithDebugLogger(tt.args.logger)(d)

			assert.NotNil(t, d.v4.Logger)
		})
	}
}

func TestWithSharedCredentials(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			c := WithSharedCredentials(tt.args.filename, tt.args.profile)()
			assert.Equal(t, tt.want, c)
		})
	}
}

func TestWithStaticCredentials(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			assert.Equal(t, tt.want, WithStaticCredentials(tt.args.accessKey, tt.args.secretKey)())
		})
	}
}
