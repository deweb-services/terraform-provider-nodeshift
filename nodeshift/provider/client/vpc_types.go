package client

type CreateVPCRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IPRange     string `json:"ipRange,omitempty"`
}

type createVPCResponse struct {
	UUID string `json:"id"`
}

type GetVPCResponse struct {
	UUID        string      `json:"uuid"`
	Name        string      `json:"name"`
	Status      string      `json:"status"`
	Description string      `json:"description"`
	IPRange     string      `json:"addressRange"`
	Resources   []Resources `json:"resources"`
}

type Resources struct {
	IP     string `json:"ip"`
	Status string `json:"status"`
}
