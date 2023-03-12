package main

import (
	"integration/common"
	"integration/widget"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/flopp/go-findfont"
)

// 设置中文字体
func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "Dengb.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func myapp() {
	a := app.New()
	w := a.NewWindow("Dashboard")
	w.Canvas()
	w.CenterOnScreen()
	w.SetIcon(theme.GridIcon())
	//go common.MoveHide()
	common.SystemTray() // 调用系统托盘功能
	tabs := container.NewAppTabs()
	widget.AppItemInit(tabs)              // 初始化界面
	w.SetMainMenu(widget.Thememenu(tabs)) // 设置菜单栏
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(700, 500))
	w.ShowAndRun()
}

func main() {
	common.Profile()
	myapp()
}
