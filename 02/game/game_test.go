package game

import (
	"fmt"
	"os"
	"testing"

	"github.com/ravenbluedragon/aoc-2023/common"
)

var testFile string = "../../test-data/02.txt"
var testData []game

func TestMain(m *testing.M) {
	stringData := common.LoadData(testFile)
	for _, line := range stringData {
		game, err := ParseGame(line)
		if err != nil {
			fmt.Printf("Failed to parse game: %v", err)
			os.Exit(1)
		}
		testData = append(testData, game)
	}
	os.Exit(m.Run())
}

func TestNewSample(t *testing.T) {
	s := NewSample(1, 2, 3)
	if s.red != 1 || s.green != 2 || s.blue != 3 {
		t.Errorf("NewSample(1, 2, 3) = %v, expected {1 2 3}", s)
	}
}

func TestSampleContains(t *testing.T) {
	bag := NewSample(5, 5, 5)
	sample := NewSample(3, 3, 3)
	if !bag.Contains(sample) {
		t.Errorf("Expected bag %v to contain sample %v", bag, sample)
	}
}

func TestSamplePower(t *testing.T) {
	minimums := []sample{
		{4, 2, 6},
		{1, 3, 4},
		{20, 13, 6},
		{14, 3, 15},
		{6, 3, 2},
	}
	power := []int{48, 12, 1560, 630, 36}
	for i, sample := range minimums {
		if sample.Power() != power[i] {
			t.Errorf("Power() = %d, expected %d", sample.Power(), power[i])
		}
	}
}

func TestParseGame(t *testing.T) {
	line := "Game 61: 2 red, 3 green, 4 blue; 1 red, 2 green, 3 blue"
	game, err := ParseGame(line)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if game.Id != 61 || len(game.samples) != 2 {
		t.Errorf("ParseGame(%s) = %v, expected game with Id 1 and 2 samples", line, game)
	}
}

func TestGamePossible(t *testing.T) {
	bag := NewSample(12, 13, 14)
	possible := []bool{true, true, false, false, true}
	for i, game := range testData {
		if game.Possible(bag) != possible[i] {
			t.Errorf("Expected game %v posstibe to be %v", game, possible[i])
		}
	}
}

func TestGameMinimum(t *testing.T) {
	minimums := []sample{
		{4, 2, 6},
		{1, 3, 4},
		{20, 13, 6},
		{14, 3, 15},
		{6, 3, 2},
	}
	for i, game := range testData {
		if min := game.Minimum(); min != minimums[i] {
			t.Errorf("Minimum() = %v, expected %v", min, minimums[i])
		}
	}
}
