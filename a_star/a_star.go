package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

func astar(sq SquareGrid, start point, goal point) (map[point]point, map[point]int) {
	cameFrom := make(map[point]point)
	costSoFar := make(map[point]int)
	frontier := PQ{Heap{}}

	frontier.put(dic{key: 0, value: start})
	cameFrom[start] = point{-1, -1}
	costSoFar[start] = 0

	for !frontier.empty() {

		current := frontier.get()

		if current == goal {
			break
		}

		neighbors := sq.neighbors(current)

		for i := 0; i < len(neighbors); i++ {
			next := neighbors[i]
			newCost := costSoFar[current] + 1
			if costSoFar[next] == 0 || newCost < costSoFar[next] {
				costSoFar[next] = newCost
				priority := float64(newCost) + myHeuristic(goal, next)
				frontier.put(dic{key: priority, value: next})
				cameFrom[next] = current
			}
		}
	}

	return cameFrom, costSoFar
}

func myHeuristic(a point, b point) float64 {
	return (float64(math.Sqrt(math.Pow(math.Abs(float64(b.y-a.y)), 2) + math.Pow(math.Abs(float64(b.x-a.x)), 2))))
}

func reconstructPath(cameFrom map[point]point, start point, goal point) []point {
	current := goal
	path := make([]point, 0)
	for current != start {
		path = append(path, current)
		current = cameFrom[current]
	}

	path = append(path, start)

	reversePointSlice(&path)

	return path
}

func reversePointSlice(s *[]point) {

	n := len(*s)

	appo := make([]point, n)

	for i := 0; i < n; i++ {
		appo[i] = (*s)[n-1-i]
	}

	*s = appo
}

// O(n)
func contain(a []point, v point) bool {
	i := 0

	for i < len(a) && a[i] != v {
		i++
	}

	return i < len(a)
}

func drawMatrix(sg SquareGrid, path []point, start point, goal point) {
	cost := 0

	fmt.Println("\ni) basato su un grafo NON pesato\n")

	for y := 0; y < sg.height; y++ {
		for x := 0; x < sg.width; x++ {

			actual := point{x, y}

			if actual == start {
				fmt.Print(" ")
				color.New(color.FgRed).Print(" S ")
				fmt.Print(" ")
				continue
			}

			if actual == goal {
				fmt.Print(" ")
				color.New(color.FgRed).Print(" G ")
				fmt.Print(" ")
				cost++
				continue
			}

			fmt.Print(" ")

			if contain(path, actual) {
				color.New(color.FgGreen).Print(" # ")
				cost++
			} else {
				color.New(color.FgBlue).Print(" * ")
			}

			fmt.Print(" ")
		}
		fmt.Println("\n")
	}

	fmt.Println("  costo -> ", cost)

}
