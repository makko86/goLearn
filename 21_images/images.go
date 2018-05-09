package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type MyImage struct {
	W, H int
}

func (m *MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (m *MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.W, m.H)
}

//shamelessly copied
func (m *MyImage) At(x, y int) color.Color {
	return color.RGBA{uint8(x) ^ uint8(y), uint8(x) ^ uint8(y), 255, 255}
}

func main() {
	m := MyImage{50, 50}
	pic.ShowImage(&m)
}
