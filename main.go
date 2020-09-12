package main

import (
	"fmt"

	"github.com/jxtopher/a_star/astar"
	"github.com/jxtopher/a_star/worldgenerator"
)

func main() {
	var xSize uint64 = 50
	var ySize uint64 = 50
	world := make([][]uint8, xSize)
	for i := range world {
		world[i] = make([]uint8, ySize)
	}

	w := worldgenerator.World{Xsize: xSize, Ysize: ySize, Ground: world}
	w.Init(0.0)

	var path []worldgenerator.Coordinate
	if w.Ground[0][0] == worldgenerator.Empty && w.Ground[48][48] == worldgenerator.Empty {
		path = astar.Run(
			w, worldgenerator.Coordinate{X: 0, Y: 0}, worldgenerator.Coordinate{X: 48, Y: 48},
		)
	}
	fmt.Print(path)

	astar.Plot(w, path)
}
