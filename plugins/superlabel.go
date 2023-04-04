package plugins

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SuperLabel struct {
	widget.Label
}

func (sl *SuperLabel) TappedSecondary(ev *fyne.PointEvent) {
	w := fyne.CurrentApp().Driver().AllWindows()[0]
	rename := widget.NewButton("重命名", nil)
	remove := widget.NewButton("删除", nil)
	// 右键弹出选项
	rightclicksubmenu := widget.NewPopUp(container.NewVBox(rename, remove), w.Canvas())
	// √按钮事件
	name := widget.NewEntry()
	p := widget.NewPopUp(name, w.Canvas())
	name.ActionItem = widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		appMap[""] = removeElement(appMap[""], sl.Text, name.Text, p)
		T.Refresh()
	})
	rename.OnTapped = func() { // 1.重命名按钮
		rightclicksubmenu.Hide()
		p.ShowAtPosition(ev.AbsolutePosition)
		p.Resize(fyne.NewSize(150, 30))
	}
	remove.OnTapped = func() { // 2.删除按钮
		data := GetConfig()
		appMap[""] = removeElement(appMap[""], sl.Text, "", p)
		T.Refresh()
		delete(data, sl.Text) // 删除config配置
		SaveConfig(data)
		rightclicksubmenu.Hide()
		delete(appContainerMap, sl.Text) // 删除appConatinerMap中的节点
		fmt.Printf("appContainerMap: %v\n", appContainerMap)
	}
	rightclicksubmenu.ShowAtPosition(ev.AbsolutePosition) // 面板出现在鼠标点击位置
}

func NewSuperLabel(text string) *SuperLabel {
	label := &SuperLabel{}
	label.Text = text
	label.ExtendBaseWidget(label)
	return label
}

func removeElement(arr []string, elem, newelem string, p *widget.PopUp) []string {
	var result []string
	for _, val := range arr {
		if val != elem {
			result = append(result, val)
		} else {
			// 增加判断newelem是否为空字符串且新名称不能与旧名称相同，如果不为空字符串则执行重命名功能
			if newelem != "" && elem != newelem {
				result = append(result, newelem)
				data := GetConfig()
				data[newelem] = data[elem] // 新建一个newelem=elem,然后删除elem实现重命名
				delete(data, elem)
				SaveConfig(data)
				p.Hide()
			}
		}
	}
	return result
}
