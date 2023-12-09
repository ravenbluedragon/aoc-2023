package main

import (
	"log"

	"github.com/ravenbluedragon/aoc-2023/common"
	"github.com/ravenbluedragon/aoc-2023/common/parse"
)

// boilerplate to load and solve puzzles
func main() {
	common.ChoosePart(9, solve1, solve2)
}

func solve1(filename string) any {
	lines := common.LoadData(filename)
	data := parseInput(lines)
	sum := 0
	for _, seq := range data {
		sum += extrapolateNext(seq)
	}
	return sum
}

func solve2(filename string) any {
	lines := common.LoadData(filename)
	data := parseInput(lines)
	sum := 0
	for _, seq := range data {
		sum += extrapolatePrev(seq)
	}
	return sum
}

// parseInput converts a slice of strings to a slice of int slices
func parseInput(lines []string) [][]int {
	var result [][]int
	for _, line := range lines {
		ints, err := parse.StringToIntList(line, " ")
		if err != nil {
			log.Fatalf("Failed to parse line: %s", line)
		}
		result = append(result, ints)
	}
	return result
}

// allZero returns true if all values in the slice are zero or if the slice is empty
func allZero(seq []int) bool {
	for _, i := range seq {
		if i != 0 {
			return false
		}
	}
	return true
}

// finiteDifferences returns a slice of the finite differences of the input slice
func finiteDifferences(seq []int) [][]int {
	var result [][]int
	result = append(result, seq)
	prev := seq
	for !allZero(prev) {
		current := make([]int, len(prev)-1)
		for i, v := range prev[1:] {
			current[i] = v - prev[i]
		}
		result = append(result, current)
		prev = current
	}
	return result
}

// extrapolateNext extrapolates the next value in the sequence based on successive finite differences
func extrapolateNext(seq []int) int {
	sequences := finiteDifferences(seq)
	diff := 0
	for i := len(sequences) - 1; i >= 0; i-- {
		diff += last(sequences[i])
	}
	return diff
}

// last returns the last value in the slice
func last(seq []int) int {
	return seq[len(seq)-1]
}

// extrapolatePrev extrapolates the previous value in the sequence based on successive finite differences
func extrapolatePrev(seq []int) int {
	sequences := finiteDifferences(seq)
	diff := 0
	for i := len(sequences) - 1; i >= 0; i-- {
		diff = sequences[i][0] - diff
	}
	return diff
}
