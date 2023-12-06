package main

import (
	"testing"

	"github.com/ravenbluedragon/aoc-2023/common"
)

const filename = "../test-data/06.txt"

func TestSolve1(t *testing.T) {
	expected := 288
	actual := solve1(filename)

	if actual != expected {
		t.Errorf("evaluate1(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestSolve2(t *testing.T) {
	expected := 71503
	actual := solve2(filename)

	if actual != expected {
		t.Errorf("evaluate2(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestParseRaces1(t *testing.T) {
	lines := common.LoadData(filename)
	expectedRaces := []race{
		{7, 9},
		{15, 40},
		{30, 200},
	}
	actualRaces := parseRaces1(lines)
	if len(actualRaces) != len(expectedRaces) {
		t.Errorf("parseRaces1(%s) = %v; want %v", filename, actualRaces, expectedRaces)
	}
	for i := range actualRaces {
		if actualRaces[i] != expectedRaces[i] {
			t.Errorf("parseRaces1(%s) = %v; want %v", filename, actualRaces, expectedRaces)
		}
	}
}

func TestRaceWaysToWin(t *testing.T) {
	table := []struct {
		race race
		ways int
	}{
		{race{7, 9}, 4},
		{race{15, 40}, 8},
		{race{30, 200}, 9},
	}
	for _, test := range table {
		ways := test.race.waysToWin()
		if ways != test.ways {
			t.Errorf("race.waysToWin() = %d; want %d", ways, test.ways)
		}
	}
}
