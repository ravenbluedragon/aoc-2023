package main

import (
	"testing"
)

const base_data = "../test-data/"

func TestSolve1(t *testing.T) {
	filename := base_data + "05.txt"
	expected := 35
	actual := solve1(filename)

	if actual != expected {
		t.Errorf("evaluate1(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestSolve2(t *testing.T) {
	filename := base_data + "05.txt"
	expected := 46
	actual := solve2(filename)

	if actual != expected {
		t.Errorf("evaluate2(%s) = %d, expected %d", filename, actual, expected)
	}
}
