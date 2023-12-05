package main

import (
	"log"
	"math"
	"strings"

	"github.com/ravenbluedragon/aoc-2023/common"
	"github.com/ravenbluedragon/aoc-2023/common/parse"
)

// boilerplate to load and solve puzzles
func main() {
	common.ChoosePart(5, solve1, solve2)
}

func solve1(filename string) any {
	lines := common.LoadData(filename)
	almanac := parseAlmanac(lines)
	minLocation := math.MaxInt
	for _, seed := range almanac.seeds {
		location := almanac.convert(seed)
		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}

func solve2(filename string) any {
	lines := common.LoadData(filename)
	almanac := parseAlmanac(lines)
	seedRanges := almanac.seedRanges()
	locationRanges := almanac.convertRanges(seedRanges)
	minLocation := math.MaxInt
	for _, r := range locationRanges {
		if r.start < minLocation {
			minLocation = r.start
		}
	}
	return minLocation
}

// type intRange is a range of ints
type intRange struct {
	start  int
	end    int
	length int
}

// newIntRange creates a new intRange
func newIntRange(start, length int) intRange {
	return intRange{start: start, end: start + length - 1, length: length}
}

// type rangeSplit splits a range into ranges based on intersection
type rangeSplit struct {
	intersecting    []intRange
	nonIntersecting []intRange
}

// intRange.split splits a range into ranges based on intersection
func (r intRange) split(other intRange) rangeSplit {
	var split rangeSplit
	if r.end < other.start || other.end < r.start {
		split.nonIntersecting = append(split.nonIntersecting, r)
		return split
	}
	a := r.start
	b := r.end
	if r.start < other.start {
		a = other.start
		split.nonIntersecting = append(split.nonIntersecting, newIntRange(r.start, other.start-r.start))
	}
	if r.end > other.end {
		b = other.end
		split.nonIntersecting = append(split.nonIntersecting, newIntRange(other.end+1, r.end-other.end))
	}
	split.intersecting = append(split.intersecting, newIntRange(a, b-a+1))
	return split
}

// type almanac has a list of seeds and a list of conversion maps
// note that maps are in correct order for conversions
type almanac struct {
	seeds []int
	maps  []rangeMap
}

// almanac.convert takes a seed and converts it to a location
func (a almanac) convert(seed int) int {
	value := seed
	for _, m := range a.maps {
		value = m.convert(value)
	}
	return value
}

// almanac.seedRanges returns a list of seed ranges
func (a almanac) seedRanges() []intRange {
	var ranges []intRange
	for i := 0; i < len(a.seeds); i += 2 {
		ranges = append(ranges, newIntRange(a.seeds[i], a.seeds[i+1]))
	}
	return ranges
}

// almanac.convertRanges converts ranges of seeds to ranges of locations
func (a almanac) convertRanges(r []intRange) []intRange {
	for _, m := range a.maps {
		r = m.convertRanges(r)
	}
	return r
}

// type rangeMap is a list of linear maps, with a source and destination types
type rangeMap struct {
	source      string
	destination string
	maps        []linearMap
}

// rangeMap.convert takes a source and converts it to a destination value
func (r rangeMap) convert(value int) int {
	for _, m := range r.maps {
		if val, ok := m.convert(value); ok {
			return val
		}
	}
	return value
}

// rangeMap.convertRanges takes ranges of sources and converts them to ranges of destinations
func (r rangeMap) convertRanges(src []intRange) []intRange {
	var dest []intRange
	for _, lm := range r.maps {
		var fragments []intRange
		for _, s := range src {
			split := s.split(lm.source)
			fragments = append(fragments, split.nonIntersecting...)
			for _, i := range split.intersecting {
				dest = append(dest, lm.convertRange(i))
			}
		}
		src = fragments
	}
	dest = append(dest, src...)
	return dest
}

// type linearMap has a source and destination and length
type linearMap struct {
	source      intRange
	destination int
}

// newLinearMap creates a new linear map
func newLinearMap(source, destination, length int) linearMap {
	src := newIntRange(source, length)
	return linearMap{source: src, destination: destination}
}

// linearMap.convert takes a source and converts it to a destination value
func (l linearMap) convert(value int) (int, bool) {
	if l.source.start <= value && value <= l.source.end {
		offset := value - l.source.start
		return l.destination + offset, true
	}
	return value, false
}

// linearMap.convertRange takes a range of sources and converts it to a range of destinations
func (l linearMap) convertRange(r intRange) intRange {
	start, ok := l.convert(r.start)
	if !ok {
		log.Fatalf("Error converting range: %v with map %v", r, l)
	}
	return newIntRange(start, r.length)
}

// parseAlmanac parses the almanac data into a struct
func parseAlmanac(lines []string) almanac {
	var almanac almanac

	// NOTE: int slice starts at char 7 ("seeds: ")
	seeds, err := parse.StringToIntList(lines[0][7:], " ")
	if err != nil {
		log.Fatalf("Error parsing seeds: %s", err)
	}
	almanac.seeds = seeds

	// NOTE: maps start on line 3
	var currentMap rangeMap
	for _, line := range lines[2:] {
		if line == "" {
			almanac.maps = append(almanac.maps, currentMap)
			currentMap = rangeMap{}
		} else if strings.HasSuffix(line, "map:") {
			keys := strings.Split(strings.Split(line, " ")[0], "-")
			currentMap.source = keys[0]
			currentMap.destination = keys[2]
		} else {
			ints, err := parse.StringToIntList(line, " ")
			if err != nil || len(ints) != 3 {
				log.Fatalf("Error parsing map: %s, %s", line, err)
			}
			// NOTE: layout is dest src len
			lm := newLinearMap(ints[1], ints[0], ints[2])
			currentMap.maps = append(currentMap.maps, lm)
		}
	}
	almanac.maps = append(almanac.maps, currentMap)
	return almanac
}
