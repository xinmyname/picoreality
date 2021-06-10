package pico8

import (
	"fmt"
	"math"
)

func CalculateLensLUT() (xoffsets [][]int, yoffsets [][]int) {

	width := 48
	height := 48

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

	k1 := -0.251
	k2 := 0.0

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

			d := k1*math.Pow(s, 2) + k2*math.Pow(s, 4)

			r := math.Sqrt(s*s + t*t)

			u := 0
			v := 0

			if r <= 1.0 {
				u = int(s * d * cx * -1.0)
				v = int(t * d * cy * -1.0)
			}

			xoffsets[y][x] = u
			yoffsets[y][x] = v

			/*
				s := float64(x) - cx/fx
				t := float64(y) - cy/fy

				r := math.Sqrt(s*s + t*t)
				d := 1.0 + k1*math.Pow(r, 2) + k2*math.Pow(r, 4)

				u := (s*d+(2.0*p1*s*t+p2*(r*r+2.0*s*s)))*fx + cx
				v := (t*d+(p1*(r*r+2.0*t*t)+2.0*p2*s*t))*fy + cy

				str := fmt.Sprintf("%0.1f,%0.1f", u, v)
				fmt.Printf("%7s  ", str)
			*/
			/*
				xs := x - width/2
				if xs > 0 {
					xs += 1
				}

				ys := y - height/2
				if ys > 0 {
					ys += 1
				}

				u := float64(xs) / float64(width/2)
				v := float64(ys) / float64(height/2)

				r := math.Sqrt(u*u + v*v)

				nx := 0.0
				ny := 0.0

				if r <= 1.0 {

					if xs != 0 || ys != 0 {
						m := math.Pow(r/2.0, 4) * 16
						t := math.Atan(float64(ys) / float64(xs))

						nx = math.Abs(math.Cos(t) * m * 6)
						ny = math.Abs(math.Sin(t) * m * 6)

						if xs < 0 {
							nx *= -1
						}

						if ys < 0 {
							ny *= -1
						}

					}
				}

				s := fmt.Sprintf("%d,%d", int(nx), int(ny))
				fmt.Printf("%5s  ", s)
			*/
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	return
}

/*
  // Calculate l2 norm
  float r = x*x + y*y;

  // Calculate the deflated or inflated new coordinate (reverse transform)
  float x3 = x / (1.0 - alphax * r);
  float y3 = y / (1.0 - alphay * r);
  float x2 = x / (1.0 - alphax * (x3 * x3 + y3 * y3));
  float y2 = y / (1.0 - alphay * (x3 * x3 + y3 * y3));

  // Forward transform
  // float x2 = x * (1.0 - alphax * r);
  // float y2 = y * (1.0 - alphay * r);

  // De-normalize to the original range
  float i2 = (x2 + 1.0) * 1.0 / 2.0;
  float j2 = (y2 + 1.0) * 1.0 / 2.0;

*/
