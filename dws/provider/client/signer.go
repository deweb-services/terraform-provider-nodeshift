package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type Signer struct {
	v4 *v4.Signer
}

type CredentialsOpt func() *credentials.Credentials

type SignerOpt func(signer *Signer)

func NewSigner(credentialsOpt CredentialsOpt, opts ...SignerOpt) *Signer {
	credentials := credentialsOpt()
	signer := &Signer{v4: v4.NewSigner(credentials)}

	for _, opt := range opts {
		opt(signer)
	}

	return signer
}

func (s *Signer) SignRequest(req *http.Request, body io.ReadSeeker) error {
	if err := s.signRequest(req, body); err != nil {
		return err
	}

	return nil
}

func (s *Signer) signRequest(req *http.Request, body io.ReadSeeker) error {
	_, err := s.v4.Sign(req, body, "terraform", "global", time.Now())
	if err != nil {
		return fmt.Errorf("failed to sign request: %w", err)
	}

	return nil
}

func WithStaticCredentials(accessKey, secretKey string) CredentialsOpt {
	return func() *credentials.Credentials {
		return credentials.NewStaticCredentials(accessKey, secretKey, "")
	}
}

func WithSharedCredentials(filename, profile string) CredentialsOpt {
	return func() *credentials.Credentials {
		return credentials.NewSharedCredentials(filename, profile)
	}
}

func WithAnonymousCredentials() CredentialsOpt {
	return func() *credentials.Credentials {
		return credentials.AnonymousCredentials
	}
}

func WithDebugLogger(logger aws.Logger) SignerOpt {
	return func(signer *Signer) {
		signer.v4.Debug = aws.LogDebugWithSigning
		signer.v4.Logger = logger
	}
}

type DebugLogger struct {
	context.Context
}

func (l *DebugLogger) Log(values ...interface{}) {
	for _, item := range values {
		tflog.Debug(l, fmt.Sprintf("%+v", item))
	}
}
