package main

import (
	"embed"
	"wails-selfupdate-spike/logging"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	logging.NewLogger(logging.NewLoggerParams{
		ResourceDir:          "/Users/sam/git/wails-selfupdate-spike/_logs",
		EnableDebugLogging:   true,
		EnableFileLogging:    true,
		EnableConsoleLogging: true,
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wails-selfupdate-spike",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
