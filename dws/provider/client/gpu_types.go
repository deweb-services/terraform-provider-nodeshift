package client

type GPUConfig struct {
	GPUName  string `json:"gpuName"`
	Image    string `json:"image"`
	SSHKey   string `json:"sshKey"`
	GPUCount int64  `json:"gpuCount,omitempty"`
	Region   string `json:"region"`
}

type GPUConfigResponse struct {
	Region   string `json:"region,omitempty"`
	Image    string `json:"image"`
	GPUName  string `json:"gpuName"`
	GPUCount int64  `json:"gpuCount"`
	UUID     string `json:"uuid"`
}

type RentedGpuInfoResponse struct {
	ActualStatus string `json:"actual_status,omitempty"`
	GpuName      string `json:"gpu_name"`
	NumGpus      int64  `json:"num_gpus"`
	SshHost      string `json:"ssh_host"`
	SshPort      int64  `json:"ssh_port"`
}
