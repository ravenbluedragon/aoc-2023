package main

import (
	"testing"

	"github.com/ravenbluedragon/aoc-2023/common"
)

func TestEvaluate1(t *testing.T) {
	data := []struct {
		input    []string
		expected int
	}{
		{[]string{"abc123def", "456ghi789"}, 13 + 49},
		{[]string{"xyz987uvw", "654jkl321"}, 97 + 61},
		{[]string{"1", "123", "a1a"}, 11 + 13 + 11},
	}

	for _, tc := range data {
		result := evaluate1(tc.input)
		if result != tc.expected {
			t.Errorf("evaluate1(%v) = %d, expected %d", tc.input, result, tc.expected)
		}
	}
}

func TestPart1(t *testing.T) {
	filename := "../data/01-test1.txt"
	expected := 142
	actual := evaluate1(common.LoadData(filename))

	if actual != expected {
		t.Errorf("evaluate1(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestEvaluate2(t *testing.T) {
	tests := []struct {
		lines    []string
		expected int
	}{
		{[]string{"abc123def", "456ghi789"}, 13 + 49},
		{[]string{"eightxyz987uvw", "654jkl321one"}, 87 + 61},
		{[]string{"1", "123", "a1a", "three"}, 11 + 13 + 11 + 33},
	}

	for _, tc := range tests {
		result := evaluate2(tc.lines)
		if result != tc.expected {
			t.Errorf("evaluate2(%v) = %d, expected %d", tc.lines, result, tc.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	filename := "../data/01-test2.txt"
	expected := 281
	actual := evaluate2(common.LoadData(filename))

	if actual != expected {
		t.Errorf("evaluate2(%s) = %d, expected %d", filename, actual, expected)
	}
}
