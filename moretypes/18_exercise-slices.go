package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, 0, dy)
	for y := 0; y < dy; y++ {
		a := make([]uint8, 0, dx)
		for x := 0; x < dx; x++ {
			a = append(a, uint8((x + y) / 2))
			//a = append(a, uint8(x * y))
			//a = append(a, uint8(x ^ y))
		}
		img = append(img, a)
	}
	return img
}

func main() {
	pic.Show(Pic)
}
