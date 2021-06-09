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
	DataPath    string
	MonsterPath string
	AtomicPath  string
}

func (bp blueprint) LoadMainLuaText() string {
	text, err := ioutil.ReadFile(bp.MainPath)

	if err != nil {
		log.Fatal(err.Error())
	}

	return string(text)
}

func (bp blueprint) LoadDataLuaText() string {
	text, err := ioutil.ReadFile(bp.DataPath)

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

func (bp blueprint) LoadAtomicBytes() []byte {
	f, err := os.Open(bp.AtomicPath)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	bytes, err := ioutil.ReadAll(f)

	if err != nil {
		log.Fatal(err.Error())
	}

	return bytes
}
