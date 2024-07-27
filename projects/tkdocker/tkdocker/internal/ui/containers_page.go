package ui

import (
	"tkdocker/internal/services"

	"github.com/visualfc/atk/tk"
)

type ContainersPage struct {
	*tk.Frame
	container_service *services.ContainersService
	list_frame        *ContainersFrame
}

func NewContainersPage(parent tk.Widget, containers_service *services.ContainersService) *ContainersPage {

	fr := tk.NewFrame(parent)

	// label
	lbl := tk.NewLabel(fr, "Containers")

	// list
	list := NewContainersFrame(fr)

	// controls bar
	controls := tk.NewFrame(fr)
	controls_layout := tk.NewHPackLayout(controls)
	// refresh button
	refresh_button := tk.NewButton(controls, "Refresh")
	controls_layout.AddWidget(tk.NewLayoutSpacer(controls, 1, true))
	controls_layout.AddWidget(refresh_button, tk.PackAttrSideRight())

	// layout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(lbl)
	main_layout.AddWidget(list,
		tk.PackAttrFillX(),
	)
	main_layout.AddWidget(controls, tk.PackAttrFillX(), tk.PackAttrPadx(4), tk.PackAttrPady(4))

	// fr.BindEvent("<Enter>", func(e *tk.Event) {
	// 	fmt.Println("Enter")
	// })
	//
	// fr.BindEvent("<Activate>", func(e *tk.Event) {
	// 	fmt.Println("Activate")
	// })

	page := &ContainersPage{
		fr,
		containers_service,
		list,
	}

	// bind events --------------------------------------------------------------

	// событие отображения(при переключении вкладок)
	fr.BindEvent("<Visibility>", func(e *tk.Event) {
		page.refresh()
	})

	refresh_button.OnCommand(func() {
		page.refresh()
	})

	return page
}

func (cp *ContainersPage) refresh() {
	items, _ := cp.container_service.GetAll()

	cp.list_frame.SetItems(items)
}
