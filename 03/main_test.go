package main

import (
	"testing"

	"github.com/ravenbluedragon/aoc-2023/common/grid"
)

const base_data = "../test-data/"

func TestSolve1(t *testing.T) {
	filename := base_data + "03.txt"
	expected := 4361
	actual := solve1(filename)

	if actual != expected {
		t.Errorf("evaluate1(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestSolve2(t *testing.T) {
	filename := base_data + "03.txt"
	expected := 467835
	actual := solve2(filename)

	if actual != expected {
		t.Errorf("evaluate2(%s) = %d, expected %d", filename, actual, expected)
	}
}

var expectedNumbers []PartNumber = []PartNumber{
	{grid.R2c(0, 0, 2, 0), 467},
	{grid.R2c(5, 0, 7, 0), 114},
	{grid.R2c(2, 2, 3, 2), 35},
	{grid.R2c(6, 2, 8, 2), 633},
	{grid.R2c(0, 4, 2, 4), 617},
	{grid.R2c(7, 5, 8, 5), 58},
	{grid.R2c(2, 6, 4, 6), 592},
	{grid.R2c(6, 7, 8, 7), 755},
	{grid.R2c(1, 9, 3, 9), 664},
	{grid.R2c(5, 9, 7, 9), 598},
}

var exepectedSymbols map[grid.Point2D]rune = map[grid.Point2D]rune{
	grid.P2(3, 1): '*',
	grid.P2(6, 3): '#',
	grid.P2(3, 4): '*',
	grid.P2(5, 5): '+',
	grid.P2(3, 8): '$',
	grid.P2(5, 8): '*',
}

func TestLoadNumbersAndSymbols(t *testing.T) {
	filename := base_data + "03.txt"
	numbers, symbols := loadNumbersAndSymbols(filename)

	for i, n := range numbers {
		if n != expectedNumbers[i] {
			t.Errorf("numbers[%d] = %v, expected %v", i, n, expectedNumbers[i])
		}
	}

	if len(symbols) != 6 {
		t.Errorf("len(symbols) = %d, expected 9", len(symbols))
	}
	for pos := range symbols {
		if _, ok := exepectedSymbols[pos]; !ok {
			t.Errorf("symbols[%v] = false, expected true", pos)
		}
	}
}

func TestBorder(t *testing.T) {
	intersection := []bool{true, false, true, true, true, false, true, true, true, true}
	for i, n := range expectedNumbers {
		if n.BordersSymbol(exepectedSymbols) != intersection[i] {
			t.Errorf("numbers[%d].BordersSymbol() = %v, expected %v", i, n.BordersSymbol(exepectedSymbols), intersection[i])
		}
	}
}

func TestAdjacent(t *testing.T) {
	table := []struct {
		pos      grid.Point2D
		expected []int
	}{
		{grid.P2(3, 1), []int{0, 2}},
		{grid.P2(3, 4), []int{4}},
		{grid.P2(5, 8), []int{7, 9}},
	}

	for _, test := range table {
		actual := adjacent(test.pos, expectedNumbers)
		if len(actual) != len(test.expected) {
			t.Errorf("adjacent(%v) = %v, expected indexes %v", test.pos, actual, test.expected)
		}
		for i, n := range actual {
			expected := expectedNumbers[test.expected[i]]
			if n != expected {
				t.Errorf("adjacent(%v)[%d] = %v, expected %v", test.pos, i, n, expected)
			}
		}
	}
}
