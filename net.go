package vm

import (
	"fmt"
)

type NetworkDevice struct {
	Type string // Netdev type (user, tap...)
	ID   string // Netdev ID

	Name string // TAP: interface name
	MAC    string // TAP: Interface hardware address
}

// NewNetworkDevice creates a QEMU network
// device
func NewNetworkDevice(deviceType, id string) (NetworkDevice, error) {
	var netdev NetDev

	if deviceType != "user" && deviceType != "tap" {
		return netdev, fmt.Errorf("Unsupported netdev type")
	}
	if len(id) == 0 {
		return netdev, fmt.Errorf("You must specify a netdev ID")
	}

	netdev.Type = deviceType
	netdev.ID = id

	return netdev, nil
}

// SetHostInterfaceName sets the host interface name
// for the netdev (if supported by netdev type)
func (n *NetworkDevice) SetHostInterfaceName(name string) {
	n.Name = name
}

// SetMacAddress sets the mac address of the
// netdev (if supported by netdev type)
func (n *NetworkDevice) SetMacAddress(mac string) {
	n.MAC = mac
}
