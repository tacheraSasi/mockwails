package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}
type Server struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Endpoint         string `json:"endpoint"`
	Method           string `json:"method"`
	RequestHeaders   string `json:"requestHeaders"`
	RequestBody      string `json:"requestBody"`
	ResponseStatus   int    `json:"responseStatus"`
	ResponseHeaders  string `json:"responseHeaders"`
	ResponseBody     string `json:"responseBody"`
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

func (a *App) CreateServer(data interface{}){
	fmt.Println(data,"CreateServer called")
}
