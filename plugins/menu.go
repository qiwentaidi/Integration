package plugins

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var modle bool

func MyMenu() *fyne.MainMenu {
	w := fyne.CurrentApp().Driver().AllWindows()[0]
	open := fyne.NewMenuItem("打开目录", func() {
		path, _ := os.Getwd()
		exec.Command("explorer", path).Start()
	})
	add := fyne.NewMenuItem("添加分页", func() {
		p := AddFirstNode()
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
	m := fyne.NewMainMenu(fyne.NewMenu("设置", open, add), fyne.NewMenu("其他", mstsc, regedit, memo), fyne.NewMenu("帮助", about))
	return m
}

func AddFirstNode() *widget.PopUp {
	w := fyne.CurrentApp().Driver().AllWindows()[0]
	title := widget.NewLabelWithStyle("请输入主节点名称:", fyne.TextAlignCenter, fyne.TextStyle{})
	hostnname := widget.NewEntry() // 要接收的名字
	p := widget.NewModalPopUp(nil, w.Canvas())
	save := widget.NewButton("保存", func() { // 添加按钮
		config := GetConfig()
		// 获取所有第一级子节点的 key
		for key := range config {
			modle = false
			if hostnname.Text == key {
				p.Hide()
				dialog.ShowInformation("提示", "已存在该名称的主节点，请不要重复添加", w)
				modle = true
				break
			}
		}
		if !modle {
			appMap[""] = append(appMap[""], hostnname.Text)
			T.Refresh()
			p.Hide()
			config[hostnname.Text] = make(map[string]interface{})
			SaveConfig(config)
		}
	})
	cancle := widget.NewButton("退出", func() {
		p.Hide()
	})
	p.Content = container.NewBorder(title, container.NewGridWithColumns(2, save, cancle), nil, nil, hostnname)
	p.Resize(fyne.NewSize(200, 100))
	return p
}
