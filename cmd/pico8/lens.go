package pico8

import (
	"math"
	"math/rand"
)

func CalculateLensLUT() (xoffsets [][]int, yoffsets [][]int) {

	width := 42
	height := 42

	xoffsets = make([][]int, height)
	for i := range xoffsets {
		xoffsets[i] = make([]int, width)
	}

	yoffsets = make([][]int, height)
	for i := range yoffsets {
		yoffsets[i] = make([]int, width)
	}

	cx := float64(width) / 2.0
	cy := float64(width) / 2.0

	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {

			xs := x - width/2
			if xs > 0 {
				xs += 1
			}

			ys := y - height/2
			if ys > 0 {
				ys += 1
			}

			s := float64(xs) / cx
			t := float64(ys) / cy

			r := math.Sqrt(s*s + t*t)
			m := 5.0 * math.Pow(r, 2)
			j := 0.0 //((rand.Float64() - 0.5) * 2.0) * 1
			if r < 0.7 {
				j = ((rand.Float64() - 0.5) * 2.0) * 1.5
			}

			u := 0
			v := 0

			if r <= 1.0 {
				u = int(r*m + j)

				if xs < 0 {
					u *= -1
				}

				v = int(r*m + j)

				if ys < 0 {
					v *= -1
				}
			}

			//fmt.Printf("%d ", v)

			xoffsets[y][x] = u
			yoffsets[y][x] = v
		}
		//fmt.Println()
	}

	return
}
