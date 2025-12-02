package day01

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Mat12143/adventOfCode2025/utils"
)

func Part1() {
	dial := 50
	rotations := 0

	data := utils.ReadFileByLine("day01/input.txt")
	for l := range strings.SplitSeq(data, "\n") {
		if len(l) == 0 {
			continue
		}

		direction := l[:1]
		steps := l[1:]

		stepsNumber, err := strconv.Atoi(steps)
		if err != nil {
			log.Fatal(err)
		}

		if direction == "R" {
			dial += stepsNumber
		} else {
			dial -= stepsNumber
		}

		for dial < 0 || dial >= 100 {
			if dial >= 100 {
				dial -= 100
			}
			if dial < 0 {
				dial = 100 + dial
			}
		}
		if dial == 0 {
			rotations++
		}
	}

	fmt.Printf("rotations: %v\n", rotations)
}

func mod(number int) int {
	number %= 100
	if number < 0 {
		number += 100
	}
	return number
}

func Part2() {
	dial := 50
	rotations := 0

	data := utils.ReadFileByLine("day01/input.txt")
	for l := range strings.SplitSeq(data, "\n") {

		if len(l) == 0 {
			continue
		}

		direction := l[:1]
		steps := l[1:]
		stepsNumber, err := strconv.Atoi(steps)
		if err != nil {
			log.Fatal(err)
		}

		step := 1

		if direction == "L" {
			step = -1
		}

		for range stepsNumber {
			dial = mod(dial + step)
			if dial == 0 {
				rotations++
			}
		}

	}
	fmt.Printf("rotations: %v\n", rotations)

}
