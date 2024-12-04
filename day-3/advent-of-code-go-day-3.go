package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func readInput() ([]string) {
	file, err := os.ReadFile("./input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	parsedInput := strings.Split(string(file), "\n")
  re := regexp.MustCompile(`mul\(\d{1,3}\+\d{1,3}\)`)

	var filteredInput []string
	for _, line := range parsedInput {
		if re.MatchString(line) {
			filteredInput = append(filteredInput, line)
		}
	}

	log.Printf("%v", filteredInput)

	return filteredInput
}

func part1() (int) {
	operations := readInput()

	return len(operations)
}

func part2() (int) {
	return 0
}