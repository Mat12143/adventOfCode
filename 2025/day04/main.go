package day04

import (
	"fmt"
	"strings"
)

type Point struct {
	x, y          int
	isToiletPaper bool
}

func printMatrix(matrix [][]Point) {
	for _, r := range matrix {
		for _, l := range r {
			if !l.isToiletPaper {
				fmt.Print(".")

			} else {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}
}

func getNearbyPapers(matrix [][]Point, location Point) int {
	directions := [][]int{
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
	}

	nearby := 0

	for _, d := range directions {
		x := location.x + d[0]
		y := location.y + d[1]

		// Out of bounds
		if x < 0 || x >= len(matrix) {
			continue
		}

		if y < 0 || y >= len(matrix[x]) {
			continue
		}

		if matrix[x][y].isToiletPaper {
			nearby++
		}

	}

	return nearby
}

func parseMatrix(data string) [][]Point {
	matrix := [][]Point{}

	for i, l := range strings.Split(data, "\n") {

		newline := []Point{}
		for j, c := range l {
			if string(c) == "@" {
				newline = append(newline, Point{
					i, j, true,
				})
			} else {
				newline = append(newline, Point{
					i, j, false,
				})
			}
		}
		matrix = append(matrix, newline)
	}

	return matrix

}

func Part1(data string) {

	tot := 0

	matrix := parseMatrix(data)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j].isToiletPaper {
				if getNearbyPapers(matrix, matrix[i][j]) < 4 {
					tot++
				}
			}
		}
	}

	fmt.Printf("tot: %v\n", tot)

}

func Part2(data string) {
	removed := 0

	matrix := parseMatrix(data)

	for {
		lastRemoved := 0
		for i := range matrix {
			for j := range matrix[i] {
				if matrix[i][j].isToiletPaper && getNearbyPapers(matrix, matrix[i][j]) < 4 {
					removed++
					lastRemoved++
					// Remove the toilet paper
					matrix[i][j].isToiletPaper = false
				}
			}
		}
		// No more removals
		if lastRemoved == 0 {
			break
		}
	}

	fmt.Printf("removed: %v\n", removed)
}
