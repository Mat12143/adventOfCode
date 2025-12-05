package day05

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Interval struct {
	start, end int
}

func binSearch(intervals []Interval, n int, start int, end int) bool {
	if start > end {
		return false
	}

	mid := (start + end) / 2

	if n < intervals[mid].start {
		return binSearch(intervals, n, start, mid-1)
	} else if n > intervals[mid].end {
		return binSearch(intervals, n, mid+1, end)
	} else {
		return true
	}
}

func mergeIntervals(intervals []Interval) []Interval {
	merged := []Interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		last := merged[len(merged)-1]

		if curr.start <= last.end {
			if curr.end > last.end {
				merged[len(merged)-1].end = curr.end
			}
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}

func Part1(data string) {

	intervals := []Interval{}
	numbers := []int{}

	for l := range strings.SplitSeq(data, "\n") {
		if len(l) == 0 {
			continue
		}

		if strings.Contains(l, "-") {
			interval := strings.Split(l, "-")
			start, err := strconv.Atoi(interval[0])
			if err != nil {
				log.Fatal(err)
			}
			end, err := strconv.Atoi(interval[1])
			if err != nil {
				log.Fatal(err)
			}

			intervals = append(intervals, Interval{
				start, end,
			})
		} else {
			number, _ := strconv.Atoi(l)
			numbers = append(numbers, number)
		}
	}

	slices.SortFunc(intervals, func(a Interval, b Interval) int {
		return a.start - b.start
	})

	intervals = mergeIntervals(intervals)

	tot := 0

	for _, n := range numbers {
		if binSearch(intervals, n, 0, len(intervals)-1) {
			tot++
		}

	}

	fmt.Printf("tot: %v\n", tot)

}

func Part2(data string) {
	intervals := []Interval{}
	numbers := []int{}

	for l := range strings.SplitSeq(data, "\n") {
		if len(l) == 0 {
			continue
		}

		if strings.Contains(l, "-") {
			interval := strings.Split(l, "-")
			start, err := strconv.Atoi(interval[0])
			if err != nil {
				log.Fatal(err)
			}
			end, err := strconv.Atoi(interval[1])
			if err != nil {
				log.Fatal(err)
			}

			intervals = append(intervals, Interval{
				start, end,
			})
		} else {
			number, _ := strconv.Atoi(l)
			numbers = append(numbers, number)
		}
	}

	slices.SortFunc(intervals, func(a Interval, b Interval) int {
		return a.start - b.start
	})

	intervals = mergeIntervals(intervals)

	tot := 0

	for _, in := range intervals {
		tot += (in.end - in.start) + 1
	}

	fmt.Printf("tot: %v\n", tot)
}
