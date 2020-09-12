package main

import (
	// "fmt"

	"fmt"

	"github.com/jxtopher/a_star/astar"
	"github.com/jxtopher/a_star/worldgenerator"
	// "log"
)

func main() {
	var xSize uint64 = 50
	var ySize uint64 = 50
	world := make([][]uint8, xSize)
	for i := range world {
		world[i] = make([]uint8, ySize)
	}

	w := worldgenerator.World{xSize, ySize, world}
	w.Init(0.0)

	var path []worldgenerator.Coordinate
	if w.Ground[0][0] == worldgenerator.Empty && w.Ground[48][48] == worldgenerator.Empty {
		path = astar.Run(w, worldgenerator.Coordinate{0, 0}, worldgenerator.Coordinate{48, 48})
	}
	fmt.Print(path)

	astar.Plot(w, path)
}
