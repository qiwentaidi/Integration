package widget

import (
	"encoding/json"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/viper"
)

var checkmodle bool

// 初始化TabItem
func AppItemInit(tabs *container.AppTabs) {
	// 加载 json 文件
	f, _ := os.Open("config.json")
	// 解析 json 文件
	var config map[string]interface{}
	decoder := json.NewDecoder(f)
	decoder.Decode(&config)
	// 获取所有第一级子节点的 key
	for key := range config {
		c, c2 := AppItem(tabs, key)
		for _, v := range ButtonInit(c2, key) {
			c2.Add(v)
		}
		it := container.NewTabItem(key, c)
		it.Icon = theme.LoginIcon()
		tabs.Append(it)
	}
}

// 加减按钮的功能
func Popup(title, message string, entry *widget.Entry, OnTapped func()) (p *widget.PopUp) {
	subtitle := widget.NewLabelWithStyle(title, fyne.TextAlignCenter, fyne.TextStyle{}) // 标题
	entry.FocusGained()
	form := container.NewBorder(nil, nil, widget.NewLabel(message), nil, entry)
	save := widget.NewButtonWithIcon("save", theme.ConfirmIcon(), OnTapped)
	cancel := widget.NewButtonWithIcon("cancel", theme.CancelIcon(), func() {
		p.Hide()
	})
	subheader := container.NewVBox(subtitle, form, container.NewHBox(layout.NewSpacer(), save, cancel, layout.NewSpacer()))
	subframe := container.NewBorder(widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), subheader)
	p = widget.NewModalPopUp(subframe, fyne.CurrentApp().Driver().AllWindows()[0].Canvas())
	p.Resize(fyne.NewSize(200, 100))
	return p
}

// 初始化按钮，传入标题名称，从config.json中读取
func ButtonInit(ctrl *fyne.Container, title string) (objects []*RightClickButton) {
	viper.SetConfigFile("./config.json")
	viper.ReadInConfig()
	for key := range viper.GetStringMap(title) {
		rcb := NewRightClickButton(ctrl, title, key)
		objects = append(objects, rcb)
	}
	return objects
}

func AppItem(tabs *container.AppTabs, title string) (c *fyne.Container, vbox *fyne.Container) {
	vbox = container.NewVBox() // 存放按钮
	a := widget.NewAccordion()
	ti := widget.NewAccordionItem(title, vbox)
	a.Append(ti)
	a.MultiOpen = true
	a.OpenAll()
	name := widget.NewEntry()
	addbutton := widget.NewButtonWithIcon("添加按钮", theme.ContentAddIcon(), func() {
		p := Popup("添加按钮", "按钮名称:", name, func() {
			b := NewRightClickButton(vbox, title, name.Text)
			viper.SetConfigFile("./config.json")
			if err := viper.ReadInConfig(); err != nil {
				fmt.Println(err)
			}
			m := viper.GetStringMap(title)
			checkmodle = false
			for v := range m {
				if v == name.Text {
					checkmodle = true
					break
				}
			}
			if checkmodle {
				dialog.ShowInformation("提示", "请勿添加重名按钮", fyne.CurrentApp().Driver().AllWindows()[0])
			} else {
				vbox.Add(b)
			}
			name.Text = ""
			name.Refresh()
		})
		p.Resize(fyne.NewSize(200, 100))
		p.Show()
	})
	delbutton := widget.NewButtonWithIcon("删除分页", theme.HistoryIcon(), func() {
		for i, v := range tabs.Items {
			if title == v.Text {
				tabs.RemoveIndex(i)
				viper.SetConfigFile("./config.json")
				viper.Set(title, nil)
				viper.WriteConfigAs("./config.json")
			}
		}
	})
	close := widget.NewButtonWithIcon("", theme.StorageIcon(), func() {
		if a.Items[0].Open {
			a.CloseAll()
		} else {
			a.OpenAll()
		}
	})
	c = container.NewBorder(nil, container.NewGridWithColumns(3, addbutton, delbutton, close), nil, nil, a)
	return c, vbox
}
