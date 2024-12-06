package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

const FIND_OPS_PART1 = `(mul\(\d+,\d+\))`
const FIND_OPS_PART2 = `((?:don't|do|mul)\(\d*,?\d*\))`
const PARSE_MULTIPLIER = `mul\((\d+),(\d+)\)`

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func readInput() (*os.File) {
	file, err := os.Open("./input.txt")
	
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func atoiSlice(d []string) []int {
	intSlice := make([]int, len(d))
	for i, e := range d {
		convertedElement, err := strconv.Atoi(e)
		if err != nil {
			slog.Error("error converting string to int", "error", err)
		}
		intSlice[i] = convertedElement
	}
	return intSlice
}

func findOps(inputReader io.Reader, reString string) []string {
	scanner := bufio.NewScanner(inputReader)
	re := regexp.MustCompile(reString)
	ops := make([]string, 0, 10)

	for scanner.Scan() {
		line := scanner.Text()
		ops = append(ops, re.FindAllString(line, -1)...)
	}

	return ops
}

func parseMultiply(ops []string, reString string) [][]int {
	parsedOps := make([][]int, 0, len(ops))
	re := regexp.MustCompile(reString)

	for _, op := range ops {
		matches := re.FindStringSubmatch(op)
		matchesInt := atoiSlice(matches[1:])
		parsedOps = append(parsedOps, matchesInt)
	}

	return parsedOps
}

func multiplySum(ops [][]int) (sum int) {
	for _, op := range ops {
		sum += multiply(op)
	}
	return
}

func multiply(op []int) (result int) {
	result = 1
	for _, op := range op {
		result *= op
	}
	return result
}

func part1() (int) {
	inputReader := readInput()
	foundOps := findOps(inputReader, FIND_OPS_PART1)
	multiplyOps := parseMultiply(foundOps, PARSE_MULTIPLIER)
	sum := multiplySum(multiplyOps)

	return sum
}

func part2() (int) {
	inputReader := readInput()
	defer inputReader.Close()

	foundOps := findOps(inputReader, FIND_OPS_PART2)

	enabledOps := make([]string, 0, len(foundOps)/2)
	enabled := true
	for _, op := range foundOps {
		if op == "do()" {
			enabled = true
		} else if op == "don't()" {
			enabled = false
		} else {
			if enabled {
				enabledOps = append(enabledOps, op)
			}
		}
	}

	multiplyOps := parseMultiply(enabledOps, PARSE_MULTIPLIER)
	sum := multiplySum(multiplyOps)

	return sum
}