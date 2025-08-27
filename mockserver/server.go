package mockserver

import (
	"github.com/tacheraSasi/mockwails/db"
	"github.com/tacheraSasi/mockwails/utils"
)

func Start(){
	//TODO: Start the mock server
}

func Stop(){
	//TODO: Stop a running mock server
}

func CheckStatus(server db.Server)string{
	isRunning:= utils.IsPortInUse(server.AddressAssigned.Port)
	if isRunning {
		return "active"
	}
	return "inactive"
}

