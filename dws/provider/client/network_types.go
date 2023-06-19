package client

import (
	"net"
)

type VPCConfig struct {
	ID          string
	Name        string
	Description string
	IPRange     net.IPNet
}
