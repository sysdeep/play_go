package gui

import "github.com/gizak/termui/v3/widgets"

type ContainersFrame struct {
	List *widgets.List
}

func NewContainersFrame() *ContainersFrame {

	list := widgets.NewList()
	list.SetRect(20, 20, 60, 30)
	list.Border = true
	list.Title = " Containers "

	list.Rows = append(list.Rows, "first", "second")

	return &ContainersFrame{
		List: list,
	}
}

// func(cf *ContainersFrame) GetWidget()
