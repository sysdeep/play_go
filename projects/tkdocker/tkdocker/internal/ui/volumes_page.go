package ui

import (
	"tkdocker/internal/services"

	"github.com/visualfc/atk/tk"
)

type VolumesPage struct {
	*tk.Frame
	// images_service *services.ImagesService
	// list_frame     *ImagesFrame
}

func NewVolumesPage(parent tk.Widget, images_service *services.ImagesService) *VolumesPage {

	fr := tk.NewFrame(parent)

	// label
	lbl := tk.NewLabel(fr, "Volumes")

	// list
	list := NewImagesFrame(fr)

	// refresh button
	refresh_button := tk.NewButton(fr, "Refresh")

	// lauout
	main_layout := tk.NewVPackLayout(fr)
	main_layout.AddWidget(lbl)
	main_layout.AddWidget(list,
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
	)
	main_layout.AddWidget(refresh_button)

	// fr.BindEvent("<Enter>", func(e *tk.Event) {
	// 	fmt.Println("Enter")
	// })
	//
	// fr.BindEvent("<Activate>", func(e *tk.Event) {
	// 	fmt.Println("Activate")
	// })

	page := &VolumesPage{
		fr,
		// images_service,
		// list,
	}

	// bind events --------------------------------------------------------------

	// событие отображения(при переключении вкладок)
	// fr.BindEvent("<Visibility>", func(e *tk.Event) {
	// 	page.refresh()
	// })
	//
	// refresh_button.OnCommand(func() {
	// 	page.refresh()
	// })

	return page
}

func (cp *VolumesPage) refresh() {
	// items, _ := cp.images_service.GetAll()
	//
	// cp.list_frame.SetItems(items)
}
