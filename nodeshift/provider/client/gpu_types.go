package client

type GPUConfig struct {
	GPUName        string `json:"gpuName"`
	Image          string `json:"image"`
	SSHKey         string `json:"sshKey"`
	GPUCount       int64  `json:"gpuCount,omitempty"`
	Region         string `json:"region,omitempty"`
	Disk           int64  `json:"disk,omitempty"`
	MinCudaVersion string `json:"minCudaVersion,omitempty"`
}

type GPUConfigResponse struct {
	Region   string `json:"region,omitempty"`
	Image    string `json:"image"`
	GPUName  string `json:"gpuName"`
	GPUCount int64  `json:"gpuCount"`
	UUID     string `json:"uuid"`
}

type RentedGpuInfoResponse struct {
	ActualStatus string `json:"status"`
	GpuName      string `json:"gpuName"`
	NumGpus      int64  `json:"gpusAmount"`
	SshHost      string `json:"sshHost"`
	SshPort      int64  `json:"sshPort"`
}
