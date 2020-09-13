package astar

import (
	"testing"

	"github.com/jxtopher/a_star/worldgen"
	"github.com/stretchr/testify/assert"
)

func TestGetNeighborhood(t *testing.T) {
	w := worldgen.Init(50, 50, 0)

	pick := worldgen.Coordinate{X: 5, Y: 5}
	assert.Equal(t, len(getNeighborhood(w, pick)), 8, "they should be equal")
	pick = worldgen.Coordinate{X: 0, Y: 0}
	assert.Equal(t, len(getNeighborhood(w, pick)), 3, "they should be equal")
	pick = worldgen.Coordinate{X: 49, Y: 49}
	assert.Equal(t, len(getNeighborhood(w, pick)), 3, "they should be equal")
	pick = worldgen.Coordinate{X: 0, Y: 49}
	assert.Equal(t, len(getNeighborhood(w, pick)), 3, "they should be equal")
	pick = worldgen.Coordinate{X: 49, Y: 0}
	assert.Equal(t, len(getNeighborhood(w, pick)), 3, "they should be equal")
	pick = worldgen.Coordinate{X: 49, Y: 15}
	assert.Equal(t, len(getNeighborhood(w, pick)), 5, "they should be equal")
}

func TestGetDistance(t *testing.T) {
	distance := getDistance(
		worldgen.Coordinate{X: 50, Y: 50},
		worldgen.Coordinate{X: 0, Y: 0},
	)
	assert.Equal(t, distance, 100.0, "")
}
