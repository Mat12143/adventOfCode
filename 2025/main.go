package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Mat12143/adventOfCode2025/day01"
	"github.com/Mat12143/adventOfCode2025/day02"
	"github.com/Mat12143/adventOfCode2025/day03"
	"github.com/Mat12143/adventOfCode2025/day04"
	"github.com/Mat12143/adventOfCode2025/day05"
	"github.com/Mat12143/adventOfCode2025/day06"
	"github.com/Mat12143/adventOfCode2025/day07"
	"github.com/Mat12143/adventOfCode2025/day08"
	"github.com/Mat12143/adventOfCode2025/day09"
	"github.com/Mat12143/adventOfCode2025/day10"
	"github.com/Mat12143/adventOfCode2025/day11"
	"github.com/Mat12143/adventOfCode2025/day12"
)

type Solution func()

var solutions = map[int][]Solution{
	1:  {day01.Part1, day01.Part2},
	2:  {day02.Part1, day02.Part2},
	3:  {day03.Part1, day03.Part2},
	4:  {day04.Part1, day04.Part2},
	5:  {day05.Part1, day05.Part2},
	6:  {day06.Part1, day06.Part2},
	7:  {day07.Part1, day07.Part2},
	8:  {day08.Part1, day08.Part2},
	9:  {day09.Part1, day09.Part2},
	10: {day10.Part1, day10.Part2},
	11: {day11.Part1, day11.Part2},
	12: {day12.Part1, day12.Part2},
}

func main() {
	day := flag.Int("day", 1, "Day number [1-12]")
	part := flag.Int("part", 1, "Part number [1-2]")
	trackTime := flag.Bool("time", false, "Show the time it took to complete")
	flag.Parse()

	if *day < 1 || *day > 12 {
		fmt.Println("Please specify a valid day (1-12)")
		os.Exit(1)
	}

	if *part != 1 && *part != 2 {
		fmt.Println("Part must be 1 or 2")
		os.Exit(1)
	}

	if funcs, ok := solutions[(*day)]; ok {
		var start time.Time
		if *trackTime {
			start = time.Now()
		}
		funcs[(*part - 1)]()

		if *trackTime {
			fmt.Printf("Time: %v\n", time.Since(start))
		}
	} else {
		log.Fatal("Day not found")
	}
}
