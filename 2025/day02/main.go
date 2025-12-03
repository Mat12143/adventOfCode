package day02

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Mat12143/adventOfCode2025/utils"
)

func isValid(id string) bool {
	half := len(id) / 2
	return id[:half] != id[half:]
}

func isValidComplicate(id string) bool {

	for i := 1; i <= len(id)/2; i++ {
		if len(id)%i != 0 {
			continue
		}

		part := id[:i]

		repeated := strings.Repeat(part, len(id)/i)
		if repeated == id {
			return false
		}
	}
	return true
}

func Part1() {
	sum := 0

	data := utils.ReadFileByLine("day02/input.txt")
	for l := range strings.SplitSeq(data, ",") {

		l = strings.Trim(l, "\n")

		parts := strings.Split(l, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := start; i <= end; i++ {
			if !isValid(strconv.Itoa(i)) {
				sum += i
			}
		}
	}
	fmt.Printf("sum: %v\n", sum)
}

func Part2() {
	sum := 0

	data := utils.ReadFileByLine("day02/input.txt")
	for l := range strings.SplitSeq(data, ",") {

		l = strings.Trim(l, "\n")

		parts := strings.Split(l, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := start; i <= end; i++ {
			if !isValidComplicate(strconv.Itoa(i)) {
				sum += i
			}
		}

	}
	fmt.Printf("sum: %v\n", sum)
}
