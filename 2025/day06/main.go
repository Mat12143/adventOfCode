package day06

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func intoMatrix(data string) ([][]int, []string) {

	numbers := [][]int{}
	operators := []string{}

	for line := range strings.SplitSeq(data, "\n") {

		newLineNumbers := []int{}

		for n := range strings.SplitSeq(line, " ") {
			// Remove spaces
			n = strings.Trim(n, " ")

			if len(n) == 0 {
				continue
			}

			num, err := strconv.Atoi(n)
			if err != nil {
				// Operator
				operators = append(operators, n)
			} else {
				newLineNumbers = append(newLineNumbers, num)
			}

		}
		if len(newLineNumbers) != 0 {
			numbers = append(numbers, newLineNumbers)
		}
	}

	return numbers, operators

}

func Part1(data string) {

	numbers, op := intoMatrix(data)

	tot := 0

	// columns
	for i := range numbers[0] {
		partial := numbers[0][i]

		// row of column
		for j := 1; j < len(numbers); j++ {
			if op[i] == "*" {
				partial *= numbers[j][i]
			} else {
				partial += numbers[j][i]
			}
		}

		tot += partial
	}

	fmt.Printf("tot: %v\n", tot)

}

func Part2(data string) {

	lines := []string{}

	for l := range strings.SplitSeq(data, "\n") {
		lines = append(lines, l)
	}

	op := lines[len(lines)-2:][0]

	operators := []string{}

	for i := range op {
		if string(op[i]) != " " {
			operators = append(operators, string(op[i]))
		}
	}

	lines = lines[:len(lines)-2]

	tot := 0

	partial := 0
	calcIndex := 0

	for i := range lines[0] {
		number := ""

		for j := 0; j < len(lines); j++ {
			if string(lines[j][i]) != " " {
				number += string(lines[j][i])
			}
		}

		if number != "" {
			n, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			if partial == 0 {
				partial = n
			} else {
				if operators[calcIndex] == "*" {
					partial *= n
				} else {
					partial += n
				}
			}

			// check for end
			if i+1 == len(lines[0]) {
				tot += partial
			}

		} else {
			tot += partial
			partial = 0
			calcIndex++
		}
	}
	fmt.Printf("tot: %v\n", tot)
}
