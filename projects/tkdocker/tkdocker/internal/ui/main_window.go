package ui

import (
	"tkdocker/internal/services"

	"github.com/visualfc/atk/tk"
)

type MainWindow struct {
	*tk.Window
	// containers_page *ContainersPage
}

func NewMainWindow(servs *services.Services) *MainWindow {

	mw := &MainWindow{
		tk.RootWindow(),
	}

	tabs := tk.NewNotebook(mw)

	containers_page := NewContainersPage(tabs, servs.Containers)
	tabs.AddTab(containers_page, "Containers")

	images_page := NewImagesPage(tabs, servs.Images)
	tabs.AddTab(images_page, "Images")

	volumes_page := NewVolumesPage(tabs, servs.Images)
	tabs.AddTab(volumes_page, "Volumes")

	test_label := tk.NewLabel(tabs, "test")
	tabs.AddTab(test_label, "Test")

	btn := tk.NewButton(mw, "Quit")
	btn.OnCommand(func() {
		tk.Quit()
	})

	layout := tk.NewVPackLayout(mw)
	layout.AddWidget(tabs,
		tk.PackAttrFillBoth(),
		tk.PackAttrExpand(true),
		tk.PackAttrPadx(4),
		tk.PackAttrPady(10))

	// layout.AddWidget(tk.NewLayoutSpacer(mw, 0, true))
	layout.AddWidget(btn, tk.PackAttrSideBottom())

	return mw

}
