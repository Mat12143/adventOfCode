package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/Mat12143/adventOfCode/day01"
	"github.com/Mat12143/adventOfCode/day02"
	"github.com/Mat12143/adventOfCode/day03"
	"github.com/Mat12143/adventOfCode/day04"
	"github.com/Mat12143/adventOfCode/day05"
	"github.com/Mat12143/adventOfCode/day06"
	"github.com/Mat12143/adventOfCode/day07"
	"github.com/Mat12143/adventOfCode/day08"
	"github.com/Mat12143/adventOfCode/day09"
	"github.com/Mat12143/adventOfCode/day10"
	"github.com/Mat12143/adventOfCode/day11"
	"github.com/Mat12143/adventOfCode/day12"
	"github.com/Mat12143/adventOfCode/utils"
	"github.com/urfave/cli/v3"
)

type Solution func(string)

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

func StdDev(durations []time.Duration, mean time.Duration) time.Duration {
	var sum float64
	for _, d := range durations {
		diff := float64(d - mean)
		sum += diff * diff
	}
	variance := sum / float64(len(durations))
	return time.Duration(math.Sqrt(variance))
}

func Avg(durations []time.Duration) time.Duration {
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations))
}

func main() {

	cmd := &cli.Command{
		Name:  "exec",
		Usage: "",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "avg",
				Value: false,
				Usage: "Track an avg time",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {

			dayString := cmd.Args().Get(0)
			partString := cmd.Args().Get(1)

			day, err := strconv.Atoi(dayString)
			if err != nil {
				return err
			}

			part, err := strconv.Atoi(partString)
			if err != nil {
				return err
			}

			if p, ok := solutions[day]; ok {
				if part < 1 || part > 2 {
					return errors.New("Invalid part given")
				}

				inputData := utils.ReadFile(fmt.Sprintf("day%02d/input.txt", day))

				if cmd.Bool("avg") {

					avgs := []time.Duration{}

					orig := os.Stdout
					devNull, _ := os.Open(os.DevNull)
					os.Stdout = devNull

					for range 100 {

						start := time.Now()
						p[part-1](inputData)
						avgs = append(avgs, time.Since(start))
					}

					total := time.Duration(0)
					for _, v := range avgs {
						total += v
					}

					os.Stdout = orig

					avg := Avg(avgs)

					fmt.Printf("Avg: %v ± %v\n", avg, StdDev(avgs, avg))

					return nil

				}

				start := time.Now()
				p[part-1](inputData)
				fmt.Printf("Took %v\n", time.Since(start))

			} else {
				return errors.New("Invalid day given")
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
