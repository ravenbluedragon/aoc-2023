package main

import (
	"log"

	"github.com/ravenbluedragon/aoc-2023/common"
	"github.com/ravenbluedragon/aoc-2023/common/numeric"
)

// boilerplate to load and solve puzzles
func main() {
	common.ChoosePart(8, solve1, solve2)
}

func solve1(filename string) any {
	lines := common.LoadData(filename)
	instructions, nodes := readData(lines)

	// NOTE the input is set up so AAA -> ZZZ
	return stepsToZ("AAA", instructions, nodes)

}

func solve2(filename string) any {
	lines := common.LoadData(filename)
	instructions, nodes := readData(lines)

	var cycles []int
	for node := range nodes {
		if node[2] == 'A' {
			steps := stepsToZ(node, instructions, nodes)
			cycles = append(cycles, steps)
		}
	}

	return numeric.Lcm(cycles)
}

// type Node represents a path with a left and right node
type Node struct {
	left  string
	right string
}

// readData parses the input file into a map of nodes
func readData(lines []string) (string, map[string]Node) {
	nodes := make(map[string]Node)
	for _, line := range lines[2:] {
		// "AAA = (BBB, CCC)"
		key := line[0:3]
		left := line[7:10]
		right := line[12:15]
		nodes[key] = Node{left, right}
	}
	return lines[0], nodes
}

// stepsToZ returns the number of steps to get to the node ??Z
func stepsToZ(key string, instructions string, nodes map[string]Node) int {
	steps := 0
	for key[2] != 'Z' {
		if steps < 0 {
			log.Fatal("Overflow on steps")
		}
		node := nodes[key]

		i := steps % len(instructions)
		switch instructions[i] {
		case 'L':
			key = node.left
		case 'R':
			key = node.right
		}
		steps++
	}
	return steps
}
