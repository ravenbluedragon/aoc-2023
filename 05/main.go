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
	data := common.LoadData(filename)
	almanac := parseAlmanac(data)
	min := math.MaxInt
	for _, seed := range almanac.seeds {
		value := seed
		for _, conversion := range almanac.conversions {
			value = conversion.convert(value)
		}
		if value < min {
			min = value
		}
	}

	return min
}

func solve2(filename string) any {
	data := common.LoadData(filename)
	almanac := parseAlmanac(data)
	min := math.MaxInt
	var ranges []struct {
		start int
		len   int
	}
	for i := 0; i < len(almanac.seeds); i += 2 {
		ranges = append(ranges, struct {
			start int
			len   int
		}{start: almanac.seeds[i], len: almanac.seeds[i+1]})
	}
	for _, r := range ranges {
		for seed := r.start; seed <= r.start+r.len; seed++ {
			value := seed
			for _, conversion := range almanac.conversions {
				value = conversion.convert(value)
			}
			if value < min {
				min = value
			}
		}
	}

	return min
}

type rangeMap struct {
	src int
	dst int
	len int
}

type conversionTable struct {
	source string
	dest   string
	maps   []rangeMap
}

func (c conversionTable) convert(value int) int {
	for _, r := range c.maps {
		if r.src <= value && value <= r.src+r.len {
			return r.dst + (value - r.src)
		}
	}
	return value
}

type almanac struct {
	seeds       []int
	conversions []conversionTable
}

func parseAlmanac(lines []string) almanac {
	almanac := almanac{}
	var err error
	almanac.seeds, err = parse.StringToIntList(strings.Split(lines[0], ": ")[1], " ")
	if err != nil {
		log.Fatalf("Error parsing seeds: %s", err)
	}
	var conversion conversionTable
	for _, line := range lines[2:] {
		if line == "" {
			almanac.conversions = append(almanac.conversions, conversion)
			conversion = conversionTable{}
			continue
		}
		if strings.HasSuffix(line, "map:") {
			parts := strings.Split(line, " ")
			keys := strings.Split(parts[0], "-")
			conversion.source = keys[0]
			conversion.dest = keys[2]
			continue
		}
		ints, err := parse.StringToIntList(line, " ")
		if err != nil || len(ints) != 3 {
			log.Fatalf("Error parsing map: %s", err)
		}
		conversion.maps = append(conversion.maps, rangeMap{dst: ints[0], src: ints[1], len: ints[2]})
	}
	almanac.conversions = append(almanac.conversions, conversion)

	return almanac
}
