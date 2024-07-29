package ui

import (
	"fmt"
	"tkdocker/internal/services"
	"tkdocker/internal/utils"

	"github.com/visualfc/atk/tk"
)

// область со списком контейнеров
type ImagesFrame struct {
	*tk.Frame
	tree        *tk.TreeView
	current_map map[string]services.ImageListModel
}

func NewImagesFrame(parent tk.Widget) *ImagesFrame {

	fr := tk.NewFrame(parent)

	// tree
	tree := tk.NewTreeView(fr)
	labels := []string{"tag", "ID", "Created", "Size"}
	tree.SetColumnCount(len(labels))
	tree.SetHeaderLabels(labels)

	// layout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(tree,
		tk.PackAttrFillBoth(),
		tk.PackAttrPadx(4),
		tk.PackAttrPady(4),
		tk.PackAttrExpand(true),
	)

	// instance
	current_map := make(map[string]services.ImageListModel)
	cf := &ImagesFrame{fr, tree, current_map}

	// events
	tree.BindEvent("<Double-1>", func(e *tk.Event) {
		items := tree.SelectionList()
		if len(items) > 0 {
			cf.onSelected(items[0])
		}
	})

	tree.BindEvent("<Return>", func(e *tk.Event) {
		items := tree.SelectionList()
		if len(items) > 0 {
			cf.onSelected(items[0])
		}
	})

	return cf

}

func (cf *ImagesFrame) SetItems(items []services.ImageListModel) {

	// clear all
	cf.tree.DeleteAllItems()
	for k := range cf.current_map {
		delete(cf.current_map, k)
	}

	// fill
	root := cf.tree.RootItem()
	for i, item := range items {

		for j, tag := range item.RepoTags {

			// TODO: size view

			row_id := i*100 + j
			row := root.InsertItem(row_id, tag, []string{item.ID, item.Created, fmt.Sprintf("%d", item.Size)})
			fmt.Println(row.Id())
			cf.current_map[row.Id()] = item

		}
	}
}

func (cf *ImagesFrame) onSelected(item *tk.TreeItem) {

	utils.PrintAsJson(cf.current_map)

	fmt.Println(item.Id())
	model := cf.current_map[item.Id()]
	fmt.Println(model)

	// if cf.on_container_selected_handler != nil {
	// 	cf.on_container_selected_handler(model)
	// }
}
