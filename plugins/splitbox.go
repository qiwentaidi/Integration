package plugins

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	appContainerMap = map[string]fyne.CanvasObject{}
)

func SplitBox() (sbox *container.Split) {
	T = NewHostNode(InitFirstNode())
	// 主节点被选中时,响应的事件,UID为节点名称
	sbox = container.NewHSplit(T, widget.NewLabel(""))
	T.OnSelected = func(uid widget.TreeNodeID) {
		c, ok := appContainerMap[uid] // 如果不存在appContainerMap[uid]，则创建一个c容器将其赋值给appContainerMap[uid]
		if !ok {
			e, docstab := NewTertiary(uid)
			secoundnode := InitTrailingNode(uid)
			for _, scnode := range secoundnode {
				docstab.Append(scnode)
			}
			c = container.NewBorder(nil, e, nil, nil, docstab)
			appContainerMap[uid] = c
		}
		sbox.Trailing = appContainerMap[uid] // 把c容器赋值到sbox的右边
		sbox.Refresh()
	}
	sbox.Offset = 0
	return sbox
}
