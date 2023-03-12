package widget

import (
	"fmt"
	"integration/common"
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
	"github.com/spf13/viper"
)

// 继承button的属性
type RightClickButton struct {
	widget.Button
	ContainerName string
	Parent        *fyne.Container // 获取到父类容器对象
}

// 实现TappedSecondary(*fyne.PointEvent)方法即可完成右键响应
func (r *RightClickButton) TappedSecondary(ev *fyne.PointEvent) {
	v := viper.New()
	editor := widget.NewButton("编辑", nil)
	remove := widget.NewButton("删除", nil)
	open := widget.NewButton("打开文件目录", nil)
	// 右键弹出选项
	rightclicksubmenu := widget.NewPopUp(container.NewVBox(editor, remove, open), fyne.CurrentApp().Driver().AllWindows()[0].Canvas())
	editor.OnTapped = func() { // 1.配置路径与启动方式
		rightclicksubmenu.Hide()
		p := EditorConfig(v, r)
		p.Show()
	}
	remove.OnTapped = func() { // 2.删除按钮
		rightclicksubmenu.Hide()
		DeleteSelf(v, r)
	}
	open.OnTapped = func() { // 3.打开文件配置路径所在目录
		rightclicksubmenu.Hide()
		OpenFolder(v, r.Text, r.ContainerName)
	}
	rightclicksubmenu.ShowAtPosition(ev.AbsolutePosition)
	rightclicksubmenu.Show()
}

// 增加New的方法
func NewRightClickButton(container *fyne.Container, containername, label string) *RightClickButton {
	ret := &RightClickButton{
		ContainerName: containername,
		Parent:        container,
	}
	ret.ExtendBaseWidget(ret)
	ret.SetText(label)
	// 实现RightClickButton的左键功能
	ret.OnTapped = func() {
		viper.SetConfigFile("./config.json")
		err := viper.ReadInConfig()
		if err != nil {
			dialog.ShowInformation("", fmt.Sprintf("读取配置失败:%v", err), fyne.CurrentApp().Driver().AllWindows()[0])
		}
		// 读取开启方法
		method := viper.GetString(fmt.Sprintf("%s.%s.method", containername, label))
		// 读取目录
		path := viper.GetString(fmt.Sprintf("%s.%s.path", containername, label))
		CallApp(method, path)
	}
	return ret
}

// 右键-打开文件目录功能
func OpenFolder(v *viper.Viper, buttoname, containername string) {
	v.SetConfigFile("./config.json")
	if err := v.ReadInConfig(); err != nil {
		dialog.ShowInformation("", fmt.Sprintf("读取配置失败:%v", err), fyne.CurrentApp().Driver().AllWindows()[0])
	}
	path := v.GetString(fmt.Sprintf("%s.%s.path", containername, buttoname))
	info, err2 := os.Stat(path)
	if err2 != nil {
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

// 右键-删除功能
func DeleteSelf(v *viper.Viper, r *RightClickButton) {
	r.Parent.Remove(r)
	common.Editor(v, r.ContainerName, r.Text, "", "")
}

// 右键-编辑功能
func EditorConfig(v *viper.Viper, r *RightClickButton) *widget.PopUp {
	w := fyne.CurrentApp().Driver().AllWindows()[0]
	setmethod := widget.NewSelect([]string{"folder", "exe", "java", "bat"}, nil)
	setpath := widget.NewMultiLineEntry()
	setpath.PlaceHolder = "你可以选择粘贴文件路径或者点击右侧按钮选择文件路径"
	setpath.Wrapping = fyne.TextWrapBreak
	setpath.ActionItem = widget.NewButtonWithIcon("", theme.FolderOpenIcon(), func() {
		dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
			if uc != nil {
				setpath.SetText(strings.ReplaceAll(uc.URI().Path(), "/", "\\"))
			}
		}, w)
	})
	v.SetConfigFile("config.json")
	if err := v.ReadInConfig(); err != nil {
		dialog.ShowInformation("", fmt.Sprintf("读取配置失败:%v", err), fyne.CurrentApp().Driver().AllWindows()[0])
	}
	// 读取启动方法
	setmethod.SetSelected(v.GetString(fmt.Sprintf("%s.%s.method", r.ContainerName, r.Text)))
	// 读取目录
	path := v.GetString(fmt.Sprintf("%s.%s.path", r.ContainerName, r.Text))
	setpath.SetText(path)
	form := widget.NewForm(
		widget.NewFormItem("启动方法:", setmethod),
		widget.NewFormItem("路径:", setpath),
	)
	save := widget.NewButtonWithIcon("保存", theme.ConfirmIcon(), nil)
	cancel := widget.NewButtonWithIcon("退出", theme.CancelIcon(), nil)
	subheader := container.NewVBox(widget.NewLabelWithStyle("编辑", fyne.TextAlignCenter, fyne.TextStyle{}), form, container.NewHBox(layout.NewSpacer(), save, cancel, layout.NewSpacer()))
	subframe := container.NewBorder(widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), widget.NewSeparator(), subheader)
	editor_popup := widget.NewModalPopUp(subframe, w.Canvas()) // 编辑功能控件
	editor_popup.Resize(fyne.NewSize(400, 100))
	save.OnTapped = func() {
		fmt.Println("11")
		common.Editor(v, r.ContainerName, r.Text, setpath.Text, setmethod.Selected) // 写入配置文件
		editor_popup.Hide()
	}
	cancel.OnTapped = func() {
		editor_popup.Hide()
	}
	return editor_popup
}

func CallApp(method, path string) {
	switch method {
	case "folder":
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
