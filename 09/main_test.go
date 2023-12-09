package main

import (
	"testing"

	"github.com/ravenbluedragon/aoc-2023/common"
)

const filename = "../test-data/09.txt"

func TestSolve1(t *testing.T) {
	expected := 114
	actual := solve1(filename)

	if actual != expected {
		t.Errorf("evaluate1(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestSolve2(t *testing.T) {
	expected := 2
	actual := solve2(filename)

	if actual != expected {
		t.Errorf("evaluate2(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestExtrapolateNext(t *testing.T) {
	expected := []int{18, 28, 68}
	lines := common.LoadData(filename)
	data := parseInput(lines)
	for i, seq := range data {
		actual := extrapolateNext(seq)
		if actual != expected[i] {
			t.Errorf("extrapolateNext(%v) = %d, expected %d", seq, actual, expected[i])
		}
	}
}

func TestExtrapolatePrev(t *testing.T) {
	expected := []int{-3, 0, 5}
	lines := common.LoadData(filename)
	data := parseInput(lines)
	for i, seq := range data {
		actual := extrapolatePrev(seq)
		if actual != expected[i] {
			t.Errorf("extrapolatePrev(%v) = %d, expected %d", seq, actual, expected[i])
		}
	}
}
