package main

import (
	"egame/internal/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	resource "github.com/quasilyte/ebitengine-resource"
)

func main() {
	g := &myGame{
		windowWidth:  320,
		windowHeight: 240,
		loader:       createLoader(),
	}

	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle("Ebitengine Quest")

	assets.RegisterResources(g.loader)

	g.init()

	// RunGame ожидает реализации трёх методов:
	// Update, Draw и Layout; они определены ниже.
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}

}

type myGame struct {
	windowWidth  int
	windowHeight int
	player       *Player
	loader       *resource.Loader
}

func (g *myGame) init() {
	gopher := g.loader.LoadImage(assets.ImageGopher).Data
	g.player = &Player{img: gopher}
}

func (g *myGame) Update() error {
	g.player.pos.X += 16 * (1.0 / 60.0)

	return nil

}

func (g *myGame) Draw(screen *ebiten.Image) {

	// gopher := g.loader.LoadImage(assets.ImageGopher).Data
	// var options ebiten.DrawImageOptions
	// screen.DrawImage(gopher, &options)

	var options ebiten.DrawImageOptions
	options.GeoM.Translate(g.player.pos.X, g.player.pos.Y)
	screen.DrawImage(g.player.img, &options)

	// ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *myGame) Layout(w, h int) (int, int) {
	// Layout - тема для продвинутых, поэтому нам пока
	// достаточно считать, что screen size = window size.
	return g.windowWidth, g.windowHeight
}

func createLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAsset
	return loader
}
