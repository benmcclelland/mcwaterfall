package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

//BuildWaterFalls builds n, s, e, w waterfalls
func BuildWaterFalls(basename string) error {
	origin := XYZ{X: 0, Y: 0, Z: -2}

	for _, direction := range []string{"north", "east", "south", "west"} {
		fname := path.Join(basename, "waterfall_"+direction)
		f, err := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("open %v: %v", fname, err)
		}
		defer f.Close()

		obj := NewMCObject(WithOrientation(direction))
		wf := CreateWaterfall(origin, obj)
		err = WriteBoxes(f, wf)
		if err != nil {
			return fmt.Errorf("build waterfall: %v", err)
		}
	}

	return nil
}

//BuildLavaFalls builds n, s, e, w lava falls
func BuildLavaFalls(basename string) error {
	origin := XYZ{X: 0, Y: 0, Z: -2}

	for _, direction := range []string{"north", "east", "south", "west"} {
		fname := path.Join(basename, "lavafall_"+direction)
		f, err := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("open %v: %v", fname, err)
		}
		defer f.Close()

		obj := NewMCObject(WithOrientation(direction), WithType("lavafall"))
		wf := CreateWaterfall(origin, obj)
		err = WriteBoxes(f, wf)
		if err != nil {
			return fmt.Errorf("build lavafall: %v", err)
		}
	}

	return nil
}

func main() {
	var mcSavesDir string
	var mcWorldFuncDir string

	flag.StringVar(&mcSavesDir, "s", "~", "Minecraft saves directory")
	flag.StringVar(&mcWorldFuncDir, "w", "mc", "Minecraft functions directory")
	flag.Parse()

	err := BuildWaterFalls(path.Join(mcSavesDir, mcWorldFuncDir))
	if err != nil {
		log.Fatalln(err)
	}
	err = BuildLavaFalls(path.Join(mcSavesDir, mcWorldFuncDir))
	if err != nil {
		log.Fatalln(err)
	}
}
