package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("long.txt")

	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}
	similarity := 0

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "   ")

		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	for i := 0; i < len(left); i++ {
		appears := 0

		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				appears++
			}
		}
		similarity += left[i] * appears
	}

	fmt.Printf("similarity: %v\n", similarity)
}
