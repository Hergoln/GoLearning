package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	data [][]uint8
	rect image.Rectangle
}

// image 256x256
func NewStaticFuncImage() *Image {
	img := Image{make([][]uint8, 256),
		image.Rect(0, 0, 255, 255)}
	// alternate
	// rect := Rectangle{0, 0, 255, 255}
	// data := make([][]uint8, 256)
	// img := Image{data, rect}

	for d := range img.data {
		img.data[d] = make([]uint8, 256)
	}

	for y := 0; y < 256; y++ {
		for x := 0; x < 256; x++ {
			img.data[y][x] = uint8(x * y)
		}
	}

	return &img
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return img.rect
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{img.data[x][y], img.data[x][y], 255, 255}
}

func main() {
	m := NewStaticFuncImage()
	pic.ShowImage(m)
}
