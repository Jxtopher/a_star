package main

import (
	"flag"

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

func main() {
	// command-Line Flags
	worldPtr := flag.String("world", "", "a string")
	flag.Parse()

	var w worldgen.World
	if *worldPtr == "" {
		w = worldgen.Init(7, 7, 0.1)
		w.Show()
	} else {
		// go run .\main.go -world=".\dataworld\world.json"
		w = worldgen.Loadjson(*worldPtr)
		w.Show()
	}

	// Find path
	// var path []worldgen.Coordinate
	// if w.Ground[2][0] == worldgen.Empty && w.Ground[2][3] == worldgen.Empty {
	// 	path = astar.Run(
	// 		w, worldgen.Coordinate{X: 2, Y: 0}, worldgen.Coordinate{X: 6, Y: 6},
	// 	)
	// }
	// fmt.Print(path)

	astar.Plot("hello", w, []worldgen.Coordinate{})
}
