package main

import (
	"log"
	"os/exec"
)

// For now i want to move the bin to a generic location
// and then call it from the main app

func main(){
	binName := "mockwails"
	binPath := "../build/bin/MockWails.app/Contents/MacOS/"+binName
	err := exec.Command("chmod", "+x", binPath, "&&", "sudo movebin", binPath).Run()
	if err != nil {
		log.Fatalf("Failed to start mockwails: %v", err)
	}
	log.Printf("Successfully moved %s to /usr/local/bin/%s", binName, binName)
	log.Printf("You can now run the application using the command: %s", binName)
}