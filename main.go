package main

import (
	"fmt"

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
	w := worldgen.Init(7, 7, 0.1)
	// w := worldgen.Loadjson("./dataworld/world.json")
	w.Show()
	fmt.Print(w.GetBox(4, 3))

	// var xSize uint64 = 7
	// var ySize uint64 = 7
	// world := make([][]uint8, xSize)
	// for i := range world {
	// 	world[i] = make([]uint8, ySize)
	// }

	// w := worldgen.World{Xsize: xSize, Ysize: ySize, Ground: world}
	// w.Ground = [][]uint8{
	// 	{0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0, 0},
	// }

	// var path []worldgen.Coordinate
	// if w.Ground[2][0] == worldgen.Empty && w.Ground[2][3] == worldgen.Empty {
	// 	path = astar.Run(
	// 		w, worldgen.Coordinate{X: 2, Y: 0}, worldgen.Coordinate{X: 6, Y: 6},
	// 	)
	// }
	// fmt.Print(path)

	// astar.Plot(w, path)
}
