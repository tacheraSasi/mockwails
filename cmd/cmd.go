package main

import (
	"log"
	"os/exec"
)

// For now i want to move the bin to a generic location
// and then call it from the main app

func main() {
	binName := "mockwails"
	binPath := "../build/bin/mockwails.app/Contents/MacOS/mockwails"
	destPath := "/usr/local/bin/" + binName

	// Making the binary executable
	err := exec.Command("chmod", "+x", binPath).Run()
	if err != nil {
		log.Fatalf("Failed to make binary executable: %v", err)
	}

	// Copying the binary to /usr/local/bin
	err = exec.Command("sudo", "cp", binPath, destPath).Run()
	if err != nil {
		log.Fatalf("Failed to copy mockwails to /usr/local/bin: %v", err)
	}

	log.Printf("Successfully copied %s to %s", binName, destPath)
	log.Printf("You can now run the application using the command: %s", binName)
}
