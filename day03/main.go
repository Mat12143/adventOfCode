package day03

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/Mat12143/adventOfCode2025/utils"
)

type Battery struct {
	jolt int
	pos  int
}

func Part1() {
	data := utils.ReadFileByLine("day03/input.txt")

	sum := 0

	for l := range strings.SplitSeq(data, "\n") {
		if len(l) == 0 {
			continue
		}

		numbers := []int{}

		for _, d := range l {
			numbers = append(numbers, int(d-'0'))
		}
		sum += getMaxJoltage(numbers, 2)

	}
	fmt.Printf("sum: %v\n", sum)
}

func getMaxJoltage(list []int, dim int) int {
	if dim == 0 {
		return 0
	}

	maxValue := slices.Max(list[:len(list)-dim+1])
	maxIndex := slices.Index(list, maxValue)

	return (int(math.Pow10(dim-1)) * maxValue) + getMaxJoltage(list[maxIndex+1:], dim-1)

}

func Part2() {
	sum := 0

	data := utils.ReadFileByLine("day03/input.txt")

	for l := range strings.SplitSeq(data, "\n") {

		if len(l) == 0 {
			continue
		}
		numbers := []int{}

		for _, d := range l {
			numbers = append(numbers, int(d-'0'))
		}

		test := getMaxJoltage(numbers, 12)
		sum += test

	}

	fmt.Printf("sum: %v\n", sum)
}
