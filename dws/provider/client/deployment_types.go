package client

import "time"

/*
 * Deployment config represents payload that
 * Contains configuration of deployment to create
 */
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

// CreatedDeployment config
type CreatedDeployment struct {
	UUID         string    `json:"uuid"`
	Status       int       `json:"status"`
	IP           string    `json:"ip"`
	TaskID       string    `json:"taskId"`
	Cru          int       `json:"cru"`
	Mru          int       `json:"mru"`
	Sru          int       `json:"sru"`
	Hru          int       `json:"hru"`
	HddType      int       `json:"hddType"`
	Provider     int       `json:"provider"`
	Hostname     string    `json:"hostname"`
	Ipv6         int       `json:"ipv6"`
	SSHKey       string    `json:"sshKey"`
	Image        int       `json:"image"`
	ImageVersion string    `json:"imageVersion"`
	ChosenPlanID int       `json:"chosenPlanId"`
	Price        string    `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
}

// All structs declared below related to the asynchronous API
// It is different from the standard API

type AsyncAPIDeploymentResponse struct {
	StartTime    int64                   `json:"startTime"`
	ServiceType  string                  `json:"serviceType"`
	EndTime      *int64                  `json:"endTime"`
	IsError      bool                    `json:"isError"`
	Data         *DeploymentResponseData `json:"data"`
	FailedReason string                  `json:"failedReason"`

	// Not presented in this response, but
	// We still have to declare this property
	// So as to assign ID to deployment from the AsyncAPIDeploymentTask
	ID string `json:"-"`
}

type DeploymentResponseData struct {
	IP   string                     `json:"ip"`
	IPv6 string                     `json:"ipv6"`
	Ygg  string                     `json:"ygg"`
	Plan DeploymentResponseDataPlan `json:"plan"`
}

type DeploymentResponseDataPlan struct {
	ID      int    `json:"id"`
	CPU     int    `json:"cpu"`
	RAM     int    `json:"ram"`
	Hdd     int    `json:"hdd"`
	HddType string `json:"hddType"`
}

type AsyncAPIDeploymentTask struct {
	ID     string `json:"uuid"`
	TaskID string `json:"taskId"`
}
