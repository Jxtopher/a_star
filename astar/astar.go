package astar

import (
	"fmt"
	"math"

	"github.com/jxtopher/a_star/worldgenerator"
)

// Give all free neighord nodes of the node chose
func getNeighborhood(
	w worldgenerator.World, pick worldgenerator.Coordinate,
) []worldgenerator.Coordinate {
	var neighborhood []worldgenerator.Coordinate

	if pick.Y < w.Ysize-1 && w.Ground[pick.X][pick.Y+1] == worldgenerator.Empty {
		neighborhood = append(
			neighborhood, worldgenerator.Coordinate{X: pick.X, Y: pick.Y + 1},
		)
	}
	if pick.X < w.Xsize-1 && pick.Y < w.Ysize-1 &&
		w.Ground[pick.X+1][pick.Y+1] == worldgenerator.Empty {
		// then
		neighborhood = append(neighborhood, worldgenerator.Coordinate{X: pick.X + 1, Y: pick.Y + 1})
	}
	if pick.X < w.Xsize-1 && w.Ground[pick.X+1][pick.Y] == worldgenerator.Empty {
		neighborhood = append(neighborhood, worldgenerator.Coordinate{X: pick.X + 1, Y: pick.Y})
	}
	if pick.X < w.Xsize-1 && pick.Y > 0 && w.Ground[pick.X+1][pick.Y-1] == worldgenerator.Empty {
		neighborhood = append(neighborhood, worldgenerator.Coordinate{X: pick.X + 1, Y: pick.Y - 1})
	}
	if pick.Y > 0 && w.Ground[pick.X][pick.Y-1] == worldgenerator.Empty {
		neighborhood = append(neighborhood, worldgenerator.Coordinate{X: pick.X, Y: pick.Y - 1})
	}
	if pick.X > 0 && pick.Y > 0 && w.Ground[pick.X-1][pick.Y-1] == worldgenerator.Empty {
		neighborhood = append(neighborhood, worldgenerator.Coordinate{X: pick.X - 1, Y: pick.Y - 1})
	}
	if pick.X > 0 && w.Ground[pick.X-1][pick.Y] == worldgenerator.Empty {
		neighborhood = append(neighborhood, worldgenerator.Coordinate{X: pick.X - 1, Y: pick.Y})
	}
	if pick.X > 0 && pick.Y < w.Ysize-1 && w.Ground[pick.X-1][pick.Y+1] == worldgenerator.Empty {
		neighborhood = append(neighborhood, worldgenerator.Coordinate{X: pick.X - 1, Y: pick.Y + 1})
	}

	return neighborhood
}

// Give distance between two nodes
func getDistance(a worldgenerator.Coordinate, b worldgenerator.Coordinate) float64 {
	var A = math.Abs(float64(int64(a.X - b.X)))
	var B = math.Abs(float64(int64(a.Y - b.Y)))
	return math.Sqrt(A*A + B*B)
}

type pair struct {
	distance float64
	cell     worldgenerator.Coordinate
}

// Returns an index on the node's set with the minimum fScore
func minScore(set []worldgenerator.Coordinate, fScore map[worldgenerator.Coordinate]float64) int {
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
func Find(slice []worldgenerator.Coordinate, val worldgenerator.Coordinate) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// reconstructionPath build the final path
func reconstructPath(
	cameFrom map[worldgenerator.Coordinate]worldgenerator.Coordinate, end worldgenerator.Coordinate,
) []worldgenerator.Coordinate {
	var ret []worldgenerator.Coordinate
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
	w worldgenerator.World,
	start worldgenerator.Coordinate,
	end worldgenerator.Coordinate,
) []worldgenerator.Coordinate {
	var openSet []worldgenerator.Coordinate
	openSet = append(openSet, start)
	cameFrom := make(map[worldgenerator.Coordinate]worldgenerator.Coordinate)

	fScore := make(map[worldgenerator.Coordinate]float64)
	gScore := make(map[worldgenerator.Coordinate]float64)
	gScore[start] = 0.0
	fScore[start] = getDistance(start, end)

	for len(openSet) != 0 {
		current := minScore(openSet, fScore)
		pick := openSet[current]
		openSet = append(openSet[:current], openSet[current+1:]...)

		fmt.Println(pick)
		if pick == end {

			return reconstructPath(cameFrom, end)
		}

		neighborhood := getNeighborhood(w, pick)

		for _, neighbor := range neighborhood {
			_, found := gScore[neighbor]
			if !found {
				gScore[neighbor] = gScore[pick] + 1
				fScore[neighbor] = gScore[pick] + 1 + getDistance(neighbor, end)
				cameFrom[neighbor] = pick
				_, elementFound := Find(openSet, neighbor)
				if !elementFound {
					openSet = append(openSet, neighbor)
				}
			} else {
				if gScore[pick]+1 < gScore[neighbor] {
					gScore[neighbor] = gScore[pick] + 1
					fScore[neighbor] = gScore[pick] + 1 + getDistance(neighbor, end)
					cameFrom[neighbor] = pick
					_, elementFound := Find(openSet, neighbor)
					if !elementFound {
						openSet = append(openSet, neighbor)
					}
				}
			}
		}
	}
	return []worldgenerator.Coordinate{}
}
