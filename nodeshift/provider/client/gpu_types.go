package client

type CreateGPURequest struct {
	GPUName        string `json:"gpuName"`
	Image          string `json:"image"`
	SSHKey         string `json:"sshKey"`
	GPUCount       int64  `json:"gpuCount,omitempty"`
	Region         string `json:"region,omitempty"`
	Disk           int64  `json:"disk,omitempty"`
	MinCudaVersion string `json:"minCudaVersion,omitempty"`
}

type CreateGPUResponse struct {
	UUID     string `json:"uuid"`
	Region   string `json:"region,omitempty"`
	Image    string `json:"image"`
	GPUName  string `json:"gpuName"`
	GPUCount int64  `json:"gpuCount"`
}

type GetGPUResponse struct {
	UUID    string `json:"uuid"`
	GpuName string `json:"gpuName"`
	NumGpus int64  `json:"gpusAmount"`
	SSHHost string `json:"sshHost"`
	SSHPort int64  `json:"sshPort"`
	Status  string `json:"status"`
}
