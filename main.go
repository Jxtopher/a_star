package main

import (
	"fmt"

	"github.com/jxtopher/a_star/astar"
	"github.com/jxtopher/a_star/worldgen"
)

func main() {
	var xSize uint64 = 4
	var ySize uint64 = 4
	world := make([][]uint8, xSize)
	for i := range world {
		world[i] = make([]uint8, ySize)
	}

	w := worldgen.World{Xsize: xSize, Ysize: ySize, Ground: world}
	w.Ground = [][]uint8{
		{0, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}

	var path []worldgen.Coordinate
	if w.Ground[2][0] == worldgen.Empty && w.Ground[2][3] == worldgen.Empty {
		path = astar.Run(
			w, worldgen.Coordinate{X: 2, Y: 0}, worldgen.Coordinate{X: 2, Y: 3},
		)
	}
	fmt.Print(path)

	astar.Plot(w, path)
}
