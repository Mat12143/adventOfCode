package day07

import (
	"fmt"
	"strings"
)

func Part1(data string) {

	matrix := [][]string{}

	splitted := 0

	for l := range strings.SplitSeq(data, "\n") {

		if len(l) == 0 {
			continue
		}

		line := []string{}
		for _, c := range l {
			line = append(line, string(c))
		}
		matrix = append(matrix, line)
	}

	for i := 0; i < len(matrix)-1; i++ {
		for j := range matrix[i] {
			curr := matrix[i][j]

			switch curr {
			case "S":
				matrix[i+1][j] = "|"
			case "|":
				if matrix[i+1][j] == "^" {
					splitted++
					matrix[i+1][j-1] = "|"
					matrix[i+1][j+1] = "|"

				} else {
					matrix[i+1][j] = "|"

				}
			}
		}
	}

	fmt.Printf("splitted: %v\n", splitted)

}

func Part2(data string) {
	matrix := [][]string{}

	for l := range strings.SplitSeq(data, "\n") {
		if len(l) == 0 {
			continue
		}

		line := []string{}
		for _, c := range l {
			line = append(line, string(c))
		}
		matrix = append(matrix, line)
	}

	R := len(matrix)
	C := len(matrix[0])

	numberMatrix := make([][]int, R)
	for i := range R {
		numberMatrix[i] = make([]int, C)
	}

	for i := 0; i < R-1; i++ {
		for j := range C {

			curr := numberMatrix[i][j]

			cell := matrix[i][j]

			switch cell {
			case "S":
				numberMatrix[i+1][j] = 1
			case ".":
				if curr != 0 {
					numberMatrix[i+1][j] += curr
				}
			case "^":
				numberMatrix[i+1][j-1] += curr
				numberMatrix[i+1][j+1] += curr
			}
		}
	}

	tot := 0
	for _, v := range numberMatrix[R-1] {
		tot += v
	}

	fmt.Printf("tot: %v\n", tot)
}
