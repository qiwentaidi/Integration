package plugins

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// 升级版的TreeWithString 多了个右键功能
func NewHostNode(data map[string][]string) (t *widget.Tree) {
	t = &widget.Tree{
		ChildUIDs: func(uid string) (c []string) {
			c = data[uid]
			return
		},
		IsBranch: func(uid string) (b bool) {
			_, b = data[uid]
			return
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return NewSuperLabel("Template Object")
		},
		UpdateNode: func(uid string, branch bool, node fyne.CanvasObject) {
			node.(*SuperLabel).SetText(uid)
		},
	}
	t.ExtendBaseWidget(t)
	return t
}
