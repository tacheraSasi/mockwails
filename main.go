package main

import (
	"embed"
	"strings"

	"github.com/tacheraSasi/mockwails/db"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	db.AutoMigrate()
	app := NewApp()

	err := wails.Run(&options.App{
		Title:         "MockWails",
		Width:         1200,
		Height:        700,
		DisableResize: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title: "MockWails",
				Message: strings.Join([]string{
					"© 2025 Tachera Sasi",
				}, "\n"),
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
