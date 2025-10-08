package client

import "time"

/*
 * Deployment config represents payload that
 * Contains configuration of deployment to create.
 */
type CreateDeploymentRequest struct {
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
	SSHKeyName   string `json:"sshKeyName"`
	HostName     string `json:"hostName"`
	NetworkUUID  string `json:"networkUuid,omitempty"`
}

type createDeploymentResponse struct {
	UUID string `json:"uuid"`
}

// GetDeploymentResponse ...
type GetDeploymentResponse struct {
	UUID         string    `json:"uuid"`
	Status       string    `json:"status"`
	IP           string    `json:"ip"`
	Cru          int       `json:"cru"`
	Mru          int       `json:"mru"`
	Sru          int       `json:"sru"`
	Hru          int       `json:"hru"`
	HddType      int       `json:"hddType"`
	Provider     int       `json:"provider"`
	Hostname     string    `json:"hostname"`
	Ipv6         int       `json:"ipv6"`
	SSHKey       string    `json:"sshKey"`
	SSHKeyName   string    `json:"sshKeyName"`
	Image        int       `json:"image"`
	ImageVersion string    `json:"imageVersion"`
	ChosenPlanID int       `json:"chosenPlanId"`
	Price        string    `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
}
