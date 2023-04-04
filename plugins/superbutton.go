package plugins

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// 继承button的属性
type SuperButton struct {
	widget.Button
	Parent        *fyne.Container // 获取到父类容器对象
	NodeName      string
	ContainerName string
}

// 实现TappedSecondary(*fyne.PointEvent)方法即可完成右键响应
func (sb *SuperButton) TappedSecondary(ev *fyne.PointEvent) {
	editor := widget.NewButton("编辑", nil)
	remove := widget.NewButton("删除", nil)
	open := widget.NewButton("打开文件目录", nil)
	// 右键弹出选项
	rightclicksubmenu := widget.NewPopUp(container.NewVBox(editor, remove, open), fyne.CurrentApp().Driver().AllWindows()[0].Canvas())
	editor.OnTapped = func() { // 1.配置路径与启动方式
		rightclicksubmenu.Hide()
		ep := editorConfig(sb)
		ep.Show()
	}
	remove.OnTapped = func() { // 2.删除按钮
		rightclicksubmenu.Hide()
		sb.Parent.Remove(sb)
		data := GetConfig()
		delete(data[sb.NodeName].(map[string]interface{})[sb.ContainerName].(map[string]interface{}), sb.Text)
		SaveConfig(data)
	}
	open.OnTapped = func() { // 3.打开文件配置路径所在目录
		rightclicksubmenu.Hide()
		data := GetConfig()
		path := data[sb.NodeName].(map[string]interface{})[sb.ContainerName].(map[string]interface{})[sb.Text].(map[string]interface{})["path"].(string)
		OpenFolder(path)
	}
	rightclicksubmenu.ShowAtPosition(ev.AbsolutePosition)
}

// 增加New的方法
func NewSuperButton(label string, c *fyne.Container, nodename, containername string) *SuperButton {
	sb := &SuperButton{
		Parent:        c,
		ContainerName: containername,
		NodeName:      nodename,
	}
	sb.ExtendBaseWidget(sb)
	sb.SetText(label)
	// 实现SuperButton的左键功能
	sb.OnTapped = func() {
		data := GetConfig()
		// 读取启动方法和目录
		method := data[sb.NodeName].(map[string]interface{})[sb.ContainerName].(map[string]interface{})[sb.Text].(map[string]interface{})["method"].(string)
		path := data[sb.NodeName].(map[string]interface{})[sb.ContainerName].(map[string]interface{})[sb.Text].(map[string]interface{})["path"].(string)
		CallApp(method, path)
	}
	return sb
}

// 右键-打开文件目录功能
func OpenFolder(path string) {
	info, err := os.Stat(path)
	if err != nil {
		dialog.ShowInformation("提示", "未配置目录文件或路径错误", fyne.CurrentApp().Driver().AllWindows()[0])
	} else {
		if info.IsDir() {
			c := exec.Command("explorer", path)
			c.Start()
		} else {
			dir := filepath.Dir(path)
			c := exec.Command("explorer", dir)
			c.Start()
		}
	}
}

// 右键-编辑功能
func editorConfig(sb *SuperButton) *widget.PopUp {
	w := fyne.CurrentApp().Driver().AllWindows()[0]
	name := widget.NewEntry()
	name.SetText(sb.Text)
	setmethod := widget.NewSelect([]string{"cmd", "exe", "java", "bat"}, nil)
	setpath := widget.NewMultiLineEntry()
	setpath.PlaceHolder = "你可以选择粘贴文件路径或者点击右侧按钮选择文件路径"
	setpath.Wrapping = fyne.TextWrapBreak
	// 选择目录控件
	setpath.ActionItem = widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
		dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
			if uc != nil {
				setpath.SetText(strings.ReplaceAll(uc.URI().Path(), "/", "\\")) // 获取的路径为/，将其替换成\\
			}
		}, w)
	})
	data := GetConfig()
	// 读取启动方法和目录
	method := data[sb.NodeName].(map[string]interface{})[sb.ContainerName].(map[string]interface{})[sb.Text].(map[string]interface{})["method"].(string)
	path := data[sb.NodeName].(map[string]interface{})[sb.ContainerName].(map[string]interface{})[sb.Text].(map[string]interface{})["path"].(string)
	setmethod.SetSelected(method)
	setpath.SetText(path)
	form := widget.NewForm(
		widget.NewFormItem("名称:", name),
		widget.NewFormItem("启动方法:", setmethod),
		widget.NewFormItem("路径:", setpath),
	)
	save := widget.NewButtonWithIcon("保存", theme.ConfirmIcon(), nil)
	cancel := widget.NewButtonWithIcon("退出", theme.CancelIcon(), nil)
	subheader := container.NewVBox(widget.NewLabelWithStyle("编辑", fyne.TextAlignCenter, fyne.TextStyle{}), form, container.NewHBox(layout.NewSpacer(), save, cancel, layout.NewSpacer()))
	subframe := container.NewBorder(widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), subheader)
	ep := widget.NewModalPopUp(subframe, w.Canvas()) // 编辑功能控件
	ep.Resize(fyne.NewSize(400, 100))
	save.OnTapped = func() {
		delete(data[sb.NodeName].(map[string]interface{})[sb.ContainerName].(map[string]interface{}), sb.Text)
		data[sb.NodeName].(map[string]interface{})[sb.ContainerName].(map[string]interface{})[name.Text] = map[string]interface{}{
			"method": setmethod.Selected,
			"path":   setpath.Text,
		}
		SaveConfig(data)
		sb.SetText(name.Text)
		sb.Refresh()
		ep.Hide()
	}
	cancel.OnTapped = func() {
		ep.Hide()
	}
	return ep
}

func CallApp(method, path string) {
	switch method {
	case "cmd":
		info, _ := os.Stat(path)
		go func() {
			if info.IsDir() {
				c := exec.Command("cmd", "/k", "start", "cd", "/d", path)
				c.Start()
			} else {
				dir := filepath.Dir(path)
				c := exec.Command("cmd", "/k", "start", "cd", "/d", dir)
				c.Start()
			}
		}()
	case "exe":
		go func() {
			// 启动外部exe文件
			cmd := exec.Command(path)
			cmd.Start()
		}()
	case "java":
		go func() {
			// 启动外部java文件
			cmd := exec.Command("java", "-jar", path)
			cmd.Start()
		}()
	case "bat":
		go func() {
			// 启动外部bat文件
			cmd := exec.Command("cmd", "/C", path)
			cmd.CombinedOutput()
		}()
	}
}
