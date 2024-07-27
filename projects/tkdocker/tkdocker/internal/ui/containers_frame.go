package ui

import (
	"tkdocker/internal/services"

	"github.com/visualfc/atk/tk"
)

// область со списком контейнеров
type ContainersFrame struct {
	*tk.Frame
	tree *tk.TreeView
}

func NewContainersFrame(parent tk.Widget) *ContainersFrame {

	fr := tk.NewFrame(parent)

	lbl := tk.NewLabel(fr, "Containers list")

	// tree
	tree := tk.NewTreeView(fr)
	labels := []string{"state", "name"}
	tree.SetColumnCount(len(labels))
	tree.SetHeaderLabels(labels)

	// layout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(lbl)
	main_layout.AddWidget(tree)

	return &ContainersFrame{fr, tree}
}

func (cf *ContainersFrame) SetItems(items []services.ContainerListModel) {

	// clear all
	cf.tree.DeleteAllItems()

	// fill
	root := cf.tree.RootItem()
	for i, item := range items {
		root.InsertItem(i, item.State, []string{item.Name, "", ""})
	}
}
