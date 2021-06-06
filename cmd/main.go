package main

import (
	"cartgen/pico8"
	"fmt"
	"image"
	"image/draw"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	if len(os.Args) < 2 {

		fmt.Printf("Usage: %s <blueprint file> \n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	bp := pico8.LoadBlueprint(os.Args[1])
	mainLuaText := bp.LoadMainLuaText()

	titleImage := bp.LoadTitleImage()
	dstTitlePal, dstTitlePicoPal := pico8.OptimumPalette(titleImage)
	dstTitleImage := image.NewPaletted(titleImage.Bounds(), dstTitlePal)
	draw.FloydSteinberg.Draw(dstTitleImage, titleImage.Bounds(), titleImage, image.Point{})

	monsterImage := bp.LoadMonsterImage()
	dstMonsterPal, dstMonsterPicoPal := pico8.CustomPalette([]int{0, 1, 130, 131, 4, 5, 134, 128, 129, 132, 133, 141, 140, 13, 143, 15})
	dstMonsterImage := image.NewPaletted(monsterImage.Bounds(), dstMonsterPal)
	draw.FloydSteinberg.Draw(dstMonsterImage, monsterImage.Bounds(), monsterImage, image.Point{})

	fmt.Println("pico-8 cartridge // http://www.pico-8.com")
	fmt.Println("version 32")
	fmt.Println("__lua__")
	fmt.Println(mainLuaText)
	fmt.Println("-->8")
	fmt.Println("data={}")

	fmt.Printf("data.title_pal={%s}\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dstTitlePicoPal)), ","), "[]"))
	fmt.Printf("data.monster_pal={%s}\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dstMonsterPicoPal)), ","), "[]"))

	fmt.Printf("data.monster_image=\"")
	monsterBytes := pico8.ImageToBytes(dstMonsterImage)
	pico8.WriteBytesAsText(monsterBytes, os.Stdout)
	fmt.Printf("\"\n")

	fmt.Println("__gfx__")

	for line := range pico8.ImageToHexLines(dstTitleImage) {
		fmt.Println(line)
	}

	fmt.Println("__map__")
	fmt.Println("000102030405060708090a0b0c0d0e0f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("101112131415161718191a1b1c1d1e1f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("202122232425262728292a2b2c2d2e2f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("303132333435363738393a3b3c3d3e3f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("404142434445464748494a4b4c4d4e4f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("505152535455565758595a5b5c5d5e5f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("606162636465666768696a6b6c6d6e6f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("707172737475767778797a7b7c7d7e7f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("808182838485868788898a8b8c8d8e8f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("909192939495969798999a9b9c9d9e9f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("a0a1a2a3a4a5a6a7a8a9aaabacadaeaf00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("b0b1b2b3b4b5b6b7b8b9babbbcbdbebf00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("c0c1c2c3c4c5c6c7c8c9cacbcccdcecf00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("d0d1d2d3d4d5d6d7d8d9dadbdcdddedf00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("e0e1e2e3e4e5e6e7e8e9eaebecedeeef00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	fmt.Println("f0f1f2f3f4f5f6f7f8f9fafbfcfdfeff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
}
