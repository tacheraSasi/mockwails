package mockserver

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/tacheraSasi/mockwails/db"
)

// SharedServerManager manages HTTP servers that can handle multiple endpoints
type SharedServerManager struct {
	servers map[int]*http.Server
	muxes   map[int]*http.ServeMux
	mutex   sync.RWMutex
}

var serverManager = &SharedServerManager{
	servers: make(map[int]*http.Server),
	muxes:   make(map[int]*http.ServeMux),
}

// GetOrCreateSharedServer gets an existing server for the port or creates a new one
func (sm *SharedServerManager) GetOrCreateSharedServer(port int) (*http.ServeMux, error) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	// If server already exists for this port, return its mux
	if mux, exists := sm.muxes[port]; exists {
		return mux, nil
	}

	// Create new mux and server for this port
	mux := http.NewServeMux()
	sm.muxes[port] = mux

	// Create the main handler that routes to all active endpoints on this port
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sm.handleRequest(w, r, port)
	})

	addr := ":" + strconv.Itoa(port)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	sm.servers[port] = server

	// Start the server in a goroutine
	go func() {
		log.Printf("Starting shared mock server on port %d", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Shared mock server on port %d stopped: %v", port, err)
		}
	}()

	return mux, nil
}

// handleRequest handles incoming requests and routes them to the appropriate endpoint
func (sm *SharedServerManager) handleRequest(w http.ResponseWriter, r *http.Request, port int) {
	endpoint := r.URL.Path
	method := r.Method

	// Get all active endpoints for this port
	servers, err := db.GetAvailableEndpointsForPort(port)
	if err != nil {
		log.Printf("Error getting endpoints for port %d: %v", port, err)
		serveCustom404(w, r, endpoint, port)
		return
	}

	// Find matching endpoint and method
	var matchingServer *db.Server
	for _, server := range servers {
		if server.Endpoint == endpoint && server.Method == method {
			matchingServer = &server
			break
		}
	}

	if matchingServer == nil {
		// Check if endpoint exists with different method
		for _, server := range servers {
			if server.Endpoint == endpoint {
				serveMethodNotAllowed(w, r, endpoint, port, server.Method)
				return
			}
		}
		// No matching endpoint found
		serveCustom404(w, r, endpoint, port)
		return
	}

	// Serve the response from the matching server
	responseHeaders := parseHeaders(matchingServer.ResponseHeaders)
	for k, v := range responseHeaders {
		w.Header().Set(k, v)
	}
	w.WriteHeader(matchingServer.ResponseStatus)
	w.Write([]byte(matchingServer.ResponseBody))
}

// StopServer stops the server for a specific port if no more endpoints are using it
func (sm *SharedServerManager) StopServer(port int) error {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	// Check if there are any active endpoints still using this port
	servers, err := db.GetAvailableEndpointsForPort(port)
	if err != nil {
		return err
	}

	// If there are still active endpoints, don't stop the server
	if len(servers) > 0 {
		return nil
	}

	// Stop the server if it exists
	if server, exists := sm.servers[port]; exists {
		log.Printf("Stopping shared server on port %d", port)
		server.Close()
		delete(sm.servers, port)
		delete(sm.muxes, port)
	}

	return nil
}

// IsPortManaged returns true if the port is managed by the shared server manager
func (sm *SharedServerManager) IsPortManaged(port int) bool {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	_, exists := sm.servers[port]
	return exists
}