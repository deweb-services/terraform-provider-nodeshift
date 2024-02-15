package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newServer(t *testing.T, requestType string) (*httptest.Server, *AsyncAPIDeploymentResponse, INodeshiftClient) {
	mustPollTimes := 10
	expectedResponse := &AsyncAPIDeploymentResponse{
		StartTime:   time.Now().Unix(),
		ServiceType: "Backend Service",
		Data: &DeploymentResponseData{
			IP:           "190.12.32.19",
			IPv6:         "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			Ygg:          "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			ProviderPlan: "",
		},
		IsError: false,
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/api/terraform/deployment"):
			t.Logf("received %s deployment request", requestType)
			response := AsyncAPIDeploymentTask{
				ID:     "ea50370f-ae1f-4a7b-8626-1d389020922e",
				TaskID: "67e1c297-0d78-4668-b23a-6a268000a392",
			}
			b, _ := json.Marshal(response)
			//nolint:errcheck
			w.Write(b)
			return
		case strings.HasPrefix(r.URL.Path, "/api/task/terraform"):
			t.Log("received poll deployment task request")
			mustPollTimes--
			if mustPollTimes == 0 {
				expectedResponse.EndTime = Int64P(time.Now().Unix())
				b, _ := json.Marshal(expectedResponse)
				//nolint:errcheck
				w.Write(b)
				return
			}
			response := AsyncAPIDeploymentResponse{
				StartTime:   time.Now().Unix(),
				ServiceType: "Backend Service",
				EndTime:     nil,
				Data:        nil,
				IsError:     false,
			}
			b, _ := json.Marshal(response)
			//nolint:errcheck
			w.Write(b)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))

	client := NewClient(context.TODO(), NodeshiftProviderConfiguration{
		AccessKey:       "access_key",
		SecretAccessKey: "secret_access_key",
	}, ClientOptWithURL(mockServer.URL))
	return mockServer, expectedResponse, client
}

func Test_DeploymentCreate(t *testing.T) {
	mockServer, expectedResponse, client := newServer(t, "create")
	defer mockServer.Close()

	response, err := client.CreateDeployment(context.TODO(), &DeploymentConfig{
		ImageVersion: "Ubuntu-v22.04",
		Region:       "USA",
		CPU:          4,
		RAM:          2,
		Hdd:          50,
		HddType:      "ssd",
		Ipv4:         true,
		Ipv6:         true,
		SSHKey:       "test",
	})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.EndTime, response.EndTime)
	assert.Equal(t, expectedResponse.Data, response.Data)
	assert.Equal(t, expectedResponse.ServiceType, response.ServiceType)
}

func Test_VPCCreate(t *testing.T) {
	expectedResponse := VPCConfig{
		ID:          "4165b26c-1536-427d-bfc5-b1feeca7acd9",
		Name:        "test-vpc",
		Description: "test-vpc description",
		IPRange:     "10.0.0.0/16",
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("header: %+v", r.Header)
		b, _ := json.Marshal(expectedResponse)
		//nolint:errcheck
		w.Write(b)
	}))
	defer mockServer.Close()

	client := NewClient(context.TODO(), NodeshiftProviderConfiguration{
		AccessKey:       "access_key",
		SecretAccessKey: "secret_access_key",
	}, ClientOptWithURL(mockServer.URL))

	response, err := client.CreateVPC(context.TODO(), &VPCConfig{
		Name:        "test-vpc",
		Description: "test-vpc description",
		IPRange:     "10.0.0.0/16",
	})
	assert.NoError(t, err)

	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.Name, response.Name)
	assert.Equal(t, expectedResponse.Description, response.Description)
	assert.Equal(t, expectedResponse.IPRange, response.IPRange)
}

func Int64P(n int64) *int64 {
	return &n
}

func Test_GPUCreate(t *testing.T) {
	expectedResponse := GPUConfigResponse{
		Region:   "Central America",
		Image:    "ubuntu:latest",
		GPUName:  "RTX A4000",
		GPUCount: 2,
		UUID:     uuid.NewString(),
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("header: %+v", r.Header)
		b, _ := json.Marshal(expectedResponse)
		//nolint:errcheck
		w.Write(b)
	}))
	defer mockServer.Close()

	client := NewClient(context.TODO(), NodeshiftProviderConfiguration{
		AccessKey:       "access_key",
		SecretAccessKey: "secret_access_key",
	}, ClientOptWithURL(mockServer.URL))

	response, err := client.CreateGPU(context.TODO(), &GPUConfig{
		GPUName:  "RTX A4000",
		Image:    "ubuntu:latest",
		SSHKey:   "ssh-rsa ...",
		GPUCount: 2,
		Region:   "Central America",
	})
	assert.NoError(t, err)

	assert.Equal(t, expectedResponse.GPUName, response.GPUName)
	assert.Equal(t, expectedResponse.GPUCount, response.GPUCount)
	assert.Equal(t, expectedResponse.Image, response.Image)
	assert.Equal(t, expectedResponse.Region, response.Region)
}

func TestNodeshiftClient_DeleteDeployment(t *testing.T) {
	mockServer, _, client := newServer(t, "delete")
	defer mockServer.Close()

	err := client.DeleteDeployment(context.TODO(), "id")
	assert.NoError(t, err)
}

func TestNodeshiftClient_GetDeployment(t *testing.T) {
	mockServer, _, client := newServer(t, "get")
	defer mockServer.Close()
	_, err := client.GetDeployment(context.TODO(), "id")
	assert.NoError(t, err)
}

func TestNodeshiftClient_UpdateDeployment(t *testing.T) {
	mockServer, _, client := newServer(t, "update")
	defer mockServer.Close()
	resp, err := client.UpdateDeployment(context.TODO(), "id", &DeploymentConfig{})
	assert.Nil(t, resp)
	assert.EqualError(t, err, "update not implemented")
}

func TestNodeshiftClient_PollDeploymentTask(t *testing.T) {
	mockServer, expectedResponse, client := newServer(t, "create")
	defer mockServer.Close()

	response, err := client.PollDeploymentTask(context.TODO(), "taskID")
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.EndTime, response.EndTime)
	assert.Equal(t, expectedResponse.Data, response.Data)
	assert.Equal(t, expectedResponse.ServiceType, response.ServiceType)
}
