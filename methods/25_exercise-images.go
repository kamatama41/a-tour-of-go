package main

import (
	"image/color"
	"image"
	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBA64Model
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (i Image) At(x, y int) color.Color {
	//rg := uint8((x + y) / 2)
	//rg := uint8(x * y)
	rg := uint8(x ^ y)
	return color.RGBA{rg, rg, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
