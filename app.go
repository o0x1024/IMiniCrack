package main

import (
	"IMiniCrack/pkg/util"
	"context"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	runtime2 "runtime"
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
func (a *App) OpenFile(curPath string) string {
	//file := p.rt.Dialog.SelectFile()
	parent := ""
	if curPath == "" {
		osType := runtime2.GOOS
		switch osType {
		case "windows":
			parent = "C:\\"
		case "linux":
			parent = "/"
		case "darwin":
			parent = "/"
		}
	} else {
		parent = util.GetParentDirectory(curPath)

	}

	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Select File",
		DefaultDirectory: parent,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "wxpack (*.wxapkg)",
				Pattern:     "*.wxapkg",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return file
}

func (a *App) OpenDir() string {

	dd := ""
	osType := runtime2.GOOS
	switch osType {
	case "windows":
		dd = "C:\\"
	case "linux":
		dd = "/"
	case "darwin":
		dd = "/"
	}
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Select Directory",
		DefaultDirectory: dd,
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
