package main

import (
	"fmt"
	"log"
	"os"
)

func PartOne() {

	file, err := os.ReadFile("long.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := make([]string, 0)

	fileStr := string(file)
	line := ""

	for i := range fileStr {
		if fileStr[i] != '\n' {
			line += string(fileStr[i])
		} else {
			lines = append(lines, line)
			line = ""
		}
	}

	count := 0

	for i := range lines {
		for j := 0; j < len(lines[i]); j++ {
			char := string(lines[i][j])

			if char == "X" {

				// Up check
				if string(lines[i][j-1]) == "M" && string(lines[i][j-2]) == "A" && string(lines[i][j-3]) == "S" {
					count++
				}
				// Down check
				if string(lines[i][j+1]) == "M" && string(lines[i][j+2]) == "A" && string(lines[i][j+3]) == "S" {
					count++
				}

				// Right check
				if string(lines[i+1][j]) == "M" && string(lines[i+2][j]) == "A" && string(lines[i+3][j]) == "S" {
					count++
				}

				// Left check
				if string(lines[i-1][j]) == "M" && string(lines[i-2][j]) == "A" && string(lines[i-3][j]) == "S" {
					count++
				}

				// Right up diagonal
				if string(lines[i+1][j+1]) == "M" && string(lines[i+2][j+2]) == "A" && string(lines[i+3][j+3]) == "S" {
					count++
				}

				// Left up diagonal
				if string(lines[i-1][j+1]) == "M" && string(lines[i-2][j+2]) == "A" && string(lines[i-3][j+3]) == "S" {
					count++
				}

				// Right down diagonal
				if string(lines[i-1][j-1]) == "M" && string(lines[i-2][j-2]) == "A" && string(lines[i-3][j-3]) == "S" {
					count++
				}

				// Left down diagonal
				if string(lines[i+1][j-1]) == "M" && string(lines[i+2][j-2]) == "A" && string(lines[i+3][j-3]) == "S" {
					count++
				}

			}
		}
	}

	fmt.Printf("count: %v\n", count)
}

func PartTwo() {

	file, err := os.ReadFile("long.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := make([]string, 0)

	fileStr := string(file)
	line := ""

	for i := range fileStr {
		if fileStr[i] != '\n' {
			line += string(fileStr[i])
		} else {
			lines = append(lines, line)
			line = ""
		}
	}

	count := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {

			if string(lines[i][j]) == "A" {

				if string(lines[i-1][j-1]) == "M" && string(lines[i-1][j+1]) == "M" && string(lines[i+1][j+1]) == "S" && string(lines[i+1][j-1]) == "S" {
					count++
				}

				if string(lines[i-1][j-1]) == "S" && string(lines[i-1][j+1]) == "S" && string(lines[i+1][j+1]) == "M" && string(lines[i+1][j-1]) == "M" {
					count++
				}

				if string(lines[i-1][j-1]) == "M" && string(lines[i+1][j-1]) == "M" && string(lines[i-1][j+1]) == "S" && string(lines[i+1][j+1]) == "S" {
					count++
				}

				if string(lines[i-1][j-1]) == "S" && string(lines[i+1][j-1]) == "S" && string(lines[i-1][j+1]) == "M" && string(lines[i+1][j+1]) == "M" {
					count++
				}
			}

		}
	}

	fmt.Printf("count: %v\n", count)
}

func main() {
	PartOne()
	PartTwo()
}
