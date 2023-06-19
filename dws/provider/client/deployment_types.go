package client

type VMConfig struct {
	ImageVersion string `json:"imageVersion"`
	Region       string `json:"region"`
	CPU          int    `json:"cpu"`
	RAM          int    `json:"ram"`
	Hdd          int    `json:"hdd"`
	HddType      string `json:"hddType"`
	Ipv4         bool   `json:"ipv4,omitempty"`
	Ipv6         bool   `json:"ipv6,omitempty"`
	SSHKey       string `json:"sshKey"`
	HostName     string `json:"hostName"`
	NetworkUUID  string `json:"networkUuid,omitempty"`
}

type VMResponse struct {
	StartTime   int64  `json:"startTime"`
	ServiceType string `json:"serviceType"`
	EndTime     *int64 `json:"endTime"`
	IsError     bool   `json:"isError"`
	Data        *struct {
		IP   string `json:"ip"`
		IPv6 string `json:"ipv6"`
		Ygg  string `json:"ygg"`
		Plan struct {
			ID      int    `json:"id"`
			CPU     int    `json:"cpu"`
			RAM     int    `json:"ram"`
			Hdd     int    `json:"hdd"`
			HddType string `json:"hddType"`
		} `json:"plan"`
	} `json:"data"`
}

type VMResponseTask struct {
	ID     string `json:"id"`
	TaskID string `json:"task_id"`
}
