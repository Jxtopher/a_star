package main

import (
	// "fmt"

	"github.com/jxtopher/a_star/aStar"
	"github.com/jxtopher/a_star/worldGenerator"
	// "log"
)

func main() {
	var x_size uint64 = 50
	var y_size uint64 = 50
	world := make([][]uint8, x_size)
	for i := range world {
		world[i] = make([]uint8, y_size)
	}

	w := worldGenerator.World{x_size, y_size, world}
	w.Init(0.0)

	var path []worldGenerator.Coordinate
	if w.Ground[0][0] == worldGenerator.Empty && w.Ground[48][48] == worldGenerator.Empty {
		path = aStar.Run(w, worldGenerator.Coordinate{0, 0}, worldGenerator.Coordinate{48, 48})
	}

	aStar.Plot(w, path)
	//
}
