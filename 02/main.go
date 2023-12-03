package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/ravenbluedragon/aoc-2023/common"
)

// boilerplate to load and solve puzzles
func main() {
	common.ChoosePart(2, solve1, solve2)
}

type sample struct {
	red   int
	green int
	blue  int
}

type game struct {
	id      int
	samples []sample
}

func parse(line string) game {
	game := game{}
	parts := strings.Split(line, ": ")
	if id, err := strconv.Atoi(parts[0][5:]); err != nil {
		log.Fatalf("Unable to parse id: %s", parts[0][5:])
	} else {
		game.id = id
	}

	for _, draw := range strings.Split(parts[1], "; ") {
		var sample sample
		for _, pair := range strings.Split(draw, ", ") {
			parts := strings.Split(pair, " ")
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalf("Unable to parse count: %s", parts[0])
			}
			switch parts[1] {
			case "red":
				sample.red = count
			case "green":
				sample.green = count
			case "blue":
				sample.blue = count
			}
		}
		game.samples = append(game.samples, sample)
	}

	return game
}

func contains(bag sample, sample sample) bool {
	return bag.red >= sample.red &&
		bag.green >= sample.green &&
		bag.blue >= sample.blue
}

func possible(game game, bag sample) bool {
	for _, sample := range game.samples {
		if !contains(bag, sample) {
			return false
		}
	}
	return true
}

func solve1(filename string) any {
	data := common.LoadData(filename)
	bag := sample{12, 13, 14}
	total := 0
	for _, line := range data {
		game := parse(line)
		if possible(game, bag) {
			total += game.id
		}
	}
	return total
}

func minimum(samples []sample) sample {
	min := samples[0]
	for _, sample := range samples {
		if sample.red > min.red {
			min.red = sample.red
		}
		if sample.green > min.green {
			min.green = sample.green
		}
		if sample.blue > min.blue {
			min.blue = sample.blue
		}
	}
	return min
}

func power(sample sample) int {
	return sample.red * sample.green * sample.blue
}

func solve2(filename string) any {
	data := common.LoadData(filename)
	sum := 0
	for _, line := range data {
		game := parse(line)
		sum += power(minimum(game.samples))
	}
	return sum
}
