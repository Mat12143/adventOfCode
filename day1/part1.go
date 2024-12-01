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
	differences := 0

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[3])

		left = append(left, leftNum)
		right = append(right, rightNum)

	}

	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = diff * -1
		}

		differences += diff
	}

	fmt.Printf("differences: %v\n", differences)

}
