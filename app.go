package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tacheraSasi/mockwails/db"
	"github.com/tacheraSasi/mockwails/mockserver"
	"github.com/tacheraSasi/mockwails/utils"
)

type App struct {
	ctx context.Context
}
type Server = db.Server

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	fmt.Println("Seeding servers...")
	if err := db.SeedServers(); err != nil {
		fmt.Println("Failed to seed servers:", err)
		return
	}
	fmt.Println("App started...")
	//TODO: Here i should start all the servers from the db where status is active
	//Meaning they where not stoped but the app was shutdown
	//TODO: In the future i will add a column auto-start for specifically this use
	servers, err := db.GetAllActiveServers()
	if err != nil {
		fmt.Println("Failed to get active servers:", err)
		return
	}
	for _, server := range servers {
		response := a.StartServer(server.ID)
		if !response.Success {
			fmt.Printf("Failed to start server %s: %s\n", server.Name, response.Message)
		}
	}
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	fmt.Println("App shutting down...")
	servers, err := db.GetAllActiveServers()
	if err != nil {
		fmt.Println("Failed to get active servers:", err)
		return
	}
	for _, server := range servers {
		response := a.StopServer(server.ID)
		if !response.Success {
			fmt.Printf("Failed to stop server %s: %s\n", server.Name, response.Message)
		}
	}
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// CreateServer creates a new server in the database
func (a *App) CreateServer(data map[string]interface{}) utils.Response {
	fmt.Println("CreateServer called with data:", data)
	var server Server
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("Failed to marshal data:", err)
		return utils.Response{Success: false, Message: "Failed to create server: " + err.Error()}
	}
	if err := json.Unmarshal(b, &server); err != nil {
		log.Println("Failed to unmarshal server data:", err)
		return utils.Response{Success: false, Message: "Failed to create server: " + err.Error()}
	}

	// Set the server status to inactive initially
	server.Status = "inactive"

	err = db.CreateServer(&server)
	if err != nil {
		log.Println("Failed to create server:", err)
		return utils.Response{Success: false, Message: "Failed to create server: " + err.Error()}
	}
	log.Println("CreateServer called with server:", server)
	log.Println("SERVER NAME:", server.Name)

	// Start the server after creation
	startResponse := a.StartServer(server.ID)
	if !startResponse.Success {
		log.Println("Failed to start server:", startResponse.Message)
		return utils.Response{Success: false, Message: "Server created but failed to start: " + startResponse.Message}
	}

	log.Println("SERVER started:", server.Name)
	return utils.Response{Success: true, Message: "Server created and started successfully", Data: server}
}

// GetAllServers returns all servers from the database
func (a *App) GetAllServers() ([]Server, error) {
	servers, err := db.GetAllServers()
	if err != nil {
		log.Println("Failed to get servers:", err)
		return nil, err
	}
	return servers, nil
}

// GetServerByID retrieves a server by its ID
func (a *App) GetServerByID(id uint) (*Server, error) {
	servers, err := db.GetAllServers()
	if err != nil {
		log.Println("Failed to get servers:", err)
		return nil, err
	}
	for _, server := range servers {
		if server.ID == id {
			return &server, nil
		}
	}
	return nil, fmt.Errorf("server with ID %d not found", id)
}

// UpdateServer updates a server in the database
func (a *App) UpdateServer(data map[string]any) utils.Response {
	var server Server
	b, err := json.Marshal(data)
	if err != nil {
		log.Println("Failed to marshal data:", err)
		return utils.Response{Success: false, Message: "Failed to update server: " + err.Error()}
	}
	if err := json.Unmarshal(b, &server); err != nil {
		log.Println("Failed to unmarshal server data:", err)
		return utils.Response{Success: false, Message: "Failed to update server: " + err.Error()}
	}
	err = db.UpdateServer(&server)
	if err != nil {
		log.Println("Failed to update server:", err)
		return utils.Response{Success: false, Message: "Failed to update server: " + err.Error()}
	}
	log.Println("UpdateServer called with server:", server)
	return utils.Response{Success: true, Message: "Server updated successfully", Data: server}
}

// DeleteServer deletes a server by ID
func (a *App) DeleteServer(id uint) utils.Response {
	// First stop the server if it's running
	server, err := db.GetServerByID(id)
	if err != nil {
		log.Println("Failed to get server:", err)
		return utils.Response{Success: false, Message: "Failed to get server: " + err.Error()}
	}

	if server.Status == "active" {
		stopResponse := a.StopServer(id)
		if !stopResponse.Success {
			log.Println("Failed to stop server before deletion:", stopResponse.Message)
			return utils.Response{Success: false, Message: "Failed to stop server before deletion: " + stopResponse.Message}
		}
	}

	err = db.DeleteServer(id)
	if err != nil {
		log.Println("Failed to delete server:", err)
		return utils.Response{Success: false, Message: "Failed to delete server: " + err.Error()}
	}
	log.Println("DeleteServer called for ID:", id)
	return utils.Response{Success: true, Message: "Server deleted successfully"}
}

// StartServer starts a server by ID
func (a *App) StartServer(id uint) utils.Response {
	server, err := db.GetServerByID(id)
	if err != nil {
		log.Println("Failed to get server:", err)
		return utils.Response{Success: false, Message: "Failed to get server: " + err.Error()}
	}
	log.Println("StartServer called for server:", server)

	if server.Status == "inactive" {
		err = db.ToggleServerStatus(server.ID)
		if err != nil {
			return utils.Response{Success: false, Message: "Failed to toggle server status: " + err.Error()}
		}
	}

	mockserver.CheckStatus(*server)
	err = mockserver.Start(*server)
	if err != nil {
		log.Println("Failed to start mock server:", err)
		if server.Status == "inactive" {
			db.ToggleServerStatus(server.ID)
		}
		return utils.Response{Success: false, Message: "Failed to start mock server: " + err.Error()}
	}

	return utils.Response{Success: true, Message: "Server started successfully", Data: server}
}

func (a *App) StopServer(id uint) utils.Response {
	server, err := db.GetServerByID(id)
	if err != nil {
		log.Println("Failed to get server:", err)
		return utils.Response{Success: false, Message: "Failed to get server: " + err.Error()}
	}
	log.Println("StopServer called for server:", server)

	err = mockserver.Stop(*server)
	if err != nil {
		log.Println("Failed to stop server:", err)
		return utils.Response{Success: false, Message: "Failed to stop server: " + err.Error()}
	}

	mockserver.CheckStatus(*server)
	return utils.Response{Success: true, Message: "Server stopped successfully", Data: server}
}
