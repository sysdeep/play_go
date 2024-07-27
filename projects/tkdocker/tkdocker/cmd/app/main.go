package main

import (
	"fmt"
	"tkdocker/internal/services"
	"tkdocker/internal/ui"

	"github.com/docker/docker/client"
	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

func main() {

	// docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// init services
	servs := services.NewServices(cli)

	// start mainloop
	tk.MainLoop(func() {
		mw := ui.NewMainWindow(servs)
		mw.ResizeN(800, 600)
		// mw := NewWindow()
		mw.SetTitle("ATK Sample")
		mw.Center(nil)
		mw.ShowNormal()
		fmt.Println(tk.DumpWidget(mw))
	})

}
