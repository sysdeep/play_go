package containers_page

import (
	"fmt"
	"hdu/internal/services"
	"sort"

	"github.com/visualfc/atk/tk"
)

// type onContainerSelectedHandler = func(model *services.ContainerListModel)

// область со списком контейнеров
type ContainersFrame struct {
	*tk.Frame
	tree        *tk.TreeView
	current_map map[string]services.ContainerListModel
	// on_container_selected_handler onContainerSelectedHandler
	vm *ContainersVM
}

func NewContainersFrame(parent tk.Widget, vm *ContainersVM) *ContainersFrame {

	fr := tk.NewFrame(parent)

	// tree
	tree := tk.NewTreeView(fr)
	labels := []string{"state", "name", "image"}
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

	current_map := make(map[string]services.ContainerListModel)
	cf := &ContainersFrame{
		Frame:       fr,
		tree:        tree,
		current_map: current_map,
		vm:          vm,
		// on_container_selected_handler: nil,
	}

	vm.register(cf)

	// events
	tree.BindEvent("<Double-1>", func(e *tk.Event) {
		items := tree.SelectionList()
		// fmt.Println(items)
		if len(items) > 0 {
			cf.onSelected(items[0])
		}
	})

	tree.BindEvent("<Return>", func(e *tk.Event) {
		items := tree.SelectionList()
		// fmt.Println(items)
		if len(items) > 0 {
			cf.onSelected(items[0])
		}
	})

	return cf
}

// func (cf *ContainersFrame) SetItems(items []services.ContainerListModel) {
//
// 	// clear all
// 	cf.tree.DeleteAllItems()
// 	for k := range cf.current_map {
// 		delete(cf.current_map, k)
// 	}
//
// 	// fill
// 	root := cf.tree.RootItem()
// 	for i, item := range items {
// 		row := root.InsertItem(i*10, item.State, []string{item.Name, item.Image})
// 		fmt.Println(row.Id())
// 		cf.current_map[row.Id()] = item
// 	}
// }

func (cf *ContainersFrame) onSelected(item *tk.TreeItem) {
	// fmt.Println(item)
	// fmt.Println(item.Id(), item.Index(), item.Values())

	model := cf.current_map[item.Id()]
	// utils.PrintAsJson(model)

	// if cf.on_container_selected_handler != nil {
	// 	cf.on_container_selected_handler(&model)
	// }

	cf.vm.show_container(model)
}

// func (cf *ContainersFrame) ConnectContainerSelected(handler onContainerSelectedHandler) {
// 	cf.on_container_selected_handler = handler
// }

func (cf *ContainersFrame) update(ev string) {
	switch ev {
	case EV_REFRESH:
		cf.refresh()
		break

	}
}

func (cf *ContainersFrame) refresh() {

	items := cf.vm.get_containers()

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Name < items[j].Name
	})

	// clear all
	cf.tree.DeleteAllItems()
	for k := range cf.current_map {
		delete(cf.current_map, k)
	}

	// fill
	root := cf.tree.RootItem()
	for i, item := range items {
		row := root.InsertItem(i*10, item.State, []string{item.Name, item.Image})
		fmt.Println(row.Id())
		cf.current_map[row.Id()] = item
	}

}
