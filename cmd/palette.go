package main

import (
	"image"
	"image/color"
	"math"
	"sort"
)

type LabColor struct {
	L, a, b float64
}

func rgb2xyz(R uint32, G uint32, B uint32) (float64, float64, float64) {

	r := float64(R) / 65535.0
	g := float64(G) / 65535.0
	b := float64(B) / 65535.0

	if r > 0.04045 {
		r = math.Pow((r+0.055)/1.055, 2.4)
	} else {
		r = r / 12.92
	}

	if g > 0.04045 {
		g = math.Pow((g+0.055)/1.055, 2.4)
	} else {
		g = g / 12.92
	}

	if b > 0.04045 {
		b = math.Pow((b+0.055)/1.055, 2.4)
	} else {
		b = b / 12.92
	}

	r *= 100.0
	g *= 100.0
	b *= 100.0

	return (r*0.4124 + g*0.3576 + b*0.1805),
		(r*0.2126 + g*0.7152 + b*0.0722),
		(r*0.0193 + g*0.1192 + b*0.9505)
}

func xyz2Lab(X float64, Y float64, Z float64) (float64, float64, float64) {

	const REF_X = 95.047 // Observer= 2°, Illuminant= D65
	const REF_Y = 100.000
	const REF_Z = 108.883

	x := X / REF_X
	y := Y / REF_Y
	z := Z / REF_Z

	if x > 0.008856 {
		x = math.Pow(x, 1.0/3.0)
	} else {
		x = (7.787 * x) + (16.0 / 116.0)
	}

	if y > 0.008856 {
		y = math.Pow(y, 1.0/3.0)
	} else {
		y = (7.787 * y) + (16.0 / 116.0)
	}

	if z > 0.008856 {
		z = math.Pow(z, 1.0/3.0)
	} else {
		z = (7.787 * z) + (16.0 / 116.0)
	}

	return (116.0 * y) - 16.0,
		500.0 * (x - y),
		200.0 * (y - z)
}

func makeLabColor(c color.Color) LabColor {
	rc, gc, bc, _ := c.RGBA()
	x, y, z := rgb2xyz(rc, gc, bc)
	L, a, b := xyz2Lab(x, y, z)

	return LabColor{L, a, b}
}

func (c LabColor) ToRGB() color.Color {
	return color.RGBA{0, 0, 0, 0}
}

type LabPalette []LabColor

func makeLabPalette(pal []color.Color) LabPalette {
	var labpal []LabColor

	for _, rgb := range pal {
		labpal = append(labpal, makeLabColor(rgb))
	}

	return labpal
}

func LabDistance(a LabColor, b LabColor) float64 {
	return math.Sqrt(math.Pow(a.L-b.L, 2) + math.Pow(a.a-b.a, 2) + math.Pow(a.b-b.b, 2))
}

func (pal LabPalette) ClosestIndex(color color.Color) int {
	index := -1
	lastDist := math.MaxFloat64
	labColor := makeLabColor(color)

	for i, palColor := range pal {
		dist := LabDistance(palColor, labColor)
		if lastDist >= dist {
			lastDist = dist
			index = i
		}
	}

	return index
}

func OptimumPalette(img image.Image) (color.Palette, []int) {

	var rgbpal = color.Palette{
		color.RGBA{0, 0, 0, 0xff},       // 0 : black
		color.RGBA{29, 43, 83, 0xff},    // 1 : dark-blue
		color.RGBA{126, 37, 83, 0xff},   // 2 : dark-purple
		color.RGBA{0, 135, 81, 0xff},    // 3 : dark-green
		color.RGBA{171, 82, 54, 0xff},   // 4 : brown
		color.RGBA{95, 87, 79, 0xff},    // 5 : dark-grey
		color.RGBA{194, 195, 199, 0xff}, // 6 : light-grey
		color.RGBA{255, 241, 232, 0xff}, // 7 : white
		color.RGBA{255, 0, 77, 0xff},    // 8 : red
		color.RGBA{255, 163, 0, 0xff},   // 9 : orange
		color.RGBA{255, 236, 39, 0xff},  // 10 : yellow
		color.RGBA{0, 228, 54, 0xff},    // 11 : green
		color.RGBA{41, 173, 255, 0xff},  // 12 : blue
		color.RGBA{131, 118, 156, 0xff}, // 13 : lavender
		color.RGBA{255, 119, 168, 0xff}, // 14 : pink
		color.RGBA{255, 204, 170, 0xff}, // 15 : light-peach
		color.RGBA{41, 24, 20, 0xff},    // 128 : darkest-grey
		color.RGBA{17, 29, 53, 0xff},    // 129 : darker-blue
		color.RGBA{66, 33, 54, 0xff},    // 130 : darker-purple
		color.RGBA{18, 83, 89, 0xff},    // 131 : blue-green
		color.RGBA{116, 47, 41, 0xff},   // 132 : dark-brown
		color.RGBA{73, 51, 59, 0xff},    // 133 : darker-grey
		color.RGBA{162, 136, 121, 0xff}, // 134 : medium-grey
		color.RGBA{243, 239, 125, 0xff}, // 135 : light-yellow
		color.RGBA{190, 18, 80, 0xff},   // 136 : dark-red
		color.RGBA{255, 108, 36, 0xff},  // 137 : dark-orange
		color.RGBA{168, 231, 46, 0xff},  // 138 : light-green
		color.RGBA{0, 181, 67, 0xff},    // 139 : medium-green
		color.RGBA{6, 90, 181, 0xff},    // 140 : medium-blue
		color.RGBA{117, 70, 101, 0xff},  // 141 : mauve
		color.RGBA{255, 110, 89, 0xff},  // 142 : dark-peach
		color.RGBA{255, 157, 129, 0xff}, // 143 : peach
	}

	labpal := makeLabPalette(rgbpal)
	inventory := [32]int{}
	var optpal []color.Color

	for y := 0; y < img.Bounds().Dy(); y += 1 {
		for x := 0; x < img.Bounds().Dx(); x += 1 {
			i := labpal.ClosestIndex(img.At(x, y))
			inventory[i] += 1
		}
	}

	hist := make(map[int]int)
	indices := make([]int, 0, 32)

	for i := 0; i < 32; i += 1 {
		indices = append(indices, inventory[i])
		hist[inventory[i]] = i
	}

	pico8pal := make([]int, 0, 16)

	sort.Sort(sort.Reverse(sort.IntSlice(indices)))

	for i := 0; i < 16; i += 1 {
		str := indices[i]
		index := hist[str]
		optpal = append(optpal, rgbpal[index])

		val := index
		if index > 15 {
			val += 112
		}

		pico8pal = append(pico8pal, val)
	}

	return optpal, pico8pal
}
