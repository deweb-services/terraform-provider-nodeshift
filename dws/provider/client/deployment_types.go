package client

type DeploymentConfig struct {
	ImageVersion string `json:"imageVersion"`
	Region       string `json:"region"`
	CPU          int    `json:"cpu"`
	RAM          int    `json:"ram"`
	Hdd          int    `json:"hdd"`
	HddType      string `json:"hddType"`
	Ipv4         bool   `json:"ipv4,omitempty"`
	Ipv6         bool   `json:"ipv6,omitempty"`
	Ygg          bool   `json:"ygg,omitempty"`
	SSHKey       string `json:"sshKey"`
	HostName     string `json:"hostName"`
	NetworkUUID  string `json:"networkUuid,omitempty"`
}

type CreatedDeployment struct {
	StartTime   int64                  `json:"startTime"`
	ServiceType string                 `json:"serviceType"`
	EndTime     *int64                 `json:"endTime"`
	IsError     bool                   `json:"isError"`
	Data        *CreatedDeploymentData `json:"data"`

	//
	ID string `json:"-"`
}

type CreatedDeploymentData struct {
	IP   string                    `json:"ip"`
	IPv6 string                    `json:"ipv6"`
	Ygg  string                    `json:"ygg"`
	Plan CreatedDeploymentDataPlan `json:"plan"`
}

type CreatedDeploymentDataPlan struct {
	ID      int    `json:"id"`
	CPU     int    `json:"cpu"`
	RAM     int    `json:"ram"`
	Hdd     int    `json:"hdd"`
	HddType string `json:"hddType"`
}

type DeploymentCreateTask struct {
	ID     string `json:"id"`
	TaskID string `json:"taskId"`
}
