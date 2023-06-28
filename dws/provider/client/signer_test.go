package client

import (
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws/credentials"
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
	req, err := http.NewRequest("POST", "https://example.com/api", nil)
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
	assert.Equal(t, "Credential=ACCESS_KEY/20230628/global/terraform/aws4_request", authHeaderValues[1])
	assert.Equal(t, "SignedHeaders=host;x-amz-date", authHeaderValues[2])
}
