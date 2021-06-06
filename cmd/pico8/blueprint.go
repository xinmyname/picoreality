package pico8

import (
	"encoding/json"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
)

func LoadBlueprint(path string) blueprint {
	var bp blueprint

	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	if err = json.NewDecoder(f).Decode(&bp); err != nil {
		log.Fatal(err.Error())
	}

	return bp
}

type blueprint struct {
	TitlePath   string
	MainPath    string
	MonsterPath string
}

func (bp blueprint) LoadMainLuaText() string {
	text, err := ioutil.ReadFile(bp.MainPath)

	if err != nil {
		log.Fatal(err.Error())
	}

	return string(text)
}

func (bp blueprint) LoadTitleImage() image.Image {

	f, err := os.Open(bp.TitlePath)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	img, _, err := image.Decode(f)

	if err != nil {
		log.Fatal(err.Error())
	}

	return img
}

func (bp blueprint) LoadMonsterImage() image.Image {
	f, err := os.Open(bp.MonsterPath)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	img, _, err := image.Decode(f)

	if err != nil {
		log.Fatal(err.Error())
	}

	return img
}
