package worldGenerator

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
	X_size uint64
	Y_size uint64
	Ground [][]uint8
}

func (w World) Init(p float64) {
	var i, j uint64
	for i = 0; i < w.X_size; i++ {
		for j = 0; j < w.Y_size; j++ {
			if rand.Float64() < p {
				w.Ground[i][j] = Bloc
			}
		}
	}
}

func (w World) Show() {
	var i, j uint64
	for i = 0; i < w.X_size; i++ {
		for j = 0; j < w.Y_size; j++ {
			//   fmt.Printf("a[%d][%d] = %d\n", i,j, w.Ground[i][j] )
			fmt.Printf("%d ", w.Ground[i][j])
		}
		fmt.Printf("\n")
	}
}
