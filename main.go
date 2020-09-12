package main

import (
	"fmt"

	"github.com/jxtopher/a_star/astar"
	"github.com/jxtopher/a_star/worldgen"
)

func main() {
	var xSize uint64 = 50
	var ySize uint64 = 50
	world := make([][]uint8, xSize)
	for i := range world {
		world[i] = make([]uint8, ySize)
	}

	w := worldgen.World{Xsize: xSize, Ysize: ySize, Ground: world}
	w.Init(0.0)

	var path []worldgen.Coordinate
	if w.Ground[0][0] == worldgen.Empty && w.Ground[48][48] == worldgen.Empty {
		path = astar.Run(
			w, worldgen.Coordinate{X: 0, Y: 0}, worldgen.Coordinate{X: 48, Y: 48},
		)
	}
	fmt.Print(path)

	astar.Plot(w, path)
}
