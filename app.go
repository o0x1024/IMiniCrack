package main

import (
	"IMiniCrack/pkg/model"
	"IMiniCrack/pkg/scan"
	"context"
	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"os/exec"
	"os/user"
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
	runtime.WindowSetBackgroundColour(a.ctx, 255, 255, 255, 255)
}

// 保存正则信息
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	resp := scan.Sc.SaveRegex()
	if resp.Err != "" {
		dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:    runtime.QuestionDialog,
			Title:   "Quit?",
			Message: "正则保存失败，原因：" + resp.Err + ",确定要关闭么？",
		})

		if err != nil {
			return false
		}
		return dialog != "Yes"
	}
	return false
}

func (a *App) OpenDecDir(path string) string {
	osType := runtime2.GOOS
	curPath, err := os.Getwd()
	if err != nil {
		return err.Error()
	}
	switch osType {
	case "windows":
		exec.Command(`cmd`, `/c`, `explorer`, path).Start()
	//case "linux":
	//	dd = "/"
	case "darwin":
		exec.Command(`open`, curPath).Start()
	}
	return ""
}

func (a *App) SelectOpenFile() (resp model.Response) {
	user, _ := user.Current()
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Select Directory",
		DefaultDirectory: user.HomeDir + "\\Documents",
		ShowHiddenFiles:  true,
	})
	if err != nil {
		resp.Err = err.Error()
		return resp
	}
	resp.Data = path
	return resp
}

func (a *App) OpenDisFile(path string) (resp model.Response) {

	data, err := os.ReadFile(path)
	if err != nil {
		resp.Err = err.Error()
		return resp
	}

	resp.Data = string(data)
	return resp
}

func (a *App) OpenWxPackDir(curPath string) string {
	//file := p.rt.Dialog.SelectFile()
	//curPath, err := os.Getwd()
	user, err := user.Current()
	if err != nil {
		return "OpenWxPackDir" + err.Error()
	}
	if err != nil {
		return err.Error()
	}

	wxPath := ""

	osType := runtime2.GOOS
	switch osType {
	case "windows":
		wxPath = user.HomeDir + "\\Documents\\WeChat Files\\Applet"
	case "linux":
		wxPath = user.HomeDir
	case "darwin":
		wxPath = user.HomeDir
	}

	file, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Select wxpack Dir",
		DefaultDirectory: wxPath,
	})
	if err != nil {
		panic(err)
	}
	return file
}

func (a *App) GetDefaultOutPath() string {
	user, err := user.Current()
	if err != nil {
		return "OpenWxPackDir" + err.Error()
	}

	return user.HomeDir + "\\Documents"
}

func (a *App) OpenDir() string {

	user, err := user.Current()
	if err != nil {
		return "OpenWxPackDir" + err.Error()
	}
	dd := ""
	osType := runtime2.GOOS
	switch osType {
	case "windows":
		dd = user.HomeDir + "\\Documents"
	case "linux":
		dd = user.HomeDir
	case "darwin":
		dd = user.HomeDir
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

func (a *App) OpenScanDir(path string) string {
	if path == "" {
		osType := runtime2.GOOS
		switch osType {
		case "windows":
			path = "C:\\"
		case "linux":
			path = "/"
		case "darwin":
			path = "/"
		}
	}
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "Select Directory",
		DefaultDirectory: path,
	})
	if err != nil {
		panic(err)
	}
	return dir
}

func (p *App) WailsInit(runtime *wails.Runtime) error {
	p.rt = runtime

	return nil
}
