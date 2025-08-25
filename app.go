package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
    log.Println("CreateServer called with server:", server)
    log.Println("SERVER NAME:", server.Name)
}