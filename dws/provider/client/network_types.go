package client

import (
	"net"
)

type NetworkConfig struct {
	ID          string
	Name        string
	Description string
	IPRange     net.IPNet
	// Nodes        []uint32
	// AddWGAccess  bool
	// SolutionType string
	// NodesIPRange map[uint32]net.IPNet

	// computed
	// AccessWGConfig   string
	// ExternalIP       *net.IPNet
	// ExternalSK       wireGuard.Key
	// PublicNodeID     uint32
	// NodeDeploymentID map[uint32]uint64

	// WGPort map[uint32]int
	// Keys   map[uint32]wireGuard.Key
}
