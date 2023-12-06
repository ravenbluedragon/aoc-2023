package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/ravenbluedragon/aoc-2023/common"
	"github.com/ravenbluedragon/aoc-2023/common/parse"
)

// boilerplate to load and solve puzzles
func main() {
	common.ChoosePart(6, solve1, solve2)
}

func solve1(filename string) any {
	lines := common.LoadData(filename)
	races := parseRaces1(lines)
	prod := 1
	for _, race := range races {
		prod *= race.waysToWin()
	}
	return prod
}

func solve2(filename string) any {
	lines := common.LoadData(filename)
	race := parseRace2(lines)
	return race.waysToWin()
}

// type race holds a time and distance
type race struct {
	time     int
	distance int
}

// parseRaces1 converts an input to a slice of races
func parseRaces1(lines []string) []race {
	times, err := parse.StringToIntList(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ")
	if err != nil {
		log.Fatalf("failed to parse times: %s", err)
	}
	distances, err := parse.StringToIntList(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " ")
	if err != nil {
		log.Fatalf("failed to parse distances: %s", err)
	}
	var races []race
	if len(times) != len(distances) {
		log.Fatalf("times and distances are not the same length")
	}
	for i := range times {
		races = append(races, race{times[i], distances[i]})
	}
	return races
}

// parseRace2 converts an input to a race by eliminating spaces
func parseRace2(lines []string) race {
	var race race
	var err error

	// split the line at the colon, then remove all spaces
	timeStr := strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")
	if race.time, err = strconv.Atoi(timeStr); err != nil {
		log.Fatalf("failed to parse time: %s", err)
	}

	distanceStr := strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")
	if race.distance, err = strconv.Atoi(distanceStr); err != nil {
		log.Fatalf("failed to parse distance: %s", err)
	}

	return race
}

// race.waysToWin calculates the ways to win for a given time and distance to beat
func (r race) waysToWin() int {
	ways := 0
	for speed := 1; speed < r.time; speed++ {
		if r.distance < speed*(r.time-speed) {
			ways++
		}
	}
	return ways
}
