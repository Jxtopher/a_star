package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
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
	log.Println("Starting the search for a shorter path with algorithm A*")
	// command-Line Flags
	worldPtr := flag.String("world", "", "a string")
	flag.Parse()

	var w worldgen.World
	var points startEnd

	if *worldPtr == "" {
		w = worldgen.Init(50, 50, 0.3)
		points.Start = []uint64{0, 0}
		points.End = []uint64{49, 49}
		// w.Show()
	} else {
		// go run .\main.go -world=".\dataworld\world.json"
		w = worldgen.Loadjson(*worldPtr)
		// w.Show()

		jsonFile, err := os.Open(*worldPtr)
		if err != nil {
			log.Fatal(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &points)
	}

	// Find path
	var path []worldgen.Coordinate
	path = astar.Run(
		w,
		worldgen.Coordinate{X: points.Start[0], Y: points.Start[1]},
		worldgen.Coordinate{X: points.End[0], Y: points.End[1]},
	)

	log.Println("Plot the solution")
	astar.Plot("hello", w, path)
}
