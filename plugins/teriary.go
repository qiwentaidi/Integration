package plugins

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// 新建三级节点
func NewTertiary(node string) (*widget.Entry, *container.DocTabs) {
	config := GetConfig()
	c := container.NewDocTabs()
	tertiaryName := widget.NewEntry()
	tertiaryName.PlaceHolder = "在此处输入需要添加的二级节点名称,若不填写则右上角加号按钮无法添加节点"
	c.CreateTab = func() (tertiaryContainer *container.TabItem) {
		if tertiaryName.Text != "" {
			tertiaryContainer = container.NewTabItem(tertiaryName.Text, nil)
			b := NewGetParentButton(node, tertiaryContainer.Text, nil) // 每个容器需要一个初始化按钮,作用为添加新的按钮
			content := ContentTertiary(b)
			b.OnTapped = func() {
				p := AddContentPopUp(content, b.NodeName, b.ContainerName)
				p.Show()
			}
			tertiaryContainer.Content = content
			config[node].(map[string]interface{})[tertiaryContainer.Text] = make(map[string]interface{})
			SaveConfig(config)
			tertiaryName.Text = ""
			tertiaryName.Refresh()
			return tertiaryContainer
		} else {
			dialog.ShowInformation("提示", "请输入要添加的二级节点名称", fyne.CurrentApp().Driver().AllWindows()[0])
		}
		return
	}
	c.CloseIntercept = func(ti *container.TabItem) {
		delete(config[node].(map[string]interface{}), ti.Text)
		SaveConfig(config)
		c.Remove(ti)
		c.Refresh()
	}
	return tertiaryName, c
}

// 内容为一个5列每行的容器
func ContentTertiary(objects ...fyne.CanvasObject) *fyne.Container {
	c := container.NewGridWithColumns(5, objects...)
	return c
}

// nodename, containername string
func AddContentPopUp(c *fyne.Container, nodename, containername string) *widget.PopUp {
	w := fyne.CurrentApp().Driver().AllWindows()[0]
	title := widget.NewLabelWithStyle("请输入按钮名称:", fyne.TextAlignCenter, fyne.TextStyle{})
	buttonname := widget.NewEntry() // 要接收的名字
	p := widget.NewModalPopUp(nil, w.Canvas())
	save := widget.NewButton("保存", func() { // 添加按钮
		sb := NewSuperButton(buttonname.Text, c, nodename, containername)
		data := GetConfig()
		data[nodename].(map[string]interface{})[containername].(map[string]interface{})[buttonname.Text] = make(map[string]interface{})
		data[nodename].(map[string]interface{})[containername].(map[string]interface{})[buttonname.Text].(map[string]interface{})["method"] = ""
		data[nodename].(map[string]interface{})[containername].(map[string]interface{})[buttonname.Text].(map[string]interface{})["path"] = ""
		SaveConfig(data)
		c.Add(sb)
		p.Hide()
	})
	cancle := widget.NewButton("退出", func() {
		p.Hide()
	})
	p.Content = container.NewBorder(title, container.NewGridWithColumns(2, save, cancle), nil, nil, buttonname)
	p.Resize(fyne.NewSize(200, 100))
	return p
}
