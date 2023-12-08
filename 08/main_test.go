package main

import (
	"testing"

	"github.com/ravenbluedragon/aoc-2023/common"
)

const base_folder = "../test-data/"

func TestSolve1(t *testing.T) {
	table := []struct {
		filename string
		expected int
	}{
		{"08-1.txt", 2},
		{"08-2.txt", 6},
	}

	for _, test := range table {
		actual := solve1(base_folder + test.filename)
		if actual != test.expected {
			t.Errorf("evaluate1(%s) = %d, expected %d", test.filename, actual, test.expected)
		}
	}
}

func TestSolve2(t *testing.T) {
	filename := base_folder + "08-3.txt"
	expected := 6
	actual := solve2(filename)

	if actual != expected {
		t.Errorf("evaluate2(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestReadMap(t *testing.T) {
	filename := base_folder + "08-1.txt"
	expected := map[string][2]string{
		"AAA": {"BBB", "CCC"},
		"BBB": {"DDD", "EEE"},
		"CCC": {"ZZZ", "GGG"},
		"DDD": {"DDD", "DDD"},
		"EEE": {"EEE", "EEE"},
		"GGG": {"GGG", "GGG"},
		"ZZZ": {"ZZZ", "ZZZ"},
	}
	instructions, nodes := readData(common.LoadData(filename))
	if instructions != "RL" {
		t.Errorf("readData(%s) = %s, expected AAA", filename, instructions)
	}
	for key, value := range expected {
		if nodes[key] != value {
			t.Errorf("readData(%s) = %s, expected %s", filename, nodes[key], value)
		}
	}

}
