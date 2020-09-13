package astar

import (
	"testing"

	"github.com/jxtopher/a_star/worldgen"
)

func TestPlot(t *testing.T) {
	var xSize uint64 = 7
	var ySize uint64 = 7
	world := make([][]uint8, xSize)
	for i := range world {
		world[i] = make([]uint8, ySize)
	}

	w := worldgen.World{Xsize: xSize, Ysize: ySize, Ground: world}
	w.Ground = [][]uint8{
		{0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0},
	}

	path := []worldgen.Coordinate{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}
	Plot("test_plot1", w, path)
}
