package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func extractNums(line string) []int {

	intSlice := []int{}

	numbersStr := strings.Split(line, ",")

	for _, ns := range numbersStr {
		n, _ := strconv.Atoi(ns)
		intSlice = append(intSlice, n)
	}

	return intSlice
}

func checkForBefore(index int, num int, slice []int) (bool, int) {

	for i := index; i >= 0; i-- {
		if num == slice[i] {
			return true, i
		}
	}

	return false, -1
}

func isPageValid(numbers []int, rules [][]int) (int, bool) {

	for i, num := range numbers {
		for _, rule := range rules {

			if num == rule[0] {
				if v, _ := checkForBefore(i, rule[1], numbers); v {
					return 0, false
				}
			}

		}
	}

	return numbers[len(numbers)/2], true
}

func getNumIndex(num int, numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		if num == numbers[i] {
			return i
		}
	}
	return 0
}

func reOrder(numbers []int, rules [][]int) []int {
	for i, num := range numbers {
		for _, rule := range rules {

			if num == rule[0] {
				v, index := checkForBefore(i, rule[1], numbers)

				if v {
					numbers[getNumIndex(num, numbers)] = numbers[index]
					numbers[index] = num
				}
			}
		}
	}
	return numbers
}

func PartOne() {

	file, _ := os.Open("long.txt")

	sc := bufio.NewScanner(file)

	rules := [][]int{}
	isOnRules := true
	sum := 0

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			isOnRules = false
		}

		if isOnRules {
			numbers := strings.Split(line, "|")

			first, err := strconv.Atoi(numbers[0])
			second, err := strconv.Atoi(numbers[1])

			if err != nil {
				log.Fatal(err)
			}

			rules = append(rules, []int{first, second})
		} else if line != "" && !isOnRules {

			numbers := extractNums(line)
			middle, valid := isPageValid(numbers, rules)

			if valid {
				sum += middle
			}
		}
	}

	fmt.Printf("sum: %v\n", sum)

}

func PartTwo() {

	file, _ := os.Open("long.txt")

	sc := bufio.NewScanner(file)

	rules := [][]int{}
	isOnRules := true
	sum := 0

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			isOnRules = false
		}

		if isOnRules {
			numbers := strings.Split(line, "|")

			first, err := strconv.Atoi(numbers[0])
			second, err := strconv.Atoi(numbers[1])

			if err != nil {
				log.Fatal(err)
			}

			rules = append(rules, []int{first, second})
		} else if line != "" && !isOnRules {

			numbers := extractNums(line)
			middle, valid := isPageValid(numbers, rules)

			if !valid {
				for !valid {
					numbers = reOrder(numbers, rules)
					middle, valid = isPageValid(numbers, rules)
				}
				sum += middle

			}
		}
	}

	fmt.Printf("sum: %v\n", sum)

}

func main() {
	PartOne()
	PartTwo()
}
