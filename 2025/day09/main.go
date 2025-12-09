package day09

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func getRectArea(a, b Point) int {
	return (int(math.Abs(float64(a.x-b.x))) + 1) * (int(math.Abs(float64(a.y-b.y))) + 1)
}

func Part1(data string) {

	points := []Point{}

	for l := range strings.SplitSeq(data, "\n") {

		if len(l) == 0 {
			continue
		}

		numbers := strings.Split(l, ",")
		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])

		points = append(points, Point{x, y})
	}

	maxArea := 0

	for _, p := range points {
		for _, p2 := range points {
			if p == p2 {
				continue
			}
			maxArea = max(maxArea, getRectArea(p, p2))
		}
	}

	fmt.Printf("maxArea: %v\n", maxArea)

}

func Part2(data string) {
	// TODO: implement Part2
}
