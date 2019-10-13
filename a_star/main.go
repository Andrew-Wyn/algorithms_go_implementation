package main

import (
	"fmt"
)

func main() {
	fmt.Println("\t*** A-STAR ***")

	start := point{0, 0}
	goal := point{7, 9}

	sg := SquareGrid{height: 10, width: 10}

	points, _ := astar(sg, start, goal)

	drawMatrix(sg, reconstructPath(points, start, goal), start, goal)

}
