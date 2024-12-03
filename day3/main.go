package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func PartOne() {

	file, _ := os.Open("long.txt")
	sc := bufio.NewScanner(file)

	sum := 0

	for sc.Scan() {

		line := sc.Text()

		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := re.FindAllStringSubmatch(line, -1)

		for _, n := range matches {
			first, _ := strconv.Atoi(n[1])
			second, _ := strconv.Atoi(n[2])

			sum += first * second
		}
	}
	fmt.Printf("sum: %v\n", sum)
}

func PartTwo() {
	file, _ := os.Open("long.txt")

	sc := bufio.NewScanner(file)

	sum := 0
	mulEnabled := true

	for sc.Scan() {

		line := sc.Text()

		mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		doRe := regexp.MustCompile(`do\(\)`)
		dontRe := regexp.MustCompile(`don't\(\)`)

		matches := mulRe.FindAllStringSubmatch(line, -1)
		matchesIndex := mulRe.FindAllStringIndex(line, -1)

		doIndex := doRe.FindAllStringIndex(line, -1)
		dontIndex := dontRe.FindAllStringIndex(line, -1)

		doI, dontI, matchI := 0, 0, 0

		for i := 0; i < len(line); i++ {
			matchId := matchesIndex[matchI][0]
			doId := doIndex[doI][0]
			dontId := dontIndex[dontI][0]

			if i == doId {
				mulEnabled = true
				if len(doIndex)-1 != doI {
					doI++
				}
			}
			if i == dontId {
				mulEnabled = false
				if len(dontIndex)-1 != dontI {
					dontI++
				}
			}

			if i == matchId {
				if mulEnabled {
					first, _ := strconv.Atoi(matches[matchI][1])
					second, _ := strconv.Atoi(matches[matchI][2])

					sum += first * second
				}
				if len(matchesIndex)-1 != matchI {
					matchI++
				}
			}
		}
	}

	fmt.Printf("sum: %v\n", sum)
}

func main() {
	PartOne()
	PartTwo()
}
