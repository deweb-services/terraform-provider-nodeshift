package client

type CreateVPCRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IPRange     string `json:"ipRange"`
}

type createVPCResponse struct {
	UUID string `json:"uuid"`
}

type GetVPCResponse struct {
	UUID        string      `json:"uuid"`
	Name        string      `json:"name"`
	State       string      `json:"state"`
	Description string      `json:"description"`
	IPRange     string      `json:"addressRange"`
	Resources   []Resources `json:"resources"`
}

type Resources struct {
	IP     string `json:"ip"`
	Status string `json:"status"`
}
