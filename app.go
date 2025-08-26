package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/tacheraSasi/mockwails/goofer"
)

type App struct {
	ctx context.Context
}
type Server struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Endpoint        string `json:"endpoint"`
	Method          string `json:"method"`
	RequestHeaders  string `json:"requestHeaders"`
	RequestBody     string `json:"requestBody"`
	ResponseStatus  int    `json:"responseStatus"`
	ResponseHeaders string `json:"responseHeaders"`
	ResponseBody    string `json:"responseBody"`
	Status          string `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`

	// goofer.ServerEntity  //TODO: Lookup a better way to handle this
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// CreateServer creates a new server in the database
func (a *App) CreateServer(data map[string]interface{}) {
	var server Server
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("Failed to marshal data:", err)
		return
	}
	if err := json.Unmarshal(b, &server); err != nil {
		log.Println("Failed to unmarshal server data:", err)
		return
	}
	err = goofer.CreateServer(goofer.ServerEntity{
		ID:              server.ID, //
		Name:            server.Name,
		Description:     server.Description,
		Endpoint:        server.Endpoint,
		Method:          server.Method,
		RequestHeaders:  server.RequestHeaders,
		RequestBody:     server.RequestBody,
		ResponseStatus:  server.ResponseStatus,
		ResponseHeaders: server.ResponseHeaders,
		ResponseBody:    server.ResponseBody,
	})
	if err != nil {
		//TODO: Will shift to a more robust error handling
		log.Fatal("Failed to create server:", err)
		return
	}
	log.Println("CreateServer called with server:", server)
	log.Println("SERVER NAME:", server.Name)
}

// GetAllServers returns all servers from the database
func (a *App) GetAllServers() ([]Server, error) {
	servers, err := goofer.GetAllServers()
	if err != nil {
		log.Println("Failed to get servers:", err)
		return nil, err
	}
	// NOTE: I Convert goofer.Server to main.Server if needed (fields are the same)
	var result []Server
	for _, s := range servers {
		result = append(result, Server{
			ID:              s.ID,
			Name:            s.Name,
			Description:     s.Description,
			Endpoint:        s.Endpoint,
			Method:          s.Method,
			RequestHeaders:  s.RequestHeaders,
			RequestBody:     s.RequestBody,
			ResponseStatus:  s.ResponseStatus,
			ResponseHeaders: s.ResponseHeaders,
			ResponseBody:    s.ResponseBody,
			// ServerEntity:    s,
		})
	}
	return result, nil
}

// UpdateServer updates a server in the database
func (a *App) UpdateServer(data map[string]interface{}) {
	var server Server
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("Failed to marshal data:", err)
		return
	}
	if err := json.Unmarshal(b, &server); err != nil {
		log.Println("Failed to unmarshal server data:", err)
		return
	}
	err = goofer.UpdateServer(goofer.ServerEntity{
		ID:              server.ID,
		Name:            server.Name,
		Description:     server.Description,
		Endpoint:        server.Endpoint,
		Method:          server.Method,
		RequestHeaders:  server.RequestHeaders,
		RequestBody:     server.RequestBody,
		ResponseStatus:  server.ResponseStatus,
		ResponseHeaders: server.ResponseHeaders,
		ResponseBody:    server.ResponseBody,
	})
	if err != nil {
		log.Println("Failed to update server:", err)
		return
	}
	log.Println("UpdateServer called with server:", server)
}

// DeleteServer deletes a server by ID
func (a *App) DeleteServer(id uint) {
	err := goofer.DeleteServer(id)
	if err != nil {
		log.Println("Failed to delete server:", err)
		return
	}
	log.Println("DeleteServer called for ID:", id)
}
