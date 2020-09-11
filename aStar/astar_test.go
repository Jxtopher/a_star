package aStar

import (
	"testing"

	"github.com/jxtopher/a_star/worldGenerator"
)

func TestGetNeighborhood(t *testing.T) {
	var xSize uint64 = 50
	var ySize uint64 = 50
	world := make([][]uint8, xSize)
	for i := range world {
		world[i] = make([]uint8, ySize)
	}

	w := worldGenerator.World{xSize, ySize, world}
	w.Init(0)

	pick := worldGenerator.Coordinate{5, 5}
	if len(getNeighborhood(w, pick)) != 8 {
		t.Error("")
	}
	pick = worldGenerator.Coordinate{0, 0}
	if len(getNeighborhood(w, pick)) != 3 {
		t.Error("")
	}
	pick = worldGenerator.Coordinate{49, 49}
	if len(getNeighborhood(w, pick)) != 3 {
		t.Error("")
	}
	pick = worldGenerator.Coordinate{0, 49}
	if len(getNeighborhood(w, pick)) != 3 {
		t.Error("")
	}
	pick = worldGenerator.Coordinate{49, 0}
	if len(getNeighborhood(w, pick)) != 3 {
		t.Error("")
	}
	pick = worldGenerator.Coordinate{49, 15}
	if len(getNeighborhood(w, pick)) != 5 {
		t.Error("")
	}
}

func TestGetDistance(t *testing.T) {
	if getDistance(worldGenerator.Coordinate{50, 50}, worldGenerator.Coordinate{0, 0})-70.71 > 0.1 {
		t.Error("")
	}
}
