package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (Image) ColorModel() color.Model {

	return color.RGBAModel
}

func (Image) Bounds() image.Rectangle {
	w, h := 20, 20

	return image.Rect(0, 0, w, h)
}

func (Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
