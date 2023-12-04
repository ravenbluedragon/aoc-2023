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
	common.ChoosePart(4, solve1, solve2)
}

func solve1(filename string) any {
	data := common.LoadData(filename)
	sum := 0
	for _, line := range data {
		c := parseCard(line)
		sum += c.score1()
	}
	return sum
}

func solve2(filename string) any {
	data := common.LoadData(filename)
	copies := make([]int, len(data))
	for i := range copies {
		copies[i] = 1
	}
	sum := 0
	for i, line := range data {
		c := parseCard(line)
		m := c.matches()
		p := copies[i]
		sum += p
		for j := 1; j <= m; j++ {
			copies[i+j] += p
		}
	}
	return sum
}

type card struct {
	id      int
	winning []int
	player  []int
}

// parseCard takes a string and returns a card
// line format: Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
func parseCard(line string) card {
	c := card{}
	var err error
	split1 := strings.Split(line, ": ")
	c.id, err = strconv.Atoi(strings.TrimSpace(split1[0][5:]))
	if err != nil {
		log.Fatalf("Error converting Card %s to int: %s", split1[0], err)
	}
	split2 := strings.Split(split1[1], " | ")
	if c.winning, err = parse.StringToIntList(split2[0], " "); err != nil {
		log.Fatalf("Error converting winning %s to int: %s", split2[0], err)
	}
	if c.player, err = parse.StringToIntList(split2[1], " "); err != nil {
		log.Fatalf("Error converting player %s to int: %s", split2[1], err)
	}

	return c
}

// matches returns the number of matches between the winning and player values
func (c *card) matches() int {
	matches := 0
	for _, w := range c.winning {
		for _, p := range c.player {
			if w == p {
				matches++
			}
		}
	}
	return matches
}

// score1 returns the score of the winning player as 2^matches
func (c *card) score1() int {
	matches := c.matches()
	if matches == 0 {
		return 0
	}
	return 1 << (matches - 1)
}
