package client

type VPCConfig struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IPRange     string `json:"ipRange"`
}

type VPCConfigResponse struct {
	ID          string `json:"uuid,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IPRange     string `json:"addressRange"`
}
