package main

import (
	"log"

	"github.com/ravenbluedragon/aoc-2023/02/game"
	"github.com/ravenbluedragon/aoc-2023/common"
)

// boilerplate to load and solve puzzles
func main() {
	common.ChoosePart(2, solve1, solve2)
}

func solve1(filename string) any {
	data := common.LoadData(filename)
	bag := game.NewSample(12, 13, 14)
	sum := 0
	for _, line := range data {
		game, err := game.ParseGame(line)
		if err != nil {
			log.Fatal(err)
		}
		if game.Possible(bag) {
			sum += game.Id
		}
	}
	return sum
}

func solve2(filename string) any {
	data := common.LoadData(filename)
	sum := 0
	for _, line := range data {
		game, err := game.ParseGame(line)
		if err != nil {
			log.Fatal(err)
		}
		min := game.Minimum()
		sum += min.Power()
	}
	return sum
}
