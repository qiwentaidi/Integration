package common

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	trayMode bool
)

// 系统托盘功能
func SystemTray() {
	w := fyne.CurrentApp().Driver().AllWindows()[0]
	if desk, ok := fyne.CurrentApp().(desktop.App); ok {
		m := fyne.NewMenu("",
			fyne.NewMenuItem("Show", func() {
				w.Show()
				trayMode = false
			}))
		desk.SetSystemTrayMenu(m)
	}
	w.SetCloseIntercept(func() {
		w.Hide()
		trayMode = true
	})
}
