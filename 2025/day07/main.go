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

	timeLines := 0

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
					timeLines += 2
					matrix[i+1][j-1] = "|"
					matrix[i+1][j+1] = "|"

				} else {
					matrix[i+1][j] = "|"

				}
			}
		}
	}

	fmt.Printf("timeLines: %v\n", timeLines)
}
