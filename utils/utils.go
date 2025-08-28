package utils

import (
	"fmt"
	"net"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func IsPortInUse(port int) bool {
	address := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return true // Port is in use
	}
	defer ln.Close()
	return false // Port is available
}
