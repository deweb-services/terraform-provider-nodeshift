package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_DeploymentCreate(t *testing.T) {
	mustPollTimes := 10

	expectedResponse := CreatedDeployment{
		StartTime:   time.Now().Unix(),
		ServiceType: "Backend Service",
		Data: &CreatedDeploymentData{
			IP:   "190.12.32.19",
			IPv6: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			Ygg:  "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			Plan: CreatedDeploymentDataPlan{
				ID:      1,
				CPU:     4,
				RAM:     2,
				Hdd:     50,
				HddType: "ssd",
			},
		},
		IsError: false,
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/api/terraform/deployment"):
			t.Log("received create deployment request")
			response := DeploymentCreateTask{
				ID:     "ea50370f-ae1f-4a7b-8626-1d389020922e",
				TaskID: "67e1c297-0d78-4668-b23a-6a268000a392",
			}
			b, _ := json.Marshal(response)
			w.Write(b)
			return
		case strings.HasPrefix(r.URL.Path, "/api/task"):
			t.Log("received poll deployment task request")
			mustPollTimes--
			if mustPollTimes == 0 {
				expectedResponse.EndTime = Int64P(time.Now().Unix())
				b, _ := json.Marshal(expectedResponse)
				w.Write(b)
				w.WriteHeader(http.StatusOK)
				return
			}
			response := CreatedDeployment{
				StartTime:   time.Now().Unix(),
				ServiceType: "Backend Service",
				EndTime:     nil,
				Data:        nil,
				IsError:     false,
			}
			b, _ := json.Marshal(response)
			w.Write(b)
		}
	}))
	defer mockServer.Close()

	client := NewClient(DWSProviderConfiguration{
		AccessKey:       "access_key",
		SecretAccessKey: "secret_access_key",
	}, ClientOptWithURL(mockServer.URL))

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
		b, _ := json.Marshal(expectedResponse)
		w.Write(b)
	}))
	defer mockServer.Close()

	client := NewClient(DWSProviderConfiguration{
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
