# Advent of Code Solutions
My solutions for [Advent of Code](https://adventofcode.com/) 2024 and 2025, written in Go (day01 2024 in C++)

## 2025 Solutions
For the 2025 solutions, I built a simple CLI to run specific days and parts with optional performance benchmarking

### Usage

```bash
go run . <day> <part> [--avg]
```

**Arguments:**
- `<day>`: The day number (1-12)
- `<part>`: The part number (1 or 2)
- `--avg`: (Optional) Runs the solution 100 times and calculates the average execution time ± standard error

**Examples:**
```bash
# Run day 5, part 1
go run . 5 1

# Run day 12, part 2 with performance metrics
go run . 12 2 --avg
```

## 2024 Solutions
The 2024 solutions are organized by day and can be run individually from their respective directories.
