package main

import (
	"fmt"
	"log"

	"github.com/ravenbluedragon/aoc-2023/common"
)

// boilerplate to load and solve puzzles
func main() {
	common.ChoosePart(8, solve1, solve2)
}

func solve1(filename string) any {
	lines := common.LoadData(filename)
	steps := 0
	node := "AAA"
	instructions, nodes := readData(lines)
	for node != "ZZZ" {
		if steps < 0 {
			log.Fatal("Too many steps")
		}
		i := steps % len(instructions)
		j := 0
		if instructions[i] == 'R' {
			j = 1
		}
		node = nodes[node][j]
		// fmt.Printf("%d %s %d\n", j, node, steps)
		steps++
	}
	return steps
}

func solve2(filename string) any {
	lines := common.LoadData(filename)
	instructions, nodes := readData(lines)

	type cycle struct {
		initial string
		final   string
		steps   int
	}

	var cycles []cycle
	for key := range nodes {
		if key[2] == 'A' {
			node := key
			steps := 0
			for node[2] != 'Z' {
				i := steps % len(instructions)
				j := 0
				if instructions[i] == 'R' {
					j = 1
				}
				node = nodes[node][j]
				steps++
			}
			cycles = append(cycles, cycle{initial: key, final: node, steps: steps})
		}
	}
	fmt.Printf("Cycles: %v\n", cycles)

	var lengths []int
	for _, cycle := range cycles {
		lengths = append(lengths, cycle.steps)
	}
	return lcm(lengths)
}

func readData(lines []string) (string, map[string][2]string) {
	nodes := make(map[string][2]string)
	for _, line := range lines[2:] {
		// "AAA = (BBB, CCC)"
		key := line[0:3]
		left := line[7:10]
		right := line[12:15]
		nodes[key] = [2]string{left, right}
	}
	return lines[0], nodes
}

func lcm(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = lcm2(result, number)
		if result < 0 {
			log.Fatalf("Overflow: %d", result)
		}
	}
	return result
}

func lcm2(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
