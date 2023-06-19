package client

type VPCConfig struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IPRange     string `json:"ipRange"`
}
