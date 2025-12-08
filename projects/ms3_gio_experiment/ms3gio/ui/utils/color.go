package utils

import "image/color"

func Torgb(c uint32) color.NRGBA {
	return Toargb(0xff000000 | c)
}

func Toargb(c uint32) color.NRGBA {
	return color.NRGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}
