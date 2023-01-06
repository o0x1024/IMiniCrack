package main

import (
	"context"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
)

// App struct
type App struct {
	rt  *wails.Runtime
	ctx context.Context
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

func (a *App) OpenDecDir(path string) {
	exec.Command(`cmd`, `/c`, `explorer`, path).Start()
}
func (a *App) OpenFile() string {
	//file := p.rt.Dialog.SelectFile()
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Select File",
		DefaultDirectory: "C:\\Users\\test\\Documents\\WeChat Files\\Applet\\",
	})
	if err != nil {
		panic(err)
	}
	return file
}

func (a *App) OpenDir() string {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Select Directory",
		DefaultDirectory: "C:\\",
	})
	if err != nil {
		panic(err)
	}
	return dir
}

// WailsInit assigns the runtime to the PortfallOS struct
func (p *App) WailsInit(runtime *wails.Runtime) error {
	p.rt = runtime
	return nil
}
