package main

import (
	"strings"

	"github.com/ravenbluedragon/aoc-2023/common"
)

// boilerplate to load and solve puzzles
func main() {
	common.ChoosePart(1, solve1, solve2)
}

func solve1(filename string) any {
	data := common.LoadData(filename)
	return evaluate1(data)
}

func solve2(filename string) any {
	data := common.LoadData(filename)
	return evaluate2(data)
}

// ctoi converts a byte to an int
func ctoi(c byte) int {
	return int(c - '0')
}

var digits = "0123456789"
var words = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func evaluate1(data []string) int {
	var result int
	for _, line := range data {
		first := strings.IndexAny(line, digits)
		last := strings.LastIndexAny(line, digits)
		result += 10*ctoi(line[first]) + ctoi(line[last])
	}
	return result
}

func evaluate2(data []string) int {
	var result int
	for _, line := range data {
		first_val := findFirstDigit(line)
		last_val := findLastDigit(line)
		result += 10*first_val + last_val
	}
	return result
}

func findDigit(
	line string,
	charSearch func(line string, chars string) int,
	wordSearch func(line string, needle string) int,
	sort func(best int, current int) bool,
) int {
	val := 0
	index := charSearch(line, digits)
	if index != -1 {
		val = ctoi(line[index])
	}
	for i, digit := range words {
		if idx := wordSearch(line, digit); idx != -1 && sort(index, idx) {
			index = idx
			val = i
		}
	}
	return val
}

func findFirstDigit(line string) int {
	return findDigit(
		line,
		strings.IndexAny,
		strings.Index,
		func(best, current int) bool { return best == -1 || current < best },
	)
}

func findLastDigit(line string) int {
	return findDigit(
		line,
		strings.LastIndexAny,
		strings.LastIndex,
		func(best, current int) bool { return current > best },
	)
}
