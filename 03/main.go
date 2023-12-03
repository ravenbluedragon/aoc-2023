package main

import (
	"github.com/ravenbluedragon/aoc-2023/common"
)

// boilerplate to load and solve puzzles
func main() {
	common.ChoosePart(3, solve1, solve2)
}

func solve1(filename string) any {
	numbers, symbols := loadNumbersAndSymbols(filename)
	sum := 0
	for _, n := range numbers {
		if n.BordersSymbol(symbols) {
			sum += n.Value
		}
	}
	return sum
}

func loadNumbersAndSymbols(filename string) ([]PartNumber, map[[2]int]rune) {
	data := common.LoadData(filename)

	var numbers []PartNumber
	symbols := make(map[[2]int]rune)

	for i, line := range data {
		numbers = append(numbers, numberPositions(i, line)...)
		for pos, symb := range symbolPositions(i, line) {
			symbols[pos] = symb
		}
	}
	return numbers, symbols
}

func solve2(filename string) any {
	numbers, symbols := loadNumbersAndSymbols(filename)
	sum := 0
	for pos, symb := range symbols {
		if symb == '*' {
			adjacent := borders(pos, numbers)
			if len(adjacent) == 2 {
				sum += adjacent[0].Value * adjacent[1].Value
			}
		}
	}
	return sum
}

type PartNumber struct {
	LineNo int
	Start  int
	End    int
	Value  int
}

func (p *PartNumber) Border() [][2]int {
	var border [][2]int
	rows := []int{p.LineNo - 1, p.LineNo, p.LineNo + 1}
	for _, i := range rows {
		if i >= 0 {
			for j := p.Start - 1; j <= p.End+1; j++ {
				if j >= 0 {
					border = append(border, [2]int{i, j})
				}
			}
		}
	}
	return border
}

func borders(pos [2]int, numbers []PartNumber) []PartNumber {
	var include []PartNumber
	for _, n := range numbers {
		for _, b := range n.Border() {
			if b == pos {
				include = append(include, n)
				break
			}
		}
	}
	return include
}

func (p *PartNumber) BordersSymbol(symbols map[[2]int]rune) bool {
	for _, pos := range p.Border() {
		if _, ok := symbols[pos]; ok {
			return true
		}
	}
	return false
}

func numberPositions(lineNo int, line string) []PartNumber {
	var positions []PartNumber
	number := 0
	start := -1
	for i, c := range line {
		if c >= '0' && c <= '9' {
			if start < 0 {
				start = i
			}
			number = number*10 + int(c-'0')
		} else if start >= 0 {
			positions = append(positions, PartNumber{lineNo, start, i - 1, number})
			start = -1
			number = 0
		}
	}
	if start >= 0 {
		positions = append(positions, PartNumber{lineNo, start, len(line) - 1, number})
	}
	return positions
}

func symbolPositions(lineNo int, line string) map[[2]int]rune {
	positions := make(map[[2]int]rune)
	for i, c := range line {
		if c != '.' && (c < '0' || c > '9') {
			positions[[2]int{lineNo, i}] = c
		}
	}
	return positions
}
