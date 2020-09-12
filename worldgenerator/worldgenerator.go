package worldgenerator

import (
	"fmt"
	"math/rand"
)

const (
	Empty uint8 = 0
	Bloc  uint8 = 1
)

type Coordinate struct {
	X uint64
	Y uint64
}

type World struct {
	Xsize  uint64
	Ysize  uint64
	Ground [][]uint8
}

// Init the world
func (w World) Init(p float64) {
	var i, j uint64
	for i = 0; i < w.Xsize; i++ {
		for j = 0; j < w.Ysize; j++ {
			if rand.Float64() < p {
				w.Ground[i][j] = Bloc
			}
		}
	}
}

// Show the world
func (w World) Show() {
	var i, j uint64
	for i = 0; i < w.Xsize; i++ {
		for j = 0; j < w.Ysize; j++ {
			//   fmt.Printf("a[%d][%d] = %d\n", i,j, w.Ground[i][j] )
			fmt.Printf("%d ", w.Ground[i][j])
		}
		fmt.Printf("\n")
	}
}
