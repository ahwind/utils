package network

import (
	"net"
)

type Interface struct {
	Name  string   // Name of the interface
	Mac   string   // Mac address of the interface
	Addrs []string // List of IP addresses of the interface
}

/*
	Get list of network interfaces
*/
func GetNetInterfaces() ([]Interface, error) {
	var interfaces []Interface
	var err error
	var netInterfaces []net.Interface
	var netInterface net.Interface

	netInterfaces, err = net.Interfaces()
	if err != nil {
		return interfaces, err
	}

	for _, netInterface = range netInterfaces {
		var netIf Interface
		netIf.Name = netInterface.Name
		netIf.Mac = netInterface.HardwareAddr.String()
		tmpNetIf, err := netInterface.Addrs()
		if err != nil {
			return interfaces, err
		}
		for _, ipAddr := range tmpNetIf {
			netIf.Addrs = append(netIf.Addrs, ipAddr.String())
		}
		interfaces = append(interfaces, netIf)
	}

	return interfaces, nil
}
