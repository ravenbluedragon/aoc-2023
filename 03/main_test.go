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
	{0, 0, 2, 467},
	{0, 5, 7, 114},
	{2, 2, 3, 35},
	{2, 6, 8, 633},
	{4, 0, 2, 617},
	{5, 7, 8, 58},
	{6, 2, 4, 592},
	{7, 6, 8, 755},
	{9, 1, 3, 664},
	{9, 5, 7, 598},
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
