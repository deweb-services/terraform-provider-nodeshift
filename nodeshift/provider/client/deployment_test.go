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
	"github.com/stretchr/testify/require"
)

// nolint: testifylint
func newServer(t *testing.T) (*httptest.Server, INodeshiftClient) {
	t.Helper()

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/api/terraform/deployment/all-countries"):
			b, err := json.Marshal([]string{"Singapore", "Germany", "Sweden"})
			require.NoError(t, err)

			_, err = w.Write(b)
			require.NoError(t, err)

			return
		case strings.HasPrefix(r.URL.Path, "/api/terraform/deployment"):
			response := GetDeploymentResponse{
				UUID:         "123e4567-e89b-12d3-a456-426614174000",
				Status:       "running",
				IP:           "192.168.1.10",
				Cru:          4,
				Mru:          8192,
				Sru:          200,
				Hru:          100,
				HddType:      1,
				Provider:     42,
				Hostname:     "test-node.local",
				Ipv6:         1,
				SSHKey:       "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD...",
				SSHKeyName:   "test-key",
				Image:        7,
				ImageVersion: "v1.2.3",
				ChosenPlanID: 101,
				Price:        "12.34",
				CreatedAt:    time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
			}
			b, err := json.Marshal(response)
			require.NoError(t, err)

			_, err = w.Write(b)
			require.NoError(t, err)

			return
		case strings.HasPrefix(r.URL.Path, "/api/terraform/vpc"):
			response := GetVPCResponse{
				UUID:        "vpc-123e4567-e89b-12d3-a456-426614174000",
				Name:        "test-vpc",
				Description: "This is a test VPC",
				IPRange:     "10.0.0.0/24",
				State:       "running",
				Resources: []Resources{
					{
						IP:     "10.0.0.2",
						Status: "running",
					},
					{
						IP:     "10.0.0.3",
						Status: "running",
					},
					{
						IP:     "10.0.0.4",
						Status: "running",
					},
				},
			}
			b, err := json.Marshal(response)
			require.NoError(t, err)

			_, err = w.Write(b)
			require.NoError(t, err)

			return
		case strings.HasPrefix(r.URL.Path, "/api/terraform/gpu"):
			response := GetGPUResponse{
				UUID:    "gpu-123e4567-e89b-12d3-a456-426614174000",
				GpuName: "NVIDIA Tesla V100",
				NumGpus: 2,
				SSHHost: "192.168.1.50",
				SSHPort: 22,
				Status:  "running",
			}
			b, err := json.Marshal(response)
			require.NoError(t, err)

			_, err = w.Write(b)
			require.NoError(t, err)

			return
		case strings.HasPrefix(r.URL.Path, "/api/terraform/load-balancer"):
			response := GetLBResponse{
				UUID:           "lb-123e4567-e89b-12d3-a456-426614174000",
				Name:           "test-lb",
				Status:         "running",
				ReplicasAmount: 2,
				CPUAmount:      8,
				PriceInUSD:     "123.45",
				CreatedAt:      time.Date(2025, 10, 7, 12, 0, 0, 0, time.UTC),
				LoadBalancerCfg: LoadBalancerDetails{
					DNSRecord: "lb.example.com",
					Replicas: []ReplicaInfo{
						{Name: "replica-1", IP: "10.0.0.2", Region: "us-east-1", Status: 1},
						{Name: "replica-2", IP: "10.0.0.3", Region: "us-east-1", Status: 1},
					},
					ForwardingRules: []ForwardingRuleV2{
						{
							In:  RuleEndpoint{Protocol: "TCP", Port: 80},
							Out: RuleEndpoint{Protocol: "UPD", Port: 8080},
							CertificateParams: map[string]json.RawMessage{
								"cert": json.RawMessage(`"dummy-cert"`),
							},
						},
					},
					Backends: []Backend{
						{
							Public:  BackendInterface{ID: "be-public-1", IP: "192.168.1.10"},
							Private: PrivateInterface{ID: "be-private-1", IP: "10.0.0.2", VPCName: "test-vpc", VPCBootstrapToken: "token123"},
						},
					},
				},
				VPC: VPCInfo{
					UUID:               "vpc-123e4567-e89b-12d3-a456-426614174000",
					Name:               "test-vpc",
					Description:        "Test VPC for LB",
					AddressRangePrefix: "10.0.0.0/24",
					NetworkSize:        256,
					Token:              "token123",
					Resources: []VPCResource{
						{
							UUID:      "res-1",
							Hostname:  "node-1",
							IP:        "10.0.0.2",
							Status:    1,
							CRU:       4,
							MRU:       8192,
							SRU:       200,
							TaskID:    "task-abc",
							CreatedAt: time.Date(2025, 10, 7, 10, 0, 0, 0, time.UTC),
						},
					},
					Tags: []Tag{
						{TagUUID: "tag-1", Name: "env:test"},
					},
					CreatedAt: time.Date(2025, 10, 7, 9, 0, 0, 0, time.UTC),
				},
				Deployments: []Deployment{
					{
						UUID:         "dep-1",
						Status:       1,
						IP:           "10.0.0.2",
						CRU:          4,
						MRU:          8192,
						SRU:          200,
						HRU:          100,
						Hostname:     "node-1",
						IPv4:         true,
						SSHKey:       "ssh-rsa AAAAB3...",
						ImageVersion: "v1.2.3",
						Price:        "12.34",
						Region:       "us-east-1",
						HDDType:      1,
						ChargeRate:   0.05,
						Tags:         []Tag{{TagUUID: "tag-1", Name: "env:test"}},
						CreatedAt:    time.Date(2025, 10, 7, 10, 0, 0, 0, time.UTC),
						AuthType:     1,
						Password:     "pass123",
					},
				},
				Tags: []Tag{
					{TagUUID: "tag-1", Name: "env:test"},
				},
			}
			b, err := json.Marshal(response)
			require.NoError(t, err)

			_, err = w.Write(b)
			require.NoError(t, err)

			return
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))

	return mockServer, NewClient(context.Background(), NodeshiftProviderConfiguration{
		AccessKey:       "access_key",
		SecretAccessKey: "secret_access_key",
	}, ClientOptWithURL(mockServer.URL))
}

func Test_DeploymentCreate(t *testing.T) {
	t.Parallel()

	mockServer, client := newServer(t)
	defer mockServer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	actual, err := client.CreateDeployment(ctx, &CreateDeploymentRequest{
		ImageVersion: "v1.2.3",
		Region:       "default-region",
		CPU:          4,
		RAM:          8192,
		Hdd:          200,
		HddType:      "1",
		Ipv4:         true,
		Ipv6:         true,
		Ygg:          false,
		SSHKey:       "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD...",
		SSHKeyName:   "test-key",
		HostName:     "test-node.local",
		NetworkUUID:  "",
	})

	require.NotEmpty(t, actual)
	require.NoError(t, err)
}

func Test_VPCCreate(t *testing.T) {
	t.Parallel()

	mockServer, client := newServer(t)
	defer mockServer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := client.CreateVPC(ctx, &CreateVPCRequest{
		Name:        "test-vpc",
		Description: "test-vpc description",
		IPRange:     "10.0.0.0/16",
	})
	require.NotEmpty(t, response)
	require.NoError(t, err)
}

func Test_GPUCreate(t *testing.T) {
	t.Parallel()

	mockServer, client := newServer(t)
	defer mockServer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := client.CreateGPU(ctx, &CreateGPURequest{
		GPUName:  "RTX A4000",
		Image:    "ubuntu:latest",
		SSHKey:   "ssh-rsa ...",
		GPUCount: 2,
		Region:   "Central America",
	})
	require.NotEmpty(t, response)
	require.NoError(t, err)
}

func Test_DeleteDeployment(t *testing.T) {
	t.Parallel()

	mockServer, client := newServer(t)
	defer mockServer.Close()

	err := client.DeleteDeployment(context.Background(), "id")
	require.NoError(t, err)
}

func Test_GetDeployment(t *testing.T) {
	t.Parallel()

	mockServer, client := newServer(t)
	defer mockServer.Close()

	resp, err := client.GetDeployment(context.Background(), "id")
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func Test_UpdateDeployment(t *testing.T) {
	t.Parallel()

	mockServer, client := newServer(t)
	defer mockServer.Close()

	resp, err := client.UpdateDeployment(context.Background(), "id", &CreateDeploymentRequest{})
	require.Empty(t, resp)
	require.Error(t, err)

	assert.ErrorIs(t, err, errNotImplemented)
}

func Test_ListRegions(t *testing.T) {
	t.Parallel()

	mockServer, client := newServer(t)
	defer mockServer.Close()

	response, err := client.ListRegions(context.Background())
	require.NotEmpty(t, response)
	require.NoError(t, err)

	assert.Equal(t, []string{"Singapore", "Germany", "Sweden"}, response)
}
