package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jxtopher/a_star/astar"
	"github.com/jxtopher/a_star/worldgen"
)

//   0  ---> x   6
// 0 0 0 0 0 0 1 0
//   0 1 0 0 0 0 0
// | 0 0 0 0 0 1 0
// | 0 0 0 0 1 0 0
// y 0 0 0 0 0 0 0
//   0 0 0 0 0 0 0
// 6 0 0 0 0 0 0 0

type startEnd struct {
	Start []uint64 `json:"start"`
	End   []uint64 `json:"end"`
}

func main() {
	// command-Line Flags
	worldPtr := flag.String("world", "", "a string")
	flag.Parse()

	var w worldgen.World
	var points startEnd

	if *worldPtr == "" {
		w = worldgen.Init(30, 30, 0.2)
		points.Start = []uint64{0, 0}
		points.End = []uint64{29, 29}
		w.Show()
	} else {
		// go run .\main.go -world=".\dataworld\world.json"
		w = worldgen.Loadjson(*worldPtr)
		w.Show()

		jsonFile, err := os.Open(*worldPtr)
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &points)
		fmt.Print(points.End)

	}

	// Find path
	var path []worldgen.Coordinate
	// if w.Ground[2][0] == worldgen.Empty && w.Ground[2][3] == worldgen.Empty {
	path = astar.Run(
		w,
		worldgen.Coordinate{X: points.Start[0], Y: points.Start[1]},
		worldgen.Coordinate{X: points.End[0], Y: points.End[1]},
	)
	// }
	// fmt.Print(path)

	astar.Plot("hello", w, path)
}
