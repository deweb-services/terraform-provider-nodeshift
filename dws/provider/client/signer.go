package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

type Signer struct {
	*v4.Signer
}

type SignerOpt func() *credentials.Credentials

func NewSigner(opt SignerOpt) *Signer {
	credentials := opt()

	signer := v4.NewSigner(credentials)
	signer.DisableRequestBodyOverwrite = true
	signer.DisableURIPathEscaping = true
	signer.DisableRequestBodyOverwrite = true

	return &Signer{
		signer,
	}
}

func (s *Signer) SignRequest(req *http.Request, body io.ReadSeeker) error {
	if err := s.signRequest(req, body); err != nil {
		return err
	}

	return s.replaceAuthenticationHeaders(req.Header)
}

func (s *Signer) signRequest(req *http.Request, body io.ReadSeeker) error {
	_, err := s.Sign(req, body, "terraform", "global", time.Now())
	if err != nil {
		return fmt.Errorf("failed to sign request: %w", err)
	}

	return nil
}

func WithStaticCredentials(accessKey, secretKey string) SignerOpt {
	return func() *credentials.Credentials {
		return credentials.NewStaticCredentials(accessKey, secretKey, "")
	}
}

func WithSharedCredentials(filename, profile string) SignerOpt {
	return func() *credentials.Credentials {
		return credentials.NewSharedCredentials(filename, profile)
	}
}

func WithAnonymousCredentials() SignerOpt {
	return func() *credentials.Credentials {
		return credentials.AnonymousCredentials
	}
}

func (s *Signer) replaceAuthenticationHeaders(header http.Header) error {
	value := header.Get("Authorization")
	authHeaderParts := strings.Split(value, ",")
	header.Del("Authorization")
	algorithmCredentials := strings.Split(authHeaderParts[0], " ")
	if len(algorithmCredentials) < 2 {
		return errors.New("invalid Authorization header")
	}
	header.Set("X-Amz-Expires", "900")
	header.Set("X-Amz-Algorithm", algorithmCredentials[0])
	header.Set("X-Amz-Credential", strings.Replace(strings.TrimPrefix(strings.TrimSpace(algorithmCredentials[1]), "Credential="), "//", "/", -1))
	header.Set("X-Amz-SignedHeaders", strings.TrimPrefix(strings.TrimSpace(authHeaderParts[1]), "SignedHeaders="))
	header.Set("X-Amz-Signature", strings.TrimPrefix(strings.TrimSpace(authHeaderParts[2]), "Signature="))

	return nil
}
