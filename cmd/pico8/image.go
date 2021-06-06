package pico8

import (
	"bytes"
	"fmt"
	"image"
	"strings"
)

func ImageToBytes(img *image.Paletted) []byte {

	var bytes bytes.Buffer

	stride := img.Bounds().Dx() / 2
	height := img.Bounds().Dy()

	for y := 0; y < height; y += 1 {
		for x := 0; x < stride; x += 1 {
			c1 := img.ColorIndexAt(x*2, y)
			c2 := img.ColorIndexAt(x*2+1, y)
			bytes.WriteByte(byte(c1<<4 | c2))
		}
	}

	return bytes.Bytes()
}

func ImageToHexLines(img *image.Paletted) <-chan string {

	ch := make(chan string)
	go func() {

		stride := img.Bounds().Dx() / 2
		height := img.Bounds().Dy()

		for y := 0; y < height; y += 1 {
			var line strings.Builder
			for x := 0; x < stride; x += 1 {
				c1 := img.ColorIndexAt(x*2, y)
				c2 := img.ColorIndexAt(x*2+1, y)
				line.WriteString(fmt.Sprintf("%x%x", c1, c2))
			}
			ch <- line.String()
		}
	}()
	return ch
}
