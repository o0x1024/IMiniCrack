package main

import (
	"IMiniCrack/pkg/crack"
	"IMiniCrack/pkg/scan"
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure

	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	app := NewApp()
	c := &crack.Crack{}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "IMiniCrack   -  gelenlen",
		Width:  1200,
		Height: 960,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnDomReady:       c.GetCtx,
		OnBeforeClose:    app.beforeClose,
		Bind: []interface{}{
			app,
			c,
			scan.Sc,
		},
		Windows: &windows.Options{},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
