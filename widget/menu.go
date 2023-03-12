package widget

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
)

var (
	modle     bool
	tabsmodle bool
)

func Thememenu(tabs *container.AppTabs) *fyne.MainMenu {
	w := fyne.CurrentApp().Driver().AllWindows()[0]
	open := fyne.NewMenuItem("打开目录", func() {
		path, _ := os.Getwd()
		exec.Command("explorer", path).Start()
	})
	add := fyne.NewMenuItem("添加分页", func() {
		e := widget.NewEntry()
		p := Popup("添加分页", "分页名称:", e, func() {
			f, _ := os.Open("config.json")
			// 解析 json 文件
			var config map[string]interface{}
			decoder := json.NewDecoder(f)
			decoder.Decode(&config)
			// 获取所有第一级子节点的 key
			for key := range config {
				modle = false
				if e.Text == key {
					dialog.ShowInformation("提示", "已存在重名分页，请不要重复添加", w)
					modle = true
					break
				}
			}
			if !modle {
				c, _ := AppItem(tabs, e.Text)
				tabs.Append(container.NewTabItem(e.Text, c))
				viper.SetConfigFile("./config.json")
				if err := viper.ReadInConfig(); err != nil {
					fmt.Printf("err: %v\n", err)
				}
				viper.Set(e.Text, "")
				if err2 := viper.WriteConfigAs("./config.json"); err2 != nil {
					fmt.Printf("err2: %v\n", err2)
				}
			}
		})
		p.Show()
	})
	regedit := fyne.NewMenuItem("注册表", func() {
		exec.Command("regedit").Run()
	})
	mstsc := fyne.NewMenuItem("远程桌面", func() {
		exec.Command("mstsc").Start()
	})
	memo := fyne.NewMenuItem("备忘录", func() {
		txt := widget.NewMultiLineEntry()
		exit := widget.NewButton("保存&退出", nil)
		if _, err := os.Stat("memo.txt"); err != nil {
			os.Create("memo.txt")
		}
		b, _ := ioutil.ReadFile("memo.txt")
		txt.SetText(string(b))
		txt.Refresh()
		p := widget.NewModalPopUp(container.NewBorder(widget.NewLabelWithStyle("备忘录", fyne.TextAlignCenter, fyne.TextStyle{}), exit, nil, nil, txt), w.Canvas())
		exit.OnTapped = func() {
			p.Hide()
			f, _ := os.OpenFile("memo.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			defer f.Close()
			io.WriteString(f, txt.Text)
		}
		p.Resize(fyne.NewSize(400, 300))
		p.Show()
	})
	about := fyne.NewMenuItem("关于", func() {
		dialog.ShowInformation("", "Faster startup app", w)
	})
	warn := fyne.NewMenuItem("提示", func() {
		dialog.ShowInformation("可能遇到的错误",
			`在添加burp bat程序指向时无法启动
			bat修改-javaagent:%~dp0burp-loader-keygen-2_1_06.jar,%~dp0为指定bat当前路径`, w)
	})

	hv := fyne.NewMenuItem("分页样式", func() {
		if tabsmodle {
			tabs.SetTabLocation(container.TabLocationTop)
			tabsmodle = false
		} else {
			tabs.SetTabLocation(container.TabLocationLeading)
			tabsmodle = true
		}
	})
	m := fyne.NewMainMenu(fyne.NewMenu("设置", open, add), fyne.NewMenu("样式", hv), fyne.NewMenu("其他", mstsc, regedit, memo), fyne.NewMenu("帮助", about, warn))
	return m
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
