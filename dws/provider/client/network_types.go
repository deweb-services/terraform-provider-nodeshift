package client

import (
	"net"
)

type NetworkConfig struct {
	ID          string
	Name        string
	Description string
	IPRange     net.IPNet
}
