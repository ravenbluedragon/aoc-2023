package main

import (
	"reflect"
	"testing"

	"github.com/ravenbluedragon/aoc-2023/common"
)

const filename = "../test-data/05.txt"

func TestSolve1(t *testing.T) {
	expected := 35
	actual := solve1(filename)

	if actual != expected {
		t.Errorf("evaluate1(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestSolve2(t *testing.T) {
	expected := 46
	actual := solve2(filename)

	if actual != expected {
		t.Errorf("evaluate2(%s) = %d, expected %d", filename, actual, expected)
	}
}

func TestParseAlmanacSeeds(t *testing.T) {
	data := common.LoadData(filename)
	almanac := parseAlmanac(data)
	expected := []int{79, 14, 55, 13}
	if len(almanac.seeds) != len(expected) {
		t.Errorf("parseAlmanac(%s) = %d seeds, expected %d", filename, len(almanac.seeds), len(expected))
	}
	for i, seed := range expected {
		if almanac.seeds[i] != seed {
			t.Errorf("parseAlmanac(%s) = %d, expected %d", filename, almanac.seeds[i], seed)
		}
	}
}

func TestParseAlmanacKeys(t *testing.T) {
	data := common.LoadData(filename)
	almanac := parseAlmanac(data)
	expectedKeys := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	if len(almanac.maps) != len(expectedKeys)-1 {
		t.Errorf("parseAlmanac(%s) = %d maps, expected %d", filename, len(almanac.maps), len(expectedKeys)-1)
	}
	if almanac.maps[0].source != expectedKeys[0] {
		t.Errorf("parseAlmanac(%s) = %s, expected %s", filename, almanac.maps[0].source, expectedKeys[0])
	}
	for i, key := range expectedKeys[1:] {
		if almanac.maps[i].destination != key {
			t.Errorf("parseAlmanac(%s) = %s, expected %s", filename, almanac.maps[i].source, key)
		}
	}
}

func TestConvertSeeds(t *testing.T) {
	data := common.LoadData(filename)
	almanac := parseAlmanac(data)
	table := []struct {
		seed     int
		location int
	}{
		{79, 82},
		{14, 43},
		{55, 86},
		{13, 35},
	}
	for _, test := range table {
		result := almanac.convert(test.seed)
		if result != test.location {
			t.Errorf("convert(%d) = %d, expected %d", test.seed, result, test.location)
		}
	}
}

func TestAlmanacSeedRanges(t *testing.T) {
	data := common.LoadData(filename)
	almanac := parseAlmanac(data)
	expected := []intRange{{start: 79, length: 14, end: 92}, {start: 55, length: 13, end: 67}}
	seedRanges := almanac.seedRanges()
	if len(seedRanges) != len(expected) {
		t.Errorf("seedRanges() = %d ranges, expected %d", len(seedRanges), len(expected))
	}
	if !reflect.DeepEqual(seedRanges, expected) {
		t.Errorf("seedRanges() = %d, expected %d", seedRanges, expected)
	}
}

func TestIntRangeSplit(t *testing.T) {
	table := []struct {
		source   intRange
		other    intRange
		expected rangeSplit
	}{
		{
			source: intRange{start: 0, length: 10, end: 9},
			other:  intRange{start: 0, length: 5, end: 4},
			expected: rangeSplit{
				intersecting:    []intRange{{start: 0, length: 5, end: 4}},
				nonIntersecting: []intRange{{start: 5, length: 5, end: 9}},
			},
		},
		{
			source: intRange{start: 0, length: 1, end: 0},
			other:  intRange{start: 10, length: 1, end: 10},
			expected: rangeSplit{
				intersecting:    nil,
				nonIntersecting: []intRange{{start: 0, length: 1, end: 0}},
			},
		},
		{
			source: intRange{start: 1, length: 100, end: 100},
			other:  intRange{start: 25, length: 50, end: 74},
			expected: rangeSplit{
				intersecting:    []intRange{{start: 25, length: 50, end: 74}},
				nonIntersecting: []intRange{{start: 1, length: 24, end: 24}, {start: 75, length: 26, end: 100}},
			},
		},
	}
	for _, test := range table {
		result := test.source.split(test.other)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("split(%d, %d) = %d, expected %d", test.source, test.other, result, test.expected)
		}
	}
}
