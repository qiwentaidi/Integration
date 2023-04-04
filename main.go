package main

/*
Integration UI重构版 - 设计思路
1、界面布局左侧 - 采用widget.Tree构建,指定成一级目录
2、界面布局右侧 - 采用container.DocTabs,配置成三级目录
3、界面右侧的container.DocTabs需继承上一版Integration的按钮功能,内部结构以container.NewGridWithColumns进行布局，一行可以显示5个按钮控件
4、一切配置均保存在config配置文件中,以为读取文件作为初始化配置的基准
*/

import (
	"integration/plugins"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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

func main() {
	plugins.CreateCofigFile()
	a := app.New()
	w := a.NewWindow("Integation-UI Reconstructed version 1.1")
	sbox := plugins.SplitBox()
	w.SetContent(sbox)
	w.SetMainMenu(plugins.MyMenu()) // 设置菜单栏
	w.Resize(fyne.NewSize(900, 600))
	w.SetIcon(theme.GridIcon())
	w.ShowAndRun()
}
