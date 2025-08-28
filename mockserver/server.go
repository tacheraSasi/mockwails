package mockserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tacheraSasi/mockwails/db"
	"github.com/tacheraSasi/mockwails/utils"
)

// NotFoundData holds the template data for the 404 page
type NotFoundData struct {
	Method             string
	Endpoint           string
	Port               int
	AvailableEndpoints []db.Server
}

// serveCustom404 serves the custom 404 HTML page or json based on the method type
func serveCustom404(w http.ResponseWriter, r *http.Request, endpoint string, port int) {
	availableEndpoints, _ := db.GetAvailableEndpointsForPort(port)

	tmpl, err := template.New("404").Parse(custom404HTML)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Endpoint Not Found"))
		return
	}

	data := NotFoundData{
		Method:             r.Method,
		Endpoint:           endpoint,
		Port:               port,
		AvailableEndpoints: availableEndpoints,
	}

	// We Serve JSON for non GET requests
	if data.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(data)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Endpoint Not Found"))
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write(buf.Bytes())
}


// serveMethodNotAllowed serves a method not allowed response
func serveMethodNotAllowed(w http.ResponseWriter, r *http.Request, endpoint string, port int, allowedMethod string) {
	w.Header().Set("Allow", allowedMethod)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)

	response := map[string]interface{}{
		"error":   "Method Not Allowed",
		"message": fmt.Sprintf("The endpoint %s only accepts %s requests", endpoint, allowedMethod),
		"details": map[string]interface{}{
			"requested_method": r.Method,
			"allowed_method":   allowedMethod,
			"endpoint":         endpoint,
			"port":             port,
		},
	}

	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}

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

	// Create a catch-all handler that checks for configured endpoints
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == endpoint {
			if r.Method != method {
				serveMethodNotAllowed(w, r, endpoint, port, method)
				return
			}

			// Check if endpoint exists and is active
			if exists, _ := db.CheckIfEndpointExists(endpoint, port); !exists {
				serveCustom404(w, r, endpoint, port)
				return
			}

			// This is our configured endpoint with correct method, handle it normally
			for k, v := range responseHeaders {
				w.Header().Set(k, v)
			}
			w.WriteHeader(responseStatus)
			w.Write([]byte(responseBody))
			return
		}

		// For any other path, we serve custom 404
		serveCustom404(w, r, r.URL.Path, port)
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
		//  JSON
		_ = json.Unmarshal([]byte(headerStr), &headers)
	} else {
		//  key:value\nkey:value
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
