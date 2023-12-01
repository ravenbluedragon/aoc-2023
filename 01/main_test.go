package main

import (
	"testing"
)

const base_data = "../test-data/"

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

func TestSolve1(t *testing.T) {
	filename := base_data + "01-1.txt"
	expected := 142
	actual := solve1(filename)

	if actual != expected {
		t.Errorf("evaluate1(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestFindFirstDigit(t *testing.T) {
	tests := []struct {
		line     string
		expected int
	}{
		{"abc123def", 1},
		{"xyz987uvw", 9},
		{"1", 1},
		{"123", 1},
		{"a1a", 1},
		{"three", 3},
		{"twoseven", 2},
	}

	for _, tc := range tests {
		result := findFirstDigit(tc.line)
		if result != tc.expected {
			t.Errorf("findFirstDigit(%s) = %d, expected %d", tc.line, result, tc.expected)
		}
	}
}

func TestFindLastDigit(t *testing.T) {
	tests := []struct {
		line     string
		expected int
	}{
		{"abc123def", 3},
		{"xyz987uvw", 7},
		{"1", 1},
		{"123", 3},
		{"a1a", 1},
		{"three", 3},
		{"twoseven", 7},
	}

	for _, tc := range tests {
		result := findLastDigit(tc.line)
		if result != tc.expected {
			t.Errorf("findLastDigit(%s) = %d, expected %d", tc.line, result, tc.expected)
		}
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

func TestSolve2(t *testing.T) {
	filename := base_data + "01-2.txt"
	expected := 281
	actual := solve2(filename)

	if actual != expected {
		t.Errorf("evaluate2(%s) = %d, expected %d", filename, actual, expected)
	}
}
