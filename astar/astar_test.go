package astar

import (
	"testing"

	"github.com/jxtopher/a_star/worldgenerator"
	"github.com/stretchr/testify/assert"
)

func TestGetNeighborhood(t *testing.T) {
	var xSize uint64 = 50
	var ySize uint64 = 50
	world := make([][]uint8, xSize)
	for i := range world {
		world[i] = make([]uint8, ySize)
	}

	w := worldgenerator.World{Xsize: xSize, Ysize: ySize, Ground: world}
	w.Init(0)

	pick := worldgenerator.Coordinate{X: 5, Y: 5}
	assert.Equal(t, len(getNeighborhood(w, pick)), 8, "they should be equal")
	pick = worldgenerator.Coordinate{X: 0, Y: 0}
	assert.Equal(t, len(getNeighborhood(w, pick)), 3, "they should be equal")
	pick = worldgenerator.Coordinate{X: 49, Y: 49}
	assert.Equal(t, len(getNeighborhood(w, pick)), 3, "they should be equal")
	pick = worldgenerator.Coordinate{X: 0, Y: 49}
	assert.Equal(t, len(getNeighborhood(w, pick)), 3, "they should be equal")
	pick = worldgenerator.Coordinate{X: 49, Y: 0}
	assert.Equal(t, len(getNeighborhood(w, pick)), 3, "they should be equal")
	pick = worldgenerator.Coordinate{X: 49, Y: 15}
	assert.Equal(t, len(getNeighborhood(w, pick)), 5, "they should be equal")
}

func TestGetDistance(t *testing.T) {
	distance := getDistance(
		worldgenerator.Coordinate{X: 50, Y: 50},
		worldgenerator.Coordinate{X: 0, Y: 0},
	)
	assert.Less(t, distance-70.71, 0.1, "")
}
