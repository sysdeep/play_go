package ui

import (
	"tkdocker/internal/services"

	"github.com/visualfc/atk/tk"
)

// область со списком контейнеров
type ImagesFrame struct {
	*tk.Frame
	tree *tk.TreeView
}

func NewImagesFrame(parent tk.Widget) *ImagesFrame {

	fr := tk.NewFrame(parent)

	lbl := tk.NewLabel(fr, "Images list")

	// tree
	tree := tk.NewTreeView(fr)
	labels := []string{"ID", "-", "-"}
	tree.SetColumnCount(len(labels))
	tree.SetHeaderLabels(labels)

	// layout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(lbl)
	main_layout.AddWidget(tree)

	return &ImagesFrame{fr, tree}
}

func (cf *ImagesFrame) SetItems(items []services.ImageListModel) {

	// clear all
	cf.tree.DeleteAllItems()

	// fill
	root := cf.tree.RootItem()
	for i, item := range items {
		root.InsertItem(i, item.ID, []string{"", ""})
	}
}
