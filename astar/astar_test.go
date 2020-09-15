package astar

import (
	"testing"

	"github.com/jxtopher/a_star/worldgen"
	"github.com/stretchr/testify/assert"
)

func TestGetNeighborhood(t *testing.T) {
	w := worldgen.Init(50, 50, 0)

	pick := worldgen.Coordinate{X: 5, Y: 5}
	assert.Equal(t, len(getNeighborhood(w, pick, true)), 8, "they should be equal")
	pick = worldgen.Coordinate{X: 0, Y: 0}
	assert.Equal(t, len(getNeighborhood(w, pick, true)), 3, "they should be equal")
	pick = worldgen.Coordinate{X: 49, Y: 49}
	assert.Equal(t, len(getNeighborhood(w, pick, true)), 3, "they should be equal")
	pick = worldgen.Coordinate{X: 0, Y: 49}
	assert.Equal(t, len(getNeighborhood(w, pick, true)), 3, "they should be equal")
	pick = worldgen.Coordinate{X: 49, Y: 0}
	assert.Equal(t, len(getNeighborhood(w, pick, true)), 3, "they should be equal")
	pick = worldgen.Coordinate{X: 49, Y: 15}
	assert.Equal(t, len(getNeighborhood(w, pick, true)), 5, "they should be equal")
}

func TestFind(t *testing.T) {
	slice := []worldgen.Coordinate{{X: 0, Y: 0}, {X: 3, Y: 4}}
	element1 := worldgen.Coordinate{X: 0, Y: 0}
	index, found := find(slice, element1)
	assert.Equal(t, index, 0, "they should be equal")
	assert.Equal(t, found, true, "they should be equal")

	element2 := worldgen.Coordinate{X: 32, Y: 0}
	index, found = find(slice, element2)
	assert.Equal(t, index, -1, "they should be equal")
	assert.Equal(t, found, false, "they should be equal")
}

func TestGetDistance(t *testing.T) {
	distance := getDistance(
		worldgen.Coordinate{X: 50, Y: 50},
		worldgen.Coordinate{X: 0, Y: 0},
	)
	assert.Equal(t, distance, 100.0, "")
}
