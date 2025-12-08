package day08

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Box struct {
	x, y, z int
}

type JunctionBox struct {
	boxes []Box
}

type Distance struct {
	a, b     Box
	distance int
}

func getDistance(a, b Box) int {
	dx := b.x - a.x
	dy := b.y - a.y
	dz := b.z - a.z
	return dx*dx + dy*dy + dz*dz
}

func addToCircuit(circuits [][]Box, a, b Box) [][]Box {
	// Get all the circuit where a or b is present
	var indexes []int
	for i, c := range circuits {
		if slices.Contains(c, a) || slices.Contains(c, b) {
			indexes = append(indexes, i)
		}
	}

	// No circuits -> New one
	if len(indexes) == 0 {
		return append(circuits, []Box{a, b})
	}

	// No multiple circuits with a or b
	if len(indexes) == 1 {
		i := indexes[0]
		if !slices.Contains(circuits[i], a) {
			circuits[i] = append(circuits[i], a)
		}
		if !slices.Contains(circuits[i], b) {
			circuits[i] = append(circuits[i], b)
		}
		return circuits
	}

	// Create a new circuit to merge the 2 circuits that overlaps
	newCircuit := []Box{}
	for _, i := range indexes {
		newCircuit = append(newCircuit, circuits[i]...)
	}

	// Add a and b, if missing in the circuit
	if !slices.Contains(newCircuit, a) {
		newCircuit = append(newCircuit, a)
	}
	if !slices.Contains(newCircuit, b) {
		newCircuit = append(newCircuit, b)
	}

	// Filter circuits in order to avoid adding the two merged one
	newCirc := [][]Box{}
	for i, c := range circuits {
		if !slices.Contains(indexes, i) {
			newCirc = append(newCirc, c)
		}
	}

	// Add new merged circuit
	return append(newCirc, newCircuit)
}

func Part1(data string) {

	boxes := []Box{}
	distances := []Distance{}
	circuits := [][]Box{}

	for l := range strings.SplitSeq(data, "\n") {

		if len(l) == 0 {
			continue
		}

		coordinates := strings.Split(l, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		z, _ := strconv.Atoi(coordinates[2])

		boxes = append(boxes, Box{x, y, z})
	}

	for i := range boxes {
		for j := range boxes {
			// Remove reverse pair
			if j <= i {
				continue
			}

			a := boxes[i]
			b := boxes[j]

			dis := getDistance(a, b)
			if dis > 0 {
				distances = append(distances, Distance{a, b, dis})
			}
		}
	}

	slices.SortFunc(distances, func(a Distance, b Distance) int {
		return int(a.distance - b.distance)
	})

	for i := range 1000 {
		curr := distances[i]
		circuits = addToCircuit(circuits, curr.a, curr.b)
	}

	tot := 1

	slices.SortFunc(circuits, func(a, b []Box) int {
		return len(b) - len(a)
	})

	for i := range 3 {
		tot *= len(circuits[i])
	}

	fmt.Printf("tot: %v\n", tot)
}

func Part2(data string) {
}
