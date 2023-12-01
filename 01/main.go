package main

import (
	"fmt"
	"strings"

	"github.com/ravenbluedragon/aoc-2023/common"
)

func main() {
	data := common.LoadData("data/01.txt")
	fmt.Printf("Part 1 result: %d\n", evaluate1(data))
	fmt.Printf("Part 2 result: %d\n", evaluate2(data))
}

func evaluate1(data []string) int {
	var result int
	for _, line := range data {
		first := strings.IndexAny(line, "0123456789")
		last := strings.LastIndexAny(line, "0123456789")
		result += 10*int(line[first]-'0') + int(line[last]-'0')
	}
	return result
}

var digits = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func evaluate2(data []string) int {
	var result int
	for _, line := range data {
		first_val := 0
		// get first digit
		first := strings.IndexAny(line, "0123456789")
		if first != -1 {
			first_val = int(line[first] - '0')
		}
		for i, digit := range digits {
			if index := strings.Index(line, digit); index != -1 && (index < first || first == -1) {
				first = index
				first_val = i
			}
		}

		// get last digit
		last_val := 0
		last := strings.LastIndexAny(line, "0123456789")
		if last != -1 {
			last_val = int(line[last] - '0')
		}
		for i, digit := range digits {
			if index := strings.LastIndex(line, digit); index != -1 && index > last {
				last = index
				last_val = i
			}
		}

		result += 10*first_val + last_val
	}
	return result
}
