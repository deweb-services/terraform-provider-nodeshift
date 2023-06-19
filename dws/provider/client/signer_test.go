package client

import (
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/stretchr/testify/assert"
)

func TestNewSigner_WithStaticCredentials(t *testing.T) {
	accessKey := "ACCESS_KEY"
	secretKey := "SECRET_KEY"

	signer := NewSigner(WithStaticCredentials(accessKey, secretKey))

	creds, _ := signer.Credentials.Get()

	assert.NotNil(t, signer)
	assert.Equal(t, accessKey, creds.AccessKeyID)
	assert.Equal(t, secretKey, creds.SecretAccessKey)
}

func TestNewSigner_WithSharedCredentials(t *testing.T) {
	filename := "credentials"
	profile := "default"

	signer := NewSigner(WithSharedCredentials(filename, profile))

	creds, _ := signer.Credentials.Get()

	assert.NotNil(t, signer)
	assert.Equal(t, credentials.SharedCredsProviderName, creds.ProviderName)
}

func TestSigner_SignRequest(t *testing.T) {
	accessKey := "ACCESS_KEY"
	secretKey := "SECRET_KEY"

	signer := NewSigner(WithStaticCredentials(accessKey, secretKey))

	// Create a sample HTTP request
	req, err := http.NewRequest("GET", "https://example.com/api", nil)
	assert.NoError(t, err)

	err = signer.SignRequest(req, nil)
	assert.NoError(t, err)

	for header, value := range req.Header {
		t.Logf("header: %s, value: %s", header, value)
	}

	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.NotEmpty(t, req.Header.Get("X-Amz-Date"))
	assert.NotEmpty(t, req.Header.Get("X-Amz-Credential"))
	assert.Equal(t, "AWS4-HMAC-SHA256", req.Header.Get("X-Amz-Algorithm"))
	assert.Equal(t, "host;x-amz-date", req.Header.Get("X-Amz-SignedHeaders"))

	// Assert other necessary headers or values as needed
}

func TestSigner_ReplaceHeaders(t *testing.T) {
	accessKey := "ACCESS_KEY"
	secretKey := "SECRET_KEY"

	signer := NewSigner(WithStaticCredentials(accessKey, secretKey))

	// Create a sample HTTP request
	req, err := http.NewRequest("GET", "https://example.com/api", nil)
	assert.NoError(t, err)

	err = signer.signRequest(req, nil)
	assert.NoError(t, err)

	err = signer.replaceAuthenticationHeaders(req.Header)
	assert.NoError(t, err)

	for header, value := range req.Header {
		t.Logf("header: %s, value: %s", header, value)
	}

	assert.Equal(t, "AWS4-HMAC-SHA256", req.Header.Get("X-Amz-Algorithm"))
}
