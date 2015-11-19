package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct{
	w int
	h int
	v uint8
}

func(i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func(i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func(i Image) At(x, y int) color.Color {
	return color.RGBA{i.v, i.v / 2, 40, 255}
}

func main() {
	//m := Image{w: 65, h: 65, v:100}
	//pic.ShowImage(m)
}
