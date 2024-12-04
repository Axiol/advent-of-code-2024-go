package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func readInput() ([]int, []int) {
	file, err := os.ReadFile("./input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	parsedInput := strings.Split(string(file), "\n")
	dist1 := make([]int, len(parsedInput))
	dist2 := make([]int, len(parsedInput))

	for i, v := range parsedInput {
		dist1[i], _ = strconv.Atoi(strings.Fields(v)[0])
		dist2[i], _ = strconv.Atoi(strings.Fields(v)[1])
	}

	sort.Ints(dist1)
	sort.Ints(dist2)
	return dist1, dist2
}

func part1() int {
	total := 0
	dist1, dist2 := readInput()

	for i, v := range dist1 {
		diff := 0

		if v < dist2[i] {
			diff = dist2[i] - v
		} else {
			diff = v - dist2[i]
		}

		total += diff
	}
	
	return total
}

func part2() int {
	total := 0
	dist1, dist2 := readInput()

	for _, v := range dist1 {
		repeat := 0

		for _, w := range dist2 {
			if v == w {
				repeat += 1
			}
		}

		total += v * repeat
	}

	return total
}