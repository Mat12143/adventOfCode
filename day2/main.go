package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isSafe(nums []int) bool {
	isIncreasing := nums[0] < nums[1]
	isDecreasing := nums[0] > nums[1]

	for i := 0; i < len(nums)-1; i++ {
		diff := math.Abs(float64(nums[i] - nums[i+1]))

		if diff < 1 || diff > 3 {
			return false
		}

		if isDecreasing && nums[i] < nums[i+1] {
			return false
		} else if isIncreasing && nums[i] > nums[i+1] {
			return false
		}

	}
	return true
}

func retryReport(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		remNums := append([]int{}, nums[:i]...)
		remNums = append(remNums, nums[i+1:]...)

		if isSafe(remNums) {
			return true
		}
	}
	return false
}

func PartTwo() int {

	file, _ := os.Open("long.txt")

	scanner := bufio.NewScanner(file)

	safeReports := 0

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Split(line, " ")
		var reports []int = make([]int, len(numbers))

		for i, n := range numbers {
			reports[i], _ = strconv.Atoi(n)
		}

		if isSafe(reports) {
			safeReports++
		} else {
			if retryReport(reports) {
				safeReports++
			}
		}
	}
	return safeReports
}

func PartOne() int {

	file, _ := os.Open("long.txt")

	scanner := bufio.NewScanner(file)

	safeReports := 0

	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Split(line, " ")
		var reports []int = make([]int, len(numbers))

		for i, n := range numbers {
			reports[i], _ = strconv.Atoi(n)
		}

		if isSafe(reports) {
			safeReports++
		}
	}
	return safeReports
}

func main() {
	fmt.Println("Part One: ", PartOne())
	fmt.Println("Part Two: ", PartTwo())
}
