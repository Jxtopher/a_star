package aStar

import (
	"fmt"
	"math"

	"github.com/jxtopher/a_star/worldGenerator"
)

// Give all free neighord node  of the node chose
func getNeighborhood(w worldGenerator.World, pick worldGenerator.Coordinate) []worldGenerator.Coordinate {
	var neighborhood []worldGenerator.Coordinate

	if pick.Y < w.Y_size-1 && w.Ground[pick.X][pick.Y+1] == worldGenerator.Empty {
		neighborhood = append(neighborhood, worldGenerator.Coordinate{pick.X, pick.Y + 1})
	}
	if pick.X < w.X_size-1 && pick.Y < w.Y_size-1 && w.Ground[pick.X+1][pick.Y+1] == worldGenerator.Empty {
		neighborhood = append(neighborhood, worldGenerator.Coordinate{pick.X + 1, pick.Y + 1})
	}
	if pick.X < w.X_size-1 && w.Ground[pick.X+1][pick.Y] == worldGenerator.Empty {
		neighborhood = append(neighborhood, worldGenerator.Coordinate{pick.X + 1, pick.Y})
	}
	if pick.X < w.X_size-1 && pick.Y > 0 && w.Ground[pick.X+1][pick.Y-1] == worldGenerator.Empty {
		neighborhood = append(neighborhood, worldGenerator.Coordinate{pick.X + 1, pick.Y - 1})
	}
	if pick.Y > 0 && w.Ground[pick.X][pick.Y-1] == worldGenerator.Empty {
		neighborhood = append(neighborhood, worldGenerator.Coordinate{pick.X, pick.Y - 1})
	}
	if pick.X > 0 && pick.Y > 0 && w.Ground[pick.X-1][pick.Y-1] == worldGenerator.Empty {
		neighborhood = append(neighborhood, worldGenerator.Coordinate{pick.X - 1, pick.Y - 1})
	}
	if pick.X > 0 && w.Ground[pick.X-1][pick.Y] == worldGenerator.Empty {
		neighborhood = append(neighborhood, worldGenerator.Coordinate{pick.X - 1, pick.Y})
	}
	if pick.X > 0 && pick.Y < w.Y_size-1 && w.Ground[pick.X-1][pick.Y+1] == worldGenerator.Empty {
		neighborhood = append(neighborhood, worldGenerator.Coordinate{pick.X - 1, pick.Y + 1})
	}

	return neighborhood
}

// Give distance between two nodes
func getDistance(a worldGenerator.Coordinate, b worldGenerator.Coordinate) float64 {
	var A = math.Abs(float64(int64(a.X - b.X)))
	var B = math.Abs(float64(int64(a.Y - b.Y)))
	return math.Sqrt(A*A + B*B)
}

type pair struct {
	distance float64
	cell     worldGenerator.Coordinate
}

func Run(w worldGenerator.World, start worldGenerator.Coordinate, end worldGenerator.Coordinate) []worldGenerator.Coordinate {
	// w.Show()
	m := make(map[worldGenerator.Coordinate]uint64)

	var path []worldGenerator.Coordinate
	worklist := make([]pair, 0)
	var pick worldGenerator.Coordinate = start
	m[pick] = 0
	for pick != end {
		fmt.Print("pick>", pick, "\n")
		neighborhood := getNeighborhood(w, pick)
		fmt.Print("neighborhood>", neighborhood, "\n")
		for _, neighbour := range neighborhood {
			_, found := m[neighbour]
			if !found {
				m[neighbour] = m[pick] + 1
				worklist = append(worklist, pair{getDistance(neighbour, end), neighbour})
			}
		}

		fmt.Print("worklist>", worklist, "\n")
		// find cell with min distance at the end
		var minIndex = 0
		for i, e := range worklist {
			if e.distance < worklist[minIndex].distance {
				minIndex = i
			}
		}
		pick = worklist[minIndex].cell
		fmt.Print(pick, "\n")
		worklist = worklist[:0]
	}

	//
	neighborhood := getNeighborhood(w, pick)
	fmt.Print("neighborhood>", neighborhood, "\n")
	for _, neighbour := range neighborhood {
		_, found := m[neighbour]
		if !found {
			m[neighbour] = m[pick] + 1
			worklist = append(worklist, pair{getDistance(neighbour, end), neighbour})
		}
	}

	// Rebuild of the path
	pick = end
	path = append(path, pick)
	for pick != start {
		neighborhood := getNeighborhood(w, pick)
		min_value := neighborhood[0]
		for _, neighbour := range neighborhood {
			fmt.Print("(", neighbour, " : ", m[neighbour], ") ")
			if m[neighbour] < m[min_value] {
				min_value = neighbour
			}
		}

		pick = min_value
		fmt.Println("\n>", pick)
		path = append(path, pick)
		// return
	}
	return path
	// fmt.Println(path)
}
