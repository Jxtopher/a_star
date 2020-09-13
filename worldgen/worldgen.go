package worldgen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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
	Xsize  uint64    `json:"xsize"`
	Ysize  uint64    `json:"ysize"`
	Ground [][]uint8 `json:"ground"`
}

// Init the World
func Init(xsize uint64, ysize uint64, p float64) World {
	w := World{Xsize: xsize, Ysize: ysize}
	fmt.Println(w.Xsize)
	w.Ground = make([][]uint8, w.Xsize)
	for i := range w.Ground {
		w.Ground[i] = make([]uint8, w.Ysize)
	}
	var i, j uint64
	for i = 0; i < w.Xsize; i++ {
		for j = 0; j < w.Ysize; j++ {
			if rand.Float64() < p {
				w.setBox(i, j, Bloc)
			}
		}
	}
	return w
}

func Loadjson(pathFile string) World {
	var w World
	jsonFile, err := os.Open(pathFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &w)
	return w
}

// Get element in box
func (w World) setBox(x uint64, y uint64, element uint8) {
	w.Ground[y][x] = element
}

// Get element in box
func (w World) GetBox(x uint64, y uint64) uint8 {
	return w.Ground[y][x]
}

// Show the World
func (w World) Show() {
	var i, j uint64
	fmt.Println(w.Xsize)

	for i = 0; i < w.Xsize; i++ {
		for j = 0; j < w.Ysize; j++ {
			fmt.Printf("%d ", w.GetBox(j, i))
		}
		fmt.Printf("\n")
	}
}
