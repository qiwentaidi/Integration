package plugins

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type GetParentButton struct {
	widget.Button
	ContainerName string // 父二级节点的名称
	NodeName      string // 父主节点的名称
}

func NewGetParentButton(nodename, containername string, tapped func()) *GetParentButton {
	gpb := &GetParentButton{
		NodeName:      nodename,
		ContainerName: containername,
	}
	gpb.ExtendBaseWidget(gpb)
	gpb.Icon = theme.HomeIcon()
	gpb.SetText("添加按钮")
	// 实现gpb的左键功能
	gpb.OnTapped = tapped
	return gpb
}
