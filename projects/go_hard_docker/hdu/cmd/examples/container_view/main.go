package main

import (
	"tkdocker/internal/services"
	"tkdocker/internal/ui"

	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

func main() {

	// init services

	// start mainloop
	tk.MainLoop(func() {
		mw := tk.RootWindow()

		view := ui.NewContainerView(mw, NewFakeContainerProvider())

		tk.NewVPackLayout(mw).AddWidget(view, tk.PackAttrFillBoth(), tk.PackAttrExpand(true))

		mw.ResizeN(800, 600)
		// mw := NewWindow()
		mw.SetTitle("ATK Sample")
		mw.Center(nil)
		mw.ShowNormal()
		// fmt.Println(tk.DumpWidget(mw))
	})

}

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
