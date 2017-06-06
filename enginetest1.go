package main

import (
	"fmt"
	"github.com/jamesbyrnes/grogue/engine"
)

func main() {
	dung := new(engine.Dungeon)
	dung.Initialize(1)

	dung.Build(0, 2)

	for _, a := range dung.Levels[0].Tiles {
		for _, b := range a {
			fmt.Print(fmt.Sprintf("%s", b.GetChar()))
		}
		fmt.Print("\n")
	}
}
