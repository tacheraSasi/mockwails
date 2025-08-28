package mockserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tacheraSasi/mockwails/db"
	"github.com/tacheraSasi/mockwails/utils"
)

// Start launches a mock HTTP server based on the Server struct details
func Start(server db.Server) error {
	port := server.AddressAssigned.Port
	endpoint := server.Endpoint
	method := strings.ToUpper(server.Method)
	responseStatus := server.ResponseStatus
	responseHeaders := parseHeaders(server.ResponseHeaders)
	responseBody := server.ResponseBody

	if utils.IsPortInUse(port) {
		return fmt.Errorf("port %d is already in use", port)
	}

	mux := http.NewServeMux()
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if exists, _ := db.CheckIfEndpointExists(endpoint, port); !exists {
			http.Error(w, "Endpoint Not Found", http.StatusNotFound)
			w.WriteHeader(http.StatusNotFound)
			//TODO: I will show here a custom not found page
			return
		}
		for k, v := range responseHeaders {
			w.Header().Set(k, v)
		}
		w.WriteHeader(responseStatus)
		w.Write([]byte(responseBody))
	})

	addr := ":" + strconv.Itoa(port)
	go func() {
		log.Printf("Starting mock server on %s%s", addr, endpoint)
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.Printf("Mock server stopped: %v", err)
		}
	}()

	return nil
}

// parseHeaders parses a JSON or key-value string into a map
func parseHeaders(headerStr string) map[string]string {
	headers := make(map[string]string)
	if strings.HasPrefix(headerStr, "{") {
		// Try JSON
		_ = json.Unmarshal([]byte(headerStr), &headers)
	} else {
		// Try key:value\nkey:value
		lines := strings.Split(headerStr, "\n")
		for _, line := range lines {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}
	}
	return headers
}

func Stop(server db.Server) error {
	err := db.ToggleServerStatus(server.ID)
	if err != nil {
		return err
	}
	//TODO: Stop a running mock server
	return nil
}

func CheckStatus(server db.Server) string {
	isRunning := utils.IsPortInUse(server.AddressAssigned.Port)
	if isRunning {
		return "active"
	}
	return "inactive"
}
