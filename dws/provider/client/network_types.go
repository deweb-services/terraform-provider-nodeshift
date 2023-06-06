package client

import (
	"net"

	wireGuard "golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type NetworkConfig struct {
	ID           string
	Name         string
	Description  string
	Nodes        []uint32
	IPRange      net.IPNet
	AddWGAccess  bool
	SolutionType string
	NodesIPRange map[uint32]net.IPNet

	// computed
	AccessWGConfig   string
	ExternalIP       *net.IPNet
	ExternalSK       wireGuard.Key
	PublicNodeID     uint32
	NodeDeploymentID map[uint32]uint64

	WGPort map[uint32]int
	Keys   map[uint32]wireGuard.Key
}
