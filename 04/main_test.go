package main

import (
	"testing"
)

const base_data = "../test-data/"

func TestSolve1(t *testing.T) {
	filename := base_data + "04.txt"
	expected := 13
	actual := solve1(filename)

	if actual != expected {
		t.Errorf("evaluate1(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestSolve2(t *testing.T) {
	filename := base_data + "04.txt"
	expected := 30
	actual := solve2(filename)

	if actual != expected {
		t.Errorf("evaluate2(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestCardMatches(t *testing.T) {
	tests := []struct {
		name     string
		winning  []int
		player   []int
		expected int
	}{
		{
			name:     "No matches",
			winning:  []int{1, 2, 3, 4, 5},
			player:   []int{6, 7, 8, 9, 10},
			expected: 0,
		},
		{
			name:     "Some matches",
			winning:  []int{1, 2, 3, 4, 5},
			player:   []int{1, 7, 3, 9, 10},
			expected: 2,
		},
		{
			name:     "All matches",
			winning:  []int{1, 2, 3, 4, 5},
			player:   []int{1, 2, 3, 4, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card{
				winning: tt.winning,
				player:  tt.player,
			}
			actual := c.matches()
			if actual != tt.expected {
				t.Errorf("card.matches() = %d, expected %d", actual, tt.expected)
			}
		})
	}
}
