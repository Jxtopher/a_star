package astar

import (
	"fmt"
	"math"

	"github.com/jxtopher/a_star/worldgen"
)

type pair struct {
	distance float64
	cell     worldgen.Coordinate
}

// Give all free neighord nodes of the node chose
func getNeighborhood(
	w worldgen.World, pick worldgen.Coordinate,
) []worldgen.Coordinate {
	var neighborhood []worldgen.Coordinate

	if pick.Y < w.Ysize-1 && w.Ground[pick.X][pick.Y+1] == worldgen.Empty {
		neighborhood = append(
			neighborhood, worldgen.Coordinate{X: pick.X, Y: pick.Y + 1},
		)
	}
	if pick.X < w.Xsize-1 && pick.Y < w.Ysize-1 &&
		w.Ground[pick.X+1][pick.Y+1] == worldgen.Empty {
		// then
		neighborhood = append(neighborhood, worldgen.Coordinate{X: pick.X + 1, Y: pick.Y + 1})
	}
	if pick.X < w.Xsize-1 && w.Ground[pick.X+1][pick.Y] == worldgen.Empty {
		neighborhood = append(neighborhood, worldgen.Coordinate{X: pick.X + 1, Y: pick.Y})
	}
	if pick.X < w.Xsize-1 && pick.Y > 0 && w.Ground[pick.X+1][pick.Y-1] == worldgen.Empty {
		neighborhood = append(neighborhood, worldgen.Coordinate{X: pick.X + 1, Y: pick.Y - 1})
	}
	if pick.Y > 0 && w.Ground[pick.X][pick.Y-1] == worldgen.Empty {
		neighborhood = append(neighborhood, worldgen.Coordinate{X: pick.X, Y: pick.Y - 1})
	}
	if pick.X > 0 && pick.Y > 0 && w.Ground[pick.X-1][pick.Y-1] == worldgen.Empty {
		neighborhood = append(neighborhood, worldgen.Coordinate{X: pick.X - 1, Y: pick.Y - 1})
	}
	if pick.X > 0 && w.Ground[pick.X-1][pick.Y] == worldgen.Empty {
		neighborhood = append(neighborhood, worldgen.Coordinate{X: pick.X - 1, Y: pick.Y})
	}
	if pick.X > 0 && pick.Y < w.Ysize-1 && w.Ground[pick.X-1][pick.Y+1] == worldgen.Empty {
		neighborhood = append(neighborhood, worldgen.Coordinate{X: pick.X - 1, Y: pick.Y + 1})
	}

	return neighborhood
}

// Give distance between two nodes
func getDistance(a worldgen.Coordinate, b worldgen.Coordinate) float64 {
	return math.Abs(float64(int64(a.X-b.X))) + math.Abs(float64(int64(a.Y-b.Y)))
}

// Returns an index on the node's set with the minimum fScore
func minScore(set []worldgen.Coordinate, fScore map[worldgen.Coordinate]float64) int {
	minElement := 0
	for i, element := range set {
		if i < 1 {
			continue
		}

		if fScore[set[minElement]] < fScore[element] {
			minElement = i
		}
	}
	return minElement
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []worldgen.Coordinate, val worldgen.Coordinate) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// reconstructionPath build the final path
func reconstructPath(
	cameFrom map[worldgen.Coordinate]worldgen.Coordinate, end worldgen.Coordinate,
) []worldgen.Coordinate {
	var ret []worldgen.Coordinate
	pick := end
	_, found := cameFrom[pick]
	for found {
		ret = append(ret, pick)
		pick = cameFrom[pick]
		_, found = cameFrom[pick]
	}
	ret = append(ret, pick)
	return ret
}

func Run(
	w worldgen.World,
	start worldgen.Coordinate,
	end worldgen.Coordinate,
) []worldgen.Coordinate {
	var openSet []worldgen.Coordinate
	openSet = append(openSet, start)
	cameFrom := make(map[worldgen.Coordinate]worldgen.Coordinate)

	fScore := make(map[worldgen.Coordinate]float64)
	gScore := make(map[worldgen.Coordinate]float64)
	gScore[start] = 0.0
	fScore[start] = getDistance(start, end)

	for len(openSet) != 0 {
		currentIndex := minScore(openSet, fScore)
		current := openSet[currentIndex]
		openSet = append(openSet[:currentIndex], openSet[currentIndex+1:]...)

		fmt.Println(current)
		if current == end {

			return reconstructPath(cameFrom, end)
		}

		neighborhood := getNeighborhood(w, current)

		for _, neighbor := range neighborhood {
			_, found := gScore[neighbor]
			if !found {
				gScore[neighbor] = gScore[current] + 1
				fScore[neighbor] = gScore[current] + 1 + getDistance(neighbor, end)
				cameFrom[neighbor] = current
				_, elementFound := Find(openSet, neighbor)
				if !elementFound {
					openSet = append(openSet, neighbor)
				}
			} else {
				if gScore[current]+1 < gScore[neighbor] {
					gScore[neighbor] = gScore[current] + 1
					fScore[neighbor] = gScore[current] + 1 + getDistance(neighbor, end)
					cameFrom[neighbor] = current
					_, elementFound := Find(openSet, neighbor)
					if !elementFound {
						openSet = append(openSet, neighbor)
					}
				}
			}
		}
	}
	return []worldgen.Coordinate{}
}
