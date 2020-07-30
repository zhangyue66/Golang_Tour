package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	yz := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		yz[i] = make([]uint8, dx)
	}

	for y := range yz {
		for x := range yz[y] {
			yz[y][x] = uint8((y + x) / 2)
		}
	}

	return yz
}

func main() {
	pic.Show(Pic)
}
