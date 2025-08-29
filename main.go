package main

import (
	"embed"
	"strings"

	"github.com/tacheraSasi/mockwails/config"
	"github.com/tacheraSasi/mockwails/db"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS



func main() {
	db.AutoMigrate()

	app := NewApp()

	err := wails.Run(&options.App{
		Title:         config.GetAppDetails().Name,
		Width:         1200,
		Height:        700,
		MinWidth:      800,
		MinHeight:     600,
		MaxWidth:      1920,
		MaxHeight:     1080,
		DisableResize: false,
		Fullscreen:    false,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1}, // Dark theme background
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
		CSSDragProperty: "--wails-draggable",
		CSSDragValue:    "drag",
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
		Mac: &mac.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title: config.GetAppDetails().Name + " v" + config.GetAppDetails().Version,
				Message: strings.Join([]string{
					"A powerful desktop application for creating and managing mock HTTP servers.",
					"",
					"Built with Wails, Go, and React.",
					"",
					"© 2025 Tachera Sasi",
				}, "\n"),
				Icon: nil,
			},
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			Theme:                             windows.SystemDefault,
		},
		Linux: &linux.Options{
			Icon:                []byte{},
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         config.GetAppDetails().Name,
		},
	})

	if err != nil {
		println("Fatal: Failed to start application:", err.Error())
	}
}
