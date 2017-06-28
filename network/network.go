package network

import (
	"net"
	"strings"
)
// Returns the local IP address of the machine
func GetLocalIP() (address string, err error) {
	interfaces,err := net.Interfaces()
	if err != nil {
		return "",err
	}
	// Browse all interfaces
	for _,i := range interfaces {
    		// Skip loopbacks
		if (i.Flags & net.FlagLoopback) != 0 {
			continue
		}
    		// Ensure interface is up
		if !((i.Flags & net.FlagUp) != 0) {
			continue
		}
    		// Parse the address
		addr,_ := i.Addrs()
		address := addr[0].String()
		pos := strings.Index(address,"/")
		if pos != -1 {
			address = address[0:pos]
		}
		return address,nil
	}
	return "",nil
}
