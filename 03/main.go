package main

import (
	"github.com/ravenbluedragon/aoc-2023/common"
	"github.com/ravenbluedragon/aoc-2023/common/grid"
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

func solve2(filename string) any {
	numbers, symbols := loadNumbersAndSymbols(filename)
	sum := 0
	for pos, symb := range symbols {
		if symb == '*' {
			adj := adjacent(pos, numbers)
			if len(adj) == 2 {
				sum += adj[0].Value * adj[1].Value
			}
		}
	}
	return sum
}

// loadNumbersAndSymbols loads the positions and values of numbers and symbols from the file
func loadNumbersAndSymbols(filename string) ([]PartNumber, map[grid.Point2D]rune) {
	data := common.LoadData(filename)

	var numbers []PartNumber
	symbols := make(map[grid.Point2D]rune)

	for i, line := range data {
		numbers = append(numbers, numberPositions(i, line)...)
		for pos, symb := range symbolPositions(i, line) {
			symbols[pos] = symb
		}
	}
	return numbers, symbols
}

// PartNumber stores the position and value of a number in the grid
type PartNumber struct {
	LineNo int
	Start  int
	End    int
	Value  int
}

// Border returns the positions adjacent to the number
func (p *PartNumber) Border() []grid.Point2D {
	var border []grid.Point2D
	rows := []int{p.LineNo - 1, p.LineNo, p.LineNo + 1}
	for _, j := range rows {
		if j >= 0 {
			for i := p.Start - 1; i <= p.End+1; i++ {
				if j >= 0 {
					border = append(border, grid.P2(i, j))
				}
			}
		}
	}
	return border
}

// adjacent returns the numbers adjacent to the given position
func adjacent(pos grid.Point2D, numbers []PartNumber) []PartNumber {
	var include []PartNumber
	for _, n := range numbers {
		if (n.Start-1 <= pos.X && pos.X <= n.End+1) && (n.LineNo-1 <= pos.Y && pos.Y <= n.LineNo+1) {
			include = append(include, n)
		}
	}
	return include
}

// BordersSymbol returns true if the number borders a symbol
func (p *PartNumber) BordersSymbol(symbols map[grid.Point2D]rune) bool {
	for _, pos := range p.Border() {
		if _, ok := symbols[pos]; ok {
			return true
		}
	}
	return false
}

// numberPositions returns the positions of numbers in the line
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

// symbolPositions returns the positions of symbols in the line
func symbolPositions(lineNo int, line string) map[grid.Point2D]rune {
	positions := make(map[grid.Point2D]rune)
	for i, c := range line {
		if c != '.' && (c < '0' || c > '9') {
			positions[grid.P2(i, lineNo)] = c
		}
	}
	return positions
}