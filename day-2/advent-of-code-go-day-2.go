package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func readInput() ([][]int) {
	file, err := os.ReadFile("./input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	parsedInput := strings.Split(string(file), "\n")
	reports := make([][]int, len(parsedInput))

	for i, v := range parsedInput {
		reportSplit := strings.Split(v, " ")
		reportLine := make([]int, len(reportSplit))

		for j, w := range reportSplit {
			reportLine[j], _ = strconv.Atoi(w)
		}

		reports[i] = reportLine
	}

	return reports
}

func isSafe(line []int) (bool) {
	safe := true
	increasing := true
	decreasing := true

	for i, v := range line {
		if i + 1 == len(line) {
			break
		}

		if v < line[i+1] {
			decreasing = false
		} else if v > line[i+1] {
			increasing = false
		}

		if !increasing && !decreasing {
			safe = false
			break
		}

		if math.Abs(float64(v - line[i+1])) > 3 || math.Abs(float64(v - line[i+1])) == 0 {
			safe = false
		}
	}

	log.Printf("%d : %t", line, safe)


	return safe
}

func part1() int {
	reports := readInput()
	total:= 0

	for _, v := range reports {
		if isSafe(v) {
			total += 1
		}
	}

	return total
}

func part2() int {
	reports := readInput()
	total:= 0

	for _, v := range reports {
		if isSafe(v) {
			total += 1
		} else {
			for i := range v {
				newV := make([]int, len(v)-1)
				copy(newV, v[:i])
				copy(newV[i:], v[i+1:])
				if isSafe(newV) {
						total += 1
						break
				}
			}
		}
	}

	return total
}