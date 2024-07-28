package ui

import (
	"fmt"
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

	list.ConnectContainerSelected(page.OnContainerSelected)

	return page
}

func (cp *ContainersPage) refresh() {
	items, _ := cp.container_service.GetAll()

	cp.list_frame.SetItems(items)
}

func (cp *ContainersPage) OnContainerSelected(model *services.ContainerListModel) {
	fmt.Println(model)

	// TODO: in new type
	top := tk.NewWindow()

	view := NewContainerView(top, NewFakeContainerProvider())
	layout := tk.NewVPackLayout(top)
	layout.AddWidget(view, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))

	top.SetTitle("Container view")
	top.ShowNormal()
}

// TODO: only for tests
type FakeContainerProvider struct{}

func NewFakeContainerProvider() *FakeContainerProvider {
	return &FakeContainerProvider{}
}

func (cp *FakeContainerProvider) GetContainer() (services.ContainerListModel, error) {

	addrs := make([]string, 0)
	ports := make([]string, 0)

	return services.ContainerListModel{
		ID:          "string",
		Name:        "string",
		Image:       "string",
		State:       "string",
		CreatedStr:  "string",
		IPAddresses: addrs,
		Ports:       ports,
	}, nil
}
