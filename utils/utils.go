package utils

import (
	"fmt"
	"net"

	"github.com/jaypipes/ghw/pkg/pci/address"
)

func IsPortInUse(port int)bool{
	address := fmt.Sprintf(":%d",port)
	ln, err := net.Listen("tcp",address)
	if err != nil {
		return false
	}
	defer ln.Close()
	return true
}