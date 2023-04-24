package client

type VMConfig struct {
	ID         string       `json:"id"`
	Deployment VMDeployment `json:"deployment"`
	CPU        []VMCPU      `json:"cpu"`
	RAM        VMRAM        `json:"ram"`
	Disk       []VMDisk     `json:"disk"`
	Protocols  VMProtocols  `json:"protocols"`
}

type VMDeployment struct {
	Name    string `json:"name"`
	Image   string `json:"image"`
	Network string `json:"network"`
	Region  string `json:"region"`
}

type VMCPU struct {
	Quantity int64  `json:"quantity"`
	Type     string `json:"type"`
}

type VMRAM struct {
	Volume int64 `json:"volume"`
}

type VMDisk struct {
	Type   string `json:"type"`
	Volume int64  `json:"volume"`
}

type VMProtocols struct {
	IP VMProtocolsIP `json:"IP"`
}

type VMProtocolsIP struct {
	V4 bool `json:"v4"`
	V6 bool `json:"v6"`
}
