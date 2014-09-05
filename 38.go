package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var thing = make([][]uint8, dy)

	for i := 0; i < dy; i++ {

		var that = make([]uint8, dx)

		for j := 0; j < dx; j++ {
			that[j] = uint8(i * j)
		}

		thing[i] = that
	}

	return thing
}

func main() {
	pic.Show(Pic)
}
