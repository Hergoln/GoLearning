package main

import ("golang.org/x/tour/pic")

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for index := range pic {
		pic[index] = make([]uint8, dx)
	}
	
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			pic[y][x] = uint8(x * y)
		}
	}
	
	return pic
}

func main() {
	pic.Show(Pic)
}