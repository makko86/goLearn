package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	sl := make([][]uint8, dy)
	for i := 0; i < len(sl); i++ {
		sl[i] = make([]uint8, dx)
		for j := 0; j < len(sl[i]); j++ {
			sl[i][j] = uint8((i + j) / 2)
		}
	}
	return sl
}

func main() {
	pic.Show(Pic)
}
