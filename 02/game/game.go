package game

import (
	"fmt"
	"strconv"
	"strings"
)

// sample represents a sample of coloured objects in a bag
type sample struct {
	red   int
	green int
	blue  int
}

// NewSample creates a new sample with the given red, green, and blue values
func NewSample(red, green, blue int) sample {
	return sample{red, green, blue}
}

// Contains returns true if the bag contains at least as many of each color as the sample
func (bag sample) Contains(sample sample) bool {
	return bag.red >= sample.red &&
		bag.green >= sample.green &&
		bag.blue >= sample.blue
}

// Power returns the product of the red, green, and blue values
func (s *sample) Power() int {
	return s.red * s.green * s.blue
}

// game represents a round of taking samples from a bag
type game struct {
	Id      int
	samples []sample
}

// ParseGame parses a line of input into a game
// The line should be in the format "Game #id: #red red, #green green, #blue blue; ..."
func ParseGame(line string) (game, error) {
	game := game{}
	parts := strings.Split(line, ": ")
	if id, err := strconv.Atoi(parts[0][5:]); err != nil {
		return game, fmt.Errorf("unable to parse id: %s", parts[0][5:])
	} else {
		game.Id = id
	}

	for _, draw := range strings.Split(parts[1], "; ") {
		var sample sample
		for _, pair := range strings.Split(draw, ", ") {
			parts := strings.Split(pair, " ")
			count, err := strconv.Atoi(parts[0])
			if err != nil {
				return game, fmt.Errorf("unable to parse count: %s", parts[0])
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

	return game, nil
}

// Possible returns true if the bag could contain all the samples
func (g *game) Possible(bag sample) bool {
	for _, sample := range g.samples {
		if !bag.Contains(sample) {
			return false
		}
	}
	return true
}

// Minimum returns the minimum number of each colour needed to contain all the samples
func (g *game) Minimum() sample {
	min := g.samples[0]
	for _, sample := range g.samples {
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
