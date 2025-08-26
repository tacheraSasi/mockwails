package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tacheraSasi/mockwails/db"
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
	err = db.CreateServer(&server)
	if err != nil {
		log.Println("Failed to create server:", err)
		return
	}
	log.Println("CreateServer called with server:", server)
	log.Println("SERVER NAME:", server.Name)
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
	err = db.UpdateServer(&server)
	if err != nil {
		log.Println("Failed to update server:", err)
		return
	}
	log.Println("UpdateServer called with server:", server)
}

// DeleteServer deletes a server by ID
func (a *App) DeleteServer(id uint) {
	err := db.DeleteServer(id)
	if err != nil {
		log.Println("Failed to delete server:", err)
		return
	}
	log.Println("DeleteServer called for ID:", id)
}
