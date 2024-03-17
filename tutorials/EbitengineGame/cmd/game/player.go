package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gmath"
)

type Player struct {
	pos gmath.Vec // {X, Y}
	img *ebiten.Image
}
