package main

import (
	"fmt"
	"image"
)

// interface image from "image package"
// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	//color.RGBA is interface defined in image/color
	fmt.Println(m.At(0, 0).RGBA())
}
